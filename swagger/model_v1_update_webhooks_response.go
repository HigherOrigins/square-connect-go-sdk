/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// V1UpdateWebhooksResponse
type V1UpdateWebhooksResponse struct {
	// A list of webhook event enums the location is currently subscribed to. See [V1WebhooksEvents](#type-v1webhooksevents) for possible values
	Events []V1WebhooksEvents `json:"events,omitempty"`
}
