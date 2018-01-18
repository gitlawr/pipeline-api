package config

import (
	"context"

	"k8s.io/client-go/tools/record"
    "k8s.io/client-go/rest"
	"github.com/rancher/norman/event"
	"github.com/rancher/norman/signal"
	"k8s.io/client-go/kubernetes"
	"github.com/rancher/norman/types"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/rancher/norman/controller"
	pipelinev3 "github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	pipelineSchema "github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3/schema"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/api/core/v1"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
)
type PipelineContext struct {
	eventBroadcaster record.EventBroadcaster

	RESTConfig        rest.Config
	UnversionedClient rest.Interface
	K8sClient         kubernetes.Interface
	Events            record.EventRecorder
	EventLogger       event.Logger
	Schemas           *types.Schemas
	Scheme            *runtime.Scheme

	Pipeline pipelinev3.Interface
}

func (c *PipelineContext) controllers() []controller.Starter {
	return []controller.Starter{
		c.Pipeline,
	}
}


func NewPipelineContext(config rest.Config) (*PipelineContext, error) {
	var err error

	context := &PipelineContext{
		RESTConfig: config,
	}

	context.Pipeline, err = pipelinev3.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	context.K8sClient, err = kubernetes.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	dynamicConfig := config
	if dynamicConfig.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		dynamicConfig.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	context.UnversionedClient, err = rest.UnversionedRESTClientFor(&dynamicConfig)
	if err != nil {
		return nil, err
	}

	context.Schemas = types.NewSchemas().
		AddSchemas(pipelineSchema.Schemas)

	context.eventBroadcaster = record.NewBroadcaster()
	context.Events = context.eventBroadcaster.NewRecorder(context.Scheme, v1.EventSource{
		Component: "CattleManagementServer",
	})
	context.EventLogger = event.NewLogger(context.Events)

	return context, err
}

func (c *PipelineContext) Start(ctx context.Context) error {
	logrus.Info("Starting pipeline controllers")

	watcher := c.eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
		Interface: c.K8sClient.CoreV1().Events(""),
	})

	go func() {
		<-ctx.Done()
		watcher.Stop()
	}()

	return controller.SyncThenStart(ctx, 5, c.controllers()...)
}

func (c *PipelineContext) StartAndWait() error {
	ctx := signal.SigTermCancelContext(context.Background())
	c.Start(ctx)
	<-ctx.Done()
	return ctx.Err()
}
