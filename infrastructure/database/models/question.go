package models

import "github.com/cirivas/challenge-quiz/core/entities"

type Question struct {
	Id            string
	Text          string
	Alternatives  map[entities.AnswerKey]string
	CorrectAnswer entities.AnswerKey
}
