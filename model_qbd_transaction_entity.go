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

// checks if the QbdTransactionEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdTransactionEntity{}

// QbdTransactionEntity The customer, vendor, employee, or person on QuickBooks's \"Other Names\" list associated with this transaction.
type QbdTransactionEntity struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdTransactionEntity QbdTransactionEntity

// NewQbdTransactionEntity instantiates a new QbdTransactionEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdTransactionEntity(id string, fullName string) *QbdTransactionEntity {
	this := QbdTransactionEntity{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdTransactionEntityWithDefaults instantiates a new QbdTransactionEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdTransactionEntityWithDefaults() *QbdTransactionEntity {
	this := QbdTransactionEntity{}
	return &this
}

// GetId returns the Id field value
func (o *QbdTransactionEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdTransactionEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdTransactionEntity) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdTransactionEntity) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdTransactionEntity) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdTransactionEntity) SetFullName(v string) {
	o.FullName = v
}

func (o QbdTransactionEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdTransactionEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdTransactionEntity) UnmarshalJSON(data []byte) (err error) {
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

	varQbdTransactionEntity := _QbdTransactionEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdTransactionEntity)

	if err != nil {
		return err
	}

	*o = QbdTransactionEntity(varQbdTransactionEntity)

	return err
}

type NullableQbdTransactionEntity struct {
	value *QbdTransactionEntity
	isSet bool
}

func (v NullableQbdTransactionEntity) Get() *QbdTransactionEntity {
	return v.value
}

func (v *NullableQbdTransactionEntity) Set(val *QbdTransactionEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdTransactionEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdTransactionEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdTransactionEntity(val *QbdTransactionEntity) *NullableQbdTransactionEntity {
	return &NullableQbdTransactionEntity{value: val, isSet: true}
}

func (v NullableQbdTransactionEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdTransactionEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


