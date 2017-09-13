package v1

import (
	"github.com/openshift/client-go/template/clientset/scheme"
	v1 "github.com/openshift/origin/pkg/template/apis/template/v1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type TemplateV1Interface interface {
	RESTClient() rest.Interface
	BrokerTemplateInstancesGetter
	TemplatesGetter
	TemplateInstancesGetter
}

// TemplateV1Client is used to interact with features provided by the template.openshift.io group.
type TemplateV1Client struct {
	restClient rest.Interface
}

func (c *TemplateV1Client) BrokerTemplateInstances() BrokerTemplateInstanceInterface {
	return newBrokerTemplateInstances(c)
}

func (c *TemplateV1Client) Templates(namespace string) TemplateResourceInterface {
	return newTemplates(c, namespace)
}

func (c *TemplateV1Client) TemplateInstances(namespace string) TemplateInstanceInterface {
	return newTemplateInstances(c, namespace)
}

// NewForConfig creates a new TemplateV1Client for the given config.
func NewForConfig(c *rest.Config) (*TemplateV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TemplateV1Client{client}, nil
}

// NewForConfigOrDie creates a new TemplateV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TemplateV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TemplateV1Client for the given RESTClient.
func New(c rest.Interface) *TemplateV1Client {
	return &TemplateV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TemplateV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
