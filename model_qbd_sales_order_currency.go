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

// checks if the QbdSalesOrderCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdSalesOrderCurrency{}

// QbdSalesOrderCurrency The sales order's currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable.
type QbdSalesOrderCurrency struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdSalesOrderCurrency QbdSalesOrderCurrency

// NewQbdSalesOrderCurrency instantiates a new QbdSalesOrderCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdSalesOrderCurrency(id string, fullName string) *QbdSalesOrderCurrency {
	this := QbdSalesOrderCurrency{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdSalesOrderCurrencyWithDefaults instantiates a new QbdSalesOrderCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdSalesOrderCurrencyWithDefaults() *QbdSalesOrderCurrency {
	this := QbdSalesOrderCurrency{}
	return &this
}

// GetId returns the Id field value
func (o *QbdSalesOrderCurrency) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderCurrency) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdSalesOrderCurrency) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdSalesOrderCurrency) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderCurrency) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdSalesOrderCurrency) SetFullName(v string) {
	o.FullName = v
}

func (o QbdSalesOrderCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdSalesOrderCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdSalesOrderCurrency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"fullName",
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

	varQbdSalesOrderCurrency := _QbdSalesOrderCurrency{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdSalesOrderCurrency)

	if err != nil {
		return err
	}

	*o = QbdSalesOrderCurrency(varQbdSalesOrderCurrency)

	return err
}

type NullableQbdSalesOrderCurrency struct {
	value *QbdSalesOrderCurrency
	isSet bool
}

func (v NullableQbdSalesOrderCurrency) Get() *QbdSalesOrderCurrency {
	return v.value
}

func (v *NullableQbdSalesOrderCurrency) Set(val *QbdSalesOrderCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdSalesOrderCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdSalesOrderCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdSalesOrderCurrency(val *QbdSalesOrderCurrency) *NullableQbdSalesOrderCurrency {
	return &NullableQbdSalesOrderCurrency{value: val, isSet: true}
}

func (v NullableQbdSalesOrderCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdSalesOrderCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


