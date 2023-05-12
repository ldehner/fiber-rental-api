package conf

import (
	"github.com/ldehner/fiber-rental-api/hasura"
	"github.com/ldehner/fiber-rental-api/repository"
)

type Conf struct {
}

// User Repositories
func (Conf) GetUserRepository() repository.UserRepository {
	return hasura.HasuraUserRepository{}
}

func (Conf) GetUserSearchProfileRepository() repository.UserSearchprofileRepository {
	return hasura.HasuraUserSearchProfileRepository{}
}

func (Conf) GetUserTenantInfoRepository() repository.UserTenantinfoRepository {
	return hasura.HasuraUserTenantinfoRepository{}
}

// Property Repositories
func (Conf) GetPropertyRepository() repository.PropertyRepository {
	return hasura.HasuraPropertyRepository{}
}

func (Conf) GetPropertyMarketProfileRepository() repository.PropertyMarketProfileRepository {
	return hasura.HasuraPropertyMarketProfileRepository{}
}

func (Conf) GetPropertyRentProfileRepository() repository.PropertyRentProfileRepository {
	return hasura.HasuraPropertyRentProfileRepository{}
}

func (Conf) GetPropertyIncidentRepository() repository.PropertyIncidentRepository {
	return hasura.HasuraPropertyIncidentRepository{}
}
