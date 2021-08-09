# cadence-poc

---
## Architecture

Currently this repo exposes 2 main components:
- [api](./api) is a mock up of the API proposed [here](https://www.notion.so/skillzblockchain/PRD-Ethereum-2-0-Validator-node-API-2c1023a26dcb4bf99e927c24596e2be6#a9c95e7c71f149c78065ed6d92b45c23).
  - It uses `cadence-web` for the server/routing and allows us to trigger workflows provided by workers.
  - At the moment, the work's been mostly around the architecture and properly understanding / adapting cadence's sample: [pageflow](https://github.com/uber-common/cadence-samples/tree/master/cmd/samples/pageflow)
    - It's very much WIP and requires some more work to function
- [workers](./workers) are mock ups of the workflow which would be triggered by our API.
  - [Activities](./workers/activities) are steps of a workflow. They should be as atomical as possible.
  - [Workflows](./workers/workflows) put activities together and handle errors which could arise. They're meant to be called by the API
  - At the moment, the work's been mostly the discovery of the components of cadence and how they work together.
  
---
## Links

- [Website](https://cadenceworkflow.io/)
- [Doc](https://cadenceworkflow.io/docs/get-started/)
- [API](https://pkg.go.dev/go.uber.org/cadence)
- [Samples](https://github.com/uber-common/cadence-samples/blob/master/cmd/samples)
  - [Mix API w/ Workflow](https://github.com/uber-common/cadence-samples/tree/master/cmd/samples/pageflow)
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