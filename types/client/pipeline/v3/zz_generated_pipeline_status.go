package client

const (
	PipelineStatusType               = "pipelineStatus"
	PipelineStatusFieldCommitInfo    = "commitInfo"
	PipelineStatusFieldLastRunId     = "lastRunId"
	PipelineStatusFieldLastRunStatus = "lastRunStatus"
	PipelineStatusFieldLastRunTime   = "lastRunTime"
	PipelineStatusFieldNextRunTime   = "nextRunTime"
	PipelineStatusFieldRunCount      = "runCount"
	PipelineStatusFieldWebHookId     = "webhookId"
	PipelineStatusFieldWebHookToken  = "webhookToken"
)

type PipelineStatus struct {
	CommitInfo    string `json:"commitInfo,omitempty"`
	LastRunId     string `json:"lastRunId,omitempty"`
	LastRunStatus string `json:"lastRunStatus,omitempty"`
	LastRunTime   *int64 `json:"lastRunTime,omitempty"`
	NextRunTime   *int64 `json:"nextRunTime,omitempty"`
	RunCount      *int64 `json:"runCount,omitempty"`
	WebHookId     *int64 `json:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty"`
}
