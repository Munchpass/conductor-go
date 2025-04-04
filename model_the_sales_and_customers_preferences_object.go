/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TheSalesAndCustomersPreferencesObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TheSalesAndCustomersPreferencesObject{}

// TheSalesAndCustomersPreferencesObject The sales and customers preferences for this company file.
type TheSalesAndCustomersPreferencesObject struct {
	DefaultShippingMethod TheSalesAndCustomersPreferencesObjectDefaultShippingMethod `json:"defaultShippingMethod"`
	// The default shipment-origin location (i.e., FOB - freight on board) from which invoiced products are shipped for this company file. This indicates the point at which ownership and liability for goods transfer from seller to buyer.
	DefaultShipmentOrigin string `json:"defaultShipmentOrigin"`
	// The default percentage that an inventory item will be marked up from its cost for this company file.
	DefaultMarkupPercentage string `json:"defaultMarkupPercentage"`
	// Indicates whether this company file is configured to track an expense and the customer's reimbursement for that expense in separate accounts. When `true`, reimbursements can be tracked as income rather than as a reduction of the original expense.
	IsTrackingReimbursedExpensesAsIncome bool `json:"isTrackingReimbursedExpensesAsIncome"`
	// Indicates whether this company file is configured to automatically apply a customer's payment to their outstanding invoices, beginning with the oldest invoice.
	IsAutoApplyingPayments bool `json:"isAutoApplyingPayments"`
	PriceLevels TheSalesAndCustomersPreferencesObjectPriceLevels `json:"priceLevels"`
}

type _TheSalesAndCustomersPreferencesObject TheSalesAndCustomersPreferencesObject

// NewTheSalesAndCustomersPreferencesObject instantiates a new TheSalesAndCustomersPreferencesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTheSalesAndCustomersPreferencesObject(defaultShippingMethod TheSalesAndCustomersPreferencesObjectDefaultShippingMethod, defaultShipmentOrigin string, defaultMarkupPercentage string, isTrackingReimbursedExpensesAsIncome bool, isAutoApplyingPayments bool, priceLevels TheSalesAndCustomersPreferencesObjectPriceLevels) *TheSalesAndCustomersPreferencesObject {
	this := TheSalesAndCustomersPreferencesObject{}
	this.DefaultShippingMethod = defaultShippingMethod
	this.DefaultShipmentOrigin = defaultShipmentOrigin
	this.DefaultMarkupPercentage = defaultMarkupPercentage
	this.IsTrackingReimbursedExpensesAsIncome = isTrackingReimbursedExpensesAsIncome
	this.IsAutoApplyingPayments = isAutoApplyingPayments
	this.PriceLevels = priceLevels
	return &this
}

// NewTheSalesAndCustomersPreferencesObjectWithDefaults instantiates a new TheSalesAndCustomersPreferencesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTheSalesAndCustomersPreferencesObjectWithDefaults() *TheSalesAndCustomersPreferencesObject {
	this := TheSalesAndCustomersPreferencesObject{}
	return &this
}

// GetDefaultShippingMethod returns the DefaultShippingMethod field value
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShippingMethod() TheSalesAndCustomersPreferencesObjectDefaultShippingMethod {
	if o == nil {
		var ret TheSalesAndCustomersPreferencesObjectDefaultShippingMethod
		return ret
	}

	return o.DefaultShippingMethod
}

// GetDefaultShippingMethodOk returns a tuple with the DefaultShippingMethod field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShippingMethodOk() (*TheSalesAndCustomersPreferencesObjectDefaultShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultShippingMethod, true
}

// SetDefaultShippingMethod sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetDefaultShippingMethod(v TheSalesAndCustomersPreferencesObjectDefaultShippingMethod) {
	o.DefaultShippingMethod = v
}

// GetDefaultShipmentOrigin returns the DefaultShipmentOrigin field value
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShipmentOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultShipmentOrigin
}

// GetDefaultShipmentOriginOk returns a tuple with the DefaultShipmentOrigin field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShipmentOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultShipmentOrigin, true
}

// SetDefaultShipmentOrigin sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetDefaultShipmentOrigin(v string) {
	o.DefaultShipmentOrigin = v
}

// GetDefaultMarkupPercentage returns the DefaultMarkupPercentage field value
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultMarkupPercentage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultMarkupPercentage
}

// GetDefaultMarkupPercentageOk returns a tuple with the DefaultMarkupPercentage field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetDefaultMarkupPercentageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMarkupPercentage, true
}

// SetDefaultMarkupPercentage sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetDefaultMarkupPercentage(v string) {
	o.DefaultMarkupPercentage = v
}

// GetIsTrackingReimbursedExpensesAsIncome returns the IsTrackingReimbursedExpensesAsIncome field value
func (o *TheSalesAndCustomersPreferencesObject) GetIsTrackingReimbursedExpensesAsIncome() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTrackingReimbursedExpensesAsIncome
}

// GetIsTrackingReimbursedExpensesAsIncomeOk returns a tuple with the IsTrackingReimbursedExpensesAsIncome field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetIsTrackingReimbursedExpensesAsIncomeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTrackingReimbursedExpensesAsIncome, true
}

// SetIsTrackingReimbursedExpensesAsIncome sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetIsTrackingReimbursedExpensesAsIncome(v bool) {
	o.IsTrackingReimbursedExpensesAsIncome = v
}

// GetIsAutoApplyingPayments returns the IsAutoApplyingPayments field value
func (o *TheSalesAndCustomersPreferencesObject) GetIsAutoApplyingPayments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoApplyingPayments
}

// GetIsAutoApplyingPaymentsOk returns a tuple with the IsAutoApplyingPayments field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetIsAutoApplyingPaymentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoApplyingPayments, true
}

// SetIsAutoApplyingPayments sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetIsAutoApplyingPayments(v bool) {
	o.IsAutoApplyingPayments = v
}

// GetPriceLevels returns the PriceLevels field value
func (o *TheSalesAndCustomersPreferencesObject) GetPriceLevels() TheSalesAndCustomersPreferencesObjectPriceLevels {
	if o == nil {
		var ret TheSalesAndCustomersPreferencesObjectPriceLevels
		return ret
	}

	return o.PriceLevels
}

// GetPriceLevelsOk returns a tuple with the PriceLevels field value
// and a boolean to check if the value has been set.
func (o *TheSalesAndCustomersPreferencesObject) GetPriceLevelsOk() (*TheSalesAndCustomersPreferencesObjectPriceLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceLevels, true
}

// SetPriceLevels sets field value
func (o *TheSalesAndCustomersPreferencesObject) SetPriceLevels(v TheSalesAndCustomersPreferencesObjectPriceLevels) {
	o.PriceLevels = v
}

func (o TheSalesAndCustomersPreferencesObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TheSalesAndCustomersPreferencesObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultShippingMethod"] = o.DefaultShippingMethod
	toSerialize["defaultShipmentOrigin"] = o.DefaultShipmentOrigin
	toSerialize["defaultMarkupPercentage"] = o.DefaultMarkupPercentage
	toSerialize["isTrackingReimbursedExpensesAsIncome"] = o.IsTrackingReimbursedExpensesAsIncome
	toSerialize["isAutoApplyingPayments"] = o.IsAutoApplyingPayments
	toSerialize["priceLevels"] = o.PriceLevels
	return toSerialize, nil
}

func (o *TheSalesAndCustomersPreferencesObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultShippingMethod",
		"defaultShipmentOrigin",
		"defaultMarkupPercentage",
		"isTrackingReimbursedExpensesAsIncome",
		"isAutoApplyingPayments",
		"priceLevels",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTheSalesAndCustomersPreferencesObject := _TheSalesAndCustomersPreferencesObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTheSalesAndCustomersPreferencesObject)

	if err != nil {
		return err
	}

	*o = TheSalesAndCustomersPreferencesObject(varTheSalesAndCustomersPreferencesObject)

	return err
}

type NullableTheSalesAndCustomersPreferencesObject struct {
	value *TheSalesAndCustomersPreferencesObject
	isSet bool
}

func (v NullableTheSalesAndCustomersPreferencesObject) Get() *TheSalesAndCustomersPreferencesObject {
	return v.value
}

func (v *NullableTheSalesAndCustomersPreferencesObject) Set(val *TheSalesAndCustomersPreferencesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTheSalesAndCustomersPreferencesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTheSalesAndCustomersPreferencesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTheSalesAndCustomersPreferencesObject(val *TheSalesAndCustomersPreferencesObject) *NullableTheSalesAndCustomersPreferencesObject {
	return &NullableTheSalesAndCustomersPreferencesObject{value: val, isSet: true}
}

func (v NullableTheSalesAndCustomersPreferencesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTheSalesAndCustomersPreferencesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


