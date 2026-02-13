package shinycsv

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanCardNames(t *testing.T) {
	tests := []struct {
		raw  string
		want string
	}{
		{"Thornbite Staff (White Border)", "Thornbite Staff"},
		{"Mountain - Full Art", "Mountain"},
		{"Forest - JP Full Art", "Forest"},
		{"Godzilla, Doom Inevitable - Yidaro, Wandering Monster", "Yidaro, Wandering Monster"},
		{"Budoka Gardener // Dokai, Weaver of Life", "Budoka Gardener"},
		{"Volrath's Stronghold - 1998 Brian Selden (STH)", "Volrath's Stronghold"},
		{"Westvale Abbey (Display Commander) - Thick Stock", "Westvale Abbey"},
		{`Henzie ""Toolbox"" Torre`, `Henzie "Toolbox" Torre`},

		{"Foo (bar) baz", "Foo"},
	}

	for _, tt := range tests {
		got := CleanCardNames(tt.raw)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("CleanCardNames(%q) mismatch (-want +got):\n%s", tt.raw, diff)
		}
	}
}

func TestValidateCard(t *testing.T) {
	tests := []struct{
		got ShinyRow
		want bool
	}{
		{ShinyRow{ProductName: "Forest", BrandName: "Magic The Gathering"}, true},
		{ShinyRow{ProductName: "Forest Token", BrandName: "Magic The Gathering"}, false},
		{ShinyRow{ProductName: "Pikachu", BrandName: "Pokemon"}, false},
		{ShinyRow{ProductName: "Forest", BrandName: "One Piece"}, false},
		{ShinyRow{ProductName: "Elspeth Resplendent Art Card", BrandName: "Magic The Gathering"}, false},
		{ShinyRow{ProductName: "Elspeth Resplendent", BrandName: "Magic The Gathering"}, true},
		{ShinyRow{ProductName: "Squirrel Token", BrandName: "Magic The Gathering"}, false},
		{ShinyRow{ProductName: "Squirrel", BrandName: "Magic The Gathering"}, true},


	}
	includeTokens := false
	includeArtCards := false
	for _, tt := range tests {
		got := ValidateCard(&tt.got, includeTokens, includeArtCards)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("CleanCardNames(%v) mismatch (-want +got):\n%s", tt.got, diff)
		}
	}
}

func TestParseCardNameFoil(t *testing.T) {
	gotName, gotFoil := ParseCardName("Island (Foil)")
	if diff := cmp.Diff("Island", gotName); diff != "" {
		t.Fatalf("ParseCardName name mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(true, gotFoil); diff != "" {
		t.Fatalf("ParseCardName foil mismatch (-want +got):\n%s", diff)
	}
}
