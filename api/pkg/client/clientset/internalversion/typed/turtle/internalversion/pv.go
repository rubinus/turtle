/*
Copyright The Kubernetes Authors.

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

package internalversion

import (
	"time"

	turtle "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle"
	scheme "git.cloud2go.cn/reedition/turtle/api/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PVsGetter has a method to return a PVInterface.
// A group's client should implement this interface.
type PVsGetter interface {
	PVs() PVInterface
}

// PVInterface has methods to work with PV resources.
type PVInterface interface {
	Create(*turtle.PV) (*turtle.PV, error)
	Update(*turtle.PV) (*turtle.PV, error)
	UpdateStatus(*turtle.PV) (*turtle.PV, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*turtle.PV, error)
	List(opts v1.ListOptions) (*turtle.PVList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.PV, err error)
	PVExpansion
}

// pVs implements PVInterface
type pVs struct {
	client rest.Interface
}

// newPVs returns a PVs
func newPVs(c *TurtleClient) *pVs {
	return &pVs{
		client: c.RESTClient(),
	}
}

// Get takes name of the pV, and returns the corresponding pV object, and an error if there is any.
func (c *pVs) Get(name string, options v1.GetOptions) (result *turtle.PV, err error) {
	result = &turtle.PV{}
	err = c.client.Get().
		Resource("pvs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PVs that match those selectors.
func (c *pVs) List(opts v1.ListOptions) (result *turtle.PVList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &turtle.PVList{}
	err = c.client.Get().
		Resource("pvs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pVs.
func (c *pVs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pvs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pV and creates it.  Returns the server's representation of the pV, and an error, if there is any.
func (c *pVs) Create(pV *turtle.PV) (result *turtle.PV, err error) {
	result = &turtle.PV{}
	err = c.client.Post().
		Resource("pvs").
		Body(pV).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pV and updates it. Returns the server's representation of the pV, and an error, if there is any.
func (c *pVs) Update(pV *turtle.PV) (result *turtle.PV, err error) {
	result = &turtle.PV{}
	err = c.client.Put().
		Resource("pvs").
		Name(pV.Name).
		Body(pV).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pVs) UpdateStatus(pV *turtle.PV) (result *turtle.PV, err error) {
	result = &turtle.PV{}
	err = c.client.Put().
		Resource("pvs").
		Name(pV.Name).
		SubResource("status").
		Body(pV).
		Do().
		Into(result)
	return
}

// Delete takes name of the pV and deletes it. Returns an error if one occurs.
func (c *pVs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pvs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pVs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pvs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pV.
func (c *pVs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.PV, err error) {
	result = &turtle.PV{}
	err = c.client.Patch(pt).
		Resource("pvs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
