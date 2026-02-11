To add a new input source

- mkdir new package in subdir here
- create `handler.go` ; see example from `moxfield/handler.go`
  - create `Handler` with the required methods from `core/interfaces.go` -> `core.SourceHandler`
  - make sure to inclue `init` to register the handler at launch
- add the package to the import in `all/handlers.go`
- add to the list of input formats in `core/types.go` -> `core.InputFormat`