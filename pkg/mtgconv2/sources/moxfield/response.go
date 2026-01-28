package moxfield

import (
	"encoding/json"
	"time"
)

// convert the Moxfield API response into an object
func MakeMoxfieldDeck(jsonStr string) (MoxfieldDeck, error) {
	var deck MoxfieldDeck
	err := json.Unmarshal([]byte(jsonStr), &deck)

	if err != nil {
		return MoxfieldDeck{}, err
	}
	deck.RetrievedAt = time.Now()
	return deck, nil
}

func MakeMoxfieldSeachResult(jsonStr string) ([]MoxfieldDeckSearchResult, error) {
	var results []MoxfieldDeckSearchResult
	err := json.Unmarshal([]byte(jsonStr), &results)

	if err != nil {
		return []MoxfieldDeckSearchResult{}, err
	}
	for i, _ := range results {
		results[i].RetrievedAt = time.Now()
	}

	return results, nil
}

// object type representing the fields present in the Moxfield API response
// NOTE: there are a lot more fields I have not included here
type MoxfieldDeck struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Format             string `json:"format"`
	Visibility         string `json:"visibility"`
	PublicURL          string `json:"publicUrl"`
	PublicID           string `json:"publicId"`
	LikeCount          int    `json:"likeCount"`
	ViewCount          int    `json:"viewCount"`
	CommentCount       int    `json:"commentCount"`
	AreCommentsEnabled bool   `json:"areCommentsEnabled"`
	IsShared           bool   `json:"isShared"`
	AuthorsCanEdit     bool   `json:"authorsCanEdit"`
	RetrievedAt time.Time

	CreatedByUser    MoxfieldUser   `json:"createdByUser"`
	Authors          []MoxfieldUser `json:"authors"`
	RequestedAuthors []MoxfieldUser `json:"requestedAuthors"`

	Main MoxfieldCard `json:"main"`

	MainboardCount int                  `json:"mainboardCount"`
	Mainboard      map[string]MoxfieldDeckEntry `json:"mainboard"`

	Hubs []MoxfieldHub `json:"hubs"`

	CreatedAtUTC     string `json:"createdAtUtc"`
	LastUpdatedAtUTC string `json:"lastUpdatedAtUtc"`
	ExportID         string `json:"exportId"`

	AuthorTags    map[string]any `json:"authorTags"`
	IsTooBeaucoup bool           `json:"isTooBeaucoup"`

	Affiliates map[string]string `json:"affiliates"`

	MainCardIdIsBackFace          bool `json:"mainCardIdIsBackFace"`
	AllowPrimerClone              bool `json:"allowPrimerClone"`
	EnableMultiplePrintings       bool `json:"enableMultiplePrintings"`
	IncludeBasicLandsInPrice      bool `json:"includeBasicLandsInPrice"`
	IncludeCommandersInPrice      bool `json:"includeCommandersInPrice"`
	IncludeSignatureSpellsInPrice bool `json:"includeSignatureSpellsInPrice"`

	Media []any `json:"media"`

	CommandersCount int                  `json:"commandersCount"`
	Commanders      map[string]MoxfieldDeckEntry `json:"commanders"`

	SideboardCount int                  `json:"sideboardCount"`
	Sideboard      map[string]MoxfieldDeckEntry `json:"sideboard"`

	Version int `json:"version"`
}

type MoxfieldUser struct {
	UserName        string `json:"userName"`
	DisplayName     string `json:"displayName"`
	ProfileImageURL string `json:"profileImageUrl"`
	Badges          []any  `json:"badges"`
}

type MoxfieldHub struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MoxfieldDeckEntry struct {
	Quantity  int    `json:"quantity"`
	BoardType string `json:"boardType"`
	Finish    string `json:"finish"`
	IsFoil    bool   `json:"isFoil"`
	IsAlter   bool   `json:"isAlter"`
	IsProxy   bool   `json:"isProxy"`
	Card      MoxfieldCard   `json:"card"`

	UseCmcOverride           bool `json:"useCmcOverride"`
	UseManaCostOverride      bool `json:"useManaCostOverride"`
	UseColorIdentityOverride bool `json:"useColorIdentityOverride"`
	ExcludedFromColor        bool `json:"excludedFromColor"`
}

type MoxfieldCard struct {
	ID           string `json:"id"`
	UniqueCardID string `json:"uniqueCardId"`
	ScryfallID   string `json:"scryfall_id"`

	Set      string  `json:"set"`
	SetName  string  `json:"set_name"`
	Name     string  `json:"name"`
	CN       string  `json:"cn"`
	Layout   string  `json:"layout"`
	CMC      float64 `json:"cmc"`
	Type     string  `json:"type"`
	TypeLine string  `json:"type_line"`

	OracleText string `json:"oracle_text"`
	ManaCost   string `json:"mana_cost"`
	Power      string `json:"power"`
	Toughness  string `json:"toughness"`

	Colors         []string `json:"colors"`
	ColorIndicator []string `json:"color_indicator"`
	ColorIdentity  []string `json:"color_identity"`

	Legalities map[string]string `json:"legalities"`

	Frame               string   `json:"frame"`
	FrameEffects        []string `json:"frame_effects"`
	Reserved            bool     `json:"reserved"`
	Digital             bool     `json:"digital"`
	Foil                bool     `json:"foil"`
	Nonfoil             bool     `json:"nonfoil"`
	Etched              bool     `json:"etched"`
	Glossy              bool     `json:"glossy"`
	Rarity              string   `json:"rarity"`
	BorderColor         string   `json:"border_color"`
	Colorshifted        bool     `json:"colorshifted"`
	Lang                string   `json:"lang"`
	Latest              bool     `json:"latest"`
	HasMultipleEditions bool     `json:"has_multiple_editions"`
	HasArenaLegal       bool     `json:"has_arena_legal"`

	Prices map[string]any `json:"prices"`

	CardFaces  []any    `json:"card_faces"`
	Artist     string   `json:"artist"`
	PromoTypes []string `json:"promo_types"`

	CardHoarderURL     string `json:"cardHoarderUrl"`
	CardKingdomURL     string `json:"cardKingdomUrl"`
	CardKingdomFoilURL string `json:"cardKingdomFoilUrl"`
	CardMarketURL      string `json:"cardMarketUrl"`
	TCGPlayerURL       string `json:"tcgPlayerUrl"`

	IsArenaLegal bool   `json:"isArenaLegal"`
	ReleasedAt   string `json:"released_at"`

	EDHRecRank    int   `json:"edhrec_rank"`
	MultiverseIDs []int `json:"multiverse_ids"`

	CardMarketID      int `json:"cardmarket_id"`
	MTGOID            int `json:"mtgo_id"`
	ArenaID           int `json:"arena_id"`
	TCGPlayerID       int `json:"tcgplayer_id"`
	CardKingdomID     int `json:"cardkingdom_id"`
	CardKingdomFoilID int `json:"cardkingdom_foil_id"`

	Reprint bool   `json:"reprint"`
	SetType string `json:"set_type"`

	CoolStuffIncURL     string `json:"coolStuffIncUrl"`
	CoolStuffIncFoilURL string `json:"coolStuffIncFoilUrl"`

	Acorn bool `json:"acorn"`

	ImageSeq int `json:"image_seq"`

	CardTraderURL     string `json:"cardTraderUrl"`
	CardTraderFoilURL string `json:"cardTraderFoilUrl"`

	ContentWarning bool `json:"content_warning"`

	IsPauperCommander bool   `json:"isPauperCommander"`
	IsCovered         bool   `json:"isCovered"`
	AllRaritiesMask   int    `json:"allRaritiesMask"`
	ManapoolURL       string `json:"manapool_url"`
	IsToken           bool   `json:"isToken"`
	DefaultFinish     string `json:"defaultFinish"`
}


type MoxfieldDeckSearchResult struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Format string `json:"format"`
	PublicUrl string `json:"publicUrl"`
	PublicId string `json:"publicId"`
	LikeCount int `json:"likeCount"`
	ViewCount int `json:"viewCount"`
	CommentCount int `json:"commentCount"`
	CreatedByUser    MoxfieldUser   `json:"createdByUser"`
	Authors          []MoxfieldUser `json:"authors"`
	IsLegal bool `json:"isLegal"`
	CreatedAtUTC     string `json:"createdAtUtc"`
	LastUpdatedAtUTC string `json:"lastUpdatedAtUtc"`
	MainboardCount int `json:"mainboardCount"`
	SideboardCount int `json:"sideboardCount"`
	MaybeboardCount int `json:"maybeboardCount"`
	Colors []string `json:"colors"`
	Bracket int `json:"bracket"`
	UserBracket int `json:"userBracket"`
	AutoBracket int `json:"autoBracket"`
	RetrievedAt time.Time
}