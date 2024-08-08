// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package config

import "errors"

type Filter struct {
	Name    *string
	Url     *string
	Default *bool
}

type PrebuildFilter struct {
	ProjectConfigName *string
}

type Store interface {
	List(filter *Filter) ([]*ProjectConfig, error)
	Find(filter *Filter) (*ProjectConfig, error)
	Save(projectConfig *ProjectConfig) error
	Delete(projectConfig *ProjectConfig) error
}

var (
	ErrProjectConfigNotFound = errors.New("project config not found")
	ErrPrebuildNotFound      = errors.New("prebuild not found")
)

func IsProjectConfigNotFound(err error) bool {
	return err.Error() == ErrProjectConfigNotFound.Error()
}

func IsPrebuildNotFound(err error) bool {
	return err.Error() == ErrPrebuildNotFound.Error()
}
