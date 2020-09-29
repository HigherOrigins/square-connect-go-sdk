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

type V1ListTimecardsRequest struct {
	Order *SortOrder `json:"order,omitempty"`
	// If provided, the endpoint returns only timecards for the employee with the specified ID.
	EmployeeId string `json:"employee_id,omitempty"`
	// If filtering results by their clockin_time field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginClockinTime string `json:"begin_clockin_time,omitempty"`
	// If filtering results by their clockin_time field, the end of the requested reporting period, in ISO 8601 format.
	EndClockinTime string `json:"end_clockin_time,omitempty"`
	// If filtering results by their clockout_time field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginClockoutTime string `json:"begin_clockout_time,omitempty"`
	// If filtering results by their clockout_time field, the end of the requested reporting period, in ISO 8601 format.
	EndClockoutTime string `json:"end_clockout_time,omitempty"`
	// If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginUpdatedAt string `json:"begin_updated_at,omitempty"`
	// If filtering results by their updated_at field, the end of the requested reporting period, in ISO 8601 format.
	EndUpdatedAt string `json:"end_updated_at,omitempty"`
	// If true, only deleted timecards are returned. If false, only valid timecards are returned.If you don't provide this parameter, both valid and deleted timecards are returned.
	Deleted bool `json:"deleted,omitempty"`
	// The maximum integer number of employee entities to return in a single response. Default 100, maximum 200.
	Limit int32 `json:"limit,omitempty"`
	// A pagination cursor to retrieve the next set of results for your original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`
}
