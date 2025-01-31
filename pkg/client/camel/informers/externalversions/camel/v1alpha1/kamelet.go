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

package v1alpha1

import (
	"context"
	time "time"

	camelv1alpha1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1alpha1"
	versioned "github.com/aanogueira/camel-k/pkg/client/camel/clientset/versioned"
	internalinterfaces "github.com/aanogueira/camel-k/pkg/client/camel/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/aanogueira/camel-k/pkg/client/camel/listers/camel/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KameletInformer provides access to a shared informer and lister for
// Kamelets.
type KameletInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KameletLister
}

type kameletInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKameletInformer constructs a new informer for Kamelet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKameletInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKameletInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKameletInformer constructs a new informer for Kamelet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKameletInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CamelV1alpha1().Kamelets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CamelV1alpha1().Kamelets(namespace).Watch(context.TODO(), options)
			},
		},
		&camelv1alpha1.Kamelet{},
		resyncPeriod,
		indexers,
	)
}

func (f *kameletInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKameletInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kameletInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&camelv1alpha1.Kamelet{}, f.defaultInformer)
}

func (f *kameletInformer) Lister() v1alpha1.KameletLister {
	return v1alpha1.NewKameletLister(f.Informer().GetIndexer())
}
