// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "open-cluster-management.io/api/addon/v1alpha1"
	scheme "open-cluster-management.io/api/client/addon/clientset/versioned/scheme"
)

// AddOnDeploymentConfigsGetter has a method to return a AddOnDeploymentConfigInterface.
// A group's client should implement this interface.
type AddOnDeploymentConfigsGetter interface {
	AddOnDeploymentConfigs() AddOnDeploymentConfigInterface
}

// AddOnDeploymentConfigInterface has methods to work with AddOnDeploymentConfig resources.
type AddOnDeploymentConfigInterface interface {
	Create(ctx context.Context, addOnDeploymentConfig *v1alpha1.AddOnDeploymentConfig, opts v1.CreateOptions) (*v1alpha1.AddOnDeploymentConfig, error)
	Update(ctx context.Context, addOnDeploymentConfig *v1alpha1.AddOnDeploymentConfig, opts v1.UpdateOptions) (*v1alpha1.AddOnDeploymentConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AddOnDeploymentConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AddOnDeploymentConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AddOnDeploymentConfig, err error)
	AddOnDeploymentConfigExpansion
}

// addOnDeploymentConfigs implements AddOnDeploymentConfigInterface
type addOnDeploymentConfigs struct {
	client rest.Interface
}

// newAddOnDeploymentConfigs returns a AddOnDeploymentConfigs
func newAddOnDeploymentConfigs(c *AddonV1alpha1Client) *addOnDeploymentConfigs {
	return &addOnDeploymentConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the addOnDeploymentConfig, and returns the corresponding addOnDeploymentConfig object, and an error if there is any.
func (c *addOnDeploymentConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AddOnDeploymentConfig, err error) {
	result = &v1alpha1.AddOnDeploymentConfig{}
	err = c.client.Get().
		Resource("addondeploymentconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AddOnDeploymentConfigs that match those selectors.
func (c *addOnDeploymentConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddOnDeploymentConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AddOnDeploymentConfigList{}
	err = c.client.Get().
		Resource("addondeploymentconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested addOnDeploymentConfigs.
func (c *addOnDeploymentConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("addondeploymentconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a addOnDeploymentConfig and creates it.  Returns the server's representation of the addOnDeploymentConfig, and an error, if there is any.
func (c *addOnDeploymentConfigs) Create(ctx context.Context, addOnDeploymentConfig *v1alpha1.AddOnDeploymentConfig, opts v1.CreateOptions) (result *v1alpha1.AddOnDeploymentConfig, err error) {
	result = &v1alpha1.AddOnDeploymentConfig{}
	err = c.client.Post().
		Resource("addondeploymentconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(addOnDeploymentConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a addOnDeploymentConfig and updates it. Returns the server's representation of the addOnDeploymentConfig, and an error, if there is any.
func (c *addOnDeploymentConfigs) Update(ctx context.Context, addOnDeploymentConfig *v1alpha1.AddOnDeploymentConfig, opts v1.UpdateOptions) (result *v1alpha1.AddOnDeploymentConfig, err error) {
	result = &v1alpha1.AddOnDeploymentConfig{}
	err = c.client.Put().
		Resource("addondeploymentconfigs").
		Name(addOnDeploymentConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(addOnDeploymentConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the addOnDeploymentConfig and deletes it. Returns an error if one occurs.
func (c *addOnDeploymentConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("addondeploymentconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *addOnDeploymentConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("addondeploymentconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched addOnDeploymentConfig.
func (c *addOnDeploymentConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AddOnDeploymentConfig, err error) {
	result = &v1alpha1.AddOnDeploymentConfig{}
	err = c.client.Patch(pt).
		Resource("addondeploymentconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
