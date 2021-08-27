# 10. Programming language

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27

## Context

Decide on a programming language inside the project.

## Decision

There is always a reason for using a different language for a different purpose but having a default goto language in a project allows for standardization.

When choosing a programming language we looked at python and golang. Python is very easy, has a ton of libraries, and is an interpreted language so it feels like bash. Golang is statically typed and a compiled language therefore takes a bit more effort to run but is as easy as python and has a lot of libraries to work with as well.
Most tools used by us in the container community space are golang applications, the chance is almost 100 percent that we will need to code in golang any way. Therefore we are deciding to use golang as our default programming language.

It is possible to use other languages in the project if there is a reason for it. That decision should be discussed and documented with arguments in an ADR.

## Consequences

Everyone will need to learn go.
