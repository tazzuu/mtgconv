package mtgconv

import (
	"fmt"
	"os"
	"encoding/json"
)

func MakeMoxfieldDeckResponse (jsonStr string) DeckResponse {
	var deck DeckResponse
	if err := json.Unmarshal([]byte(jsonStr), &deck); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON response: %v\n", err)
		os.Exit(1)
	}
	deck.JsonStr = jsonStr
	return deck
}

type DeckResponse struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	Description        string           `json:"description"`
	Format             string           `json:"format"`
	Visibility         string           `json:"visibility"`
	PublicURL          string           `json:"publicUrl"`
	PublicID           string           `json:"publicId"`
	LikeCount          int              `json:"likeCount"`
	ViewCount          int              `json:"viewCount"`
	CommentCount       int              `json:"commentCount"`
	AreCommentsEnabled bool             `json:"areCommentsEnabled"`
	IsShared           bool             `json:"isShared"`
	AuthorsCanEdit     bool             `json:"authorsCanEdit"`
	JsonStr string // the original JSON from the request response

	CreatedByUser User `json:"createdByUser"`
	Authors       []User `json:"authors"`
	RequestedAuthors []User `json:"requestedAuthors"`

	Main Card `json:"main"`

	MainboardCount int `json:"mainboardCount"`
	Mainboard      map[string]DeckEntry `json:"mainboard"`

	Hubs []Hub `json:"hubs"`

	CreatedAtUTC     string `json:"createdAtUtc"`
	LastUpdatedAtUTC string `json:"lastUpdatedAtUtc"`
	ExportID         string `json:"exportId"`

	AuthorTags map[string]any `json:"authorTags"`
	IsTooBeaucoup bool `json:"isTooBeaucoup"`

	Affiliates map[string]string `json:"affiliates"`

	MainCardIdIsBackFace     bool `json:"mainCardIdIsBackFace"`
	AllowPrimerClone         bool `json:"allowPrimerClone"`
	EnableMultiplePrintings  bool `json:"enableMultiplePrintings"`
	IncludeBasicLandsInPrice bool `json:"includeBasicLandsInPrice"`
	IncludeCommandersInPrice bool `json:"includeCommandersInPrice"`
	IncludeSignatureSpellsInPrice bool `json:"includeSignatureSpellsInPrice"`

	Media []any `json:"media"`

	CommandersCount int `json:"commandersCount"`
	Commanders map[string]DeckEntry `json:"commanders"`

	SideboardCount int `json:"sideboardCount"`
	Sideboard map[string]DeckEntry `json:"sideboard"`
}

type User struct {
	UserName        string        `json:"userName"`
	DisplayName     string        `json:"displayName"`
	ProfileImageURL string        `json:"profileImageUrl"`
	Badges          []any         `json:"badges"`
}

type Hub struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeckEntry struct {
	Quantity int    `json:"quantity"`
	BoardType string `json:"boardType"`
	Finish   string `json:"finish"`
	IsFoil   bool   `json:"isFoil"`
	IsAlter  bool   `json:"isAlter"`
	IsProxy  bool   `json:"isProxy"`
	Card     Card   `json:"card"`

	UseCmcOverride          bool `json:"useCmcOverride"`
	UseManaCostOverride     bool `json:"useManaCostOverride"`
	UseColorIdentityOverride bool `json:"useColorIdentityOverride"`
	ExcludedFromColor       bool `json:"excludedFromColor"`
}

type Card struct {
	ID          string `json:"id"`
	UniqueCardID string `json:"uniqueCardId"`
	ScryfallID  string `json:"scryfall_id"`

	Set     string `json:"set"`
	SetName string `json:"set_name"`
	Name    string `json:"name"`
	CN      string `json:"cn"`
	Layout  string `json:"layout"`
	CMC     float64 `json:"cmc"`
	Type    string `json:"type"`
	TypeLine string `json:"type_line"`

	OracleText string `json:"oracle_text"`
	ManaCost   string `json:"mana_cost"`
	Power      string `json:"power"`
	Toughness  string `json:"toughness"`

	Colors        []string `json:"colors"`
	ColorIndicator []string `json:"color_indicator"`
	ColorIdentity []string `json:"color_identity"`

	Legalities map[string]string `json:"legalities"`

	Frame       string `json:"frame"`
	FrameEffects []string `json:"frame_effects"`
	Reserved    bool `json:"reserved"`
	Digital     bool `json:"digital"`
	Foil        bool `json:"foil"`
	Nonfoil     bool `json:"nonfoil"`
	Etched      bool `json:"etched"`
	Glossy      bool `json:"glossy"`
	Rarity      string `json:"rarity"`
	BorderColor string `json:"border_color"`
	Colorshifted bool `json:"colorshifted"`
	Lang        string `json:"lang"`
	Latest      bool `json:"latest"`
	HasMultipleEditions bool `json:"has_multiple_editions"`
	HasArenaLegal bool `json:"has_arena_legal"`

	Prices map[string]any `json:"prices"`

	CardFaces []any `json:"card_faces"`
	Artist    string `json:"artist"`
	PromoTypes []string `json:"promo_types"`

	CardHoarderURL   string `json:"cardHoarderUrl"`
	CardKingdomURL   string `json:"cardKingdomUrl"`
	CardKingdomFoilURL string `json:"cardKingdomFoilUrl"`
	CardMarketURL    string `json:"cardMarketUrl"`
	TCGPlayerURL     string `json:"tcgPlayerUrl"`

	IsArenaLegal bool `json:"isArenaLegal"`
	ReleasedAt  string `json:"released_at"`

	EDHRecRank   int `json:"edhrec_rank"`
	MultiverseIDs []int `json:"multiverse_ids"`

	CardMarketID int `json:"cardmarket_id"`
	MTGOID       int `json:"mtgo_id"`
	ArenaID      int `json:"arena_id"`
	TCGPlayerID  int `json:"tcgplayer_id"`
	CardKingdomID int `json:"cardkingdom_id"`
	CardKingdomFoilID int `json:"cardkingdom_foil_id"`

	Reprint bool `json:"reprint"`
	SetType string `json:"set_type"`

	CoolStuffIncURL    string `json:"coolStuffIncUrl"`
	CoolStuffIncFoilURL string `json:"coolStuffIncFoilUrl"`

	Acorn bool `json:"acorn"`

	ImageSeq int `json:"image_seq"`

	CardTraderURL     string `json:"cardTraderUrl"`
	CardTraderFoilURL string `json:"cardTraderFoilUrl"`

	ContentWarning bool `json:"content_warning"`

	IsPauperCommander bool `json:"isPauperCommander"`
	IsCovered         bool `json:"isCovered"`
	AllRaritiesMask   int  `json:"allRaritiesMask"`
	ManapoolURL       string `json:"manapool_url"`
	IsToken           bool   `json:"isToken"`
	DefaultFinish     string `json:"defaultFinish"`
}
