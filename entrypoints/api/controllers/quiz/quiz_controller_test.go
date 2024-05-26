package quiz_controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cirivas/challenge-quiz/core/entities"
	mock_answer_quiz "github.com/cirivas/challenge-quiz/core/use_cases/answer_quiz/mock"
	mock_get_quiz "github.com/cirivas/challenge-quiz/core/use_cases/get_quiz/mock"

	quiz_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/quiz"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQuizController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QuizController test suite")
}

// ctrl := gomock.NewController(GinkgoT())
var _ = Describe("QuizController Test", func() {
	Context("Answer Quiz", func() {
		When("A request is received", func() {
			When("it's not a valid body", func() {
				It("should return internal error with message", func() {
					req, _ := http.NewRequest("POST", "/answer", strings.NewReader("some body"))
					rr := httptest.NewRecorder()

					quizController := quiz_controller.NewQuizController(nil, nil)
					handler := http.HandlerFunc(quizController.AnswerQuiz)

					handler.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusBadRequest))
					Expect(rr.Body.String()).ToNot(BeEmpty())
				})
			})

			When("getting the quiz", func() {
				It("should return internal server error if it fails", func() {
					req, _ := http.NewRequest("POST", "/answer", strings.NewReader(`{"respondent": "myself","quizId": "quizId", "answers": ["1","2","3"]}`))
					rr := httptest.NewRecorder()

					testCtrl := gomock.NewController(GinkgoT())

					getQuizMock := mock_get_quiz.NewMockGetQuizUseCase(testCtrl)

					getQuizMock.EXPECT().GetQuiz(gomock.Any()).Return(nil, errors.New("mock error")).Times(1)

					quizController := quiz_controller.NewQuizController(getQuizMock, nil)
					handler := http.HandlerFunc(quizController.AnswerQuiz)

					handler.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusInternalServerError))
					Expect(rr.Body.String()).To(ContainSubstring("internal error"))
				})

				It("should return not found error for invalid quiz id", func() {
					req, _ := http.NewRequest("POST", "/answer", strings.NewReader(`{"respondent": "myself","quizId": "quizId", "answers": ["1","2","3"]}`))
					rr := httptest.NewRecorder()

					testCtrl := gomock.NewController(GinkgoT())

					answerQuizkMock := mock_answer_quiz.NewMockAnswerQuizUseCase(testCtrl)
					getQuizMock := mock_get_quiz.NewMockGetQuizUseCase(testCtrl)

					getQuizMock.EXPECT().GetQuiz(gomock.Any()).Return(nil, nil).Times(1)

					quizController := quiz_controller.NewQuizController(getQuizMock, answerQuizkMock)
					handler := http.HandlerFunc(quizController.AnswerQuiz)

					handler.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusNotFound))
					Expect(rr.Body.String()).To(ContainSubstring("quiz not found"))
				})
			})

			When("answering the quest", func() {
				It("should return internal server error when it fails", func() {
					req, _ := http.NewRequest("POST", "/answer", strings.NewReader(`{"respondent": "myself","quizId": "quizId", "answers": ["1","2","3"]}`))
					rr := httptest.NewRecorder()

					testCtrl := gomock.NewController(GinkgoT())

					answerQuizkMock := mock_answer_quiz.NewMockAnswerQuizUseCase(testCtrl)
					getQuizMock := mock_get_quiz.NewMockGetQuizUseCase(testCtrl)

					getQuizMock.EXPECT().GetQuiz(gomock.Any()).Return(&entities.Quiz{}, nil).Times(1)
					answerQuizkMock.EXPECT().AnswerQuiz(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(0, errors.New("mock error")).Times(1)

					quizController := quiz_controller.NewQuizController(getQuizMock, answerQuizkMock)
					handler := http.HandlerFunc(quizController.AnswerQuiz)

					handler.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusInternalServerError))
					Expect(rr.Body.String()).To(ContainSubstring("internal error"))
				})

				It("should return the score of the quiz", func() {
					req, _ := http.NewRequest("POST", "/answer", strings.NewReader(`{"respondent": "myself","quizId": "quizId", "answers": ["1","2","3"]}`))
					rr := httptest.NewRecorder()

					testCtrl := gomock.NewController(GinkgoT())

					answerQuizkMock := mock_answer_quiz.NewMockAnswerQuizUseCase(testCtrl)
					getQuizMock := mock_get_quiz.NewMockGetQuizUseCase(testCtrl)

					getQuizMock.EXPECT().GetQuiz(gomock.Any()).Return(&entities.Quiz{}, nil).Times(1)
					answerQuizkMock.EXPECT().AnswerQuiz(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(2, nil).Times(1)

					quizController := quiz_controller.NewQuizController(getQuizMock, answerQuizkMock)
					handler := http.HandlerFunc(quizController.AnswerQuiz)

					handler.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusOK))
					Expect(rr.Body.String()).To(ContainSubstring("2/3"))
				})
			})
		})
	})
})
