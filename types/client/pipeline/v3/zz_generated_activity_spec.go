package client

const (
	ActivitySpecType                 = "activitySpec"
	ActivitySpecFieldActivityStages  = "activity_stages"
	ActivitySpecFieldCommitInfo      = "commitInfo"
	ActivitySpecFieldEnvVars         = "envVars"
	ActivitySpecFieldFailMessage     = "failMessage"
	ActivitySpecFieldNodeName        = "nodename"
	ActivitySpecFieldPendingStage    = "pendingStage"
	ActivitySpecFieldPipeline        = "pipelineSource"
	ActivitySpecFieldPipelineName    = "pipelineName"
	ActivitySpecFieldPipelineVersion = "pipelineVersion"
	ActivitySpecFieldRunSequence     = "runSequence"
	ActivitySpecFieldStartTS         = "start_ts"
	ActivitySpecFieldStopTS          = "stop_ts"
	ActivitySpecFieldTriggerType     = "triggerType"
)

type ActivitySpec struct {
	ActivityStages  []ActivityStage   `json:"activity_stages,omitempty"`
	CommitInfo      string            `json:"commitInfo,omitempty"`
	EnvVars         map[string]string `json:"envVars,omitempty"`
	FailMessage     string            `json:"failMessage,omitempty"`
	NodeName        string            `json:"nodename,omitempty"`
	PendingStage    *int64            `json:"pendingStage,omitempty"`
	Pipeline        *Pipeline         `json:"pipelineSource,omitempty"`
	PipelineName    string            `json:"pipelineName,omitempty"`
	PipelineVersion string            `json:"pipelineVersion,omitempty"`
	RunSequence     *int64            `json:"runSequence,omitempty"`
	StartTS         *int64            `json:"start_ts,omitempty"`
	StopTS          *int64            `json:"stop_ts,omitempty"`
	TriggerType     string            `json:"triggerType,omitempty"`
}
