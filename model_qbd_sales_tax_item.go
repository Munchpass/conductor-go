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

// checks if the QbdSalesTaxItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdSalesTaxItem{}

// QbdSalesTaxItem struct for QbdSalesTaxItem
type QbdSalesTaxItem struct {
	// The unique identifier assigned by QuickBooks to this sales-tax item. This ID is unique across all sales-tax items but not across different QuickBooks object types.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_sales_tax_item\"`.
	ObjectType string `json:"objectType"`
	// The date and time when this sales-tax item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	CreatedAt string `json:"createdAt"`
	// The date and time when this sales-tax item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	UpdatedAt string `json:"updatedAt"`
	// The current QuickBooks-assigned revision number of this sales-tax item object, which changes each time the object is modified. When updating this object, you must provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The case-insensitive unique name of this sales-tax item, unique across all sales-tax items.  **NOTE**: Sales-tax items do not have a `fullName` field because they are not hierarchical objects, which is why `name` is unique for them but not for objects that have parents.
	Name string `json:"name"`
	// The sales-tax item's barcode.
	Barcode string `json:"barcode"`
	// Indicates whether this sales-tax item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive bool `json:"isActive"`
	Class QbdSalesTaxItemClass `json:"class"`
	// The sales-tax item's description that will appear on sales forms that include this item.
	Description string `json:"description"`
	// The tax rate defined by this sales-tax item, represented as a decimal string. For example, \"7.5\" represents a 7.5% tax rate. This rate determines the amount of sales tax applied when this item is used in transactions. If a non-zero `taxRate` is specified, then the `taxVendor` field is required.
	TaxRate string `json:"taxRate"`
	TaxVendor QbdSalesTaxItemTaxVendor `json:"taxVendor"`
	SalesTaxReturnLine QbdSalesTaxItemSalesTaxReturnLine `json:"salesTaxReturnLine"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.
	ExternalId string `json:"externalId"`
	// The custom fields for the sales-tax item object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QbdCustomField `json:"customFields"`
	AdditionalProperties map[string]interface{}
}

type _QbdSalesTaxItem QbdSalesTaxItem

// NewQbdSalesTaxItem instantiates a new QbdSalesTaxItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdSalesTaxItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, barcode string, isActive bool, class QbdSalesTaxItemClass, description string, taxRate string, taxVendor QbdSalesTaxItemTaxVendor, salesTaxReturnLine QbdSalesTaxItemSalesTaxReturnLine, externalId string, customFields []QbdCustomField) *QbdSalesTaxItem {
	this := QbdSalesTaxItem{}
	this.Id = id
	this.ObjectType = objectType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.RevisionNumber = revisionNumber
	this.Name = name
	this.Barcode = barcode
	this.IsActive = isActive
	this.Class = class
	this.Description = description
	this.TaxRate = taxRate
	this.TaxVendor = taxVendor
	this.SalesTaxReturnLine = salesTaxReturnLine
	this.ExternalId = externalId
	this.CustomFields = customFields
	return &this
}

// NewQbdSalesTaxItemWithDefaults instantiates a new QbdSalesTaxItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdSalesTaxItemWithDefaults() *QbdSalesTaxItem {
	this := QbdSalesTaxItem{}
	return &this
}

// GetId returns the Id field value
func (o *QbdSalesTaxItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdSalesTaxItem) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QbdSalesTaxItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QbdSalesTaxItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *QbdSalesTaxItem) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *QbdSalesTaxItem) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *QbdSalesTaxItem) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *QbdSalesTaxItem) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QbdSalesTaxItem) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QbdSalesTaxItem) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetName returns the Name field value
func (o *QbdSalesTaxItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QbdSalesTaxItem) SetName(v string) {
	o.Name = v
}

// GetBarcode returns the Barcode field value
func (o *QbdSalesTaxItem) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetBarcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *QbdSalesTaxItem) SetBarcode(v string) {
	o.Barcode = v
}

// GetIsActive returns the IsActive field value
func (o *QbdSalesTaxItem) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *QbdSalesTaxItem) SetIsActive(v bool) {
	o.IsActive = v
}

// GetClass returns the Class field value
func (o *QbdSalesTaxItem) GetClass() QbdSalesTaxItemClass {
	if o == nil {
		var ret QbdSalesTaxItemClass
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetClassOk() (*QbdSalesTaxItemClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *QbdSalesTaxItem) SetClass(v QbdSalesTaxItemClass) {
	o.Class = v
}

// GetDescription returns the Description field value
func (o *QbdSalesTaxItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *QbdSalesTaxItem) SetDescription(v string) {
	o.Description = v
}

// GetTaxRate returns the TaxRate field value
func (o *QbdSalesTaxItem) GetTaxRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetTaxRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *QbdSalesTaxItem) SetTaxRate(v string) {
	o.TaxRate = v
}

// GetTaxVendor returns the TaxVendor field value
func (o *QbdSalesTaxItem) GetTaxVendor() QbdSalesTaxItemTaxVendor {
	if o == nil {
		var ret QbdSalesTaxItemTaxVendor
		return ret
	}

	return o.TaxVendor
}

// GetTaxVendorOk returns a tuple with the TaxVendor field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetTaxVendorOk() (*QbdSalesTaxItemTaxVendor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxVendor, true
}

// SetTaxVendor sets field value
func (o *QbdSalesTaxItem) SetTaxVendor(v QbdSalesTaxItemTaxVendor) {
	o.TaxVendor = v
}

// GetSalesTaxReturnLine returns the SalesTaxReturnLine field value
func (o *QbdSalesTaxItem) GetSalesTaxReturnLine() QbdSalesTaxItemSalesTaxReturnLine {
	if o == nil {
		var ret QbdSalesTaxItemSalesTaxReturnLine
		return ret
	}

	return o.SalesTaxReturnLine
}

// GetSalesTaxReturnLineOk returns a tuple with the SalesTaxReturnLine field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetSalesTaxReturnLineOk() (*QbdSalesTaxItemSalesTaxReturnLine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesTaxReturnLine, true
}

// SetSalesTaxReturnLine sets field value
func (o *QbdSalesTaxItem) SetSalesTaxReturnLine(v QbdSalesTaxItemSalesTaxReturnLine) {
	o.SalesTaxReturnLine = v
}

// GetExternalId returns the ExternalId field value
func (o *QbdSalesTaxItem) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *QbdSalesTaxItem) SetExternalId(v string) {
	o.ExternalId = v
}

// GetCustomFields returns the CustomFields field value
func (o *QbdSalesTaxItem) GetCustomFields() []QbdCustomField {
	if o == nil {
		var ret []QbdCustomField
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
func (o *QbdSalesTaxItem) GetCustomFieldsOk() ([]QbdCustomField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *QbdSalesTaxItem) SetCustomFields(v []QbdCustomField) {
	o.CustomFields = v
}

func (o QbdSalesTaxItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdSalesTaxItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["revisionNumber"] = o.RevisionNumber
	toSerialize["name"] = o.Name
	toSerialize["barcode"] = o.Barcode
	toSerialize["isActive"] = o.IsActive
	toSerialize["class"] = o.Class
	toSerialize["description"] = o.Description
	toSerialize["taxRate"] = o.TaxRate
	toSerialize["taxVendor"] = o.TaxVendor
	toSerialize["salesTaxReturnLine"] = o.SalesTaxReturnLine
	toSerialize["externalId"] = o.ExternalId
	toSerialize["customFields"] = o.CustomFields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdSalesTaxItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"createdAt",
		"updatedAt",
		"revisionNumber",
		"name",
		"barcode",
		"isActive",
		"class",
		"description",
		"taxRate",
		"taxVendor",
		"salesTaxReturnLine",
		"externalId",
		"customFields",
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

	varQbdSalesTaxItem := _QbdSalesTaxItem{}

	err = json.Unmarshal(data, &varQbdSalesTaxItem)

	if err != nil {
		return err
	}

	*o = QbdSalesTaxItem(varQbdSalesTaxItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "name")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "class")
		delete(additionalProperties, "description")
		delete(additionalProperties, "taxRate")
		delete(additionalProperties, "taxVendor")
		delete(additionalProperties, "salesTaxReturnLine")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdSalesTaxItem struct {
	value *QbdSalesTaxItem
	isSet bool
}

func (v NullableQbdSalesTaxItem) Get() *QbdSalesTaxItem {
	return v.value
}

func (v *NullableQbdSalesTaxItem) Set(val *QbdSalesTaxItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdSalesTaxItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdSalesTaxItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdSalesTaxItem(val *QbdSalesTaxItem) *NullableQbdSalesTaxItem {
	return &NullableQbdSalesTaxItem{value: val, isSet: true}
}

func (v NullableQbdSalesTaxItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdSalesTaxItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


