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

// Published when a customer [custom attribute](entity:CustomAttribute) that is visible to the subscribing application is created or updated. A notification is sent when: - Your application creates or updates a custom attribute owned by your application, regardless of the `visibility` setting. - Any application creates or updates a custom attribute whose `visibility` is `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.  Custom attributes set to `VISIBILITY_READ_WRITE_VALUES` can be created or updated by any application, but those set to `VISIBILITY_READ_ONLY` or `VISIBILITY_HIDDEN` can only be created or updated by the owner. Custom attributes are owned by the application that created the corresponding [custom attribute definition](entity:CustomAttributeDefinition).
type CustomerCustomAttributeVisibleUpdatedWebhook struct {
	// The ID of the seller associated with the event that triggered the event notification.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of this event. The value is `\"customer.custom_attribute.visible.updated\"`.
	Type_ string `json:"type,omitempty"`
	// A unique ID for the event notification.
	EventId string `json:"event_id,omitempty"`
	// The timestamp that indicates when the event notification was created, in RFC 3339 format.
	CreatedAt string                      `json:"created_at,omitempty"`
	Data      *CustomAttributeWebhookData `json:"data,omitempty"`
}
