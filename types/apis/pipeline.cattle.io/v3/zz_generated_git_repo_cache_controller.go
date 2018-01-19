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
	GitRepoCacheGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GitRepoCache",
	}
	GitRepoCacheResource = metav1.APIResource{
		Name:         "gitrepocaches",
		SingularName: "gitrepocache",
		Namespaced:   false,
		Kind:         GitRepoCacheGroupVersionKind.Kind,
	}
)

type GitRepoCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitRepoCache
}

type GitRepoCacheHandlerFunc func(key string, obj *GitRepoCache) error

type GitRepoCacheLister interface {
	List(namespace string, selector labels.Selector) (ret []*GitRepoCache, err error)
	Get(namespace, name string) (*GitRepoCache, error)
}

type GitRepoCacheController interface {
	Informer() cache.SharedIndexInformer
	Lister() GitRepoCacheLister
	AddHandler(handler GitRepoCacheHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GitRepoCacheInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*GitRepoCache) (*GitRepoCache, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*GitRepoCache, error)
	Get(name string, opts metav1.GetOptions) (*GitRepoCache, error)
	Update(*GitRepoCache) (*GitRepoCache, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GitRepoCacheList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GitRepoCacheController
	AddSyncHandler(sync GitRepoCacheHandlerFunc)
	AddLifecycle(name string, lifecycle GitRepoCacheLifecycle)
}

type gitRepoCacheLister struct {
	controller *gitRepoCacheController
}

func (l *gitRepoCacheLister) List(namespace string, selector labels.Selector) (ret []*GitRepoCache, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GitRepoCache))
	})
	return
}

func (l *gitRepoCacheLister) Get(namespace, name string) (*GitRepoCache, error) {
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
			Group:    GitRepoCacheGroupVersionKind.Group,
			Resource: "gitRepoCache",
		}, name)
	}
	return obj.(*GitRepoCache), nil
}

type gitRepoCacheController struct {
	controller.GenericController
}

func (c *gitRepoCacheController) Lister() GitRepoCacheLister {
	return &gitRepoCacheLister{
		controller: c,
	}
}

func (c *gitRepoCacheController) AddHandler(handler GitRepoCacheHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*GitRepoCache))
	})
}

type gitRepoCacheFactory struct {
}

func (c gitRepoCacheFactory) Object() runtime.Object {
	return &GitRepoCache{}
}

func (c gitRepoCacheFactory) List() runtime.Object {
	return &GitRepoCacheList{}
}

func (s *gitRepoCacheClient) Controller() GitRepoCacheController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.gitRepoCacheControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GitRepoCacheGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &gitRepoCacheController{
		GenericController: genericController,
	}

	s.client.gitRepoCacheControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type gitRepoCacheClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   GitRepoCacheController
}

func (s *gitRepoCacheClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *gitRepoCacheClient) Create(o *GitRepoCache) (*GitRepoCache, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GitRepoCache), err
}

func (s *gitRepoCacheClient) Get(name string, opts metav1.GetOptions) (*GitRepoCache, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GitRepoCache), err
}

func (s *gitRepoCacheClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*GitRepoCache, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*GitRepoCache), err
}

func (s *gitRepoCacheClient) Update(o *GitRepoCache) (*GitRepoCache, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GitRepoCache), err
}

func (s *gitRepoCacheClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *gitRepoCacheClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *gitRepoCacheClient) List(opts metav1.ListOptions) (*GitRepoCacheList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GitRepoCacheList), err
}

func (s *gitRepoCacheClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *gitRepoCacheClient) Patch(o *GitRepoCache, data []byte, subresources ...string) (*GitRepoCache, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*GitRepoCache), err
}

func (s *gitRepoCacheClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *gitRepoCacheClient) AddSyncHandler(sync GitRepoCacheHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *gitRepoCacheClient) AddLifecycle(name string, lifecycle GitRepoCacheLifecycle) {
	sync := NewGitRepoCacheLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
