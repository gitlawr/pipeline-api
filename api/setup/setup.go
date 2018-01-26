package setup

import (
	"context"

	"github.com/rancher/pipeline-api/api/pipeline"
	"github.com/rancher/pipeline-api/api/pipeline-hostory"
	"github.com/rancher/pipeline-api/api/git-account"
	"github.com/rancher/pipeline-api/api/scm-setting"
	"github.com/rancher/pipeline-api/store/scoped"
	"github.com/rancher/norman/api/builtin"
	"github.com/rancher/norman/pkg/subscribe"
	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
	pipelineschema "github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3/schema"
	"github.com/rancher/pipeline-api/types/client/pipeline/v3"
	"github.com/rancher/pipeline-api/types/config"
)

var (
	crdVersions = []*types.APIVersion{
		&pipelineschema.Version,
	}
)

func Schemas(ctx context.Context, pipeline *config.PipelineContext) error {
	schemas := pipeline.Schemas
	subscribe.Register(&builtin.Version, schemas)
	Pipeline(schemas)
	PipelineHistory(schemas)
	crdStore, err := crd.NewCRDStoreFromConfig(pipeline.RESTConfig)
	if err != nil {
		return err
	}

	var crdSchemas []*types.Schema
	for _, version := range crdVersions {
		for _, schema := range schemas.SchemasForVersion(*version) {
			crdSchemas = append(crdSchemas, schema)
		}
	}

	if err := crdStore.AddSchemas(ctx, crdSchemas...); err != nil {
		return err
	}

	NamespacedTypes(schemas)

	return nil
}

func NamespacedTypes(schemas *types.Schemas) {
	for _, version := range crdVersions {
		for _, schema := range schemas.SchemasForVersion(*version) {
			if schema.Scope != types.NamespaceScope || schema.Store == nil {
				continue
			}

			for _, key := range []string{"projectId", "clusterId"} {
				ns, ok := schema.ResourceFields["namespaceId"]
				if !ok {
					continue
				}

				if _, ok := schema.ResourceFields[key]; !ok {
					continue
				}

				schema.Store = scoped.NewScopedStore(key, schema.Store)
				ns.Required = false
				schema.ResourceFields["namespaceId"] = ns
				break
			}
		}
	}
}


func Pipeline(schemas *types.Schemas) {
	schema := schemas.Schema(&pipelineschema.Version, client.PipelineType)
	schema.ResourceActions = map[string]types.Action{
		"run": {},
		"update":     {},
		"activate":   {},
		"deactivate": {},
		"remove":     {},
		"export":     {},
	}
	schema.Formatter = pipeline.Formatter
	schema.ActionHandler = pipeline.PipelineActionHandler
}

func PipelineHistory(schemas *types.Schemas) {
	schema := schemas.Schema(&pipelineschema.Version, client.PipelineHistoryType)
	schema.ResourceActions = map[string]types.Action{
		"update":     {},
		"remove":     {},
		"approve": {},
		"deny":   {},
		"rerun": {},
		"stop":     {},
	}
	schema.Formatter = pipeline_hostory.Formatter
	schema.ActionHandler = pipeline_hostory.ActivityActionHandler
}
