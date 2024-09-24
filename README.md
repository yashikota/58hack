# Chronotes

[![codecov](https://codecov.io/github/yashikota/chronotes/graph/badge.svg?token=8LK1D9KWN5)](https://codecov.io/github/yashikota/chronotes)
[![Go Report Card](https://goreportcard.com/badge/github.com/yashikota/chronotes)](https://goreportcard.com/report/github.com/yashikota/chronotes)

> [!NOTE]
> フロントエンドのリポジトリは [GenichiMaruo/chronotes-front](https://github.com/GenichiMaruo/chronotes-front) にあります。  

## Development

```sh
go install github.com/go-task/task/v3/cmd/task@latest
```

```txt
task: Available tasks for this project:
* all:                Run all tasks
* default:            Display this help message
* api:lint:           Lint API documentation
* api:split:          Split OpenAPI specification
* api:tsp:            Generate Open API from TypeSpec
* docker:build:       Build Docker image
* docker:dev:         Run development environment
* docker:lint:        Lint Dockerfile
* go:fmt:             Format Go code
* go:lint:            Lint Go code
* go:test:            Run Go tests
```
