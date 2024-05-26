package get_quiz

import (
	"github.com/cirivas/challenge-quiz/core/repository"
)

type saveScore struct {
	scoreRepository repository.ScoreRepository
}

type SaveScoreUseCase interface {
	SaveScore(respondent string, score int) error
}

func NewSaveScoreUseCase(r repository.ScoreRepository) SaveScoreUseCase {
	return &saveScore{r}
}

func (uc *saveScore) SaveScore(respondent string, score int) error {
	return uc.scoreRepository.SaveScore(respondent, score)
}
