package dbmodel

type PosTag struct {
	Word  string `json:"word"`
	Value string `json:"value"`
}

type Sentiment struct {
	Polarity     float64 `json:"polarity"`
	Subjectivity float64 `json:"subjectivity"`
}

type Dependency struct {
	Word   string `json:"word"`
	DepRel string `json:"depRel"`
	Head   string `json:"head"`
}

type Verb struct {
	Base                string `json:"base"`
	Present             string `json:"present"`
	Past                string `json:"past"`
	PastParticiple      string `json:"pastParticiple"`
	Gerund              string `json:"gerund"`
	ThirdPersonSingular string `json:"thirdPersonSingular"`
}
