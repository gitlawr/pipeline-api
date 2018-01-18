package client

const (
	ActivityStageType               = "activityStage"
	ActivityStageFieldActivityId    = "activity_id"
	ActivityStageFieldActivitySteps = "activity_steps"
	ActivityStageFieldApprovers     = "approvers"
	ActivityStageFieldDuration      = "duration"
	ActivityStageFieldName          = "name"
	ActivityStageFieldNeedApproval  = "need_approval"
	ActivityStageFieldRawOutput     = "rawOutput"
	ActivityStageFieldStartTS       = "start_ts"
	ActivityStageFieldStatus        = "status"
)

type ActivityStage struct {
	ActivityId    string         `json:"activity_id,omitempty"`
	ActivitySteps []ActivityStep `json:"activity_steps,omitempty"`
	Approvers     []string       `json:"approvers,omitempty"`
	Duration      *int64         `json:"duration,omitempty"`
	Name          string         `json:"name,omitempty"`
	NeedApproval  *bool          `json:"need_approval,omitempty"`
	RawOutput     string         `json:"rawOutput,omitempty"`
	StartTS       *int64         `json:"start_ts,omitempty"`
	Status        string         `json:"status,omitempty"`
}
