package hasura

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraPropertyMarketProfileRepository struct {
}

func (HasuraPropertyMarketProfileRepository) CreateMarketProfile(marketprofile models.MarketProfile) (models.MarketProfile, error) {
	return models.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) UpdateMarketProfile(marketprofile models.MarketProfile) (models.MarketProfile, error) {
	return models.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) GetMarketProfile(id string) (models.MarketProfile, error) {
	return models.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) DeleteMarketProfile(id string) error {
	return nil
}
func (HasuraPropertyMarketProfileRepository) GetMarketProfiles() ([]models.MarketProfile, error) {
	return []models.MarketProfile{}, nil
}
