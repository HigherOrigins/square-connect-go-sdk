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

// SubscriptionEventSubscriptionEventType : The possible subscription event types.
type SubscriptionEventSubscriptionEventType string

// List of SubscriptionEventSubscriptionEventType
const (
	DEFAULT_SUBSCRIPTION_EVENT_TYPE_DO_NOT_USE_SubscriptionEventSubscriptionEventType SubscriptionEventSubscriptionEventType = "DEFAULT_SUBSCRIPTION_EVENT_TYPE_DO_NOT_USE"
	START_SUBSCRIPTION_SubscriptionEventSubscriptionEventType                         SubscriptionEventSubscriptionEventType = "START_SUBSCRIPTION"
	PLAN_CHANGE_SubscriptionEventSubscriptionEventType                                SubscriptionEventSubscriptionEventType = "PLAN_CHANGE"
	STOP_SUBSCRIPTION_SubscriptionEventSubscriptionEventType                          SubscriptionEventSubscriptionEventType = "STOP_SUBSCRIPTION"
	DEACTIVATE_SUBSCRIPTION_SubscriptionEventSubscriptionEventType                    SubscriptionEventSubscriptionEventType = "DEACTIVATE_SUBSCRIPTION"
	RESUME_SUBSCRIPTION_SubscriptionEventSubscriptionEventType                        SubscriptionEventSubscriptionEventType = "RESUME_SUBSCRIPTION"
	PAUSE_SUBSCRIPTION_SubscriptionEventSubscriptionEventType                         SubscriptionEventSubscriptionEventType = "PAUSE_SUBSCRIPTION"
)
