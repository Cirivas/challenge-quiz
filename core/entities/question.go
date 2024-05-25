package entities

type Question struct {
	Text          string
	Alternatives  map[AnswerKey]string
	CorrectAnswer AnswerKey
}
