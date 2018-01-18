package controller

import (

	"github.com/rancher/pipeline-api/types/config"
	pipelineController "github.com/rancher/pipeline-api/controller/pipeline"
)

func Register(pipeline *config.PipelineContext) {
	pipelineController.Register(pipeline)
}
