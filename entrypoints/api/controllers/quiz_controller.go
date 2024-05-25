package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cirivas/challenge-quiz/core/use_cases/get_quiz"
	"github.com/cirivas/challenge-quiz/gateways"
)

type quizController struct {
	getQuizUseCase get_quiz.GetQuizUseCase
}

type QuizController interface {
	GetQuiz(w http.ResponseWriter, r *http.Request)
}

func NewQuizController(getQuizUseCase get_quiz.GetQuizUseCase) QuizController {
	return &quizController{getQuizUseCase}
}

func (qc *quizController) GetQuiz(w http.ResponseWriter, r *http.Request) {
	quiz, err := qc.getQuizUseCase.GetQuiz("quizid")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if quiz == nil {
		w.WriteHeader(http.StatusNotFound)
		message := map[string]string{
			"error": "quiz not found",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	quizModel := gateways.QuizEntityToModel("quizId", quiz)
	quizJSON, _ := json.Marshal(quizModel)

	w.WriteHeader(http.StatusOK)
	w.Write(quizJSON)

}
