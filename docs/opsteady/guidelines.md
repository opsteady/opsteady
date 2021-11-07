# Project guidelines

These project guidelines are a direct implementation of the [ADR project guidelines](../adr/0002-project-guidelines.md)

## IDE (Visual Studio Code)

Everyone is entitled to their IDE to work with but having a common IDE that is supported by a team/project allows for a lot of standardization which leads to a much smoother way of working and helps speed up the onboarding of new members.
Visual Studio Code is free, open-source and has a big community. It also has a big library of useful extensions, it allows projects to be run in containers and many more features which make it very suitable as a default IDE.

All the recommended extensions are [saved](../../.vscode/extensions.json).

## Validation

We want consistency in the project regardless of who is creating the code, therefore we are using different linters and validations for different technologies. We are using [hadolint](https://github.com/hadolint/hadolint) for validating Dockerfiles. [golangci-lint](https://github.com/golangci/golangci-lint) for validating Go code. [yamllint](https://github.com/adrienverge/yamllint) for validating yaml files.

## Definition of Done (DoD)

We are creating a lot of different components as part of the platform, the further we go and new people join the harder it is to know what needs to be done to keep the consistency in the project. Therefore it is good to have a reminder of the things that might need to happen as a part of a DOD:

### Code
- Code is written
- Validation succeeds

### CI/CD
- Component has a cicd/component.go
- Component is added to components.go
- Component is tested
- Component is added to the platform pipeline

### Configuration
- Number of replicas
- Pod disruption budget
- (Anti-)affinity
- Pod Priority
- Component is running on our nodes
- Helm chart and containers images all retrieved from Opsteady container registry

### Deployment
- Helm chart and images are synced to Opsteady container registry
- Component works on both clouds

### Testing
- Test files with a readme are present (for as long as we do manual testing)

### Documentation
- ADR is written describing the reason for the component and any specific settings
- Added information to [security and compliance](security-and-compliance-list.md)
- Updated the [toc](toc.md)

### Lifecycle management
- Made sure renovate bot and our sync can do LCM for Helm, Docker, tools

DoD is extended when needed (later on when we start adding monitoring, dashboards, alerting, etc..)
