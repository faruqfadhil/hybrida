package merchant

type Merchant struct {
	ID          string
	Name        string
	Description string
	Address     string
	PhoneNumber string
}

type MerchantService interface{}
