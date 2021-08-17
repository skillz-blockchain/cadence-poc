package prepare_gitops

import (
	"cadence-poc/cadence/gitops"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"os"
	"time"
)

func PrepareGitOps(ctx workflow.Context, value string) error {
	logger := workflow.GetLogger(ctx).With(zap.String("where", "PrepareGitOps"))
	logger.Info("Beginning...", zap.String("value", value))

	ao := workflow.ActivityOptions{
		TaskList:               os.Getenv("WORKER_TASK_LIST"),
		ScheduleToCloseTimeout: time.Minute,
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	kf := make(gitops.KeyFile)
	if err := workflow.ExecuteActivity(ctx, gitops.ValidateKeyFile, value).Get(ctx, kf); err != nil {
		return err
	}

	workflow.ExecuteActivity(ctx, gitops.DeployWaitingKeys, kf)
	workflow.ExecuteActivity(ctx, gitops.ReviewWaitingDeployments, value)
	workflow.ExecuteActivity(ctx, gitops.CheckWaitingForFunds, value)

	return nil
}
