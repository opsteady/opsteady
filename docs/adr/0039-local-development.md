# 39. Local Development

Date: 2022-02-07

## Status

Status: Accepted on 2022-02-07

Build on [0005-multi-cloud.md](0005-multi-cloud.md) on 2022-02-07

## Context

We need a way to develop the Opsteady platform without incurring a lot of (cloud) costs or waiting on long provisioning times. Cloud environments are the main target for production workloads but they are slow(er), costly and not strictly necessary to develop the Opsteady capabilities and portal.

## Decision

We will use [k3d](https://k3d.io) to provision a Kubernetes cluster for local development. These clusters are based on k3s and run in containers on the local development workstation. The Opsteady deployment workflow will remain largely the same, except that we now target 'local' instead of the cloud environments. We cannot avoid the cloud infrastructure entirely because we need DNS and single sign-on capabilities. However, these pieces of infrastructure are quick to provision and very cheap to run.

We use Azure AD and a single Azure subscription to setup IAM and cloud infrastructure that is needed for local development purposes. Setting up the infrastructure is done in the normal way with our deployment tool.

## Consequences

Local development clusters will not be a full representation of normal cloud-based clusters. In particular, we cannot use any functionality that is directly integrated into the cloud environment, such as pod identity, load balancing, storage, etc. For some of these items there is a local alternative, to mimic the functionality, but for others we will simply ignore them during development. This means that we need regular cloud deployments to verify that the entire setup, with all features enabled, still functions as expected. This is an acceptable tradeoff because it gives us a big amount of flexibility and speed while doing development work in the 'upper' layers of the platform (capabilities/portal).
