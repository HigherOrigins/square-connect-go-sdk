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

// Provides metadata when the event `type` is `DELETE_REWARD`.
type LoyaltyEventDeleteReward struct {
	// The ID of the [loyalty program](#type-LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The ID of the deleted [loyalty reward](#type-LoyaltyReward). This field is returned only if the event source is `LOYALTY_API`.
	RewardId string `json:"reward_id,omitempty"`
	// The number of points returned to the loyalty account.
	Points int32 `json:"points"`
}
