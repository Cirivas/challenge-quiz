package repository

import (
	"fmt"

	"github.com/cirivas/challenge-quiz/core/entities"
	core_repository "github.com/cirivas/challenge-quiz/core/repository"
	app_repository "github.com/cirivas/challenge-quiz/entrypoints/repository"
	"github.com/cirivas/challenge-quiz/gateways"
)

type quizRepositoryGateway struct {
	quizRepository app_repository.QuizRepository
}

func NewQuizRepositoryGateway(quizRepository app_repository.QuizRepository) core_repository.QuizRepository {
	return &quizRepositoryGateway{quizRepository}
}

func (qrg *quizRepositoryGateway) GetQuiz(id string) (*entities.Quiz, error) {
	quizModel, err := qrg.quizRepository.GetQuiz(id)

	if err != nil {
		fmt.Printf("appRepo error: %#v\n", err)
		return nil, err
	}

	quizEntity := gateways.QuizModelToEntity(quizModel)
	return quizEntity, nil
}
