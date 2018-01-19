package scm_setting

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.Actions["update"] = apiContext.URLBuilder.Action("update", resource)
	resource.Actions["remove"] = apiContext.URLBuilder.Action("remove", resource)
	resource.Actions["oauth"] = apiContext.URLBuilder.Action("oauth", resource)
}

func SCMSettingActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Infof("do activity action:%s",actionName)
	//TODO Implement Actions
	return nil
}
