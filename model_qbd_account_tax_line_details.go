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

// QbdAccountTaxLineDetails - The account's tax line details, used for tax reporting purposes.
type QbdAccountTaxLineDetails struct {
	QbdTaxLineInfo *QbdTaxLineInfo
}

// QbdTaxLineInfoAsQbdAccountTaxLineDetails is a convenience function that returns QbdTaxLineInfo wrapped in QbdAccountTaxLineDetails
func QbdTaxLineInfoAsQbdAccountTaxLineDetails(v *QbdTaxLineInfo) QbdAccountTaxLineDetails {
	return QbdAccountTaxLineDetails{
		QbdTaxLineInfo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdAccountTaxLineDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdTaxLineInfo
	err = newStrictDecoder(data).Decode(&dst.QbdTaxLineInfo)
	if err == nil {
		jsonQbdTaxLineInfo, _ := json.Marshal(dst.QbdTaxLineInfo)
		if string(jsonQbdTaxLineInfo) == "{}" { // empty struct
			dst.QbdTaxLineInfo = nil
		} else {
			if err = validator.Validate(dst.QbdTaxLineInfo); err != nil {
				dst.QbdTaxLineInfo = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdTaxLineInfo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdTaxLineInfo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdAccountTaxLineDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdAccountTaxLineDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdAccountTaxLineDetails) MarshalJSON() ([]byte, error) {
	if src.QbdTaxLineInfo != nil {
		return json.Marshal(&src.QbdTaxLineInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdAccountTaxLineDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdTaxLineInfo != nil {
		return obj.QbdTaxLineInfo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdAccountTaxLineDetails) GetActualInstanceValue() (interface{}) {
	if obj.QbdTaxLineInfo != nil {
		return *obj.QbdTaxLineInfo
	}

	// all schemas are nil
	return nil
}

type NullableQbdAccountTaxLineDetails struct {
	value *QbdAccountTaxLineDetails
	isSet bool
}

func (v NullableQbdAccountTaxLineDetails) Get() *QbdAccountTaxLineDetails {
	return v.value
}

func (v *NullableQbdAccountTaxLineDetails) Set(val *QbdAccountTaxLineDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdAccountTaxLineDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdAccountTaxLineDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdAccountTaxLineDetails(val *QbdAccountTaxLineDetails) *NullableQbdAccountTaxLineDetails {
	return &NullableQbdAccountTaxLineDetails{value: val, isSet: true}
}

func (v NullableQbdAccountTaxLineDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdAccountTaxLineDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


