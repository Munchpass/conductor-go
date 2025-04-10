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

// checks if the QbdReceivableTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdReceivableTransaction{}

// QbdReceivableTransaction struct for QbdReceivableTransaction
type QbdReceivableTransaction struct {
	// The ID of the receivable transaction to which this payment is applied.
	TransactionId string `json:"transactionId"`
	// The type of transaction for this receivable transaction.
	TransactionType string `json:"transactionType"`
	// The date of this receivable transaction, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate string `json:"transactionDate"`
	// The case-sensitive user-defined reference number for this receivable transaction, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.
	RefNumber string `json:"refNumber"`
	// The outstanding balance of this receivable transaction after applying any credits or payments. Represented as a decimal string.
	BalanceRemaining string `json:"balanceRemaining"`
	// The monetary amount of this receivable transaction, represented as a decimal string.
	Amount string `json:"amount"`
	// The monetary amount by which to reduce the receivable transaction's receivable amount, represented as a decimal string.
	DiscountAmount string `json:"discountAmount"`
	DiscountAccount QbdReceivableTransactionDiscountAccount `json:"discountAccount"`
	DiscountClass QbdReceivableTransactionDiscountClass `json:"discountClass"`
	// The receivable transaction's linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter `includeLinkedTransactions` when fetching a list of receivable transactions to receive this field because it is not returned by default.
	LinkedTransactions []QbdLinkedTransaction `json:"linkedTransactions"`
	AdditionalProperties map[string]interface{}
}

type _QbdReceivableTransaction QbdReceivableTransaction

// NewQbdReceivableTransaction instantiates a new QbdReceivableTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdReceivableTransaction(transactionId string, transactionType string, transactionDate string, refNumber string, balanceRemaining string, amount string, discountAmount string, discountAccount QbdReceivableTransactionDiscountAccount, discountClass QbdReceivableTransactionDiscountClass, linkedTransactions []QbdLinkedTransaction) *QbdReceivableTransaction {
	this := QbdReceivableTransaction{}
	this.TransactionId = transactionId
	this.TransactionType = transactionType
	this.TransactionDate = transactionDate
	this.RefNumber = refNumber
	this.BalanceRemaining = balanceRemaining
	this.Amount = amount
	this.DiscountAmount = discountAmount
	this.DiscountAccount = discountAccount
	this.DiscountClass = discountClass
	this.LinkedTransactions = linkedTransactions
	return &this
}

// NewQbdReceivableTransactionWithDefaults instantiates a new QbdReceivableTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdReceivableTransactionWithDefaults() *QbdReceivableTransaction {
	this := QbdReceivableTransaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *QbdReceivableTransaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *QbdReceivableTransaction) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetTransactionType returns the TransactionType field value
func (o *QbdReceivableTransaction) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *QbdReceivableTransaction) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetTransactionDate returns the TransactionDate field value
func (o *QbdReceivableTransaction) GetTransactionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetTransactionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDate, true
}

// SetTransactionDate sets field value
func (o *QbdReceivableTransaction) SetTransactionDate(v string) {
	o.TransactionDate = v
}

// GetRefNumber returns the RefNumber field value
func (o *QbdReceivableTransaction) GetRefNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetRefNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefNumber, true
}

// SetRefNumber sets field value
func (o *QbdReceivableTransaction) SetRefNumber(v string) {
	o.RefNumber = v
}

// GetBalanceRemaining returns the BalanceRemaining field value
func (o *QbdReceivableTransaction) GetBalanceRemaining() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceRemaining
}

// GetBalanceRemainingOk returns a tuple with the BalanceRemaining field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetBalanceRemainingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceRemaining, true
}

// SetBalanceRemaining sets field value
func (o *QbdReceivableTransaction) SetBalanceRemaining(v string) {
	o.BalanceRemaining = v
}

// GetAmount returns the Amount field value
func (o *QbdReceivableTransaction) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *QbdReceivableTransaction) SetAmount(v string) {
	o.Amount = v
}

// GetDiscountAmount returns the DiscountAmount field value
func (o *QbdReceivableTransaction) GetDiscountAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetDiscountAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountAmount, true
}

// SetDiscountAmount sets field value
func (o *QbdReceivableTransaction) SetDiscountAmount(v string) {
	o.DiscountAmount = v
}

// GetDiscountAccount returns the DiscountAccount field value
func (o *QbdReceivableTransaction) GetDiscountAccount() QbdReceivableTransactionDiscountAccount {
	if o == nil {
		var ret QbdReceivableTransactionDiscountAccount
		return ret
	}

	return o.DiscountAccount
}

// GetDiscountAccountOk returns a tuple with the DiscountAccount field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetDiscountAccountOk() (*QbdReceivableTransactionDiscountAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountAccount, true
}

// SetDiscountAccount sets field value
func (o *QbdReceivableTransaction) SetDiscountAccount(v QbdReceivableTransactionDiscountAccount) {
	o.DiscountAccount = v
}

// GetDiscountClass returns the DiscountClass field value
func (o *QbdReceivableTransaction) GetDiscountClass() QbdReceivableTransactionDiscountClass {
	if o == nil {
		var ret QbdReceivableTransactionDiscountClass
		return ret
	}

	return o.DiscountClass
}

// GetDiscountClassOk returns a tuple with the DiscountClass field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetDiscountClassOk() (*QbdReceivableTransactionDiscountClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountClass, true
}

// SetDiscountClass sets field value
func (o *QbdReceivableTransaction) SetDiscountClass(v QbdReceivableTransactionDiscountClass) {
	o.DiscountClass = v
}

// GetLinkedTransactions returns the LinkedTransactions field value
func (o *QbdReceivableTransaction) GetLinkedTransactions() []QbdLinkedTransaction {
	if o == nil {
		var ret []QbdLinkedTransaction
		return ret
	}

	return o.LinkedTransactions
}

// GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field value
// and a boolean to check if the value has been set.
func (o *QbdReceivableTransaction) GetLinkedTransactionsOk() ([]QbdLinkedTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedTransactions, true
}

// SetLinkedTransactions sets field value
func (o *QbdReceivableTransaction) SetLinkedTransactions(v []QbdLinkedTransaction) {
	o.LinkedTransactions = v
}

func (o QbdReceivableTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdReceivableTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionId"] = o.TransactionId
	toSerialize["transactionType"] = o.TransactionType
	toSerialize["transactionDate"] = o.TransactionDate
	toSerialize["refNumber"] = o.RefNumber
	toSerialize["balanceRemaining"] = o.BalanceRemaining
	toSerialize["amount"] = o.Amount
	toSerialize["discountAmount"] = o.DiscountAmount
	toSerialize["discountAccount"] = o.DiscountAccount
	toSerialize["discountClass"] = o.DiscountClass
	toSerialize["linkedTransactions"] = o.LinkedTransactions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdReceivableTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactionId",
		"transactionType",
		"transactionDate",
		"refNumber",
		"balanceRemaining",
		"amount",
		"discountAmount",
		"discountAccount",
		"discountClass",
		"linkedTransactions",
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

	varQbdReceivableTransaction := _QbdReceivableTransaction{}

	err = json.Unmarshal(data, &varQbdReceivableTransaction)

	if err != nil {
		return err
	}

	*o = QbdReceivableTransaction(varQbdReceivableTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transactionId")
		delete(additionalProperties, "transactionType")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "balanceRemaining")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "discountAmount")
		delete(additionalProperties, "discountAccount")
		delete(additionalProperties, "discountClass")
		delete(additionalProperties, "linkedTransactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdReceivableTransaction struct {
	value *QbdReceivableTransaction
	isSet bool
}

func (v NullableQbdReceivableTransaction) Get() *QbdReceivableTransaction {
	return v.value
}

func (v *NullableQbdReceivableTransaction) Set(val *QbdReceivableTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdReceivableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdReceivableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdReceivableTransaction(val *QbdReceivableTransaction) *NullableQbdReceivableTransaction {
	return &NullableQbdReceivableTransaction{value: val, isSet: true}
}

func (v NullableQbdReceivableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdReceivableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


