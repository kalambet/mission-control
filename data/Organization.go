package data

// OrgEntity inner OrgResource type
type OrgEntity struct {
	Name                    string `json:"name"`
	BillingEnabled          bool   `json:"billing_enabled"`
	QuotaDefinitionGUID     string `json:"quota_definition_guid"`
	Status                  string `json:"status"`
	QuotaDefinitionURL      string `json:"quota_definition_url"`
	SpacesURL               string `json:"spaces_url"`
	DomainURL               string `json:"domains_url"`
	PrivateDomainsURL       string `json:"private_domains_url"`
	UsersURL                string `json:"users_url"`
	ManagersURL             string `json:"managers_url"`
	BiliingManagersURL      string `json:"billing_managers_url"`
	AuditorsURL             string `json:"auditors_url"`
	AppEventsURL            string `json:"app_events_url"`
	SpaceQuotaDefinitialURL string `json:"space_quota_definitions_url"`
}

// OrgResource inner OrgSearchRes type
type OrgResource struct {
	Metadata Metadata
	Entity   OrgEntity
}

// OrgSearchRes represents the organization search results
type OrgSearchRes struct {
	TotalResults int           `json:"total_results"`
	TotalPages   int           `json:"total_pages"`
	PrevURL      string        `json:"prev_url"`
	NextURL      string        `json:"next_url"`
	Resources    []OrgResource `json:"resources"`
}
