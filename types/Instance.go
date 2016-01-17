package types

// InstanceUsage describes instance usage
type InstanceUsage struct {
	Time   string  `json:"time"`
	CPU    float32 `json:"cpu"`
	Memory int64   `json:"mem"`
	Disk   int64   `json:"disk"`
}

// InstanceStats describes the instance statictics
type InstanceStats struct {
	Name      string        `json:"name"`
	URIs      []string      `json:"uirs"`
	Host      string        `json:"host"`
	Port      int32         `json:"port"`
	Uptime    int64         `json:"uptime"`
	MemQuota  int64         `json:"mem_quota"`
	DiskQuota int64         `json:"disk_quota"`
	FDSQuota  int64         `json:"fds_quota"`
	Usage     InstanceUsage `json:"usage"`
}

// InstanceState describes main structure of instance current state
type InstanceState struct {
	State string        `json:"state"`
	Stats InstanceStats `json:"stats"`
}
