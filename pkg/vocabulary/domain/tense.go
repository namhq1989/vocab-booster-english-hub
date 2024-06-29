package domain

type Tense string

const (
	TenseUnknown                  Tense = ""
	TensePresentSimple            Tense = "present_simple"
	TensePresentPerfect           Tense = "present_perfect"
	TensePresentContinuous        Tense = "present_continuous"
	TensePresentPerfectContinuous Tense = "present_perfect_continuous"
	TensePastSimple               Tense = "past_simple"
	TensePastPerfect              Tense = "past_perfect"
	TensePastContinuous           Tense = "past_continuous"
	TensePastPerfectContinuous    Tense = "past_perfect_continuous"
	TenseFutureSimple             Tense = "future_simple"
	TenseFuturePerfect            Tense = "future_perfect"
	TenseFutureContinuous         Tense = "future_continuous"
	TenseFuturePerfectContinuous  Tense = "future_perfect_continuous"
)

func (t Tense) IsValid() bool {
	return t != TenseUnknown
}

func (t Tense) String() string {
	return string(t)
}

func ToTense(value string) Tense {
	switch value {
	case TensePresentSimple.String():
		return TensePresentSimple
	case TensePresentPerfect.String():
		return TensePresentPerfect
	case TensePresentContinuous.String():
		return TensePresentContinuous
	case TensePresentPerfectContinuous.String():
		return TensePresentPerfectContinuous
	case TensePastSimple.String():
		return TensePastSimple
	case TensePastPerfect.String():
		return TensePastPerfect
	case TensePastContinuous.String():
		return TensePastContinuous
	case TensePastPerfectContinuous.String():
		return TensePastPerfectContinuous
	case TenseFutureSimple.String():
		return TenseFutureSimple
	case TenseFuturePerfect.String():
		return TenseFuturePerfect
	case TenseFutureContinuous.String():
		return TenseFutureContinuous
	case TenseFuturePerfectContinuous.String():
		return TenseFuturePerfectContinuous
	default:
		return TenseUnknown
	}
}
