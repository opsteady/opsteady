# Project guidelines

These project guidelines are a direct implementation of the [ADR project guidelines](../adr/0002-project-guidelines.md)

## IDE (Visual Studio Code)

Everyone is entitled to their IDE to work with but having a common IDE that is supported by a team/project allows for a lot of standardization which leads to a much smoother way of working and helps speed up the onboarding of new members.
Visual Studio Code is free, open-source and has a big community. It also has a big library of useful extensions, it allows projects to be run in containers and many more features which make it very suitable as a default IDE.

All the recommended extensions are [saved](../../.vscode/extensions.json).

## Validation

We want consistency in the project regardless of who is creating the code, therefore we are using different linters and validations for different technologies. We are using [hadolint](https://github.com/hadolint/hadolint) for validating Dockerfiles. [golangci-lint](https://github.com/golangci/golangci-lint) for validating Go code. [yamllint](https://github.com/adrienverge/yamllint) for validating yaml files.
