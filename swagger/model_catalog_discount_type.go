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

// CatalogDiscountType : How to apply a CatalogDiscount to a CatalogItem.
type CatalogDiscountType string

// List of CatalogDiscountType
const (
	FIXED_PERCENTAGE_CatalogDiscountType    CatalogDiscountType = "FIXED_PERCENTAGE"
	FIXED_AMOUNT_CatalogDiscountType        CatalogDiscountType = "FIXED_AMOUNT"
	VARIABLE_PERCENTAGE_CatalogDiscountType CatalogDiscountType = "VARIABLE_PERCENTAGE"
	VARIABLE_AMOUNT_CatalogDiscountType     CatalogDiscountType = "VARIABLE_AMOUNT"
)
