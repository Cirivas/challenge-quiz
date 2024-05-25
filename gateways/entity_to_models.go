package gateways

import (
	"github.com/cirivas/challenge-quiz/core/entities"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
)

func QuestionEntityToModel(question entities.Question) *models.Question {
	return &models.Question{
		Text:          question.Text,
		Alternatives:  question.Alternatives,
		CorrectAnswer: question.CorrectAnswer,
	}
}

func QuizEntityToModel(quizId string, quiz entities.Quiz) *models.Quiz {
	questions := make([]models.Question, len(quiz.Questions))

	for _, q := range quiz.Questions {
		questions = append(questions, *QuestionEntityToModel(q))
	}

	return &models.Quiz{
		Id:        quizId,
		Questions: questions,
	}
}
