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
	PipelineGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Pipeline",
	}
	PipelineResource = metav1.APIResource{
		Name:         "pipelines",
		SingularName: "pipeline",
		Namespaced:   false,
		Kind:         PipelineGroupVersionKind.Kind,
	}
)

type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline
}

type PipelineHandlerFunc func(key string, obj *Pipeline) error

type PipelineLister interface {
	List(namespace string, selector labels.Selector) (ret []*Pipeline, err error)
	Get(namespace, name string) (*Pipeline, error)
}

type PipelineController interface {
	Informer() cache.SharedIndexInformer
	Lister() PipelineLister
	AddHandler(handler PipelineHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PipelineInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*Pipeline) (*Pipeline, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*Pipeline, error)
	Get(name string, opts metav1.GetOptions) (*Pipeline, error)
	Update(*Pipeline) (*Pipeline, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PipelineList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PipelineController
	AddSyncHandler(sync PipelineHandlerFunc)
	AddLifecycle(name string, lifecycle PipelineLifecycle)
}

type pipelineLister struct {
	controller *pipelineController
}

func (l *pipelineLister) List(namespace string, selector labels.Selector) (ret []*Pipeline, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*Pipeline))
	})
	return
}

func (l *pipelineLister) Get(namespace, name string) (*Pipeline, error) {
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
			Group:    PipelineGroupVersionKind.Group,
			Resource: "pipeline",
		}, name)
	}
	return obj.(*Pipeline), nil
}

type pipelineController struct {
	controller.GenericController
}

func (c *pipelineController) Lister() PipelineLister {
	return &pipelineLister{
		controller: c,
	}
}

func (c *pipelineController) AddHandler(handler PipelineHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Pipeline))
	})
}

type pipelineFactory struct {
}

func (c pipelineFactory) Object() runtime.Object {
	return &Pipeline{}
}

func (c pipelineFactory) List() runtime.Object {
	return &PipelineList{}
}

func (s *pipelineClient) Controller() PipelineController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.pipelineControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PipelineGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &pipelineController{
		GenericController: genericController,
	}

	s.client.pipelineControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type pipelineClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   PipelineController
}

func (s *pipelineClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *pipelineClient) Create(o *Pipeline) (*Pipeline, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Pipeline), err
}

func (s *pipelineClient) Get(name string, opts metav1.GetOptions) (*Pipeline, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Pipeline), err
}

func (s *pipelineClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*Pipeline, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*Pipeline), err
}

func (s *pipelineClient) Update(o *Pipeline) (*Pipeline, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Pipeline), err
}

func (s *pipelineClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *pipelineClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *pipelineClient) List(opts metav1.ListOptions) (*PipelineList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PipelineList), err
}

func (s *pipelineClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *pipelineClient) Patch(o *Pipeline, data []byte, subresources ...string) (*Pipeline, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*Pipeline), err
}

func (s *pipelineClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *pipelineClient) AddSyncHandler(sync PipelineHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *pipelineClient) AddLifecycle(name string, lifecycle PipelineLifecycle) {
	sync := NewPipelineLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
