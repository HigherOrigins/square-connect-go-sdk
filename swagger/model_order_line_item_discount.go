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

// Represents a discount that applies to one or more line items in an order.  Fixed-amount, order-scoped discounts are distributed across all non-zero line item totals. The amount distributed to each line item is relative to the amount contributed by the item to the order subtotal.
type OrderLineItemDiscount struct {
	// Unique ID that identifies the discount only within this order.
	Uid string `json:"uid,omitempty"`
	// The catalog object id referencing [CatalogDiscount](#type-catalogdiscount).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The discount's name.
	Name  string                     `json:"name,omitempty"`
	Type_ *OrderLineItemDiscountType `json:"type,omitempty"`
	// The percentage of the discount, as a string representation of a decimal number. A value of `7.25` corresponds to a percentage of 7.25%.  `percentage` is not set for amount-based discounts.
	Percentage   string `json:"percentage,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	// Application-defined data attached to this discount. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (personally identifiable information, card details, etc.).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries may also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a max length of 255 characters.  An application may have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  See [Metadata](https://developer.squareup.com/docs/build-basics/metadata) for more information.
	Metadata map[string]string           `json:"metadata,omitempty"`
	Scope    *OrderLineItemDiscountScope `json:"scope,omitempty"`
	// The reward identifiers corresponding to this discount. The application and specification of discounts that have `reward_ids` are completely controlled by the backing criteria corresponding to the reward tiers of the rewards that are added to the order through the Loyalty API. To manually unapply discounts that are the result of added rewards, the rewards must be removed from the order through the Loyalty API.
	RewardIds []string `json:"reward_ids,omitempty"`
	// The object identifier of a [pricing rule](#type-CatalogPricingRule) to be applied automatically to this discount. The specification and application of the discounts, to which a `pricing_rule_id` is assigned, are completely controlled by the corresponding pricing rule.
	PricingRuleId string `json:"pricing_rule_id,omitempty"`
}
