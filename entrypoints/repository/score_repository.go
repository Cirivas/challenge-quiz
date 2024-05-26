package repository

import (
	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
)

type scoreRepository struct {
	db database.Datastore[models.Response]
}

type ScoreRepository interface {
	SaveScore(models.Response) error
	GetScore(respondent string, quizId string) (int, error)
	GetOthersScore(respondent string, quizId string) ([]int, error)
}

func NewScoreRepository(db database.Datastore[models.Response]) ScoreRepository {
	return &scoreRepository{db}
}

func (repo *scoreRepository) SaveScore(response models.Response) error {
	return repo.db.Save(response)
}

func (repo *scoreRepository) GetScore(respondent string, quizId string) (int, error) {
	return 0, nil
}

func (repo *scoreRepository) GetOthersScore(respondent string, quizId string) ([]int, error) {
	return nil, nil
}
