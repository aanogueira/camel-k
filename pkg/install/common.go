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

package install

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	k8s "k8s.io/client-go/kubernetes"

	ctrl "sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1"
	"github.com/aanogueira/camel-k/pkg/client"
	"github.com/aanogueira/camel-k/pkg/resources"
	"github.com/aanogueira/camel-k/pkg/util/kubernetes"
	"github.com/aanogueira/camel-k/pkg/util/openshift"
)

// ResourceCustomizer can be used to inject code that changes the objects before they are created
type ResourceCustomizer func(object ctrl.Object) ctrl.Object

// IdentityResourceCustomizer is a ResourceCustomizer that does nothing
var IdentityResourceCustomizer = func(object ctrl.Object) ctrl.Object {
	return object
}

// Resources installs named resources from the project resource directory
func Resources(ctx context.Context, c client.Client, namespace string, force bool, customizer ResourceCustomizer, names ...string) error {
	return ResourcesOrCollect(ctx, c, namespace, nil, force, customizer, names...)
}

// ResourcesOrCollect --
func ResourcesOrCollect(ctx context.Context, c client.Client, namespace string, collection *kubernetes.Collection,
	force bool, customizer ResourceCustomizer, names ...string) error {
	for _, name := range names {
		if err := ResourceOrCollect(ctx, c, namespace, collection, force, customizer, name); err != nil {
			return err
		}
	}
	return nil
}

// Resource installs a single named resource from the project resource directory
func Resource(ctx context.Context, c client.Client, namespace string, force bool, customizer ResourceCustomizer, name string) error {
	return ResourceOrCollect(ctx, c, namespace, nil, force, customizer, name)
}

func ResourceOrCollect(ctx context.Context, c client.Client, namespace string, collection *kubernetes.Collection,
	force bool, customizer ResourceCustomizer, name string) error {
	obj, err := kubernetes.LoadResourceFromYaml(c.GetScheme(), resources.ResourceAsString(name))
	if err != nil {
		return err
	}

	return ObjectOrCollect(ctx, c, namespace, collection, force, customizer(obj))
}

func ObjectOrCollect(ctx context.Context, c client.Client, namespace string, collection *kubernetes.Collection, force bool, obj ctrl.Object) error {
	if collection != nil {
		// Adding to the collection before setting the namespace
		collection.Add(obj)
		return nil
	}

	obj.SetNamespace(namespace)

	if obj.GetObjectKind().GroupVersionKind().Kind == "PersistentVolumeClaim" {
		if err := c.Create(ctx, obj); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
	}

	if force {
		if err := kubernetes.ReplaceResource(ctx, c, obj); err != nil {
			return err
		}
		// For some resources, also reset the status
		if obj.GetObjectKind().GroupVersionKind().Kind == v1.IntegrationKitKind ||
			obj.GetObjectKind().GroupVersionKind().Kind == v1.BuildKind ||
			obj.GetObjectKind().GroupVersionKind().Kind == v1.IntegrationPlatformKind {
			if err := c.Status().Update(ctx, obj); err != nil {
				return err
			}
		}
		return nil
	}

	// Just try to create them
	return c.Create(ctx, obj)
}

func isOpenShift(c k8s.Interface, clusterType string) (bool, error) {
	var res bool
	var err error
	if clusterType != "" {
		res = strings.EqualFold(clusterType, string(v1.IntegrationPlatformClusterOpenShift))
	} else {
		res, err = openshift.IsOpenShift(c)
		if err != nil {
			return false, err
		}
	}
	return res, nil
}
