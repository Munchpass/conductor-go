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

// checks if the QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner{}

// QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner struct for QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner
type QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner struct {
	// The sales order line group's item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry.
	ItemGroupId string `json:"itemGroupId"`
	// The quantity of the item group associated with this sales order line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group.
	Quantity *float32 `json:"quantity,omitempty"`
	// The unit-of-measure used for the `quantity` in this sales order line group. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// The site location where inventory for the item group associated with this sales order line group is stored.
	InventorySiteId *string `json:"inventorySiteId,omitempty"`
	// The specific location (e.g., bin or shelf) within the inventory site where the item group associated with this sales order line group is stored.
	InventorySiteLocationId *string `json:"inventorySiteLocationId,omitempty"`
	// The custom fields for the sales order line group object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner `json:"customFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner

// NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner instantiates a new QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner(itemGroupId string) *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner {
	this := QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner{}
	this.ItemGroupId = itemGroupId
	return &this
}

// NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner {
	this := QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner{}
	return &this
}

// GetItemGroupId returns the ItemGroupId field value
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetItemGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemGroupId
}

// GetItemGroupIdOk returns a tuple with the ItemGroupId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemGroupId, true
}

// SetItemGroupId sets field value
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetItemGroupId(v string) {
	o.ItemGroupId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetInventorySiteId returns the InventorySiteId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteId() string {
	if o == nil || IsNil(o.InventorySiteId) {
		var ret string
		return ret
	}
	return *o.InventorySiteId
}

// GetInventorySiteIdOk returns a tuple with the InventorySiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteId) {
		return nil, false
	}
	return o.InventorySiteId, true
}

// HasInventorySiteId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasInventorySiteId() bool {
	if o != nil && !IsNil(o.InventorySiteId) {
		return true
	}

	return false
}

// SetInventorySiteId gets a reference to the given string and assigns it to the InventorySiteId field.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetInventorySiteId(v string) {
	o.InventorySiteId = &v
}

// GetInventorySiteLocationId returns the InventorySiteLocationId field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteLocationId() string {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		var ret string
		return ret
	}
	return *o.InventorySiteLocationId
}

// GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		return nil, false
	}
	return o.InventorySiteLocationId, true
}

// HasInventorySiteLocationId returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasInventorySiteLocationId() bool {
	if o != nil && !IsNil(o.InventorySiteLocationId) {
		return true
	}

	return false
}

// SetInventorySiteLocationId gets a reference to the given string and assigns it to the InventorySiteLocationId field.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetInventorySiteLocationId(v string) {
	o.InventorySiteLocationId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner {
	if o == nil || IsNil(o.CustomFields) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetCustomFieldsOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner and assigns it to the CustomFields field.
func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) {
	o.CustomFields = v
}

func (o QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner := _QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner(varQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner)

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

type NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner struct {
	value *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner
	isSet bool
}

func (v NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) Get() *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) Set(val *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner(val *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) *NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner {
	return &NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


