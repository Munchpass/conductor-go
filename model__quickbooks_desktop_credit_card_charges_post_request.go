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

// checks if the QuickbooksDesktopCreditCardChargesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopCreditCardChargesPostRequest{}

// QuickbooksDesktopCreditCardChargesPostRequest struct for QuickbooksDesktopCreditCardChargesPostRequest
type QuickbooksDesktopCreditCardChargesPostRequest struct {
	// The bank or credit card account to which money is owed for this credit card charge.
	AccountId string `json:"accountId"`
	// The vendor or company from whom merchandise or services were purchased for this credit card charge.
	PayeeId *string `json:"payeeId,omitempty"`
	// The date of this credit card charge, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate string `json:"transactionDate"`
	// The case-sensitive user-defined reference number for this credit card charge, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment).
	RefNumber *string `json:"refNumber,omitempty"`
	// A memo or note for this credit card charge.
	Memo *string `json:"memo,omitempty"`
	// The sales-tax code for this credit card charge, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the payee. This can be overridden on the credit card charge's individual lines.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The market exchange rate between this credit card charge's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error.
	ExternalId *string `json:"externalId,omitempty"`
	// The credit card charge's expense lines, each representing one line in this expense.
	ExpenseLines []QuickbooksDesktopBillsPostRequestExpenseLinesInner `json:"expenseLines,omitempty"`
	// The credit card charge's item lines, each representing the purchase of a specific item or service.
	ItemLines []QuickbooksDesktopBillsPostRequestItemLinesInner `json:"itemLines,omitempty"`
	// The credit card charge's item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.
	ItemLineGroups []QuickbooksDesktopBillsPostRequestItemLineGroupsInner `json:"itemLineGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopCreditCardChargesPostRequest QuickbooksDesktopCreditCardChargesPostRequest

// NewQuickbooksDesktopCreditCardChargesPostRequest instantiates a new QuickbooksDesktopCreditCardChargesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopCreditCardChargesPostRequest(accountId string, transactionDate string) *QuickbooksDesktopCreditCardChargesPostRequest {
	this := QuickbooksDesktopCreditCardChargesPostRequest{}
	this.AccountId = accountId
	this.TransactionDate = transactionDate
	return &this
}

// NewQuickbooksDesktopCreditCardChargesPostRequestWithDefaults instantiates a new QuickbooksDesktopCreditCardChargesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopCreditCardChargesPostRequestWithDefaults() *QuickbooksDesktopCreditCardChargesPostRequest {
	this := QuickbooksDesktopCreditCardChargesPostRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetPayeeId() string {
	if o == nil || IsNil(o.PayeeId) {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeId) {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasPayeeId() bool {
	if o != nil && !IsNil(o.PayeeId) {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetTransactionDate returns the TransactionDate field value
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetTransactionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDate, true
}

// SetTransactionDate sets field value
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExpenseLines returns the ExpenseLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner {
	if o == nil || IsNil(o.ExpenseLines) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInner
		return ret
	}
	return o.ExpenseLines
}

// GetExpenseLinesOk returns a tuple with the ExpenseLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExpenseLinesOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool) {
	if o == nil || IsNil(o.ExpenseLines) {
		return nil, false
	}
	return o.ExpenseLines, true
}

// HasExpenseLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExpenseLines() bool {
	if o != nil && !IsNil(o.ExpenseLines) {
		return true
	}

	return false
}

// SetExpenseLines gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInner and assigns it to the ExpenseLines field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner) {
	o.ExpenseLines = v
}

// GetItemLines returns the ItemLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner {
	if o == nil || IsNil(o.ItemLines) {
		var ret []QuickbooksDesktopBillsPostRequestItemLinesInner
		return ret
	}
	return o.ItemLines
}

// GetItemLinesOk returns a tuple with the ItemLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLinesOk() ([]QuickbooksDesktopBillsPostRequestItemLinesInner, bool) {
	if o == nil || IsNil(o.ItemLines) {
		return nil, false
	}
	return o.ItemLines, true
}

// HasItemLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasItemLines() bool {
	if o != nil && !IsNil(o.ItemLines) {
		return true
	}

	return false
}

// SetItemLines gets a reference to the given []QuickbooksDesktopBillsPostRequestItemLinesInner and assigns it to the ItemLines field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner) {
	o.ItemLines = v
}

// GetItemLineGroups returns the ItemLineGroups field value if set, zero value otherwise.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner {
	if o == nil || IsNil(o.ItemLineGroups) {
		var ret []QuickbooksDesktopBillsPostRequestItemLineGroupsInner
		return ret
	}
	return o.ItemLineGroups
}

// GetItemLineGroupsOk returns a tuple with the ItemLineGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLineGroupsOk() ([]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool) {
	if o == nil || IsNil(o.ItemLineGroups) {
		return nil, false
	}
	return o.ItemLineGroups, true
}

// HasItemLineGroups returns a boolean if a field has been set.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasItemLineGroups() bool {
	if o != nil && !IsNil(o.ItemLineGroups) {
		return true
	}

	return false
}

// SetItemLineGroups gets a reference to the given []QuickbooksDesktopBillsPostRequestItemLineGroupsInner and assigns it to the ItemLineGroups field.
func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner) {
	o.ItemLineGroups = v
}

func (o QuickbooksDesktopCreditCardChargesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopCreditCardChargesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	if !IsNil(o.PayeeId) {
		toSerialize["payeeId"] = o.PayeeId
	}
	toSerialize["transactionDate"] = o.TransactionDate
	if !IsNil(o.RefNumber) {
		toSerialize["refNumber"] = o.RefNumber
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.ExpenseLines) {
		toSerialize["expenseLines"] = o.ExpenseLines
	}
	if !IsNil(o.ItemLines) {
		toSerialize["itemLines"] = o.ItemLines
	}
	if !IsNil(o.ItemLineGroups) {
		toSerialize["itemLineGroups"] = o.ItemLineGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopCreditCardChargesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountId",
		"transactionDate",
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

	varQuickbooksDesktopCreditCardChargesPostRequest := _QuickbooksDesktopCreditCardChargesPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopCreditCardChargesPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopCreditCardChargesPostRequest(varQuickbooksDesktopCreditCardChargesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "payeeId")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "expenseLines")
		delete(additionalProperties, "itemLines")
		delete(additionalProperties, "itemLineGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopCreditCardChargesPostRequest struct {
	value *QuickbooksDesktopCreditCardChargesPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopCreditCardChargesPostRequest) Get() *QuickbooksDesktopCreditCardChargesPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopCreditCardChargesPostRequest) Set(val *QuickbooksDesktopCreditCardChargesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopCreditCardChargesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopCreditCardChargesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopCreditCardChargesPostRequest(val *QuickbooksDesktopCreditCardChargesPostRequest) *NullableQuickbooksDesktopCreditCardChargesPostRequest {
	return &NullableQuickbooksDesktopCreditCardChargesPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopCreditCardChargesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopCreditCardChargesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


