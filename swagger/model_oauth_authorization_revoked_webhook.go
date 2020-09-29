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

// Published when a merchant/application revokes all access tokens and refresh tokens granted to an application.
type OauthAuthorizationRevokedWebhook struct {
	// The ID of the target merchant associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of event this represents, `\"oauth.authorization.revoked\"`.
	Type_ string `json:"type,omitempty"`
	// A unique ID for the webhook event.
	EventId string `json:"event_id,omitempty"`
	// Timestamp of when the webhook event was created, in RFC 3339 format.
	CreatedAt string                                `json:"created_at,omitempty"`
	Data      *OauthAuthorizationRevokedWebhookData `json:"data,omitempty"`
}
