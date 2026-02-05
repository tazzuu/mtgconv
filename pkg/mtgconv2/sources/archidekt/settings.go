package archidekt

import (
	"time"
	"golang.org/x/time/rate"
	"mtgconv/pkg/mtgconv2/core"
)

// fetch one deck
// $ curl -s -H "Accept: application/json" https://archidekt.com/api/decks/11516539/ | jq . > response.json

// search for decks
// curl 'https://archidekt.com/api/decks/v3/?deckFormat=3&edhBracket=5&orderBy=-updatedAt&page=1' | jq | less

// reference; https://github.com/linkian209/pyrchidekt/blob/main/pyrchidekt/formats.py
// https://github.com/topics/archidekt
// https://archidekt.com/forum/thread/40353

var DeckPublicUrlBase string = "https://archidekt.com/decks/"
var DeckFetchUrl string = "https://archidekt.com/api/decks/" // look up a specific deck
var DeckSearchUrl string = "https://archidekt.com/api/decks/v3/" // search all decks
var ApiSource core.APISource = core.SourceArchidekt
// API rate limit 1 query per second
var APIRateLimiter = rate.NewLimiter(rate.Every(2000*time.Millisecond), 1)


