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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IntegrationPlatformLister helps list IntegrationPlatforms.
// All objects returned here must be treated as read-only.
type IntegrationPlatformLister interface {
	// List lists all IntegrationPlatforms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.IntegrationPlatform, err error)
	// IntegrationPlatforms returns an object that can list and get IntegrationPlatforms.
	IntegrationPlatforms(namespace string) IntegrationPlatformNamespaceLister
	IntegrationPlatformListerExpansion
}

// integrationPlatformLister implements the IntegrationPlatformLister interface.
type integrationPlatformLister struct {
	indexer cache.Indexer
}

// NewIntegrationPlatformLister returns a new IntegrationPlatformLister.
func NewIntegrationPlatformLister(indexer cache.Indexer) IntegrationPlatformLister {
	return &integrationPlatformLister{indexer: indexer}
}

// List lists all IntegrationPlatforms in the indexer.
func (s *integrationPlatformLister) List(selector labels.Selector) (ret []*v1.IntegrationPlatform, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IntegrationPlatform))
	})
	return ret, err
}

// IntegrationPlatforms returns an object that can list and get IntegrationPlatforms.
func (s *integrationPlatformLister) IntegrationPlatforms(namespace string) IntegrationPlatformNamespaceLister {
	return integrationPlatformNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationPlatformNamespaceLister helps list and get IntegrationPlatforms.
// All objects returned here must be treated as read-only.
type IntegrationPlatformNamespaceLister interface {
	// List lists all IntegrationPlatforms in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.IntegrationPlatform, err error)
	// Get retrieves the IntegrationPlatform from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.IntegrationPlatform, error)
	IntegrationPlatformNamespaceListerExpansion
}

// integrationPlatformNamespaceLister implements the IntegrationPlatformNamespaceLister
// interface.
type integrationPlatformNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationPlatforms in the indexer for a given namespace.
func (s integrationPlatformNamespaceLister) List(selector labels.Selector) (ret []*v1.IntegrationPlatform, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IntegrationPlatform))
	})
	return ret, err
}

// Get retrieves the IntegrationPlatform from the indexer for a given namespace and name.
func (s integrationPlatformNamespaceLister) Get(name string) (*v1.IntegrationPlatform, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("integrationplatform"), name)
	}
	return obj.(*v1.IntegrationPlatform), nil
}
