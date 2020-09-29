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

type ListDeviceCodesRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](#paginatingresults) for more information.
	Cursor string `json:"cursor,omitempty"`
	// If specified, only returns DeviceCodes of the specified location. Returns DeviceCodes of all locations if empty.
	LocationId  string       `json:"location_id,omitempty"`
	ProductType *ProductType `json:"product_type,omitempty"`
}
