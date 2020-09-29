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

// LoyaltyRewardStatus : The status of the loyalty reward.
type LoyaltyRewardStatus string

// List of LoyaltyRewardStatus
const (
	ISSUED_LoyaltyRewardStatus   LoyaltyRewardStatus = "ISSUED"
	REDEEMED_LoyaltyRewardStatus LoyaltyRewardStatus = "REDEEMED"
	DELETED_LoyaltyRewardStatus  LoyaltyRewardStatus = "DELETED"
)
