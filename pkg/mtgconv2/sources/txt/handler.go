package txt

import (
	"context"
	"strings"
	"log/slog"
	"os"
	"bufio"

	"mtgconv/pkg/mtgconv2/core"
)

// input handler for plain txt format input

type Handler struct{}

func (h Handler) Source() core.InputSource {
	return core.InputTxt
}


func (h Handler) Fetch(ctx context.Context, input string, cfg core.Config, ovrr core.DeckMeta) (core.Deck, error) {
	_ = ctx
	_ = input
	_ = cfg
	// TODO: return error that this method not implemented
	return core.Deck{}, nil
}

func (h Handler) Import(filename string, cfg core.Config) (core.Deck, error) {
	_ = filename
	_ = cfg
	slog.Debug("Starting txt import")

	// start deck object
	deck := core.Deck{
		Meta:     core.DeckMeta{Name: "Moxfield TXT Import"},
		Sections: map[core.BoardType][]core.DeckEntry{},
	}

	// open input file handle
	fileHandle, err := os.Open(filename)
	if err != nil {
		return core.Deck{}, err
	}
	defer fileHandle.Close()

	// start scanning input file
	scanner := bufio.NewScanner(fileHandle)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	//  initial internal buffer capacity (64 KB)
	// maximum token size allowed (1 MB)

	started := false
	haveCommander := false
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// stop after the first empty line
		if line == "" {
			// stop parsing after hitting an empty line
			if started {
				break
			}
			continue // ignore leading blank lines
		}
		started = true

		// parse the line
		row, err := ParseTxtLine(line)
		if err != nil {
			return core.Deck{}, err
		}

		// put the first item in the Commander board and the rest go in the Mainboard
		board := core.BoardMain
		if !haveCommander {
			board = core.BoardCommander
			haveCommander = true
		}

		// create the deck entry and card entry
		entry := core.DeckEntry{
			Quantity: row.Quantity,
			Board: board,
			Finish: core.FinishDefault,
			Card: core.Card{
				Name: row.Name,
			},
		}

		// add it to the deck
		if err := deck.AddToSection(board, entry); err != nil {
			return core.Deck{}, err
		}

		// check for line scanning errors
		if err := scanner.Err(); err != nil {
			return core.Deck{}, err
		}

	}

	return deck, nil
}

func (h Handler) Search(ctx context.Context, cfg core.Config, scfg core.SearchConfig) ([]core.DeckMeta, error) {
	_ = ctx
	_ = cfg

	return []core.DeckMeta{}, nil
}

func init() {
	core.RegisterSource(Handler{})
}
