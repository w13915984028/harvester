/*
Copyright 2025 Rancher Labs, Inc.

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

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureADProviders implements AzureADProviderInterface
type FakeAzureADProviders struct {
	Fake *FakeManagementV3
}

var azureadprovidersResource = v3.SchemeGroupVersion.WithResource("azureadproviders")

var azureadprovidersKind = v3.SchemeGroupVersion.WithKind("AzureADProvider")

// Get takes name of the azureADProvider, and returns the corresponding azureADProvider object, and an error if there is any.
func (c *FakeAzureADProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.AzureADProvider, err error) {
	emptyResult := &v3.AzureADProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(azureadprovidersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.AzureADProvider), err
}

// List takes label and field selectors, and returns the list of AzureADProviders that match those selectors.
func (c *FakeAzureADProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.AzureADProviderList, err error) {
	emptyResult := &v3.AzureADProviderList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(azureadprovidersResource, azureadprovidersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.AzureADProviderList{ListMeta: obj.(*v3.AzureADProviderList).ListMeta}
	for _, item := range obj.(*v3.AzureADProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureADProviders.
func (c *FakeAzureADProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(azureadprovidersResource, opts))
}

// Create takes the representation of a azureADProvider and creates it.  Returns the server's representation of the azureADProvider, and an error, if there is any.
func (c *FakeAzureADProviders) Create(ctx context.Context, azureADProvider *v3.AzureADProvider, opts v1.CreateOptions) (result *v3.AzureADProvider, err error) {
	emptyResult := &v3.AzureADProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(azureadprovidersResource, azureADProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.AzureADProvider), err
}

// Update takes the representation of a azureADProvider and updates it. Returns the server's representation of the azureADProvider, and an error, if there is any.
func (c *FakeAzureADProviders) Update(ctx context.Context, azureADProvider *v3.AzureADProvider, opts v1.UpdateOptions) (result *v3.AzureADProvider, err error) {
	emptyResult := &v3.AzureADProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(azureadprovidersResource, azureADProvider, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.AzureADProvider), err
}

// Delete takes name of the azureADProvider and deletes it. Returns an error if one occurs.
func (c *FakeAzureADProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(azureadprovidersResource, name, opts), &v3.AzureADProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureADProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(azureadprovidersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.AzureADProviderList{})
	return err
}

// Patch applies the patch and returns the patched azureADProvider.
func (c *FakeAzureADProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AzureADProvider, err error) {
	emptyResult := &v3.AzureADProvider{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(azureadprovidersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.AzureADProvider), err
}
