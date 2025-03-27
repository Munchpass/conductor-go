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

// checks if the QuickbooksDesktopClassesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopClassesPostRequest{}

// QuickbooksDesktopClassesPostRequest struct for QuickbooksDesktopClassesPostRequest
type QuickbooksDesktopClassesPostRequest struct {
	// The case-insensitive name of this class. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like `fullName` does. For example, two classes could both have the `name` \"Marketing\", but they could have unique `fullName` values, such as \"Department:Marketing\" and \"Internal:Marketing\".  Maximum length: 31 characters.
	Name string `json:"name"`
	// Indicates whether this class is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The parent class one level above this one in the hierarchy. For example, if this class has a `fullName` of \"Department:Marketing\", its parent has a `fullName` of \"Department\". If this class is at the top level, this field will be `null`.
	ParentId *string `json:"parentId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopClassesPostRequest QuickbooksDesktopClassesPostRequest

// NewQuickbooksDesktopClassesPostRequest instantiates a new QuickbooksDesktopClassesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopClassesPostRequest(name string) *QuickbooksDesktopClassesPostRequest {
	this := QuickbooksDesktopClassesPostRequest{}
	this.Name = name
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewQuickbooksDesktopClassesPostRequestWithDefaults instantiates a new QuickbooksDesktopClassesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopClassesPostRequestWithDefaults() *QuickbooksDesktopClassesPostRequest {
	this := QuickbooksDesktopClassesPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *QuickbooksDesktopClassesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopClassesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QuickbooksDesktopClassesPostRequest) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopClassesPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopClassesPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopClassesPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopClassesPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *QuickbooksDesktopClassesPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopClassesPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *QuickbooksDesktopClassesPostRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *QuickbooksDesktopClassesPostRequest) SetParentId(v string) {
	o.ParentId = &v
}

func (o QuickbooksDesktopClassesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopClassesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopClassesPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopClassesPostRequest := _QuickbooksDesktopClassesPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopClassesPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopClassesPostRequest(varQuickbooksDesktopClassesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "parentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopClassesPostRequest struct {
	value *QuickbooksDesktopClassesPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopClassesPostRequest) Get() *QuickbooksDesktopClassesPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopClassesPostRequest) Set(val *QuickbooksDesktopClassesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopClassesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopClassesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopClassesPostRequest(val *QuickbooksDesktopClassesPostRequest) *NullableQuickbooksDesktopClassesPostRequest {
	return &NullableQuickbooksDesktopClassesPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopClassesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopClassesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


