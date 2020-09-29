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

// Represents a bulk update request for `TeamMember` objects.
type BulkUpdateTeamMembersRequest struct {
	// The data which will be used to update the `TeamMember` objects. Each key is the `team_member_id` that maps to the `UpdateTeamMemberRequest`.
	TeamMembers map[string]UpdateTeamMemberRequest `json:"team_members"`
}
