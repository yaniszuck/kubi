// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/ca-gip/kubi/pkg/apis/ca-gip/v1"
	scheme "github.com/ca-gip/kubi/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkPolicyConfigsGetter has a method to return a NetworkPolicyConfigInterface.
// A group's client should implement this interface.
type NetworkPolicyConfigsGetter interface {
	NetworkPolicyConfigs() NetworkPolicyConfigInterface
}

// NetworkPolicyConfigInterface has methods to work with NetworkPolicyConfig resources.
type NetworkPolicyConfigInterface interface {
	Create(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.CreateOptions) (*v1.NetworkPolicyConfig, error)
	Update(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.UpdateOptions) (*v1.NetworkPolicyConfig, error)
	UpdateStatus(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.UpdateOptions) (*v1.NetworkPolicyConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkPolicyConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkPolicyConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicyConfig, err error)
	NetworkPolicyConfigExpansion
}

// networkPolicyConfigs implements NetworkPolicyConfigInterface
type networkPolicyConfigs struct {
	client rest.Interface
}

// newNetworkPolicyConfigs returns a NetworkPolicyConfigs
func newNetworkPolicyConfigs(c *CagipV1Client) *networkPolicyConfigs {
	return &networkPolicyConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the networkPolicyConfig, and returns the corresponding networkPolicyConfig object, and an error if there is any.
func (c *networkPolicyConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkPolicyConfig, err error) {
	result = &v1.NetworkPolicyConfig{}
	err = c.client.Get().
		Resource("networkpolicyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkPolicyConfigs that match those selectors.
func (c *networkPolicyConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkPolicyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkPolicyConfigList{}
	err = c.client.Get().
		Resource("networkpolicyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkPolicyConfigs.
func (c *networkPolicyConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networkpolicyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkPolicyConfig and creates it.  Returns the server's representation of the networkPolicyConfig, and an error, if there is any.
func (c *networkPolicyConfigs) Create(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.CreateOptions) (result *v1.NetworkPolicyConfig, err error) {
	result = &v1.NetworkPolicyConfig{}
	err = c.client.Post().
		Resource("networkpolicyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicyConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkPolicyConfig and updates it. Returns the server's representation of the networkPolicyConfig, and an error, if there is any.
func (c *networkPolicyConfigs) Update(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.UpdateOptions) (result *v1.NetworkPolicyConfig, err error) {
	result = &v1.NetworkPolicyConfig{}
	err = c.client.Put().
		Resource("networkpolicyconfigs").
		Name(networkPolicyConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicyConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkPolicyConfigs) UpdateStatus(ctx context.Context, networkPolicyConfig *v1.NetworkPolicyConfig, opts metav1.UpdateOptions) (result *v1.NetworkPolicyConfig, err error) {
	result = &v1.NetworkPolicyConfig{}
	err = c.client.Put().
		Resource("networkpolicyconfigs").
		Name(networkPolicyConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicyConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkPolicyConfig and deletes it. Returns an error if one occurs.
func (c *networkPolicyConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networkpolicyconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkPolicyConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networkpolicyconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkPolicyConfig.
func (c *networkPolicyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicyConfig, err error) {
	result = &v1.NetworkPolicyConfig{}
	err = c.client.Patch(pt).
		Resource("networkpolicyconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
