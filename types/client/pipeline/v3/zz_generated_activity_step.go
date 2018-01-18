package client

const (
	ActivityStepType          = "activityStep"
	ActivityStepFieldDuration = "duration"
	ActivityStepFieldMessage  = "message"
	ActivityStepFieldName     = "name"
	ActivityStepFieldStartTS  = "start_ts"
	ActivityStepFieldStatus   = "status"
)

type ActivityStep struct {
	Duration *int64 `json:"duration,omitempty"`
	Message  string `json:"message,omitempty"`
	Name     string `json:"name,omitempty"`
	StartTS  *int64 `json:"start_ts,omitempty"`
	Status   string `json:"status,omitempty"`
}
