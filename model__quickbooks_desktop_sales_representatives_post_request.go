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

// checks if the QuickbooksDesktopSalesRepresentativesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesRepresentativesPostRequest{}

// QuickbooksDesktopSalesRepresentativesPostRequest struct for QuickbooksDesktopSalesRepresentativesPostRequest
type QuickbooksDesktopSalesRepresentativesPostRequest struct {
	// The initials of this sales representative's name.
	Initial string `json:"initial"`
	// Indicates whether this sales representative is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The sales representative's corresponding person entity in QuickBooks, stored as either an employee, vendor, or other-name entry.
	EntityId string `json:"entityId"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesRepresentativesPostRequest QuickbooksDesktopSalesRepresentativesPostRequest

// NewQuickbooksDesktopSalesRepresentativesPostRequest instantiates a new QuickbooksDesktopSalesRepresentativesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesRepresentativesPostRequest(initial string, entityId string) *QuickbooksDesktopSalesRepresentativesPostRequest {
	this := QuickbooksDesktopSalesRepresentativesPostRequest{}
	this.Initial = initial
	var isActive bool = true
	this.IsActive = &isActive
	this.EntityId = entityId
	return &this
}

// NewQuickbooksDesktopSalesRepresentativesPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesRepresentativesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesRepresentativesPostRequestWithDefaults() *QuickbooksDesktopSalesRepresentativesPostRequest {
	this := QuickbooksDesktopSalesRepresentativesPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetInitial returns the Initial field value
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetInitial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Initial
}

// GetInitialOk returns a tuple with the Initial field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetInitialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Initial, true
}

// SetInitial sets field value
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetInitial(v string) {
	o.Initial = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEntityId returns the EntityId field value
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetEntityId(v string) {
	o.EntityId = v
}

func (o QuickbooksDesktopSalesRepresentativesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesRepresentativesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["initial"] = o.Initial
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	toSerialize["entityId"] = o.EntityId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopSalesRepresentativesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"initial",
		"entityId",
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

	varQuickbooksDesktopSalesRepresentativesPostRequest := _QuickbooksDesktopSalesRepresentativesPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesRepresentativesPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesRepresentativesPostRequest(varQuickbooksDesktopSalesRepresentativesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "initial")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "entityId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopSalesRepresentativesPostRequest struct {
	value *QuickbooksDesktopSalesRepresentativesPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopSalesRepresentativesPostRequest) Get() *QuickbooksDesktopSalesRepresentativesPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesRepresentativesPostRequest) Set(val *QuickbooksDesktopSalesRepresentativesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesRepresentativesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesRepresentativesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesRepresentativesPostRequest(val *QuickbooksDesktopSalesRepresentativesPostRequest) *NullableQuickbooksDesktopSalesRepresentativesPostRequest {
	return &NullableQuickbooksDesktopSalesRepresentativesPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesRepresentativesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesRepresentativesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


