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

// Defines an appointment slot that encapsulates the appointment segments, location and starting time available for booking.
type Availability struct {
	// The RFC 3339 timestamp specifying the beginning time of the slot available for booking.
	StartAt string `json:"start_at,omitempty"`
	// The ID of the location available for booking.
	LocationId string `json:"location_id,omitempty"`
	// The list of appointment segments available for booking
	AppointmentSegments []AppointmentSegment `json:"appointment_segments,omitempty"`
}
