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

package integrationplatform

import (
	"context"
	"testing"
	"time"

	"github.com/aanogueira/camel-k/pkg/platform"
	"github.com/rs/xid"

	v1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1"
	"github.com/aanogueira/camel-k/pkg/util/log"
	"github.com/aanogueira/camel-k/pkg/util/test"

	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestTimeouts_Default(t *testing.T) {
	ip := v1.IntegrationPlatform{}
	ip.Namespace = "ns"
	ip.Name = xid.New().String()
	ip.Spec.Cluster = v1.IntegrationPlatformClusterOpenShift
	ip.Spec.Profile = v1.TraitProfileOpenShift

	c, err := test.NewFakeClient(&ip)
	assert.Nil(t, err)

	assert.Nil(t, platform.ConfigureDefaults(context.TODO(), c, &ip, false))

	h := NewInitializeAction()
	h.InjectLogger(log.Log)
	h.InjectClient(c)

	answer, err := h.Handle(context.TODO(), &ip)
	assert.Nil(t, err)
	assert.NotNil(t, answer)

	n := answer.Status.Build.GetTimeout().Duration.Seconds() * 0.75
	d := (time.Duration(n) * time.Second).Truncate(time.Second)

	assert.Equal(t, d, answer.Status.Build.Maven.GetTimeout().Duration)
	assert.Equal(t, 5*time.Minute, answer.Status.Build.GetTimeout().Duration)
}

func TestTimeouts_MavenComputedFromBuild(t *testing.T) {
	ip := v1.IntegrationPlatform{}
	ip.Namespace = "ns"
	ip.Name = xid.New().String()
	ip.Spec.Cluster = v1.IntegrationPlatformClusterOpenShift
	ip.Spec.Profile = v1.TraitProfileOpenShift

	timeout, err := time.ParseDuration("1m1ms")
	assert.Nil(t, err)

	ip.Spec.Build.Timeout = &metav1.Duration{
		Duration: timeout,
	}

	c, err := test.NewFakeClient(&ip)
	assert.Nil(t, err)

	assert.Nil(t, platform.ConfigureDefaults(context.TODO(), c, &ip, false))

	h := NewInitializeAction()
	h.InjectLogger(log.Log)
	h.InjectClient(c)

	answer, err := h.Handle(context.TODO(), &ip)
	assert.Nil(t, err)
	assert.NotNil(t, answer)

	n := answer.Status.Build.GetTimeout().Duration.Seconds() * 0.75
	d := (time.Duration(n) * time.Second).Truncate(time.Second)

	assert.Equal(t, d, answer.Status.Build.Maven.GetTimeout().Duration)
	assert.Equal(t, 1*time.Minute, answer.Status.Build.GetTimeout().Duration)
}

func TestTimeouts_Truncated(t *testing.T) {
	ip := v1.IntegrationPlatform{}
	ip.Namespace = "ns"
	ip.Name = xid.New().String()
	ip.Spec.Cluster = v1.IntegrationPlatformClusterOpenShift
	ip.Spec.Profile = v1.TraitProfileOpenShift

	bt, err := time.ParseDuration("5m1ms")
	assert.Nil(t, err)

	ip.Spec.Build.Timeout = &metav1.Duration{
		Duration: bt,
	}

	mt, err := time.ParseDuration("2m1ms")
	assert.Nil(t, err)

	ip.Spec.Build.Maven.Timeout = &metav1.Duration{
		Duration: mt,
	}

	c, err := test.NewFakeClient(&ip)
	assert.Nil(t, err)

	assert.Nil(t, platform.ConfigureDefaults(context.TODO(), c, &ip, false))

	h := NewInitializeAction()
	h.InjectLogger(log.Log)
	h.InjectClient(c)

	answer, err := h.Handle(context.TODO(), &ip)
	assert.Nil(t, err)
	assert.NotNil(t, answer)

	assert.Equal(t, 2*time.Minute, answer.Status.Build.Maven.GetTimeout().Duration)
	assert.Equal(t, 5*time.Minute, answer.Status.Build.GetTimeout().Duration)
}

func TestDefaultMavenSettingsApplied(t *testing.T) {
	ip := v1.IntegrationPlatform{}
	ip.Namespace = "ns"
	ip.Name = "test-platform"
	ip.Spec.Cluster = v1.IntegrationPlatformClusterOpenShift
	ip.Spec.Profile = v1.TraitProfileOpenShift

	c, err := test.NewFakeClient(&ip)
	assert.Nil(t, err)

	assert.Nil(t, platform.ConfigureDefaults(context.TODO(), c, &ip, false))

	h := NewInitializeAction()
	h.InjectLogger(log.Log)
	h.InjectClient(c)

	answer, err := h.Handle(context.TODO(), &ip)
	assert.Nil(t, err)
	assert.NotNil(t, answer)

	assert.NotNil(t, answer.Status.Build.Maven.Settings.ConfigMapKeyRef)
	assert.Nil(t, answer.Spec.Build.Maven.Settings.ConfigMapKeyRef)
	assert.Equal(t, "test-platform-maven-settings", answer.Status.Build.Maven.Settings.ConfigMapKeyRef.Name)
	assert.Equal(t, "settings.xml", answer.Status.Build.Maven.Settings.ConfigMapKeyRef.Key)
}
