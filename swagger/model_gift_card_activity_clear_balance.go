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

// Represents details about a `CLEAR_BALANCE` [gift card activity type](entity:GiftCardActivityType).
type GiftCardActivityClearBalance struct {
	Reason *GiftCardActivityClearBalanceReason `json:"reason"`
}
