package data

// Phonetics contain information regarding the pronunciation of the word
type Phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

// Definitions contain information regarding the definition, synonims and uses of the word
type Definitions struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Example    string   `json:"example"`
}

// Meanings contain broader information including definitions
type Meanings struct {
	PartOfSpeech string         `json:"partOfSpeech"`
	Definitions  []*Definitions `json:"definitions"`
}

// DictResponse contains the words, as well as its phonetics and meanings described previously
type DictResponse struct {
	Word      string       `json:"word"`
	Phonetics []*Phonetics `json:"phonetics"`
	Meanings  []*Meanings  `json:"meanings"`
}
