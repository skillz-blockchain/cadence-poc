package prepare_gitops

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

type KeyFile map[string]interface{}

func ValidateKeyFile(ctx context.Context, fileContent []byte) (KeyFile, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Checking Key File's content...", zap.ByteString("content", fileContent))

	kf := make(KeyFile)
	if err := json.Unmarshal(fileContent, kf); err != nil {
		logger.Error("ValidateKeyFile failed to marshal content.", zap.Error(err))
		return nil, err
	}

	if v, ok := kf["yolo"].(string); !ok || v != "hello_world" {
		err := errors.New("yolo != hello_world")
		logger.Error("ValidateKeyFile bad key file", zap.Error(err))
		return nil, err
	}

	logger.Info("ValidateKeyFile succeed")
	return kf, nil
}

func DeployWaitingKeys(ctx context.Context, kfs []KeyFile) error {
	//TODO: write to some tmp file
	return nil
}

func ReviewWaitingDeployments(ctx context.Context) error {
	//TODO: read some file
	return nil
}

func CheckWaitingForFunds(ctx context.Context) error {
	//TODO: sleep for a while
	return nil
}
