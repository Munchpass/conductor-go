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

// checks if the QbdVendorCreditVendor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdVendorCreditVendor{}

// QbdVendorCreditVendor The vendor who sent this vendor credit for goods or services purchased.
type QbdVendorCreditVendor struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdVendorCreditVendor QbdVendorCreditVendor

// NewQbdVendorCreditVendor instantiates a new QbdVendorCreditVendor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdVendorCreditVendor(id string, fullName string) *QbdVendorCreditVendor {
	this := QbdVendorCreditVendor{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdVendorCreditVendorWithDefaults instantiates a new QbdVendorCreditVendor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdVendorCreditVendorWithDefaults() *QbdVendorCreditVendor {
	this := QbdVendorCreditVendor{}
	return &this
}

// GetId returns the Id field value
func (o *QbdVendorCreditVendor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdVendorCreditVendor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdVendorCreditVendor) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdVendorCreditVendor) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdVendorCreditVendor) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdVendorCreditVendor) SetFullName(v string) {
	o.FullName = v
}

func (o QbdVendorCreditVendor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdVendorCreditVendor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdVendorCreditVendor) UnmarshalJSON(data []byte) (err error) {
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

	varQbdVendorCreditVendor := _QbdVendorCreditVendor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdVendorCreditVendor)

	if err != nil {
		return err
	}

	*o = QbdVendorCreditVendor(varQbdVendorCreditVendor)

	return err
}

type NullableQbdVendorCreditVendor struct {
	value *QbdVendorCreditVendor
	isSet bool
}

func (v NullableQbdVendorCreditVendor) Get() *QbdVendorCreditVendor {
	return v.value
}

func (v *NullableQbdVendorCreditVendor) Set(val *QbdVendorCreditVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdVendorCreditVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdVendorCreditVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdVendorCreditVendor(val *QbdVendorCreditVendor) *NullableQbdVendorCreditVendor {
	return &NullableQbdVendorCreditVendor{value: val, isSet: true}
}

func (v NullableQbdVendorCreditVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdVendorCreditVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


