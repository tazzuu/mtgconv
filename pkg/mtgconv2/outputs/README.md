To add a new output format

- mkdir new package in subdir here
- create `handler.go` ; see example from `txt/handler.go`
  - create `Handler` with the required methods from `core/interfaces.go` -> `core....`
  - make sure to inclue `init` to register the handler at launch
- add the package to the import in `all/handlers.go`
- add to the list of output formats in `core/types.go` -> `core....`