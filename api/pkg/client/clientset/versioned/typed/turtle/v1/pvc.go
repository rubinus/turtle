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

// PVCsGetter has a method to return a PVCInterface.
// A group's client should implement this interface.
type PVCsGetter interface {
	PVCs() PVCInterface
}

// PVCInterface has methods to work with PVC resources.
type PVCInterface interface {
	Create(*v1.PVC) (*v1.PVC, error)
	Update(*v1.PVC) (*v1.PVC, error)
	UpdateStatus(*v1.PVC) (*v1.PVC, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PVC, error)
	List(opts metav1.ListOptions) (*v1.PVCList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PVC, err error)
	PVCExpansion
}

// pVCs implements PVCInterface
type pVCs struct {
	client rest.Interface
}

// newPVCs returns a PVCs
func newPVCs(c *TurtleV1Client) *pVCs {
	return &pVCs{
		client: c.RESTClient(),
	}
}

// Get takes name of the pVC, and returns the corresponding pVC object, and an error if there is any.
func (c *pVCs) Get(name string, options metav1.GetOptions) (result *v1.PVC, err error) {
	result = &v1.PVC{}
	err = c.client.Get().
		Resource("pvcs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PVCs that match those selectors.
func (c *pVCs) List(opts metav1.ListOptions) (result *v1.PVCList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PVCList{}
	err = c.client.Get().
		Resource("pvcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pVCs.
func (c *pVCs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pvcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pVC and creates it.  Returns the server's representation of the pVC, and an error, if there is any.
func (c *pVCs) Create(pVC *v1.PVC) (result *v1.PVC, err error) {
	result = &v1.PVC{}
	err = c.client.Post().
		Resource("pvcs").
		Body(pVC).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pVC and updates it. Returns the server's representation of the pVC, and an error, if there is any.
func (c *pVCs) Update(pVC *v1.PVC) (result *v1.PVC, err error) {
	result = &v1.PVC{}
	err = c.client.Put().
		Resource("pvcs").
		Name(pVC.Name).
		Body(pVC).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pVCs) UpdateStatus(pVC *v1.PVC) (result *v1.PVC, err error) {
	result = &v1.PVC{}
	err = c.client.Put().
		Resource("pvcs").
		Name(pVC.Name).
		SubResource("status").
		Body(pVC).
		Do().
		Into(result)
	return
}

// Delete takes name of the pVC and deletes it. Returns an error if one occurs.
func (c *pVCs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pvcs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pVCs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pvcs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pVC.
func (c *pVCs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PVC, err error) {
	result = &v1.PVC{}
	err = c.client.Patch(pt).
		Resource("pvcs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}