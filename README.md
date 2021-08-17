# cadence-poc

---
## Architecture
Everything runs via docker-compose, using the template given by Uber.
We expose 3 extra services:
- The [API](./api), a dummy API used as an exemple to trigger workflows
- The [Preparer](./cmd/workers.go), a Cadence Worker which runs a workflow invoked by the API.
  - In theory, this workflow simply checks the data from the API is correct, upload it to vault & set a status in DB for the cronjob
- The [CronJob](./cmd/workers.go), another Cadence Worker which runs a workflow as a cronjob.
  - This workflow picks up the keys prepared by the preparer and create a PR on GitHub to deploy the validators
  - In theory, it is essentially the gitops manager


- The [cmd](./cmd) directory
  - Stores the mains of the 3 services
- The [cadence](./cadence) directory
  - Provides a wrapper around Cadence to easily build clients and workers
  - Exposes our workflows & their activities
    - Each workflow has its subdirectory with its function and activities
- The [api](./api) directory
  - Exposes the dummy API used to call the workflows

---
## How To Run ?

`docker-compose up --build` should be sufficient

---
## TODOs

- Make the API mock up call a workflow
  - [Framework to test workflows]()
    - Seems fairly simple once you get the notions


---
## Links

- [Presentation of Cadence @Uber Open Summit 2018](https://www.youtube.com/watch?v=llmsBGKOuWI)
- [Website](https://cadenceworkflow.io/)
- [Doc](https://cadenceworkflow.io/docs/get-started/)
- [API](https://pkg.go.dev/go.uber.org/cadence)
- [Samples](https://github.com/uber-common/cadence-samples/blob/master/cmd/samples)
  - [Mix API w/ Workflow](https://github.com/uber-common/cadence-samples/tree/master/cmd/samples/pageflow)
- [Building your first Cadence Workflow](https://medium.com/stashaway-engineering/building-your-first-cadence-workflow-e61a0b29785)
- [Testing](https://cadenceworkflow.io/docs/go-client/workflow-testing)

---
## Use Cases (-> do we have a use case for it?)

In summary, Cadence fills a lot of use case and could help us scale up; however, at the moment
it is simply too big / complex to be worth it.

- [Periodic execution](https://cadenceworkflow.io/docs/use-cases/periodic-execution/) -> yes
- [Orchestration](https://cadenceworkflow.io/docs/use-cases/orchestration/) -> yes
- [Polling](https://cadenceworkflow.io/docs/use-cases/polling/) -> yes
- [Event driven application](https://cadenceworkflow.io/docs/use-cases/event-driven/) -> yes (does it via `signals` (not UNIX signals))
- [Storage scan](https://cadenceworkflow.io/docs/use-cases/partitioned-scan/) -> no?
- [Batch job](https://cadenceworkflow.io/docs/use-cases/batch-job/) -> maybe?
- [Infrastructure provisioning](https://cadenceworkflow.io/docs/use-cases/provisioning/) -> no?
- [Deployment](https://cadenceworkflow.io/docs/use-cases/deployment/) -> yes (gitops manager essentially become (a) workflow(s))
- [Operational management](https://cadenceworkflow.io/docs/use-cases/operational-management/) -> no 
- [Interactive application](https://cadenceworkflow.io/docs/use-cases/interactive/) -> no
- [DSL workflows](https://cadenceworkflow.io/docs/use-cases/dsl/) -> no
- [Big data and ML](https://cadenceworkflow.io/docs/use-cases/big-ml/) -> no


---
- [activities](https://cadenceworkflow.io/docs/concepts/activities/#activities): basic unit of work
  - User defined
  - State not recovered in case of failures (shouldn't matter as they should be atomic)
  - Invoked asynchronously through a task list (queue)
- [workflow](https://cadenceworkflow.io/docs/concepts/workflows/#overview): can execute activities
  - User defined
  - Controls activity exec options (timeout, retry policy, ..., see cadence/workflow.ActivityOptions)
  - Can react to events
  - Can return its internal state through queries
  - The state of the workflow code, including local variables and threads it creates, is immune to process and Cadence service failures.
  - It encapsulates state, processing threads, durable timers and event handlers.
  - [How does it work internally ?](https://stackoverflow.com/questions/62904129/what-exactly-is-a-cadence-decision-task/63964726#63964726)
- [worker](https://cadenceworkflow.io/docs/go-client/workers/):  service which hosts a workflow & activities
  - User defined
  - Polls Cadence service for tasks
  - Perform them
  - Report back
---