package repository

import "github.com/cirivas/challenge-quiz/core/entities"

type QuizRepository interface {
	GetQuiz(id string) (*entities.Quiz, error)
}
