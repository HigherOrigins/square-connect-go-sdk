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

type OauthAuthorizationRevokedWebhookRevocationObject struct {
	// Timestamp of when the revocation event occurred, in RFC 3339 format.
	RevokedAt   string                                       `json:"revoked_at,omitempty"`
	RevokerType *OauthAuthorizationRevokedWebhookRevokerType `json:"revoker_type,omitempty"`
}
