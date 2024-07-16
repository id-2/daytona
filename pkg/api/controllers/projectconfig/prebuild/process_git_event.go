// Copyright 2024 Daytona Platforms Inctx.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/gitprovider"
	_ "github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/workspace/project"
	"github.com/daytonaio/daytona/pkg/workspace/project/buildconfig"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/gin-gonic/gin"
)

// ProcessGitEvent 			godoc
//
//	@Tags			prebuild
//	@Summary		ProcessGitEvent
//	@Description	ProcessGitEvent
//	@Param			workspace	body	interface{}	true	"Webhook event"
//	@Success		200
//	@Router			/project-config/prebuild/process-git-event [post]
//
//	@id				ProcessGitEvent
func ProcessGitEvent(ctx *gin.Context) {
	server := server.GetInstance(nil)

	gitProvider, err := server.GitProviderService.GetGitProviderForHttpRequest(ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get git provider for request: %s", err.Error()))
		return
	}

	var payload map[string]interface{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	obj, err := gitProvider.ParseWebhookEvent(payload)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to process webhook event: %s", err.Error()))
		return
	}

	p := &project.Project{
		ProjectConfig: config.ProjectConfig{
			Repository: &gitprovider.GitRepository{
				Url: obj.Url,
			},
		},
	}

	p.Repository = &gitprovider.GitRepository{
		Url:    obj.Url,
		Branch: &obj.Branch,
	}

	// Autodetect
	p.BuildConfig = &buildconfig.ProjectBuildConfig{}

	// gc, _ := server.GitProviderService.GetConfigForUrl(project.Repository.Url)

	// err = server.WorkspaceService.PrebuildProject(project, gc)
	// if err != nil {
	// 	ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create workspace: %s", err.Error()))
	// 	return
	// }

	ctx.Status(200)
}
