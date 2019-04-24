<!--
title: example-serverless-golang-api
description: Example port of the current Go API to Golang.
layout: Doc
framework: v1
platform: AWS
language: Go
authorLink: 'https://github.com/mmacheerpuppy'
authorName: 'Callum Dempsey Leach'
-->

## Guide to get going with Go

1. Install the Go langauge, make sure your GOPATH is set to where you want your go files to be.
2. Install `dep` the official tool for dependency management, if you're on windows you take the dep binary and copy it to your GOROOT bin dir (here)
3. Start a project with `dep init` within \$GOPATH/src/<Project Name> -- reason https://github.com/golang/dep/issues/2054#issuecomment-444066202. If you're cloning a repository from github then you should use the **case-sensitive** form \$GOPATH/src/github.com/username/projectname
4. Install the VS Code Go plugin because it rocks!

## Guide to building

Skip to step 3 if you just want to build a binary without any additional development time.

1. You should understand how dependencies work in Go. Fundementally in normal case, the import path for each `.go` file must be the path of the package relative to `src/` in `$GOPATH`. But since we're using deps this is not the case (https://github.com/golang/dep/blob/master/docs/FAQ.md#how-does-dep-work-without-changing-my-packages-imports). For instance, with deps, import github.com/user/awesome-project will be found in the project's /vendor/github.com/user/awesome-project before looking to \$GOPATH/src/github.com/user/awesome-project, before looking to HTTP resources. Please note that absolute paths are case sensitive for local package resolution.
2. Install dependencies with `dep ensure` (installs to vendor).
3. Make the makefile (use Cygwin if on Windows). The makefile will run `dep ensure -v` prior to build, so you can skip the second stage if all you want to do is create a binary. Binaries will be created in the `bin/` folder from Go packages at `pkg/`

## Guide to building a Lambda

1. Run `npm install && make` (Unix).
2. Push to a stage. This will push the compiled binary to AWS. Do not push without specifying a stage. You should know serverless framework well enough if you get this far.
