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

// ShiftStatus : Enumerates the possible status of a `Shift`.
type ShiftStatus string

// List of ShiftStatus
const (
	UNKNOWN_STATUS_ShiftStatus ShiftStatus = "UNKNOWN_STATUS"
	OPEN_ShiftStatus           ShiftStatus = "OPEN"
	CLOSED_ShiftStatus         ShiftStatus = "CLOSED"
)
