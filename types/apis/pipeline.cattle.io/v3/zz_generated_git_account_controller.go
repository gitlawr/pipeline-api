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
	GitAccountGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GitAccount",
	}
	GitAccountResource = metav1.APIResource{
		Name:         "gitaccounts",
		SingularName: "gitaccount",
		Namespaced:   false,
		Kind:         GitAccountGroupVersionKind.Kind,
	}
)

type GitAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitAccount
}

type GitAccountHandlerFunc func(key string, obj *GitAccount) error

type GitAccountLister interface {
	List(namespace string, selector labels.Selector) (ret []*GitAccount, err error)
	Get(namespace, name string) (*GitAccount, error)
}

type GitAccountController interface {
	Informer() cache.SharedIndexInformer
	Lister() GitAccountLister
	AddHandler(handler GitAccountHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GitAccountInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*GitAccount) (*GitAccount, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*GitAccount, error)
	Get(name string, opts metav1.GetOptions) (*GitAccount, error)
	Update(*GitAccount) (*GitAccount, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GitAccountList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GitAccountController
	AddSyncHandler(sync GitAccountHandlerFunc)
	AddLifecycle(name string, lifecycle GitAccountLifecycle)
}

type gitAccountLister struct {
	controller *gitAccountController
}

func (l *gitAccountLister) List(namespace string, selector labels.Selector) (ret []*GitAccount, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GitAccount))
	})
	return
}

func (l *gitAccountLister) Get(namespace, name string) (*GitAccount, error) {
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
			Group:    GitAccountGroupVersionKind.Group,
			Resource: "gitAccount",
		}, name)
	}
	return obj.(*GitAccount), nil
}

type gitAccountController struct {
	controller.GenericController
}

func (c *gitAccountController) Lister() GitAccountLister {
	return &gitAccountLister{
		controller: c,
	}
}

func (c *gitAccountController) AddHandler(handler GitAccountHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*GitAccount))
	})
}

type gitAccountFactory struct {
}

func (c gitAccountFactory) Object() runtime.Object {
	return &GitAccount{}
}

func (c gitAccountFactory) List() runtime.Object {
	return &GitAccountList{}
}

func (s *gitAccountClient) Controller() GitAccountController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.gitAccountControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GitAccountGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &gitAccountController{
		GenericController: genericController,
	}

	s.client.gitAccountControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type gitAccountClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   GitAccountController
}

func (s *gitAccountClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *gitAccountClient) Create(o *GitAccount) (*GitAccount, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GitAccount), err
}

func (s *gitAccountClient) Get(name string, opts metav1.GetOptions) (*GitAccount, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GitAccount), err
}

func (s *gitAccountClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*GitAccount, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*GitAccount), err
}

func (s *gitAccountClient) Update(o *GitAccount) (*GitAccount, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GitAccount), err
}

func (s *gitAccountClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *gitAccountClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *gitAccountClient) List(opts metav1.ListOptions) (*GitAccountList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GitAccountList), err
}

func (s *gitAccountClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *gitAccountClient) Patch(o *GitAccount, data []byte, subresources ...string) (*GitAccount, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*GitAccount), err
}

func (s *gitAccountClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *gitAccountClient) AddSyncHandler(sync GitAccountHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *gitAccountClient) AddLifecycle(name string, lifecycle GitAccountLifecycle) {
	sync := NewGitAccountLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
