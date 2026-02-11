package shinycsv

import (
	"regexp"
	"strings"
)

// remove trailing parenthesis; "Thornbite Staff (White Border)"
var reParen *regexp.Regexp = regexp.MustCompile(`\(.*\)`)

// remove trailing descriptors
// "Mountain - Full Art"
var reFullArt *regexp.Regexp = regexp.MustCompile(` - Full Art$`)
var reJPFullArt *regexp.Regexp = regexp.MustCompile(` - JP Full Art$`)


// catch some special cards with alternate names embedded
// "Godzilla, Doom Inevitable - Yidaro, Wandering Monster" -> Yidaro, Wandering Monster
var reAltName *regexp.Regexp = regexp.MustCompile(`^.* - `)

// NOTE: some of these are backwards;
// Volrath's Stronghold - 1998 Brian Selden (STH) -> Volrath's Stronghold
// Westvale Abbey (Display Commander) - Thick Stock -> Westvale Abbey

// Budoka Gardener // Dokai, Weaver of Life -> Budoka Gardener
var reSplit *regexp.Regexp = regexp.MustCompile(`\s*//.*$`)

// #13 -> 13
var reHash *regexp.Regexp = regexp.MustCompile(`^#`)


func ParseCardName(raw string) (string, bool) {
	var isFoil bool
	var newName string

	newName = CleanCardNames(raw)
	isFoil = strings.Contains(strings.ToLower(raw), "foil")

	return newName, isFoil
}

func CleanCardNames(raw string) string {
	newName := reParen.ReplaceAllString(raw, "")
	newName = reFullArt.ReplaceAllString(newName, "")
	newName = reJPFullArt.ReplaceAllString(newName, "")
	newName = reAltName.ReplaceAllString(newName, "")
	newName = reSplit.ReplaceAllString(newName, "")
	return strings.TrimSpace(newName)
}

func CleanCollectorNumber(raw string) string {
	var newVal string
	newVal = reHash.ReplaceAllString(newVal, "")
	return newVal
}