package userroutes

import (
	"github.com/ldehner/fiber-rental-api/models"
)

func CreateResponseUser(userModel models.User) User {
	return User{
		Id:          userModel.Id,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Phone:       userModel.Phone,
		Mail:        userModel.Mail,
		Country:     userModel.Country,
		City:        userModel.City,
		Street:      userModel.Street,
		Housenumber: userModel.Housenumber,
		Apartment:   userModel.Apartment,
	}
}

func CreateResponseContactInfo(contactInfoModel models.ContactInfo) ContactInfo {
	return ContactInfo{
		Phone: contactInfoModel.Phone,
		Mail:  contactInfoModel.Mail,
	}
}

func CreateResponseTenantInfo(tenantInfoModel models.TenantInfo) TenantInfo {
	return TenantInfo{
		CriminalRecord: tenantInfoModel.CriminalRecord,
		Income:         tenantInfoModel.Income,
		IncomeProof:    tenantInfoModel.IncomeProof,
	}
}

func CreateResponseSearchProfile(searchProfileModel models.SearchProfile) SearchProfile {
	return SearchProfile{
		Budget:  searchProfileModel.Budget,
		City:    searchProfileModel.City,
		Country: searchProfileModel.Country,
		Radius:  searchProfileModel.Radius,
		Rooms:   searchProfileModel.Rooms,
		Size:    searchProfileModel.Size,
		Street:  searchProfileModel.Street,
		Type:    searchProfileModel.Type,
		Zipcode: searchProfileModel.Zipcode,
	}
}
