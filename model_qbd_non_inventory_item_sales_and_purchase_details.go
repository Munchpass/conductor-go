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

// QbdNonInventoryItemSalesAndPurchaseDetails - Details for non-inventory items that are both purchased and sold, such as reimbursable expenses or inventory items that are bought from vendors and sold to customers.  **IMPORTANT**: A non-inventory item will have either `salesAndPurchaseDetails` or `salesOrPurchaseDetails`, but never both because an item cannot have both configurations.
type QbdNonInventoryItemSalesAndPurchaseDetails struct {
	QbdSalesAndPurchaseDetails *QbdSalesAndPurchaseDetails
}

// QbdSalesAndPurchaseDetailsAsQbdNonInventoryItemSalesAndPurchaseDetails is a convenience function that returns QbdSalesAndPurchaseDetails wrapped in QbdNonInventoryItemSalesAndPurchaseDetails
func QbdSalesAndPurchaseDetailsAsQbdNonInventoryItemSalesAndPurchaseDetails(v *QbdSalesAndPurchaseDetails) QbdNonInventoryItemSalesAndPurchaseDetails {
	return QbdNonInventoryItemSalesAndPurchaseDetails{
		QbdSalesAndPurchaseDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdNonInventoryItemSalesAndPurchaseDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdSalesAndPurchaseDetails
	err = newStrictDecoder(data).Decode(&dst.QbdSalesAndPurchaseDetails)
	if err == nil {
		jsonQbdSalesAndPurchaseDetails, _ := json.Marshal(dst.QbdSalesAndPurchaseDetails)
		if string(jsonQbdSalesAndPurchaseDetails) == "{}" { // empty struct
			dst.QbdSalesAndPurchaseDetails = nil
		} else {
			if err = validator.Validate(dst.QbdSalesAndPurchaseDetails); err != nil {
				dst.QbdSalesAndPurchaseDetails = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdSalesAndPurchaseDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdSalesAndPurchaseDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdNonInventoryItemSalesAndPurchaseDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdNonInventoryItemSalesAndPurchaseDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdNonInventoryItemSalesAndPurchaseDetails) MarshalJSON() ([]byte, error) {
	if src.QbdSalesAndPurchaseDetails != nil {
		return json.Marshal(&src.QbdSalesAndPurchaseDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdNonInventoryItemSalesAndPurchaseDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdSalesAndPurchaseDetails != nil {
		return obj.QbdSalesAndPurchaseDetails
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdNonInventoryItemSalesAndPurchaseDetails) GetActualInstanceValue() (interface{}) {
	if obj.QbdSalesAndPurchaseDetails != nil {
		return *obj.QbdSalesAndPurchaseDetails
	}

	// all schemas are nil
	return nil
}

type NullableQbdNonInventoryItemSalesAndPurchaseDetails struct {
	value *QbdNonInventoryItemSalesAndPurchaseDetails
	isSet bool
}

func (v NullableQbdNonInventoryItemSalesAndPurchaseDetails) Get() *QbdNonInventoryItemSalesAndPurchaseDetails {
	return v.value
}

func (v *NullableQbdNonInventoryItemSalesAndPurchaseDetails) Set(val *QbdNonInventoryItemSalesAndPurchaseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdNonInventoryItemSalesAndPurchaseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdNonInventoryItemSalesAndPurchaseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdNonInventoryItemSalesAndPurchaseDetails(val *QbdNonInventoryItemSalesAndPurchaseDetails) *NullableQbdNonInventoryItemSalesAndPurchaseDetails {
	return &NullableQbdNonInventoryItemSalesAndPurchaseDetails{value: val, isSet: true}
}

func (v NullableQbdNonInventoryItemSalesAndPurchaseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdNonInventoryItemSalesAndPurchaseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


