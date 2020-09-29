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

// The hourly wage rate that an employee will earn on a `Shift` for doing the job specified by the `title` property of this object. Deprecated at verison 2020-08-26. Use `TeamMemberWage` instead.
type EmployeeWage struct {
	// UUID for this object.
	Id string `json:"id,omitempty"`
	// The `Employee` that this wage is assigned to.
	EmployeeId string `json:"employee_id,omitempty"`
	// The job title that this wage relates to.
	Title      string `json:"title,omitempty"`
	HourlyRate *Money `json:"hourly_rate,omitempty"`
}
