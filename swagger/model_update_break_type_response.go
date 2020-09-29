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

// A response to a request to update a `BreakType`. Contains the requested `BreakType` objects. May contain a set of `Error` objects if the request resulted in errors.
type UpdateBreakTypeResponse struct {
	BreakType *BreakType `json:"break_type,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
