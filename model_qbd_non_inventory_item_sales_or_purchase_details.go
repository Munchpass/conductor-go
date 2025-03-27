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

// QbdNonInventoryItemSalesOrPurchaseDetails - Details for non-inventory items that are exclusively sold or exclusively purchased, but not both. This typically applies to non-inventory items (like a purchased office supply that isn't resold) or service items (like consulting services that are sold but not purchased).  **IMPORTANT**: A non-inventory item will have either `salesAndPurchaseDetails` or `salesOrPurchaseDetails`, but never both because an item cannot have both configurations.
type QbdNonInventoryItemSalesOrPurchaseDetails struct {
	QbdSalesOrPurchaseDetails *QbdSalesOrPurchaseDetails
}

// QbdSalesOrPurchaseDetailsAsQbdNonInventoryItemSalesOrPurchaseDetails is a convenience function that returns QbdSalesOrPurchaseDetails wrapped in QbdNonInventoryItemSalesOrPurchaseDetails
func QbdSalesOrPurchaseDetailsAsQbdNonInventoryItemSalesOrPurchaseDetails(v *QbdSalesOrPurchaseDetails) QbdNonInventoryItemSalesOrPurchaseDetails {
	return QbdNonInventoryItemSalesOrPurchaseDetails{
		QbdSalesOrPurchaseDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdNonInventoryItemSalesOrPurchaseDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdSalesOrPurchaseDetails
	err = newStrictDecoder(data).Decode(&dst.QbdSalesOrPurchaseDetails)
	if err == nil {
		jsonQbdSalesOrPurchaseDetails, _ := json.Marshal(dst.QbdSalesOrPurchaseDetails)
		if string(jsonQbdSalesOrPurchaseDetails) == "{}" { // empty struct
			dst.QbdSalesOrPurchaseDetails = nil
		} else {
			if err = validator.Validate(dst.QbdSalesOrPurchaseDetails); err != nil {
				dst.QbdSalesOrPurchaseDetails = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdSalesOrPurchaseDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdSalesOrPurchaseDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdNonInventoryItemSalesOrPurchaseDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdNonInventoryItemSalesOrPurchaseDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdNonInventoryItemSalesOrPurchaseDetails) MarshalJSON() ([]byte, error) {
	if src.QbdSalesOrPurchaseDetails != nil {
		return json.Marshal(&src.QbdSalesOrPurchaseDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdNonInventoryItemSalesOrPurchaseDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdSalesOrPurchaseDetails != nil {
		return obj.QbdSalesOrPurchaseDetails
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdNonInventoryItemSalesOrPurchaseDetails) GetActualInstanceValue() (interface{}) {
	if obj.QbdSalesOrPurchaseDetails != nil {
		return *obj.QbdSalesOrPurchaseDetails
	}

	// all schemas are nil
	return nil
}

type NullableQbdNonInventoryItemSalesOrPurchaseDetails struct {
	value *QbdNonInventoryItemSalesOrPurchaseDetails
	isSet bool
}

func (v NullableQbdNonInventoryItemSalesOrPurchaseDetails) Get() *QbdNonInventoryItemSalesOrPurchaseDetails {
	return v.value
}

func (v *NullableQbdNonInventoryItemSalesOrPurchaseDetails) Set(val *QbdNonInventoryItemSalesOrPurchaseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdNonInventoryItemSalesOrPurchaseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdNonInventoryItemSalesOrPurchaseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdNonInventoryItemSalesOrPurchaseDetails(val *QbdNonInventoryItemSalesOrPurchaseDetails) *NullableQbdNonInventoryItemSalesOrPurchaseDetails {
	return &NullableQbdNonInventoryItemSalesOrPurchaseDetails{value: val, isSet: true}
}

func (v NullableQbdNonInventoryItemSalesOrPurchaseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdNonInventoryItemSalesOrPurchaseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


