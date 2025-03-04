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

package v3

import (
	"context"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ActiveDirectoryProvidersGetter has a method to return a ActiveDirectoryProviderInterface.
// A group's client should implement this interface.
type ActiveDirectoryProvidersGetter interface {
	ActiveDirectoryProviders() ActiveDirectoryProviderInterface
}

// ActiveDirectoryProviderInterface has methods to work with ActiveDirectoryProvider resources.
type ActiveDirectoryProviderInterface interface {
	Create(ctx context.Context, activeDirectoryProvider *v3.ActiveDirectoryProvider, opts v1.CreateOptions) (*v3.ActiveDirectoryProvider, error)
	Update(ctx context.Context, activeDirectoryProvider *v3.ActiveDirectoryProvider, opts v1.UpdateOptions) (*v3.ActiveDirectoryProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ActiveDirectoryProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ActiveDirectoryProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ActiveDirectoryProvider, err error)
	ActiveDirectoryProviderExpansion
}

// activeDirectoryProviders implements ActiveDirectoryProviderInterface
type activeDirectoryProviders struct {
	*gentype.ClientWithList[*v3.ActiveDirectoryProvider, *v3.ActiveDirectoryProviderList]
}

// newActiveDirectoryProviders returns a ActiveDirectoryProviders
func newActiveDirectoryProviders(c *ManagementV3Client) *activeDirectoryProviders {
	return &activeDirectoryProviders{
		gentype.NewClientWithList[*v3.ActiveDirectoryProvider, *v3.ActiveDirectoryProviderList](
			"activedirectoryproviders",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v3.ActiveDirectoryProvider { return &v3.ActiveDirectoryProvider{} },
			func() *v3.ActiveDirectoryProviderList { return &v3.ActiveDirectoryProviderList{} }),
	}
}
