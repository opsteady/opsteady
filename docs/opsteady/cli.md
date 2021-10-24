# Opsteady CLI

Opsteady CLI is used to do all kind of tasks regarding local work and CI/CD as explained in the [ADR CI/CD](../adr/0012-ci-cd.md) and [CLI](../adr/0016-cli.md).

## How to run the CLI?

To run the CLI locally use the project container or run locally. If running locally make sure you have all the required tools. To run the CLI in the container:

```bash
docker run -it --rm \
  -p 8250:8250 \
  -v $(pwd):/data \
  -v ${HOME}/.cache:/home/platform/.cache \
  -v ${HOME}/.cache/opsteady-go:/home/platform/go \
  -v /var/run/docker.sock:/var/run/docker.sock \
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

## Folder structure

To simplify and standardize the code and the repository there are some guidelines for folders used in the CLI. Every component (management-bootstrap, foundation, DNS, etc..) has a **cicd** folder where the golang code is created. This code is used to initialize settings like Helm charts, adjust the default execution order or create a completely new flow for deploy, build, etc... if necessary.

This is a full example of a root folder of a component that we recommend to be used:

```bash
aws # The files can be either in the root, aws or azure folder if they differ in config
aws/cicd # component golang code used for executing the component
aws/crd # Kubernetes CRD yaml files that need to be created before other steps
aws/helm/nginx/values.yaml # helm folder can have multiple charts with a values.yaml file
aws/helm/dns/values.yaml # helm folder can have multiple charts with a values.yaml file
aws/terraform # Terraform code to create resources
aws/kubernetes # Kubernetes yaml files to be applied to Kubernetes
aws/code # golang/javascript/java code with a program where the code can be any name
aws/code/docker # contains the Dockerfile for creating a container image
aws/code/chart # contains the chart for this code
```

You can always override this in a component golang code but it is not recommended!

## Default environment variables available

The following environment variables are always available in every step:

- vault_address
- vault_token
- platform_version
- platform_environment_name
- platform_component_name
