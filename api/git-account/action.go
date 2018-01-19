package git_account

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.Actions["share"] = apiContext.URLBuilder.Action("share", resource)
	resource.Actions["unshare"] = apiContext.URLBuilder.Action("unshare", resource)
	resource.Actions["remove"] = apiContext.URLBuilder.Action("remove", resource)
	resource.Actions["refreshrepo"] = apiContext.URLBuilder.Action("refreshrepo", resource)
}

func GitAccountActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Infof("do git account action:%s",actionName)
	//TODO Implement Actions
	return nil
}
