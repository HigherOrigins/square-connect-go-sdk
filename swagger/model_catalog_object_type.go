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

// CatalogObjectType : Possible types of CatalogObjects returned from the catalog, each containing type-specific properties in the `*_data` field corresponding to the specfied object type.
type CatalogObjectType string

// List of CatalogObjectType
const (
	CATALOG_OBJECT_TYPE_DO_NOT_USE_CatalogObjectType CatalogObjectType = "CATALOG_OBJECT_TYPE_DO_NOT_USE"
	ITEM_CatalogObjectType                           CatalogObjectType = "ITEM"
	IMAGE_CatalogObjectType                          CatalogObjectType = "IMAGE"
	CATEGORY_CatalogObjectType                       CatalogObjectType = "CATEGORY"
	ITEM_VARIATION_CatalogObjectType                 CatalogObjectType = "ITEM_VARIATION"
	TAX_CatalogObjectType                            CatalogObjectType = "TAX"
	DISCOUNT_CatalogObjectType                       CatalogObjectType = "DISCOUNT"
	MODIFIER_LIST_CatalogObjectType                  CatalogObjectType = "MODIFIER_LIST"
	MODIFIER_CatalogObjectType                       CatalogObjectType = "MODIFIER"
	DINING_OPTION_CatalogObjectType                  CatalogObjectType = "DINING_OPTION"
	TAX_EXEMPTION_CatalogObjectType                  CatalogObjectType = "TAX_EXEMPTION"
	SERVICE_CHARGE_CatalogObjectType                 CatalogObjectType = "SERVICE_CHARGE"
	PRICING_RULE_CatalogObjectType                   CatalogObjectType = "PRICING_RULE"
	PRODUCT_SET_CatalogObjectType                    CatalogObjectType = "PRODUCT_SET"
	TIME_PERIOD_CatalogObjectType                    CatalogObjectType = "TIME_PERIOD"
	MEASUREMENT_UNIT_CatalogObjectType               CatalogObjectType = "MEASUREMENT_UNIT"
	SUBSCRIPTION_PLAN_CatalogObjectType              CatalogObjectType = "SUBSCRIPTION_PLAN"
	ITEM_OPTION_CatalogObjectType                    CatalogObjectType = "ITEM_OPTION"
	ITEM_OPTION_VAL_CatalogObjectType                CatalogObjectType = "ITEM_OPTION_VAL"
	CUSTOM_ATTRIBUTE_DEFINITION_CatalogObjectType    CatalogObjectType = "CUSTOM_ATTRIBUTE_DEFINITION"
	QUICK_AMOUNTS_SETTINGS_CatalogObjectType         CatalogObjectType = "QUICK_AMOUNTS_SETTINGS"
	COMPONENT_CatalogObjectType                      CatalogObjectType = "COMPONENT"
	COMPOSITION_CatalogObjectType                    CatalogObjectType = "COMPOSITION"
	RESOURCE_CatalogObjectType                       CatalogObjectType = "RESOURCE"
	CHECKOUT_LINK_CatalogObjectType                  CatalogObjectType = "CHECKOUT_LINK"
	ADDRESS_CatalogObjectType                        CatalogObjectType = "ADDRESS"
)
