package registry

import "github.com/cirivas/challenge-quiz/entrypoints/api/controllers"

type registry struct{}

type Registry interface {
	NewController() controllers.Controllers
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewController() controllers.Controllers {
	return controllers.Controllers{
		QuizController: r.NewQuizController(),
	}
}
