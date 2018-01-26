package server

import (
	"context"
	"net/http"
	"github.com/rancher/pipeline-api/types/config"
	normanapi "github.com/rancher/norman/api"
	//"github.com/rancher/norman/types"
	//"github.com/rancher/norman/store/crd"
	//"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3/schema"
	//"github.com/sirupsen/logrus"
	"github.com/rancher/pipeline-api/controller"
	"github.com/rancher/pipeline-api/api"
)

func New(ctx context.Context, pipeline *config.PipelineContext) (http.Handler, error) {

	api.Schemas(ctx,pipeline)
	api.SetupProvider()
	/*
		store, err := crd.NewCRDStoreFromConfig(pipeline.RESTConfig)
		if err != nil {
			return nil,err
		}

		var crdSchemas []*types.Schema

		for _, schema := range pipeline.Schemas.SchemasForVersion(schema.Version) {
			crdSchemas = append(crdSchemas, schema)
		}

		if err := store.AddSchemas(ctx, crdSchemas...); err != nil {
			logrus.Error(err)
		}
	*/

	server := normanapi.NewAPIServer()
	if err := server.AddSchemas(pipeline.Schemas); err != nil {
		return nil,err
	}
	controller.Register(pipeline)

	return server,nil
}
