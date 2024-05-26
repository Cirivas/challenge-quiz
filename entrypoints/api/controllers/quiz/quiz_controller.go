package quiz_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cirivas/challenge-quiz/core/entities"
	"github.com/cirivas/challenge-quiz/core/use_cases/answer_quiz"
	"github.com/cirivas/challenge-quiz/core/use_cases/get_quiz"
	"github.com/gorilla/mux"
)

type quizController struct {
	getQuizUseCase    get_quiz.GetQuizUseCase
	answerQuizUseCase answer_quiz.AnswerQuizUseCase
}

type QuizController interface {
	GetQuiz(w http.ResponseWriter, r *http.Request)
	AnswerQuiz(w http.ResponseWriter, r *http.Request)
}

func NewQuizController(
	getQuizUseCase get_quiz.GetQuizUseCase,
	answerQuizUseCase answer_quiz.AnswerQuizUseCase,
) QuizController {
	return &quizController{
		getQuizUseCase,
		answerQuizUseCase,
	}
}

func (qc *quizController) GetQuiz(w http.ResponseWriter, r *http.Request) {
	log.Println("QuizController.GetQuiz")

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

	quiz, err := qc.getQuizUseCase.GetQuiz(quizId)

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

	quizDTO := QuizEntityToDTO(quizId, quiz)
	quizJSON, _ := json.Marshal(quizDTO)

	w.WriteHeader(http.StatusOK)
	w.Write(quizJSON)
}

func (qc *quizController) AnswerQuiz(w http.ResponseWriter, r *http.Request) {
	bodyRaw, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	var request struct {
		QuizId     string               `json:"quizId"`
		Respondent string               `json:"respondent"`
		Answers    []entities.AnswerKey `json:"answers"`
	}

	if err = json.Unmarshal(bodyRaw, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	quiz, err := qc.getQuizUseCase.GetQuiz(request.QuizId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := map[string]string{
			"error": "internal error",
		}
		json.NewEncoder(w).Encode(message)
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

	correctAnswers, err := qc.answerQuizUseCase.AnswerQuiz(request.QuizId, request.Respondent, quiz, request.Answers)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := map[string]string{
			"error": "internal error",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	w.WriteHeader(http.StatusOK)
	message := map[string]string{
		"score": fmt.Sprintf("%d/%d", correctAnswers, len(request.Answers)),
	}
	json.NewEncoder(w).Encode(message)

}
