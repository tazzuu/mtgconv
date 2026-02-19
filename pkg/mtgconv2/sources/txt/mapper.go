package txt

import (
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"mtgconv/pkg/mtgconv2/core"
)

// Matches:
// "1 Abrade" -> qty=1, name=Abrade
// "Abrade"   -> qty="", name=Abrade
var txtLineRE = regexp.MustCompile(`^\s*(?:(\d+)\s+)?(.+?)\s*$`)


func ParseTxtLine(line string) (TxtRow, error) {
	line = strings.TrimSpace(line)
	match := txtLineRE.FindStringSubmatch(line)
	if match == nil {
		return TxtRow{}, &core.LineParseError{Message: fmt.Sprintf("invalid txt line: %q", line)}
	}

	qty := 1 // default when no leading integer
	if match[1] != "" {
		n, err := strconv.Atoi(match[1])
		if err != nil {
			return TxtRow{}, &core.QuantityParseError{Quantity: match[1]}
		}
		qty = n
	}

	return TxtRow{Quantity: qty, Name: strings.TrimSpace(match[2])}, nil
}