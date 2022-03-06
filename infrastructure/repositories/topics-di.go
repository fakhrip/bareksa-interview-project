package repositories

import (
	application "bareksa-interview-project/application"
	repositories "bareksa-interview-project/domain/repositories"
	persistence "bareksa-interview-project/infrastructure/persistence"
)

func CreateTopicsRepositoryResolve(dbPass string) repositories.ITopicsRepository {
	return createTopicsRepository(persistence.CreateDatabase(dbPass))
}

func CreateTopicsServiceResolve(dbPass string) application.TopicsService {
	return application.CreateTopicsService(CreateTopicsRepositoryResolve(dbPass))
}
