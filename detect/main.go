/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/build-system-buildpack/gradle"
	"github.com/cloudfoundry/build-system-buildpack/maven"
	detectPkg "github.com/cloudfoundry/libcfbuildpack/detect"
)

func main() {
	detect, err := detectPkg.DefaultDetect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize Detect: %s\n", err)
		os.Exit(101)
	}

	if code, err := d(detect); err != nil {
		detect.Logger.Info(err.Error())
		os.Exit(code)
	} else {
		os.Exit(code)
	}
}

func d(detect detectPkg.Detect) (int, error) {
	if gradle.IsGradle(detect.Application) {
		detect.Logger.Debug("Gradle application")
		return detect.Pass(gradle.BuildPlanContribution())
	}

	if maven.IsMaven(detect.Application) {
		detect.Logger.Debug("Maven application")
		return detect.Pass(maven.BuildPlanContribution())
	}

	return detect.Fail(), nil
}
