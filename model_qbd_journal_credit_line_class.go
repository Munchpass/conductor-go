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

// checks if the QbdJournalCreditLineClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdJournalCreditLineClass{}

// QbdJournalCreditLineClass The journal credit line's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all journal credit lines unless overridden here, at the transaction line level.
type QbdJournalCreditLineClass struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdJournalCreditLineClass QbdJournalCreditLineClass

// NewQbdJournalCreditLineClass instantiates a new QbdJournalCreditLineClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdJournalCreditLineClass(id string, fullName string) *QbdJournalCreditLineClass {
	this := QbdJournalCreditLineClass{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdJournalCreditLineClassWithDefaults instantiates a new QbdJournalCreditLineClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdJournalCreditLineClassWithDefaults() *QbdJournalCreditLineClass {
	this := QbdJournalCreditLineClass{}
	return &this
}

// GetId returns the Id field value
func (o *QbdJournalCreditLineClass) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdJournalCreditLineClass) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdJournalCreditLineClass) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdJournalCreditLineClass) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdJournalCreditLineClass) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdJournalCreditLineClass) SetFullName(v string) {
	o.FullName = v
}

func (o QbdJournalCreditLineClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdJournalCreditLineClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdJournalCreditLineClass) UnmarshalJSON(data []byte) (err error) {
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

	varQbdJournalCreditLineClass := _QbdJournalCreditLineClass{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdJournalCreditLineClass)

	if err != nil {
		return err
	}

	*o = QbdJournalCreditLineClass(varQbdJournalCreditLineClass)

	return err
}

type NullableQbdJournalCreditLineClass struct {
	value *QbdJournalCreditLineClass
	isSet bool
}

func (v NullableQbdJournalCreditLineClass) Get() *QbdJournalCreditLineClass {
	return v.value
}

func (v *NullableQbdJournalCreditLineClass) Set(val *QbdJournalCreditLineClass) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdJournalCreditLineClass) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdJournalCreditLineClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdJournalCreditLineClass(val *QbdJournalCreditLineClass) *NullableQbdJournalCreditLineClass {
	return &NullableQbdJournalCreditLineClass{value: val, isSet: true}
}

func (v NullableQbdJournalCreditLineClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdJournalCreditLineClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


