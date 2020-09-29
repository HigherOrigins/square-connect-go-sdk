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

// Contains all information related to a single order to process with Square, including line items that specify the products to purchase. Order objects also include information on any associated tenders, refunds, and returns.  All Connect V2 Transactions have all been converted to Orders including all associated itemization data.
type Order struct {
	// The order's unique ID.
	Id string `json:"id,omitempty"`
	// The ID of the merchant location this order is associated with.
	LocationId string `json:"location_id"`
	// A client specified identifier to associate an entity in another system with this order.
	ReferenceId string       `json:"reference_id,omitempty"`
	Source      *OrderSource `json:"source,omitempty"`
	// The [Customer](#type-customer) ID of the customer associated with the order.
	CustomerId string `json:"customer_id,omitempty"`
	// The line items included in the order.
	LineItems []OrderLineItem `json:"line_items,omitempty"`
	// The list of all taxes associated with the order.  Taxes can be scoped to either `ORDER` or `LINE_ITEM`. For taxes with `LINE_ITEM` scope, an `OrderLineItemAppliedTax` must be added to each line item that the tax applies to. For taxes with `ORDER` scope, the server will generate an `OrderLineItemAppliedTax` for every line item.  On reads, each tax in the list will include the total amount of that tax applied to the order.  __IMPORTANT__: If `LINE_ITEM` scope is set on any taxes in this field, usage of the deprecated `line_items.taxes` field will result in an error. Please use `line_items.applied_taxes` instead.
	Taxes []OrderLineItemTax `json:"taxes,omitempty"`
	// The list of all discounts associated with the order.  Discounts can be scoped to either `ORDER` or `LINE_ITEM`. For discounts scoped to `LINE_ITEM`, an `OrderLineItemAppliedDiscount` must be added to each line item that the discount applies to. For discounts with `ORDER` scope, the server will generate an `OrderLineItemAppliedDiscount` for every line item.  __IMPORTANT__: If `LINE_ITEM` scope is set on any discounts in this field, usage of the deprecated `line_items.discounts` field will result in an error. Please use `line_items.applied_discounts` instead.
	Discounts []OrderLineItemDiscount `json:"discounts,omitempty"`
	// A list of service charges applied to the order.
	ServiceCharges []OrderServiceCharge `json:"service_charges,omitempty"`
	// Details on order fulfillment.  Orders can only be created with at most one fulfillment. However, orders returned by the API may contain multiple fulfillments.
	Fulfillments []OrderFulfillment `json:"fulfillments,omitempty"`
	// Collection of items from sale Orders being returned in this one. Normally part of an Itemized Return or Exchange.  There will be exactly one `Return` object per sale Order being referenced.
	Returns            []OrderReturn            `json:"returns,omitempty"`
	ReturnAmounts      *OrderMoneyAmounts       `json:"return_amounts,omitempty"`
	NetAmounts         *OrderMoneyAmounts       `json:"net_amounts,omitempty"`
	RoundingAdjustment *OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	// The Tenders which were used to pay for the Order.
	Tenders []Tender `json:"tenders,omitempty"`
	// The Refunds that are part of this Order.
	Refunds []Refund `json:"refunds,omitempty"`
	// Application-defined data attached to this order. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (personally identifiable information, card details, etc.).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries may also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a max length of 255 characters.  An application may have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  See [Metadata](https://developer.squareup.com/docs/build-basics/metadata) for more information.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Timestamp for when the order was created. In RFC 3339 format, e.g., \"2016-09-04T23:59:33.123Z\".
	CreatedAt string `json:"created_at,omitempty"`
	// Timestamp for when the order was last updated. In RFC 3339 format, e.g., \"2016-09-04T23:59:33.123Z\".
	UpdatedAt string `json:"updated_at,omitempty"`
	// Timestamp for when the order reached a terminal [state](#property-state). In RFC 3339 format, e.g., \"2016-09-04T23:59:33.123Z\".
	ClosedAt string      `json:"closed_at,omitempty"`
	State    *OrderState `json:"state,omitempty"`
	// Version number which is incremented each time an update is committed to the order. Orders that were not created through the API will not include a version and thus cannot be updated.  [Read more about working with versions](https://developer.squareup.com/docs/orders-api/manage-orders#update-orders).
	Version                 int32                `json:"version,omitempty"`
	TotalMoney              *Money               `json:"total_money,omitempty"`
	TotalTaxMoney           *Money               `json:"total_tax_money,omitempty"`
	TotalDiscountMoney      *Money               `json:"total_discount_money,omitempty"`
	TotalTipMoney           *Money               `json:"total_tip_money,omitempty"`
	TotalServiceChargeMoney *Money               `json:"total_service_charge_money,omitempty"`
	PricingOptions          *OrderPricingOptions `json:"pricing_options,omitempty"`
	// A set-like list of rewards that have been added to the order.
	Rewards []OrderReward `json:"rewards,omitempty"`
}
