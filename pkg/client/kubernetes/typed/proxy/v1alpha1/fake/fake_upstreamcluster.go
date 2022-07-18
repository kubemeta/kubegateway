/*
Copyright 2022 ByteDance and its affiliates.

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

package fake

import (
	"context"

	v1alpha1 "github.com/kubewharf/kubegateway/pkg/apis/proxy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpstreamClusters implements UpstreamClusterInterface
type FakeUpstreamClusters struct {
	Fake *FakeProxyV1alpha1
}

var upstreamclustersResource = schema.GroupVersionResource{Group: "proxy.kubegateway.io", Version: "v1alpha1", Resource: "upstreamclusters"}

var upstreamclustersKind = schema.GroupVersionKind{Group: "proxy.kubegateway.io", Version: "v1alpha1", Kind: "UpstreamCluster"}

// Get takes name of the upstreamCluster, and returns the corresponding upstreamCluster object, and an error if there is any.
func (c *FakeUpstreamClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UpstreamCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(upstreamclustersResource, name), &v1alpha1.UpstreamCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpstreamCluster), err
}

// List takes label and field selectors, and returns the list of UpstreamClusters that match those selectors.
func (c *FakeUpstreamClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UpstreamClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(upstreamclustersResource, upstreamclustersKind, opts), &v1alpha1.UpstreamClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UpstreamClusterList{ListMeta: obj.(*v1alpha1.UpstreamClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.UpstreamClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upstreamClusters.
func (c *FakeUpstreamClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(upstreamclustersResource, opts))
}

// Create takes the representation of a upstreamCluster and creates it.  Returns the server's representation of the upstreamCluster, and an error, if there is any.
func (c *FakeUpstreamClusters) Create(ctx context.Context, upstreamCluster *v1alpha1.UpstreamCluster, opts v1.CreateOptions) (result *v1alpha1.UpstreamCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(upstreamclustersResource, upstreamCluster), &v1alpha1.UpstreamCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpstreamCluster), err
}

// Update takes the representation of a upstreamCluster and updates it. Returns the server's representation of the upstreamCluster, and an error, if there is any.
func (c *FakeUpstreamClusters) Update(ctx context.Context, upstreamCluster *v1alpha1.UpstreamCluster, opts v1.UpdateOptions) (result *v1alpha1.UpstreamCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(upstreamclustersResource, upstreamCluster), &v1alpha1.UpstreamCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpstreamCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUpstreamClusters) UpdateStatus(ctx context.Context, upstreamCluster *v1alpha1.UpstreamCluster, opts v1.UpdateOptions) (*v1alpha1.UpstreamCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(upstreamclustersResource, "status", upstreamCluster), &v1alpha1.UpstreamCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpstreamCluster), err
}

// Delete takes name of the upstreamCluster and deletes it. Returns an error if one occurs.
func (c *FakeUpstreamClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(upstreamclustersResource, name), &v1alpha1.UpstreamCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpstreamClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(upstreamclustersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UpstreamClusterList{})
	return err
}

// Patch applies the patch and returns the patched upstreamCluster.
func (c *FakeUpstreamClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UpstreamCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(upstreamclustersResource, name, pt, data, subresources...), &v1alpha1.UpstreamCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpstreamCluster), err
}
