package propertyroutes

import (
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

func CreateResponseProperty(property storemodels.Property) responsemodels.Property {
	return responsemodels.Property{
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

func CreateResponseMarketProfile(marketProfile storemodels.MarketProfile) responsemodels.MarketProfile {
	return responsemodels.MarketProfile{
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

func CreateResponseRentProfile(rentProfile storemodels.RentProfile) responsemodels.RentProfile {
	return responsemodels.RentProfile{
		ContractDue: rentProfile.ContractDue,
		HeatCosts:   rentProfile.HeatCosts,
		Deposit:     rentProfile.Deposit,
		Id:          rentProfile.Id,
		Minimum:     rentProfile.Minimum,
		Rent:        rentProfile.Rent,
		RentalStart: rentProfile.RentalStart,
	}
}

func CreateResponseInvite(invite storemodels.Invite) responsemodels.Invite {
	return responsemodels.Invite{
		Id:       invite.Id,
		Property: invite.Property,
		Tenant:   invite.Tenant,
		Landlord: invite.Landlord,
		ValidDue: invite.ValidDue,
	}
}
func CreateResponseIncident(incident storemodels.Incident) responsemodels.Incident {
	return responsemodels.Incident{
		Description: incident.Description,
		Id:          incident.Id,
		Type:        incident.Type,
		TenantId:    incident.TenantId,
		Status:      incident.Status,
		PropertyId:  incident.PropertyId,
		LandlordId:  incident.LandlordId,
	}

}

func CreateStoreProperty(property requestmodels.Property, id string) storemodels.Property {
	return storemodels.Property{
		Id:          id,
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

func CreateStoreMarketProfile(marketProfile requestmodels.MarketProfile, id string) storemodels.MarketProfile {
	return storemodels.MarketProfile{
		Rent:             marketProfile.Rent,
		PowerPrice:       marketProfile.PowerPrice,
		Period:           marketProfile.Period,
		MinimumPeriod:    marketProfile.MinimumPeriod,
		MinimumIncome:    marketProfile.MinimumIncome,
		Id:               id,
		HeatPrice:        marketProfile.HeatPrice,
		Description:      marketProfile.Description,
		AvailabilityDate: marketProfile.AvailabilityDate,
		Deposit:          marketProfile.Deposit,
	}
}

func CreateStoreRentProfile(rentProfile requestmodels.RentProfile, id string) storemodels.RentProfile {
	return storemodels.RentProfile{
		ContractDue: rentProfile.ContractDue,
		HeatCosts:   rentProfile.HeatCosts,
		Deposit:     rentProfile.Deposit,
		Id:          id,
		Minimum:     rentProfile.Minimum,
		Rent:        rentProfile.Rent,
		RentalStart: rentProfile.RentalStart,
	}
}

func CreateStoreInvite(invite requestmodels.Invite, id string) storemodels.Invite {
	return storemodels.Invite{
		Id:       id,
		Property: invite.Property,
		Tenant:   invite.Tenant,
		Landlord: invite.Landlord,
		ValidDue: invite.ValidDue,
	}
}
func CreateStoreIncident(incident requestmodels.Incident, id string, propertyId string) storemodels.Incident {
	return storemodels.Incident{
		Description: incident.Description,
		Id:          id,
		Type:        incident.Type,
		TenantId:    incident.TenantId,
		Status:      incident.Status,
		PropertyId:  propertyId,
		LandlordId:  incident.LandlordId,
	}

}
