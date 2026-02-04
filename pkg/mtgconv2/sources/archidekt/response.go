package archidekt

import "time"

// curl -s -H "Accept: application/json" https://archidekt.com/api/decks/19632970/
type DeckResponse struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeckFormat     int             `json:"deckFormat"`
	EdhBracket     int             `json:"edhBracket"`
	Game           *string         `json:"game"`
	Description    string          `json:"description"`
	ViewCount      int             `json:"viewCount"`
	Featured       string          `json:"featured"`
	CustomFeatured string          `json:"customFeatured"`
	Private        bool            `json:"private"`
	Unlisted       bool            `json:"unlisted"`
	Theorycrafted  bool            `json:"theorycrafted"`
	Points         int             `json:"points"`
	UserInput      int             `json:"userInput"`
	Owner          Owner           `json:"owner"`
	CommentRoot    int             `json:"commentRoot"`
	Editors        *string         `json:"editors"`
	ParentFolder   int             `json:"parentFolder"`
	Bookmarked     bool            `json:"bookmarked"`
	Categories     []Category      `json:"categories"`
	DeckTags       []string        `json:"deckTags"`
	PlaygroupDeck  *string         `json:"playgroupDeckUrl"`
	CardPackage    *string         `json:"cardPackage"`
	Cards          []DeckCardEntry `json:"cards"`
}

type Owner struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	Avatar       string  `json:"avatar"`
	Frame        *string `json:"frame"`
	CKAffiliate  string  `json:"ckAffiliate"`
	TCGAffiliate string  `json:"tcgAffiliate"`
	ReferrerEnum *string `json:"referrerEnum"`
}

type Category struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	IsPremier        bool   `json:"isPremier"`
	IncludedInDeck   bool   `json:"includedInDeck"`
	IncludedInPrice  bool   `json:"includedInPrice"`
}

type DeckCardEntry struct {
	ID               int      `json:"id"`
	Categories       []string `json:"categories"`
	Companion        bool     `json:"companion"`
	FlippedDefault   bool     `json:"flippedDefault"`
	Label            string   `json:"label"`
	Modifier         string   `json:"modifier"`
	Quantity         int      `json:"quantity"`
	CustomCmc        *int     `json:"customCmc"`
	RemovedCategories *string `json:"removedCategories"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DeletedAt        *time.Time `json:"deletedAt"`
	Notes            *string   `json:"notes"`
	Card             Card      `json:"card"`
	Owned            int       `json:"owned"`
	PinnedStatus     int       `json:"pinnedStatus"`
	Prices           Prices    `json:"prices"`
	Rarity           string    `json:"rarity"`
	GlobalCategories []string  `json:"globalCategories"`
}

type Card struct {
	ID              int         `json:"id"`
	Artist          string      `json:"artist"`
	TcgProductID    int         `json:"tcgProductId"`
	CkFoilID        int         `json:"ckFoilId"`
	CkNormalID      int         `json:"ckNormalId"`
	CmEd            *string     `json:"cmEd"`
	ScgSku          string      `json:"scgSku"`
	ScgFoilSku      *string     `json:"scgFoilSku"`
	CollectorNumber string      `json:"collectorNumber"`
	MultiverseID    int         `json:"multiverseid"`
	MtgoFoilID      int         `json:"mtgoFoilId"`
	MtgoNormalID    int         `json:"mtgoNormalId"`
	UID             string      `json:"uid"`
	DisplayName     *string     `json:"displayName"`
	ReleasedAt      string      `json:"releasedAt"`
	ContentWarning  bool        `json:"contentWarning"`
	Edition         Edition     `json:"edition"`
	Flavor          string      `json:"flavor"`
	Games           []string    `json:"games"`
	Options         []string    `json:"options"`
	ScryfallImageHash string    `json:"scryfallImageHash"`
	OracleCard      OracleCard  `json:"oracleCard"`
	Owned           int         `json:"owned"`
	PinnedStatus    int         `json:"pinnedStatus"`
	Prices          Prices      `json:"prices"`
	Rarity          string      `json:"rarity"`
	GlobalCategories []string   `json:"globalCategories"`
}

type Edition struct {
	EditionCode string  `json:"editioncode"`
	EditionName string  `json:"editionname"`
	EditionDate string  `json:"editiondate"`
	EditionType string  `json:"editiontype"`
	MtgoCode    *string `json:"mtgoCode"`
}

type OracleCard struct {
	ID              int         `json:"id"`
	Cmc             int         `json:"cmc"`
	ColorIdentity   []string    `json:"colorIdentity"`
	Colors          []string    `json:"colors"`
	EdhrecRank      int         `json:"edhrecRank"`
	Faces           []any       `json:"faces"`
	Layout          string      `json:"layout"`
	UID             string      `json:"uid"`
	Legalities      Legalities  `json:"legalities"`
	ManaCost        string      `json:"manaCost"`
	ManaProduction  ManaProduction `json:"manaProduction"`
	Name            string      `json:"name"`
	Power           string      `json:"power"`
	Salt            float64     `json:"salt"`
	SubTypes        []string    `json:"subTypes"`
	SuperTypes      []string    `json:"superTypes"`
	Keywords        []string    `json:"keywords"`
	Text            string      `json:"text"`
	Tokens          []any       `json:"tokens"`
	Toughness       string      `json:"toughness"`
	Types           []string    `json:"types"`
	Loyalty         *string     `json:"loyalty"`
	CanlanderPoints *int        `json:"canlanderPoints"`
	IsPDHCommander  bool        `json:"isPDHCommander"`
	DefaultCategory *string     `json:"defaultCategory"`
	GameChanger     bool        `json:"gameChanger"`
	ExtraTurns      bool        `json:"extraTurns"`
	Tutor           bool        `json:"tutor"`
	MassLandDenial  bool        `json:"massLandDenial"`
	TwoCardComboSingleton bool  `json:"twoCardComboSingelton"`
	TwoCardComboIDs []int       `json:"twoCardComboIds"`
	AtomicCombos    []any       `json:"atomicCombos"`
	PotentialCombos []any       `json:"potentialCombos"`
	Lang            string      `json:"lang"`
}

type Legalities struct {
	Alchemy          string  `json:"alchemy"`
	Legacy           string  `json:"legacy"`
	Oldschool        string  `json:"oldschool"`
	Modern           string  `json:"modern"`
	Vintage          string  `json:"vintage"`
	Oathbreaker      string  `json:"oathbreaker"`
	OneVOne          string  `json:"1v1"`
	HistoricBrawl    string  `json:"historicbrawl"`
	Premodern        string  `json:"premodern"`
	Historic         string  `json:"historic"`
	Commander        string  `json:"commander"`
	PauperCommander  string  `json:"paupercommander"`
	Gladiator        string  `json:"gladiator"`
	Explorer         *string `json:"explorer"`
	Brawl            string  `json:"brawl"`
	Penny            string  `json:"penny"`
	Pioneer          string  `json:"pioneer"`
	Duel             string  `json:"duel"`
	Pauper           string  `json:"pauper"`
	Standard         string  `json:"standard"`
	Future           string  `json:"future"`
	PreDH            string  `json:"predh"`
	Timeless         string  `json:"timeless"`
	Canlander        string  `json:"canlander"`
}

type ManaProduction struct {
	W *int `json:"W"`
	U *int `json:"U"`
	B *int `json:"B"`
	R *int `json:"R"`
	G *int `json:"G"`
	C *int `json:"C"`
}

type Prices struct {
	CK         float64 `json:"ck"`
	CKFoil     float64 `json:"ckfoil"`
	CM         float64 `json:"cm"`
	CMFoil     float64 `json:"cmfoil"`
	MTGO       float64 `json:"mtgo"`
	MTGOFoil   float64 `json:"mtgofoil"`
	TCG        float64 `json:"tcg"`
	TCGFoil    float64 `json:"tcgfoil"`
	SCG        float64 `json:"scg"`
	SCGFoil    float64 `json:"scgfoil"`
	MP         float64 `json:"mp"`
	MPFoil     float64 `json:"mpfoil"`
	TCGLand    float64 `json:"tcgLand"`
	TCGLandFoil float64 `json:"tcgLandFoil"`
}


// curl 'https://archidekt.com/api/decks/v3/?deckFormat=3&edhBracket=5&orderBy=-updatedAt&page=1'
type DeckSearchResponse struct {
	Count   int              `json:"count"`
	Next    *string          `json:"next"`
	Results []DeckSearchItem `json:"results"`
}

type DeckSearchItem struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Size          int           `json:"size"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	CreatedAt     time.Time     `json:"createdAt"`
	DeckFormat    int           `json:"deckFormat"`
	EdhBracket    int           `json:"edhBracket"`
	Featured      string        `json:"featured"`
	CustomFeatured string       `json:"customFeatured"`
	ViewCount     int           `json:"viewCount"`
	Private       bool          `json:"private"`
	Unlisted      bool          `json:"unlisted"`
	Theorycrafted bool          `json:"theorycrafted"`
	Game          *int          `json:"game"`
	HasDescription bool         `json:"hasDescription"`
	Tags          []DeckTag     `json:"tags"`
	ParentFolderID int          `json:"parentFolderId"`
	Owner         SearchOwner   `json:"owner"`
	Colors        DeckColors    `json:"colors"`
	CardPackage   *string       `json:"cardPackage"`
	Contest       *string       `json:"contest"`
}

type DeckTag struct {
	ID       int    `json:"id"`
	Tag      int    `json:"tag"`
	Deck     int    `json:"deck"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type SearchOwner struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Avatar      string   `json:"avatar"`
	Moderator   bool     `json:"moderator"`
	PledgeLevel *int     `json:"pledgeLevel"`
	Roles       []string `json:"roles"`
}

type DeckColors struct {
	W int `json:"W"`
	U int `json:"U"`
	B int `json:"B"`
	R int `json:"R"`
	G int `json:"G"`
}
