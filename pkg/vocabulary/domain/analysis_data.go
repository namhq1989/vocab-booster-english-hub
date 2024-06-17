package domain

type PosTag struct {
	Word  string
	Value PartOfSpeech
}

type Sentiment struct {
	Polarity     float64
	Subjectivity float64
}

type Dependency struct {
	Word   string
	DepRel string
	Head   string
}

type Verb struct {
	Base                string
	Past                string
	PastParticiple      string
	Gerund              string
	ThirdPersonSingular string
}
