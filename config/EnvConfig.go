package config

// EnvConfig represents the MC configuration
type EnvConfig struct {
	APIEndpoint string   `json:"api_endpoint"`
	Login       string   `json:"login"`
	Password    string   `json:"password"`
	Org         string   `json:"org"`
	Space       string   `json:"space"`
	Services    []string `json:"services"`
	DBURI       string   `json:"db_uri"`
}

// GetCredentials returns PaaS credentails
func (c *EnvConfig) GetCredentials() (login string, pass string) {
	login = c.Login
	pass = c.Password
	return
}

// GetFullDetails return all suitable PaaS details to find the services
func (c *EnvConfig) GetFullDetails() (org string, space string, services []string) {
	org = c.Org
	space = c.Space
	services = c.Services
	return
}

// GetServiceList returns services list
func (c *EnvConfig) GetServiceList() []string {
	return c.Services
}
