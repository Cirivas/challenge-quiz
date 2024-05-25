package registry

import (
	"github.com/cirivas/challenge-quiz/core/repository"
	"github.com/cirivas/challenge-quiz/core/use_cases/get_quiz"
	"github.com/cirivas/challenge-quiz/entrypoints/api/controllers"
)

func (r *registry) NewQuizController() controllers.QuizController {
	return controllers.NewQuizController(r.NewGetQuizUseCase())
}

func (r *registry) NewGetQuizUseCase() get_quiz.GetQuizUseCase {
	return get_quiz.NewGetQuizUseCase(r.NewQuizRepository())
}

func (r *registry) NewQuizRepository() repository.QuizRepository {
	return nil
}
