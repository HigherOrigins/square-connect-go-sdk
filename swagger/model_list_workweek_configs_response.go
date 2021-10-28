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

// The response to a request for a set of `WorkweekConfig` objects. The response contains the requested `WorkweekConfig` objects and might contain a set of `Error` objects if the request resulted in errors.
type ListWorkweekConfigsResponse struct {
	// A page of `EmployeeWage` results.
	WorkweekConfigs []WorkweekConfig `json:"workweek_configs,omitempty"`
	// The value supplied in the subsequent request to fetch the next page of `EmployeeWage` results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
