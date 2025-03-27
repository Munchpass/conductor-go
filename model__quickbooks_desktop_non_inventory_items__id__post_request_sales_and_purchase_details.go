/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
)

// checks if the QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails{}

// QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails Details for non-inventory items that are both purchased and sold, such as reimbursable expenses or inventory items that are bought from vendors and sold to customers.  **IMPORTANT**: You cannot specify both `salesAndPurchaseDetails` and `salesOrPurchaseDetails` when modifying a non-inventory item because an item cannot have both configurations.
type QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails struct {
	// The description of this item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers.
	SalesDescription *string `json:"salesDescription,omitempty"`
	// The price at which this item is sold to customers, represented as a decimal string.
	SalesPrice *string `json:"salesPrice,omitempty"`
	// The income account used to track revenue from sales of this item.
	IncomeAccountId *string `json:"incomeAccountId,omitempty"`
	// The description of this item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors.
	PurchaseDescription *string `json:"purchaseDescription,omitempty"`
	// The cost at which this item is purchased from vendors, represented as a decimal string.
	PurchaseCost *string `json:"purchaseCost,omitempty"`
	// The tax code applied to purchases of this item. Applicable in regions where purchase taxes are used, such as Canada or the UK.
	PurchaseTaxCodeId *string `json:"purchaseTaxCodeId,omitempty"`
	// The expense account used to track costs from purchases of this item.
	ExpenseAccountId *string `json:"expenseAccountId,omitempty"`
	// The preferred vendor from whom this item is typically purchased.
	PreferredVendorId *string `json:"preferredVendorId,omitempty"`
	// When `true`, applies the new income account (specified by the `incomeAccountId` field) to all existing transactions that use this item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes.
	UpdateExistingTransactionsIncomeAccount *bool `json:"updateExistingTransactionsIncomeAccount,omitempty"`
	// When `true`, applies the new expense account (specified by the `expenseAccountId` field) to all existing transactions that use this item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes.
	UpdateExistingTransactionsExpenseAccount *bool `json:"updateExistingTransactionsExpenseAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails

// NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails() *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails {
	this := QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails{}
	return &this
}

// NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetailsWithDefaults instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetailsWithDefaults() *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails {
	this := QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails{}
	return &this
}

// GetSalesDescription returns the SalesDescription field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesDescription() string {
	if o == nil || IsNil(o.SalesDescription) {
		var ret string
		return ret
	}
	return *o.SalesDescription
}

// GetSalesDescriptionOk returns a tuple with the SalesDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SalesDescription) {
		return nil, false
	}
	return o.SalesDescription, true
}

// HasSalesDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasSalesDescription() bool {
	if o != nil && !IsNil(o.SalesDescription) {
		return true
	}

	return false
}

// SetSalesDescription gets a reference to the given string and assigns it to the SalesDescription field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetSalesDescription(v string) {
	o.SalesDescription = &v
}

// GetSalesPrice returns the SalesPrice field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesPrice() string {
	if o == nil || IsNil(o.SalesPrice) {
		var ret string
		return ret
	}
	return *o.SalesPrice
}

// GetSalesPriceOk returns a tuple with the SalesPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesPriceOk() (*string, bool) {
	if o == nil || IsNil(o.SalesPrice) {
		return nil, false
	}
	return o.SalesPrice, true
}

// HasSalesPrice returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasSalesPrice() bool {
	if o != nil && !IsNil(o.SalesPrice) {
		return true
	}

	return false
}

// SetSalesPrice gets a reference to the given string and assigns it to the SalesPrice field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetSalesPrice(v string) {
	o.SalesPrice = &v
}

// GetIncomeAccountId returns the IncomeAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetIncomeAccountId() string {
	if o == nil || IsNil(o.IncomeAccountId) {
		var ret string
		return ret
	}
	return *o.IncomeAccountId
}

// GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetIncomeAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.IncomeAccountId) {
		return nil, false
	}
	return o.IncomeAccountId, true
}

// HasIncomeAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasIncomeAccountId() bool {
	if o != nil && !IsNil(o.IncomeAccountId) {
		return true
	}

	return false
}

// SetIncomeAccountId gets a reference to the given string and assigns it to the IncomeAccountId field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetIncomeAccountId(v string) {
	o.IncomeAccountId = &v
}

// GetPurchaseDescription returns the PurchaseDescription field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseDescription() string {
	if o == nil || IsNil(o.PurchaseDescription) {
		var ret string
		return ret
	}
	return *o.PurchaseDescription
}

// GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseDescription) {
		return nil, false
	}
	return o.PurchaseDescription, true
}

// HasPurchaseDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseDescription() bool {
	if o != nil && !IsNil(o.PurchaseDescription) {
		return true
	}

	return false
}

// SetPurchaseDescription gets a reference to the given string and assigns it to the PurchaseDescription field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseDescription(v string) {
	o.PurchaseDescription = &v
}

// GetPurchaseCost returns the PurchaseCost field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseCost() string {
	if o == nil || IsNil(o.PurchaseCost) {
		var ret string
		return ret
	}
	return *o.PurchaseCost
}

// GetPurchaseCostOk returns a tuple with the PurchaseCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseCostOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseCost) {
		return nil, false
	}
	return o.PurchaseCost, true
}

// HasPurchaseCost returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseCost() bool {
	if o != nil && !IsNil(o.PurchaseCost) {
		return true
	}

	return false
}

// SetPurchaseCost gets a reference to the given string and assigns it to the PurchaseCost field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseCost(v string) {
	o.PurchaseCost = &v
}

// GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeId() string {
	if o == nil || IsNil(o.PurchaseTaxCodeId) {
		var ret string
		return ret
	}
	return *o.PurchaseTaxCodeId
}

// GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseTaxCodeId) {
		return nil, false
	}
	return o.PurchaseTaxCodeId, true
}

// HasPurchaseTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseTaxCodeId() bool {
	if o != nil && !IsNil(o.PurchaseTaxCodeId) {
		return true
	}

	return false
}

// SetPurchaseTaxCodeId gets a reference to the given string and assigns it to the PurchaseTaxCodeId field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseTaxCodeId(v string) {
	o.PurchaseTaxCodeId = &v
}

// GetExpenseAccountId returns the ExpenseAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetExpenseAccountId() string {
	if o == nil || IsNil(o.ExpenseAccountId) {
		var ret string
		return ret
	}
	return *o.ExpenseAccountId
}

// GetExpenseAccountIdOk returns a tuple with the ExpenseAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetExpenseAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseAccountId) {
		return nil, false
	}
	return o.ExpenseAccountId, true
}

// HasExpenseAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasExpenseAccountId() bool {
	if o != nil && !IsNil(o.ExpenseAccountId) {
		return true
	}

	return false
}

// SetExpenseAccountId gets a reference to the given string and assigns it to the ExpenseAccountId field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetExpenseAccountId(v string) {
	o.ExpenseAccountId = &v
}

// GetPreferredVendorId returns the PreferredVendorId field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPreferredVendorId() string {
	if o == nil || IsNil(o.PreferredVendorId) {
		var ret string
		return ret
	}
	return *o.PreferredVendorId
}

// GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPreferredVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredVendorId) {
		return nil, false
	}
	return o.PreferredVendorId, true
}

// HasPreferredVendorId returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPreferredVendorId() bool {
	if o != nil && !IsNil(o.PreferredVendorId) {
		return true
	}

	return false
}

// SetPreferredVendorId gets a reference to the given string and assigns it to the PreferredVendorId field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPreferredVendorId(v string) {
	o.PreferredVendorId = &v
}

// GetUpdateExistingTransactionsIncomeAccount returns the UpdateExistingTransactionsIncomeAccount field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsIncomeAccount() bool {
	if o == nil || IsNil(o.UpdateExistingTransactionsIncomeAccount) {
		var ret bool
		return ret
	}
	return *o.UpdateExistingTransactionsIncomeAccount
}

// GetUpdateExistingTransactionsIncomeAccountOk returns a tuple with the UpdateExistingTransactionsIncomeAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsIncomeAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateExistingTransactionsIncomeAccount) {
		return nil, false
	}
	return o.UpdateExistingTransactionsIncomeAccount, true
}

// HasUpdateExistingTransactionsIncomeAccount returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasUpdateExistingTransactionsIncomeAccount() bool {
	if o != nil && !IsNil(o.UpdateExistingTransactionsIncomeAccount) {
		return true
	}

	return false
}

// SetUpdateExistingTransactionsIncomeAccount gets a reference to the given bool and assigns it to the UpdateExistingTransactionsIncomeAccount field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetUpdateExistingTransactionsIncomeAccount(v bool) {
	o.UpdateExistingTransactionsIncomeAccount = &v
}

// GetUpdateExistingTransactionsExpenseAccount returns the UpdateExistingTransactionsExpenseAccount field value if set, zero value otherwise.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsExpenseAccount() bool {
	if o == nil || IsNil(o.UpdateExistingTransactionsExpenseAccount) {
		var ret bool
		return ret
	}
	return *o.UpdateExistingTransactionsExpenseAccount
}

// GetUpdateExistingTransactionsExpenseAccountOk returns a tuple with the UpdateExistingTransactionsExpenseAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsExpenseAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateExistingTransactionsExpenseAccount) {
		return nil, false
	}
	return o.UpdateExistingTransactionsExpenseAccount, true
}

// HasUpdateExistingTransactionsExpenseAccount returns a boolean if a field has been set.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasUpdateExistingTransactionsExpenseAccount() bool {
	if o != nil && !IsNil(o.UpdateExistingTransactionsExpenseAccount) {
		return true
	}

	return false
}

// SetUpdateExistingTransactionsExpenseAccount gets a reference to the given bool and assigns it to the UpdateExistingTransactionsExpenseAccount field.
func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetUpdateExistingTransactionsExpenseAccount(v bool) {
	o.UpdateExistingTransactionsExpenseAccount = &v
}

func (o QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SalesDescription) {
		toSerialize["salesDescription"] = o.SalesDescription
	}
	if !IsNil(o.SalesPrice) {
		toSerialize["salesPrice"] = o.SalesPrice
	}
	if !IsNil(o.IncomeAccountId) {
		toSerialize["incomeAccountId"] = o.IncomeAccountId
	}
	if !IsNil(o.PurchaseDescription) {
		toSerialize["purchaseDescription"] = o.PurchaseDescription
	}
	if !IsNil(o.PurchaseCost) {
		toSerialize["purchaseCost"] = o.PurchaseCost
	}
	if !IsNil(o.PurchaseTaxCodeId) {
		toSerialize["purchaseTaxCodeId"] = o.PurchaseTaxCodeId
	}
	if !IsNil(o.ExpenseAccountId) {
		toSerialize["expenseAccountId"] = o.ExpenseAccountId
	}
	if !IsNil(o.PreferredVendorId) {
		toSerialize["preferredVendorId"] = o.PreferredVendorId
	}
	if !IsNil(o.UpdateExistingTransactionsIncomeAccount) {
		toSerialize["updateExistingTransactionsIncomeAccount"] = o.UpdateExistingTransactionsIncomeAccount
	}
	if !IsNil(o.UpdateExistingTransactionsExpenseAccount) {
		toSerialize["updateExistingTransactionsExpenseAccount"] = o.UpdateExistingTransactionsExpenseAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails := _QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails{}

	err = json.Unmarshal(data, &varQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails(varQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "salesDescription")
		delete(additionalProperties, "salesPrice")
		delete(additionalProperties, "incomeAccountId")
		delete(additionalProperties, "purchaseDescription")
		delete(additionalProperties, "purchaseCost")
		delete(additionalProperties, "purchaseTaxCodeId")
		delete(additionalProperties, "expenseAccountId")
		delete(additionalProperties, "preferredVendorId")
		delete(additionalProperties, "updateExistingTransactionsIncomeAccount")
		delete(additionalProperties, "updateExistingTransactionsExpenseAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails struct {
	value *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails
	isSet bool
}

func (v NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) Get() *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails {
	return v.value
}

func (v *NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) Set(val *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails(val *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) *NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails {
	return &NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


