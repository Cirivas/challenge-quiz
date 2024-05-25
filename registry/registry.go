package registry

import (
	"github.com/cirivas/challenge-quiz/entrypoints/api/controllers"
	"github.com/cirivas/challenge-quiz/infrastructure/database"
)

type registry struct {
	dbClient database.DatastoreClient
}

type Registry interface {
	NewController(dbClient database.DatastoreClient) controllers.Controllers
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewController(dbClient database.DatastoreClient) controllers.Controllers {
	r.dbClient = dbClient
	return controllers.Controllers{
		QuizController: r.NewQuizController(),
	}
}
