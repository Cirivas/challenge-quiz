package registry

import (
	"github.com/cirivas/challenge-quiz/core/repository"
	"github.com/cirivas/challenge-quiz/entrypoints/api/controllers"
	app_repository "github.com/cirivas/challenge-quiz/entrypoints/repository"
	gateway_repository "github.com/cirivas/challenge-quiz/gateways/repository"
	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
	redis_db "github.com/cirivas/challenge-quiz/infrastructure/database/redis"
	"github.com/redis/go-redis/v9"
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
		QuizController:    r.NewQuizController(),
		RankingController: r.NewRankingController(),
	}
}

// repositories
func (r *registry) NewQuizRepository() repository.QuizRepository {
	store := redis_db.NewRedisCollection[models.Quiz]("quiz", r.dbClient.Client().(*redis.Client))
	return gateway_repository.NewQuizRepositoryGateway(app_repository.NewQuizRepository(store))
}

func (r *registry) NewScoreRepository() repository.ScoreRepository {
	store := redis_db.NewRedisCollection[models.Response]("response", r.dbClient.Client().(*redis.Client))
	return gateway_repository.NewScoreRepositoryGateway(app_repository.NewScoreRepository(store))
}
