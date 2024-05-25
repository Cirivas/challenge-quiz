package answer_quiz

import (
	"errors"

	"github.com/cirivas/challenge-quiz/core/entities"
)

type answerQuiz struct{}

type AnswerQuizUseCase interface {
	AnswerQuiz(quiz entities.Quiz, answers []entities.AnswerKey) (int, error)
}

func NewAnswerQuizUseCase() AnswerQuizUseCase {
	return &answerQuiz{}
}

func (uc *answerQuiz) AnswerQuiz(quiz entities.Quiz, answers []entities.AnswerKey) (int, error) {
	if len(answers) == 0 {
		return 0, errors.New("no answers error")
	}

	if len(answers) != len(quiz.Questions) {
		return 0, errors.New("non matching answers to quiz")
	}

	totalCorrectAnswers := 0

	for i, question := range quiz.Questions {
		if question.CorrectAnswer == answers[i] {
			totalCorrectAnswers++
		}
	}

	return totalCorrectAnswers, nil
}
