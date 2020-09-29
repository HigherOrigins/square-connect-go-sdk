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

// Pricing options for an order. The options affect how the order's price is calculated. They can be used, for example, to apply automatic price adjustments that are based on pre-configured [pricing rules](https://developer.squareup.com/docs/reference/square/objects/CatalogPricingRule).
type OrderPricingOptions struct {
	// The option to determine whether or not pricing rule-based discounts are automatically applied to an order.
	AutoApplyDiscounts bool `json:"auto_apply_discounts,omitempty"`
}
