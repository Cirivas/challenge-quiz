package models

import "github.com/cirivas/challenge-quiz/core/entities"

type Question struct {
	Text          string                        `json:"text"`
	Alternatives  map[entities.AnswerKey]string `json:"alternatives"`
	CorrectAnswer entities.AnswerKey            `json:"correctAnswer"`
}
