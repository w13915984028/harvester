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

package v1beta1

import (
	"context"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ResourceQuotasGetter has a method to return a ResourceQuotaInterface.
// A group's client should implement this interface.
type ResourceQuotasGetter interface {
	ResourceQuotas(namespace string) ResourceQuotaInterface
}

// ResourceQuotaInterface has methods to work with ResourceQuota resources.
type ResourceQuotaInterface interface {
	Create(ctx context.Context, resourceQuota *v1beta1.ResourceQuota, opts v1.CreateOptions) (*v1beta1.ResourceQuota, error)
	Update(ctx context.Context, resourceQuota *v1beta1.ResourceQuota, opts v1.UpdateOptions) (*v1beta1.ResourceQuota, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, resourceQuota *v1beta1.ResourceQuota, opts v1.UpdateOptions) (*v1beta1.ResourceQuota, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ResourceQuota, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ResourceQuotaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ResourceQuota, err error)
	ResourceQuotaExpansion
}

// resourceQuotas implements ResourceQuotaInterface
type resourceQuotas struct {
	*gentype.ClientWithList[*v1beta1.ResourceQuota, *v1beta1.ResourceQuotaList]
}

// newResourceQuotas returns a ResourceQuotas
func newResourceQuotas(c *HarvesterhciV1beta1Client, namespace string) *resourceQuotas {
	return &resourceQuotas{
		gentype.NewClientWithList[*v1beta1.ResourceQuota, *v1beta1.ResourceQuotaList](
			"resourcequotas",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.ResourceQuota { return &v1beta1.ResourceQuota{} },
			func() *v1beta1.ResourceQuotaList { return &v1beta1.ResourceQuotaList{} }),
	}
}
