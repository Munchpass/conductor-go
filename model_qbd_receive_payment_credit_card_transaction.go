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

// QbdReceivePaymentCreditCardTransaction - The credit card transaction data for this receive-payment's payment when using QuickBooks Merchant Services (QBMS).
type QbdReceivePaymentCreditCardTransaction struct {
	QbdCreditCardTransaction *QbdCreditCardTransaction
}

// QbdCreditCardTransactionAsQbdReceivePaymentCreditCardTransaction is a convenience function that returns QbdCreditCardTransaction wrapped in QbdReceivePaymentCreditCardTransaction
func QbdCreditCardTransactionAsQbdReceivePaymentCreditCardTransaction(v *QbdCreditCardTransaction) QbdReceivePaymentCreditCardTransaction {
	return QbdReceivePaymentCreditCardTransaction{
		QbdCreditCardTransaction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QbdReceivePaymentCreditCardTransaction) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QbdCreditCardTransaction
	err = newStrictDecoder(data).Decode(&dst.QbdCreditCardTransaction)
	if err == nil {
		jsonQbdCreditCardTransaction, _ := json.Marshal(dst.QbdCreditCardTransaction)
		if string(jsonQbdCreditCardTransaction) == "{}" { // empty struct
			dst.QbdCreditCardTransaction = nil
		} else {
			if err = validator.Validate(dst.QbdCreditCardTransaction); err != nil {
				dst.QbdCreditCardTransaction = nil
			} else {
				match++
			}
		}
	} else {
		dst.QbdCreditCardTransaction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QbdCreditCardTransaction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QbdReceivePaymentCreditCardTransaction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QbdReceivePaymentCreditCardTransaction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QbdReceivePaymentCreditCardTransaction) MarshalJSON() ([]byte, error) {
	if src.QbdCreditCardTransaction != nil {
		return json.Marshal(&src.QbdCreditCardTransaction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QbdReceivePaymentCreditCardTransaction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QbdCreditCardTransaction != nil {
		return obj.QbdCreditCardTransaction
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj QbdReceivePaymentCreditCardTransaction) GetActualInstanceValue() (interface{}) {
	if obj.QbdCreditCardTransaction != nil {
		return *obj.QbdCreditCardTransaction
	}

	// all schemas are nil
	return nil
}

type NullableQbdReceivePaymentCreditCardTransaction struct {
	value *QbdReceivePaymentCreditCardTransaction
	isSet bool
}

func (v NullableQbdReceivePaymentCreditCardTransaction) Get() *QbdReceivePaymentCreditCardTransaction {
	return v.value
}

func (v *NullableQbdReceivePaymentCreditCardTransaction) Set(val *QbdReceivePaymentCreditCardTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdReceivePaymentCreditCardTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdReceivePaymentCreditCardTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdReceivePaymentCreditCardTransaction(val *QbdReceivePaymentCreditCardTransaction) *NullableQbdReceivePaymentCreditCardTransaction {
	return &NullableQbdReceivePaymentCreditCardTransaction{value: val, isSet: true}
}

func (v NullableQbdReceivePaymentCreditCardTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdReceivePaymentCreditCardTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


