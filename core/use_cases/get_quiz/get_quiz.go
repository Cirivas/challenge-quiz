package get_quiz

import (
	"github.com/cirivas/challenge-quiz/core/entities"
	"github.com/cirivas/challenge-quiz/core/repository"
)

type getQuiz struct {
	quizRepository repository.QuizRepository
}

type GetQuizUseCase interface {
	GetQuiz(quizId string) (*entities.Quiz, error)
}

func NewGetQuizUseCase(r repository.QuizRepository) GetQuizUseCase {
	return &getQuiz{r}
}

func (uc *getQuiz) GetQuiz(quizId string) (*entities.Quiz, error) {
	return uc.quizRepository.GetQuiz(quizId)
}
