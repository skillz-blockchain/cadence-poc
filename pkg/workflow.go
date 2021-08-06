package pkg

import (
    "go.uber.org/cadence/workflow"
    "go.uber.org/zap"
    "time"
)

const (
	goodKeyFile = `
{
  'yolo': 'hello_world'
}
`

	badKeyFile = `
{
  'yolo': 'nope',
  '123': '456'
}
`
)

func BasicWorkflow(ctx workflow.Context, value string) error {
    logger := workflow.GetLogger(ctx).With(zap.String("where", "BasicWorkflow"))
    logger.Info("Beginning...", zap.String("value", value))

    ao := workflow.ActivityOptions{
        TaskList:               "sampleTaskList",
        ScheduleToCloseTimeout: time.Second * 60,
        ScheduleToStartTimeout: time.Second * 60,
        StartToCloseTimeout:    time.Second * 60,
        HeartbeatTimeout:       time.Second * 10,
        WaitForCancellation:    false,
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    kf := make(KeyFile)
    if err := workflow.ExecuteActivity(ctx, ValidateKeyFile, value).Get(ctx, kf); err != nil {
        return err
    }

    workflow.ExecuteActivity(ctx, DeployWaitingKeys, kf)
    workflow.ExecuteActivity(ctx, ReviewWaitingDeployments, value)
    workflow.ExecuteActivity(ctx, CheckWaitingForFunds, value)

    return nil
}
