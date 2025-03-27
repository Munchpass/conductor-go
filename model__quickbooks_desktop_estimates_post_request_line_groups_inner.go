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

// checks if the QuickbooksDesktopEstimatesPostRequestLineGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopEstimatesPostRequestLineGroupsInner{}

// QuickbooksDesktopEstimatesPostRequestLineGroupsInner struct for QuickbooksDesktopEstimatesPostRequestLineGroupsInner
type QuickbooksDesktopEstimatesPostRequestLineGroupsInner struct {
	// The estimate line group's item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry.
	ItemGroupId string `json:"itemGroupId"`
	// The quantity of the item group associated with this estimate line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group.
	Quantity *float32 `json:"quantity,omitempty"`
	// The unit-of-measure used for the `quantity` in this estimate line group. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// The site location where inventory for the item group associated with this estimate line group is stored.
	InventorySiteId *string `json:"inventorySiteId,omitempty"`
	// The specific location (e.g., bin or shelf) within the inventory site where the item group associated with this estimate line group is stored.
	InventorySiteLocationId *string `json:"inventorySiteLocationId,omitempty"`
	// The custom fields for the estimate line group object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner `json:"customFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopEstimatesPostRequestLineGroupsInner QuickbooksDesktopEstimatesPostRequestLineGroupsInner

// NewQuickbooksDesktopEstimatesPostRequestLineGroupsInner instantiates a new QuickbooksDesktopEstimatesPostRequestLineGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopEstimatesPostRequestLineGroupsInner(itemGroupId string) *QuickbooksDesktopEstimatesPostRequestLineGroupsInner {
	this := QuickbooksDesktopEstimatesPostRequestLineGroupsInner{}
	this.ItemGroupId = itemGroupId
	return &this
}

// NewQuickbooksDesktopEstimatesPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopEstimatesPostRequestLineGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopEstimatesPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopEstimatesPostRequestLineGroupsInner {
	this := QuickbooksDesktopEstimatesPostRequestLineGroupsInner{}
	return &this
}

// GetItemGroupId returns the ItemGroupId field value
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetItemGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemGroupId
}

// GetItemGroupIdOk returns a tuple with the ItemGroupId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemGroupId, true
}

// SetItemGroupId sets field value
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetItemGroupId(v string) {
	o.ItemGroupId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetInventorySiteId returns the InventorySiteId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetInventorySiteId() string {
	if o == nil || IsNil(o.InventorySiteId) {
		var ret string
		return ret
	}
	return *o.InventorySiteId
}

// GetInventorySiteIdOk returns a tuple with the InventorySiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetInventorySiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteId) {
		return nil, false
	}
	return o.InventorySiteId, true
}

// HasInventorySiteId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) HasInventorySiteId() bool {
	if o != nil && !IsNil(o.InventorySiteId) {
		return true
	}

	return false
}

// SetInventorySiteId gets a reference to the given string and assigns it to the InventorySiteId field.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetInventorySiteId(v string) {
	o.InventorySiteId = &v
}

// GetInventorySiteLocationId returns the InventorySiteLocationId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetInventorySiteLocationId() string {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		var ret string
		return ret
	}
	return *o.InventorySiteLocationId
}

// GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetInventorySiteLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		return nil, false
	}
	return o.InventorySiteLocationId, true
}

// HasInventorySiteLocationId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) HasInventorySiteLocationId() bool {
	if o != nil && !IsNil(o.InventorySiteLocationId) {
		return true
	}

	return false
}

// SetInventorySiteLocationId gets a reference to the given string and assigns it to the InventorySiteLocationId field.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetInventorySiteLocationId(v string) {
	o.InventorySiteLocationId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner {
	if o == nil || IsNil(o.CustomFields) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) GetCustomFieldsOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner and assigns it to the CustomFields field.
func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) {
	o.CustomFields = v
}

func (o QuickbooksDesktopEstimatesPostRequestLineGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopEstimatesPostRequestLineGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemGroupId"] = o.ItemGroupId
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	if !IsNil(o.InventorySiteId) {
		toSerialize["inventorySiteId"] = o.InventorySiteId
	}
	if !IsNil(o.InventorySiteLocationId) {
		toSerialize["inventorySiteLocationId"] = o.InventorySiteLocationId
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemGroupId",
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

	varQuickbooksDesktopEstimatesPostRequestLineGroupsInner := _QuickbooksDesktopEstimatesPostRequestLineGroupsInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopEstimatesPostRequestLineGroupsInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopEstimatesPostRequestLineGroupsInner(varQuickbooksDesktopEstimatesPostRequestLineGroupsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemGroupId")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "inventorySiteId")
		delete(additionalProperties, "inventorySiteLocationId")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner struct {
	value *QuickbooksDesktopEstimatesPostRequestLineGroupsInner
	isSet bool
}

func (v NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) Get() *QuickbooksDesktopEstimatesPostRequestLineGroupsInner {
	return v.value
}

func (v *NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) Set(val *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner(val *QuickbooksDesktopEstimatesPostRequestLineGroupsInner) *NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner {
	return &NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopEstimatesPostRequestLineGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


