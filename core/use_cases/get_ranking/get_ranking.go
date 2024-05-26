package get_ranking

import (
	"slices"

	"github.com/cirivas/challenge-quiz/core/repository"
)

type getRanking struct {
	scoreRepository repository.ScoreRepository
}

type GetRankingUseCase interface {
	GetRanking(respondent string, quizId string) (float64, error)
}

func NewGetRankingUseCase(scoreRepository repository.ScoreRepository) GetRankingUseCase {
	return &getRanking{scoreRepository}
}

func (uc *getRanking) GetRanking(respondent string, quizId string) (float64, error) {
	score, err := uc.scoreRepository.GetScore(respondent, quizId)

	if err != nil {
		return 0.0, err
	}

	allScores, err := uc.scoreRepository.GetOthersScore(respondent, quizId)

	if err != nil {
		return 0.0, err
	}

	slices.Sort(allScores)

	firstScoreIndex := slices.Index(allScores, score)

	if firstScoreIndex == -1 {
		if len(allScores) > 0 {
			if score > allScores[len(allScores)-1] {
				// high score
				return 1.0, nil
			}
		}
		// no other quizzers
		return 0, nil
	}

	totalAnswers := len(allScores)

	ranking := float64(len(allScores[:firstScoreIndex])) / float64(totalAnswers)

	return ranking, nil
}
