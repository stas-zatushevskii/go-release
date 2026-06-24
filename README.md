# go-release

## Configure yaml file first

1. change project_name
2. configure `builds.goos`, `builds.goarch`
3. check actual path to main.go in your project `builds.main`
4. configure path to project version variable in codebase `builds.ldflags`

## Create build 

```sh
goreleaser release --snapshot --clean
```

## Result:
* runs `go mod tidy`
* runs `lintes`
* runs `go test`
* build project, set version in project variable, zip build and cfg
