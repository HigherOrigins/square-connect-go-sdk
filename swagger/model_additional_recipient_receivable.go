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

// Represents a monetary distribution of part of a [Transaction](#type-transaction)'s amount for Transactions which included additional recipients. The location of this receivable is that same as the one specified in the [AdditionalRecipient](#type-additionalrecipient).
type AdditionalRecipientReceivable struct {
	// The additional recipient receivable's unique ID, issued by Square payments servers.
	Id string `json:"id"`
	// The ID of the transaction that the additional recipient receivable was applied to.
	TransactionId string `json:"transaction_id"`
	// The ID of the location that created the receivable. This is the location ID on the associated transaction.
	TransactionLocationId string `json:"transaction_location_id"`
	AmountMoney           *Money `json:"amount_money"`
	// The time when the additional recipient receivable was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// Any refunds of the receivable that have been applied.
	Refunds []AdditionalRecipientReceivableRefund `json:"refunds,omitempty"`
}
