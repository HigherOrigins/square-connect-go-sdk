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

// Represents a set of SearchSubscriptionsQuery filters used to limit the set of Subscriptions returned by SearchSubscriptions.
type SearchSubscriptionsFilter struct {
	// A filter to select subscriptions based on the customer.
	CustomerIds []string `json:"customer_ids,omitempty"`
	// A filter to select subscriptions based the location.
	LocationIds []string `json:"location_ids,omitempty"`
	// A filter to select subscriptions based on the source application.
	SourceNames []string `json:"source_names,omitempty"`
}
