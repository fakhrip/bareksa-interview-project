package repositories

import (
	application "bareksa-interview-project/application"
	repositories "bareksa-interview-project/domain/repositories"
	persistence "bareksa-interview-project/infrastructure/persistence"
)

func CreateNewsRepositoryResolve(dbPass string) repositories.INewsRepository {
	return createNewsRepository(persistence.CreateDatabase(dbPass))
}

func CreateNewsServiceResolve(dbPass string) application.NewsService {
	return application.CreateNewsService(CreateNewsRepositoryResolve(dbPass))
}
