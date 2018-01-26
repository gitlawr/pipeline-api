package pipeline

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
	"github.com/rancher/pipeline-api/types/client/pipeline/v3"
	"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	//"github.com/rancher/pipeline-api/api/provider"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/pipeline-api/api/provider"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.Actions["run"] = apiContext.URLBuilder.Action("run", resource)
	resource.Actions["update"] = apiContext.URLBuilder.Action("update", resource)
	resource.Actions["activate"] = apiContext.URLBuilder.Action("activate", resource)
	resource.Actions["deactivate"] = apiContext.URLBuilder.Action("deactivate", resource)
	resource.Actions["remove"] = apiContext.URLBuilder.Action("remove", resource)
	resource.Actions["export"] = apiContext.URLBuilder.Action("export", resource)
}

func PipelineActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Infof("do pipeline action:%s",actionName)
	logrus.Infof("action input:%s",action.Input)
	logrus.Infof("action output:%s",action.Output)

	if actionName == "run"{
		return RunPipeline(apiContext)
	}

	//TODO Implement Actions
	/*
	store := apiContext.Schema.Store

	data, err := store.ByID(apiContext, apiContext.Schema, apiContext.ID)
	if err != nil {
		return err
	}
	data["lastRefreshTimestamp"] = time.Now().Format(time.RFC3339)

	_, err = store.Update(apiContext, apiContext.Schema, data, apiContext.ID)
	if err != nil {
		return err
	}
	*/
	return nil
}

func RunPipeline(apiContext *types.APIContext) error{

	pipelineSchema:= apiContext.Schema
	pipelineStore :=pipelineSchema.Store
	data, err := pipelineStore.ByID(apiContext, apiContext.Schema, apiContext.ID)
	if err != nil {
		return err
	}
	pipeline := v3.Pipeline{}

	convert.ToObj(data,&pipeline)

	activity,err:=provider.PipelineProvidor.RunPipeline(&pipeline,"manul")
	if err != nil{
		return err
	}

	activityMap,err := convert.EncodeToMap(activity)
	if err != nil{
		return err
	}
	activitySchema := apiContext.Schemas.Schema(&pipelineSchema.Version,client.ActivityType)
	activityStore := activitySchema.Store
	_,err = activityStore.Create(apiContext,activitySchema,activityMap)
	if err != nil{
		return err
	}
	return nil
}
