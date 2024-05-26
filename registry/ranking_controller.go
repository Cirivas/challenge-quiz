package registry

import (
	"github.com/cirivas/challenge-quiz/core/use_cases/get_ranking"
	ranking_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/ranking"
)

func (r *registry) NewRankingController() ranking_controller.RankingController {
	return ranking_controller.NewRankingController(r.NewGetRankingUseCase())
}

func (r *registry) NewGetRankingUseCase() get_ranking.GetRankingUseCase {
	return get_ranking.NewGetRankingUseCase(r.NewScoreRepository())
}
