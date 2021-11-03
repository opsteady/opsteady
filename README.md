# ![logo](docs/images/logo.png)

Opsteady is an opioniated reference implementation, based on years of platform building experience, for a container-based platform that provides an exceptional developer experience. Next to developers, it also caters to the needs of other organisational roles that are closely involved in the software development lifecycle like security officers, testers, site reliability engineers, etc.

## Background

Building a container-based platform requires a lot of knowledge in areas such as infrastructure, Cloud, Kubernetes, operations, security, tooling, programming, software delivery processes, etc. Operating such a platform, with the necessary stability and performance guarantees, requires ongoing lifecycle management which is both time-consuming and complex. We believe that organizations should not build their own (non-differentiating) platform but instead leverage a managed solution that fit their needs.

Opsteady is designed to deliver and maintain hundreds of platforms with a relatively small team of people. We suggest that you take inspiration from Opsteady to learn what is needed to deliver a full-blown platform in a controlled and secure way. It is possible to run a production setup with Opsteady directly, either by copying or forking the code, but it requires a full understanding of everything related to the used technologies, architecture, and processes.

## Goals

These are the goals of Opsteady:

- Simplify the creation and maintenance container-based platforms
- Automate lifecycle management to lower the maintenance burden
- Staying in control by implementing a solid operations process
- Help and guide developers on their path to production
- Empower developers to own the full lifecycle of their applications
- Embed SRE and DevOps best practices in the software delivery lifecycle
- Cater to the needs of all organisational roles that are closely involved in the software development process

## Opsteady Structure

Opsteady focuses on **what** constitutes a platform, the so-called layers of the platform, as well as **how** to deliver a platform.

### Platform Layers

An Opsteady platform consists of four layers: **foundation**, **Kubernetes**, **capabilities** and **interface**

### Foundation

This layer focuses on delivering the Cloud infrastructure needed for the platform. Fundamental pieces of infrastructure, like networks, DNS zones, connectivity are all part of the foundation.

### Kubernetes

The Kubernetes layer provides a Kubernetes cluster on top of the foundation with all the necessary features like nodes, load balancing, storage, etc.

### Capabilities

A capabilities layer is added to the Kubernetes cluster to ensure the platform is useful to its users. Concerns like DNS, certificates, ingress, logging, metrics, and more are all part of this layer.

### Interface

The capabilities need to be exposed and integrated via a unified UI to configure different aspects of the platform. They also need to be made consistently available through the Kubernetes API. The interface layer takes care of this.

[The list of all the features](docs/opsteady/features.md)

### Platform Delivery

Opsteady uses a [management environment](../adr/../opsteady/docs/adr/0006-management-environment.md) together with an automated [CI/CD](../adr/../opsteady/docs/adr/0012-ci-cd.md) process and a [custom CLI](../adr/../opsteady/docs/adr/0016-cli.md) to standardize the process of creating hundreds of platforms in a robust and scalable way.

## Why is Opsteady open source?

We are standing on the shoulders of giants by using a lot of open source projects within the Opsteady platform. Therefore, we feel it's logical to contribute back as much as we can to the open source community. We have already contributed to various open source projects in the past and will continue to do so, while also creating personal open source projects. We welcome any contributions to the Opsteady platform and look forward to collaborating with the community.

With Opsteady we want to showcase a full-featured container platform and what some of the possibilities and trade-offs are while building one. We also want to show how a platform evolves after being created and emphasize the why of all the decisions that we make while creating the platform.

## Origins

The initial idea for a platform started in 2015 and evolved over many years of us helping organizations improve their software delivery processes. We focused on delivering and maintaining platforms, while at the same time also changing the way organizations work. We were and still are, always looking to increase the quality of the software delivery lifecycle to make sure that developers have a delightful experience while building and operating their software.

Opsteady is a creation of Ivo and Armin ([author of: Container platform for the Enterprise](http://leanpub.com/cpe)), powered by Xebia.

## Getting started

- [Architectural decisions](docs/adr/toc.md)
- [Opsteady documentation](docs/opsteady/toc.md)
