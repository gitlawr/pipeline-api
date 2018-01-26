package client

const (
	GibhubConfigType              = "gibhubConfig"
	GibhubConfigFieldClientId     = "clientId"
	GibhubConfigFieldClientSecret = "clientSecret"
	GibhubConfigFieldHost         = "host"
	GibhubConfigFieldScheme       = "githubConfig"
)

type GibhubConfig struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Host         string `json:"host,omitempty"`
	Scheme       string `json:"githubConfig,omitempty"`
}
