package controllers

import (
	"cadence-poc/cadence/prepare_gitops"
	"cadence-poc/cadence/workers"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/cadence/client"
	"net/http"
)

func ValidatorsDeploy(ctx *gin.Context) {
	wfc, _ := ctx.MustGet("wfc").(workers.CadenceAdapter)
	wfo := client.StartWorkflowOptions{}

	exec, err := wfc.CadenceClient.StartWorkflow(context.Background(), wfo, prepare_gitops.PrepareGitOps)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	js, err := json.Marshal(exec)
	ctx.JSON(http.StatusOK, js)
}

func ValidatorsUnstake(ctx *gin.Context) {
	ctx.String(http.StatusOK, "not implemented yet")
}
