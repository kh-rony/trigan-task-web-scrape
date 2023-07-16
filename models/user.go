package models

type User struct {
	Image            string `json:"image"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phone-number"`
	Email            string `json:"email"`
	IP               string `json:"ip"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	CreditCardNumber string `json:"credit-card-number"`
	ExpirationDate   string `json:"expiration-date"`
	IBAN             string `json:"iban"`
	SwiftBicNumber   string `json:"swift-bic-number"`
	Company          string `json:"company"`
}
