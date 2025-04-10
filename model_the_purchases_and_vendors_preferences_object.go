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

// checks if the ThePurchasesAndVendorsPreferencesObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThePurchasesAndVendorsPreferencesObject{}

// ThePurchasesAndVendorsPreferencesObject The purchases and vendors preferences for this company file.
type ThePurchasesAndVendorsPreferencesObject struct {
	// Indicates whether this company file has inventory-related features enabled.
	IsUsingInventory bool `json:"isUsingInventory"`
	// The default number of days after receipt when bills are due for this company file.
	DaysBillsAreDue float32 `json:"daysBillsAreDue"`
	// Indicates whether this company file is configured to automatically apply available vendor discounts or credits when paying bills.
	IsAutomaticallyUsingDiscounts bool `json:"isAutomaticallyUsingDiscounts"`
	DefaultDiscountAccount ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount `json:"defaultDiscountAccount"`
}

type _ThePurchasesAndVendorsPreferencesObject ThePurchasesAndVendorsPreferencesObject

// NewThePurchasesAndVendorsPreferencesObject instantiates a new ThePurchasesAndVendorsPreferencesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThePurchasesAndVendorsPreferencesObject(isUsingInventory bool, daysBillsAreDue float32, isAutomaticallyUsingDiscounts bool, defaultDiscountAccount ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount) *ThePurchasesAndVendorsPreferencesObject {
	this := ThePurchasesAndVendorsPreferencesObject{}
	this.IsUsingInventory = isUsingInventory
	this.DaysBillsAreDue = daysBillsAreDue
	this.IsAutomaticallyUsingDiscounts = isAutomaticallyUsingDiscounts
	this.DefaultDiscountAccount = defaultDiscountAccount
	return &this
}

// NewThePurchasesAndVendorsPreferencesObjectWithDefaults instantiates a new ThePurchasesAndVendorsPreferencesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThePurchasesAndVendorsPreferencesObjectWithDefaults() *ThePurchasesAndVendorsPreferencesObject {
	this := ThePurchasesAndVendorsPreferencesObject{}
	return &this
}

// GetIsUsingInventory returns the IsUsingInventory field value
func (o *ThePurchasesAndVendorsPreferencesObject) GetIsUsingInventory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsingInventory
}

// GetIsUsingInventoryOk returns a tuple with the IsUsingInventory field value
// and a boolean to check if the value has been set.
func (o *ThePurchasesAndVendorsPreferencesObject) GetIsUsingInventoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsingInventory, true
}

// SetIsUsingInventory sets field value
func (o *ThePurchasesAndVendorsPreferencesObject) SetIsUsingInventory(v bool) {
	o.IsUsingInventory = v
}

// GetDaysBillsAreDue returns the DaysBillsAreDue field value
func (o *ThePurchasesAndVendorsPreferencesObject) GetDaysBillsAreDue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysBillsAreDue
}

// GetDaysBillsAreDueOk returns a tuple with the DaysBillsAreDue field value
// and a boolean to check if the value has been set.
func (o *ThePurchasesAndVendorsPreferencesObject) GetDaysBillsAreDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysBillsAreDue, true
}

// SetDaysBillsAreDue sets field value
func (o *ThePurchasesAndVendorsPreferencesObject) SetDaysBillsAreDue(v float32) {
	o.DaysBillsAreDue = v
}

// GetIsAutomaticallyUsingDiscounts returns the IsAutomaticallyUsingDiscounts field value
func (o *ThePurchasesAndVendorsPreferencesObject) GetIsAutomaticallyUsingDiscounts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomaticallyUsingDiscounts
}

// GetIsAutomaticallyUsingDiscountsOk returns a tuple with the IsAutomaticallyUsingDiscounts field value
// and a boolean to check if the value has been set.
func (o *ThePurchasesAndVendorsPreferencesObject) GetIsAutomaticallyUsingDiscountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomaticallyUsingDiscounts, true
}

// SetIsAutomaticallyUsingDiscounts sets field value
func (o *ThePurchasesAndVendorsPreferencesObject) SetIsAutomaticallyUsingDiscounts(v bool) {
	o.IsAutomaticallyUsingDiscounts = v
}

// GetDefaultDiscountAccount returns the DefaultDiscountAccount field value
func (o *ThePurchasesAndVendorsPreferencesObject) GetDefaultDiscountAccount() ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount {
	if o == nil {
		var ret ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount
		return ret
	}

	return o.DefaultDiscountAccount
}

// GetDefaultDiscountAccountOk returns a tuple with the DefaultDiscountAccount field value
// and a boolean to check if the value has been set.
func (o *ThePurchasesAndVendorsPreferencesObject) GetDefaultDiscountAccountOk() (*ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDiscountAccount, true
}

// SetDefaultDiscountAccount sets field value
func (o *ThePurchasesAndVendorsPreferencesObject) SetDefaultDiscountAccount(v ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount) {
	o.DefaultDiscountAccount = v
}

func (o ThePurchasesAndVendorsPreferencesObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThePurchasesAndVendorsPreferencesObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isUsingInventory"] = o.IsUsingInventory
	toSerialize["daysBillsAreDue"] = o.DaysBillsAreDue
	toSerialize["isAutomaticallyUsingDiscounts"] = o.IsAutomaticallyUsingDiscounts
	toSerialize["defaultDiscountAccount"] = o.DefaultDiscountAccount
	return toSerialize, nil
}

func (o *ThePurchasesAndVendorsPreferencesObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isUsingInventory",
		"daysBillsAreDue",
		"isAutomaticallyUsingDiscounts",
		"defaultDiscountAccount",
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

	varThePurchasesAndVendorsPreferencesObject := _ThePurchasesAndVendorsPreferencesObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThePurchasesAndVendorsPreferencesObject)

	if err != nil {
		return err
	}

	*o = ThePurchasesAndVendorsPreferencesObject(varThePurchasesAndVendorsPreferencesObject)

	return err
}

type NullableThePurchasesAndVendorsPreferencesObject struct {
	value *ThePurchasesAndVendorsPreferencesObject
	isSet bool
}

func (v NullableThePurchasesAndVendorsPreferencesObject) Get() *ThePurchasesAndVendorsPreferencesObject {
	return v.value
}

func (v *NullableThePurchasesAndVendorsPreferencesObject) Set(val *ThePurchasesAndVendorsPreferencesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableThePurchasesAndVendorsPreferencesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableThePurchasesAndVendorsPreferencesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThePurchasesAndVendorsPreferencesObject(val *ThePurchasesAndVendorsPreferencesObject) *NullableThePurchasesAndVendorsPreferencesObject {
	return &NullableThePurchasesAndVendorsPreferencesObject{value: val, isSet: true}
}

func (v NullableThePurchasesAndVendorsPreferencesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThePurchasesAndVendorsPreferencesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


