package answer_quiz

import (
	"errors"

	"github.com/cirivas/challenge-quiz/core/entities"
	"github.com/cirivas/challenge-quiz/core/repository"
)

type answerQuiz struct {
	scoreRepository repository.ScoreRepository
}

type AnswerQuizUseCase interface {
	AnswerQuiz(respondent string, quiz *entities.Quiz, answers []entities.AnswerKey) (int, error)
}

func NewAnswerQuizUseCase(scoreRepository repository.ScoreRepository) AnswerQuizUseCase {
	return &answerQuiz{scoreRepository}
}

func (uc *answerQuiz) AnswerQuiz(respondent string, quiz *entities.Quiz, answers []entities.AnswerKey) (int, error) {
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

	if err := uc.scoreRepository.SaveScore(respondent, totalCorrectAnswers); err != nil {
		return 0, err
	}

	return totalCorrectAnswers, nil
}
