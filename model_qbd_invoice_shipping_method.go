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

// checks if the QbdInvoiceShippingMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdInvoiceShippingMethod{}

// QbdInvoiceShippingMethod The shipping method used for this invoice, such as standard mail or overnight delivery.
type QbdInvoiceShippingMethod struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdInvoiceShippingMethod QbdInvoiceShippingMethod

// NewQbdInvoiceShippingMethod instantiates a new QbdInvoiceShippingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdInvoiceShippingMethod(id string, fullName string) *QbdInvoiceShippingMethod {
	this := QbdInvoiceShippingMethod{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdInvoiceShippingMethodWithDefaults instantiates a new QbdInvoiceShippingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdInvoiceShippingMethodWithDefaults() *QbdInvoiceShippingMethod {
	this := QbdInvoiceShippingMethod{}
	return &this
}

// GetId returns the Id field value
func (o *QbdInvoiceShippingMethod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdInvoiceShippingMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdInvoiceShippingMethod) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdInvoiceShippingMethod) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdInvoiceShippingMethod) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdInvoiceShippingMethod) SetFullName(v string) {
	o.FullName = v
}

func (o QbdInvoiceShippingMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdInvoiceShippingMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdInvoiceShippingMethod) UnmarshalJSON(data []byte) (err error) {
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

	varQbdInvoiceShippingMethod := _QbdInvoiceShippingMethod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdInvoiceShippingMethod)

	if err != nil {
		return err
	}

	*o = QbdInvoiceShippingMethod(varQbdInvoiceShippingMethod)

	return err
}

type NullableQbdInvoiceShippingMethod struct {
	value *QbdInvoiceShippingMethod
	isSet bool
}

func (v NullableQbdInvoiceShippingMethod) Get() *QbdInvoiceShippingMethod {
	return v.value
}

func (v *NullableQbdInvoiceShippingMethod) Set(val *QbdInvoiceShippingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdInvoiceShippingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdInvoiceShippingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdInvoiceShippingMethod(val *QbdInvoiceShippingMethod) *NullableQbdInvoiceShippingMethod {
	return &NullableQbdInvoiceShippingMethod{value: val, isSet: true}
}

func (v NullableQbdInvoiceShippingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdInvoiceShippingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


