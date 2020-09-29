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

type LoyaltyAccountCreatedWebhookData struct {
	// Name of the affected object’s type, `\"loyalty.account\"`.
	Type_ string `json:"type,omitempty"`
	// ID of the affected loyalty account.
	Id     string          `json:"id,omitempty"`
	Object *LoyaltyAccount `json:"object,omitempty"`
}
