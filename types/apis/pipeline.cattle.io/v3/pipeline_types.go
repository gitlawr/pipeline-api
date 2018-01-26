package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	//###Spec PipelineSpec `json:"spec"`
	PipelineSpec
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	//###Status PipelineStatus `json:"status"`
	PipelineStatus
}

type PipelineStatus struct {
	//VersionSequence string `json:"-" yaml:"-"`
	RunCount      int    `json:"runCount" yaml:"runCount,omitempty"`
	LastRunId     string `json:"lastRunId,omitempty" yaml:"lastRunId,omitempty"`
	LastRunStatus string `json:"lastRunStatus,omitempty" yaml:"lastRunStatus,omitempty"`
	LastRunTime   int64  `json:"lastRunTime,omitempty" yaml:"lastRunTime,omitempty"`
	NextRunTime   int64  `json:"nextRunTime,omitempty" yaml:"nextRunTime,omitempty"`
	CommitInfo    string `json:"commitInfo,omitempty" yaml:"commitInfo,omitempty"`
	WebHookId     int    `json:"webhookId,omitempty" yaml:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty" yaml:"webhookToken,omitempty"`
}

type PipelineSpec struct {
	Id         string `json:"id,omitempty"`
	IsActivate bool   `json:"isActivate" yaml:"isActivate"`
	//user defined environment variables
	Parameters []string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	//for import
	Templates map[string]string `json:"templates,omitempty" yaml:"templates,omitempty"`
	//trigger
	CronTrigger   CronTrigger `json:"cronTrigger,omitempty" yaml:"cronTrigger,omitempty"`
	Stages        []Stage     `json:"stages,omitempty" yaml:"stages,omitempty"`
	KeepWorkspace bool        `json:"keepWorkspace,omitempty" yaml:"keepWorkspace,omitempty"`
}

type CronTrigger struct {
	TriggerOnUpdate bool   `json:"triggerOnUpdate" yaml:"triggerOnUpdate,omitempty"`
	Spec            string `json:"spec,omitempty" yaml:"spec,omitempty"`
	Timezone        string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
}

type Stage struct {
	Name        string             `json:"name,omitempty" yaml:"name,omitempty"`
	NeedApprove bool               `json:"needApprove" yaml:"needApprove,omitempty"`
	Parallel    bool               `json:"parallel" yaml:"parallel,omitempty"`
	Condition   string             `json:"condition,omitempty" yaml:"condition,omitempty"`
	Conditions  PipelineConditions `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Approvers   []string           `json:"approvers,omitempty" yaml:"approvers,omitempty"`
	Steps       []Step             `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type Step struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	//Step timeout in minutes
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	//Condition  string             `json:"condition,omitempty" yaml:"condition,omitempty"`
	Conditions PipelineConditions `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	//---SCM step
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
	Branch     string `json:"branch,omitempty" yaml:"branch,omitempty"`
	GitUser    string `json:"gitUser,omitempty" yaml:"gitUser,omitempty"`
	Webhook    bool   `json:"webhook" yaml:"webhook,omitempty"`
	//---Build step
	Dockerfile     string `json:"dockerFileContent,omitempty" yaml:"dockerFileContent,omitempty"`
	BuildPath      string `json:"buildPath,omitempty" yaml:"buildPath,omitempty"`
	DockerfilePath string `json:"dockerFilePath,omittempty" yaml:"dockerFilePath,omitempty"`
	TargetImage    string `json:"targetImage,omitempty" yaml:"targetImage,omitempty"`
	PushFlag       bool   `json:"push" yaml:"push,omitempty"`

	//---task step
	Image       string      `json:"image,omitempty" yaml:"image,omitempty"`
	IsService   bool        `json:"isService" yaml:"isService,omitempty"`
	Alias       string      `json:"alias,omitempty" yaml:"alias,omitempty"`
	ShellScript string      `json:"shellScript,omitempty" yaml:"shellScript,omitempty"`
	Entrypoint  string      `json:"entrypoint,omitempty" yaml:"enrtypoint,omitempty"`
	Args        string      `json:"args,omitempty" yaml:"args,omitempty"`
	Env         []string    `json:"env,omitempty" yaml:"env,omitempty"`
	Services    []CIService `json:"services,omitempty" yaml:"services,omitempty"`

	//---upgradeService step
	ImageTag        string            `json:"imageTag,omitempty" yaml:"imageTag,omitempty"`
	ServiceSelector map[string]string `json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
	BatchSize       int               `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`
	Interval        int               `json:"interval,omitempty" yaml:"interval,omitempty"`
	StartFirst      bool              `json:"startFirst" yaml:"startFirst,omitempty"`
	Endpoint        string            `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Accesskey       string            `json:"accesskey,omitempty" yaml:"accesskey,omitempty"`
	Secretkey       string            `json:"secretkey,omitempty" yaml:"secretkey,omitempty"`

	//---upgradeStack step
	//Endpoint,Accesskey,Secretkey
	StackName      string `json:"stackName,omitempty" yaml:"stackName,omitempty"`
	DockerCompose  string `json:"dockerCompose,omitempty" yaml:"dockerCompose,omitempty"`
	RancherCompose string `json:"rancherCompose,omitempty" yaml:"rancherCompose,omitempty"`

	//---upgradeCatalog step
	//Endpoint,Accesskey,Secretkey,StackName,
	ExternalId string            `json:"externalId,omitempty" yaml:"externalId,omitempty"`
	DeployFlag bool              `json:"deploy" yaml:"deploy,omitempty"`
	Templates  map[string]string `json:"templates,omitempty" yaml:"templates,omitempty"`
	Answers    string            `json:"answerString,omitempty" yaml:"answerString,omitempty"`
}

type PipelineConditions struct {
	All []string `json:"all,omitempty" yaml:"all,omitempty"`
	Any []string `json:"any,omitempty" yaml:"any,omitempty"`
}

type CIService struct {
	ContainerName string `json:"containerName,omitempty"`
	Name          string `json:"name,omitempty"`
	Image         string `json:"image,omitempty"`
	Entrypoint    string `json:"entrypoint,omitempty"`
	Command       string `json:"command,omitempty"`
}
