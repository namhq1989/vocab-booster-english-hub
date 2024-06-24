package domain

type VerbForm string

const (
	VerbFormUnknown             VerbForm = ""
	VerbFormInfinitive          VerbForm = "infinitive"
	VerbFormPast                VerbForm = "past"
	VerbFormPastParticiple      VerbForm = "past_participle"
	VerbFormGerund              VerbForm = "gerund"
	VerbFormThirdPersonSingular VerbForm = "third_person_singular"
)

func (s VerbForm) String() string {
	switch s {
	case VerbFormInfinitive, VerbFormPast, VerbFormPastParticiple, VerbFormGerund, VerbFormThirdPersonSingular:
		return string(s)
	default:
		return ""
	}
}

func (s VerbForm) IsValid() bool {
	return s != VerbFormUnknown
}

func ToVerbForm(value string) VerbForm {
	switch value {
	case VerbFormInfinitive.String():
		return VerbFormInfinitive
	case VerbFormPast.String():
		return VerbFormPast
	case VerbFormPastParticiple.String():
		return VerbFormPastParticiple
	case VerbFormGerund.String():
		return VerbFormGerund
	case VerbFormThirdPersonSingular.String():
		return VerbFormThirdPersonSingular
	default:
		return VerbFormUnknown
	}
}
