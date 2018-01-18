package client

const (
	PipelineSpecType               = "pipelineSpec"
	PipelineSpecFieldCronTrigger   = "cronTrigger"
	PipelineSpecFieldIsActivate    = "isActivate"
	PipelineSpecFieldKeepWorkspace = "keepWorkspace"
	PipelineSpecFieldParameters    = "parameters"
	PipelineSpecFieldStages        = "stages"
	PipelineSpecFieldTemplates     = "templates"
)

type PipelineSpec struct {
	CronTrigger   *CronTrigger      `json:"cronTrigger,omitempty"`
	IsActivate    *bool             `json:"isActivate,omitempty"`
	KeepWorkspace *bool             `json:"keepWorkspace,omitempty"`
	Parameters    []string          `json:"parameters,omitempty"`
	Stages        []Stage           `json:"stages,omitempty"`
	Templates     map[string]string `json:"templates,omitempty"`
}
