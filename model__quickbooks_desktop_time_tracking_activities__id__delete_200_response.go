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

// checks if the QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response{}

// QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response struct for QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response
type QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response struct {
	// The QuickBooks-assigned unique identifier of the deleted time tracking activity.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_time_tracking_activity\"`.
	ObjectType string `json:"objectType"`
	// The case-sensitive user-defined reference number of the deleted time tracking activity.
	RefNumber string `json:"refNumber"`
	// Indicates whether the time tracking activity was deleted.
	Deleted bool `json:"deleted"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response

// NewQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response instantiates a new QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response(id string, objectType string, refNumber string, deleted bool) *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response {
	this := QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response{}
	this.Id = id
	this.ObjectType = objectType
	this.RefNumber = refNumber
	this.Deleted = deleted
	return &this
}

// NewQuickbooksDesktopTimeTrackingActivitiesIdDelete200ResponseWithDefaults instantiates a new QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopTimeTrackingActivitiesIdDelete200ResponseWithDefaults() *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response {
	this := QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response{}
	return &this
}

// GetId returns the Id field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRefNumber returns the RefNumber field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetRefNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetRefNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefNumber, true
}

// SetRefNumber sets field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) SetRefNumber(v string) {
	o.RefNumber = v
}

// GetDeleted returns the Deleted field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) SetDeleted(v bool) {
	o.Deleted = v
}

func (o QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) ToMap() (map[string]interface{}, error) {
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

func (o *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response := _QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response{}

	err = json.Unmarshal(data, &varQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response(varQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response)

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

type NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response struct {
	value *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response
	isSet bool
}

func (v NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) Get() *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response {
	return v.value
}

func (v *NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) Set(val *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response(val *QuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) *NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response {
	return &NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopTimeTrackingActivitiesIdDelete200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


