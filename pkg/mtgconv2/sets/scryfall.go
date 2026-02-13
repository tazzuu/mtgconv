package sets

import (
	"encoding/json"
	"strings"
)

// attempt to retrieve sets from scryfall API https://scryfall.com/docs/api/sets/all
// fallback to load from static file

// TODO: use go:generate to convert the JSON into static go code instead to pre-build the index at compile time

func parseSetsJSON(setsJSONstr string) (SetList, error) {
	var list SetList
	err := json.NewDecoder(strings.NewReader(setsJSONstr)).Decode(&list)
	return list, err
}

func GetSetIndex() *SetIndex {
	sets, _ := AllSets()
	return buildSetIndex(sets)
}

func AllSets() ([]Set, error) {
	setList, err := parseSetsJSON(string(setsJSONstr))
	if err != nil {
		return []Set{}, nil
	}
	return setList.Data, nil
}

func buildSetIndex(sets []Set) *SetIndex {
	index := &SetIndex{
		ByName: make(map[string]*Set),
		ByCode: make(map[string]*Set),
	}
	for i := range sets {
		set := &sets[i]
		index.ByName[strings.ToLower(set.Name)] = set
		index.ByCode[strings.ToLower(set.Code)] = set
	}
	return index
}

// model to parse the JSON
type SetList struct {
	Object  string `json:"object"`
	HasMore bool   `json:"has_more"`
	Data    []Set  `json:"data"`
}

// object from the JSON
type Set struct {
	Object        string  `json:"object"`
	ID            string  `json:"id"`
	Code          string  `json:"code"`
	MTGOCode      *string `json:"mtgo_code,omitempty"`
	ArenaCode     *string `json:"arena_code,omitempty"`
	TCGPlayerID   *int    `json:"tcgplayer_id,omitempty"`
	Name          string  `json:"name"`
	// URI           string  `json:"uri"` // skip these they take up a lot of space
	// ScryfallURI   string  `json:"scryfall_uri"`
	// SearchURI     string  `json:"search_uri"`
	ReleasedAt    *string `json:"released_at,omitempty"`
	SetType       string  `json:"set_type"`
	CardCount     int     `json:"card_count"`
	ParentSetCode *string `json:"parent_set_code,omitempty"`
	Digital       bool    `json:"digital"`
	NonfoilOnly   bool    `json:"nonfoil_only"`
	FoilOnly      bool    `json:"foil_only"`
	// IconSVGURI    string  `json:"icon_svg_uri"`
	BlockCode     *string `json:"block_code,omitempty"`
	Block         *string `json:"block,omitempty"`
}

// map for lookups
type SetIndex struct {
	ByName map[string]*Set
	ByCode map[string]*Set
}
func (s SetIndex) NameExists(name string) bool {
	_, ok := s.ByName[strings.ToLower(name)]
	return ok
}
func (s SetIndex) CodeExists(code string) bool {
	_, ok := s.ByCode[strings.ToLower(code)]
	return ok
}
func (s SetIndex) GetByName(name string) *Set {
	val, _ := s.ByName[strings.ToLower(name)]
	return val
}
func (s SetIndex) GetByCode(code string) *Set {
	val, _ := s.ByCode[strings.ToLower(code)]
	return val
}
