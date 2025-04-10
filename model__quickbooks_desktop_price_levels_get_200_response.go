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

// checks if the QuickbooksDesktopPriceLevelsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopPriceLevelsGet200Response{}

// QuickbooksDesktopPriceLevelsGet200Response struct for QuickbooksDesktopPriceLevelsGet200Response
type QuickbooksDesktopPriceLevelsGet200Response struct {
	// The type of object. This value is always `\"list\"`.
	ObjectType string `json:"objectType"`
	// The endpoint URL where this list can be accessed.
	Url string `json:"url"`
	// The array of price levels.
	Data []QbdPriceLevel `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopPriceLevelsGet200Response QuickbooksDesktopPriceLevelsGet200Response

// NewQuickbooksDesktopPriceLevelsGet200Response instantiates a new QuickbooksDesktopPriceLevelsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopPriceLevelsGet200Response(objectType string, url string, data []QbdPriceLevel) *QuickbooksDesktopPriceLevelsGet200Response {
	this := QuickbooksDesktopPriceLevelsGet200Response{}
	this.ObjectType = objectType
	this.Url = url
	this.Data = data
	return &this
}

// NewQuickbooksDesktopPriceLevelsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopPriceLevelsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopPriceLevelsGet200ResponseWithDefaults() *QuickbooksDesktopPriceLevelsGet200Response {
	this := QuickbooksDesktopPriceLevelsGet200Response{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) SetUrl(v string) {
	o.Url = v
}

// GetData returns the Data field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetData() []QbdPriceLevel {
	if o == nil {
		var ret []QbdPriceLevel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopPriceLevelsGet200Response) GetDataOk() ([]QbdPriceLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *QuickbooksDesktopPriceLevelsGet200Response) SetData(v []QbdPriceLevel) {
	o.Data = v
}

func (o QuickbooksDesktopPriceLevelsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopPriceLevelsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectType"] = o.ObjectType
	toSerialize["url"] = o.Url
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopPriceLevelsGet200Response) UnmarshalJSON(data []byte) (err error) {
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

	varQuickbooksDesktopPriceLevelsGet200Response := _QuickbooksDesktopPriceLevelsGet200Response{}

	err = json.Unmarshal(data, &varQuickbooksDesktopPriceLevelsGet200Response)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopPriceLevelsGet200Response(varQuickbooksDesktopPriceLevelsGet200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "url")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopPriceLevelsGet200Response struct {
	value *QuickbooksDesktopPriceLevelsGet200Response
	isSet bool
}

func (v NullableQuickbooksDesktopPriceLevelsGet200Response) Get() *QuickbooksDesktopPriceLevelsGet200Response {
	return v.value
}

func (v *NullableQuickbooksDesktopPriceLevelsGet200Response) Set(val *QuickbooksDesktopPriceLevelsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopPriceLevelsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopPriceLevelsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopPriceLevelsGet200Response(val *QuickbooksDesktopPriceLevelsGet200Response) *NullableQuickbooksDesktopPriceLevelsGet200Response {
	return &NullableQuickbooksDesktopPriceLevelsGet200Response{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopPriceLevelsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopPriceLevelsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


