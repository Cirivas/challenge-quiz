package gateways

import (
	"github.com/cirivas/challenge-quiz/core/entities"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
)

func QuestionModelToEntity(question models.Question) *entities.Question {
	return &entities.Question{
		Text:          question.Text,
		CorrectAnswer: question.CorrectAnswer,
		Alternatives:  question.Alternatives,
	}
}

func QuizModelToEntity(quiz models.Quiz) *entities.Quiz {
	questions := make([]entities.Question, len(quiz.Questions))

	for _, q := range quiz.Questions {
		questions = append(questions, *QuestionModelToEntity(q))
	}

	return &entities.Quiz{
		Questions: questions,
	}
}
