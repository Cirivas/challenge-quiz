package ranking_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cirivas/challenge-quiz/core/use_cases/get_ranking"
	"github.com/gorilla/mux"
)

type rankingController struct {
	getRankingUseCase get_ranking.GetRankingUseCase
}

type RankingController interface {
	GetRanking(w http.ResponseWriter, r *http.Request)
}

func NewRankingController(getRankingUseCase get_ranking.GetRankingUseCase) RankingController {
	return &rankingController{getRankingUseCase}
}

func (rc *rankingController) GetRanking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	quizId, ok := vars["quizId"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]string{
			"error": "missing param",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	respondent, ok := vars["respondent"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]string{
			"error": "missing param",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	ranking, err := rc.getRankingUseCase.GetRanking(respondent, quizId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	message := map[string]string{
		"ranking": fmt.Sprintf("You were better than %v%% of all quizzer", ranking*100),
	}
	json.NewEncoder(w).Encode(message)
}
