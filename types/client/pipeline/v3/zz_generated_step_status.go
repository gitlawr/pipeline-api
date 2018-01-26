package client

const (
	StepStatusType           = "stepStatus"
	StepStatusFieldEndTime   = "endTime"
	StepStatusFieldStartTime = "startTime"
	StepStatusFieldStatus    = "status"
)

type StepStatus struct {
	EndTime   *int64 `json:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	Status    string `json:"status,omitempty"`
}
