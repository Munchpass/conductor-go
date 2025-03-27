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

// checks if the QbdReceivePaymentCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdReceivePaymentCustomer{}

// QbdReceivePaymentCustomer The customer or customer-job to which the payment for this receive-payment is credited.
type QbdReceivePaymentCustomer struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdReceivePaymentCustomer QbdReceivePaymentCustomer

// NewQbdReceivePaymentCustomer instantiates a new QbdReceivePaymentCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdReceivePaymentCustomer(id string, fullName string) *QbdReceivePaymentCustomer {
	this := QbdReceivePaymentCustomer{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdReceivePaymentCustomerWithDefaults instantiates a new QbdReceivePaymentCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdReceivePaymentCustomerWithDefaults() *QbdReceivePaymentCustomer {
	this := QbdReceivePaymentCustomer{}
	return &this
}

// GetId returns the Id field value
func (o *QbdReceivePaymentCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdReceivePaymentCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdReceivePaymentCustomer) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdReceivePaymentCustomer) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdReceivePaymentCustomer) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdReceivePaymentCustomer) SetFullName(v string) {
	o.FullName = v
}

func (o QbdReceivePaymentCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdReceivePaymentCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdReceivePaymentCustomer) UnmarshalJSON(data []byte) (err error) {
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

	varQbdReceivePaymentCustomer := _QbdReceivePaymentCustomer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdReceivePaymentCustomer)

	if err != nil {
		return err
	}

	*o = QbdReceivePaymentCustomer(varQbdReceivePaymentCustomer)

	return err
}

type NullableQbdReceivePaymentCustomer struct {
	value *QbdReceivePaymentCustomer
	isSet bool
}

func (v NullableQbdReceivePaymentCustomer) Get() *QbdReceivePaymentCustomer {
	return v.value
}

func (v *NullableQbdReceivePaymentCustomer) Set(val *QbdReceivePaymentCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdReceivePaymentCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdReceivePaymentCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdReceivePaymentCustomer(val *QbdReceivePaymentCustomer) *NullableQbdReceivePaymentCustomer {
	return &NullableQbdReceivePaymentCustomer{value: val, isSet: true}
}

func (v NullableQbdReceivePaymentCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdReceivePaymentCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


