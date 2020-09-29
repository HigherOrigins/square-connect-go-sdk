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

// CashDrawerShiftState : The current state of a cash drawer shift.
type CashDrawerShiftState string

// List of CashDrawerShiftState
const (
	OPEN_CashDrawerShiftState   CashDrawerShiftState = "OPEN"
	ENDED_CashDrawerShiftState  CashDrawerShiftState = "ENDED"
	CLOSED_CashDrawerShiftState CashDrawerShiftState = "CLOSED"
)
