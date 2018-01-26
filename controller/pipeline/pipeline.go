package pipeline

import (
	"github.com/rancher/norman/event"
	"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	"github.com/rancher/pipeline-api/types/config"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"github.com/sirupsen/logrus"
)

const (
	defaultEngineInstallURL = "https://releases.rancher.com/install-docker/17.03.2.sh"
)

func Register(pipeline *config.PipelineContext) {
	pipelineClient := pipeline.Pipeline.Pipelines("")

	pipelineLifecycle := &Lifecycle{
		pipelineClient:                pipelineClient,
		configMapGetter:              pipeline.K8sClient.CoreV1(),
		logger:                       pipeline.EventLogger,
	}

	pipelineClient.AddLifecycle("pipeline-controller", pipelineLifecycle)
}

type Lifecycle struct {
	pipelineClient                v3.PipelineInterface
	configMapGetter              typedv1.ConfigMapsGetter
	logger                       event.Logger
}

func (p *Lifecycle) Create(obj *v3.Pipeline) (*v3.Pipeline, error) {
	logrus.Infof("create apis:%v",obj.Name)
	return nil, nil
}

func (p *Lifecycle) Remove(obj *v3.Pipeline) (*v3.Pipeline, error) {
	logrus.Infof("remove apis:%v",obj.Name)
	return nil, nil
}

func (p *Lifecycle) Updated(obj *v3.Pipeline) (*v3.Pipeline, error) {
	logrus.Infof("update apis:%v",obj.Name)
	return nil, nil
}
