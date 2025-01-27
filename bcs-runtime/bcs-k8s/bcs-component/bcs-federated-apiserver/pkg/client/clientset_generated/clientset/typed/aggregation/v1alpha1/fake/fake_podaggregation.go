/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodAggregations implements PodAggregationInterface
type FakePodAggregations struct {
	Fake *FakeAggregationV1alpha1
	ns   string
}

var podaggregationsResource = schema.GroupVersionResource{Group: "aggregation", Version: "v1alpha1", Resource: "podaggregations"}

var podaggregationsKind = schema.GroupVersionKind{Group: "aggregation", Version: "v1alpha1", Kind: "PodAggregation"}

// Get takes name of the podAggregation, and returns the corresponding podAggregation object, and an error if there is any.
func (c *FakePodAggregations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podaggregationsResource, c.ns, name), &v1alpha1.PodAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAggregation), err
}

// List takes label and field selectors, and returns the list of PodAggregations that match those selectors.
func (c *FakePodAggregations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodAggregationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podaggregationsResource, podaggregationsKind, c.ns, opts), &v1alpha1.PodAggregationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodAggregationList{ListMeta: obj.(*v1alpha1.PodAggregationList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodAggregationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podAggregations.
func (c *FakePodAggregations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podaggregationsResource, c.ns, opts))

}

// Create takes the representation of a podAggregation and creates it.  Returns the server's representation of the podAggregation, and an error, if there is any.
func (c *FakePodAggregations) Create(ctx context.Context, podAggregation *v1alpha1.PodAggregation, opts v1.CreateOptions) (result *v1alpha1.PodAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podaggregationsResource, c.ns, podAggregation), &v1alpha1.PodAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAggregation), err
}

// Update takes the representation of a podAggregation and updates it. Returns the server's representation of the podAggregation, and an error, if there is any.
func (c *FakePodAggregations) Update(ctx context.Context, podAggregation *v1alpha1.PodAggregation, opts v1.UpdateOptions) (result *v1alpha1.PodAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podaggregationsResource, c.ns, podAggregation), &v1alpha1.PodAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAggregation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodAggregations) UpdateStatus(ctx context.Context, podAggregation *v1alpha1.PodAggregation, opts v1.UpdateOptions) (*v1alpha1.PodAggregation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podaggregationsResource, "status", c.ns, podAggregation), &v1alpha1.PodAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAggregation), err
}

// Delete takes name of the podAggregation and deletes it. Returns an error if one occurs.
func (c *FakePodAggregations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podaggregationsResource, c.ns, name), &v1alpha1.PodAggregation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodAggregations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podaggregationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodAggregationList{})
	return err
}

// Patch applies the patch and returns the patched podAggregation.
func (c *FakePodAggregations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodAggregation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podaggregationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PodAggregation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodAggregation), err
}
