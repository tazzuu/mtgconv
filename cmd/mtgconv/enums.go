package main

import (
	"strings"
	"mtgconv/pkg/mtgconv2/core"
)

func InputSources() string {
	format := []string{}
	for _, v := range core.InputSources() {
		format = append(format, string(v))
	}
	return strings.Join(format, ",")
}

func OutputFormats() string {
	out := []string{}
	for _, v := range core.OutputFormats() {
		out = append(out, string(v))
	}
	return strings.Join(out, ",")
}

func DeckBrackets() string {
	out := []string{}
	for _, v := range core.CommanderBrackets() {
		out = append(out, v.String())
	}
	return strings.Join(out, ",")
}

func SortDirections() string {
	out := []string{}
	for _, v := range core.SortDirections() {
		out = append(out, string(v))
	}
	return strings.Join(out, ",")
}

func SearchSortTypes() string {
	out := []string{}
	for _, v := range core.SortTypes() {
		out = append(out, string(v))
	}
	return strings.Join(out, ",")
}

func SearchDeckFormats() string {
	out := []string{}
	for _, v := range core.DeckFormats() {
		out = append(out, string(v))
	}
	return strings.Join(out, ",")
}

func SearchAPISources() string {
	out := []string{}
	for _, v := range core.APISources() {
		out = append(out, string(v))
	}
	return strings.Join(out, ",")
}