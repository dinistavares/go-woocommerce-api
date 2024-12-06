package woocommerce

import (
	"net/http"
)

// Customer service
type CustomersService service

// Customer object. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#customer-properties
type Customer struct {
	ID               int         `json:"id,omitempty"`
	DateCreated      string      `json:"date_created,omitempty"`
	DateCreatedGmt   string      `json:"date_created_gmt,omitempty"`
	DateModified     string      `json:"date_modified,omitempty"`
	DateModifiedGmt  string      `json:"date_modified_gmt,omitempty"`
	Email            string      `json:"email,omitempty"`
	FirstName        string      `json:"first_name,omitempty"`
	LastName         string      `json:"last_name,omitempty"`
	Role             string      `json:"role,omitempty"`
	Username         string      `json:"username,omitempty"`
	AvatarURL        string      `json:"avatar_url,omitempty"`
	IsPayingCustomer bool        `json:"is_paying_customer"`
	MetaData         *[]MetaData `json:"meta_data,omitempty"`
	Billing          *Billing    `json:"billing,omitempty"`
	Shipping         *Shipping   `json:"shipping,omitempty"`
	Links            *Links      `json:"_links,omitempty"`
}

type ListCustomerParams struct {
	Context string `url:"context,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Search  string `url:"search,omitempty"`
	Exclude *[]int `url:"exclude,omitempty"`
	Include *[]int `url:"include,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"orderby,omitempty"`

	Email string `url:"email,omitempty"`
	Role  string `url:"role,omitempty"`
}

type DeleteCustomerParams struct {
	Force    string `json:"force,omitempty"`
	Reassign int    `json:"reassign,omitempty"`
}

type BatchCustomerUpdate struct {
	Create *[]Customer `json:"create,omitempty"`
	Update *[]Customer `json:"update,omitempty"`
	Delete *[]int      `json:"delete,omitempty"`
}

type BatchCustomerUpdateResponse struct {
	Create *[]Customer `json:"create,omitempty"`
	Update *[]Customer `json:"update,omitempty"`
	Delete *[]Customer `json:"delete,omitempty"`
}

type CustomerDownload struct {
	DownloadId         string `json:"download_id,omitempty"`
	DownloadUrl        string `json:"download_url,omitempty"`
	ProductId          int    `json:"product_id,omitempty"`
	ProductName        string `json:"product_name,omitempty"`
	DownloadName       string `json:"download_name,omitempty"`
	OrderId            int    `json:"order_id,omitempty"`
	OrderKey           string `json:"order_key,omitempty"`
	DownloadsRemaining string `json:"downloads_remaining,omitempty"`
	AccessExpires      string `json:"access_expires,omitempty"`
	AccessExpiresGmt   string `json:"access_expires_gmt,omitempty"`
	File               *File  `json:"file,omitempty"`
	Links              *Links `json:"_links,omitempty"`
}

type File struct {
	Name string `json:"name,omitempty"`
	File string `json:"file,omitempty"`
}

// Create a customer. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-customer
func (service *CustomersService) Create(customer *Customer) (*Customer, *http.Response, error) {
	_url := "/customers"
	req, _ := service.client.NewRequest("POST", _url, nil, customer)

	createdCustomer := new(Customer)
	response, err := service.client.Do(req, createdCustomer)

	if err != nil {
		return nil, response, err
	}

	return createdCustomer, response, nil
}

// Get a customer. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-customer
func (service *CustomersService) Get(customerID string) (*Customer, *http.Response, error) {
	_url := "/customers/" + customerID
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	customer := new(Customer)
	response, err := service.client.Do(req, customer)

	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}

// List customers. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-customers
func (service *CustomersService) List(opts *ListCustomerParams) (*[]Customer, *http.Response, error) {
	_url := "/customers"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	customers := new([]Customer)
	response, err := service.client.Do(req, customers)

	if err != nil {
		return nil, response, err
	}

	return customers, response, nil
}

// Update a customer. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-customer
func (service *CustomersService) Update(customerID string, customer *Customer) (*Customer, *http.Response, error) {
	_url := "/customers/" + customerID
	req, _ := service.client.NewRequest("PUT", _url, nil, customer)

	updatedCustomer := new(Customer)
	response, err := service.client.Do(req, updatedCustomer)

	if err != nil {
		return nil, response, err
	}

	return updatedCustomer, response, nil
}

// Delete a customer. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-customer
func (service *CustomersService) Delete(customerID string, opts *DeleteCustomerParams) (*Customer, *http.Response, error) {
	_url := "/customers/" + customerID
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	customer := new(Customer)
	response, err := service.client.Do(req, customer)

	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}

// Batch update customers. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-customers
func (service *CustomersService) Batch(opts *BatchCustomerUpdate) (*BatchCustomerUpdateResponse, *http.Response, error) {
	_url := "/customers/batch"
	req, _ := service.client.NewRequest("POST", _url, nil, opts)

	customers := new(BatchCustomerUpdateResponse)
	response, err := service.client.Do(req, customers)

	if err != nil {
		return nil, response, err
	}

	return customers, response, nil
}

// Get customer downloads. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-customer-downloads
func (service *CustomersService) GetDownloads(customerID string) (*[]CustomerDownload, *http.Response, error) {
	_url := "/customers/" + customerID + "/downloads"
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	downloads := new([]CustomerDownload)
	response, err := service.client.Do(req, downloads)

	if err != nil {
		return nil, response, err
	}

	return downloads, response, nil
}
