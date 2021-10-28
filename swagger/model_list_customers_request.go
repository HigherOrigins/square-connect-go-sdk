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

// Defines the query parameters that can be included in a request to the `ListCustomers` endpoint.
type ListCustomersRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. The limit is ignored if it is less than 1 or greater than 100. The default value is 100.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Limit     int32              `json:"limit,omitempty"`
	SortField *CustomerSortField `json:"sort_field,omitempty"`
	SortOrder *SortOrder         `json:"sort_order,omitempty"`
}
