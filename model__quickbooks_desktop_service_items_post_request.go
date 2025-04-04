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

// checks if the QuickbooksDesktopServiceItemsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopServiceItemsPostRequest{}

// QuickbooksDesktopServiceItemsPostRequest struct for QuickbooksDesktopServiceItemsPostRequest
type QuickbooksDesktopServiceItemsPostRequest struct {
	// The case-insensitive name of this service item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like `fullName` does. For example, two service items could both have the `name` \"Web-Design\", but they could have unique `fullName` values, such as \"Consulting:Web-Design\" and \"Contracting:Web-Design\".  Maximum length: 31 characters.
	Name string `json:"name"`
	Barcode *QuickbooksDesktopServiceItemsPostRequestBarcode `json:"barcode,omitempty"`
	// Indicates whether this service item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The service item's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default.
	ClassId *string `json:"classId,omitempty"`
	// The parent service item one level above this one in the hierarchy. For example, if this service item has a `fullName` of \"Consulting:Web-Design\", its parent has a `fullName` of \"Consulting\". If this service item is at the top level, this field will be `null`.
	ParentId *string `json:"parentId,omitempty"`
	// The unit-of-measure set associated with this service item, which consists of a base unit and related units.
	UnitOfMeasureSetId *string `json:"unitOfMeasureSetId,omitempty"`
	// The default sales-tax code for this service item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	SalesOrPurchaseDetails *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails `json:"salesOrPurchaseDetails,omitempty"`
	SalesAndPurchaseDetails *QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails `json:"salesAndPurchaseDetails,omitempty"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error.
	ExternalId *string `json:"externalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopServiceItemsPostRequest QuickbooksDesktopServiceItemsPostRequest

// NewQuickbooksDesktopServiceItemsPostRequest instantiates a new QuickbooksDesktopServiceItemsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopServiceItemsPostRequest(name string) *QuickbooksDesktopServiceItemsPostRequest {
	this := QuickbooksDesktopServiceItemsPostRequest{}
	this.Name = name
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewQuickbooksDesktopServiceItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopServiceItemsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopServiceItemsPostRequestWithDefaults() *QuickbooksDesktopServiceItemsPostRequest {
	this := QuickbooksDesktopServiceItemsPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *QuickbooksDesktopServiceItemsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QuickbooksDesktopServiceItemsPostRequest) SetName(v string) {
	o.Name = v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetBarcode() QuickbooksDesktopServiceItemsPostRequestBarcode {
	if o == nil || IsNil(o.Barcode) {
		var ret QuickbooksDesktopServiceItemsPostRequestBarcode
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopServiceItemsPostRequestBarcode, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given QuickbooksDesktopServiceItemsPostRequestBarcode and assigns it to the Barcode field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetBarcode(v QuickbooksDesktopServiceItemsPostRequestBarcode) {
	o.Barcode = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetClassId(v string) {
	o.ClassId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetUnitOfMeasureSetId() string {
	if o == nil || IsNil(o.UnitOfMeasureSetId) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasureSetId
}

// GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasureSetId) {
		return nil, false
	}
	return o.UnitOfMeasureSetId, true
}

// HasUnitOfMeasureSetId returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasUnitOfMeasureSetId() bool {
	if o != nil && !IsNil(o.UnitOfMeasureSetId) {
		return true
	}

	return false
}

// SetUnitOfMeasureSetId gets a reference to the given string and assigns it to the UnitOfMeasureSetId field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetUnitOfMeasureSetId(v string) {
	o.UnitOfMeasureSetId = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesOrPurchaseDetails() QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails {
	if o == nil || IsNil(o.SalesOrPurchaseDetails) {
		var ret QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails
		return ret
	}
	return *o.SalesOrPurchaseDetails
}

// GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesOrPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails, bool) {
	if o == nil || IsNil(o.SalesOrPurchaseDetails) {
		return nil, false
	}
	return o.SalesOrPurchaseDetails, true
}

// HasSalesOrPurchaseDetails returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesOrPurchaseDetails() bool {
	if o != nil && !IsNil(o.SalesOrPurchaseDetails) {
		return true
	}

	return false
}

// SetSalesOrPurchaseDetails gets a reference to the given QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails and assigns it to the SalesOrPurchaseDetails field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesOrPurchaseDetails(v QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) {
	o.SalesOrPurchaseDetails = &v
}

// GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesAndPurchaseDetails() QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails {
	if o == nil || IsNil(o.SalesAndPurchaseDetails) {
		var ret QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails
		return ret
	}
	return *o.SalesAndPurchaseDetails
}

// GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesAndPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails, bool) {
	if o == nil || IsNil(o.SalesAndPurchaseDetails) {
		return nil, false
	}
	return o.SalesAndPurchaseDetails, true
}

// HasSalesAndPurchaseDetails returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesAndPurchaseDetails() bool {
	if o != nil && !IsNil(o.SalesAndPurchaseDetails) {
		return true
	}

	return false
}

// SetSalesAndPurchaseDetails gets a reference to the given QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails and assigns it to the SalesAndPurchaseDetails field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesAndPurchaseDetails(v QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails) {
	o.SalesAndPurchaseDetails = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *QuickbooksDesktopServiceItemsPostRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *QuickbooksDesktopServiceItemsPostRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o QuickbooksDesktopServiceItemsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopServiceItemsPostRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.UnitOfMeasureSetId) {
		toSerialize["unitOfMeasureSetId"] = o.UnitOfMeasureSetId
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.SalesOrPurchaseDetails) {
		toSerialize["salesOrPurchaseDetails"] = o.SalesOrPurchaseDetails
	}
	if !IsNil(o.SalesAndPurchaseDetails) {
		toSerialize["salesAndPurchaseDetails"] = o.SalesAndPurchaseDetails
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopServiceItemsPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopServiceItemsPostRequest := _QuickbooksDesktopServiceItemsPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopServiceItemsPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopServiceItemsPostRequest(varQuickbooksDesktopServiceItemsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "unitOfMeasureSetId")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "salesOrPurchaseDetails")
		delete(additionalProperties, "salesAndPurchaseDetails")
		delete(additionalProperties, "externalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopServiceItemsPostRequest struct {
	value *QuickbooksDesktopServiceItemsPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopServiceItemsPostRequest) Get() *QuickbooksDesktopServiceItemsPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopServiceItemsPostRequest) Set(val *QuickbooksDesktopServiceItemsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopServiceItemsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopServiceItemsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopServiceItemsPostRequest(val *QuickbooksDesktopServiceItemsPostRequest) *NullableQuickbooksDesktopServiceItemsPostRequest {
	return &NullableQuickbooksDesktopServiceItemsPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopServiceItemsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopServiceItemsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


