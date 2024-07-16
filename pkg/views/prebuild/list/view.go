// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package list

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views"
	views_util "github.com/daytonaio/daytona/pkg/views/util"
	"golang.org/x/term"
)

type RowData struct {
	ProjectConfigName string
	Branch            string
	TriggerFiles      string
	CommitInterval    string
}

func ListPrebuilds(prebuildList []apiclient.PrebuildDTO) {
	re := lipgloss.NewRenderer(os.Stdout)

	headers := []string{"Project Config", "Branch", "Trigger files", "Commit interval"}

	data := [][]string{}

	for _, pc := range prebuildList {
		var rowData *RowData
		var row []string

		rowData = getTableRowData(pc)
		row = getRowFromRowData(*rowData)
		data = append(data, row)
	}

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println(data)
		return
	}

	breakpointWidth := views.GetContainerBreakpointWidth(terminalWidth)

	minWidth := views_util.GetTableMinimumWidth(data)

	if breakpointWidth == 0 || minWidth > breakpointWidth {
		renderUnstyledList(prebuildList)
		return
	}

	t := table.New().
		Headers(headers...).
		Rows(data...).
		BorderStyle(re.NewStyle().Foreground(views.LightGray)).
		BorderRow(false).BorderColumn(false).BorderLeft(false).BorderRight(false).BorderTop(false).BorderBottom(false).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return views.TableHeaderStyle
			}
			return views.BaseCellStyle
		}).Width(breakpointWidth - 2*views.BaseTableStyleHorizontalPadding - 1)

	fmt.Println(views.BaseTableStyle.Render(t.String()))
}

func renderUnstyledList(prebuildList []apiclient.PrebuildDTO) {
	// todo
}

func getRowFromRowData(rowData RowData) []string {
	var isDefault string

	if rowData.CommitInterval == "" {
		isDefault = views.InactiveStyle.Render("/")
	} else {
		isDefault = views.ActiveStyle.Render("Yes")
	}

	row := []string{
		views.NameStyle.Render(rowData.ProjectConfigName),
		views.DefaultRowDataStyle.Render(rowData.Branch),
		views.DefaultRowDataStyle.Render(rowData.TriggerFiles),
		isDefault,
	}

	return row
}

func getTableRowData(prebuildConfig apiclient.PrebuildDTO) *RowData {
	rowData := RowData{"", "", "", ""}

	rowData.ProjectConfigName = *prebuildConfig.ProjectConfigName + views_util.AdditionalPropertyPadding
	rowData.Branch = *prebuildConfig.Branch
	rowData.CommitInterval = strconv.Itoa(int(*prebuildConfig.CommitInterval))

	rowData.TriggerFiles = getTriggerFilesString(prebuildConfig.TriggerFiles)

	return &rowData
}

func getTriggerFilesString(triggerFiles []string) string {
	if len(triggerFiles) == 0 {
		return views.InactiveStyle.Render("None")
	}

	var result string

	for i, triggerFile := range triggerFiles {
		result += triggerFile
		if i != len(triggerFiles)-1 {
			result += ", "
		}
	}

	return result
}
