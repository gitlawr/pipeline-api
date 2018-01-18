//go:generate go run generator/cleanup/main.go
//go:generate go run generate-code.go

package main

import (
	pipelineSchema "github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3/schema"
	"github.com/rancher/pipeline-api/generator"
)

func main() {
	generator.Generate(pipelineSchema.Schemas)
}
