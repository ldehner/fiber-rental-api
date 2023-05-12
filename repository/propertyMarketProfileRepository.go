package repository

import "github.com/ldehner/fiber-rental-api/models"

type PropertyMarketProfileRepository interface {
	CreateMarketProfile(marketprofile models.MarketProfile) (models.MarketProfile, error)
	UpdateMarketProfile(marketprofile models.MarketProfile) (models.MarketProfile, error)
	GetMarketProfile(id string) (models.MarketProfile, error)
	DeleteMarketProfile(id string) error
	GetMarketProfiles() ([]models.MarketProfile, error)
}
