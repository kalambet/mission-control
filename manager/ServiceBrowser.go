package manager

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/kalambet/mission-control/config"
	"github.com/kalambet/mission-control/services"
	"github.com/kalambet/mission-control/types"
)

// ServiceBrowser logs into Bluemix
// gets the token and finds ids for
// wanted services
type ServiceBrowser struct {
	serviceConfig *config.EnvConfig
	token         string
}

// GetBearerToken logins into Bluemix and returns bearer token
func (browser *ServiceBrowser) getBearerToken() (string, error) {
	client := http.Client{}
	login, password := browser.serviceConfig.GetCredentials()

	loginForm := url.Values{}
	loginForm.Set("grant_type", "password")
	loginForm.Add("scope", " ")
	loginForm.Add("username", login)
	loginForm.Add("password", password)

	// Request bearer token to access PaaS resource
	req, err := http.NewRequest(
		"POST",
		"https://login.ng.bluemix.net/UAALoginServerWAR/oauth/token",
		bytes.NewBufferString(loginForm.Encode()))
	if err != nil {
		return "", err
	}

	encodedToken := base64.StdEncoding.EncodeToString([]byte("cf:"))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+encodedToken)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(res.Body)

	var loginData types.LoginData
	err = decoder.Decode(&loginData)

	token := loginData.AccessToken
	fmt.Println("Token is : " + token)

	return token, nil
}

// GetServices searches the services by name in PaaS
func (browser *ServiceBrowser) GetServices() ([]*services.Service, error) {
	// 0. Get bearer token
	//token, err := browser.getBearerToken()
	//browser.token = token

	// 1. Serach for organization ID
	spacesURL, err := browser.getSpacesURLByOrgName()
	if err != nil {
		return nil, err
	}
	fmt.Println("Spaces URL: " + spacesURL)

	// 2. Search for space ID
	appsURL, err := browser.getAppsURLBySpacesURL(spacesURL)
	if err != nil {
		return nil, err
	}
	fmt.Println("Apps URL: " + appsURL)

	// 3. Search for services
	serviceList, err := browser.getServicesByName(appsURL)
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}

func (browser *ServiceBrowser) getSpacesURLByOrgName() (string, error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+"/v2/organizations?q=name:"+url.QueryEscape(browser.serviceConfig.Org),
		nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("[%s] on getting organization info", res.Status)
	}

	decoder := json.NewDecoder(res.Body)

	var searchResults types.OrgSearchRes
	err = decoder.Decode(&searchResults)

	if searchResults.TotalResults != 1 {
		return "", errors.New("There are more than one organization with this name pattern")
	}

	return searchResults.Resources[0].Entity.SpacesURL, nil
}

func (browser *ServiceBrowser) getAppsURLBySpacesURL(spacesURL string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+spacesURL+"?q=name:"+url.QueryEscape(browser.serviceConfig.Space),
		nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("[%s] on getting application info", res.Status)
	}

	decoder := json.NewDecoder(res.Body)

	var searchResults types.SpaceSearchRes
	err = decoder.Decode(&searchResults)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if searchResults.TotalResults != 1 {
		return "", errors.New("There are more than one space with this name pattern")
	}

	return searchResults.Resources[0].Entity.AppsURL, nil
}

func (browser *ServiceBrowser) getServicesByName(appsURL string) ([]*services.Service, error) {
	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+appsURL,
		nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("[%s] on getting service details", res.Status)
	}

	decoder := json.NewDecoder(res.Body)

	var searchResults types.AppsSearchRes
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

// CollectServiceStatus collects status of the service
// and than stores it in the persistant storage
func (browser *ServiceBrowser) CollectServiceStatus(service *services.Service) (*services.ServiceStatus, error) {
	fmt.Printf("Checking service %+v\n", *service)

	req, err := http.NewRequest(
		"GET",
		browser.serviceConfig.APIEndpoint+service.URL+"/stats",
		nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+browser.token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("[%s] on getting State of application", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()

	var serviceStates map[string]*types.InstanceState
	err = decoder.Decode(&serviceStates)
	if err != nil {
		return nil, err
	}

	// We really do nedd map in our instances state list
	var instancesStateList = make([]types.InstanceState, len(serviceStates))
	idx := 0
	for _, state := range serviceStates {
		instancesStateList[idx] = *state
		idx++
	}

	// Prepare Service Status
	var status = services.ServiceStatus{
		Name:         service.Name,
		Organization: browser.serviceConfig.Org,
		Space:        browser.serviceConfig.Space,
		Instances:    instancesStateList,
		UpdateTime:   time.Now()}

	return &status, nil
}
