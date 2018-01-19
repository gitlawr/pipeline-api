package client

import (
	"github.com/rancher/norman/types"
)

const (
	SCMSettingType                      = "scmSetting"
	SCMSettingFieldAnnotations          = "annotations"
	SCMSettingFieldClientID             = "clientID"
	SCMSettingFieldClientSecret         = "clientSecret"
	SCMSettingFieldCreated              = "created"
	SCMSettingFieldCreatorID            = "creatorId"
	SCMSettingFieldHomePage             = "homepage"
	SCMSettingFieldHostName             = "hostName"
	SCMSettingFieldIsAuth               = "isAuth"
	SCMSettingFieldLabels               = "labels"
	SCMSettingFieldName                 = "name"
	SCMSettingFieldOwnerReferences      = "ownerReferences"
	SCMSettingFieldRedirectURL          = "redirectURL"
	SCMSettingFieldRemoved              = "removed"
	SCMSettingFieldScheme               = "scheme"
	SCMSettingFieldScmType              = "scmType"
	SCMSettingFieldState                = "state"
	SCMSettingFieldStatus               = "status"
	SCMSettingFieldTransitioning        = "transitioning"
	SCMSettingFieldTransitioningMessage = "transitioningMessage"
	SCMSettingFieldUuid                 = "uuid"
)

type SCMSetting struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty"`
	ClientID             string            `json:"clientID,omitempty"`
	ClientSecret         string            `json:"clientSecret,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	HomePage             string            `json:"homepage,omitempty"`
	HostName             string            `json:"hostName,omitempty"`
	IsAuth               *bool             `json:"isAuth,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Name                 string            `json:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	RedirectURL          string            `json:"redirectURL,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	Scheme               string            `json:"scheme,omitempty"`
	ScmType              string            `json:"scmType,omitempty"`
	State                string            `json:"state,omitempty"`
	Status               *SCMSettingStatus `json:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type SCMSettingCollection struct {
	types.Collection
	Data   []SCMSetting `json:"data,omitempty"`
	client *SCMSettingClient
}

type SCMSettingClient struct {
	apiClient *Client
}

type SCMSettingOperations interface {
	List(opts *types.ListOpts) (*SCMSettingCollection, error)
	Create(opts *SCMSetting) (*SCMSetting, error)
	Update(existing *SCMSetting, updates interface{}) (*SCMSetting, error)
	ByID(id string) (*SCMSetting, error)
	Delete(container *SCMSetting) error
}

func newSCMSettingClient(apiClient *Client) *SCMSettingClient {
	return &SCMSettingClient{
		apiClient: apiClient,
	}
}

func (c *SCMSettingClient) Create(container *SCMSetting) (*SCMSetting, error) {
	resp := &SCMSetting{}
	err := c.apiClient.Ops.DoCreate(SCMSettingType, container, resp)
	return resp, err
}

func (c *SCMSettingClient) Update(existing *SCMSetting, updates interface{}) (*SCMSetting, error) {
	resp := &SCMSetting{}
	err := c.apiClient.Ops.DoUpdate(SCMSettingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SCMSettingClient) List(opts *types.ListOpts) (*SCMSettingCollection, error) {
	resp := &SCMSettingCollection{}
	err := c.apiClient.Ops.DoList(SCMSettingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SCMSettingCollection) Next() (*SCMSettingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SCMSettingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SCMSettingClient) ByID(id string) (*SCMSetting, error) {
	resp := &SCMSetting{}
	err := c.apiClient.Ops.DoByID(SCMSettingType, id, resp)
	return resp, err
}

func (c *SCMSettingClient) Delete(container *SCMSetting) error {
	return c.apiClient.Ops.DoResourceDelete(SCMSettingType, &container.Resource)
}
