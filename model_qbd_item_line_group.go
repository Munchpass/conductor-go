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

// checks if the QbdItemLineGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdItemLineGroup{}

// QbdItemLineGroup struct for QbdItemLineGroup
type QbdItemLineGroup struct {
	// The unique identifier assigned by QuickBooks to this item line group. This ID is unique across all transaction line types.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_item_line_group\"`.
	ObjectType string `json:"objectType"`
	ItemGroup QbdItemLineGroupItemGroup `json:"itemGroup"`
	// A description of this item line group.
	Description string `json:"description"`
	// The quantity of the item group associated with this item line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group.
	Quantity float32 `json:"quantity"`
	// The unit-of-measure used for the `quantity` in this item line group. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure string `json:"unitOfMeasure"`
	OverrideUnitOfMeasureSet QbdItemLineGroupOverrideUnitOfMeasureSet `json:"overrideUnitOfMeasureSet"`
	// The total monetary amount of this item line group, equivalent to the sum of the amounts in `lines`, represented as a decimal string.
	TotalAmount string `json:"totalAmount"`
	// The item line group's item lines, each representing the purchase of a specific item or service.
	ItemLines []QbdItemLine `json:"itemLines"`
	// The custom fields for the item line group object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QbdCustomField `json:"customFields"`
	AdditionalProperties map[string]interface{}
}

type _QbdItemLineGroup QbdItemLineGroup

// NewQbdItemLineGroup instantiates a new QbdItemLineGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdItemLineGroup(id string, objectType string, itemGroup QbdItemLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdItemLineGroupOverrideUnitOfMeasureSet, totalAmount string, itemLines []QbdItemLine, customFields []QbdCustomField) *QbdItemLineGroup {
	this := QbdItemLineGroup{}
	this.Id = id
	this.ObjectType = objectType
	this.ItemGroup = itemGroup
	this.Description = description
	this.Quantity = quantity
	this.UnitOfMeasure = unitOfMeasure
	this.OverrideUnitOfMeasureSet = overrideUnitOfMeasureSet
	this.TotalAmount = totalAmount
	this.ItemLines = itemLines
	this.CustomFields = customFields
	return &this
}

// NewQbdItemLineGroupWithDefaults instantiates a new QbdItemLineGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdItemLineGroupWithDefaults() *QbdItemLineGroup {
	this := QbdItemLineGroup{}
	return &this
}

// GetId returns the Id field value
func (o *QbdItemLineGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdItemLineGroup) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QbdItemLineGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QbdItemLineGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetItemGroup returns the ItemGroup field value
func (o *QbdItemLineGroup) GetItemGroup() QbdItemLineGroupItemGroup {
	if o == nil {
		var ret QbdItemLineGroupItemGroup
		return ret
	}

	return o.ItemGroup
}

// GetItemGroupOk returns a tuple with the ItemGroup field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetItemGroupOk() (*QbdItemLineGroupItemGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemGroup, true
}

// SetItemGroup sets field value
func (o *QbdItemLineGroup) SetItemGroup(v QbdItemLineGroupItemGroup) {
	o.ItemGroup = v
}

// GetDescription returns the Description field value
func (o *QbdItemLineGroup) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *QbdItemLineGroup) SetDescription(v string) {
	o.Description = v
}

// GetQuantity returns the Quantity field value
func (o *QbdItemLineGroup) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *QbdItemLineGroup) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *QbdItemLineGroup) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *QbdItemLineGroup) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field value
func (o *QbdItemLineGroup) GetOverrideUnitOfMeasureSet() QbdItemLineGroupOverrideUnitOfMeasureSet {
	if o == nil {
		var ret QbdItemLineGroupOverrideUnitOfMeasureSet
		return ret
	}

	return o.OverrideUnitOfMeasureSet
}

// GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdItemLineGroupOverrideUnitOfMeasureSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideUnitOfMeasureSet, true
}

// SetOverrideUnitOfMeasureSet sets field value
func (o *QbdItemLineGroup) SetOverrideUnitOfMeasureSet(v QbdItemLineGroupOverrideUnitOfMeasureSet) {
	o.OverrideUnitOfMeasureSet = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *QbdItemLineGroup) GetTotalAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetTotalAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *QbdItemLineGroup) SetTotalAmount(v string) {
	o.TotalAmount = v
}

// GetItemLines returns the ItemLines field value
func (o *QbdItemLineGroup) GetItemLines() []QbdItemLine {
	if o == nil {
		var ret []QbdItemLine
		return ret
	}

	return o.ItemLines
}

// GetItemLinesOk returns a tuple with the ItemLines field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetItemLinesOk() ([]QbdItemLine, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemLines, true
}

// SetItemLines sets field value
func (o *QbdItemLineGroup) SetItemLines(v []QbdItemLine) {
	o.ItemLines = v
}

// GetCustomFields returns the CustomFields field value
func (o *QbdItemLineGroup) GetCustomFields() []QbdCustomField {
	if o == nil {
		var ret []QbdCustomField
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
func (o *QbdItemLineGroup) GetCustomFieldsOk() ([]QbdCustomField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *QbdItemLineGroup) SetCustomFields(v []QbdCustomField) {
	o.CustomFields = v
}

func (o QbdItemLineGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdItemLineGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["itemGroup"] = o.ItemGroup
	toSerialize["description"] = o.Description
	toSerialize["quantity"] = o.Quantity
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	toSerialize["overrideUnitOfMeasureSet"] = o.OverrideUnitOfMeasureSet
	toSerialize["totalAmount"] = o.TotalAmount
	toSerialize["itemLines"] = o.ItemLines
	toSerialize["customFields"] = o.CustomFields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdItemLineGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"itemGroup",
		"description",
		"quantity",
		"unitOfMeasure",
		"overrideUnitOfMeasureSet",
		"totalAmount",
		"itemLines",
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

	varQbdItemLineGroup := _QbdItemLineGroup{}

	err = json.Unmarshal(data, &varQbdItemLineGroup)

	if err != nil {
		return err
	}

	*o = QbdItemLineGroup(varQbdItemLineGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "itemGroup")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "overrideUnitOfMeasureSet")
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "itemLines")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdItemLineGroup struct {
	value *QbdItemLineGroup
	isSet bool
}

func (v NullableQbdItemLineGroup) Get() *QbdItemLineGroup {
	return v.value
}

func (v *NullableQbdItemLineGroup) Set(val *QbdItemLineGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdItemLineGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdItemLineGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdItemLineGroup(val *QbdItemLineGroup) *NullableQbdItemLineGroup {
	return &NullableQbdItemLineGroup{value: val, isSet: true}
}

func (v NullableQbdItemLineGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdItemLineGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


