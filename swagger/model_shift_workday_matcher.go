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

// ShiftWorkdayMatcher : Defines the logic used to apply a workday filter.
type ShiftWorkdayMatcher string

// List of ShiftWorkdayMatcher
const (
	START_AT_ShiftWorkdayMatcher     ShiftWorkdayMatcher = "START_AT"
	END_AT_ShiftWorkdayMatcher       ShiftWorkdayMatcher = "END_AT"
	INTERSECTION_ShiftWorkdayMatcher ShiftWorkdayMatcher = "INTERSECTION"
)
