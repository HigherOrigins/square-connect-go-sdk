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

type PaymentOptions struct {
	// Indicates whether the Payment objects created from this `TerminalCheckout` will automatically be COMPLETED or left in an APPROVED state for later modification.
	Autocomplete bool `json:"autocomplete,omitempty"`
}
