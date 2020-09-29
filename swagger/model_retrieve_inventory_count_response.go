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

type RetrieveInventoryCountResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The current calculated inventory counts for the requested object and locations.
	Counts []InventoryCount `json:"counts,omitempty"`
	// The pagination cursor to be used in a subsequent request. If unset, this is the final response.  See the [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
}
