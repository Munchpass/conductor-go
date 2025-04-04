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

// checks if the QuickbooksDesktopItemReceiptsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopItemReceiptsPostRequest{}

// QuickbooksDesktopItemReceiptsPostRequest struct for QuickbooksDesktopItemReceiptsPostRequest
type QuickbooksDesktopItemReceiptsPostRequest struct {
	// The vendor who sent this item receipt for goods or services purchased.
	VendorId string `json:"vendorId"`
	// The Accounts-Payable (A/P) account to which this item receipt is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this item receipt is linked to other transactions, this A/P account must match the `payablesAccount` used in those other transactions.
	PayablesAccountId *string `json:"payablesAccountId,omitempty"`
	// The date of this item receipt, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate string `json:"transactionDate"`
	// The case-sensitive user-defined reference number for this item receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment).
	RefNumber *string `json:"refNumber,omitempty"`
	// A memo or note for this item receipt.
	Memo *string `json:"memo,omitempty"`
	// The sales-tax code for this item receipt, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the item receipt's individual lines.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The market exchange rate between this item receipt's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error.
	ExternalId *string `json:"externalId,omitempty"`
	// IDs of existing purchase orders that you wish to link to this item receipt. Note that this links entire transactions, not individual transaction lines. If you want to link individual lines in a transaction, instead use the field `linkToTransactionLine` on this item receipt's lines, if available.  Transactions can only be linked when creating this item receipt and cannot be unlinked later.  You can use both `linkToTransactionIds` (on this item receipt) and `linkToTransactionLine` (on its transaction lines) as long as they do NOT link to the same transaction (otherwise, QuickBooks will return an error). QuickBooks will also return an error if you attempt to link a transaction that is empty or already closed.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint's response even when this request is successful. To see the transactions linked via this field, refetch the item receipt and check the `linkedTransactions` response field. If fetching a list of item receipts, you must also specify the parameter `includeLinkedTransactions=true` to see the `linkedTransactions` response field.
	LinkToTransactionIds []string `json:"linkToTransactionIds,omitempty"`
	// The item receipt's expense lines, each representing one line in this expense.
	ExpenseLines []QuickbooksDesktopBillsPostRequestExpenseLinesInner `json:"expenseLines,omitempty"`
	// The item receipt's item lines, each representing the purchase of a specific item or service.
	ItemLines []QuickbooksDesktopBillsPostRequestItemLinesInner `json:"itemLines,omitempty"`
	// The item receipt's item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.
	ItemLineGroups []QuickbooksDesktopBillsPostRequestItemLineGroupsInner `json:"itemLineGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopItemReceiptsPostRequest QuickbooksDesktopItemReceiptsPostRequest

// NewQuickbooksDesktopItemReceiptsPostRequest instantiates a new QuickbooksDesktopItemReceiptsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopItemReceiptsPostRequest(vendorId string, transactionDate string) *QuickbooksDesktopItemReceiptsPostRequest {
	this := QuickbooksDesktopItemReceiptsPostRequest{}
	this.VendorId = vendorId
	this.TransactionDate = transactionDate
	return &this
}

// NewQuickbooksDesktopItemReceiptsPostRequestWithDefaults instantiates a new QuickbooksDesktopItemReceiptsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopItemReceiptsPostRequestWithDefaults() *QuickbooksDesktopItemReceiptsPostRequest {
	this := QuickbooksDesktopItemReceiptsPostRequest{}
	return &this
}

// GetVendorId returns the VendorId field value
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetVendorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetVendorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorId, true
}

// SetVendorId sets field value
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetVendorId(v string) {
	o.VendorId = v
}

// GetPayablesAccountId returns the PayablesAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetPayablesAccountId() string {
	if o == nil || IsNil(o.PayablesAccountId) {
		var ret string
		return ret
	}
	return *o.PayablesAccountId
}

// GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetPayablesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayablesAccountId) {
		return nil, false
	}
	return o.PayablesAccountId, true
}

// HasPayablesAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasPayablesAccountId() bool {
	if o != nil && !IsNil(o.PayablesAccountId) {
		return true
	}

	return false
}

// SetPayablesAccountId gets a reference to the given string and assigns it to the PayablesAccountId field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetPayablesAccountId(v string) {
	o.PayablesAccountId = &v
}

// GetTransactionDate returns the TransactionDate field value
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetTransactionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDate, true
}

// SetTransactionDate sets field value
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetLinkToTransactionIds returns the LinkToTransactionIds field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetLinkToTransactionIds() []string {
	if o == nil || IsNil(o.LinkToTransactionIds) {
		var ret []string
		return ret
	}
	return o.LinkToTransactionIds
}

// GetLinkToTransactionIdsOk returns a tuple with the LinkToTransactionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetLinkToTransactionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LinkToTransactionIds) {
		return nil, false
	}
	return o.LinkToTransactionIds, true
}

// HasLinkToTransactionIds returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasLinkToTransactionIds() bool {
	if o != nil && !IsNil(o.LinkToTransactionIds) {
		return true
	}

	return false
}

// SetLinkToTransactionIds gets a reference to the given []string and assigns it to the LinkToTransactionIds field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetLinkToTransactionIds(v []string) {
	o.LinkToTransactionIds = v
}

// GetExpenseLines returns the ExpenseLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner {
	if o == nil || IsNil(o.ExpenseLines) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInner
		return ret
	}
	return o.ExpenseLines
}

// GetExpenseLinesOk returns a tuple with the ExpenseLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetExpenseLinesOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool) {
	if o == nil || IsNil(o.ExpenseLines) {
		return nil, false
	}
	return o.ExpenseLines, true
}

// HasExpenseLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasExpenseLines() bool {
	if o != nil && !IsNil(o.ExpenseLines) {
		return true
	}

	return false
}

// SetExpenseLines gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInner and assigns it to the ExpenseLines field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner) {
	o.ExpenseLines = v
}

// GetItemLines returns the ItemLines field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner {
	if o == nil || IsNil(o.ItemLines) {
		var ret []QuickbooksDesktopBillsPostRequestItemLinesInner
		return ret
	}
	return o.ItemLines
}

// GetItemLinesOk returns a tuple with the ItemLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetItemLinesOk() ([]QuickbooksDesktopBillsPostRequestItemLinesInner, bool) {
	if o == nil || IsNil(o.ItemLines) {
		return nil, false
	}
	return o.ItemLines, true
}

// HasItemLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasItemLines() bool {
	if o != nil && !IsNil(o.ItemLines) {
		return true
	}

	return false
}

// SetItemLines gets a reference to the given []QuickbooksDesktopBillsPostRequestItemLinesInner and assigns it to the ItemLines field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner) {
	o.ItemLines = v
}

// GetItemLineGroups returns the ItemLineGroups field value if set, zero value otherwise.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner {
	if o == nil || IsNil(o.ItemLineGroups) {
		var ret []QuickbooksDesktopBillsPostRequestItemLineGroupsInner
		return ret
	}
	return o.ItemLineGroups
}

// GetItemLineGroupsOk returns a tuple with the ItemLineGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) GetItemLineGroupsOk() ([]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool) {
	if o == nil || IsNil(o.ItemLineGroups) {
		return nil, false
	}
	return o.ItemLineGroups, true
}

// HasItemLineGroups returns a boolean if a field has been set.
func (o *QuickbooksDesktopItemReceiptsPostRequest) HasItemLineGroups() bool {
	if o != nil && !IsNil(o.ItemLineGroups) {
		return true
	}

	return false
}

// SetItemLineGroups gets a reference to the given []QuickbooksDesktopBillsPostRequestItemLineGroupsInner and assigns it to the ItemLineGroups field.
func (o *QuickbooksDesktopItemReceiptsPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner) {
	o.ItemLineGroups = v
}

func (o QuickbooksDesktopItemReceiptsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopItemReceiptsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vendorId"] = o.VendorId
	if !IsNil(o.PayablesAccountId) {
		toSerialize["payablesAccountId"] = o.PayablesAccountId
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
	if !IsNil(o.LinkToTransactionIds) {
		toSerialize["linkToTransactionIds"] = o.LinkToTransactionIds
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

func (o *QuickbooksDesktopItemReceiptsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vendorId",
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

	varQuickbooksDesktopItemReceiptsPostRequest := _QuickbooksDesktopItemReceiptsPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopItemReceiptsPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopItemReceiptsPostRequest(varQuickbooksDesktopItemReceiptsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vendorId")
		delete(additionalProperties, "payablesAccountId")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "linkToTransactionIds")
		delete(additionalProperties, "expenseLines")
		delete(additionalProperties, "itemLines")
		delete(additionalProperties, "itemLineGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopItemReceiptsPostRequest struct {
	value *QuickbooksDesktopItemReceiptsPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopItemReceiptsPostRequest) Get() *QuickbooksDesktopItemReceiptsPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopItemReceiptsPostRequest) Set(val *QuickbooksDesktopItemReceiptsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopItemReceiptsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopItemReceiptsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopItemReceiptsPostRequest(val *QuickbooksDesktopItemReceiptsPostRequest) *NullableQuickbooksDesktopItemReceiptsPostRequest {
	return &NullableQuickbooksDesktopItemReceiptsPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopItemReceiptsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopItemReceiptsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


