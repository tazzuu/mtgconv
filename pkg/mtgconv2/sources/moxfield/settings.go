package moxfield

import (
	"time"
	"golang.org/x/time/rate"
	"mtgconv/pkg/mtgconv2/core"
)

var MoxfieldBaseUrl string = "https://api.moxfield.com/v2/decks/all" // look up a specific deck
var MoxfieldDeckSearchUrl string = "https://api2.moxfield.com/v2/decks/search" // search all decks
var ApiSource core.APISource = core.SourceMoxfield
// API rate limit 1 query per second for Moxfield
var MoxfieldAPIRateLimiter = rate.NewLimiter(rate.Every(time.Second), 1)

