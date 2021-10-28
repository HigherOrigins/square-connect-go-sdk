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

// Represents the tax IDs for an invoice recipient. The country of the seller account determines  whether the corresponding `tax_ids` field is available for the customer. For more information,  see [Invoice recipient tax IDs](https://developer.squareup.com/docs/invoices-api/overview#recipient-tax-ids).
type InvoiceRecipientTaxIds struct {
	// The EU VAT identification number for the invoice recipient. For example, `IE3426675K`.
	EuVat string `json:"eu_vat,omitempty"`
}
