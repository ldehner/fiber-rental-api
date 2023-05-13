package userroutes

import (
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

func CreateResponseUser(userModel storemodels.User) responsemodels.User {
	return responsemodels.User{
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

func CreateResponseContactInfo(contactInfoModel storemodels.ContactInfo) responsemodels.ContactInfo {
	return responsemodels.ContactInfo{
		Phone: contactInfoModel.Phone,
		Mail:  contactInfoModel.Mail,
	}
}

func CreateResponseTenantInfo(tenantInfoModel storemodels.TenantInfo) responsemodels.TenantInfo {
	return responsemodels.TenantInfo{
		CriminalRecord: tenantInfoModel.CriminalRecord,
		Income:         tenantInfoModel.Income,
		IncomeProof:    tenantInfoModel.IncomeProof,
	}
}

func CreateResponseSearchProfile(searchProfileModel storemodels.SearchProfile) responsemodels.SearchProfile {
	return responsemodels.SearchProfile{
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

func CreateStoreUser(userModel requestmodels.User) storemodels.User {
	return storemodels.User{
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

func CreateStoreContactInfo(contactInfoModel requestmodels.ContactInfo, id string) storemodels.ContactInfo {
	return storemodels.ContactInfo{
		Phone: contactInfoModel.Phone,
		Mail:  contactInfoModel.Mail,
		ID:    id,
	}
}

func CreateStoreTenantInfo(tenantInfoModel requestmodels.TenantInfo, id string) storemodels.TenantInfo {
	return storemodels.TenantInfo{
		CriminalRecord: tenantInfoModel.CriminalRecord,
		Income:         tenantInfoModel.Income,
		IncomeProof:    tenantInfoModel.IncomeProof,
		Id:             id,
	}
}

func CreateStoreSearchProfile(searchProfileModel requestmodels.SearchProfile, id string) storemodels.SearchProfile {
	return storemodels.SearchProfile{
		Budget:  searchProfileModel.Budget,
		City:    searchProfileModel.City,
		Country: searchProfileModel.Country,
		Radius:  searchProfileModel.Radius,
		Rooms:   searchProfileModel.Rooms,
		Size:    searchProfileModel.Size,
		Street:  searchProfileModel.Street,
		Type:    searchProfileModel.Type,
		Zipcode: searchProfileModel.Zipcode,
		UserId:  id,
	}
}
