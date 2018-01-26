package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
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
		AddMapperForType(&Version, v3.PipelineConfig{}).
		MustImport(&Version, v3.PipelineConfig{}).
		AddMapperForType(&Version, v3.Pipeline{},
			&mapper.NamespaceIDMapper{},
			&m.DisplayName{},
		).
		MustImport(&Version, v3.Pipeline{}).
		AddMapperForType(&Version, v3.PipelineHistory{},
			&mapper.NamespaceIDMapper{},
		).
		MustImport(&Version, v3.PipelineHistory{})
	//AddMapperForType(&Version, v3.GitAccount{}).
	//AddMapperForType(&Version, v3.GitRepoCache{}).
	//MustImport(&Version, v3.GitAccount{}).
	//MustImport(&Version, v3.GitRepoCache{})
}
