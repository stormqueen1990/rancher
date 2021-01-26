/*
Copyright 2021 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ProjectMonitorGraphHandler func(string, *v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error)

type ProjectMonitorGraphController interface {
	generic.ControllerMeta
	ProjectMonitorGraphClient

	OnChange(ctx context.Context, name string, sync ProjectMonitorGraphHandler)
	OnRemove(ctx context.Context, name string, sync ProjectMonitorGraphHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ProjectMonitorGraphCache
}

type ProjectMonitorGraphClient interface {
	Create(*v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error)
	Update(*v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectMonitorGraph, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ProjectMonitorGraphList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ProjectMonitorGraph, err error)
}

type ProjectMonitorGraphCache interface {
	Get(namespace, name string) (*v3.ProjectMonitorGraph, error)
	List(namespace string, selector labels.Selector) ([]*v3.ProjectMonitorGraph, error)

	AddIndexer(indexName string, indexer ProjectMonitorGraphIndexer)
	GetByIndex(indexName, key string) ([]*v3.ProjectMonitorGraph, error)
}

type ProjectMonitorGraphIndexer func(obj *v3.ProjectMonitorGraph) ([]string, error)

type projectMonitorGraphController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewProjectMonitorGraphController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ProjectMonitorGraphController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &projectMonitorGraphController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromProjectMonitorGraphHandlerToHandler(sync ProjectMonitorGraphHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ProjectMonitorGraph
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ProjectMonitorGraph))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *projectMonitorGraphController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ProjectMonitorGraph))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateProjectMonitorGraphDeepCopyOnChange(client ProjectMonitorGraphClient, obj *v3.ProjectMonitorGraph, handler func(obj *v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error)) (*v3.ProjectMonitorGraph, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *projectMonitorGraphController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *projectMonitorGraphController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *projectMonitorGraphController) OnChange(ctx context.Context, name string, sync ProjectMonitorGraphHandler) {
	c.AddGenericHandler(ctx, name, FromProjectMonitorGraphHandlerToHandler(sync))
}

func (c *projectMonitorGraphController) OnRemove(ctx context.Context, name string, sync ProjectMonitorGraphHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromProjectMonitorGraphHandlerToHandler(sync)))
}

func (c *projectMonitorGraphController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *projectMonitorGraphController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *projectMonitorGraphController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *projectMonitorGraphController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *projectMonitorGraphController) Cache() ProjectMonitorGraphCache {
	return &projectMonitorGraphCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *projectMonitorGraphController) Create(obj *v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error) {
	result := &v3.ProjectMonitorGraph{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *projectMonitorGraphController) Update(obj *v3.ProjectMonitorGraph) (*v3.ProjectMonitorGraph, error) {
	result := &v3.ProjectMonitorGraph{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectMonitorGraphController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *projectMonitorGraphController) Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectMonitorGraph, error) {
	result := &v3.ProjectMonitorGraph{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *projectMonitorGraphController) List(namespace string, opts metav1.ListOptions) (*v3.ProjectMonitorGraphList, error) {
	result := &v3.ProjectMonitorGraphList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *projectMonitorGraphController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *projectMonitorGraphController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ProjectMonitorGraph, error) {
	result := &v3.ProjectMonitorGraph{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type projectMonitorGraphCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *projectMonitorGraphCache) Get(namespace, name string) (*v3.ProjectMonitorGraph, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ProjectMonitorGraph), nil
}

func (c *projectMonitorGraphCache) List(namespace string, selector labels.Selector) (ret []*v3.ProjectMonitorGraph, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ProjectMonitorGraph))
	})

	return ret, err
}

func (c *projectMonitorGraphCache) AddIndexer(indexName string, indexer ProjectMonitorGraphIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ProjectMonitorGraph))
		},
	}))
}

func (c *projectMonitorGraphCache) GetByIndex(indexName, key string) (result []*v3.ProjectMonitorGraph, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ProjectMonitorGraph, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ProjectMonitorGraph))
	}
	return result, nil
}
