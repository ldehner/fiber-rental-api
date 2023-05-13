package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type PropertyMarketProfileRepository interface {
	CreateMarketProfile(marketprofile storemodels.MarketProfile) (storemodels.MarketProfile, error)
	UpdateMarketProfile(marketprofile storemodels.MarketProfile) (storemodels.MarketProfile, error)
	GetMarketProfile(id string) (storemodels.MarketProfile, error)
	DeleteMarketProfile(id string) error
	GetMarketProfiles() ([]storemodels.MarketProfile, error)
}
