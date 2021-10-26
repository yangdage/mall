package models

type AppAddressFormParam struct {
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

type AppAddressUpdateParam struct {
	Id              uint   `form:"id"`
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

type AppAddressDeleteParam struct {
	AddressId uint `form:"addressId"`
}

type AppAddressInfoParam struct {
	AddressId uint `form:"addressId"`
}

type AppAddressListParam struct {
	UserId string `form:"userId" json:"userId"`
}
