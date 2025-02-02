/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/fire-ant/consularis/pkg/apis/consularis.io/v1alpha1"
	scheme "github.com/fire-ant/consularis/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsulObjectsGetter has a method to return a ConsulObjectInterface.
// A group's client should implement this interface.
type ConsulObjectsGetter interface {
	ConsulObjects(namespace string) ConsulObjectInterface
}

// ConsulObjectInterface has methods to work with ConsulObject resources.
type ConsulObjectInterface interface {
	Create(*v1alpha1.ConsulObject) (*v1alpha1.ConsulObject, error)
	Update(*v1alpha1.ConsulObject) (*v1alpha1.ConsulObject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConsulObject, error)
	List(opts v1.ListOptions) (*v1alpha1.ConsulObjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConsulObject, err error)
	ConsulObjectExpansion
}

// consulObjects implements ConsulObjectInterface
type consulObjects struct {
	client rest.Interface
	ns     string
}

// newConsulObjects returns a ConsulObjects
func newConsulObjects(c *ConsularisV1alpha1Client, namespace string) *consulObjects {
	return &consulObjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the consulObject, and returns the corresponding consulObject object, and an error if there is any.
func (c *consulObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.ConsulObject, err error) {
	result = &v1alpha1.ConsulObject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consulobjects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsulObjects that match those selectors.
func (c *consulObjects) List(opts v1.ListOptions) (result *v1alpha1.ConsulObjectList, err error) {
	result = &v1alpha1.ConsulObjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consulobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consulObjects.
func (c *consulObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("consulobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a consulObject and creates it.  Returns the server's representation of the consulObject, and an error, if there is any.
func (c *consulObjects) Create(consulObject *v1alpha1.ConsulObject) (result *v1alpha1.ConsulObject, err error) {
	result = &v1alpha1.ConsulObject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("consulobjects").
		Body(consulObject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a consulObject and updates it. Returns the server's representation of the consulObject, and an error, if there is any.
func (c *consulObjects) Update(consulObject *v1alpha1.ConsulObject) (result *v1alpha1.ConsulObject, err error) {
	result = &v1alpha1.ConsulObject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("consulobjects").
		Name(consulObject.Name).
		Body(consulObject).
		Do().
		Into(result)
	return
}

// Delete takes name of the consulObject and deletes it. Returns an error if one occurs.
func (c *consulObjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consulobjects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consulObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consulobjects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched consulObject.
func (c *consulObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConsulObject, err error) {
	result = &v1alpha1.ConsulObject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("consulobjects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
