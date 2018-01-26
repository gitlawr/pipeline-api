package client

import (
	"github.com/rancher/norman/types"
)

const (
	PipelineType                 = "pipeline"
	PipelineFieldAnnotations     = "annotations"
	PipelineFieldCommitInfo      = "commitInfo"
	PipelineFieldCreated         = "created"
	PipelineFieldCreatorID       = "creatorId"
	PipelineFieldCronTrigger     = "cronTrigger"
	PipelineFieldId              = "id"
	PipelineFieldIsActivate      = "isActivate"
	PipelineFieldKeepWorkspace   = "keepWorkspace"
	PipelineFieldLabels          = "labels"
	PipelineFieldLastRunId       = "lastRunId"
	PipelineFieldLastRunStatus   = "lastRunStatus"
	PipelineFieldLastRunTime     = "lastRunTime"
	PipelineFieldName            = "name"
	PipelineFieldNextRunTime     = "nextRunTime"
	PipelineFieldOwnerReferences = "ownerReferences"
	PipelineFieldParameters      = "parameters"
	PipelineFieldRemoved         = "removed"
	PipelineFieldRunCount        = "runCount"
	PipelineFieldStages          = "stages"
	PipelineFieldTemplates       = "templates"
	PipelineFieldUuid            = "uuid"
	PipelineFieldWebHookId       = "webhookId"
	PipelineFieldWebHookToken    = "webhookToken"
)

type Pipeline struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty"`
	CommitInfo      string            `json:"commitInfo,omitempty"`
	Created         string            `json:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty"`
	CronTrigger     *CronTrigger      `json:"cronTrigger,omitempty"`
	Id              string            `json:"id,omitempty"`
	IsActivate      *bool             `json:"isActivate,omitempty"`
	KeepWorkspace   *bool             `json:"keepWorkspace,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	LastRunId       string            `json:"lastRunId,omitempty"`
	LastRunStatus   string            `json:"lastRunStatus,omitempty"`
	LastRunTime     *int64            `json:"lastRunTime,omitempty"`
	Name            string            `json:"name,omitempty"`
	NextRunTime     *int64            `json:"nextRunTime,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty"`
	Parameters      []string          `json:"parameters,omitempty"`
	Removed         string            `json:"removed,omitempty"`
	RunCount        *int64            `json:"runCount,omitempty"`
	Stages          []Stage           `json:"stages,omitempty"`
	Templates       map[string]string `json:"templates,omitempty"`
	Uuid            string            `json:"uuid,omitempty"`
	WebHookId       *int64            `json:"webhookId,omitempty"`
	WebHookToken    string            `json:"webhookToken,omitempty"`
}
type PipelineCollection struct {
	types.Collection
	Data   []Pipeline `json:"data,omitempty"`
	client *PipelineClient
}

type PipelineClient struct {
	apiClient *Client
}

type PipelineOperations interface {
	List(opts *types.ListOpts) (*PipelineCollection, error)
	Create(opts *Pipeline) (*Pipeline, error)
	Update(existing *Pipeline, updates interface{}) (*Pipeline, error)
	ByID(id string) (*Pipeline, error)
	Delete(container *Pipeline) error
}

func newPipelineClient(apiClient *Client) *PipelineClient {
	return &PipelineClient{
		apiClient: apiClient,
	}
}

func (c *PipelineClient) Create(container *Pipeline) (*Pipeline, error) {
	resp := &Pipeline{}
	err := c.apiClient.Ops.DoCreate(PipelineType, container, resp)
	return resp, err
}

func (c *PipelineClient) Update(existing *Pipeline, updates interface{}) (*Pipeline, error) {
	resp := &Pipeline{}
	err := c.apiClient.Ops.DoUpdate(PipelineType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PipelineClient) List(opts *types.ListOpts) (*PipelineCollection, error) {
	resp := &PipelineCollection{}
	err := c.apiClient.Ops.DoList(PipelineType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PipelineCollection) Next() (*PipelineCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PipelineCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PipelineClient) ByID(id string) (*Pipeline, error) {
	resp := &Pipeline{}
	err := c.apiClient.Ops.DoByID(PipelineType, id, resp)
	return resp, err
}

func (c *PipelineClient) Delete(container *Pipeline) error {
	return c.apiClient.Ops.DoResourceDelete(PipelineType, &container.Resource)
}
