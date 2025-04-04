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

// checks if the QuickbooksDesktopSalesTaxItemsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesTaxItemsPostRequest{}

// QuickbooksDesktopSalesTaxItemsPostRequest struct for QuickbooksDesktopSalesTaxItemsPostRequest
type QuickbooksDesktopSalesTaxItemsPostRequest struct {
	// The case-insensitive unique name of this sales-tax item, unique across all sales-tax items.  **NOTE**: Sales-tax items do not have a `fullName` field because they are not hierarchical objects, which is why `name` is unique for them but not for objects that have parents.  Maximum length: 31 characters.
	Name string `json:"name"`
	Barcode *QuickbooksDesktopSalesTaxItemsPostRequestBarcode `json:"barcode,omitempty"`
	// Indicates whether this sales-tax item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The sales-tax item's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default.
	ClassId *string `json:"classId,omitempty"`
	// The sales-tax item's description that will appear on sales forms that include this item.
	Description *string `json:"description,omitempty"`
	// The tax rate defined by this sales-tax item, represented as a decimal string. For example, \"7.5\" represents a 7.5% tax rate. This rate determines the amount of sales tax applied when this item is used in transactions. If a non-zero `taxRate` is specified, then the `taxVendor` field is required.
	TaxRate *string `json:"taxRate,omitempty"`
	// The tax agency (vendor) to whom collected sales taxes are owed for this sales-tax item. This field refers to a vendor in QuickBooks that represents the tax authority. If a non-zero `taxRate` is specified, then `taxVendor` is required.
	TaxVendorId *string `json:"taxVendorId,omitempty"`
	// The specific line on the sales tax return form where the tax collected using this sales-tax item should be reported.
	SalesTaxReturnLineId *string `json:"salesTaxReturnLineId,omitempty"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error.
	ExternalId *string `json:"externalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesTaxItemsPostRequest QuickbooksDesktopSalesTaxItemsPostRequest

// NewQuickbooksDesktopSalesTaxItemsPostRequest instantiates a new QuickbooksDesktopSalesTaxItemsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesTaxItemsPostRequest(name string) *QuickbooksDesktopSalesTaxItemsPostRequest {
	this := QuickbooksDesktopSalesTaxItemsPostRequest{}
	this.Name = name
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewQuickbooksDesktopSalesTaxItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxItemsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesTaxItemsPostRequestWithDefaults() *QuickbooksDesktopSalesTaxItemsPostRequest {
	this := QuickbooksDesktopSalesTaxItemsPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetName(v string) {
	o.Name = v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetBarcode() QuickbooksDesktopSalesTaxItemsPostRequestBarcode {
	if o == nil || IsNil(o.Barcode) {
		var ret QuickbooksDesktopSalesTaxItemsPostRequestBarcode
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopSalesTaxItemsPostRequestBarcode, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given QuickbooksDesktopSalesTaxItemsPostRequestBarcode and assigns it to the Barcode field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetBarcode(v QuickbooksDesktopSalesTaxItemsPostRequestBarcode) {
	o.Barcode = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetClassId(v string) {
	o.ClassId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxRate() string {
	if o == nil || IsNil(o.TaxRate) {
		var ret string
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxRateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given string and assigns it to the TaxRate field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetTaxRate(v string) {
	o.TaxRate = &v
}

// GetTaxVendorId returns the TaxVendorId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxVendorId() string {
	if o == nil || IsNil(o.TaxVendorId) {
		var ret string
		return ret
	}
	return *o.TaxVendorId
}

// GetTaxVendorIdOk returns a tuple with the TaxVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxVendorId) {
		return nil, false
	}
	return o.TaxVendorId, true
}

// HasTaxVendorId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasTaxVendorId() bool {
	if o != nil && !IsNil(o.TaxVendorId) {
		return true
	}

	return false
}

// SetTaxVendorId gets a reference to the given string and assigns it to the TaxVendorId field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetTaxVendorId(v string) {
	o.TaxVendorId = &v
}

// GetSalesTaxReturnLineId returns the SalesTaxReturnLineId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetSalesTaxReturnLineId() string {
	if o == nil || IsNil(o.SalesTaxReturnLineId) {
		var ret string
		return ret
	}
	return *o.SalesTaxReturnLineId
}

// GetSalesTaxReturnLineIdOk returns a tuple with the SalesTaxReturnLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetSalesTaxReturnLineIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxReturnLineId) {
		return nil, false
	}
	return o.SalesTaxReturnLineId, true
}

// HasSalesTaxReturnLineId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasSalesTaxReturnLineId() bool {
	if o != nil && !IsNil(o.SalesTaxReturnLineId) {
		return true
	}

	return false
}

// SetSalesTaxReturnLineId gets a reference to the given string and assigns it to the SalesTaxReturnLineId field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetSalesTaxReturnLineId(v string) {
	o.SalesTaxReturnLineId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o QuickbooksDesktopSalesTaxItemsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesTaxItemsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.ClassId) {
		toSerialize["classId"] = o.ClassId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TaxRate) {
		toSerialize["taxRate"] = o.TaxRate
	}
	if !IsNil(o.TaxVendorId) {
		toSerialize["taxVendorId"] = o.TaxVendorId
	}
	if !IsNil(o.SalesTaxReturnLineId) {
		toSerialize["salesTaxReturnLineId"] = o.SalesTaxReturnLineId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopSalesTaxItemsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varQuickbooksDesktopSalesTaxItemsPostRequest := _QuickbooksDesktopSalesTaxItemsPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesTaxItemsPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesTaxItemsPostRequest(varQuickbooksDesktopSalesTaxItemsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "taxRate")
		delete(additionalProperties, "taxVendorId")
		delete(additionalProperties, "salesTaxReturnLineId")
		delete(additionalProperties, "externalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopSalesTaxItemsPostRequest struct {
	value *QuickbooksDesktopSalesTaxItemsPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopSalesTaxItemsPostRequest) Get() *QuickbooksDesktopSalesTaxItemsPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesTaxItemsPostRequest) Set(val *QuickbooksDesktopSalesTaxItemsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesTaxItemsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesTaxItemsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesTaxItemsPostRequest(val *QuickbooksDesktopSalesTaxItemsPostRequest) *NullableQuickbooksDesktopSalesTaxItemsPostRequest {
	return &NullableQuickbooksDesktopSalesTaxItemsPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesTaxItemsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesTaxItemsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


