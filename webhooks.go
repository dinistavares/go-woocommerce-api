package woocommerce

import (
	"net/http"
)

// Webhooks service
type WebhookService service

type Webhook struct {
	Id              int      `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Status          string   `json:"status,omitempty"`
	Topic           string   `json:"topic,omitempty"`
	Resource        string   `json:"resource,omitempty"`
	Event           string   `json:"event,omitempty"`
	Hooks           []string `json:"hooks,omitempty"`
	DeliveryUrl     string   `json:"delivery_url,omitempty"`
	Secret          string   `json:"secret,omitempty"`
	DateCreated     string   `json:"date_created,omitempty"`
	DateCreatedGmt  string   `json:"date_created_gmt,omitempty"`
	DateModified    string   `json:"date_modified,omitempty"`
	DateModifiedGmt string   `json:"date_modified_gmt,omitempty"`
	Links           *Links   `json:"links,omitempty"`
}

type ListWebhooksParams struct {
	Context string `url:"context,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Search  string `url:"search,omitempty"`
	Exclude *[]int `url:"exclude,omitempty"`
	Include *[]int `url:"include,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"orderby,omitempty"`

	After  string `url:"after,omitempty"`
	Before string `url:"before,omitempty"`
	Status string `url:"status,omitempty"`
}

type DeleteWebhookParams struct {
	Force string `json:"force,omitempty"`
}

type BatchWebhookUpdate struct {
	Create *[]Webhook `json:"create,omitempty"`
	Delete *[]int     `json:"delete,omitempty"`
}

type BatchWebhookUpdateResponse struct {
	Create *[]Webhook `json:"create,omitempty"`
	Delete *[]Webhook `json:"delete,omitempty"`
}

// Create a webhook. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-webhook
func (service *WebhookService) Create(webhook *Webhook) (*Webhook, *http.Response, error) {
	_url := "/webhooks"
	req, _ := service.client.NewRequest("POST", _url, nil, webhook)

	createdWebhook := new(Webhook)
	response, err := service.client.Do(req, createdWebhook)

	if err != nil {
		return nil, response, err
	}

	return createdWebhook, response, nil
}

// Get a wehook. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-webhook
func (service *WebhookService) Get(webhookID string) (*Webhook, *http.Response, error) {
	_url := "/webhooks/" + webhookID
	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	webhook := new(Webhook)
	response, err := service.client.Do(req, webhook)

	if err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

// List Webhooks. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-webhooks
func (service *WebhookService) List(opts *ListWebhooksParams) (*[]Webhook, *http.Response, error) {
	_url := "/webhooks"
	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	webhooks := new([]Webhook)
	response, err := service.client.Do(req, webhooks)

	if err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}

// Update a webhook. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-webhook
func (service *WebhookService) Update(webhookID string, webhook *Webhook) (*Webhook, *http.Response, error) {
	_url := "/webhooks/" + webhookID
	req, _ := service.client.NewRequest("PUT", _url, nil, webhook)

	updatedWebhook := new(Webhook)
	response, err := service.client.Do(req, updatedWebhook)

	if err != nil {
		return nil, response, err
	}

	return updatedWebhook, response, nil
}

// Delete a webhook. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-webhook
func (service *WebhookService) Delete(webhookID string, opts *DeleteWebhookParams) (*Webhook, *http.Response, error) {
	_url := "/webhooks/" + webhookID
	req, _ := service.client.NewRequest("DELETE", _url, opts, nil)

	webhook := new(Webhook)
	response, err := service.client.Do(req, webhook)

	if err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

// Batch update webhooks. Reference: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-webhooks
func (service *WebhookService) Batch(opts *BatchWebhookUpdate) (*BatchWebhookUpdateResponse, *http.Response, error) {
	_url := "/webhooks/batch"
	req, _ := service.client.NewRequest("POST", _url, nil, opts)

	webhooks := new(BatchWebhookUpdateResponse)
	response, err := service.client.Do(req, webhooks)

	if err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}
