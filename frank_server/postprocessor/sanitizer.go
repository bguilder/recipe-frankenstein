package postprocessor

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// Sanitizer filters the results of the scraper
type Sanitizer struct {
	StopWords map[string]interface{}
}

func loadStopWords() map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(stopWords), &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// NewSanitizer returns a Sanitizer
func NewSanitizer() Sanitizer {
	// load the list of stop words
	return Sanitizer{StopWords: loadStopWords()}
}

// TODO:
// - Remove punctuation
// - Lower case everything
// -

func (s *Sanitizer) hasStopWord(word string) bool {

	for i := 0; i < len(s.StopWords); i++ {
		// Make sure the first characters is a letter
		if word != "" {
			if !unicode.IsLetter([]rune(word)[0]) {
				return true
			}
		}
		if _, ok := s.StopWords[word]; ok {
			fmt.Printf("has stop word... %s", word)
			return true
		}
	}
	return false
}

func (s *Sanitizer) RemovePunctuation(word string) string {
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		if unicode.IsPunct(runes[i]) {
			runes = append(runes[:i], runes[i+1:]...)
			i--
		}
	}
	return string(runes)
}

func (s *Sanitizer) ToLower(word string) string {
	return strings.ToLower(word)
}

const stopWords = `{
	"cup": {},
	"cups": {},
	"ounces": {},
	"ounce": {},
	"tablespoon": {},
	"tablespoons": {},
	"chopped": {},
	"teaspoon": {},
	"optional": {},
	"sliced": {},
	"pound": {},
	"taste": {},
	"cut": {},
	"frying": {},
	"needed": {},
	"cubed": {},
	"quartered": {},
	"torn": {},
	",": {},
	"-": {},
	"divided": {},
	"Goya": {},
	"ounce)": {},
	"ourselves": {},
	"hers": {},
	"between": {},
	"yourself": {},
	"but": {},
	"again": {},
	"there": {},
	"about": {},
	"once": {},
	"during": {},
	"out": {},
	"very": {},
	"having": {},
	"with": {},
	"they": {},
	"own": {},
	"an": {},
	"be": {},
	"some": {},
	"for": {},
	"do": {},
	"its": {},
	"yours": {},
	"such": {},
	"into": {},
	"of": {},
	"most": {},
	"itself": {},
	"other": {},
	"off": {},
	"is": {},
	"s": {},
	"am": {},
	"or": {},
	"who": {},
	"as": {},
	"from": {},
	"him": {},
	"each": {},
	"the": {},
	"until": {},
	"below": {},
	"are": {},
	"we": {},
	"these": {},
	"your": {},
	"through": {},
	"me": {},
	"were": {},
	"more": {},
	"this": {},
	"down": {},
	"should": {},
	"their": {},
	"while": {},
	"above": {},
	"both": {},
	"up": {},
	"to": {},
	"ours": {},
	"had": {},
	"she": {},
	"all": {},
	"no": {},
	"when": {},
	"at": {},
	"any": {},
	"before": {},
	"them": {},
	"same": {},
	"and": {},
	"been": {},
	"have": {},
	"in": {},
	"will": {},
	"on": {},
	"does": {},
	"yourselves": {},
	"then": {},
	"that": {},
	"because": {},
	"what": {},
	"over": {},
	"why": {},
	"so": {},
	"can": {},
	"did": {},
	"not": {},
	"now": {},
	"under": {},
	"he": {},
	"you": {},
	"herself": {},
	"has": {},
	"just": {},
	"where": {},
	"too": {},
	"only": {},
	"myself": {},
	"which": {},
	"those": {},
	"i": {},
	"after": {},
	"few": {},
	"whom": {},
	"t": {},
	"being": {},
	"if": {},
	"theirs": {},
	"my": {},
	"against": {},
	"a": {},
	"by": {},
	"doing": {},
	"it": {},
	"how": {},
	"further": {},
	"was": {},
	"here": {},
	"pinch": {},
	"freshly": {},
	"bunch": {},
	"teaspoons": {},
	"jar": {},
	"jars": {},
	"split": {},
	"lukewarm": {},
	"large": {},
	"crushed": {},
	"dried": {},
	"fresh": {},
	"finely": {},
	"small": {},
	"diced": {},
	"container": {},
	"pounds": {},
	"minced": {},
	"fine": {},
	"packed": {},
	"diagonally": {},
	"inch": {},
	"pieces": {},
	"medium": {},
	"package": {},
	"thinly": {},
	"thawed": {},
	"frozen": {},
	"sheets": {},
	"grated": {},
	"slightly": {},
	"beaten": {},
	"softened": {},
	"warm": {},
	"lightly": {},
	"thick": {},
	"cans": {},
	"drained": {},
	"cheap": {},
	"canned": {},
	"rinsed": {},
	"clove": {},
	"seeded": {},
	"fluid": {},
	"undrained": {},
	"crumbles": {},
	"packages": {},
	"envelope": {},
	"halves": {},
	"cut ": {},
	"cubes": {},
	"than": {}
}`
