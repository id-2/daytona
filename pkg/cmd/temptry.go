// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

var tempTryCmd = &cobra.Command{
	Use:   "temptry",
	Short: "temptry",
	Run: func(cmd *cobra.Command, args []string) {
		// server := server.GetInstance(nil)
	},
}
