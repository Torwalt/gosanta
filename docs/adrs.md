# Architecture decision records

## Move from multiple commands to one

When starting out the application, I was not sure where to put the commands to
start the HTTP server and the background jobs. Having multiple folders in the
cmd directory seemed like a good idea: Having multiple commands gives us
atleast the advantage, to deploy binaries for each of the "tasks" of this
application separately, and, potentially as microservices. E.g., each binary
could run as an AWS lambda function (some adjustments would need to be done
here first).

But I realized there is more downsides than upsides, specifically as individual
processes will consume more resources than one process with multiple goroutines
running. Also, if there will be a need to have different parts of the program
run separately somewhere then you can always split this thing up again.

As an alternative, the service has now only one command, which starts the HTTP
server and other needed components in goroutines. This simplifies building and
deploying the program as only one binary is generated.

To keep the main package dumb and the service packages simple, an additional
package `eventbroker` is created. The `eventbroker.AwardNotifier` handles the
orchestration and dataflow from all relevant components for the awarding.

