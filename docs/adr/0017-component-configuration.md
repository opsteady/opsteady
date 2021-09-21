# 17. Component configuration

Date: 2021-09-26

## Status

Status: Accepted on 2021-09-26
Builds on [0016-cli.md](0016-cli.md) on 2021-09-26
Builds on [0008-vault.md](0008-vault.md) on 2021-09-26

## Context

Deciding how to configure components when run using the CLI.

## Decision

A component needs multiple sets of configurations which can be normal settings, other components information, and potentially secrets. A component also might need a platform specifics configuration because platform A has different needs than platform B. To make sure we don't mix different ways of handling all of these configurations and secrets we are going to use Vault for all of them. For now, it means we treat the config as it is a secret.

It is important to know that Vault uses paths like `/config/abcd` where `abcd` can contain one or more key values, in our case, we will in most cases store multiple key values inside one path!

Every component needs its own manual configuration that is preset by a person as a default. We will store every component in its path and every path holds all the default values for that component:

```bash
/config/component/componentA-manual
/config/component/componentB-manual
/config/component/componentC-manual
```

As mentioned before we also need platform specific values for a component which will override the default values. We could store them as following:

```bash
/config/platform/env111/compA
/config/platform/env111/compB
/config/platform/env111/compC
/config/platform/env555/compA
/config/platform/env555/compB
/config/platform/env555/compC
```

This works fine when you are always using the "master" version but we also want to know which values have been used in which deployment. Therefore whenever we do a deployment of the same version the same values are used while we continue editing the values for new releases. Configuration and secrets used need to be immutable. To achieve that we will prefix all the config with a main version `v0` which will be used to add and adjust values and then when we deploy a new release we will copy all the Vault paths to the new version and threat that path as read-only. This will look like this:

```bash
/config/v0/component/compA
/config/v0/component/compB
/config/v0/platform/env111/compA
/config/v0/platform/env111/compB

/config/v1.0.2/component/compA
/config/v1.0.2/component/compB
/config/v1.0.2/platform/env111/compA
/config/v1.0.2/platform/env111/compB

/config/v1.5.1/component/compA
/config/v1.5.1/component/compB
/config/v1.5.1/platform/env111/compA
/config/v1.5.1/platform/env111/compB
```

Besides the manual configuration, some components or "technologies" inside the same component require information from other components to run. For example a DNS zone or a subnet ID, etc. This data could be fetched with data sources in Terraform but as not everything is done in Terraform we would need other means to pass this information around when needed. We have decided to push this data to Vault from the component. Vault doesn't have a way to merge credentials therefore we have decided that every "technology" in a component, like Terraform or Helm, will output its data to Vault by itself which means we need to suffix the paths in Vault when they are created. As we only foresee this now in Terraform, here is only an example for it:

```bash
/config/v1.0.2/component/compA
/config/v1.0.2/component/compB
/config/v1.0.2/platform/env111/compA
/config/v1.0.2/platform/env111/compA-tf
/config/v1.0.2/platform/env111/compB
/config/v1.0.2/platform/env111/compB-tf

/config/v1.5.1/component/compA
/config/v1.5.1/component/compB
/config/v1.5.1/platform/env111/compA
/config/v1.5.1/platform/env111/compA-tf
/config/v1.5.1/platform/env111/compB
/config/v1.5.1/platform/env111/compB-tf
```

The 'componentName-tf' is an optional path and will only be created if we output something from Terraform to Vault which will only happen if other somethings depend on the data from that component. In the future we might also get 'componentName-helm' or 'componentName-kube' if needed.

When a component needs values from another component the values from that component will be fetched as well. To make sure key names are not the same in two different components we will always prefix the key names with the name of the component inside Vault. This makes the key names unique and recognizable from which component the value is coming from when used inside Terraform or Helm. Not all technologies like dashes therefore all names will only use underscores in key names. An example is "management_bootstrap_terraform_state_account_name" instead of just "terraform_state_account_name".

Because we need to fetch multiple paths at the same time we will fetch them parallel from the CLI.

If you provide the information to the component manually you can still use kubectl/helm/etc... to run it without the CLI.

## Consequences

We need to follow a way of working to comply which might not always fit. A path under, for example "/config/v1.5.1/platform/env111/" should stay immutable but because we are also pushing data from a component to that path this will not be true for paths ending with "-tf".
