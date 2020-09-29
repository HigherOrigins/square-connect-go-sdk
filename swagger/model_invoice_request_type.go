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

// InvoiceRequestType : Identifies the type of the payment request. For more information,  see [Payment request](TBD).
type InvoiceRequestType string

// List of InvoiceRequestType
const (
	BALANCE_InvoiceRequestType     InvoiceRequestType = "BALANCE"
	DEPOSIT_InvoiceRequestType     InvoiceRequestType = "DEPOSIT"
	INSTALLMENT_InvoiceRequestType InvoiceRequestType = "INSTALLMENT"
)
