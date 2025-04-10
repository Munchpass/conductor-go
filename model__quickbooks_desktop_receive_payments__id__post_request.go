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

// checks if the QuickbooksDesktopReceivePaymentsIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopReceivePaymentsIdPostRequest{}

// QuickbooksDesktopReceivePaymentsIdPostRequest struct for QuickbooksDesktopReceivePaymentsIdPostRequest
type QuickbooksDesktopReceivePaymentsIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the receive-payment object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The customer or customer-job to which the payment for this receive-payment is credited.
	CustomerId *string `json:"customerId,omitempty"`
	// The Accounts-Receivable (A/R) account to which this receive-payment is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/R account.  **IMPORTANT**: If this receive-payment is linked to other transactions, this A/R account must match the `receivablesAccount` used in all linked transactions. For example, when refunding a credit card payment, the A/R account must match the one used in the original credit transactions being refunded.
	ReceivablesAccountId *string `json:"receivablesAccountId,omitempty"`
	// The date of this receive-payment, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate *string `json:"transactionDate,omitempty"`
	// The case-sensitive user-defined reference number for this receive-payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.
	RefNumber *string `json:"refNumber,omitempty"`
	// The total monetary amount of this receive-payment, represented as a decimal string.  **NOTE**: The sum of the `paymentAmount` amounts in the `applyToTransactions` array cannot exceed the `totalAmount`, or you will receive an error.
	TotalAmount *string `json:"totalAmount,omitempty"`
	// The market exchange rate between this receive-payment's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// The receive-payment's payment method (e.g., cash, check, credit card).
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	// A memo or note for this receive-payment that will be displayed at the beginning of reports containing details about this receive-payment.
	Memo *string `json:"memo,omitempty"`
	// The account where the funds for this receive-payment will be or have been deposited.
	DepositToAccountId *string `json:"depositToAccountId,omitempty"`
	CreditCardTransaction *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction `json:"creditCardTransaction,omitempty"`
	// The invoices to be paid by this receive-payment. This will create a link between this receive-payment and the specified invoices.  **IMPORTANT**: In each `applyToTransactions` object, you must specify either `paymentAmount`, `applyCredits`, `discountAmount`, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target invoice must have `isPaid=false`, otherwise, QuickBooks will report this object as \"cannot be found\".
	ApplyToTransactions []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner `json:"applyToTransactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopReceivePaymentsIdPostRequest QuickbooksDesktopReceivePaymentsIdPostRequest

// NewQuickbooksDesktopReceivePaymentsIdPostRequest instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopReceivePaymentsIdPostRequest(revisionNumber string) *QuickbooksDesktopReceivePaymentsIdPostRequest {
	this := QuickbooksDesktopReceivePaymentsIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopReceivePaymentsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopReceivePaymentsIdPostRequestWithDefaults() *QuickbooksDesktopReceivePaymentsIdPostRequest {
	this := QuickbooksDesktopReceivePaymentsIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetReceivablesAccountId returns the ReceivablesAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetReceivablesAccountId() string {
	if o == nil || IsNil(o.ReceivablesAccountId) {
		var ret string
		return ret
	}
	return *o.ReceivablesAccountId
}

// GetReceivablesAccountIdOk returns a tuple with the ReceivablesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetReceivablesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivablesAccountId) {
		return nil, false
	}
	return o.ReceivablesAccountId, true
}

// HasReceivablesAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasReceivablesAccountId() bool {
	if o != nil && !IsNil(o.ReceivablesAccountId) {
		return true
	}

	return false
}

// SetReceivablesAccountId gets a reference to the given string and assigns it to the ReceivablesAccountId field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetReceivablesAccountId(v string) {
	o.ReceivablesAccountId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTotalAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetDepositToAccountId returns the DepositToAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetDepositToAccountId() string {
	if o == nil || IsNil(o.DepositToAccountId) {
		var ret string
		return ret
	}
	return *o.DepositToAccountId
}

// GetDepositToAccountIdOk returns a tuple with the DepositToAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetDepositToAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.DepositToAccountId) {
		return nil, false
	}
	return o.DepositToAccountId, true
}

// HasDepositToAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasDepositToAccountId() bool {
	if o != nil && !IsNil(o.DepositToAccountId) {
		return true
	}

	return false
}

// SetDepositToAccountId gets a reference to the given string and assigns it to the DepositToAccountId field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetDepositToAccountId(v string) {
	o.DepositToAccountId = &v
}

// GetCreditCardTransaction returns the CreditCardTransaction field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCreditCardTransaction() QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction {
	if o == nil || IsNil(o.CreditCardTransaction) {
		var ret QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction
		return ret
	}
	return *o.CreditCardTransaction
}

// GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCreditCardTransactionOk() (*QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction, bool) {
	if o == nil || IsNil(o.CreditCardTransaction) {
		return nil, false
	}
	return o.CreditCardTransaction, true
}

// HasCreditCardTransaction returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasCreditCardTransaction() bool {
	if o != nil && !IsNil(o.CreditCardTransaction) {
		return true
	}

	return false
}

// SetCreditCardTransaction gets a reference to the given QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction and assigns it to the CreditCardTransaction field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetCreditCardTransaction(v QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction) {
	o.CreditCardTransaction = &v
}

// GetApplyToTransactions returns the ApplyToTransactions field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner {
	if o == nil || IsNil(o.ApplyToTransactions) {
		var ret []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner
		return ret
	}
	return o.ApplyToTransactions
}

// GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetApplyToTransactionsOk() ([]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool) {
	if o == nil || IsNil(o.ApplyToTransactions) {
		return nil, false
	}
	return o.ApplyToTransactions, true
}

// HasApplyToTransactions returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasApplyToTransactions() bool {
	if o != nil && !IsNil(o.ApplyToTransactions) {
		return true
	}

	return false
}

// SetApplyToTransactions gets a reference to the given []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner and assigns it to the ApplyToTransactions field.
func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) {
	o.ApplyToTransactions = v
}

func (o QuickbooksDesktopReceivePaymentsIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopReceivePaymentsIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.ReceivablesAccountId) {
		toSerialize["receivablesAccountId"] = o.ReceivablesAccountId
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !IsNil(o.RefNumber) {
		toSerialize["refNumber"] = o.RefNumber
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.DepositToAccountId) {
		toSerialize["depositToAccountId"] = o.DepositToAccountId
	}
	if !IsNil(o.CreditCardTransaction) {
		toSerialize["creditCardTransaction"] = o.CreditCardTransaction
	}
	if !IsNil(o.ApplyToTransactions) {
		toSerialize["applyToTransactions"] = o.ApplyToTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopReceivePaymentsIdPostRequest := _QuickbooksDesktopReceivePaymentsIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopReceivePaymentsIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopReceivePaymentsIdPostRequest(varQuickbooksDesktopReceivePaymentsIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "receivablesAccountId")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "paymentMethodId")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "depositToAccountId")
		delete(additionalProperties, "creditCardTransaction")
		delete(additionalProperties, "applyToTransactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopReceivePaymentsIdPostRequest struct {
	value *QuickbooksDesktopReceivePaymentsIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopReceivePaymentsIdPostRequest) Get() *QuickbooksDesktopReceivePaymentsIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopReceivePaymentsIdPostRequest) Set(val *QuickbooksDesktopReceivePaymentsIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopReceivePaymentsIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopReceivePaymentsIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopReceivePaymentsIdPostRequest(val *QuickbooksDesktopReceivePaymentsIdPostRequest) *NullableQuickbooksDesktopReceivePaymentsIdPostRequest {
	return &NullableQuickbooksDesktopReceivePaymentsIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopReceivePaymentsIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopReceivePaymentsIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


