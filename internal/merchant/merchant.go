package merchant

type Merchant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type MerchantService interface{}
