package answer_quiz_test

import (
	"testing"

	"github.com/cirivas/challenge-quiz/core/entities"
	mock_repository "github.com/cirivas/challenge-quiz/core/repository/mock"
	"github.com/cirivas/challenge-quiz/core/use_cases/answer_quiz"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnswerQuiz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AnswerQuiz test suite")
}

var _ = Describe("Test AnswerQuiz", func() {
	questions := []entities.Question{
		{
			Text: "First Question",
			Alternatives: map[entities.AnswerKey]string{
				entities.First:  "First alternative",
				entities.Second: "Second Alternative",
				entities.Third:  "Third alternative",
				entities.Fourth: "Fourth alternative",
			},
			CorrectAnswer: entities.Third,
		},
		{
			Text: "Second Question",
			Alternatives: map[entities.AnswerKey]string{
				entities.First:  "First alternative",
				entities.Second: "Second Alternative",
				entities.Third:  "Third alternative",
				entities.Fourth: "Fourth alternative",
			},
			CorrectAnswer: entities.Second,
		},
		{
			Text: "Third Question",
			Alternatives: map[entities.AnswerKey]string{
				entities.First:  "First alternative",
				entities.Second: "Second Alternative",
				entities.Third:  "Third alternative",
				entities.Fourth: "Fourth alternative",
			},
			CorrectAnswer: entities.Second,
		},
	}

	quiz := &entities.Quiz{Questions: questions}

	When("There were no answers provided", func() {
		It("Should return error", func() {
			useCase := answer_quiz.NewAnswerQuizUseCase(nil)

			_, err := useCase.AnswerQuiz("", quiz, nil)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("no answers error"))
		})
	})

	When("There's a mismatch in answers", func() {
		It("Should return a mismatch error", func() {
			useCase := answer_quiz.NewAnswerQuizUseCase(nil)

			answers := []entities.AnswerKey{entities.First, entities.Fourth}

			_, err := useCase.AnswerQuiz("", quiz, answers)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("non matching answers to quiz"))
		})
	})

	When("Answers are provided correctly", func() {
		It("Should count correct answers", func() {
			testController := gomock.NewController(GinkgoT())
			scoreRepositoryMock := mock_repository.NewMockScoreRepository(testController)

			scoreRepositoryMock.EXPECT().SaveScore("respondent", 1).Return(nil).Times(1)

			useCase := answer_quiz.NewAnswerQuizUseCase(scoreRepositoryMock)

			answers := []entities.AnswerKey{entities.First, entities.Fourth, entities.Second}

			correctAnswers, err := useCase.AnswerQuiz("respondent", quiz, answers)

			Expect(err).To(BeNil())
			Expect(correctAnswers).To(Equal(1))
		})
	})
})
