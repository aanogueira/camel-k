/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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
	v1 "github.com/aanogueira/camel-k/pkg/client/camel/clientset/versioned/typed/camel/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCamelV1 struct {
	*testing.Fake
}

func (c *FakeCamelV1) Builds(namespace string) v1.BuildInterface {
	return &FakeBuilds{c, namespace}
}

func (c *FakeCamelV1) CamelCatalogs(namespace string) v1.CamelCatalogInterface {
	return &FakeCamelCatalogs{c, namespace}
}

func (c *FakeCamelV1) Integrations(namespace string) v1.IntegrationInterface {
	return &FakeIntegrations{c, namespace}
}

func (c *FakeCamelV1) IntegrationKits(namespace string) v1.IntegrationKitInterface {
	return &FakeIntegrationKits{c, namespace}
}

func (c *FakeCamelV1) IntegrationPlatforms(namespace string) v1.IntegrationPlatformInterface {
	return &FakeIntegrationPlatforms{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCamelV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
