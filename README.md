# go-release

## Configure yaml file first

1. change project_name
2. configure `builds.goos`, `builds.goarch`
3. check actual path to main.go in your project `builds.main`

## Create build 

```sh
goreleaser release --snapshot --clean
```
