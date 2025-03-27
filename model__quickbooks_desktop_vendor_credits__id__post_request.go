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

// checks if the QuickbooksDesktopVendorCreditsIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopVendorCreditsIdPostRequest{}

// QuickbooksDesktopVendorCreditsIdPostRequest struct for QuickbooksDesktopVendorCreditsIdPostRequest
type QuickbooksDesktopVendorCreditsIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the vendor credit object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The vendor who sent this vendor credit for goods or services purchased.
	VendorId *string `json:"vendorId,omitempty"`
	// The Accounts-Payable (A/P) account to which this vendor credit is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this vendor credit is linked to other transactions, this A/P account must match the `payablesAccount` used in those other transactions.
	PayablesAccountId *string `json:"payablesAccountId,omitempty"`
	// The date of this vendor credit, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate *string `json:"transactionDate,omitempty"`
	// The case-sensitive user-defined reference number for this vendor credit, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.
	RefNumber *string `json:"refNumber,omitempty"`
	// A memo or note for this vendor credit.
	Memo *string `json:"memo,omitempty"`
	// The sales-tax code for this vendor credit, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the vendor credit's individual lines.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The market exchange rate between this vendor credit's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// When `true`, removes all existing expense lines associated with this vendor credit. To modify or add individual expense lines, use the field `expenseLines` instead.
	ClearExpenseLines *bool `json:"clearExpenseLines,omitempty"`
	// The vendor credit's expense lines, each representing one line in this expense.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing expense lines for the vendor credit with this array. To keep any existing expense lines, you must include them in this array even if they have not changed. **Any expense lines not included will be removed.**  2. To add a new expense line, include it here with the `id` field set to `-1`.  3. If you do not wish to modify any expense lines, omit this field entirely to keep them unchanged.
	ExpenseLines []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner `json:"expenseLines,omitempty"`
	// When `true`, removes all existing item lines associated with this vendor credit. To modify or add individual item lines, use the field `itemLines` instead.
	ClearItemLines *bool `json:"clearItemLines,omitempty"`
	// The vendor credit's item lines, each representing the purchase of a specific item or service.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item lines for the vendor credit with this array. To keep any existing item lines, you must include them in this array even if they have not changed. **Any item lines not included will be removed.**  2. To add a new item line, include it here with the `id` field set to `-1`.  3. If you do not wish to modify any item lines, omit this field entirely to keep them unchanged.
	ItemLines []QuickbooksDesktopBillsIdPostRequestItemLinesInner `json:"itemLines,omitempty"`
	// The vendor credit's item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item group lines for the vendor credit with this array. To keep any existing item group lines, you must include them in this array even if they have not changed. **Any item group lines not included will be removed.**  2. To add a new item group line, include it here with the `id` field set to `-1`.  3. If you do not wish to modify any item group lines, omit this field entirely to keep them unchanged.
	ItemLineGroups []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner `json:"itemLineGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopVendorCreditsIdPostRequest QuickbooksDesktopVendorCreditsIdPostRequest

// NewQuickbooksDesktopVendorCreditsIdPostRequest instantiates a new QuickbooksDesktopVendorCreditsIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopVendorCreditsIdPostRequest(revisionNumber string) *QuickbooksDesktopVendorCreditsIdPostRequest {
	this := QuickbooksDesktopVendorCreditsIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopVendorCreditsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopVendorCreditsIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopVendorCreditsIdPostRequestWithDefaults() *QuickbooksDesktopVendorCreditsIdPostRequest {
	this := QuickbooksDesktopVendorCreditsIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetVendorId(v string) {
	o.VendorId = &v
}

// GetPayablesAccountId returns the PayablesAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetPayablesAccountId() string {
	if o == nil || IsNil(o.PayablesAccountId) {
		var ret string
		return ret
	}
	return *o.PayablesAccountId
}

// GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetPayablesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayablesAccountId) {
		return nil, false
	}
	return o.PayablesAccountId, true
}

// HasPayablesAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasPayablesAccountId() bool {
	if o != nil && !IsNil(o.PayablesAccountId) {
		return true
	}

	return false
}

// SetPayablesAccountId gets a reference to the given string and assigns it to the PayablesAccountId field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetPayablesAccountId(v string) {
	o.PayablesAccountId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetClearExpenseLines returns the ClearExpenseLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetClearExpenseLines() bool {
	if o == nil || IsNil(o.ClearExpenseLines) {
		var ret bool
		return ret
	}
	return *o.ClearExpenseLines
}

// GetClearExpenseLinesOk returns a tuple with the ClearExpenseLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetClearExpenseLinesOk() (*bool, bool) {
	if o == nil || IsNil(o.ClearExpenseLines) {
		return nil, false
	}
	return o.ClearExpenseLines, true
}

// HasClearExpenseLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasClearExpenseLines() bool {
	if o != nil && !IsNil(o.ClearExpenseLines) {
		return true
	}

	return false
}

// SetClearExpenseLines gets a reference to the given bool and assigns it to the ClearExpenseLines field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetClearExpenseLines(v bool) {
	o.ClearExpenseLines = &v
}

// GetExpenseLines returns the ExpenseLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetExpenseLines() []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner {
	if o == nil || IsNil(o.ExpenseLines) {
		var ret []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner
		return ret
	}
	return o.ExpenseLines
}

// GetExpenseLinesOk returns a tuple with the ExpenseLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetExpenseLinesOk() ([]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner, bool) {
	if o == nil || IsNil(o.ExpenseLines) {
		return nil, false
	}
	return o.ExpenseLines, true
}

// HasExpenseLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasExpenseLines() bool {
	if o != nil && !IsNil(o.ExpenseLines) {
		return true
	}

	return false
}

// SetExpenseLines gets a reference to the given []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner and assigns it to the ExpenseLines field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) {
	o.ExpenseLines = v
}

// GetClearItemLines returns the ClearItemLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetClearItemLines() bool {
	if o == nil || IsNil(o.ClearItemLines) {
		var ret bool
		return ret
	}
	return *o.ClearItemLines
}

// GetClearItemLinesOk returns a tuple with the ClearItemLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetClearItemLinesOk() (*bool, bool) {
	if o == nil || IsNil(o.ClearItemLines) {
		return nil, false
	}
	return o.ClearItemLines, true
}

// HasClearItemLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasClearItemLines() bool {
	if o != nil && !IsNil(o.ClearItemLines) {
		return true
	}

	return false
}

// SetClearItemLines gets a reference to the given bool and assigns it to the ClearItemLines field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetClearItemLines(v bool) {
	o.ClearItemLines = &v
}

// GetItemLines returns the ItemLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetItemLines() []QuickbooksDesktopBillsIdPostRequestItemLinesInner {
	if o == nil || IsNil(o.ItemLines) {
		var ret []QuickbooksDesktopBillsIdPostRequestItemLinesInner
		return ret
	}
	return o.ItemLines
}

// GetItemLinesOk returns a tuple with the ItemLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetItemLinesOk() ([]QuickbooksDesktopBillsIdPostRequestItemLinesInner, bool) {
	if o == nil || IsNil(o.ItemLines) {
		return nil, false
	}
	return o.ItemLines, true
}

// HasItemLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasItemLines() bool {
	if o != nil && !IsNil(o.ItemLines) {
		return true
	}

	return false
}

// SetItemLines gets a reference to the given []QuickbooksDesktopBillsIdPostRequestItemLinesInner and assigns it to the ItemLines field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetItemLines(v []QuickbooksDesktopBillsIdPostRequestItemLinesInner) {
	o.ItemLines = v
}

// GetItemLineGroups returns the ItemLineGroups field value if set, zero value otherwise.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner {
	if o == nil || IsNil(o.ItemLineGroups) {
		var ret []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner
		return ret
	}
	return o.ItemLineGroups
}

// GetItemLineGroupsOk returns a tuple with the ItemLineGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) GetItemLineGroupsOk() ([]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner, bool) {
	if o == nil || IsNil(o.ItemLineGroups) {
		return nil, false
	}
	return o.ItemLineGroups, true
}

// HasItemLineGroups returns a boolean if a field has been set.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) HasItemLineGroups() bool {
	if o != nil && !IsNil(o.ItemLineGroups) {
		return true
	}

	return false
}

// SetItemLineGroups gets a reference to the given []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner and assigns it to the ItemLineGroups field.
func (o *QuickbooksDesktopVendorCreditsIdPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner) {
	o.ItemLineGroups = v
}

func (o QuickbooksDesktopVendorCreditsIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopVendorCreditsIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.VendorId) {
		toSerialize["vendorId"] = o.VendorId
	}
	if !IsNil(o.PayablesAccountId) {
		toSerialize["payablesAccountId"] = o.PayablesAccountId
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
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
	if !IsNil(o.ClearExpenseLines) {
		toSerialize["clearExpenseLines"] = o.ClearExpenseLines
	}
	if !IsNil(o.ExpenseLines) {
		toSerialize["expenseLines"] = o.ExpenseLines
	}
	if !IsNil(o.ClearItemLines) {
		toSerialize["clearItemLines"] = o.ClearItemLines
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

func (o *QuickbooksDesktopVendorCreditsIdPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopVendorCreditsIdPostRequest := _QuickbooksDesktopVendorCreditsIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopVendorCreditsIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopVendorCreditsIdPostRequest(varQuickbooksDesktopVendorCreditsIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "vendorId")
		delete(additionalProperties, "payablesAccountId")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "clearExpenseLines")
		delete(additionalProperties, "expenseLines")
		delete(additionalProperties, "clearItemLines")
		delete(additionalProperties, "itemLines")
		delete(additionalProperties, "itemLineGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopVendorCreditsIdPostRequest struct {
	value *QuickbooksDesktopVendorCreditsIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopVendorCreditsIdPostRequest) Get() *QuickbooksDesktopVendorCreditsIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopVendorCreditsIdPostRequest) Set(val *QuickbooksDesktopVendorCreditsIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopVendorCreditsIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopVendorCreditsIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopVendorCreditsIdPostRequest(val *QuickbooksDesktopVendorCreditsIdPostRequest) *NullableQuickbooksDesktopVendorCreditsIdPostRequest {
	return &NullableQuickbooksDesktopVendorCreditsIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopVendorCreditsIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopVendorCreditsIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


