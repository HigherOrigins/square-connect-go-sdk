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

// The set of line items, service charges, taxes, discounts, tips, etc. being returned in an Order.
type OrderReturn struct {
	// Unique ID that identifies the return only within this order.
	Uid string `json:"uid,omitempty"`
	// Order which contains the original sale of these returned line items. This will be unset for unlinked returns.
	SourceOrderId string `json:"source_order_id,omitempty"`
	// Collection of line items which are being returned.
	ReturnLineItems []OrderReturnLineItem `json:"return_line_items,omitempty"`
	// Collection of service charges which are being returned.
	ReturnServiceCharges []OrderReturnServiceCharge `json:"return_service_charges,omitempty"`
	// Collection of references to taxes being returned for an order, including the total applied tax amount to be returned. The taxes must reference a top-level tax ID from the source order.
	ReturnTaxes []OrderReturnTax `json:"return_taxes,omitempty"`
	// Collection of references to discounts being returned for an order, including the total applied discount amount to be returned. The discounts must reference a top-level discount ID from the source order.
	ReturnDiscounts    []OrderReturnDiscount    `json:"return_discounts,omitempty"`
	RoundingAdjustment *OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	ReturnAmounts      *OrderMoneyAmounts       `json:"return_amounts,omitempty"`
}
