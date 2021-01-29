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

// JobsGetter has a method to return a JobInterface.
// A group's client should implement this interface.
type JobsGetter interface {
	Jobs() JobInterface
}

// JobInterface has methods to work with Job resources.
type JobInterface interface {
	Create(*turtle.Job) (*turtle.Job, error)
	Update(*turtle.Job) (*turtle.Job, error)
	UpdateStatus(*turtle.Job) (*turtle.Job, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*turtle.Job, error)
	List(opts v1.ListOptions) (*turtle.JobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.Job, err error)
	JobExpansion
}

// jobs implements JobInterface
type jobs struct {
	client rest.Interface
}

// newJobs returns a Jobs
func newJobs(c *TurtleClient) *jobs {
	return &jobs{
		client: c.RESTClient(),
	}
}

// Get takes name of the job, and returns the corresponding job object, and an error if there is any.
func (c *jobs) Get(name string, options v1.GetOptions) (result *turtle.Job, err error) {
	result = &turtle.Job{}
	err = c.client.Get().
		Resource("jobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Jobs that match those selectors.
func (c *jobs) List(opts v1.ListOptions) (result *turtle.JobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &turtle.JobList{}
	err = c.client.Get().
		Resource("jobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *jobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("jobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a job and creates it.  Returns the server's representation of the job, and an error, if there is any.
func (c *jobs) Create(job *turtle.Job) (result *turtle.Job, err error) {
	result = &turtle.Job{}
	err = c.client.Post().
		Resource("jobs").
		Body(job).
		Do().
		Into(result)
	return
}

// Update takes the representation of a job and updates it. Returns the server's representation of the job, and an error, if there is any.
func (c *jobs) Update(job *turtle.Job) (result *turtle.Job, err error) {
	result = &turtle.Job{}
	err = c.client.Put().
		Resource("jobs").
		Name(job.Name).
		Body(job).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *jobs) UpdateStatus(job *turtle.Job) (result *turtle.Job, err error) {
	result = &turtle.Job{}
	err = c.client.Put().
		Resource("jobs").
		Name(job.Name).
		SubResource("status").
		Body(job).
		Do().
		Into(result)
	return
}

// Delete takes name of the job and deletes it. Returns an error if one occurs.
func (c *jobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("jobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("jobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched job.
func (c *jobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.Job, err error) {
	result = &turtle.Job{}
	err = c.client.Patch(pt).
		Resource("jobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}