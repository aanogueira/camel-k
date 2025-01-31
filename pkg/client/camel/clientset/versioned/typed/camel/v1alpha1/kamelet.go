/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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
	"context"
	"time"

	v1alpha1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1alpha1"
	scheme "github.com/aanogueira/camel-k/pkg/client/camel/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KameletsGetter has a method to return a KameletInterface.
// A group's client should implement this interface.
type KameletsGetter interface {
	Kamelets(namespace string) KameletInterface
}

// KameletInterface has methods to work with Kamelet resources.
type KameletInterface interface {
	Create(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.CreateOptions) (*v1alpha1.Kamelet, error)
	Update(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.UpdateOptions) (*v1alpha1.Kamelet, error)
	UpdateStatus(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.UpdateOptions) (*v1alpha1.Kamelet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Kamelet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KameletList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Kamelet, err error)
	KameletExpansion
}

// kamelets implements KameletInterface
type kamelets struct {
	client rest.Interface
	ns     string
}

// newKamelets returns a Kamelets
func newKamelets(c *CamelV1alpha1Client, namespace string) *kamelets {
	return &kamelets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kamelet, and returns the corresponding kamelet object, and an error if there is any.
func (c *kamelets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Kamelet, err error) {
	result = &v1alpha1.Kamelet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kamelets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Kamelets that match those selectors.
func (c *kamelets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KameletList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KameletList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kamelets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kamelets.
func (c *kamelets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kamelets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kamelet and creates it.  Returns the server's representation of the kamelet, and an error, if there is any.
func (c *kamelets) Create(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.CreateOptions) (result *v1alpha1.Kamelet, err error) {
	result = &v1alpha1.Kamelet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kamelets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kamelet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kamelet and updates it. Returns the server's representation of the kamelet, and an error, if there is any.
func (c *kamelets) Update(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.UpdateOptions) (result *v1alpha1.Kamelet, err error) {
	result = &v1alpha1.Kamelet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kamelets").
		Name(kamelet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kamelet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kamelets) UpdateStatus(ctx context.Context, kamelet *v1alpha1.Kamelet, opts v1.UpdateOptions) (result *v1alpha1.Kamelet, err error) {
	result = &v1alpha1.Kamelet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kamelets").
		Name(kamelet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kamelet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kamelet and deletes it. Returns an error if one occurs.
func (c *kamelets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kamelets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kamelets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kamelets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kamelet.
func (c *kamelets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Kamelet, err error) {
	result = &v1alpha1.Kamelet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kamelets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
