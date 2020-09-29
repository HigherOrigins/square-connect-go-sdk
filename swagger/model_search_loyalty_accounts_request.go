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

// A request to search for loyalty accounts.
type SearchLoyaltyAccountsRequest struct {
	Query *SearchLoyaltyAccountsRequestLoyaltyAccountQuery `json:"query,omitempty"`
	// The maximum number of results to include in the response.
	Limit int32 `json:"limit,omitempty"`
	// A pagination cursor returned by a previous call to  this endpoint. Provide this to retrieve the next set of  results for the original query.  For more information,  see [Pagination](https://developer.squareup.com/docs/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`
}
