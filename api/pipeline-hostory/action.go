package pipeline_hostory

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.AddAction(apiContext, "update")
	resource.AddAction(apiContext, "remove")
	resource.AddAction(apiContext, "rerun")
	resource.AddAction(apiContext, "stop")
}

func ActivityActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Infof("do pipeline-hostory action:%s",actionName)
	//TODO Implement Actions
	return nil
}
