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

// checks if the QuickbooksDesktopDateDrivenTermsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopDateDrivenTermsGet200Response{}

// QuickbooksDesktopDateDrivenTermsGet200Response struct for QuickbooksDesktopDateDrivenTermsGet200Response
type QuickbooksDesktopDateDrivenTermsGet200Response struct {
	// The type of object. This value is always `\"list\"`.
	ObjectType string `json:"objectType"`
	// The endpoint URL where this list can be accessed.
	Url string `json:"url"`
	// The array of date-driven terms.
	Data []QbdDateDrivenTerm `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopDateDrivenTermsGet200Response QuickbooksDesktopDateDrivenTermsGet200Response

// NewQuickbooksDesktopDateDrivenTermsGet200Response instantiates a new QuickbooksDesktopDateDrivenTermsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopDateDrivenTermsGet200Response(objectType string, url string, data []QbdDateDrivenTerm) *QuickbooksDesktopDateDrivenTermsGet200Response {
	this := QuickbooksDesktopDateDrivenTermsGet200Response{}
	this.ObjectType = objectType
	this.Url = url
	this.Data = data
	return &this
}

// NewQuickbooksDesktopDateDrivenTermsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopDateDrivenTermsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopDateDrivenTermsGet200ResponseWithDefaults() *QuickbooksDesktopDateDrivenTermsGet200Response {
	this := QuickbooksDesktopDateDrivenTermsGet200Response{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) SetUrl(v string) {
	o.Url = v
}

// GetData returns the Data field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetData() []QbdDateDrivenTerm {
	if o == nil {
		var ret []QbdDateDrivenTerm
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) GetDataOk() ([]QbdDateDrivenTerm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *QuickbooksDesktopDateDrivenTermsGet200Response) SetData(v []QbdDateDrivenTerm) {
	o.Data = v
}

func (o QuickbooksDesktopDateDrivenTermsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopDateDrivenTermsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectType"] = o.ObjectType
	toSerialize["url"] = o.Url
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopDateDrivenTermsGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectType",
		"url",
		"data",
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

	varQuickbooksDesktopDateDrivenTermsGet200Response := _QuickbooksDesktopDateDrivenTermsGet200Response{}

	err = json.Unmarshal(data, &varQuickbooksDesktopDateDrivenTermsGet200Response)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopDateDrivenTermsGet200Response(varQuickbooksDesktopDateDrivenTermsGet200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "url")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopDateDrivenTermsGet200Response struct {
	value *QuickbooksDesktopDateDrivenTermsGet200Response
	isSet bool
}

func (v NullableQuickbooksDesktopDateDrivenTermsGet200Response) Get() *QuickbooksDesktopDateDrivenTermsGet200Response {
	return v.value
}

func (v *NullableQuickbooksDesktopDateDrivenTermsGet200Response) Set(val *QuickbooksDesktopDateDrivenTermsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopDateDrivenTermsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopDateDrivenTermsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopDateDrivenTermsGet200Response(val *QuickbooksDesktopDateDrivenTermsGet200Response) *NullableQuickbooksDesktopDateDrivenTermsGet200Response {
	return &NullableQuickbooksDesktopDateDrivenTermsGet200Response{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopDateDrivenTermsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopDateDrivenTermsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


