package client

const (
	PipelineSpecType             = "pipelineSpec"
	PipelineSpecFieldCronTrigger = "cronTrigger"
	PipelineSpecFieldDisplayName = "displayName"
	PipelineSpecFieldStages      = "stages"
)

type PipelineSpec struct {
	CronTrigger *CronTrigger `json:"cronTrigger,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	Stages      []Stage      `json:"stages,omitempty"`
}
