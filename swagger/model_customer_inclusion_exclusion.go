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

// CustomerInclusionExclusion : Indicates whether customers should be included in, or excluded from, the result set when they match the filtering criteria.
type CustomerInclusionExclusion string

// List of CustomerInclusionExclusion
const (
	INCLUDE_CustomerInclusionExclusion CustomerInclusionExclusion = "INCLUDE"
	EXCLUDE_CustomerInclusionExclusion CustomerInclusionExclusion = "EXCLUDE"
)
