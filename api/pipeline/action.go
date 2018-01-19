package pipeline

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
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

func RunPipeline(){

}
