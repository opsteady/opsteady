# 17. Component configuration

Date: 2021-09-26

## Status

Status: Accepted on 2021-09-26
Builds on [0016-cli.md](0016-cli.md) on 2021-09-26
Builds on [0008-vault.md](0008-vault.md) on 2021-09-26
Foundation for [0018-vault-setup.md](0018-vault-setup.md) on 2021-10-06

## Context

Deciding how to configure components when run using the CLI.

## Decision

A component needs multiple sets of configurations which can be normal settings, other components information, and potentially secrets. A component also might need a platform specific configuration because platform A has different needs than platform B. To make sure we don't mix different ways of handling all of these configurations and secrets, we are going to use Vault for all of them. For now, it means that we treat the entire config as a secret.

It is important to know that Vault uses paths like `/config/abcd` where `abcd` can contain one or more key values. However, we will store multiple key values inside one path in most cases!

Every component might need its own default configuration that is preset by a person as a default. We will store every component in its path and every path holds all the default values for that component:

```bash
/config/component/componentA-default
/config/component/componentB-default
/config/component/componentC-default
```

As mentioned before we also need platform specific values for a component, which will override the default values. We could store them as follows:

```bash
/config/platform/env111/compA
/config/platform/env111/compB
/config/platform/env111/compC
/config/platform/env555/compA
/config/platform/env555/compB
/config/platform/env555/compC
```

The difference between the component and platform values is that component values should not be changing across the platform while the platform values do change. It doesn't mean every component has to have these values, it is possible to keep all the values in the platform. Therefore it can happen that no values are eventually in the component path.

This works fine when you are always using the "master" version but we also want to know which values have been used in which deployment. Therefore, whenever we do a deployment of the same version, the same values are used while we continue editing the values for new releases. Configuration and secrets need to be immutable. To achieve this, we will prefix all the config with a main version `v0` which will be used to add and adjust values. We then deploy a new release which copies all the Vault paths to the new version and threat that path as read-only. This will look like:

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

Besides the manual configuration, some components or "technologies" inside the same component require information from other components to run. For example a DNS zone or a subnet ID, etc. This data could be fetched with data sources in Terraform but as not everything is done in Terraform, we would need other means to pass this information around when needed. We have decided to push this data to Vault from the component. Vault doesn't have a way to merge credentials, therefore we have decided that every "technology" in a component, like Terraform or Helm, will output its data to Vault by itself. This means that we need to suffix the paths in Vault when they are created. As we only foresee this now in Terraform, here is an example of this:

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

The 'componentName-tf' is an optional path and will only be created if we output something from Terraform to Vault. This will only happen if other things depend on the data from that component. In the future we might also get 'componentName-helm' or 'componentName-kube' if needed.

When a component needs values from another component, the values from that component will be fetched as well. To make sure key names are not the same in two different components we will always prefix the key names with the name of the component inside Vault. This makes the key names unique and traceable from which component the value is coming, when used inside Terraform or Helm. Not all technologies accept dashes, therefore all names will only use underscores in key names. An example is "management_bootstrap_terraform_state_account_name" instead of just "terraform_state_account_name".

Because we need to fetch multiple paths at the same time, we will fetch them in parallel from the CLI.

If you provide the information to the component manually, you can still use kubectl/helm/etc... to run it without the CLI.

## Consequences

We need to comply to a way of working which might not always fit. A path under, for example "/config/v1.5.1/platform/env111/" should stay immutable but because we are also pushing data from a component to that path, this will not be true for paths ending with "-tf".
