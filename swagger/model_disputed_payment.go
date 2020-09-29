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

// The payment the cardholder disputed.
type DisputedPayment struct {
	// Square-generated unique ID of the payment being disputed.
	PaymentId string `json:"payment_id,omitempty"`
}
