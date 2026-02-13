package shinycsv

import (
	"context"
	// "net/http"
	"log/slog"
	// "strconv"
	"os"
	"io"

	"github.com/gocarina/gocsv"

	"mtgconv/pkg/mtgconv2/core"

)

// input handler for decks from Moxfield

type Handler struct{}

func (h Handler) Source() core.InputSource {
	return core.InputShinyCSV
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
	slog.Debug("Starting Shiny csv import")

	fileHandle, err := os.Open(filename)
	if err != nil {
		return core.Deck{}, err
	}
	defer fileHandle.Close()

	rows := []*ShinyRow{}

	// https://github.com/gocarina/gocsv/issues/186
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in)
	})

	if err := gocsv.UnmarshalFile(fileHandle, &rows); err != nil {
		return core.Deck{}, err
	}
	slog.Debug("got rows", "n", len(rows))
	slog.Debug("first row", "row", rows[0])

	deck := ShinyRowsToCoreDeck(rows, cfg)

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
