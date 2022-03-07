package repositories

import (
	application "bareksa-interview-project/application"
	repositories "bareksa-interview-project/domain/repositories"
	persistence "bareksa-interview-project/infrastructure/persistence"
)

func CreateTopicsRepositoryResolve(dbPass string, redisPass string) repositories.ITopicsRepository {
	return createTopicsRepository(persistence.CreateDatabase(dbPass), persistence.CreateRedisCache(redisPass))
}

func CreateTopicsServiceResolve(dbPass string, redisPass string) application.TopicsService {
	return application.CreateTopicsService(CreateTopicsRepositoryResolve(dbPass, redisPass))
}
