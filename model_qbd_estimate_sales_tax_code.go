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

// checks if the QbdEstimateSalesTaxCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdEstimateSalesTaxCode{}

// QbdEstimateSalesTaxCode The sales-tax code for this estimate, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
type QbdEstimateSalesTaxCode struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdEstimateSalesTaxCode QbdEstimateSalesTaxCode

// NewQbdEstimateSalesTaxCode instantiates a new QbdEstimateSalesTaxCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdEstimateSalesTaxCode(id string, fullName string) *QbdEstimateSalesTaxCode {
	this := QbdEstimateSalesTaxCode{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdEstimateSalesTaxCodeWithDefaults instantiates a new QbdEstimateSalesTaxCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdEstimateSalesTaxCodeWithDefaults() *QbdEstimateSalesTaxCode {
	this := QbdEstimateSalesTaxCode{}
	return &this
}

// GetId returns the Id field value
func (o *QbdEstimateSalesTaxCode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdEstimateSalesTaxCode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdEstimateSalesTaxCode) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdEstimateSalesTaxCode) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdEstimateSalesTaxCode) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdEstimateSalesTaxCode) SetFullName(v string) {
	o.FullName = v
}

func (o QbdEstimateSalesTaxCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdEstimateSalesTaxCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdEstimateSalesTaxCode) UnmarshalJSON(data []byte) (err error) {
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

	varQbdEstimateSalesTaxCode := _QbdEstimateSalesTaxCode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdEstimateSalesTaxCode)

	if err != nil {
		return err
	}

	*o = QbdEstimateSalesTaxCode(varQbdEstimateSalesTaxCode)

	return err
}

type NullableQbdEstimateSalesTaxCode struct {
	value *QbdEstimateSalesTaxCode
	isSet bool
}

func (v NullableQbdEstimateSalesTaxCode) Get() *QbdEstimateSalesTaxCode {
	return v.value
}

func (v *NullableQbdEstimateSalesTaxCode) Set(val *QbdEstimateSalesTaxCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdEstimateSalesTaxCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdEstimateSalesTaxCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdEstimateSalesTaxCode(val *QbdEstimateSalesTaxCode) *NullableQbdEstimateSalesTaxCode {
	return &NullableQbdEstimateSalesTaxCode{value: val, isSet: true}
}

func (v NullableQbdEstimateSalesTaxCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdEstimateSalesTaxCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


