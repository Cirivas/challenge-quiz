package registry

import (
	"github.com/cirivas/challenge-quiz/core/use_cases/answer_quiz"
	"github.com/cirivas/challenge-quiz/core/use_cases/get_quiz"
	quiz_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/quiz"
)

func (r *registry) NewQuizController() quiz_controller.QuizController {
	return quiz_controller.NewQuizController(r.NewGetQuizUseCase(), r.NewAnswerQuizUseCase())
}

func (r *registry) NewAnswerQuizUseCase() answer_quiz.AnswerQuizUseCase {
	return answer_quiz.NewAnswerQuizUseCase(r.NewScoreRepository())
}

func (r *registry) NewGetQuizUseCase() get_quiz.GetQuizUseCase {
	return get_quiz.NewGetQuizUseCase(r.NewQuizRepository())
}
