/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// QbdSalesOrderShippingAddress - The sales order's shipping address.
type QbdSalesOrderShippingAddress struct {
	QbdAddress *QbdAddress
}

// QbdAddressAsQbdSalesOrderShippingAddress is a convenience function that returns QbdAddress wrapped in QbdSalesOrderShippingAddress
func QbdAddressAsQbdSalesOrderShippingAddress(v *QbdAddress) QbdSalesOrderShippingAddress {
	return QbdSalesOrderShippingAddress{
		QbdAddress: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdSalesOrderShippingAddress) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdAddress
	err = newStrictDecoder(data).Decode(&dst.QbdAddress)
	if err == nil {
		jsonQbdAddress, _ := json.Marshal(dst.QbdAddress)
		if string(jsonQbdAddress) == "{}" { // empty struct
			dst.QbdAddress = nil
		} else {
			if err = validator.Validate(dst.QbdAddress); err != nil {
				dst.QbdAddress = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdAddress = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdAddress = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdSalesOrderShippingAddress)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdSalesOrderShippingAddress)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdSalesOrderShippingAddress) MarshalJSON() ([]byte, error) {
	if src.QbdAddress != nil {
		return json.Marshal(&src.QbdAddress)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdSalesOrderShippingAddress) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdAddress != nil {
		return obj.QbdAddress
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdSalesOrderShippingAddress) GetActualInstanceValue() (interface{}) {
	if obj.QbdAddress != nil {
		return *obj.QbdAddress
	}

	// all schemas are nil
	return nil
}

type NullableQbdSalesOrderShippingAddress struct {
	value *QbdSalesOrderShippingAddress
	isSet bool
}

func (v NullableQbdSalesOrderShippingAddress) Get() *QbdSalesOrderShippingAddress {
	return v.value
}

func (v *NullableQbdSalesOrderShippingAddress) Set(val *QbdSalesOrderShippingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdSalesOrderShippingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdSalesOrderShippingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdSalesOrderShippingAddress(val *QbdSalesOrderShippingAddress) *NullableQbdSalesOrderShippingAddress {
	return &NullableQbdSalesOrderShippingAddress{value: val, isSet: true}
}

func (v NullableQbdSalesOrderShippingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdSalesOrderShippingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


