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

// Sets the day of the week and hour of the day that a business starts a workweek. This is used to calculate overtime pay.
type WorkweekConfig struct {
	// The UUID for this object.
	Id          string   `json:"id,omitempty"`
	StartOfWeek *Weekday `json:"start_of_week"`
	// The local time at which a business week ends. Represented as a string in `HH:MM` format (`HH:MM:SS` is also accepted, but seconds are truncated).
	StartOfDayLocalTime string `json:"start_of_day_local_time"`
	// Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If not provided, Square executes a blind write; potentially overwriting data from another write.
	Version int32 `json:"version,omitempty"`
	// A read-only timestamp in RFC 3339 format; presented in UTC.
	CreatedAt string `json:"created_at,omitempty"`
	// A read-only timestamp in RFC 3339 format; presented in UTC.
	UpdatedAt string `json:"updated_at,omitempty"`
}
