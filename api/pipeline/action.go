package pipeline

import (

	"github.com/rancher/norman/httperror"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)
func PipelineActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Info("do run pipeline")
	if actionName != "run" {
		return httperror.NewAPIError(httperror.NotFound, "not found")
	}
	//TODO RUN
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
