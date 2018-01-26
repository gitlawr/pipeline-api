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
	PipelineHistoriesGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	pipelineControllers        map[string]PipelineController
	pipelineHistoryControllers map[string]PipelineHistoryController
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

		pipelineControllers:        map[string]PipelineController{},
		pipelineHistoryControllers: map[string]PipelineHistoryController{},
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

type PipelineHistoriesGetter interface {
	PipelineHistories(namespace string) PipelineHistoryInterface
}

func (c *Client) PipelineHistories(namespace string) PipelineHistoryInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &PipelineHistoryResource, PipelineHistoryGroupVersionKind, pipelineHistoryFactory{})
	return &pipelineHistoryClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
