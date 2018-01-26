package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


type PipelineConfig struct {
	GibhubConfig *GibhubConfig `json:"githubConfig,omitempty"`
}

type GibhubConfig struct {
	Scheme string `json:"githubConfig,omitempty"`
	Host string `json:"host,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
}


type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec PipelineSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status PipelineStatus `json:"status"`
}

type PipelineStatus struct {
	//VersionSequence string `json:"-" yaml:"-"`
	NextRunNumber      int    `json:"nextRunNumber" yaml:"nextRunNumber,omitempty"`
	LastRunId     string `json:"lastRunId,omitempty" yaml:"lastRunId,omitempty"`
	LastRunStatus string `json:"lastRunStatus,omitempty" yaml:"lastRunStatus,omitempty"`
	LastRunTime   int64  `json:"lastRunTime,omitempty" yaml:"lastRunTime,omitempty"`
	NextRunTime   int64  `json:"nextRunTime,omitempty" yaml:"nextRunTime,omitempty"`
	WebHookId     string    `json:"webhookId,omitempty" yaml:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty" yaml:"webhookToken,omitempty"`
}

type PipelineSpec struct {
	Id         string `json:"id,omitempty"`
	CronTrigger   CronTrigger `json:"cronTrigger,omitempty" yaml:"cronTrigger,omitempty"`
	Stages        []Stage     `json:"stages,omitempty" yaml:"stages,omitempty"`
}

type CronTrigger struct {
	Timezone        string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
	Spec            string `json:"spec,omitempty" yaml:"spec,omitempty"`
	TriggerOnUpdate bool   `json:"triggerOnUpdate" yaml:"triggerOnUpdate,omitempty"`
}

type Stage struct {
	Name        string             `json:"name,omitempty" yaml:"name,omitempty"`
	Steps       []Step             `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type Step struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	SourceCodeStepConfig *SourceCodeStepConfig
	RunScriptStepConfig *RunScriptStepConfig
	BuildImageStepConfig *BuildImageStepConfig

	//Step timeout in minutes
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

}

type SourceCodeStepConfig struct {
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
	Branch     string `json:"branch,omitempty" yaml:"branch,omitempty"`
	GitUser    string `json:"gitUser,omitempty" yaml:"gitUser,omitempty"`
}

type RunScriptStepConfig struct {
	Image       string      `json:"image,omitempty" yaml:"image,omitempty"`
	ShellScript string      `json:"shellScript,omitempty" yaml:"shellScript,omitempty"`
	Entrypoint  string      `json:"entrypoint,omitempty" yaml:"enrtypoint,omitempty"`
	Args        string      `json:"args,omitempty" yaml:"args,omitempty"`
	Env         []string    `json:"env,omitempty" yaml:"env,omitempty"`
}

type BuildImageStepConfig struct {
	DockerfilePath string `json:"dockerFilePath,omittempty" yaml:"dockerFilePath,omitempty"`
	BuildPath      string `json:"buildPath,omitempty" yaml:"buildPath,omitempty"`
	TargetImage    string `json:"targetImage,omitempty" yaml:"targetImage,omitempty"`
	Push       bool   `json:"push" yaml:"push,omitempty"`
}




type PipelineHistory struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec PipelineHistorySpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status PipelineHistoryStatus `json:"status"`
}

type PipelineHistorySpec struct {
	Id              string            `json:"id,omitempty"`
	RunNumber     int               `json:"runNumber,omitempty"`
	TriggerType     string            `json:"triggerType,omitempty"`
	Pipeline        PipelineSpec          `json:"pipeline,omitempty"`
}

type PipelineHistoryStatus struct {
	CommitInfo      string            `json:"commitInfo,omitempty"`
	EnvVars         map[string]string `json:"envVars,omitempty"`
	Status          string            `json:"status,omitempty"`
	StartTime         int64             `json:"startTime,omitempty"`
	EndTime          int64             `json:"endTime,omitempty"`
	StageStatus	[]StageStatus `json:"stageStatus,omitempty"`
}

type StageStatus struct {
	Status          string            `json:"status,omitempty"`
	StartTime         int64             `json:"startTime,omitempty"`
	EndTime          int64             `json:"endTime,omitempty"`
	StepStatus	[]StepStatus `json:"stepStatus,omitempty"`
}

type StepStatus struct {
	Status          string            `json:"status,omitempty"`
	StartTime         int64             `json:"startTime,omitempty"`
	EndTime          int64             `json:"endTime,omitempty"`
}



