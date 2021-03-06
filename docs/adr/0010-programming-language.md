# 10. Programming language

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27
Foundation on [0012-ci-cd.md](0012-ci-cd.md) on 2021-08-29
Foundation for [0016-cli.md](0016-cli.md) on 2021-09-26

## Context

Decide on a programming language inside the project.

## Decision

There is always a reason for using a different language for a different purpose but having a default goto language in a project allows for standardization.

When choosing a programming language we looked at python and golang. Python is very easy, has a ton of libraries, and is an interpreted language so it feels like bash. Golang is statically typed and a compiled language. Although it takes a bit more effort to run, golang being a statically typed language gives us build-time safety. Go is in our opinion as easy as python and has a lot of libraries to work with as well.
Most tools used by us in the container community space are golang applications, the chance is almost 100 percent that we will need to code in golang anyway. Therefore we are deciding to use golang as our default programming language.

It is possible to use other languages in the project if there is a reason for it. That decision should be discussed and documented with arguments in an ADR.

## Consequences

Everyone will need to learn Go.
