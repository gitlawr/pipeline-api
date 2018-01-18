package client

import (
	"github.com/rancher/norman/types"
)

const (
	ActivityType                      = "activity"
	ActivityFieldActivityStages       = "activity_stages"
	ActivityFieldAnnotations          = "annotations"
	ActivityFieldCommitInfo           = "commitInfo"
	ActivityFieldCreated              = "created"
	ActivityFieldCreatorID            = "creatorId"
	ActivityFieldEnvVars              = "envVars"
	ActivityFieldFailMessage          = "failMessage"
	ActivityFieldLabels               = "labels"
	ActivityFieldName                 = "name"
	ActivityFieldNodeName             = "nodename"
	ActivityFieldOwnerReferences      = "ownerReferences"
	ActivityFieldPendingStage         = "pendingStage"
	ActivityFieldPipeline             = "pipelineSource"
	ActivityFieldPipelineName         = "pipelineName"
	ActivityFieldPipelineVersion      = "pipelineVersion"
	ActivityFieldRemoved              = "removed"
	ActivityFieldRunSequence          = "runSequence"
	ActivityFieldStartTS              = "start_ts"
	ActivityFieldState                = "state"
	ActivityFieldStatus               = "status"
	ActivityFieldStopTS               = "stop_ts"
	ActivityFieldTransitioning        = "transitioning"
	ActivityFieldTransitioningMessage = "transitioningMessage"
	ActivityFieldTriggerType          = "triggerType"
	ActivityFieldUuid                 = "uuid"
)

type Activity struct {
	types.Resource
	ActivityStages       []ActivityStage   `json:"activity_stages,omitempty"`
	Annotations          map[string]string `json:"annotations,omitempty"`
	CommitInfo           string            `json:"commitInfo,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	EnvVars              map[string]string `json:"envVars,omitempty"`
	FailMessage          string            `json:"failMessage,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NodeName             string            `json:"nodename,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	PendingStage         *int64            `json:"pendingStage,omitempty"`
	Pipeline             *Pipeline         `json:"pipelineSource,omitempty"`
	PipelineName         string            `json:"pipelineName,omitempty"`
	PipelineVersion      string            `json:"pipelineVersion,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	RunSequence          *int64            `json:"runSequence,omitempty"`
	StartTS              *int64            `json:"start_ts,omitempty"`
	State                string            `json:"state,omitempty"`
	Status               *ActivityStatus   `json:"status,omitempty"`
	StopTS               *int64            `json:"stop_ts,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	TriggerType          string            `json:"triggerType,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type ActivityCollection struct {
	types.Collection
	Data   []Activity `json:"data,omitempty"`
	client *ActivityClient
}

type ActivityClient struct {
	apiClient *Client
}

type ActivityOperations interface {
	List(opts *types.ListOpts) (*ActivityCollection, error)
	Create(opts *Activity) (*Activity, error)
	Update(existing *Activity, updates interface{}) (*Activity, error)
	ByID(id string) (*Activity, error)
	Delete(container *Activity) error
}

func newActivityClient(apiClient *Client) *ActivityClient {
	return &ActivityClient{
		apiClient: apiClient,
	}
}

func (c *ActivityClient) Create(container *Activity) (*Activity, error) {
	resp := &Activity{}
	err := c.apiClient.Ops.DoCreate(ActivityType, container, resp)
	return resp, err
}

func (c *ActivityClient) Update(existing *Activity, updates interface{}) (*Activity, error) {
	resp := &Activity{}
	err := c.apiClient.Ops.DoUpdate(ActivityType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ActivityClient) List(opts *types.ListOpts) (*ActivityCollection, error) {
	resp := &ActivityCollection{}
	err := c.apiClient.Ops.DoList(ActivityType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ActivityCollection) Next() (*ActivityCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ActivityCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ActivityClient) ByID(id string) (*Activity, error) {
	resp := &Activity{}
	err := c.apiClient.Ops.DoByID(ActivityType, id, resp)
	return resp, err
}

func (c *ActivityClient) Delete(container *Activity) error {
	return c.apiClient.Ops.DoResourceDelete(ActivityType, &container.Resource)
}
