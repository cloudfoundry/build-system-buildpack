/*
 * Copyright 2019-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/cloudfoundry/build-system-cnb/buildsystem"
	"github.com/cloudfoundry/build-system-cnb/cache"
	"github.com/cloudfoundry/build-system-cnb/runner"
	"github.com/cloudfoundry/libcfbuildpack/build"
)

func main() {
	build, err := build.DefaultBuild()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize Build: %s\n", err)
		os.Exit(101)
	}

	if code, err := b(build); err != nil {
		build.Logger.TerminalError(build.Buildpack, err.Error())
		os.Exit(code)
	} else {
		os.Exit(code)
	}
}

func b(build build.Build) (int, error) {
	build.Logger.Title(build.Buildpack)

	if b, ok, err := buildsystem.NewGradleBuildSystem(build); err != nil {
		return build.Failure(102), err
	} else if ok {
		if err = b.Contribute(); err != nil {
			return build.Failure(103), err
		}

		if cache, err := cache.NewGradleCache(build); err != nil {
			return build.Failure(102), err
		} else {
			if err = cache.Contribute(); err != nil {
				return build.Failure(103), err
			}
		}

		if runner, err := runner.NewGradleRunner(build, b); err != nil {
			return build.Failure(102), err
		} else {
			if err = runner.Contribute(); err != nil {
				return build.Failure(103), err
			}
		}
	}

	if buildSystem, ok, err := buildsystem.NewMavenBuildSystem(build); err != nil {
		return build.Failure(102), err
	} else if ok {
		if err = buildSystem.Contribute(); err != nil {
			return build.Failure(103), err
		}

		if cache, err := cache.NewMavenCache(build); err != nil {
			return build.Failure(102), err
		} else {
			if err = cache.Contribute(); err != nil {
				return build.Failure(103), err
			}
		}

		if runner, err := runner.NewMavenRunner(build, buildSystem); err != nil {
			return build.Failure(102), err
		} else {
			if err = runner.Contribute(); err != nil {
				return build.Failure(103), err
			}
		}
	}

	return build.Success()
}
