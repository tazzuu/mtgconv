package moxfieldcollection

// output row format
type MoxfieldCollectionRow struct {
	Count          int     `csv:"Count"`
	TradelistCount int     `csv:"Tradelist Count"`
	Name           string  `csv:"Name"`
	Edition        string  `csv:"Edition"`
	Condition      string  `csv:"Condition"`
	Language       string  `csv:"Language"`
	Foil           string  `csv:"Foil"` // empty or "foil"
	Tags           string  `csv:"Tags"`
	LastModified   string  `csv:"Last Modified"` // or time.Time with parsing
	CollectorNumber string `csv:"Collector Number"`
	Alter          string  `csv:"Alter"`
	Proxy          string  `csv:"Proxy"`         // "TRUE"/"FALSE" or empty
	PlaytestCard   string  `csv:"Playtest Card"` // "TRUE"/"FALSE" or empty
	PurchasePrice  string  `csv:"Purchase Price"`

}