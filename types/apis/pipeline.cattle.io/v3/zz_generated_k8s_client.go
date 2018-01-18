package v3

import (
	"context"
	"sync"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	PipelinesGetter
	ActivitiesGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	pipelineControllers map[string]PipelineController
	activityControllers map[string]ActivityController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		config.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	restClient, err := rest.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		pipelineControllers: map[string]PipelineController{},
		activityControllers: map[string]ActivityController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type PipelinesGetter interface {
	Pipelines(namespace string) PipelineInterface
}

func (c *Client) Pipelines(namespace string) PipelineInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &PipelineResource, PipelineGroupVersionKind, pipelineFactory{})
	return &pipelineClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ActivitiesGetter interface {
	Activities(namespace string) ActivityInterface
}

func (c *Client) Activities(namespace string) ActivityInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ActivityResource, ActivityGroupVersionKind, activityFactory{})
	return &activityClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
