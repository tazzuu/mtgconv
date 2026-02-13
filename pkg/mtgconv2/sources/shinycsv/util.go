package shinycsv

import (
	"regexp"
	"strings"
	"log/slog"
	"mtgconv/pkg/mtgconv2/sets"
)

// remove trailing parenthesis; "Thornbite Staff (White Border)" -> Thornbite Staff
var reParen *regexp.Regexp = regexp.MustCompile(`\s*\(.*$`) // `\(.*\)` ; remove anything inside (...)

// remove trailing descriptors
// "Mountain - Full Art"
var reFullArt *regexp.Regexp = regexp.MustCompile(` - Full Art$`)
var reJPFullArt *regexp.Regexp = regexp.MustCompile(` - JP Full Art$`)

// some cards have year and name embedded
// Volrath's Stronghold - 1998 Brian Selden (STH) -> Volrath's Stronghold
// Westvale Abbey (Display Commander) - Thick Stock -> Westvale Abbey
var reYearName *regexp.Regexp = regexp.MustCompile(`\s-\s(?:19|20)\d{2}\b.*$`)

// catch some special cards with alternate names embedded
// "Godzilla, Doom Inevitable - Yidaro, Wandering Monster" -> Yidaro, Wandering Monster
var reAltName *regexp.Regexp = regexp.MustCompile(`^.* - `)

// Budoka Gardener // Dokai, Weaver of Life -> Budoka Gardener
var reSplit *regexp.Regexp = regexp.MustCompile(`\s*//.*$`)

// Henzie ""Toolbox"" Torre -> Henzie "Toolbox" Torre
var reDoubleQuote *regexp.Regexp = regexp.MustCompile(`""`)

// #13 -> 13
var reHash *regexp.Regexp = regexp.MustCompile(`^#`)

var reToken *regexp.Regexp = regexp.MustCompile(` Token$`)
var reArtCard *regexp.Regexp = regexp.MustCompile(` Art Card$`)

// determine if a card is valid; invalid cards should get excluded from downstream processing
func ValidateCard(row *ShinyRow, includeTokens bool, includeArtCards bool) bool {
	// slog.Debug("validating Shiny row", "row", row)
	var isValid bool = true
	if row == nil {
		slog.Debug("Shiny ValidateCard: empty card entry", "row", row)
		return false // or true, whatever makes sense
	}
	if row.BrandName != "Magic The Gathering" {
		slog.Debug("Shiny ValidateCard: invalid brand name", "BrandName", row.BrandName)
		return false
	}
	if includeTokens == false {
		if reToken.MatchString(row.ProductName) {
			slog.Debug("Shiny ValidateCard: is Token", "ProductName", row.ProductName)
			return false
		}
	}
	if includeArtCards == false {
		if reArtCard.MatchString(row.ProductName) {
			slog.Debug("Shiny ValidateCard: is Art Card", "ProductName", row.ProductName)
			return false
		}
	}
	return isValid
}

func ParseCardName(raw string) (string, bool) {
	var isFoil bool
	var newName string

	newName = CleanCardNames(raw)
	isFoil = strings.Contains(strings.ToLower(raw), "foil")

	return newName, isFoil
}

// fix messed up set details in the card row
func ValidateSet(row *ShinyRow, setIndex *sets.SetIndex) (string, string, string) {
	var setName string
	var setCode string
	var setType string

	// HOTFIX: pretty much all the 'one' set cards are mislabeled so kick back empty string for them all
	if row.SetName == "Phyrexia: All Will Be One" {
		slog.Debug("Got invalid card; dropping set codes from all cards from set one", "row", row)
		return setName, setCode, setType
	}

	// if the name is actually the code
	if setIndex.CodeExists(setName) {
		// swap the name and the code
		setCode = setName
		// get the set details
		set := setIndex.GetByCode(setCode)
		setType = set.SetType
		// get the real name
		setName = set.Name
	} else if setIndex.NameExists(row.SetName) {
		set := setIndex.GetByName(row.SetName)
		setCode = set.Code
		setType = set.SetType
	}

	return setName, setCode, setType

}

func CleanCardNames(raw string) string {
	newName := reParen.ReplaceAllString(raw, "")
	newName = reFullArt.ReplaceAllString(newName, "")
	newName = reJPFullArt.ReplaceAllString(newName, "")
	newName = reYearName.ReplaceAllString(newName, "")
	newName = reAltName.ReplaceAllString(newName, "")
	newName = reSplit.ReplaceAllString(newName, "")
	newName = reDoubleQuote.ReplaceAllString(newName, `"`)
	return strings.TrimSpace(newName)
}

func CleanCollectorNumber(raw string) string {
	var newVal string
	newVal = reHash.ReplaceAllString(newVal, "")
	return newVal
}