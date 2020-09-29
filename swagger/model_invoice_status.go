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

// InvoiceStatus : Indicates the status of an invoice.
type InvoiceStatus string

// List of InvoiceStatus
const (
	DRAFT_InvoiceStatus              InvoiceStatus = "DRAFT"
	UNPAID_InvoiceStatus             InvoiceStatus = "UNPAID"
	SCHEDULED_InvoiceStatus          InvoiceStatus = "SCHEDULED"
	PARTIALLY_PAID_InvoiceStatus     InvoiceStatus = "PARTIALLY_PAID"
	PAID_InvoiceStatus               InvoiceStatus = "PAID"
	PARTIALLY_REFUNDED_InvoiceStatus InvoiceStatus = "PARTIALLY_REFUNDED"
	REFUNDED_InvoiceStatus           InvoiceStatus = "REFUNDED"
	CANCELED_InvoiceStatus           InvoiceStatus = "CANCELED"
	FAILED_InvoiceStatus             InvoiceStatus = "FAILED"
	PAYMENT_PENDING_InvoiceStatus    InvoiceStatus = "PAYMENT_PENDING"
)
