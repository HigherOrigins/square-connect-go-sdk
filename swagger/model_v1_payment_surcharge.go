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

// V1PaymentSurcharge
type V1PaymentSurcharge struct {
	// The name of the surcharge.
	Name         string   `json:"name,omitempty"`
	AppliedMoney *V1Money `json:"applied_money,omitempty"`
	// The amount of the surcharge as a percentage. The percentage is provided as a string representing the decimal equivalent of the percentage. For example, \"0.7\" corresponds to a 7% surcharge. Exactly one of rate or amount_money should be set.
	Rate        string                  `json:"rate,omitempty"`
	AmountMoney *V1Money                `json:"amount_money,omitempty"`
	Type_       *V1PaymentSurchargeType `json:"type,omitempty"`
	// Indicates whether the surcharge is taxable.
	Taxable bool `json:"taxable,omitempty"`
	// The list of taxes that should be applied to the surcharge.
	Taxes []V1PaymentTax `json:"taxes,omitempty"`
	// A Square-issued unique identifier associated with the surcharge.
	SurchargeId string `json:"surcharge_id,omitempty"`
}
