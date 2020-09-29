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

// Contains details necessary to fulfill a pickup order.
type OrderFulfillmentPickupDetails struct {
	Recipient *OrderFulfillmentRecipient `json:"recipient,omitempty"`
	// The [timestamp](#workingwithdates) indicating when this fulfillment will expire if it is not accepted. Must be in RFC 3339 format e.g., \"2016-09-04T23:59:33.123Z\". Expiration time can only be set up to 7 days in the future. If `expires_at` is not set, this pickup fulfillment will be automatically accepted when placed.
	ExpiresAt string `json:"expires_at,omitempty"`
	// The duration of time after which an open and accepted pickup fulfillment will automatically move to the `COMPLETED` state. Must be in RFC3339 duration format e.g., \"P1W3D\".  If not set, this pickup fulfillment will remain accepted until it is canceled or completed.
	AutoCompleteDuration string                                     `json:"auto_complete_duration,omitempty"`
	ScheduleType         *OrderFulfillmentPickupDetailsScheduleType `json:"schedule_type,omitempty"`
	// The [timestamp](#workingwithdates) that represents the start of the pickup window. Must be in RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\". For fulfillments with the schedule type `ASAP`, this is automatically set to the current time plus the expected duration to prepare the fulfillment.
	PickupAt string `json:"pickup_at,omitempty"`
	// The window of time in which the order should be picked up after the `pickup_at` timestamp. Must be in RFC3339 duration format, e.g., \"P1W3D\". Can be used as an informational guideline for merchants.
	PickupWindowDuration string `json:"pickup_window_duration,omitempty"`
	// The duration of time it takes to prepare this fulfillment. Must be in RFC3339 duration format, e.g., \"P1W3D\".
	PrepTimeDuration string `json:"prep_time_duration,omitempty"`
	// A note meant to provide additional instructions about the pickup fulfillment displayed in the Square Point of Sale and set by the API.
	Note string `json:"note,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment was placed. Must be in RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	PlacedAt string `json:"placed_at,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment was accepted. In RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	AcceptedAt string `json:"accepted_at,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment was rejected. In RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	RejectedAt string `json:"rejected_at,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment is marked as ready for pickup. In RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	ReadyAt string `json:"ready_at,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment expired. In RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	ExpiredAt string `json:"expired_at,omitempty"`
	// The [timestamp](#workingwithdates) indicating when the fulfillment was picked up by the recipient. In RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	PickedUpAt string `json:"picked_up_at,omitempty"`
	// The [timestamp](#workingwithdates) in RFC3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\", indicating when the fulfillment was canceled.
	CanceledAt string `json:"canceled_at,omitempty"`
	// A description of why the pickup was canceled. Max length: 100 characters.
	CancelReason string `json:"cancel_reason,omitempty"`
	// If true, indicates this pickup order is for curbside pickup, not in-store pickup.
	IsCurbsidePickup      bool                                                `json:"is_curbside_pickup,omitempty"`
	CurbsidePickupDetails *OrderFulfillmentPickupDetailsCurbsidePickupDetails `json:"curbside_pickup_details,omitempty"`
}
