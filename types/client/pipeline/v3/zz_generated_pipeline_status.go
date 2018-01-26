package client

const (
	PipelineStatusType               = "pipelineStatus"
	PipelineStatusFieldLastRunId     = "lastRunId"
	PipelineStatusFieldLastRunStatus = "lastRunStatus"
	PipelineStatusFieldLastRunTime   = "lastRunTime"
	PipelineStatusFieldNextRunNumber = "nextRunNumber"
	PipelineStatusFieldNextRunTime   = "nextRunTime"
	PipelineStatusFieldWebHookId     = "webhookId"
	PipelineStatusFieldWebHookToken  = "webhookToken"
)

type PipelineStatus struct {
	LastRunId     string `json:"lastRunId,omitempty"`
	LastRunStatus string `json:"lastRunStatus,omitempty"`
	LastRunTime   *int64 `json:"lastRunTime,omitempty"`
	NextRunNumber *int64 `json:"nextRunNumber,omitempty"`
	NextRunTime   *int64 `json:"nextRunTime,omitempty"`
	WebHookId     string `json:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty"`
}
