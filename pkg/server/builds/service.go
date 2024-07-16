// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package builds

import (
	"github.com/daytonaio/daytona/pkg/build"
)

type IBuildService interface {
	Delete(hash string) error
	Find(hash string) (*build.Build, error)
	List(filter *build.BuildFilter) ([]*build.Build, error)
	Save(*build.Build) error
}

type BuildServiceConfig struct {
	BuildStore build.Store
}

type BuildService struct {
	buildStore build.Store
}

func NewBuildService(config BuildServiceConfig) IBuildService {
	return &BuildService{
		buildStore: config.BuildStore,
	}
}

func (s *BuildService) List(filter *build.BuildFilter) ([]*build.Build, error) {
	return s.buildStore.List(filter)
}

func (s *BuildService) Find(hash string) (*build.Build, error) {
	return s.buildStore.Find(hash)
}

func (s *BuildService) Save(b *build.Build) error {
	return s.buildStore.Save(b)
}

func (s *BuildService) Delete(hash string) error {
	return s.buildStore.Delete(hash)
}
