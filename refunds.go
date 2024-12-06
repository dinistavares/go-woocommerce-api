package woocommerce

import (
	"net/http"
)

// Refunds service
type RefundsService service

// Refund object. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#order-refund-properties
type Refund struct {
	Id              int               `json:"id,omitempty"`
	DateCreated     string            `json:"date_created,omitempty"`
	DateCreatedGmt  string            `json:"date_created_gmt,omitempty"`
	Amount          string            `json:"amount,omitempty"`
	Reason          string            `json:"reason,omitempty"`
	RefundedBy      int               `json:"refunded_by,omitempty"`
	RefundedPayment bool              `json:"refunded_payment,omitempty"`
	ApiRefund       bool              `json:"api_refund,omitempty"`
	MetaData        *[]MetaData       `json:"meta_data,omitempty"`
	LineItems       *[]RefundLineItem `json:"line_items,omitempty"`
}

type RefundLineItem struct {
	Id          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	ProductId   int          `json:"product_id,omitempty"`
	VariationId int          `json:"variation_id,omitempty"`
	Quantity    int          `json:"quantity,omitempty"`
	TaxClass    int          `json:"tax_class,omitempty"`
	Subtotal    string       `json:"subtotal,omitempty"`
	SubtotalTax string       `json:"subtotal_tax,omitempty"`
	Total       string       `json:"total,omitempty"`
	TotalTax    string       `json:"total_tax,omitempty"`
	Sku         string       `json:"sku,omitempty"`
	Price       string       `json:"price,omitempty"`
	RefundTotal float64      `json:"refund_total,omitempty"`
	Taxes       *[]RefundTax `json:"taxes,omitempty"`
	MetaData    *[]MetaData  `json:"meta_data,omitempty"`
}

type RefundTax struct {
	Id          int     `json:"id,omitempty"`
	Total       string  `json:"total,omitempty"`
	Subtotal    string  `json:"subtotal,omitempty"`
	RefundTotal float64 `json:"refund_total,omitempty"`
}

type ListRefundParams struct {
	Context       string      `url:"context,omitempty"`
	Page          int         `url:"page,omitempty"`
	PerPage       int         `url:"per_page,omitempty"`
	Search        string      `url:"search,omitempty"`
	Exclude       *[]int      `url:"exclude,omitempty"`
	Include       *[]int      `url:"include,omitempty"`
	Offset        int         `url:"offset,omitempty"`
	Order         string      `url:"order,omitempty"`
	OrderBy       string      `url:"orderby,omitempty"`
	After         string      `url:"after,omitempty"`
	Before        string      `url:"before,omitempty"`
	Orderby       string      `url:"orderby,omitempty"`
	Parent        interface{} `url:"parent,omitempty"`
	ParentExclude interface{} `url:"parent_exclude,omitempty"`
	Dp            int         `url:"dp,omitempty"`
}

type DeleteRefundParams struct {
	Force bool `url:"force"`
}

// Create a refund. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-refund
func (service *RefundsService) Create(orderId string, refund *Refund) (*Refund, *http.Response, error) {
	_url := "/orders/" + orderId + "/refunds"
	req, _ := service.client.NewRequest("POST", _url, nil, refund)

	createdRefund := new(Refund)
	response, err := service.client.Do(req, createdRefund)

	if err != nil {
		return nil, response, err
	}

	return createdRefund, response, nil
}

// Get a refund. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-refund
func (service *RefundsService) Get(orderId string, refundId string) (*Refund, *http.Response, error) {
	_url := "/orders/" + orderId + "/refunds/" + refundId
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	refund := new(Refund)
	response, err := service.client.Do(req, refund)

	if err != nil {
		return nil, response, err
	}

	return refund, response, nil
}

// List orders. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-refunds
func (service *RefundsService) List(orderId string, opts *ListRefundParams) (*[]Refund, *http.Response, error) {
	_url := "/orders/" + orderId + "/refunds"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orders := new([]Refund)
	response, err := service.client.Do(req, orders)

	if err != nil {
		return nil, response, err
	}

	return orders, response, nil
}

// Delete a refund. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-refund
func (service *RefundsService) Delete(orderId string, refundId string, opts *DeleteRefundParams) (*Refund, *http.Response, error) {
	_url := "/orders/" + orderId + "/refunds/" + refundId
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	refund := new(Refund)
	response, err := service.client.Do(req, refund)

	if err != nil {
		return nil, response, err
	}

	return refund, response, nil
}
