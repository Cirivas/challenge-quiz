package repository

import (
	"fmt"

	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
)

type quizRepository struct {
	db database.Datastore[models.Quiz]
}

type QuizRepository interface {
	GetQuiz(id string) (*models.Quiz, error)
}

func NewQuizRepository(db database.Datastore[models.Quiz]) QuizRepository {
	return &quizRepository{db}
}

func (repo *quizRepository) GetQuiz(id string) (*models.Quiz, error) {
	quiz, err := repo.db.GetById(id)

	if err != nil {
		fmt.Printf("DB error: %#v\n", err)
		return nil, err
	}

	return quiz, nil
}
