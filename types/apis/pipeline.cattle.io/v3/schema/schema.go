package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "pipeline.cattle.io",
		Path:    "/v3/pipeline.cattle.io",
		SubContexts: map[string]bool{
			"clusters": true,
		},
	}

	Schemas = factory.Schemas(&Version).
		Init(pipelineTypes)
)

func pipelineTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.Pipeline{}).
		AddMapperForType(&Version, v3.Activity{}).
		AddMapperForType(&Version, v3.GitAccount{}).
		AddMapperForType(&Version, v3.GitRepoCache{}).
		AddMapperForType(&Version, v3.SCMSetting{}).
		MustImport(&Version, v3.Pipeline{}).
		MustImport(&Version, v3.Activity{}).
		MustImport(&Version, v3.GitAccount{}).
		MustImport(&Version, v3.GitRepoCache{}).
		MustImport(&Version, v3.SCMSetting{})
}
