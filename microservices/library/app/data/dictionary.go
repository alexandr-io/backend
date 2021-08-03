package data

type Phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type Definitions struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Example    string   `json:"example"`
}

type Meanings struct {
	PartOfSpeech string         `json:"partOfSpeech"`
	Definitions  []*Definitions `json:"definitions"`
}

type DictResponse struct {
	Word      string       `json:"word"`
	Phonetics []*Phonetics `json:"phonetics"`
	Meanings  []*Meanings  `json:"meanings"`
}
