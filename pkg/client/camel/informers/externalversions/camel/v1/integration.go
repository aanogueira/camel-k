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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	camelv1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1"
	versioned "github.com/aanogueira/camel-k/pkg/client/camel/clientset/versioned"
	internalinterfaces "github.com/aanogueira/camel-k/pkg/client/camel/informers/externalversions/internalinterfaces"
	v1 "github.com/aanogueira/camel-k/pkg/client/camel/listers/camel/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IntegrationInformer provides access to a shared informer and lister for
// Integrations.
type IntegrationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.IntegrationLister
}

type integrationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIntegrationInformer constructs a new informer for Integration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIntegrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIntegrationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIntegrationInformer constructs a new informer for Integration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIntegrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CamelV1().Integrations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CamelV1().Integrations(namespace).Watch(context.TODO(), options)
			},
		},
		&camelv1.Integration{},
		resyncPeriod,
		indexers,
	)
}

func (f *integrationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIntegrationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *integrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&camelv1.Integration{}, f.defaultInformer)
}

func (f *integrationInformer) Lister() v1.IntegrationLister {
	return v1.NewIntegrationLister(f.Informer().GetIndexer())
}
