package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraPropertyMarketProfileRepository struct {
}

func (HasuraPropertyMarketProfileRepository) CreateMarketProfile(marketprofile storemodels.MarketProfile) (storemodels.MarketProfile, error) {
	return storemodels.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) UpdateMarketProfile(marketprofile storemodels.MarketProfile) (storemodels.MarketProfile, error) {
	return storemodels.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) GetMarketProfile(id string) (storemodels.MarketProfile, error) {
	return storemodels.MarketProfile{}, nil
}
func (HasuraPropertyMarketProfileRepository) DeleteMarketProfile(id string) error {
	return nil
}
func (HasuraPropertyMarketProfileRepository) GetMarketProfiles() ([]storemodels.MarketProfile, error) {
	return []storemodels.MarketProfile{}, nil
}
