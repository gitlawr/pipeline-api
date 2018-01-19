package client

import (
	"github.com/rancher/norman/types"
)

const (
	GitAccountType                      = "gitAccount"
	GitAccountFieldAccessToken          = "accessToken"
	GitAccountFieldAccountType          = "accountType"
	GitAccountFieldAnnotations          = "annotations"
	GitAccountFieldAvatarURL            = "avatar_url"
	GitAccountFieldCreated              = "created"
	GitAccountFieldCreatorID            = "creatorId"
	GitAccountFieldHTMLURL              = "html_url"
	GitAccountFieldLabels               = "labels"
	GitAccountFieldLogin                = "login"
	GitAccountFieldName                 = "name"
	GitAccountFieldOwnerReferences      = "ownerReferences"
	GitAccountFieldPrivate              = "private"
	GitAccountFieldRancherUserID        = "rancherUserId"
	GitAccountFieldRemoved              = "removed"
	GitAccountFieldState                = "state"
	GitAccountFieldStatus               = "status"
	GitAccountFieldTransitioning        = "transitioning"
	GitAccountFieldTransitioningMessage = "transitioningMessage"
	GitAccountFieldUuid                 = "uuid"
)

type GitAccount struct {
	types.Resource
	AccessToken          string            `json:"accessToken,omitempty"`
	AccountType          string            `json:"accountType,omitempty"`
	Annotations          map[string]string `json:"annotations,omitempty"`
	AvatarURL            string            `json:"avatar_url,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	HTMLURL              string            `json:"html_url,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Login                string            `json:"login,omitempty"`
	Name                 string            `json:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	Private              *bool             `json:"private,omitempty"`
	RancherUserID        string            `json:"rancherUserId,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	State                string            `json:"state,omitempty"`
	Status               *GitAccountStatus `json:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type GitAccountCollection struct {
	types.Collection
	Data   []GitAccount `json:"data,omitempty"`
	client *GitAccountClient
}

type GitAccountClient struct {
	apiClient *Client
}

type GitAccountOperations interface {
	List(opts *types.ListOpts) (*GitAccountCollection, error)
	Create(opts *GitAccount) (*GitAccount, error)
	Update(existing *GitAccount, updates interface{}) (*GitAccount, error)
	ByID(id string) (*GitAccount, error)
	Delete(container *GitAccount) error
}

func newGitAccountClient(apiClient *Client) *GitAccountClient {
	return &GitAccountClient{
		apiClient: apiClient,
	}
}

func (c *GitAccountClient) Create(container *GitAccount) (*GitAccount, error) {
	resp := &GitAccount{}
	err := c.apiClient.Ops.DoCreate(GitAccountType, container, resp)
	return resp, err
}

func (c *GitAccountClient) Update(existing *GitAccount, updates interface{}) (*GitAccount, error) {
	resp := &GitAccount{}
	err := c.apiClient.Ops.DoUpdate(GitAccountType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GitAccountClient) List(opts *types.ListOpts) (*GitAccountCollection, error) {
	resp := &GitAccountCollection{}
	err := c.apiClient.Ops.DoList(GitAccountType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GitAccountCollection) Next() (*GitAccountCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GitAccountCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GitAccountClient) ByID(id string) (*GitAccount, error) {
	resp := &GitAccount{}
	err := c.apiClient.Ops.DoByID(GitAccountType, id, resp)
	return resp, err
}

func (c *GitAccountClient) Delete(container *GitAccount) error {
	return c.apiClient.Ops.DoResourceDelete(GitAccountType, &container.Resource)
}
