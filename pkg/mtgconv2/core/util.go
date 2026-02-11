package core

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"sort"
	"strings"
	"time"
	"regexp"
	// "errors"
)

func GetDateStr() string {
	return time.Now().Format("2006-01-02")
}

func ConvertDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func CollectAuthors(m DeckMeta) string {
	return strings.Join(m.Authors, ",")
}

func SortDeckEntries(entries []DeckEntry) []DeckEntry {
	out := make([]DeckEntry, len(entries))
	copy(out, entries)
	sort.Slice(out, func(i, j int) bool {
        a := out[i]
        b := out[j]
        return a.Card.Name < b.Card.Name
    })
	return out
}

func CollectCommanders(d Deck) []DeckEntry {
	// check if the deck has Commander section
	val, ok := d.Sections[BoardCommander]
	// If the key exists
	if ok {
		return SortDeckEntries(val)
	}
	return []DeckEntry{}
}

func CollectMainboard(d Deck) []DeckEntry {
	val, ok := d.Sections[BoardMain]
	if ok {
		return SortDeckEntries(val)
	}
	return []DeckEntry{}
}

func CollectSideboard(d Deck) []DeckEntry {
	val, ok := d.Sections[BoardSideboard]
	if ok {
		return SortDeckEntries(val)
	}
	return []DeckEntry{}
}


// returns the correct BoardType enum from the available options
func ParseBoardType(raw string) (BoardType, error) {
	switch strings.ToLower(raw) {
	case string(BoardMain):
		return BoardMain, nil
	case string(BoardCommander):
		return BoardCommander, nil
	case string(BoardSideboard):
		return BoardSideboard, nil
	default:
		return "", &UnknownBoardType{raw}
	}
}

// returns the correct FinishType enum from the available options
func ParseFinishType(raw string) (FinishType, error) {
	switch strings.ToLower(raw) {
	case string(FinishDefault):
		return FinishDefault, nil
	case string(FinishFoil):
		return FinishFoil, nil
	case string(FinishNonfoil):
		return FinishNonfoil, nil
	case string(FinishEtched):
		return FinishEtched, nil
	default:
		return "", &UnknownFinishType{raw}
	}
}

// returns the correct FinishType enum from the available options
func ParseSortType(raw string) (SortType, error) {
	switch strings.ToLower(raw) {
	case string(SortLikes):
		return SortLikes, nil
	case string(SortViews):
		return SortViews, nil
	case string(SortUpdated):
		return SortUpdated, nil
	default:
		return "", &UnknownSortType{raw}
	}
}

func ParseSortDirection(raw string) (SortDirection, error) {
	switch strings.ToLower(raw) {
	case string(SortAsc):
		return SortAsc, nil
	case string(SortDesc):
		return SortDesc, nil
	default:
		return "", &UnknownSortDirection{raw}
	}
}

func ParseDeckFormat(raw string) (DeckFormat, error) {
	switch strings.ToLower(raw) {
	case string(DeckFormatCommander):
		return DeckFormatCommander, nil
	default:
		return "", &UnknownDeckFormat{raw}
	}
}

func ParseBracket(raw int) (CommanderBracket, error) {
	switch raw {
	case 0: // unset Bracket defaults to Bracket 1
		return CommanderBracket1, nil
	case int(CommanderBracket1):
		return CommanderBracket1, nil
	case int(CommanderBracket2):
		return CommanderBracket2, nil
	case int(CommanderBracket3):
		return CommanderBracket3, nil
	case int(CommanderBracket4):
		return CommanderBracket4, nil
	case int(CommanderBracket5):
		return CommanderBracket5, nil
	default:
		return 0, &UnknownBracket{raw}
	}
}

// regex to match and non-alphanumeric characters
var safeFileRe = regexp.MustCompile(`[^A-Za-z0-9._-]+`)

// return a name that is safe for use as an output filename
func SanitizeFilename(name string) string {
	// remove leading and trailing whitespace
	n := strings.TrimSpace(name)
	// if empty string just return the name 'deck'
	if n == "" {
		return "deck"
	}
	// replace all non-alphanumeric characters with _
	n = safeFileRe.ReplaceAllString(n, "_")
	// trim excess leading and trailing characters
	n = strings.Trim(n, "._-")
	// double check that string is not empty
	if n == "" {
		return "deck"
	}
	// truncate to 120 characters
	if len(n) > 120 {
		n = n[:120]
	}
	return n
}

// return a safe formatted filename with embedded metadata and file extension
func GenerateSafeFilename(name string, version int, extension string) string {
	// NOTE: removed date // deck.Meta.Date.Format("20060102"), so that we can re-export more easily
	var output string = fmt.Sprintf(
		"%s_v%d.%s",
		SanitizeFilename(name),
		version,
		extension)
	return output
}

// returns indented JSON
func PrettyJSON(raw string) (string, error) {
    var out bytes.Buffer
    if err := json.Indent(&out, []byte(raw), "", "  "); err != nil {
        return "", err
    }
    return out.String(), nil
}

// saves text to a file
func SaveTxtToFile(filename string, input string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
		// NOTE: this attempt to return error from defer is tricky and suspicious
		defer func() {
			_ = out.Close()
		}()
	_, err = fmt.Fprintln(out, input)
	return err
}


// check if a file with the given name exists
// func FileExists(filename string) bool {
// 	_, error := os.Stat(filename)
// 	return !errors.Is(error, os.ErrNotExist)
// }

// if the card name has '//' in it, return the parts
func SplitMultiFaceName(raw string) []string {
	parts := strings.Split(raw, "//")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

