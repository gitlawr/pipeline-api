package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ActivityLifecycle interface {
	Create(obj *Activity) (*Activity, error)
	Remove(obj *Activity) (*Activity, error)
	Updated(obj *Activity) (*Activity, error)
}

type activityLifecycleAdapter struct {
	lifecycle ActivityLifecycle
}

func (w *activityLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*Activity))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activityLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*Activity))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activityLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*Activity))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewActivityLifecycleAdapter(name string, client ActivityInterface, l ActivityLifecycle) ActivityHandlerFunc {
	adapter := &activityLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, adapter, client.ObjectClient())
	return func(key string, obj *Activity) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
