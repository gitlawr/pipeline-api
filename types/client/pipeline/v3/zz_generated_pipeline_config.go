package client

const (
	PipelineConfigType              = "pipelineConfig"
	PipelineConfigFieldGibhubConfig = "githubConfig"
)

type PipelineConfig struct {
	GibhubConfig *GibhubConfig `json:"githubConfig,omitempty"`
}
