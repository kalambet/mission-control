package manager

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kalambet/mission-control/config"
	"github.com/kalambet/mission-control/data"
	"github.com/kalambet/mission-control/services"
)

// ServiceBrowser logs into Bluemix
// gets the token and finds ids for
// wanted services
type ServiceBrowser struct {
	serviceConfig *config.EnvConfig
	token         string
}

// GetBearerToken logins into Bluemix and returns bearer token
func (browser *ServiceBrowser) getBearerToken() (token string, err error) {
	client := http.Client{}
	login, password := browser.serviceConfig.GetCredentials()

	loginForm := url.Values{}
	loginForm.Set("grant_type", "password")
	loginForm.Add("scope", " ")
	loginForm.Add("username", login)
	loginForm.Add("password", password)

	fmt.Println("Login Form: ")
	fmt.Print(loginForm)

	// Request bearer token to access PaaS resource
	req, err := http.NewRequest(
		"POST",
		"https://login.ng.bluemix.net/UAALoginServerWAR/oauth/token",
		bytes.NewBufferString(loginForm.Encode()))

	encodedToken := base64.StdEncoding.EncodeToString([]byte("cf:"))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+encodedToken)

	res, err := client.Do(req)
	decoder := json.NewDecoder(res.Body)

	var data data.LoginData
	err = decoder.Decode(&data)

	token = data.AccessToken
	fmt.Println("Token is : " + token)

	return
}

// GetServices searches the services by name in PaaS
func (browser *ServiceBrowser) GetServices() (serviceList []*services.Service, err error) {
	// 0. Get bearer token
	//token, err := browser.getBearerToken()

	//browser.token = token

	// 1. Serach for organization ID
	spacesURL, err := browser.getSpacesURLByOrgName()
	fmt.Println("Spaces URL: " + spacesURL)
	if err != nil {
		return nil, err
	}

	// 2. Search for space ID
	appsURL, err := browser.getAppsURLBySpacesURL(spacesURL)
	fmt.Println("Apps URL: " + appsURL)
	if err != nil {
		return nil, err
	}

	// 3. Search for services
	serviceList, err = browser.getServicesByName(appsURL)
	if err != nil {
		return nil, err
	}

	return
}

func (browser *ServiceBrowser) getSpacesURLByOrgName() (_url string, err error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+"/v2/organizations?q=name:"+url.QueryEscape(browser.serviceConfig.Org),
		nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	decoder := json.NewDecoder(res.Body)

	var searchResults data.OrgSearchRes
	err = decoder.Decode(&searchResults)

	if searchResults.TotalResults != 1 {
		err = errors.New("There are more than one organization with this name pattern")
		fmt.Println(err)
		return "", err
	}

	_url = searchResults.Resources[0].Entity.SpacesURL
	return
}

func (browser *ServiceBrowser) getAppsURLBySpacesURL(spacesURL string) (_url string, err error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+spacesURL+"?q=name:"+url.QueryEscape(browser.serviceConfig.Space),
		nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	decoder := json.NewDecoder(res.Body)

	var searchResults data.SpaceSearchRes
	err = decoder.Decode(&searchResults)

	if searchResults.TotalResults != 1 {
		err = errors.New("There are more than one space with this name pattern")
		fmt.Println(err)
		return "", err
	}

	_url = searchResults.Resources[0].Entity.AppsURL
	return
}

func (browser *ServiceBrowser) getServicesByName(appsURL string) (serviceList []*services.Service, err error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+appsURL,
		nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)

	var searchResults data.AppsSearchRes
	err = decoder.Decode(&searchResults)
	result := make([]*services.Service, len(browser.serviceConfig.Services))

	for _, appResource := range searchResults.Resources {
		for idx, serviceName := range browser.serviceConfig.Services {
			if appResource.Entity.Name == serviceName {
				result[idx] = &services.Service{
					GUID: appResource.Metadata.GUID,
					Name: serviceName,
					URL:  appResource.Metadata.URL}
			}
		}
	}
	return result, nil
}
