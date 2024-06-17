package dbmodel

type PosTag struct {
	Word  string `bson:"word"`
	Value string `bson:"value"`
}

type Sentiment struct {
	Polarity     float64 `bson:"polarity"`
	Subjectivity float64 `bson:"subjectivity"`
}

type Dependency struct {
	Word   string `bson:"word"`
	DepRel string `bson:"depRel"`
	Head   string `bson:"head"`
}

type Verb struct {
	Base                string `bson:"base"`
	Past                string `bson:"past"`
	PastParticiple      string `bson:"pastParticiple"`
	Gerund              string `bson:"gerund"`
	ThirdPersonSingular string `bson:"thirdPersonSingular"`
}
