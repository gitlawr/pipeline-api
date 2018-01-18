package v3

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ActivityGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Activity",
	}
	ActivityResource = metav1.APIResource{
		Name:         "activities",
		SingularName: "activity",
		Namespaced:   false,
		Kind:         ActivityGroupVersionKind.Kind,
	}
)

type ActivityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Activity
}

type ActivityHandlerFunc func(key string, obj *Activity) error

type ActivityLister interface {
	List(namespace string, selector labels.Selector) (ret []*Activity, err error)
	Get(namespace, name string) (*Activity, error)
}

type ActivityController interface {
	Informer() cache.SharedIndexInformer
	Lister() ActivityLister
	AddHandler(handler ActivityHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ActivityInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*Activity) (*Activity, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*Activity, error)
	Get(name string, opts metav1.GetOptions) (*Activity, error)
	Update(*Activity) (*Activity, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ActivityList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ActivityController
	AddSyncHandler(sync ActivityHandlerFunc)
	AddLifecycle(name string, lifecycle ActivityLifecycle)
}

type activityLister struct {
	controller *activityController
}

func (l *activityLister) List(namespace string, selector labels.Selector) (ret []*Activity, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*Activity))
	})
	return
}

func (l *activityLister) Get(namespace, name string) (*Activity, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ActivityGroupVersionKind.Group,
			Resource: "activity",
		}, name)
	}
	return obj.(*Activity), nil
}

type activityController struct {
	controller.GenericController
}

func (c *activityController) Lister() ActivityLister {
	return &activityLister{
		controller: c,
	}
}

func (c *activityController) AddHandler(handler ActivityHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Activity))
	})
}

type activityFactory struct {
}

func (c activityFactory) Object() runtime.Object {
	return &Activity{}
}

func (c activityFactory) List() runtime.Object {
	return &ActivityList{}
}

func (s *activityClient) Controller() ActivityController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.activityControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ActivityGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &activityController{
		GenericController: genericController,
	}

	s.client.activityControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type activityClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ActivityController
}

func (s *activityClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *activityClient) Create(o *Activity) (*Activity, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Activity), err
}

func (s *activityClient) Get(name string, opts metav1.GetOptions) (*Activity, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Activity), err
}

func (s *activityClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*Activity, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*Activity), err
}

func (s *activityClient) Update(o *Activity) (*Activity, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Activity), err
}

func (s *activityClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *activityClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *activityClient) List(opts metav1.ListOptions) (*ActivityList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ActivityList), err
}

func (s *activityClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *activityClient) Patch(o *Activity, data []byte, subresources ...string) (*Activity, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*Activity), err
}

func (s *activityClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *activityClient) AddSyncHandler(sync ActivityHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *activityClient) AddLifecycle(name string, lifecycle ActivityLifecycle) {
	sync := NewActivityLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
