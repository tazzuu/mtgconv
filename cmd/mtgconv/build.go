package main

import (
	"mtgconv/pkg/mtgconv2/core"
)

// set up the build options for the program

// overwrite this at build time ;
// -ldflags="-X 'main.Version=someversion'"
// also gets set by goreleaser; https://goreleaser.com/cookbooks/using-main.version/
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	program = "mtgconv"
)

// use this internally
var BuildInfo core.BuildInfo = core.BuildInfo{
	Version: version,
	Commit: commit,
	Date: date,
	Program: program,
}
