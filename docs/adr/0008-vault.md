# 8. Vault

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27

## Context

We need to store and use secrets throughout the whole CI/CD process for the creation and maintenance of environments. Besides secretes, we need to store different properties and configurations for rolled-out environments.

## Decision

One of the options was to use a cloud-native solution but as we are using multiple clouds this would mean extra complexity and differences. Therefore we have chosen to use HashiCorp Vault as it is a very versatile solution, it is built with security in mind and has APIs that can be used to automate the process.
We will be storing the secrets and our configuration inside Vault, this way we have one uniform way for both.

Besides that, we will also use Vault for getting access to all the environments so that we have one entry point for getting credentials for the environments we created.

## Consequences

It gives us a uniform way to handle the secrets, configuration, and access to environments but it does lock us in on Vault and maintenance of it.
