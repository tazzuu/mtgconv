package shinycsv

import (
	"strings"
)

type ShinyRow struct {
	ID            string  `csv:"id"`
	ProductName   string  `csv:"product_name"`
	SetName       string  `csv:"set_name"`
	BrandName     string  `csv:"brand_name"` // "Magic the Gathering"
	Discriminator string  `csv:"discriminator"`
	Rarity        string  `csv:"rarity"`
	Quantity      int     `csv:"quantity"`
	// ValueTotal    float64 `csv:"value_total"` // NOTE: ignore these monetary values because they are not always floats some are "Unavailable"
	// ValuePerUnit  float64 `csv:"value_per_unit"`
	// ValueCurrency string  `csv:"value_currency"`
	// PaidTotal     float64 `csv:"paid_total"`
	// PaidPerUnit   float64 `csv:"paid_per_unit"`
	// PaidCurrency  string  `csv:"paid_currency"`
	GradeType     string  `csv:"grade_type"`
	GradeSubtype  string  `csv:"grade_subtype"`
	GroupName     string  `csv:"group_name"`
	GroupWishlist bool    `csv:"group_wishlist"`
	DateAdded     string  `csv:"date_added"`
	Tag           string  `csv:"tag"`
}

func (s *ShinyRow) IsFoil() bool {
	return strings.Contains(strings.ToLower(s.Rarity), "foil")
}
func (s *ShinyRow) IsProxy() bool {
	return strings.Contains(strings.ToLower(s.Tag), "proxy")
}