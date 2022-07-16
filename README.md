# GoSanta

Golang service to assign awards to users.

This is test project to explore DDD & Hexagonal design patterns in golang. Inspired by: [goddd](https://github.com/marcusolsson/goddd)

## Prerequisites

- golang 1.17
- docker
- make

## Init

- create env file from example env
  - `cp example.env .env`
  - fill out missing env vars
- export env, e.g.:
  - `export $(cat .env | xargs)` or
  - `set -o allexport source .env set +o allexport`
- install tools and dependencies
  - `make install-tools`
- build and run containers
  - `make run-docker-composed`
- create db schema
  - `make init-db`

**Not End-to-End tested yet**

The service requires and external queue to connect to (i.e. AWS SQS) so it does
not make sense to run it locally. I need to provide a mock implementation of
ports.EventQueue for that.

## Running Tests

- db connection with docker needed before running tests
  - `make`
- for coverage report
  - `make test-cov`

## Domain

To test the human component of a company's IT security, a phishing test
provider can send test phishing emails to users. A `User` can interact with an email by:

1. Ignoring it
2. Clicking on the phishing link
3. Opening the email and doing nothing else
4. Or "reporting" it.

To improve interaction numbers (i.e. everything but ignoring), a gamification
component is added to the educational phishing, which is implemented by this application.

## Architecture

### Hexagonal

Or `Ports & Adapters` is a design pattern in which application code is split
into domain logic and infrastructure implementations. A domain, here `awards`,
consists of `PhishingEvent`s being processed to assing `PhishingAward`s to
`User`s belonging to `Company`s. The orchestration of that domain is done
through services. A service, e.g. `awarding`, receives a phishing event as input,
applies the domain logic to it and creates (writes to a repository) the
appropriate `PhishingAward`. Such a domain service can be interacted with
through different `ports`, e.g. an HTTP interface or a cron job. The main point
is, that the domain and domain services are not coupled to infrastructure
implementing the domain's functionality, e.g. the `AwardService` depends on an
`ports.AwardRepository` interface which is implemented by
`postgres.AwardRepository`. Additionally, the *outside* ports or *driving*
ports, e.g. the (HTTP) `server` are also decoupled from the actual service implementation.

### DDD

I am not very knowledgeble in DDD and thus I do not know all ideas it encompasses. The
main idea used here from DDD is the naming, which is explained further
[here](#package-and-folder-naming). Services of the application should explain
what the application does, and domain types should mirror business objects.

## TODO

1. User sync.
   Ideally, user data would be shared by a user service over events. E.g.
   UserCreated, UserUpdated, etc. Those events could be put into the same queue
   as the phishing events.
2. Swagger/OpenAPI doc gen.
   Should be put into the pkg directory.
3. Logging and metrics.
4. Auth?
5. CI/CD?

## Considerations

### Ports package

A crucial part of hexagonal design patterns are the ports. Ports can be driven
or drive the application. E.g. an HTTP interface can drive the application
trough some application service port, e.g. here the `AwardReadingService`.

In which package should the interface be defined? It is considered an
anti-pattern, to put interfaces in the same package as an implementation of
that interface, as the idea of an interface is to decouple a dependency from a
specific implementation.

On the other hand, putting that interface on, e.g. the `server` package, would
require to duplicate the interface definition in the case of another package
requiring the same interface, e.g. an RPC server.

Having a separate package for *all* (internal) ports, eliminates that question
with, for now, no obvious downsides.

### Multiple commands

Deprecated. Have a look [here](docs/adrs.md).

### Package and folder naming

The root of the repo should follow the [standard golang project layout](https://github.com/golang-standards/project-layout).

As we want to have a screaming architecture, one way to name folders could be
having the application services named as verbs and architecture (port
implementations) as the used technology. E.g.:

1. The application creates *awards* for correct interactions of users with test phishing mails -> awarding (service)
2. The application provides read views on earned awards and ranks users of a company by a score -> ranking (service)
3. The application persists phishing and user events coming from some source -> eventlogging (not a great name imho)

Packages that implement ports, e.g. all the repository ports are implemented
with postgres, are named after the technology used. In the case of the postgres
package, we could also call it `persistence` or `rdbs` or so. But naming it
postgres signals to a new *developer* looking at the repo, that this application
deals (among other things) with data persisting, specifically through a RDBS
and with a specific RDBS, named postgresql. This leaves also room for other
packages that could also deal with data persistence but not with a RDBS but a
NoSQL db like mongodb. Or with other RDBS like MySQL. Also, this makes the
folder structure flatter, which is more readable imho.

With such a naming convention, folders of application services can be
differentiated from infrastructure just by a glance on the folder structure.

### Testing

Ideally, all functionality is covered with unit tests. This is done more or
less, as the ports are naturally placed interfaces that can be easily mocked.
However, a very strict observer would not classify the service tests as unit
tests, as they also test (without mocking) the functionality of the domain
package `awards`. But I think we can live with that.

The database functionality is only covered in integration tests. I see no point
in creating unit tests and mocking all the behaviour the `bun` package
provides. As the `postgres` package only uses the functionality of `bun` and
casting the DB types to domain types, there is not much to test if we would
mock `bun`. Also, spinning up a postgresql db in a container is very easy.
`bun` also provides functionality to seed such a test db. All that makes
writing integration tests more feasible, fun and useful.

Additionally, we could mark such integration tests so that unit tests can be run
separate from the integration tests.

