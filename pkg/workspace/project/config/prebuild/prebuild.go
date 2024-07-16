// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// PrebuildConfig holds configuration for the prebuild process
type PrebuildConfig struct {
	Id             string   `json:"id"`
	Branch         string   `json:"branch"`
	CommitInterval *int     `json:"commitInterval"`
	TriggerFiles   []string `json:"triggerFiles"`
} // @name PrebuildConfig

func (p *PrebuildConfig) GenerateId() error {
	triggerFilesJson, err := json.Marshal(p.TriggerFiles)
	if err != nil {
		return err
	}

	data := string(p.Branch) + string(triggerFilesJson)
	hash := sha256.Sum256([]byte(data))
	hashStr := hex.EncodeToString(hash[:])

	p.Id = hashStr

	return nil
}
