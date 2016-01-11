package data

import "encoding/json"

// AppsEntity inner AppsResource type
type AppsEntity struct {
	Name                    string          `json:"name"`
	Production              bool            `json:"production"`
	SpaceGUID               string          `json:"space_guid"`
	StackGUID               string          `json:"stack_guid"`
	Buildpack               string          `json:"buildpack"`
	DetectedBuildpack       string          `json:"detected_buildpack"`
	AppEnv                  json.RawMessage `json:"environment_json"`
	Memory                  int             `json:"memory"`
	Instances               int             `json:"instances"`
	DiskQuota               int             `json:"disk_quota"`
	State                   string          `json:"state"`
	Version                 string          `json:"version"`
	Command                 string          `json:"command"`
	Console                 bool            `json:"console"`
	Debug                   string          `json:"debug"`
	StagingTaskID           string          `json:"staging_task_id"`
	PackageState            string          `json:"package_state"`
	HealthCheckType         string          `json:"health_check_type"`
	HealthCheckTimeout      string          `json:"health_check_timeout"`
	StagingFailedReason     string          `json:"staging_failed_reason"`
	StagingFailsDescription string          `json:"staging_failed_description"`
	Diego                   string          `json:"diego"`
	DockerImage             string          `json:"docker_image"`
	PackageUpdatedAt        string          `json:"package_updated_at"`
	DetectedStartCommand    string          `json:"detected_start_command"`
	EnableSSH               bool            `json:"enable_ssh"`
	DockerCreds             json.RawMessage `json:"docker_credentials_json"`
	SpacesURL               string          `json:"spaces_url"`
	StackURL                string          `json:"stack_url"`
	EventsURL               string          `json:"events_url"`
	ServiceBindingsURL      string          `json:"service_bindings_url"`
	RoutesURL               string          `json:"routes_url"`
}

// AppsResource inner AppsSearchRes type
type AppsResource struct {
	Metadata Metadata
	Entity   AppsEntity
}

// AppsSearchRes represents the organization search results
type AppsSearchRes struct {
	TotalResults int            `json:"total_results"`
	TotalPages   int            `json:"total_pages"`
	PrevURL      string         `json:"prev_url"`
	NextURL      string         `json:"next_url"`
	Resources    []AppsResource `json:"resources"`
}

// AppCurrentState contains information about app instances
type AppCurrentState struct {
	Instances []InstanceState
}
