package ranking_controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_get_ranking "github.com/cirivas/challenge-quiz/core/use_cases/get_ranking/mock"
	ranking_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/ranking"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRankingController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QuizController test suite")
}

var _ = Describe("Test Ranking Controller", func() {
	It("should deliver a fixed text", func() {
		testController := gomock.NewController(GinkgoT())

		getRankingUseCaseMock := mock_get_ranking.NewMockGetRankingUseCase(testController)
		getRankingUseCaseMock.EXPECT().GetRanking(gomock.Any(), gomock.Any()).Return(0.45, nil).Times(1)

		rankingController := ranking_controller.NewRankingController(getRankingUseCaseMock)
		handler := http.HandlerFunc(rankingController.GetRanking)

		req, _ := http.NewRequest("GET", "/quiz/quizid/ranking/myself", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/quiz/{quizId}/ranking/{respondent}", handler)

		router.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(ContainSubstring("45%"))
	})
})
