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

// checks if the QuickbooksDesktopSalesReceiptsIdDelete200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesReceiptsIdDelete200Response{}

// QuickbooksDesktopSalesReceiptsIdDelete200Response struct for QuickbooksDesktopSalesReceiptsIdDelete200Response
type QuickbooksDesktopSalesReceiptsIdDelete200Response struct {
	// The QuickBooks-assigned unique identifier of the deleted sales receipt.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_sales_receipt\"`.
	ObjectType string `json:"objectType"`
	// The case-sensitive user-defined reference number of the deleted sales receipt.
	RefNumber string `json:"refNumber"`
	// Indicates whether the sales receipt was deleted.
	Deleted bool `json:"deleted"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesReceiptsIdDelete200Response QuickbooksDesktopSalesReceiptsIdDelete200Response

// NewQuickbooksDesktopSalesReceiptsIdDelete200Response instantiates a new QuickbooksDesktopSalesReceiptsIdDelete200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesReceiptsIdDelete200Response(id string, objectType string, refNumber string, deleted bool) *QuickbooksDesktopSalesReceiptsIdDelete200Response {
	this := QuickbooksDesktopSalesReceiptsIdDelete200Response{}
	this.Id = id
	this.ObjectType = objectType
	this.RefNumber = refNumber
	this.Deleted = deleted
	return &this
}

// NewQuickbooksDesktopSalesReceiptsIdDelete200ResponseWithDefaults instantiates a new QuickbooksDesktopSalesReceiptsIdDelete200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesReceiptsIdDelete200ResponseWithDefaults() *QuickbooksDesktopSalesReceiptsIdDelete200Response {
	this := QuickbooksDesktopSalesReceiptsIdDelete200Response{}
	return &this
}

// GetId returns the Id field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRefNumber returns the RefNumber field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetRefNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetRefNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefNumber, true
}

// SetRefNumber sets field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) SetRefNumber(v string) {
	o.RefNumber = v
}

// GetDeleted returns the Deleted field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) SetDeleted(v bool) {
	o.Deleted = v
}

func (o QuickbooksDesktopSalesReceiptsIdDelete200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesReceiptsIdDelete200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["refNumber"] = o.RefNumber
	toSerialize["deleted"] = o.Deleted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopSalesReceiptsIdDelete200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"refNumber",
		"deleted",
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

	varQuickbooksDesktopSalesReceiptsIdDelete200Response := _QuickbooksDesktopSalesReceiptsIdDelete200Response{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesReceiptsIdDelete200Response)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesReceiptsIdDelete200Response(varQuickbooksDesktopSalesReceiptsIdDelete200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "deleted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopSalesReceiptsIdDelete200Response struct {
	value *QuickbooksDesktopSalesReceiptsIdDelete200Response
	isSet bool
}

func (v NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) Get() *QuickbooksDesktopSalesReceiptsIdDelete200Response {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) Set(val *QuickbooksDesktopSalesReceiptsIdDelete200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesReceiptsIdDelete200Response(val *QuickbooksDesktopSalesReceiptsIdDelete200Response) *NullableQuickbooksDesktopSalesReceiptsIdDelete200Response {
	return &NullableQuickbooksDesktopSalesReceiptsIdDelete200Response{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesReceiptsIdDelete200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


