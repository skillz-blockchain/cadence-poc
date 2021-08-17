package main

import (
    "cadence-poc/cadence/gitops"
    "cadence-poc/cadence/prepare_gitops"
    "cadence-poc/cadence/workers"
    "flag"
    "go.uber.org/cadence/worker"
    "go.uber.org/zap"
    "os"
)

var (
    domain = os.Getenv("CADENCE_DOMAIN")

    preparerTaskList = os.Getenv("PREPARER_TASK_LIST")
    preparerName = os.Getenv("PREPARER_NAME")

    cronTaskList = os.Getenv("CRON_TASK_LIST")
    cronName = os.Getenv("CRON_NAME")
)

func main() {
    // Declare CLI arguments
    var mode string
    flag.StringVar(&mode, "m", "", "Mode is preparer or gitops.")
    flag.Parse()

    // Create Cadence Worker
    var c workers.CadenceAdapter
    c.Setup()

    workerOptions := worker.Options{
        MetricsScope: c.Scope,
        Logger:       c.Logger,
    }

    var cadenceWorker worker.Worker

    // Register workflow & activities based on the mode we run
    switch mode {
    case "preparer":
        cadenceWorker = worker.New(c.ServiceClient, domain, preparerTaskList, workerOptions)
        cadenceWorker.RegisterWorkflow(prepare_gitops.PrepareGitOps)
        cadenceWorker.RegisterActivity(gitops.ValidateKeyFile)

    case "gitops":
        cadenceWorker = worker.New(c.ServiceClient, domain, cronTaskList, workerOptions)
        cadenceWorker.RegisterWorkflow(gitops.GitOpsCron)
        cadenceWorker.RegisterActivity(gitops.DeployWaitingKeys)
        _ = os.Getenv("CRON_STRING")
        //TODO: use cronStr: see https://github.com/uber-common/cadence-samples/blob/master/cmd/samples/cron/
    }

    // Start the worker
    err := cadenceWorker.Start()
    if err != nil {
        c.Logger.Error("Failed to start workers.", zap.Error(err))
        panic("Failed to start workers")
    }

    // The workers are supposed to be long running process that should not exit.
    select {}
}