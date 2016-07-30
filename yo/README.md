# generic API

## Setup

The project expects to be in a valid `GOPATH`. Namely, it needs to be in a directory workspace which conforms to [this](https://golang.org/doc/code.html#Organization).

Most importantly, the repo should be cloned to a directory with the name `$GOPATH/src/ultimatesoftware.com/notifications`.

## Dependencies

Dependencies are managed using the [godep](https://github.com/tools/godep) tool. Thus, if you need to add a dependency, make sure you use the `godep` tool to do that. Then the dependency will be stored in the `Godeps/Godeps.json` file. Note that when you add a dependency, **do not** use the `-r` option, which will rewrite the import path.

## Build

A Rakefile contains all the tasks to build the project. Running `rake` will install all dependencies and build the binary.

## Docker

If you have [Docker](https://www.docker.com/) and [Docker Compose](https://www.docker.com/docker-compose) installed, you can quickly run all of the individual components of the Notification Service on your local machine.

All of the Dockerfiles are in the `docker/` directory.

There are a set of `rake` tasks to help with using Docker.

    $ rake -T docker
    ...

To build all of the images and start containers from them, execute `rake docker`.

Use `rake docker:build` build build the Docker images from scratch.

After you've built the images, you can start all of the services with `rake docker:start`. The `start` task also takes an optional parameter to detach and start the services in the background. You can invoke it with `rake docker:start[true]`.

Stop, kill, and remove containers with the `stop, kill, and rm` tasks respectively.

When you use `docker:start`, Docker Compose will start up 1 container:

  * The HTTP API. You can access the service on port `8000`.

All of the ports mentioned above will be on your Docker Machine's IP address. So if you're using a Mac or Windows, this **will not** be localhost. Use `docker-machine` to find out the IP address of your Docker virtual machine.

    $ docker-machine ip default
    192.168.99.100

Except for the `rake docker:build` task, all of the `docker:*` tasks are simply thin wrappers around the `docker-compose` command. If you want to access any of the more advanced features of `docker-compose`, such as looking at logs or starting/stopping individual containers, you'll want to call this command directly. Look at the tasks in `docker/docker.rake` to see how these tasks use the `docker-compose` command. Executing `docker-compose` with no arguments will give a useful help message.

    $ docker-compose
    Define and run multi-container applications with Docker.
    ...
    version            Show the Docker-Compose version information

If you execute `docker-compose` on it's own, make sure you are in the root directory of the project repo. Docker compose uses the `docker-compose.yml` stored in the directory, and will throw errors if it can't find it.

## Deployments

The tasks `create`, `list`, `delete` are available. They interact with SPADE to create new deployments, list existing deployments and delete them respectively.

There also support for **self reliant** deployments. Those are deployments that can be triggered locally without having the artifacts published to artifactory. The task to trigger a self reliant deployment is `create_self[name, port]` and it will:
  1. Build and Package the artifacts
  2. Copy the artifacts to the directory `app_root/ci/artifacts_host/generic/0.0.1` and host them using Python SimpleHTTPServer. (The port is specified as a parameter to `create_self`).
  3. Modify `config` to use the hosted artifacts instead of artifactory.
  4. Create a new deployment.
