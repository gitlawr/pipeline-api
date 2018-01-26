package client

const (
	PipelineHistoryStatusType             = "pipelineHistoryStatus"
	PipelineHistoryStatusFieldCommitInfo  = "commitInfo"
	PipelineHistoryStatusFieldEndTime     = "endTime"
	PipelineHistoryStatusFieldEnvVars     = "envVars"
	PipelineHistoryStatusFieldStageStatus = "stageStatus"
	PipelineHistoryStatusFieldStartTime   = "startTime"
	PipelineHistoryStatusFieldStatus      = "status"
)

type PipelineHistoryStatus struct {
	CommitInfo  string            `json:"commitInfo,omitempty"`
	EndTime     *int64            `json:"endTime,omitempty"`
	EnvVars     map[string]string `json:"envVars,omitempty"`
	StageStatus []StageStatus     `json:"stageStatus,omitempty"`
	StartTime   *int64            `json:"startTime,omitempty"`
	Status      string            `json:"status,omitempty"`
}
