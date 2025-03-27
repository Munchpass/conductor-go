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

// QbdPreferencesItemsAndInventory - The item inventory preferences for this company file.
type QbdPreferencesItemsAndInventory struct {
	QbdItemsAndInventoryPreferences *QbdItemsAndInventoryPreferences
}

// QbdItemsAndInventoryPreferencesAsQbdPreferencesItemsAndInventory is a convenience function that returns QbdItemsAndInventoryPreferences wrapped in QbdPreferencesItemsAndInventory
func QbdItemsAndInventoryPreferencesAsQbdPreferencesItemsAndInventory(v *QbdItemsAndInventoryPreferences) QbdPreferencesItemsAndInventory {
	return QbdPreferencesItemsAndInventory{
		QbdItemsAndInventoryPreferences: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdPreferencesItemsAndInventory) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdItemsAndInventoryPreferences
	err = newStrictDecoder(data).Decode(&dst.QbdItemsAndInventoryPreferences)
	if err == nil {
		jsonQbdItemsAndInventoryPreferences, _ := json.Marshal(dst.QbdItemsAndInventoryPreferences)
		if string(jsonQbdItemsAndInventoryPreferences) == "{}" { // empty struct
			dst.QbdItemsAndInventoryPreferences = nil
		} else {
			if err = validator.Validate(dst.QbdItemsAndInventoryPreferences); err != nil {
				dst.QbdItemsAndInventoryPreferences = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdItemsAndInventoryPreferences = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdItemsAndInventoryPreferences = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdPreferencesItemsAndInventory)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdPreferencesItemsAndInventory)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdPreferencesItemsAndInventory) MarshalJSON() ([]byte, error) {
	if src.QbdItemsAndInventoryPreferences != nil {
		return json.Marshal(&src.QbdItemsAndInventoryPreferences)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdPreferencesItemsAndInventory) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdItemsAndInventoryPreferences != nil {
		return obj.QbdItemsAndInventoryPreferences
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdPreferencesItemsAndInventory) GetActualInstanceValue() (interface{}) {
	if obj.QbdItemsAndInventoryPreferences != nil {
		return *obj.QbdItemsAndInventoryPreferences
	}

	// all schemas are nil
	return nil
}

type NullableQbdPreferencesItemsAndInventory struct {
	value *QbdPreferencesItemsAndInventory
	isSet bool
}

func (v NullableQbdPreferencesItemsAndInventory) Get() *QbdPreferencesItemsAndInventory {
	return v.value
}

func (v *NullableQbdPreferencesItemsAndInventory) Set(val *QbdPreferencesItemsAndInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdPreferencesItemsAndInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdPreferencesItemsAndInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdPreferencesItemsAndInventory(val *QbdPreferencesItemsAndInventory) *NullableQbdPreferencesItemsAndInventory {
	return &NullableQbdPreferencesItemsAndInventory{value: val, isSet: true}
}

func (v NullableQbdPreferencesItemsAndInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdPreferencesItemsAndInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


