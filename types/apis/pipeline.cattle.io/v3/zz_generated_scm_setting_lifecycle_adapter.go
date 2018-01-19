package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type SCMSettingLifecycle interface {
	Create(obj *SCMSetting) (*SCMSetting, error)
	Remove(obj *SCMSetting) (*SCMSetting, error)
	Updated(obj *SCMSetting) (*SCMSetting, error)
}

type scmSettingLifecycleAdapter struct {
	lifecycle SCMSettingLifecycle
}

func (w *scmSettingLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*SCMSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *scmSettingLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*SCMSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *scmSettingLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*SCMSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewSCMSettingLifecycleAdapter(name string, client SCMSettingInterface, l SCMSettingLifecycle) SCMSettingHandlerFunc {
	adapter := &scmSettingLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, adapter, client.ObjectClient())
	return func(key string, obj *SCMSetting) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
