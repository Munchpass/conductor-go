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
)

// checks if the QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine{}

// QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine An existing transaction line that you wish to link to this item line. Note that this only links to a single transaction line item, not an entire transaction. If you want to link an entire transaction and bring in all its lines, instead use the field `linkToTransactionIds` on the parent transaction, if available. If the parent transaction is a bill or an item receipt, you can only link to purchase orders; QuickBooks does not support linking these transactions to other transaction types.  Transaction lines can only be linked when creating this item line and cannot be unlinked later.  **IMPORTANT**: If you use `linkToTransactionLine` on this item line, you cannot use the field `item` on this line (QuickBooks will return an error) because this field brings in all of the item information you need. You can, however, specify whatever `quantity` or `rate` that you want, or any other transaction line element other than `item`.  If the parent transaction supports the `linkToTransactionIds` field, you can use both `linkToTransactionLine` (on this item line) and `linkToTransactionIds` (on its parent transaction) in the same request as long as they do NOT link to the same transaction (otherwise, QuickBooks will return an error). QuickBooks will also return an error if you attempt to link a transaction that is empty or already closed.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transaction line in this endpoint's response even when this request is successful. To see the transaction line linked via this field, refetch the parent transaction and check the `linkedTransactions` response field. If fetching a list of transactions, you must also specify the parameter `includeLinkedTransactions=true` to see the `linkedTransactions` response field.
type QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine struct {
	// The ID of the transaction to which to link this transaction.
	TransactionId string `json:"transactionId"`
	// The ID of the transaction line to which to link this transaction.
	TransactionLineId string `json:"transactionLineId"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine

// NewQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine(transactionId string, transactionLineId string) *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine {
	this := QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine{}
	this.TransactionId = transactionId
	this.TransactionLineId = transactionLineId
	return &this
}

// NewQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLineWithDefaults instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLineWithDefaults() *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine {
	this := QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetTransactionLineId returns the TransactionLineId field value
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) GetTransactionLineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionLineId
}

// GetTransactionLineIdOk returns a tuple with the TransactionLineId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) GetTransactionLineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionLineId, true
}

// SetTransactionLineId sets field value
func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) SetTransactionLineId(v string) {
	o.TransactionLineId = v
}

func (o QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionId"] = o.TransactionId
	toSerialize["transactionLineId"] = o.TransactionLineId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactionId",
		"transactionLineId",
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

	varQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine := _QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine{}

	err = json.Unmarshal(data, &varQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine(varQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transactionId")
		delete(additionalProperties, "transactionLineId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine struct {
	value *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine
	isSet bool
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) Get() *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine {
	return v.value
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) Set(val *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine(val *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) *NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine {
	return &NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


