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

// checks if the QbdTimeTrackingPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdTimeTrackingPreferences{}

// QbdTimeTrackingPreferences struct for QbdTimeTrackingPreferences
type QbdTimeTrackingPreferences struct {
	// The first day of a weekly timesheet period for this company file.
	FirstDayOfWeek string `json:"firstDayOfWeek"`
}

type _QbdTimeTrackingPreferences QbdTimeTrackingPreferences

// NewQbdTimeTrackingPreferences instantiates a new QbdTimeTrackingPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdTimeTrackingPreferences(firstDayOfWeek string) *QbdTimeTrackingPreferences {
	this := QbdTimeTrackingPreferences{}
	this.FirstDayOfWeek = firstDayOfWeek
	return &this
}

// NewQbdTimeTrackingPreferencesWithDefaults instantiates a new QbdTimeTrackingPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdTimeTrackingPreferencesWithDefaults() *QbdTimeTrackingPreferences {
	this := QbdTimeTrackingPreferences{}
	return &this
}

// GetFirstDayOfWeek returns the FirstDayOfWeek field value
func (o *QbdTimeTrackingPreferences) GetFirstDayOfWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstDayOfWeek
}

// GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field value
// and a boolean to check if the value has been set.
func (o *QbdTimeTrackingPreferences) GetFirstDayOfWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstDayOfWeek, true
}

// SetFirstDayOfWeek sets field value
func (o *QbdTimeTrackingPreferences) SetFirstDayOfWeek(v string) {
	o.FirstDayOfWeek = v
}

func (o QbdTimeTrackingPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdTimeTrackingPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstDayOfWeek"] = o.FirstDayOfWeek
	return toSerialize, nil
}

func (o *QbdTimeTrackingPreferences) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"firstDayOfWeek",
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

	varQbdTimeTrackingPreferences := _QbdTimeTrackingPreferences{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdTimeTrackingPreferences)

	if err != nil {
		return err
	}

	*o = QbdTimeTrackingPreferences(varQbdTimeTrackingPreferences)

	return err
}

type NullableQbdTimeTrackingPreferences struct {
	value *QbdTimeTrackingPreferences
	isSet bool
}

func (v NullableQbdTimeTrackingPreferences) Get() *QbdTimeTrackingPreferences {
	return v.value
}

func (v *NullableQbdTimeTrackingPreferences) Set(val *QbdTimeTrackingPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdTimeTrackingPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdTimeTrackingPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdTimeTrackingPreferences(val *QbdTimeTrackingPreferences) *NullableQbdTimeTrackingPreferences {
	return &NullableQbdTimeTrackingPreferences{value: val, isSet: true}
}

func (v NullableQbdTimeTrackingPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdTimeTrackingPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


