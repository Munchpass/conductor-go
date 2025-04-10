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

// QbdPreferencesMultiLocationInventory - The multi-location inventory preferences for this company file.
type QbdPreferencesMultiLocationInventory struct {
	QbdMultiLocationInventoryPreferences *QbdMultiLocationInventoryPreferences
}

// QbdMultiLocationInventoryPreferencesAsQbdPreferencesMultiLocationInventory is a convenience function that returns QbdMultiLocationInventoryPreferences wrapped in QbdPreferencesMultiLocationInventory
func QbdMultiLocationInventoryPreferencesAsQbdPreferencesMultiLocationInventory(v *QbdMultiLocationInventoryPreferences) QbdPreferencesMultiLocationInventory {
	return QbdPreferencesMultiLocationInventory{
		QbdMultiLocationInventoryPreferences: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdPreferencesMultiLocationInventory) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdMultiLocationInventoryPreferences
	err = newStrictDecoder(data).Decode(&dst.QbdMultiLocationInventoryPreferences)
	if err == nil {
		jsonQbdMultiLocationInventoryPreferences, _ := json.Marshal(dst.QbdMultiLocationInventoryPreferences)
		if string(jsonQbdMultiLocationInventoryPreferences) == "{}" { // empty struct
			dst.QbdMultiLocationInventoryPreferences = nil
		} else {
			if err = validator.Validate(dst.QbdMultiLocationInventoryPreferences); err != nil {
				dst.QbdMultiLocationInventoryPreferences = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdMultiLocationInventoryPreferences = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdMultiLocationInventoryPreferences = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdPreferencesMultiLocationInventory)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdPreferencesMultiLocationInventory)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdPreferencesMultiLocationInventory) MarshalJSON() ([]byte, error) {
	if src.QbdMultiLocationInventoryPreferences != nil {
		return json.Marshal(&src.QbdMultiLocationInventoryPreferences)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdPreferencesMultiLocationInventory) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdMultiLocationInventoryPreferences != nil {
		return obj.QbdMultiLocationInventoryPreferences
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdPreferencesMultiLocationInventory) GetActualInstanceValue() (interface{}) {
	if obj.QbdMultiLocationInventoryPreferences != nil {
		return *obj.QbdMultiLocationInventoryPreferences
	}

	// all schemas are nil
	return nil
}

type NullableQbdPreferencesMultiLocationInventory struct {
	value *QbdPreferencesMultiLocationInventory
	isSet bool
}

func (v NullableQbdPreferencesMultiLocationInventory) Get() *QbdPreferencesMultiLocationInventory {
	return v.value
}

func (v *NullableQbdPreferencesMultiLocationInventory) Set(val *QbdPreferencesMultiLocationInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdPreferencesMultiLocationInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdPreferencesMultiLocationInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdPreferencesMultiLocationInventory(val *QbdPreferencesMultiLocationInventory) *NullableQbdPreferencesMultiLocationInventory {
	return &NullableQbdPreferencesMultiLocationInventory{value: val, isSet: true}
}

func (v NullableQbdPreferencesMultiLocationInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdPreferencesMultiLocationInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


