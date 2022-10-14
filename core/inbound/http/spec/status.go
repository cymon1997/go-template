package spec

const (
	StatusOK    InstanceStatus = "OK"
	StatusError InstanceStatus = "error"
)

type InstanceStatus string

type Instance struct {
	Name   string         `json:"name"`
	Status InstanceStatus `json:"status"`
}

type StatusResponse struct {
	Instances []Instance `json:"instances"`
}
