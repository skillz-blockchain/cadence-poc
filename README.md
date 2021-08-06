# cadence-poc

---
## Links

- [Website](https://cadenceworkflow.io/)
- [Doc](https://cadenceworkflow.io/docs/get-started/)
- [API](https://pkg.go.dev/go.uber.org/cadence)
- [Samples](https://github.com/uber-common/cadence-samples/blob/master/cmd/samples)

---
## Use Case (-> do we have a use case for it?)

- [Periodic execution](https://cadenceworkflow.io/docs/use-cases/periodic-execution/) -> yes
- [Orchestration](https://cadenceworkflow.io/docs/use-cases/orchestration/) -> yes
- [Polling](https://cadenceworkflow.io/docs/use-cases/polling/) -> yes
- [Event driven application](https://cadenceworkflow.io/docs/use-cases/event-driven/) -> yes (does it via `signals` (not UNIX signals))
- [Storage scan](https://cadenceworkflow.io/docs/use-cases/partitioned-scan/) -> no?
- [Batch job](https://cadenceworkflow.io/docs/use-cases/batch-job/) -> maybe?
- [Infrastructure provisioning](https://cadenceworkflow.io/docs/use-cases/provisioning/) -> no?
- [Deployment](https://cadenceworkflow.io/docs/use-cases/deployment/) -> no (gitops manager)
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