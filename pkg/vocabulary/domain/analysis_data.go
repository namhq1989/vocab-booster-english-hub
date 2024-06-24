package domain

type PosTag struct {
	Word  string
	Value PartOfSpeech
	Level int
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
