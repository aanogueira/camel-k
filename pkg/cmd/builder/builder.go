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

package builder

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"k8s.io/apimachinery/pkg/types"

	controller "sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	v1 "github.com/aanogueira/camel-k/pkg/apis/camel/v1"
	"github.com/aanogueira/camel-k/pkg/builder"
	"github.com/aanogueira/camel-k/pkg/client"
	"github.com/aanogueira/camel-k/pkg/util/cancellable"
	"github.com/aanogueira/camel-k/pkg/util/defaults"
	logger "github.com/aanogueira/camel-k/pkg/util/log"
	"github.com/aanogueira/camel-k/pkg/util/patch"
)

var log = logger.WithName("builder")

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	log.Info(fmt.Sprintf("Camel K Version: %v", defaults.Version))
}

// Run a build resource in the specified namespace
func Run(namespace string, buildName string, taskName string) {
	logf.SetLogger(zap.New(func(o *zap.Options) {
		o.Development = false
	}))

	rand.Seed(time.Now().UTC().UnixNano())
	printVersion()

	c, err := client.NewClient(false)
	exitOnError(err, "")

	ctx := cancellable.NewContext()

	build := &v1.Build{}
	exitOnError(
		c.Get(ctx, types.NamespacedName{Namespace: namespace, Name: buildName}, build), "",
	)

	status := builder.New(c).Build(build).TaskByName(taskName).Do(ctx)
	target := build.DeepCopy()
	target.Status = status
	// Copy the failure field from the build to persist recovery state
	target.Status.Failure = build.Status.Failure
	// Patch the build status with the result
	p, err := patch.PositiveMergePatch(build, target)
	exitOnError(err, "cannot create merge patch")
	if len(p) > 0 {
		exitOnError(
			c.Status().Patch(ctx, target, controller.RawPatch(types.MergePatchType, p)),
			fmt.Sprintf("\n--- patch ---\n%s\n-------------\n", string(p)),
		)
	} else {
		log.Info("Patch not applied (no difference)")
	}
	build.Status = target.Status

	switch build.Status.Phase {
	case v1.BuildPhaseFailed:
		log.Error(nil, build.Status.Error)
		os.Exit(1)
	default:
		os.Exit(0)
	}
}

func exitOnError(err error, msg string) {
	if err != nil {
		log.Error(err, msg)
		os.Exit(1)
	}
}
