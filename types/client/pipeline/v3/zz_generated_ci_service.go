package client

const (
	CIServiceType               = "ciService"
	CIServiceFieldCommand       = "command"
	CIServiceFieldContainerName = "containerName"
	CIServiceFieldEntrypoint    = "entrypoint"
	CIServiceFieldImage         = "image"
	CIServiceFieldName          = "name"
)

type CIService struct {
	Command       string `json:"command,omitempty"`
	ContainerName string `json:"containerName,omitempty"`
	Entrypoint    string `json:"entrypoint,omitempty"`
	Image         string `json:"image,omitempty"`
	Name          string `json:"name,omitempty"`
}
