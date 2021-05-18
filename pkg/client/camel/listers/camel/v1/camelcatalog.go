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

// CamelCatalogLister helps list CamelCatalogs.
// All objects returned here must be treated as read-only.
type CamelCatalogLister interface {
	// List lists all CamelCatalogs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CamelCatalog, err error)
	// CamelCatalogs returns an object that can list and get CamelCatalogs.
	CamelCatalogs(namespace string) CamelCatalogNamespaceLister
	CamelCatalogListerExpansion
}

// camelCatalogLister implements the CamelCatalogLister interface.
type camelCatalogLister struct {
	indexer cache.Indexer
}

// NewCamelCatalogLister returns a new CamelCatalogLister.
func NewCamelCatalogLister(indexer cache.Indexer) CamelCatalogLister {
	return &camelCatalogLister{indexer: indexer}
}

// List lists all CamelCatalogs in the indexer.
func (s *camelCatalogLister) List(selector labels.Selector) (ret []*v1.CamelCatalog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CamelCatalog))
	})
	return ret, err
}

// CamelCatalogs returns an object that can list and get CamelCatalogs.
func (s *camelCatalogLister) CamelCatalogs(namespace string) CamelCatalogNamespaceLister {
	return camelCatalogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CamelCatalogNamespaceLister helps list and get CamelCatalogs.
// All objects returned here must be treated as read-only.
type CamelCatalogNamespaceLister interface {
	// List lists all CamelCatalogs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CamelCatalog, err error)
	// Get retrieves the CamelCatalog from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CamelCatalog, error)
	CamelCatalogNamespaceListerExpansion
}

// camelCatalogNamespaceLister implements the CamelCatalogNamespaceLister
// interface.
type camelCatalogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CamelCatalogs in the indexer for a given namespace.
func (s camelCatalogNamespaceLister) List(selector labels.Selector) (ret []*v1.CamelCatalog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CamelCatalog))
	})
	return ret, err
}

// Get retrieves the CamelCatalog from the indexer for a given namespace and name.
func (s camelCatalogNamespaceLister) Get(name string) (*v1.CamelCatalog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("camelcatalog"), name)
	}
	return obj.(*v1.CamelCatalog), nil
}
