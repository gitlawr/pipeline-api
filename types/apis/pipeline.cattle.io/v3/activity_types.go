package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Activity struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec ActivitySpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status ActivityStatus `json:"status"`
}

type ActivityStatus struct {
	TBD string `json:"tbd"`
}

type ActivitySpec struct {
	Pipeline        Pipeline          `json:"pipelineSource,omitempty"`
	PipelineName    string            `json:"pipelineName,omitempty"`
	PipelineVersion string            `json:"pipelineVersion,omitempty"`
	RunSequence     int               `json:"runSequence,omitempty"`
	CommitInfo      string            `json:"commitInfo,omitempty"`
	FailMessage     string            `json:"failMessage,omitempty"`
	PendingStage    int               `json:"pendingStage,omitempty"`
	StartTS         int64             `json:"start_ts,omitempty"`
	StopTS          int64             `json:"stop_ts,omitempty"`
	NodeName        string            `json:"nodename,omitempty"`
	ActivityStages  []ActivityStage   `json:"activity_stages,omitempty"`
	EnvVars         map[string]string `json:"envVars,omitempty"`
	TriggerType     string            `json:"triggerType,omitempty"`
}

type ActivityStage struct {
	ActivityId    string         `json:"activity_id,omitempty"`
	Name          string         `json:"name,omitempty"`
	NeedApproval  bool           `json:"need_approval,omitempty"`
	Approvers     []string       `json:"approvers,omitempty"`
	ActivitySteps []ActivityStep `json:"activity_steps,omitempty"`
	StartTS       int64          `json:"start_ts,omitempty"`
	Duration      int64          `json:"duration,omitempty"`
	Status        string         `json:"status,omitempty"`
	RawOutput     string         `json:"rawOutput,omitempty"`
}

type ActivityStep struct {
	Name     string `json:"name,omitempty"`
	Message  string `json:"message,omitempty"`
	Status   string `json:"status,omitempty"`
	StartTS  int64  `json:"start_ts,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}
