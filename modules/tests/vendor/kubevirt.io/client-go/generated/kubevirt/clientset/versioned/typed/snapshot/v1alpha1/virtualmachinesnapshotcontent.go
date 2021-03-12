/*
Copyright 2021 The KubeVirt Authors.

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "kubevirt.io/client-go/apis/snapshot/v1alpha1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachineSnapshotContentsGetter has a method to return a VirtualMachineSnapshotContentInterface.
// A group's client should implement this interface.
type VirtualMachineSnapshotContentsGetter interface {
	VirtualMachineSnapshotContents(namespace string) VirtualMachineSnapshotContentInterface
}

// VirtualMachineSnapshotContentInterface has methods to work with VirtualMachineSnapshotContent resources.
type VirtualMachineSnapshotContentInterface interface {
	Create(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.CreateOptions) (*v1alpha1.VirtualMachineSnapshotContent, error)
	Update(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineSnapshotContent, error)
	UpdateStatus(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineSnapshotContent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualMachineSnapshotContent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualMachineSnapshotContentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineSnapshotContent, err error)
	VirtualMachineSnapshotContentExpansion
}

// virtualMachineSnapshotContents implements VirtualMachineSnapshotContentInterface
type virtualMachineSnapshotContents struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineSnapshotContents returns a VirtualMachineSnapshotContents
func newVirtualMachineSnapshotContents(c *SnapshotV1alpha1Client, namespace string) *virtualMachineSnapshotContents {
	return &virtualMachineSnapshotContents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineSnapshotContent, and returns the corresponding virtualMachineSnapshotContent object, and an error if there is any.
func (c *virtualMachineSnapshotContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineSnapshotContents that match those selectors.
func (c *virtualMachineSnapshotContents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineSnapshotContentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineSnapshotContentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineSnapshotContents.
func (c *virtualMachineSnapshotContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineSnapshotContent and creates it.  Returns the server's representation of the virtualMachineSnapshotContent, and an error, if there is any.
func (c *virtualMachineSnapshotContents) Create(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineSnapshotContent and updates it. Returns the server's representation of the virtualMachineSnapshotContent, and an error, if there is any.
func (c *virtualMachineSnapshotContents) Update(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(virtualMachineSnapshotContent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineSnapshotContents) UpdateStatus(ctx context.Context, virtualMachineSnapshotContent *v1alpha1.VirtualMachineSnapshotContent, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(virtualMachineSnapshotContent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineSnapshotContent and deletes it. Returns an error if one occurs.
func (c *virtualMachineSnapshotContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineSnapshotContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineSnapshotContent.
func (c *virtualMachineSnapshotContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineSnapshotContent, err error) {
	result = &v1alpha1.VirtualMachineSnapshotContent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinesnapshotcontents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
