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

// checks if the QuickbooksDesktopSalesTaxCodesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesTaxCodesPostRequest{}

// QuickbooksDesktopSalesTaxCodesPostRequest struct for QuickbooksDesktopSalesTaxCodesPostRequest
type QuickbooksDesktopSalesTaxCodesPostRequest struct {
	// The case-insensitive unique name of this sales-tax code, unique across all sales-tax codes. This short name will appear on sales forms to identify the tax status of an item.  **NOTE**: Sales-tax codes do not have a `fullName` field because they are not hierarchical objects, which is why `name` is unique for them but not for objects that have parents.  Maximum length: 3 characters.
	Name string `json:"name"`
	// Indicates whether this sales-tax code is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// Indicates whether this sales-tax code is tracking taxable sales. This field cannot be modified once the sales-tax code has been used in a transaction.
	IsTaxable bool `json:"isTaxable"`
	// A description of this sales-tax code.
	Description *string `json:"description,omitempty"`
	// The sales-tax item used to calculate the actual tax amount for this sales-tax code's transactions by applying a specific tax rate collected for a single tax agency. Unlike `salesTaxCode`, which only indicates general taxability, this field drives the actual tax calculation and reporting.
	SalesTaxItemId *string `json:"salesTaxItemId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesTaxCodesPostRequest QuickbooksDesktopSalesTaxCodesPostRequest

// NewQuickbooksDesktopSalesTaxCodesPostRequest instantiates a new QuickbooksDesktopSalesTaxCodesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesTaxCodesPostRequest(name string, isTaxable bool) *QuickbooksDesktopSalesTaxCodesPostRequest {
	this := QuickbooksDesktopSalesTaxCodesPostRequest{}
	this.Name = name
	var isActive bool = true
	this.IsActive = &isActive
	this.IsTaxable = isTaxable
	return &this
}

// NewQuickbooksDesktopSalesTaxCodesPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxCodesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesTaxCodesPostRequestWithDefaults() *QuickbooksDesktopSalesTaxCodesPostRequest {
	this := QuickbooksDesktopSalesTaxCodesPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsTaxable returns the IsTaxable field value
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsTaxable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTaxable
}

// GetIsTaxableOk returns a tuple with the IsTaxable field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsTaxableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTaxable, true
}

// SetIsTaxable sets field value
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetIsTaxable(v bool) {
	o.IsTaxable = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSalesTaxItemId returns the SalesTaxItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetSalesTaxItemId() string {
	if o == nil || IsNil(o.SalesTaxItemId) {
		var ret string
		return ret
	}
	return *o.SalesTaxItemId
}

// GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetSalesTaxItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxItemId) {
		return nil, false
	}
	return o.SalesTaxItemId, true
}

// HasSalesTaxItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasSalesTaxItemId() bool {
	if o != nil && !IsNil(o.SalesTaxItemId) {
		return true
	}

	return false
}

// SetSalesTaxItemId gets a reference to the given string and assigns it to the SalesTaxItemId field.
func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetSalesTaxItemId(v string) {
	o.SalesTaxItemId = &v
}

func (o QuickbooksDesktopSalesTaxCodesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesTaxCodesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	toSerialize["isTaxable"] = o.IsTaxable
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SalesTaxItemId) {
		toSerialize["salesTaxItemId"] = o.SalesTaxItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopSalesTaxCodesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"isTaxable",
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

	varQuickbooksDesktopSalesTaxCodesPostRequest := _QuickbooksDesktopSalesTaxCodesPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesTaxCodesPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesTaxCodesPostRequest(varQuickbooksDesktopSalesTaxCodesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "isTaxable")
		delete(additionalProperties, "description")
		delete(additionalProperties, "salesTaxItemId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopSalesTaxCodesPostRequest struct {
	value *QuickbooksDesktopSalesTaxCodesPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopSalesTaxCodesPostRequest) Get() *QuickbooksDesktopSalesTaxCodesPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesTaxCodesPostRequest) Set(val *QuickbooksDesktopSalesTaxCodesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesTaxCodesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesTaxCodesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesTaxCodesPostRequest(val *QuickbooksDesktopSalesTaxCodesPostRequest) *NullableQuickbooksDesktopSalesTaxCodesPostRequest {
	return &NullableQuickbooksDesktopSalesTaxCodesPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesTaxCodesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesTaxCodesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


