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

// Information about fulfillment updates.
type OrderFulfillmentUpdatedUpdate struct {
	// Unique ID that identifies the fulfillment only within this order.
	FulfillmentUid string                 `json:"fulfillment_uid,omitempty"`
	OldState       *OrderFulfillmentState `json:"old_state,omitempty"`
	NewState       *OrderFulfillmentState `json:"new_state,omitempty"`
}
