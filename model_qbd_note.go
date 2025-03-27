/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QbdNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdNote{}

// QbdNote struct for QbdNote
type QbdNote struct {
	// The auto-incrementing identifier assigned by QuickBooks to this note.
	Id float32 `json:"id"`
	// The date this note was last updated, in ISO 8601 format (YYYY-MM-DD).
	Date NullableString `json:"date"`
	// The text of this note.
	Note string `json:"note"`
}

type _QbdNote QbdNote

// NewQbdNote instantiates a new QbdNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdNote(id float32, date NullableString, note string) *QbdNote {
	this := QbdNote{}
	this.Id = id
	this.Date = date
	this.Note = note
	return &this
}

// NewQbdNoteWithDefaults instantiates a new QbdNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdNoteWithDefaults() *QbdNote {
	this := QbdNote{}
	return &this
}

// GetId returns the Id field value
func (o *QbdNote) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdNote) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdNote) SetId(v float32) {
	o.Id = v
}

// GetDate returns the Date field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdNote) GetDate() string {
	if o == nil || o.Date.Get() == nil {
		var ret string
		return ret
	}

	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdNote) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// SetDate sets field value
func (o *QbdNote) SetDate(v string) {
	o.Date.Set(&v)
}

// GetNote returns the Note field value
func (o *QbdNote) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *QbdNote) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *QbdNote) SetNote(v string) {
	o.Note = v
}

func (o QbdNote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["date"] = o.Date.Get()
	toSerialize["note"] = o.Note
	return toSerialize, nil
}

func (o *QbdNote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"date",
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

	varQbdNote := _QbdNote{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdNote)

	if err != nil {
		return err
	}

	*o = QbdNote(varQbdNote)

	return err
}

type NullableQbdNote struct {
	value *QbdNote
	isSet bool
}

func (v NullableQbdNote) Get() *QbdNote {
	return v.value
}

func (v *NullableQbdNote) Set(val *QbdNote) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdNote) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdNote(val *QbdNote) *NullableQbdNote {
	return &NullableQbdNote{value: val, isSet: true}
}

func (v NullableQbdNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


