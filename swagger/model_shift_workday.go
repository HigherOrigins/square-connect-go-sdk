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

// A `Shift` search query filter parameter that sets a range of days that a `Shift` must start or end in before passing the filter condition.
type ShiftWorkday struct {
	DateRange     *DateRange           `json:"date_range,omitempty"`
	MatchShiftsBy *ShiftWorkdayMatcher `json:"match_shifts_by,omitempty"`
	// Location-specific timezones convert workdays to datetime filters. Every location included in the query must have a timezone, or this field must be provided as a fallback. Format: the IANA timezone database identifier for the relevant timezone.
	DefaultTimezone string `json:"default_timezone,omitempty"`
}
