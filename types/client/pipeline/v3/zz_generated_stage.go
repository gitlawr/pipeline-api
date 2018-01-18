package client

const (
	StageType             = "stage"
	StageFieldApprovers   = "approvers"
	StageFieldCondition   = "condition"
	StageFieldConditions  = "conditions"
	StageFieldName        = "name"
	StageFieldNeedApprove = "needApprove"
	StageFieldParallel    = "parallel"
	StageFieldSteps       = "steps"
)

type Stage struct {
	Approvers   []string            `json:"approvers,omitempty"`
	Condition   string              `json:"condition,omitempty"`
	Conditions  *PipelineConditions `json:"conditions,omitempty"`
	Name        string              `json:"name,omitempty"`
	NeedApprove *bool               `json:"needApprove,omitempty"`
	Parallel    *bool               `json:"parallel,omitempty"`
	Steps       []Step              `json:"steps,omitempty"`
}
