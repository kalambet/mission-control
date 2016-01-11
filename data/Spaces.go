package data

// SpaceEntity inner SpaceResource type
type SpaceEntity struct {
	Name                     string `json:"name"`
	OrganizationGUI          string `json:"organization_guid"`
	SpaceQuotaDefinitionGUID string `json:"quota_definition_guid"`
	AllowSSH                 bool   `json:"allow_ssh"`
	OrganizationURL          string `json:"organization_url"`
	DevelopersURL            string `json:"developers_url"`
	ManagersURL              string `json:"managers_url"`
	AuditorsURL              string `json:"auditors_url"`
	AppsURL                  string `json:"apps_url"`
	RoutesURL                string `json:"routes_url"`
	DomainURL                string `json:"domains_url"`
	ServiceInstancesURL      string `json:"service_instances_url"`
	AppEventsURL             string `json:"app_events_url"`
	EventsURL                string `json:"events_url"`
	SecurityGroupsURL        string `json:"security_groups_url"`
}

// SpaceResource inner SpaceSearchRes type
type SpaceResource struct {
	Metadata Metadata
	Entity   SpaceEntity
}

// SpaceSearchRes represents the organization search results
type SpaceSearchRes struct {
	TotalResults int             `json:"total_results"`
	TotalPages   int             `json:"total_pages"`
	PrevURL      string          `json:"prev_url"`
	NextURL      string          `json:"next_url"`
	Resources    []SpaceResource `json:"resources"`
}
