package woocommerce

import (
	"net/http"
)

// Orders service
type OrdersService service

// Order object. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#order-properties
type Order struct {
	ID                 int              `json:"id,omitempty"`
	ParentID           int              `json:"parent_id,omitempty"`
	CustomerID         int              `json:"customer_id,omitempty"`
	PricesIncludeTax   bool             `json:"prices_include_tax"`
	NeedsPayment       bool             `json:"needs_payment"`
	NeedsProcessing    bool             `json:"needs_processing"`
	IsEditable         bool             `json:"is_editable"`
	CurrencySymbol     string           `json:"currency_symbol,omitempty"`
	Number             string           `json:"number,omitempty"`
	OrderKey           string           `json:"order_key,omitempty"`
	CreatedVia         string           `json:"created_via,omitempty"`
	Version            string           `json:"version,omitempty"`
	Status             string           `json:"status,omitempty"`
	Currency           string           `json:"currency,omitempty"`
	DateCreated        string           `json:"date_created,omitempty"`
	DateCreatedGmt     string           `json:"date_created_gmt,omitempty"`
	DateModified       string           `json:"date_modified,omitempty"`
	DateModifiedGmt    string           `json:"date_modified_gmt,omitempty"`
	DiscountTotal      string           `json:"discount_total,omitempty"`
	DiscountTax        string           `json:"discount_tax,omitempty"`
	ShippingTotal      string           `json:"shipping_total,omitempty"`
	ShippingTax        string           `json:"shipping_tax,omitempty"`
	CartTax            string           `json:"cart_tax,omitempty"`
	Total              string           `json:"total,omitempty"`
	TotalTax           string           `json:"total_tax,omitempty"`
	CustomerIPAddress  string           `json:"customer_ip_address,omitempty"`
	CustomerUserAgent  string           `json:"customer_user_agent,omitempty"`
	CustomerNote       string           `json:"customer_note,omitempty"`
	PaymentURL         string           `json:"payment_url,omitempty"`
	PaymentMethod      string           `json:"payment_method,omitempty"`
	PaymentMethodTitle string           `json:"payment_method_title,omitempty"`
	TransactionID      string           `json:"transaction_id,omitempty"`
	DatePaid           string           `json:"date_paid,omitempty"`
	DatePaidGmt        string           `json:"date_paid_gmt,omitempty"`
	DateCompleted      string           `json:"date_completed,omitempty"`
	DateCompletedGmt   string           `json:"date_completed_gmt,omitempty"`
	CartHash           string           `json:"cart_hash,omitempty"`
	Billing            *Billing         `json:"billing,omitempty"`
	Shipping           *Shipping        `json:"shipping,omitempty"`
	Links              *Links           `json:"_links,omitempty"`
	FeeLines           *[]FeeLine       `json:"fee_lines,omitempty"`
	Refunds            *[]OrderRefund   `json:"refunds,omitempty"`
	MetaData           *[]MetaData      `json:"meta_data,omitempty"`
	CouponLines        *[]CouponLine    `json:"coupon_lines,omitempty"`
	LineItems          *[]LineItems     `json:"line_items,omitempty"`
	TaxLines           *[]TaxLines      `json:"tax_lines,omitempty"`
	ShippingLines      *[]ShippingLines `json:"shipping_lines,omitempty"`
}

type OrderRefund struct {
	ID     int    `json:"id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Total  string `json:"total,omitempty"`
}

type CouponLine struct {
	Id          int       `json:"id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Discount    string    `json:"discount,omitempty"`
	DiscountTax string    `json:"discountTax,omitempty"`
	MetaData    *MetaData `json:"metaData,omitempty"`
}

type FeeLine struct {
	Id        int         `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	TaxClass  string      `json:"tax_class,omitempty"`
	TaxStatus string      `json:"tax_status,omitempty"`
	Amount    string      `json:"amount,omitempty"`
	Total     string      `json:"total,omitempty"`
	TotalTax  string      `json:"total_tax,omitempty"`
	Taxes     *[]Taxes    `json:"taxes,omitempty"`
	MetaData  *[]MetaData `json:"meta_data,omitempty"`
}

type Taxes struct {
	Id               int    `json:"id,omitempty"`
	Rate_code        string `json:"rate_code,omitempty"`
	Rate_id          string `json:"rate_id,omitempty"`
	Label            string `json:"label,omitempty"`
	Compound         bool   `json:"compound"`
	TaxTotal         string `json:"tax_total,omitempty"`
	ShippingTaxTotal string `json:"shipping_tax_total,omitempty"`
}

type LineItems struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	ProductID   int         `json:"product_id,omitempty"`
	VariationID int         `json:"variation_id,omitempty"`
	Quantity    int         `json:"quantity,omitempty"`
	TaxClass    string      `json:"tax_class,omitempty"`
	Subtotal    string      `json:"subtotal,omitempty"`
	SubtotalTax string      `json:"subtotal_tax,omitempty"`
	Total       string      `json:"total,omitempty"`
	TotalTax    string      `json:"total_tax,omitempty"`
	Taxes       *[]Taxes    `json:"taxes,omitempty"`
	MetaData    *[]MetaData `json:"meta_data,omitempty"`
	Sku         string      `json:"sku,omitempty"`
	Price       float64     `json:"price,omitempty"`
	Image       *Image      `json:"image,omitempty"`
	ParentName  interface{} `json:"parent_name,omitempty"`
}

type TaxLines struct {
	ID               int            `json:"id,omitempty"`
	RateCode         string         `json:"rate_code,omitempty"`
	RateID           int            `json:"rate_id,omitempty"`
	Label            string         `json:"label,omitempty"`
	Compound         bool           `json:"compound"`
	TaxTotal         string         `json:"tax_total,omitempty"`
	ShippingTaxTotal string         `json:"shipping_tax_total,omitempty"`
	RatePercent      float64        `json:"rate_percent"`
	MetaData         *[]interface{} `json:"meta_data,omitempty"`
}

type ShippingLines struct {
	ID          int            `json:"id,omitempty"`
	MethodTitle string         `json:"method_title,omitempty"`
	MethodID    string         `json:"method_id,omitempty"`
	InstanceID  string         `json:"instance_id,omitempty"`
	Total       string         `json:"total,omitempty"`
	TotalTax    string         `json:"total_tax,omitempty"`
	Taxes       *[]interface{} `json:"taxes,omitempty"`
	MetaData    *[]interface{} `json:"meta_data,omitempty"`
}

type ListOrdersParams struct {
	Context        string    `url:"context,omitempty"`
	Page           int       `url:"page,omitempty"`
	PerPage        int       `url:"per_page,omitempty"`
	Search         string    `url:"search,omitempty"`
	Exclude        *[]int    `url:"exclude,omitempty"`
	Include        *[]int    `url:"include,omitempty"`
	Offset         int       `url:"offset,omitempty"`
	Order          string    `url:"order,omitempty"`
	OrderBy        string    `url:"orderby,omitempty"`
	Customer       int       `url:"customer,omitempty"`
	Product        int       `url:"product,omitempty"`
	DecimalPoints  int       `url:"dp,omitempty"`
	Parent         *[]int    `url:"parent,omitempty"`
	ParentExclude  *[]int    `url:"parent_exclude,omitempty"`
	DatesAreGTM    bool      `url:"dates_are_gmt"`
	After          string    `url:"after,omitempty"`
	Before         string    `url:"before,omitempty"`
	ModifiedAfter  string    `url:"modified_after,omitempty"`
	ModifiedBefore string    `url:"modified_before,omitempty"`
	Status         *[]string `url:"status,omitempty"`
}

type GetOrderParams struct {
	DecimalPoints int `url:"dp,omitempty"`
}

type DeleteOrderParams struct {
	Force bool `url:"force"`
}

type BatchOrderUpdate struct {
	Create *[]Order `json:"create,omitempty"`
	Update *[]Order `json:"update,omitempty"`
	Delete *[]int   `json:"delete,omitempty"`
}

type BatchOrderUpdateResponse struct {
	Create *[]Order `json:"create,omitempty"`
	Update *[]Order `json:"update,omitempty"`
	Delete *[]Order `json:"delete,omitempty"`
}

// Create an order. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-an-order
func (service *OrdersService) Create(order *Order) (*Order, *http.Response, error) {
	_url := "/orders"
	req, _ := service.client.NewRequest("POST", _url, nil, order)

	createdOrder := new(Order)
	response, err := service.client.Do(req, createdOrder)

	if err != nil {
		return nil, response, err
	}

	return createdOrder, response, nil
}

// Get an order. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-an-order
func (service *OrdersService) Get(orderId string, opts *GetOrderParams) (*Order, *http.Response, error) {
	_url := "/orders/" + orderId
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	order := new(Order)
	response, err := service.client.Do(req, order)

	if err != nil {
		return nil, response, err
	}

	return order, response, nil
}

// List orders. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-orders
func (service *OrdersService) List(opts *ListOrdersParams) (*[]Order, *http.Response, error) {
	_url := "/orders"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orders := new([]Order)
	response, err := service.client.Do(req, orders)

	if err != nil {
		return nil, response, err
	}

	return orders, response, nil
}

// Update an order. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-an-order
func (service *OrdersService) Update(orderId string, order *Order) (*Order, *http.Response, error) {
	_url := "/orders/" + orderId
	req, _ := service.client.NewRequest("PUT", _url, nil, order)

	updatedOrder := new(Order)
	response, err := service.client.Do(req, updatedOrder)

	if err != nil {
		return nil, response, err
	}

	return updatedOrder, response, nil
}

// Delete an order. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-an-order
func (service *OrdersService) Delete(orderId string, opts *DeleteOrderParams) (*Order, *http.Response, error) {
	_url := "/orders/" + orderId
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	order := new(Order)
	response, err := service.client.Do(req, order)

	if err != nil {
		return nil, response, err
	}

	return order, response, nil
}

// Batch update orders. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-orders
func (service *OrdersService) Batch(opts *BatchOrderUpdate) (*BatchOrderUpdateResponse, *http.Response, error) {
	_url := "/orders/batch"
	req, _ := service.client.NewRequest("POST", _url, opts, nil)

	orders := new(BatchOrderUpdateResponse)
	response, err := service.client.Do(req, orders)

	if err != nil {
		return nil, response, err
	}

	return orders, response, nil
}
