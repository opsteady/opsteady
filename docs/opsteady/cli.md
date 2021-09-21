# Opsteady CLI

Opsteady CLI is used to do all kind of tasks regarding local work and CI/CD as explained in the [ADR CI/CD](../adr/0012-ci-cd.md) and [CLI](../adr/0016-cli.md).

## How to run the CLI?

To run the CLI locally use the project container or run locally. If running locally make sure you have all the required tools. To run the CLI in the container:

```bash
docker run -it --rm \
  --net host \
  -v $(pwd):/data \
  -v ${HOME}/.cache:/home/platform/.cache \
  -v ${HOME}/.cache/opsteady-go:/home/platform/go \
  dev-management.azurecr.io/cicd:1.0.0 /bin/bash
go run . help
```

The **.cache** and **opsteady-go** folder are needed to save cached data and golang build. This will save a lot of time when restarting the container.

## CLI settings

Every setting for Opsteady CLI can be provided using the command line flags. Some of the settings can also be provided using the **config.yaml** or the **environments variables**.

The **default-config.yaml** is meant as an example, please provide your own **config.yaml** which doesn't have to be committed.

## Example commands

Try the deployment to management bootstrap out without actually deploying it:

```bash
go run . deploy --component management-bootstrap --azure-id management --dry-run
```

Deploy management bootstrap and use the cached credentials:

```bash
go run . deploy --component management-bootstrap --azure-id management --cache
```

Destroy management bootstrap:

```bash
go run . destroy --component management-bootstrap --azure-id management
```
