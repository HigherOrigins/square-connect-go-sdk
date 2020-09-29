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

// Defines the fields included in the request body for the SearchCustomers endpoint.
type SearchCustomersRequest struct {
	// Include the pagination cursor in subsequent calls to this endpoint to retrieve the next set of results associated with the original query.  See the [Pagination guide](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
	// A limit on the number of results to be returned in a single page. The limit is advisory - the implementation may return more or fewer results. If the supplied limit is negative, zero, or is higher than the maximum limit of 100, it will be ignored.
	Limit int64          `json:"limit,omitempty"`
	Query *CustomerQuery `json:"query,omitempty"`
}
