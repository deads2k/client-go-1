package v1

import (
	scheme "github.com/openshift/client-go/template/clientset/scheme"
	v1 "github.com/openshift/origin/pkg/template/apis/template/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BrokerTemplateInstancesGetter has a method to return a BrokerTemplateInstanceInterface.
// A group's client should implement this interface.
type BrokerTemplateInstancesGetter interface {
	BrokerTemplateInstances() BrokerTemplateInstanceInterface
}

// BrokerTemplateInstanceInterface has methods to work with BrokerTemplateInstance resources.
type BrokerTemplateInstanceInterface interface {
	Create(*v1.BrokerTemplateInstance) (*v1.BrokerTemplateInstance, error)
	Update(*v1.BrokerTemplateInstance) (*v1.BrokerTemplateInstance, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.BrokerTemplateInstance, error)
	List(opts meta_v1.ListOptions) (*v1.BrokerTemplateInstanceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BrokerTemplateInstance, err error)
	BrokerTemplateInstanceExpansion
}

// brokerTemplateInstances implements BrokerTemplateInstanceInterface
type brokerTemplateInstances struct {
	client rest.Interface
}

// newBrokerTemplateInstances returns a BrokerTemplateInstances
func newBrokerTemplateInstances(c *TemplateV1Client) *brokerTemplateInstances {
	return &brokerTemplateInstances{
		client: c.RESTClient(),
	}
}

// Get takes name of the brokerTemplateInstance, and returns the corresponding brokerTemplateInstance object, and an error if there is any.
func (c *brokerTemplateInstances) Get(name string, options meta_v1.GetOptions) (result *v1.BrokerTemplateInstance, err error) {
	result = &v1.BrokerTemplateInstance{}
	err = c.client.Get().
		Resource("brokertemplateinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BrokerTemplateInstances that match those selectors.
func (c *brokerTemplateInstances) List(opts meta_v1.ListOptions) (result *v1.BrokerTemplateInstanceList, err error) {
	result = &v1.BrokerTemplateInstanceList{}
	err = c.client.Get().
		Resource("brokertemplateinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested brokerTemplateInstances.
func (c *brokerTemplateInstances) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("brokertemplateinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a brokerTemplateInstance and creates it.  Returns the server's representation of the brokerTemplateInstance, and an error, if there is any.
func (c *brokerTemplateInstances) Create(brokerTemplateInstance *v1.BrokerTemplateInstance) (result *v1.BrokerTemplateInstance, err error) {
	result = &v1.BrokerTemplateInstance{}
	err = c.client.Post().
		Resource("brokertemplateinstances").
		Body(brokerTemplateInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a brokerTemplateInstance and updates it. Returns the server's representation of the brokerTemplateInstance, and an error, if there is any.
func (c *brokerTemplateInstances) Update(brokerTemplateInstance *v1.BrokerTemplateInstance) (result *v1.BrokerTemplateInstance, err error) {
	result = &v1.BrokerTemplateInstance{}
	err = c.client.Put().
		Resource("brokertemplateinstances").
		Name(brokerTemplateInstance.Name).
		Body(brokerTemplateInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the brokerTemplateInstance and deletes it. Returns an error if one occurs.
func (c *brokerTemplateInstances) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("brokertemplateinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *brokerTemplateInstances) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("brokertemplateinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched brokerTemplateInstance.
func (c *brokerTemplateInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BrokerTemplateInstance, err error) {
	result = &v1.BrokerTemplateInstance{}
	err = c.client.Patch(pt).
		Resource("brokertemplateinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
