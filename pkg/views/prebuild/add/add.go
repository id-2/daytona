// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package add

import (
	"errors"
	"log"

	"github.com/daytonaio/daytona/pkg/views"

	"github.com/charmbracelet/huh"
)

type PrebuildAddView struct {
	ProjectConfigName string
	Branch            string
	CommitInterval    string
	TriggerFiles      string
}

func PrebuildCreationView(prebuildAddView *PrebuildAddView, editing bool) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Commit interval").
				Description("Commit interval").
				Value(&prebuildAddView.CommitInterval).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("commit interval can not be blank")
					}

					return nil
				}),
			huh.NewInput().
				Title("Trigger files").
				Value(&prebuildAddView.TriggerFiles).
				Validate(func(str string) error {
					return nil
				}),
		),
	).WithTheme(views.GetCustomTheme())

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
