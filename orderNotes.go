package woocommerce

import (
	"net/http"
)

// Order Notes service
type OrderNotesService service

// OrderNote object. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#order-note-properties
type OrderNote struct {
	Id             int    `json:"id,omitempty"`
	Author         string `json:"author,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateCreatedGmt string `json:"date_created_gmt,omitempty"`
	Note           string `json:"note,omitempty"`
	CustomerNote   bool   `json:"customer_note,omitempty"`
	AddedByUser    bool   `json:"added_by_user,omitempty"`
}

type ListOrderNotesParams struct {
	Context string `json:"context,omitempty"`
	Type    string `json:"type,omitempty"`
}

type DeleteOrderNoteParams struct {
	Force bool `url:"force"`
}

// Create an order Note. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-an-order-note
func (service *OrderNotesService) Create(orderId string, orderNote *OrderNote) (*OrderNote, *http.Response, error) {
	_url := "/orders/" + orderId + "/notes"
	req, _ := service.client.NewRequest("POST", _url, nil, orderNote)

	createdOrder := new(OrderNote)
	response, err := service.client.Do(req, createdOrder)

	if err != nil {
		return nil, response, err
	}

	return createdOrder, response, nil
}

// Get an order Note. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-an-order-note
func (service *OrderNotesService) Get(orderId string, noteId string) (*OrderNote, *http.Response, error) {
	_url := "/orders/" + orderId + "/notes/" + noteId
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	orderNote := new(OrderNote)
	response, err := service.client.Do(req, orderNote)

	if err != nil {
		return nil, response, err
	}

	return orderNote, response, nil
}

// List order Notes. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-order-notes
func (service *OrderNotesService) List(orderId string, opts *ListOrderNotesParams) (*[]OrderNote, *http.Response, error) {
	_url := "/orders/" + orderId + "/notes"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orders := new([]OrderNote)
	response, err := service.client.Do(req, orders)

	if err != nil {
		return nil, response, err
	}

	return orders, response, nil
}

// Delete an order Note. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-an-order-note
func (service *OrderNotesService) Delete(orderId string, noteId string, opts *DeleteOrderParams) (*OrderNote, *http.Response, error) {
	_url := "/orders/" + orderId + "/notes/" + noteId
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	orderNote := new(OrderNote)
	response, err := service.client.Do(req, orderNote)

	if err != nil {
		return nil, response, err
	}

	return orderNote, response, nil
}
