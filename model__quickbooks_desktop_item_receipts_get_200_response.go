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

// checks if the QuickbooksDesktopItemReceiptsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopItemReceiptsGet200Response{}

// QuickbooksDesktopItemReceiptsGet200Response struct for QuickbooksDesktopItemReceiptsGet200Response
type QuickbooksDesktopItemReceiptsGet200Response struct {
	// The type of object. This value is always `\"list\"`.
	ObjectType string `json:"objectType"`
	// The endpoint URL where this list can be accessed.
	Url string `json:"url"`
	// The array of item receipts.
	Data []QbdItemReceipt `json:"data"`
	// The `nextCursor` is a pagination token returned in the response when you use the `limit` parameter in your request. To retrieve subsequent pages of results, include this token as the value of the `cursor` request parameter in your following API calls.  **NOTE**: The `nextCursor` value remains constant throughout the pagination process for a specific list instance; continue to use the same `nextCursor` token in each request to fetch additional pages.
	NextCursor string `json:"nextCursor"`
	// The number of objects remaining to be fetched.
	RemainingCount float32 `json:"remainingCount"`
	// Indicates whether there are more objects to be fetched.
	HasMore bool `json:"hasMore"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopItemReceiptsGet200Response QuickbooksDesktopItemReceiptsGet200Response

// NewQuickbooksDesktopItemReceiptsGet200Response instantiates a new QuickbooksDesktopItemReceiptsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopItemReceiptsGet200Response(objectType string, url string, data []QbdItemReceipt, nextCursor string, remainingCount float32, hasMore bool) *QuickbooksDesktopItemReceiptsGet200Response {
	this := QuickbooksDesktopItemReceiptsGet200Response{}
	this.ObjectType = objectType
	this.Url = url
	this.Data = data
	this.NextCursor = nextCursor
	this.RemainingCount = remainingCount
	this.HasMore = hasMore
	return &this
}

// NewQuickbooksDesktopItemReceiptsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopItemReceiptsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopItemReceiptsGet200ResponseWithDefaults() *QuickbooksDesktopItemReceiptsGet200Response {
	this := QuickbooksDesktopItemReceiptsGet200Response{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetUrl(v string) {
	o.Url = v
}

// GetData returns the Data field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetData() []QbdItemReceipt {
	if o == nil {
		var ret []QbdItemReceipt
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetDataOk() ([]QbdItemReceipt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetData(v []QbdItemReceipt) {
	o.Data = v
}

// GetNextCursor returns the NextCursor field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetNextCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextCursor, true
}

// SetNextCursor sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetNextCursor(v string) {
	o.NextCursor = v
}

// GetRemainingCount returns the RemainingCount field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetRemainingCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemainingCount
}

// GetRemainingCountOk returns a tuple with the RemainingCount field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetRemainingCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingCount, true
}

// SetRemainingCount sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetRemainingCount(v float32) {
	o.RemainingCount = v
}

// GetHasMore returns the HasMore field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopItemReceiptsGet200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *QuickbooksDesktopItemReceiptsGet200Response) SetHasMore(v bool) {
	o.HasMore = v
}

func (o QuickbooksDesktopItemReceiptsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopItemReceiptsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectType"] = o.ObjectType
	toSerialize["url"] = o.Url
	toSerialize["data"] = o.Data
	toSerialize["nextCursor"] = o.NextCursor
	toSerialize["remainingCount"] = o.RemainingCount
	toSerialize["hasMore"] = o.HasMore

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopItemReceiptsGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectType",
		"url",
		"data",
		"nextCursor",
		"remainingCount",
		"hasMore",
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

	varQuickbooksDesktopItemReceiptsGet200Response := _QuickbooksDesktopItemReceiptsGet200Response{}

	err = json.Unmarshal(data, &varQuickbooksDesktopItemReceiptsGet200Response)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopItemReceiptsGet200Response(varQuickbooksDesktopItemReceiptsGet200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "url")
		delete(additionalProperties, "data")
		delete(additionalProperties, "nextCursor")
		delete(additionalProperties, "remainingCount")
		delete(additionalProperties, "hasMore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopItemReceiptsGet200Response struct {
	value *QuickbooksDesktopItemReceiptsGet200Response
	isSet bool
}

func (v NullableQuickbooksDesktopItemReceiptsGet200Response) Get() *QuickbooksDesktopItemReceiptsGet200Response {
	return v.value
}

func (v *NullableQuickbooksDesktopItemReceiptsGet200Response) Set(val *QuickbooksDesktopItemReceiptsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopItemReceiptsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopItemReceiptsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopItemReceiptsGet200Response(val *QuickbooksDesktopItemReceiptsGet200Response) *NullableQuickbooksDesktopItemReceiptsGet200Response {
	return &NullableQuickbooksDesktopItemReceiptsGet200Response{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopItemReceiptsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopItemReceiptsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


