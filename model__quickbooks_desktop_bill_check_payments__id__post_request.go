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

// checks if the QuickbooksDesktopBillCheckPaymentsIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopBillCheckPaymentsIdPostRequest{}

// QuickbooksDesktopBillCheckPaymentsIdPostRequest struct for QuickbooksDesktopBillCheckPaymentsIdPostRequest
type QuickbooksDesktopBillCheckPaymentsIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the bill check payment object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The date of this bill check payment, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate *string `json:"transactionDate,omitempty"`
	// The bank account from which the funds are being drawn for this bill check payment; e.g., Checking or Savings. This bill check payment will decrease the balance of this account.
	BankAccountId *string `json:"bankAccountId,omitempty"`
	// The monetary amount of this bill check payment, represented as a decimal string.
	Amount *string `json:"amount,omitempty"`
	// The market exchange rate between this bill check payment's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// Indicates whether this bill check payment is included in the queue of documents for QuickBooks to print.
	IsQueuedForPrint *bool `json:"isQueuedForPrint,omitempty"`
	// The case-sensitive user-defined reference number for this bill check payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.  **IMPORTANT**: For checks, this field is the check number.
	RefNumber *string `json:"refNumber,omitempty"`
	// A memo or note for this bill check payment.
	Memo *string `json:"memo,omitempty"`
	// The bills to be paid by this bill check payment. This will create a link between this bill check payment and the specified bills.  **IMPORTANT**: In each `applyToTransactions` object, you must specify either `paymentAmount`, `applyCredits`, `discountAmount`, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target bill must have `isPaid=false`, otherwise, QuickBooks will report this object as \"cannot be found\".
	ApplyToTransactions []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner `json:"applyToTransactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopBillCheckPaymentsIdPostRequest QuickbooksDesktopBillCheckPaymentsIdPostRequest

// NewQuickbooksDesktopBillCheckPaymentsIdPostRequest instantiates a new QuickbooksDesktopBillCheckPaymentsIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopBillCheckPaymentsIdPostRequest(revisionNumber string) *QuickbooksDesktopBillCheckPaymentsIdPostRequest {
	this := QuickbooksDesktopBillCheckPaymentsIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopBillCheckPaymentsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopBillCheckPaymentsIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopBillCheckPaymentsIdPostRequestWithDefaults() *QuickbooksDesktopBillCheckPaymentsIdPostRequest {
	this := QuickbooksDesktopBillCheckPaymentsIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetBankAccountId() string {
	if o == nil || IsNil(o.BankAccountId) {
		var ret string
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetBankAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given string and assigns it to the BankAccountId field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetBankAccountId(v string) {
	o.BankAccountId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetAmount(v string) {
	o.Amount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetIsQueuedForPrint returns the IsQueuedForPrint field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetIsQueuedForPrint() bool {
	if o == nil || IsNil(o.IsQueuedForPrint) {
		var ret bool
		return ret
	}
	return *o.IsQueuedForPrint
}

// GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQueuedForPrint) {
		return nil, false
	}
	return o.IsQueuedForPrint, true
}

// HasIsQueuedForPrint returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasIsQueuedForPrint() bool {
	if o != nil && !IsNil(o.IsQueuedForPrint) {
		return true
	}

	return false
}

// SetIsQueuedForPrint gets a reference to the given bool and assigns it to the IsQueuedForPrint field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetIsQueuedForPrint(v bool) {
	o.IsQueuedForPrint = &v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetApplyToTransactions returns the ApplyToTransactions field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner {
	if o == nil || IsNil(o.ApplyToTransactions) {
		var ret []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner
		return ret
	}
	return o.ApplyToTransactions
}

// GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetApplyToTransactionsOk() ([]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool) {
	if o == nil || IsNil(o.ApplyToTransactions) {
		return nil, false
	}
	return o.ApplyToTransactions, true
}

// HasApplyToTransactions returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasApplyToTransactions() bool {
	if o != nil && !IsNil(o.ApplyToTransactions) {
		return true
	}

	return false
}

// SetApplyToTransactions gets a reference to the given []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner and assigns it to the ApplyToTransactions field.
func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) {
	o.ApplyToTransactions = v
}

func (o QuickbooksDesktopBillCheckPaymentsIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopBillCheckPaymentsIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !IsNil(o.BankAccountId) {
		toSerialize["bankAccountId"] = o.BankAccountId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.IsQueuedForPrint) {
		toSerialize["isQueuedForPrint"] = o.IsQueuedForPrint
	}
	if !IsNil(o.RefNumber) {
		toSerialize["refNumber"] = o.RefNumber
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.ApplyToTransactions) {
		toSerialize["applyToTransactions"] = o.ApplyToTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revisionNumber",
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

	varQuickbooksDesktopBillCheckPaymentsIdPostRequest := _QuickbooksDesktopBillCheckPaymentsIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopBillCheckPaymentsIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopBillCheckPaymentsIdPostRequest(varQuickbooksDesktopBillCheckPaymentsIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "bankAccountId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "isQueuedForPrint")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "applyToTransactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest struct {
	value *QuickbooksDesktopBillCheckPaymentsIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) Get() *QuickbooksDesktopBillCheckPaymentsIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) Set(val *QuickbooksDesktopBillCheckPaymentsIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopBillCheckPaymentsIdPostRequest(val *QuickbooksDesktopBillCheckPaymentsIdPostRequest) *NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest {
	return &NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopBillCheckPaymentsIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


