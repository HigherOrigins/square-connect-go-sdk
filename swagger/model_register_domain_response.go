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

// Defines the fields that are included in the response body of a request to the [RegisterDomain](#endpoint-registerdomain) endpoint.  Either `errors` or `status` will be present in a given response (never both).
type RegisterDomainResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError                  `json:"errors,omitempty"`
	Status *RegisterDomainResponseStatus `json:"status,omitempty"`
}
