// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/workspace/project/buildconfig"
	"github.com/daytonaio/daytona/pkg/workspace/project/config/prebuild"
)

type ProjectConfig struct {
	Name        string                          `json:"name"`
	Image       string                          `json:"image"`
	User        string                          `json:"user"`
	BuildConfig *buildconfig.ProjectBuildConfig `json:"buildConfig"`
	Repository  *gitprovider.GitRepository      `json:"repository"`
	EnvVars     map[string]string               `json:"envVars"`
	IsDefault   bool                            `json:"default"`
	Prebuilds   []*prebuild.PrebuildConfig      `json:"prebuilds"`
} // @name ProjectConfig

func (pc *ProjectConfig) SetPrebuild(p *prebuild.PrebuildConfig) error {
	newPrebuild := prebuild.PrebuildConfig{
		Id:             p.Id,
		Branch:         p.Branch,
		CommitInterval: p.CommitInterval,
		TriggerFiles:   p.TriggerFiles,
	}

	for _, pb := range pc.Prebuilds {
		if pb.Id == p.Id {
			pb = &newPrebuild
			return nil
		}
	}

	pc.Prebuilds = append(pc.Prebuilds, &newPrebuild)
	return nil
}

func (pc *ProjectConfig) RemovePrebuild(p *prebuild.PrebuildConfig) error {
	for i, pb := range pc.Prebuilds {
		if pb.Id == p.Id {
			pc.Prebuilds = append(pc.Prebuilds[:i], pc.Prebuilds[i+1:]...)
		}
	}
	return nil
}
