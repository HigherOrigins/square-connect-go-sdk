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

type BatchChangeInventoryRequest struct {
	// A client-supplied, universally unique identifier (UUID) for the request.  See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) in the [API Development 101](https://developer.squareup.com/docs/basics/api101/overview) section for more information.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// The set of physical counts and inventory adjustments to be made. Changes are applied based on the client-supplied timestamp and may be sent out of order.
	Changes []InventoryChange `json:"changes,omitempty"`
	// Indicates whether the current physical count should be ignored if the quantity is unchanged since the last physical count. Default: `true`.
	IgnoreUnchangedCounts bool `json:"ignore_unchanged_counts,omitempty"`
}
