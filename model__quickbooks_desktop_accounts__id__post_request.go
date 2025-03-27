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

// checks if the QuickbooksDesktopAccountsIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopAccountsIdPostRequest{}

// QuickbooksDesktopAccountsIdPostRequest struct for QuickbooksDesktopAccountsIdPostRequest
type QuickbooksDesktopAccountsIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the account object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The case-insensitive name of this account. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like `fullName` does. For example, two accounts could both have the `name` \"Accounts-Payable\", but they could have unique `fullName` values, such as \"Corporate:Accounts-Payable\" and \"Finance:Accounts-Payable\".  Maximum length: 31 characters.
	Name *string `json:"name,omitempty"`
	// Indicates whether this account is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The parent account one level above this one in the hierarchy. For example, if this account has a `fullName` of \"Corporate:Accounts-Payable\", its parent has a `fullName` of \"Corporate\". If this account is at the top level, this field will be `null`.
	ParentId *string `json:"parentId,omitempty"`
	// The classification of this account, indicating its purpose within the chart of accounts.  **NOTE**: You cannot create an account of type `non_posting` through the API because QuickBooks creates these accounts behind the scenes.
	AccountType *string `json:"accountType,omitempty"`
	// The account's account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \"Use Account Numbers\" preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The bank account number or identifying note for this account. Access to this field may be restricted based on permissions.
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// A description of this account.
	Description *string `json:"description,omitempty"`
	// The amount of money in, or the value of, this account as of `openingBalanceDate`. On a bank statement, this would be the amount of money in the account at the beginning of the statement period.
	OpeningBalance *string `json:"openingBalance,omitempty"`
	// The date of the opening balance of this account, in ISO 8601 format (YYYY-MM-DD).
	OpeningBalanceDate *string `json:"openingBalanceDate,omitempty"`
	// The default sales-tax code for transactions with this account, determining whether the transactions are taxable or non-taxable. This can be overridden at the transaction or transaction-line level.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The identifier of the tax line associated with this account.
	TaxLineId *float32 `json:"taxLineId,omitempty"`
	// The account's currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable.
	CurrencyId *string `json:"currencyId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopAccountsIdPostRequest QuickbooksDesktopAccountsIdPostRequest

// NewQuickbooksDesktopAccountsIdPostRequest instantiates a new QuickbooksDesktopAccountsIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopAccountsIdPostRequest(revisionNumber string) *QuickbooksDesktopAccountsIdPostRequest {
	this := QuickbooksDesktopAccountsIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopAccountsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopAccountsIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopAccountsIdPostRequestWithDefaults() *QuickbooksDesktopAccountsIdPostRequest {
	this := QuickbooksDesktopAccountsIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopAccountsIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopAccountsIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetName(v string) {
	o.Name = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetBankAccountNumber() string {
	if o == nil || IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasBankAccountNumber() bool {
	if o != nil && !IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOpeningBalance returns the OpeningBalance field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetOpeningBalance() string {
	if o == nil || IsNil(o.OpeningBalance) {
		var ret string
		return ret
	}
	return *o.OpeningBalance
}

// GetOpeningBalanceOk returns a tuple with the OpeningBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetOpeningBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.OpeningBalance) {
		return nil, false
	}
	return o.OpeningBalance, true
}

// HasOpeningBalance returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasOpeningBalance() bool {
	if o != nil && !IsNil(o.OpeningBalance) {
		return true
	}

	return false
}

// SetOpeningBalance gets a reference to the given string and assigns it to the OpeningBalance field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetOpeningBalance(v string) {
	o.OpeningBalance = &v
}

// GetOpeningBalanceDate returns the OpeningBalanceDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetOpeningBalanceDate() string {
	if o == nil || IsNil(o.OpeningBalanceDate) {
		var ret string
		return ret
	}
	return *o.OpeningBalanceDate
}

// GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetOpeningBalanceDateOk() (*string, bool) {
	if o == nil || IsNil(o.OpeningBalanceDate) {
		return nil, false
	}
	return o.OpeningBalanceDate, true
}

// HasOpeningBalanceDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasOpeningBalanceDate() bool {
	if o != nil && !IsNil(o.OpeningBalanceDate) {
		return true
	}

	return false
}

// SetOpeningBalanceDate gets a reference to the given string and assigns it to the OpeningBalanceDate field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetOpeningBalanceDate(v string) {
	o.OpeningBalanceDate = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetTaxLineId returns the TaxLineId field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetTaxLineId() float32 {
	if o == nil || IsNil(o.TaxLineId) {
		var ret float32
		return ret
	}
	return *o.TaxLineId
}

// GetTaxLineIdOk returns a tuple with the TaxLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetTaxLineIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxLineId) {
		return nil, false
	}
	return o.TaxLineId, true
}

// HasTaxLineId returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasTaxLineId() bool {
	if o != nil && !IsNil(o.TaxLineId) {
		return true
	}

	return false
}

// SetTaxLineId gets a reference to the given float32 and assigns it to the TaxLineId field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetTaxLineId(v float32) {
	o.TaxLineId = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *QuickbooksDesktopAccountsIdPostRequest) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *QuickbooksDesktopAccountsIdPostRequest) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

func (o QuickbooksDesktopAccountsIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopAccountsIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.OpeningBalance) {
		toSerialize["openingBalance"] = o.OpeningBalance
	}
	if !IsNil(o.OpeningBalanceDate) {
		toSerialize["openingBalanceDate"] = o.OpeningBalanceDate
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.TaxLineId) {
		toSerialize["taxLineId"] = o.TaxLineId
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopAccountsIdPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopAccountsIdPostRequest := _QuickbooksDesktopAccountsIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopAccountsIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopAccountsIdPostRequest(varQuickbooksDesktopAccountsIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "name")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "accountNumber")
		delete(additionalProperties, "bankAccountNumber")
		delete(additionalProperties, "description")
		delete(additionalProperties, "openingBalance")
		delete(additionalProperties, "openingBalanceDate")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "taxLineId")
		delete(additionalProperties, "currencyId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopAccountsIdPostRequest struct {
	value *QuickbooksDesktopAccountsIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopAccountsIdPostRequest) Get() *QuickbooksDesktopAccountsIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopAccountsIdPostRequest) Set(val *QuickbooksDesktopAccountsIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopAccountsIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopAccountsIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopAccountsIdPostRequest(val *QuickbooksDesktopAccountsIdPostRequest) *NullableQuickbooksDesktopAccountsIdPostRequest {
	return &NullableQuickbooksDesktopAccountsIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopAccountsIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopAccountsIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


