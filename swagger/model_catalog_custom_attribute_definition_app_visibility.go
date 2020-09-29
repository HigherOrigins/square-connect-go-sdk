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

// CatalogCustomAttributeDefinitionAppVisibility : Defines the visibility of a custom attribute to applications other than their creating application.
type CatalogCustomAttributeDefinitionAppVisibility string

// List of CatalogCustomAttributeDefinitionAppVisibility
const (
	HIDDEN_CatalogCustomAttributeDefinitionAppVisibility            CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_HIDDEN"
	READ_ONLY_CatalogCustomAttributeDefinitionAppVisibility         CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_READ_ONLY"
	READ_WRITE_VALUES_CatalogCustomAttributeDefinitionAppVisibility CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_READ_WRITE_VALUES"
)
