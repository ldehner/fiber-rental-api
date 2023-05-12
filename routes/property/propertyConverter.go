package propertyroutes

import (
	"github.com/ldehner/fiber-rental-api/models"
)

func CreateResponseProperty(property models.Property) Property {
	return Property{
		Id:          property.Id,
		Apartment:   property.Apartment,
		Balcony:     property.Balcony,
		City:        property.City,
		Country:     property.Country,
		Garage:      property.Garage,
		Garden:      property.Garden,
		HeatType:    property.HeatType,
		Housenumber: property.Housenumber,
		LandlordId:  property.LandlordId,
		Rooms:       property.Rooms,
		Size:        property.Size,
		Status:      property.Status,
		Street:      property.Street,
		TenantId:    property.TenantId,
		Type:        property.Type,
	}
}

func CreateResponseMarketProfile(marketProfile models.MarketProfile) MarketProfile {
	return MarketProfile{
		Rent:             marketProfile.Rent,
		PowerPrice:       marketProfile.PowerPrice,
		Period:           marketProfile.Period,
		MinimumPeriod:    marketProfile.MinimumPeriod,
		MinimumIncome:    marketProfile.MinimumIncome,
		Id:               marketProfile.Id,
		HeatPrice:        marketProfile.HeatPrice,
		Description:      marketProfile.Description,
		AvailabilityDate: marketProfile.AvailabilityDate,
		Deposit:          marketProfile.Deposit,
	}
}

func CreateResponseRentProfile(rentProfile models.RentProfile) RentProfile {
	return RentProfile{
		ContractDue: rentProfile.ContractDue,
		HeatCosts:   rentProfile.HeatCosts,
		Deposit:     rentProfile.Deposit,
		Id:          rentProfile.Id,
		Minimum:     rentProfile.Minimum,
		Rent:        rentProfile.Rent,
		RentalStart: rentProfile.RentalStart,
	}
}

func CreateResponseInvite(invite models.Invite) Invite {
	return Invite{
		Id:       invite.Id,
		Property: invite.Property,
		Tenant:   invite.Tenant,
		Landlord: invite.Landlord,
		ValidDue: invite.ValidDue,
	}
}
func CreateResponseIncident(incident models.Incident) Incident {
	return Incident{
		Description: incident.Description,
		Id:          incident.Id,
		Type:        incident.Type,
		TenantId:    incident.TenantId,
		Status:      incident.Status,
		PropertyId:  incident.PropertyId,
		LandlordId:  incident.LandlordId,
	}

}
