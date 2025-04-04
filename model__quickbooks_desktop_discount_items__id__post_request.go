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

// checks if the QuickbooksDesktopDiscountItemsIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopDiscountItemsIdPostRequest{}

// QuickbooksDesktopDiscountItemsIdPostRequest struct for QuickbooksDesktopDiscountItemsIdPostRequest
type QuickbooksDesktopDiscountItemsIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the discount item object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The case-insensitive name of this discount item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like `fullName` does. For example, two discount items could both have the `name` \"10% labor discount\", but they could have unique `fullName` values, such as \"Discounts:10% labor discount\" and \"Promotions:10% labor discount\".  Maximum length: 31 characters.
	Name *string `json:"name,omitempty"`
	Barcode *QuickbooksDesktopDiscountItemsPostRequestBarcode `json:"barcode,omitempty"`
	// Indicates whether this discount item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The discount item's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default.
	ClassId *string `json:"classId,omitempty"`
	// The parent discount item one level above this one in the hierarchy. For example, if this discount item has a `fullName` of \"Discounts:10% labor discount\", its parent has a `fullName` of \"Discounts\". If this discount item is at the top level, this field will be `null`.
	ParentId *string `json:"parentId,omitempty"`
	// The discount item's description that will appear on sales forms that include this item.
	Description *string `json:"description,omitempty"`
	// The default sales-tax code for this discount item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The monetary amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A flat rate discount applies to ALL lines recorded above it and distributes the discount amount equally across those lines, which affects tax calculations. For example, a $10 discount applied to a $100 taxable item and $100 non-taxable item would result in a $5 taxable discount and $5 non-taxable discount.
	DiscountRate *string `json:"discountRate,omitempty"`
	// The percentage amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A percentage discount only applies to the line immediately above it, so tax implications only affect that specific line.
	DiscountRatePercent *string `json:"discountRatePercent,omitempty"`
	// The posting account to which transactions involving this discount item are posted for tracking discounts.
	AccountId *string `json:"accountId,omitempty"`
	// When `true`, applies the new account (specified by the `accountId` field) to all existing transactions associated with this discount item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes.
	UpdateExistingTransactionsAccount *bool `json:"updateExistingTransactionsAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopDiscountItemsIdPostRequest QuickbooksDesktopDiscountItemsIdPostRequest

// NewQuickbooksDesktopDiscountItemsIdPostRequest instantiates a new QuickbooksDesktopDiscountItemsIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopDiscountItemsIdPostRequest(revisionNumber string) *QuickbooksDesktopDiscountItemsIdPostRequest {
	this := QuickbooksDesktopDiscountItemsIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopDiscountItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopDiscountItemsIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopDiscountItemsIdPostRequestWithDefaults() *QuickbooksDesktopDiscountItemsIdPostRequest {
	this := QuickbooksDesktopDiscountItemsIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetName(v string) {
	o.Name = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetBarcode() QuickbooksDesktopDiscountItemsPostRequestBarcode {
	if o == nil || IsNil(o.Barcode) {
		var ret QuickbooksDesktopDiscountItemsPostRequestBarcode
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopDiscountItemsPostRequestBarcode, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given QuickbooksDesktopDiscountItemsPostRequestBarcode and assigns it to the Barcode field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetBarcode(v QuickbooksDesktopDiscountItemsPostRequestBarcode) {
	o.Barcode = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetClassId(v string) {
	o.ClassId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRate() string {
	if o == nil || IsNil(o.DiscountRate) {
		var ret string
		return ret
	}
	return *o.DiscountRate
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRateOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountRate) {
		return nil, false
	}
	return o.DiscountRate, true
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDiscountRate() bool {
	if o != nil && !IsNil(o.DiscountRate) {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given string and assigns it to the DiscountRate field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDiscountRate(v string) {
	o.DiscountRate = &v
}

// GetDiscountRatePercent returns the DiscountRatePercent field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRatePercent() string {
	if o == nil || IsNil(o.DiscountRatePercent) {
		var ret string
		return ret
	}
	return *o.DiscountRatePercent
}

// GetDiscountRatePercentOk returns a tuple with the DiscountRatePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRatePercentOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountRatePercent) {
		return nil, false
	}
	return o.DiscountRatePercent, true
}

// HasDiscountRatePercent returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDiscountRatePercent() bool {
	if o != nil && !IsNil(o.DiscountRatePercent) {
		return true
	}

	return false
}

// SetDiscountRatePercent gets a reference to the given string and assigns it to the DiscountRatePercent field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDiscountRatePercent(v string) {
	o.DiscountRatePercent = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetUpdateExistingTransactionsAccount returns the UpdateExistingTransactionsAccount field value if set, zero value otherwise.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetUpdateExistingTransactionsAccount() bool {
	if o == nil || IsNil(o.UpdateExistingTransactionsAccount) {
		var ret bool
		return ret
	}
	return *o.UpdateExistingTransactionsAccount
}

// GetUpdateExistingTransactionsAccountOk returns a tuple with the UpdateExistingTransactionsAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetUpdateExistingTransactionsAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateExistingTransactionsAccount) {
		return nil, false
	}
	return o.UpdateExistingTransactionsAccount, true
}

// HasUpdateExistingTransactionsAccount returns a boolean if a field has been set.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasUpdateExistingTransactionsAccount() bool {
	if o != nil && !IsNil(o.UpdateExistingTransactionsAccount) {
		return true
	}

	return false
}

// SetUpdateExistingTransactionsAccount gets a reference to the given bool and assigns it to the UpdateExistingTransactionsAccount field.
func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetUpdateExistingTransactionsAccount(v bool) {
	o.UpdateExistingTransactionsAccount = &v
}

func (o QuickbooksDesktopDiscountItemsIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopDiscountItemsIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.ClassId) {
		toSerialize["classId"] = o.ClassId
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.DiscountRate) {
		toSerialize["discountRate"] = o.DiscountRate
	}
	if !IsNil(o.DiscountRatePercent) {
		toSerialize["discountRatePercent"] = o.DiscountRatePercent
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.UpdateExistingTransactionsAccount) {
		toSerialize["updateExistingTransactionsAccount"] = o.UpdateExistingTransactionsAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopDiscountItemsIdPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopDiscountItemsIdPostRequest := _QuickbooksDesktopDiscountItemsIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopDiscountItemsIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopDiscountItemsIdPostRequest(varQuickbooksDesktopDiscountItemsIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "name")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "discountRate")
		delete(additionalProperties, "discountRatePercent")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "updateExistingTransactionsAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopDiscountItemsIdPostRequest struct {
	value *QuickbooksDesktopDiscountItemsIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopDiscountItemsIdPostRequest) Get() *QuickbooksDesktopDiscountItemsIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopDiscountItemsIdPostRequest) Set(val *QuickbooksDesktopDiscountItemsIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopDiscountItemsIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopDiscountItemsIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopDiscountItemsIdPostRequest(val *QuickbooksDesktopDiscountItemsIdPostRequest) *NullableQuickbooksDesktopDiscountItemsIdPostRequest {
	return &NullableQuickbooksDesktopDiscountItemsIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopDiscountItemsIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopDiscountItemsIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


