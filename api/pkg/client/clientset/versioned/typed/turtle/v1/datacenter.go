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

package v1

import (
	"time"

	v1 "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle/v1"
	scheme "git.cloud2go.cn/reedition/turtle/api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataCentersGetter has a method to return a DataCenterInterface.
// A group's client should implement this interface.
type DataCentersGetter interface {
	DataCenters() DataCenterInterface
}

// DataCenterInterface has methods to work with DataCenter resources.
type DataCenterInterface interface {
	Create(*v1.DataCenter) (*v1.DataCenter, error)
	Update(*v1.DataCenter) (*v1.DataCenter, error)
	UpdateStatus(*v1.DataCenter) (*v1.DataCenter, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.DataCenter, error)
	List(opts metav1.ListOptions) (*v1.DataCenterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DataCenter, err error)
	DataCenterExpansion
}

// dataCenters implements DataCenterInterface
type dataCenters struct {
	client rest.Interface
}

// newDataCenters returns a DataCenters
func newDataCenters(c *TurtleV1Client) *dataCenters {
	return &dataCenters{
		client: c.RESTClient(),
	}
}

// Get takes name of the dataCenter, and returns the corresponding dataCenter object, and an error if there is any.
func (c *dataCenters) Get(name string, options metav1.GetOptions) (result *v1.DataCenter, err error) {
	result = &v1.DataCenter{}
	err = c.client.Get().
		Resource("datacenters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataCenters that match those selectors.
func (c *dataCenters) List(opts metav1.ListOptions) (result *v1.DataCenterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DataCenterList{}
	err = c.client.Get().
		Resource("datacenters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataCenters.
func (c *dataCenters) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("datacenters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataCenter and creates it.  Returns the server's representation of the dataCenter, and an error, if there is any.
func (c *dataCenters) Create(dataCenter *v1.DataCenter) (result *v1.DataCenter, err error) {
	result = &v1.DataCenter{}
	err = c.client.Post().
		Resource("datacenters").
		Body(dataCenter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataCenter and updates it. Returns the server's representation of the dataCenter, and an error, if there is any.
func (c *dataCenters) Update(dataCenter *v1.DataCenter) (result *v1.DataCenter, err error) {
	result = &v1.DataCenter{}
	err = c.client.Put().
		Resource("datacenters").
		Name(dataCenter.Name).
		Body(dataCenter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataCenters) UpdateStatus(dataCenter *v1.DataCenter) (result *v1.DataCenter, err error) {
	result = &v1.DataCenter{}
	err = c.client.Put().
		Resource("datacenters").
		Name(dataCenter.Name).
		SubResource("status").
		Body(dataCenter).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataCenter and deletes it. Returns an error if one occurs.
func (c *dataCenters) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("datacenters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataCenters) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("datacenters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataCenter.
func (c *dataCenters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DataCenter, err error) {
	result = &v1.DataCenter{}
	err = c.client.Patch(pt).
		Resource("datacenters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
