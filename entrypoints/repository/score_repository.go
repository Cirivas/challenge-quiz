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
	responses, _ := repo.db.Get()

	for _, v := range responses {
		if v.Respondent == respondent && v.QuizId == quizId {
			return v.Score, nil
		}
	}
	return 0, nil
}

func (repo *scoreRepository) GetOthersScore(respondent string, quizId string) ([]int, error) {
	responses, _ := repo.db.Get()

	result := make([]int, 0)
	for _, v := range responses {
		if v.Respondent == respondent && v.QuizId == quizId {
			continue
		}
		if v.QuizId == quizId {
			result = append(result, v.Score)
		}
	}

	return result, nil
}
