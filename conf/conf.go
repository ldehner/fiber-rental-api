package conf

import "github.com/ldehner/fiber-rental-api/repository"

type Conf struct {
}

func (Conf) GetUserRepository() repository.UserRepository {
	return repository.HasuraUserRepository{}
}
