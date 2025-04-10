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

// checks if the QuickbooksDesktopCustomersPostRequestAdditionalNotesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopCustomersPostRequestAdditionalNotesInner{}

// QuickbooksDesktopCustomersPostRequestAdditionalNotesInner struct for QuickbooksDesktopCustomersPostRequestAdditionalNotesInner
type QuickbooksDesktopCustomersPostRequestAdditionalNotesInner struct {
	// The text of this note.
	Note string `json:"note"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopCustomersPostRequestAdditionalNotesInner QuickbooksDesktopCustomersPostRequestAdditionalNotesInner

// NewQuickbooksDesktopCustomersPostRequestAdditionalNotesInner instantiates a new QuickbooksDesktopCustomersPostRequestAdditionalNotesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopCustomersPostRequestAdditionalNotesInner(note string) *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner {
	this := QuickbooksDesktopCustomersPostRequestAdditionalNotesInner{}
	this.Note = note
	return &this
}

// NewQuickbooksDesktopCustomersPostRequestAdditionalNotesInnerWithDefaults instantiates a new QuickbooksDesktopCustomersPostRequestAdditionalNotesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopCustomersPostRequestAdditionalNotesInnerWithDefaults() *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner {
	this := QuickbooksDesktopCustomersPostRequestAdditionalNotesInner{}
	return &this
}

// GetNote returns the Note field value
func (o *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) SetNote(v string) {
	o.Note = v
}

func (o QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["note"] = o.Note

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varQuickbooksDesktopCustomersPostRequestAdditionalNotesInner := _QuickbooksDesktopCustomersPostRequestAdditionalNotesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopCustomersPostRequestAdditionalNotesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopCustomersPostRequestAdditionalNotesInner(varQuickbooksDesktopCustomersPostRequestAdditionalNotesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner struct {
	value *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner
	isSet bool
}

func (v NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) Get() *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) Set(val *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner(val *QuickbooksDesktopCustomersPostRequestAdditionalNotesInner) *NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner {
	return &NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopCustomersPostRequestAdditionalNotesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


