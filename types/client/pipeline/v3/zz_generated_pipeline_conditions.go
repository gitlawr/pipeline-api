package client

const (
	PipelineConditionsType     = "pipelineConditions"
	PipelineConditionsFieldAll = "all"
	PipelineConditionsFieldAny = "any"
)

type PipelineConditions struct {
	All []string `json:"all,omitempty"`
	Any []string `json:"any,omitempty"`
}
