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

// checks if the QbdPurchaseOrderLineSalesTaxCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdPurchaseOrderLineSalesTaxCode{}

// QbdPurchaseOrderLineSalesTaxCode The sales-tax code for this purchase order line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
type QbdPurchaseOrderLineSalesTaxCode struct {
	// The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types.
	Id string `json:"id"`
	// The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own `name`, separated by colons. Not case-sensitive.
	FullName string `json:"fullName"`
}

type _QbdPurchaseOrderLineSalesTaxCode QbdPurchaseOrderLineSalesTaxCode

// NewQbdPurchaseOrderLineSalesTaxCode instantiates a new QbdPurchaseOrderLineSalesTaxCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdPurchaseOrderLineSalesTaxCode(id string, fullName string) *QbdPurchaseOrderLineSalesTaxCode {
	this := QbdPurchaseOrderLineSalesTaxCode{}
	this.Id = id
	this.FullName = fullName
	return &this
}

// NewQbdPurchaseOrderLineSalesTaxCodeWithDefaults instantiates a new QbdPurchaseOrderLineSalesTaxCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdPurchaseOrderLineSalesTaxCodeWithDefaults() *QbdPurchaseOrderLineSalesTaxCode {
	this := QbdPurchaseOrderLineSalesTaxCode{}
	return &this
}

// GetId returns the Id field value
func (o *QbdPurchaseOrderLineSalesTaxCode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdPurchaseOrderLineSalesTaxCode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdPurchaseOrderLineSalesTaxCode) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *QbdPurchaseOrderLineSalesTaxCode) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *QbdPurchaseOrderLineSalesTaxCode) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *QbdPurchaseOrderLineSalesTaxCode) SetFullName(v string) {
	o.FullName = v
}

func (o QbdPurchaseOrderLineSalesTaxCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdPurchaseOrderLineSalesTaxCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *QbdPurchaseOrderLineSalesTaxCode) UnmarshalJSON(data []byte) (err error) {
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

	varQbdPurchaseOrderLineSalesTaxCode := _QbdPurchaseOrderLineSalesTaxCode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdPurchaseOrderLineSalesTaxCode)

	if err != nil {
		return err
	}

	*o = QbdPurchaseOrderLineSalesTaxCode(varQbdPurchaseOrderLineSalesTaxCode)

	return err
}

type NullableQbdPurchaseOrderLineSalesTaxCode struct {
	value *QbdPurchaseOrderLineSalesTaxCode
	isSet bool
}

func (v NullableQbdPurchaseOrderLineSalesTaxCode) Get() *QbdPurchaseOrderLineSalesTaxCode {
	return v.value
}

func (v *NullableQbdPurchaseOrderLineSalesTaxCode) Set(val *QbdPurchaseOrderLineSalesTaxCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdPurchaseOrderLineSalesTaxCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdPurchaseOrderLineSalesTaxCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdPurchaseOrderLineSalesTaxCode(val *QbdPurchaseOrderLineSalesTaxCode) *NullableQbdPurchaseOrderLineSalesTaxCode {
	return &NullableQbdPurchaseOrderLineSalesTaxCode{value: val, isSet: true}
}

func (v NullableQbdPurchaseOrderLineSalesTaxCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdPurchaseOrderLineSalesTaxCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


