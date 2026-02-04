package main

import (
	"mtgconv/pkg/mtgconv2/core"
)

// apply the cli context settings to the base default Config
func ApplyConfig(ctx Context) core.Config {
	config := core.DefaultConfig(BuildInfo)
	config.Debug = ctx.Debug
	config.Verbose = ctx.Verbose
	config.OutputDir = ctx.OutputDir
	config.UserAgent = ctx.UserAgent
	config.CompatibilityMode = ctx.CompatibilityMode
	config.SaveJSON = ctx.SaveJSON
	return config
}