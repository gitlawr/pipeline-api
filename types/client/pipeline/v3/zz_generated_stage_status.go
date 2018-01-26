package client

const (
	StageStatusType            = "stageStatus"
	StageStatusFieldEndTime    = "endTime"
	StageStatusFieldStartTime  = "startTime"
	StageStatusFieldStatus     = "status"
	StageStatusFieldStepStatus = "stepStatus"
)

type StageStatus struct {
	EndTime    *int64       `json:"endTime,omitempty"`
	StartTime  *int64       `json:"startTime,omitempty"`
	Status     string       `json:"status,omitempty"`
	StepStatus []StepStatus `json:"stepStatus,omitempty"`
}
