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

// Sorting criteria for a SearchOrders request. Results can only be sorted by a timestamp field.
type SearchOrdersSort struct {
	SortField *SearchOrdersSortField `json:"sort_field"`
	SortOrder *SortOrder             `json:"sort_order,omitempty"`
}
