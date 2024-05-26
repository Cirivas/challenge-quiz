package registry

import (
	"github.com/cirivas/challenge-quiz/core/repository"
	"github.com/cirivas/challenge-quiz/core/use_cases/answer_quiz"
	"github.com/cirivas/challenge-quiz/core/use_cases/get_quiz"
	quiz_controller "github.com/cirivas/challenge-quiz/entrypoints/api/controllers/quiz"
	app_repository "github.com/cirivas/challenge-quiz/entrypoints/repository"
	gateway_repository "github.com/cirivas/challenge-quiz/gateways/repository"
	"github.com/cirivas/challenge-quiz/infrastructure/database/models"
	redis_db "github.com/cirivas/challenge-quiz/infrastructure/database/redis"
	"github.com/redis/go-redis/v9"
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

func (r *registry) NewQuizRepository() repository.QuizRepository {
	store := redis_db.NewRedisCollection[models.Quiz]("quiz", r.dbClient.Client().(*redis.Client))
	return gateway_repository.NewQuizRepositoryGateway(app_repository.NewQuizRepository(store))
}

func (r *registry) NewScoreRepository() repository.ScoreRepository {
	store := redis_db.NewRedisCollection[models.Response]("response", r.dbClient.Client().(*redis.Client))
	return gateway_repository.NewScoreRepositoryGateway(app_repository.NewScoreRepository(store))
}
