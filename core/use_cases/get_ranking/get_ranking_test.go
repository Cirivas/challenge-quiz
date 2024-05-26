package get_ranking_test

import (
	"errors"
	"testing"

	mock_repository "github.com/cirivas/challenge-quiz/core/repository/mock"
	"github.com/cirivas/challenge-quiz/core/use_cases/get_ranking"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGetRanking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AnswerQuiz test suite")
}

var _ = Describe("Test Get Ranking", func() {
	When("a repository fails", func() {
		It("should return error if GetScore fails", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			scoreRepositoryMock.EXPECT().GetScore(gomock.Any(), gomock.Any()).Return(0, errors.New("mock error")).Times(1)

			useCase := get_ranking.NewGetRankingUseCase(scoreRepositoryMock)

			_, err := useCase.GetRanking("respondent", "quizId")

			Expect(err).To(Not(BeNil()))
		})

		It("should return error if GetOtherScore fails", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			scoreRepositoryMock.EXPECT().GetScore(gomock.Any(), gomock.Any()).Return(1, nil).Times(1)
			scoreRepositoryMock.EXPECT().GetOthersScore(gomock.Any(), gomock.All()).Return(nil, errors.New("mock error")).Times(1)

			useCase := get_ranking.NewGetRankingUseCase(scoreRepositoryMock)

			_, err := useCase.GetRanking("respondent", "quizId")

			Expect(err).To(Not(BeNil()))
		})
	})

	When("everything is fine", func() {
		It("should return better than 30 percent (0.3) position", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			othersScores := []int{1, 3, 1, 2, 3, 2, 2, 3, 1, 2}

			scoreRepositoryMock.EXPECT().GetScore(gomock.Any(), gomock.Any()).Return(2, nil).Times(1)
			scoreRepositoryMock.EXPECT().GetOthersScore(gomock.Any(), gomock.All()).Return(othersScores, nil).Times(1)

			useCase := get_ranking.NewGetRankingUseCase(scoreRepositoryMock)

			ranking, err := useCase.GetRanking("respondent", "quizId")

			Expect(err).To((BeNil()))
			Expect(ranking).To(Equal(0.3))
		})

		It("should return better than 70 percent (0.7) position", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			othersScores := []int{1, 3, 1, 2, 3, 2, 2, 3, 1, 2}

			scoreRepositoryMock.EXPECT().GetScore(gomock.Any(), gomock.Any()).Return(3, nil).Times(1)
			scoreRepositoryMock.EXPECT().GetOthersScore(gomock.Any(), gomock.All()).Return(othersScores, nil).Times(1)

			useCase := get_ranking.NewGetRankingUseCase(scoreRepositoryMock)

			ranking, err := useCase.GetRanking("respondent", "quizId")

			Expect(err).To((BeNil()))
			Expect(ranking).To(Equal(0.7))
		})

		It("sohuld return better than everyone (1) position", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			othersScores := []int{1, 3, 1, 2, 3, 2, 2, 3, 1, 2}

			scoreRepositoryMock.EXPECT().GetScore(gomock.Any(), gomock.Any()).Return(4, nil).Times(1)
			scoreRepositoryMock.EXPECT().GetOthersScore(gomock.Any(), gomock.All()).Return(othersScores, nil).Times(1)

			useCase := get_ranking.NewGetRankingUseCase(scoreRepositoryMock)

			ranking, err := useCase.GetRanking("respondent", "quizId")

			Expect(err).To((BeNil()))
			Expect(ranking).To(Equal(1.0))
		})
	})
})
