package controllers

import (
	quiz_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/quiz"
	ranking_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/ranking"
)

type Controllers struct {
	QuizController    quiz_controller.QuizController
	RankingController ranking_controller.RankingController
}
