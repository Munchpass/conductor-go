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

// checks if the QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner{}

// QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner struct for QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner
type QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner struct {
	// The ID of the note to update.
	Id float32 `json:"id"`
	// The text of this note.
	Note string `json:"note"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner

// NewQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner instantiates a new QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner(id float32, note string) *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner {
	this := QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner{}
	this.Id = id
	this.Note = note
	return &this
}

// NewQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInnerWithDefaults instantiates a new QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInnerWithDefaults() *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner {
	this := QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner{}
	return &this
}

// GetId returns the Id field value
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) SetId(v float32) {
	o.Id = v
}

// GetNote returns the Note field value
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) SetNote(v string) {
	o.Note = v
}

func (o QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["note"] = o.Note

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"note",
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

	varQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner := _QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner(varQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner struct {
	value *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner
	isSet bool
}

func (v NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) Get() *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) Set(val *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner(val *QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) *NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner {
	return &NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


