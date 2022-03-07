package repositories

import (
	application "bareksa-interview-project/application"
	repositories "bareksa-interview-project/domain/repositories"
	persistence "bareksa-interview-project/infrastructure/persistence"
)

func CreateNewsRepositoryResolve(dbPass string, redisPass string) repositories.INewsRepository {
	return createNewsRepository(persistence.CreateDatabase(dbPass), persistence.CreateRedisCache(redisPass))
}

func CreateNewsServiceResolve(dbPass string, redisPass string) application.NewsService {
	return application.CreateNewsService(CreateNewsRepositoryResolve(dbPass, redisPass))
}
