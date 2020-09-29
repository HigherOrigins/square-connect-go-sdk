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

type V1ListItemsRequest struct {
	// A pagination cursor to retrieve the next set of results for your original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`
}
