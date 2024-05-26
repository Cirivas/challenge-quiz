package repository

import (
	core_repository "github.com/cirivas/challenge-quiz/core/repository"
	app_repository "github.com/cirivas/challenge-quiz/entrypoints/repository"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
)

type scoreRepositoryGateway struct {
	scoreRepository app_repository.ScoreRepository
}

func NewScoreRepositoryGateway(scoreRepository app_repository.ScoreRepository) core_repository.ScoreRepository {
	return &scoreRepositoryGateway{scoreRepository}
}

func (srg *scoreRepositoryGateway) SaveScore(respondent string, score int, quizId string) error {
	response := models.Response{
		Respondent: respondent,
		Score:      score,
		QuizId:     quizId,
	}

	return srg.scoreRepository.SaveScore(response)
}

func (srg *scoreRepositoryGateway) GetScore(respondent string, quizId string) (int, error) {
	return srg.scoreRepository.GetScore(respondent, quizId)
}

func (srg *scoreRepositoryGateway) GetOthersScore(respondent string, quizId string) ([]int, error) {
	return srg.scoreRepository.GetOthersScore(respondent, quizId)
}
