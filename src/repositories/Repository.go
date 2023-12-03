package repositories

import "dana/anabul-rest-api/src/database"

type Repositories interface {
	UserRepository
}

type RepositoriesImpl struct {
	*UserRepositoryImpl
}

func NewRepository(db *database.PostgresImpl) *RepositoriesImpl {
	return &RepositoriesImpl{
		UserRepositoryImpl: &UserRepositoryImpl{db: db.DB},
	}
}
