/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/aiter/k8s-controller-custom-resource/pkg/apis/samplecrd/v1"
	"github.com/aiter/k8s-controller-custom-resource/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SamplecrdV1Interface interface {
	RESTClient() rest.Interface
	NetworksGetter
}

// SamplecrdV1Client is used to interact with features provided by the samplecrd.k8s.io group.
type SamplecrdV1Client struct {
	restClient rest.Interface
}

func (c *SamplecrdV1Client) Networks(namespace string) NetworkInterface {
	return newNetworks(c, namespace)
}

// NewForConfig creates a new SamplecrdV1Client for the given config.
func NewForConfig(c *rest.Config) (*SamplecrdV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SamplecrdV1Client{client}, nil
}

// NewForConfigOrDie creates a new SamplecrdV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SamplecrdV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SamplecrdV1Client for the given RESTClient.
func New(c rest.Interface) *SamplecrdV1Client {
	return &SamplecrdV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SamplecrdV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
