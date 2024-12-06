package woocommerce

type MetaData struct {
	ID         int         `json:"id,omitempty"`
	Key        string      `json:"key,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	DisplayKey string      `json:"display_key"`
}

type Self struct {
	Href string `json:"href,omitempty"`
}

type Collection struct {
	Href string `json:"href,omitempty"`
}

type Links struct {
	Self       []Self       `json:"self,omitempty"`
	Collection []Collection `json:"collection,omitempty"`
}

type Billing struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address_1,omitempty"`
	Address2  string `json:"address_2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type Shipping struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address_1,omitempty"`
	Address2  string `json:"address_2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
