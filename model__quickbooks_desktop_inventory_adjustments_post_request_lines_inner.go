/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
)

// checks if the QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner{}

// QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner struct for QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner
type QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner struct {
	// The inventory item associated with this inventory adjustment line.
	ItemId *string `json:"itemId,omitempty"`
	AdjustQuantity *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity `json:"adjustQuantity,omitempty"`
	AdjustValue *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue `json:"adjustValue,omitempty"`
	AdjustSerialNumber *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber `json:"adjustSerialNumber,omitempty"`
	AdjustLotNumber *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber `json:"adjustLotNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner

// NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner {
	this := QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner{}
	return &this
}

// NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner {
	this := QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) SetItemId(v string) {
	o.ItemId = &v
}

// GetAdjustQuantity returns the AdjustQuantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustQuantity() QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity {
	if o == nil || IsNil(o.AdjustQuantity) {
		var ret QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity
		return ret
	}
	return *o.AdjustQuantity
}

// GetAdjustQuantityOk returns a tuple with the AdjustQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustQuantityOk() (*QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity, bool) {
	if o == nil || IsNil(o.AdjustQuantity) {
		return nil, false
	}
	return o.AdjustQuantity, true
}

// HasAdjustQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) HasAdjustQuantity() bool {
	if o != nil && !IsNil(o.AdjustQuantity) {
		return true
	}

	return false
}

// SetAdjustQuantity gets a reference to the given QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity and assigns it to the AdjustQuantity field.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) SetAdjustQuantity(v QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) {
	o.AdjustQuantity = &v
}

// GetAdjustValue returns the AdjustValue field value if set, zero value otherwise.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustValue() QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue {
	if o == nil || IsNil(o.AdjustValue) {
		var ret QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue
		return ret
	}
	return *o.AdjustValue
}

// GetAdjustValueOk returns a tuple with the AdjustValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustValueOk() (*QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue, bool) {
	if o == nil || IsNil(o.AdjustValue) {
		return nil, false
	}
	return o.AdjustValue, true
}

// HasAdjustValue returns a boolean if a field has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) HasAdjustValue() bool {
	if o != nil && !IsNil(o.AdjustValue) {
		return true
	}

	return false
}

// SetAdjustValue gets a reference to the given QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue and assigns it to the AdjustValue field.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) SetAdjustValue(v QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) {
	o.AdjustValue = &v
}

// GetAdjustSerialNumber returns the AdjustSerialNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustSerialNumber() QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber {
	if o == nil || IsNil(o.AdjustSerialNumber) {
		var ret QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber
		return ret
	}
	return *o.AdjustSerialNumber
}

// GetAdjustSerialNumberOk returns a tuple with the AdjustSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustSerialNumberOk() (*QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber, bool) {
	if o == nil || IsNil(o.AdjustSerialNumber) {
		return nil, false
	}
	return o.AdjustSerialNumber, true
}

// HasAdjustSerialNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) HasAdjustSerialNumber() bool {
	if o != nil && !IsNil(o.AdjustSerialNumber) {
		return true
	}

	return false
}

// SetAdjustSerialNumber gets a reference to the given QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber and assigns it to the AdjustSerialNumber field.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) SetAdjustSerialNumber(v QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) {
	o.AdjustSerialNumber = &v
}

// GetAdjustLotNumber returns the AdjustLotNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustLotNumber() QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber {
	if o == nil || IsNil(o.AdjustLotNumber) {
		var ret QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber
		return ret
	}
	return *o.AdjustLotNumber
}

// GetAdjustLotNumberOk returns a tuple with the AdjustLotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) GetAdjustLotNumberOk() (*QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber, bool) {
	if o == nil || IsNil(o.AdjustLotNumber) {
		return nil, false
	}
	return o.AdjustLotNumber, true
}

// HasAdjustLotNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) HasAdjustLotNumber() bool {
	if o != nil && !IsNil(o.AdjustLotNumber) {
		return true
	}

	return false
}

// SetAdjustLotNumber gets a reference to the given QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber and assigns it to the AdjustLotNumber field.
func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) SetAdjustLotNumber(v QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) {
	o.AdjustLotNumber = &v
}

func (o QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["itemId"] = o.ItemId
	}
	if !IsNil(o.AdjustQuantity) {
		toSerialize["adjustQuantity"] = o.AdjustQuantity
	}
	if !IsNil(o.AdjustValue) {
		toSerialize["adjustValue"] = o.AdjustValue
	}
	if !IsNil(o.AdjustSerialNumber) {
		toSerialize["adjustSerialNumber"] = o.AdjustSerialNumber
	}
	if !IsNil(o.AdjustLotNumber) {
		toSerialize["adjustLotNumber"] = o.AdjustLotNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner := _QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner(varQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemId")
		delete(additionalProperties, "adjustQuantity")
		delete(additionalProperties, "adjustValue")
		delete(additionalProperties, "adjustSerialNumber")
		delete(additionalProperties, "adjustLotNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner struct {
	value *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner
	isSet bool
}

func (v NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) Get() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) Set(val *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner(val *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) *NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner {
	return &NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


