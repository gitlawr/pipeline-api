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
	SCMSettingGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "SCMSetting",
	}
	SCMSettingResource = metav1.APIResource{
		Name:         "scmsettings",
		SingularName: "scmsetting",
		Namespaced:   false,
		Kind:         SCMSettingGroupVersionKind.Kind,
	}
)

type SCMSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SCMSetting
}

type SCMSettingHandlerFunc func(key string, obj *SCMSetting) error

type SCMSettingLister interface {
	List(namespace string, selector labels.Selector) (ret []*SCMSetting, err error)
	Get(namespace, name string) (*SCMSetting, error)
}

type SCMSettingController interface {
	Informer() cache.SharedIndexInformer
	Lister() SCMSettingLister
	AddHandler(handler SCMSettingHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type SCMSettingInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*SCMSetting) (*SCMSetting, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*SCMSetting, error)
	Get(name string, opts metav1.GetOptions) (*SCMSetting, error)
	Update(*SCMSetting) (*SCMSetting, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*SCMSettingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() SCMSettingController
	AddSyncHandler(sync SCMSettingHandlerFunc)
	AddLifecycle(name string, lifecycle SCMSettingLifecycle)
}

type scmSettingLister struct {
	controller *scmSettingController
}

func (l *scmSettingLister) List(namespace string, selector labels.Selector) (ret []*SCMSetting, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*SCMSetting))
	})
	return
}

func (l *scmSettingLister) Get(namespace, name string) (*SCMSetting, error) {
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
			Group:    SCMSettingGroupVersionKind.Group,
			Resource: "scmSetting",
		}, name)
	}
	return obj.(*SCMSetting), nil
}

type scmSettingController struct {
	controller.GenericController
}

func (c *scmSettingController) Lister() SCMSettingLister {
	return &scmSettingLister{
		controller: c,
	}
}

func (c *scmSettingController) AddHandler(handler SCMSettingHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*SCMSetting))
	})
}

type scmSettingFactory struct {
}

func (c scmSettingFactory) Object() runtime.Object {
	return &SCMSetting{}
}

func (c scmSettingFactory) List() runtime.Object {
	return &SCMSettingList{}
}

func (s *scmSettingClient) Controller() SCMSettingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.scmSettingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(SCMSettingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &scmSettingController{
		GenericController: genericController,
	}

	s.client.scmSettingControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type scmSettingClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   SCMSettingController
}

func (s *scmSettingClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *scmSettingClient) Create(o *SCMSetting) (*SCMSetting, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*SCMSetting), err
}

func (s *scmSettingClient) Get(name string, opts metav1.GetOptions) (*SCMSetting, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*SCMSetting), err
}

func (s *scmSettingClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*SCMSetting, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*SCMSetting), err
}

func (s *scmSettingClient) Update(o *SCMSetting) (*SCMSetting, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*SCMSetting), err
}

func (s *scmSettingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *scmSettingClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *scmSettingClient) List(opts metav1.ListOptions) (*SCMSettingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*SCMSettingList), err
}

func (s *scmSettingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *scmSettingClient) Patch(o *SCMSetting, data []byte, subresources ...string) (*SCMSetting, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*SCMSetting), err
}

func (s *scmSettingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *scmSettingClient) AddSyncHandler(sync SCMSettingHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *scmSettingClient) AddLifecycle(name string, lifecycle SCMSettingLifecycle) {
	sync := NewSCMSettingLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
