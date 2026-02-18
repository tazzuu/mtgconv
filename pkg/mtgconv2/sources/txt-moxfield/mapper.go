package txtmoxfield

import (
	"fmt"
	"mtgconv/pkg/mtgconv2/core"
	"regexp"
	"strings"
	"strconv"
)

// Matches: 1 Abrade (VOW) 139
var txtMoxfieldLineRE = regexp.MustCompile(`^\s*(\d+)\s+(.+?)\s+\(([^)]+)\)\s+([^\s]+)(?:\s+.*)?\s*$`)

func ParseTxtLine(line string) (TxtRow, error) {
	line = strings.TrimSpace(line)
	match := txtMoxfieldLineRE.FindStringSubmatch(line)
	if match == nil {
		return TxtRow{}, &core.LineParseError{Message:fmt.Sprintf("invalid moxfield txt line: %q", line)}
	}


	name := match[2]
	set := match[3]
	collectorNum := match[4]

	quantity, err := strconv.Atoi(match[1])
	if err != nil {
		return TxtRow{}, &core.QuantityParseError{Quantity: fmt.Sprintf("invalid quantity %q: %w", match[1], err)}
	}



	return TxtRow{
		Name: name,
		Quantity: quantity,
		SetCode: set,
		CollectorNumber: collectorNum,
	}, nil
}