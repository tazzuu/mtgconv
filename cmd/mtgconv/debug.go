package main

import (
	"fmt"
	"log/slog"
	"mtgconv/pkg/mtgconv2/core"
	"mtgconv/pkg/mtgconv2/sets"
)

func debug(config core.Config){
	slog.Debug("debug func got config", "config", config)
	fmt.Println(sets.AllSets())
}