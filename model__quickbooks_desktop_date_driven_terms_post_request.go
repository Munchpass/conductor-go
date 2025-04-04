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

// checks if the QuickbooksDesktopDateDrivenTermsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopDateDrivenTermsPostRequest{}

// QuickbooksDesktopDateDrivenTermsPostRequest struct for QuickbooksDesktopDateDrivenTermsPostRequest
type QuickbooksDesktopDateDrivenTermsPostRequest struct {
	// The case-insensitive unique name of this date-driven term, unique across all date-driven terms.  **NOTE**: Date-driven terms do not have a `fullName` field because they are not hierarchical objects, which is why `name` is unique for them but not for objects that have parents.  Maximum length: 31 characters.
	Name string `json:"name"`
	// Indicates whether this date-driven term is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// The day of the month when full payment is due without discount.
	DueDayOfMonth float32 `json:"dueDayOfMonth"`
	// The number of days before `dueDayOfMonth` when an invoice or bill issued within this threshold is considered due the following month. For example, with `dueDayOfMonth` set to 15 and `gracePeriodDays` set to 2, an invoice issued on the 13th would be due on the 15th of the next month, while an invoice issued on the 12th would be due on the 15th of the current month.
	GracePeriodDays *float32 `json:"gracePeriodDays,omitempty"`
	// The day of the month within which payment must be received to qualify for the discount specified by `discountPercentage`.
	DiscountDayOfMonth *float32 `json:"discountDayOfMonth,omitempty"`
	// The discount percentage applied to the payment if received on or before the specified `discountDayOfMonth`. The value is between 0 and 100.
	DiscountPercentage *string `json:"discountPercentage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopDateDrivenTermsPostRequest QuickbooksDesktopDateDrivenTermsPostRequest

// NewQuickbooksDesktopDateDrivenTermsPostRequest instantiates a new QuickbooksDesktopDateDrivenTermsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopDateDrivenTermsPostRequest(name string, dueDayOfMonth float32) *QuickbooksDesktopDateDrivenTermsPostRequest {
	this := QuickbooksDesktopDateDrivenTermsPostRequest{}
	this.Name = name
	var isActive bool = true
	this.IsActive = &isActive
	this.DueDayOfMonth = dueDayOfMonth
	return &this
}

// NewQuickbooksDesktopDateDrivenTermsPostRequestWithDefaults instantiates a new QuickbooksDesktopDateDrivenTermsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopDateDrivenTermsPostRequestWithDefaults() *QuickbooksDesktopDateDrivenTermsPostRequest {
	this := QuickbooksDesktopDateDrivenTermsPostRequest{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDueDayOfMonth returns the DueDayOfMonth field value
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDueDayOfMonth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DueDayOfMonth
}

// GetDueDayOfMonthOk returns a tuple with the DueDayOfMonth field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDueDayOfMonthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDayOfMonth, true
}

// SetDueDayOfMonth sets field value
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDueDayOfMonth(v float32) {
	o.DueDayOfMonth = v
}

// GetGracePeriodDays returns the GracePeriodDays field value if set, zero value otherwise.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetGracePeriodDays() float32 {
	if o == nil || IsNil(o.GracePeriodDays) {
		var ret float32
		return ret
	}
	return *o.GracePeriodDays
}

// GetGracePeriodDaysOk returns a tuple with the GracePeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetGracePeriodDaysOk() (*float32, bool) {
	if o == nil || IsNil(o.GracePeriodDays) {
		return nil, false
	}
	return o.GracePeriodDays, true
}

// HasGracePeriodDays returns a boolean if a field has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasGracePeriodDays() bool {
	if o != nil && !IsNil(o.GracePeriodDays) {
		return true
	}

	return false
}

// SetGracePeriodDays gets a reference to the given float32 and assigns it to the GracePeriodDays field.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetGracePeriodDays(v float32) {
	o.GracePeriodDays = &v
}

// GetDiscountDayOfMonth returns the DiscountDayOfMonth field value if set, zero value otherwise.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountDayOfMonth() float32 {
	if o == nil || IsNil(o.DiscountDayOfMonth) {
		var ret float32
		return ret
	}
	return *o.DiscountDayOfMonth
}

// GetDiscountDayOfMonthOk returns a tuple with the DiscountDayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountDayOfMonthOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountDayOfMonth) {
		return nil, false
	}
	return o.DiscountDayOfMonth, true
}

// HasDiscountDayOfMonth returns a boolean if a field has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasDiscountDayOfMonth() bool {
	if o != nil && !IsNil(o.DiscountDayOfMonth) {
		return true
	}

	return false
}

// SetDiscountDayOfMonth gets a reference to the given float32 and assigns it to the DiscountDayOfMonth field.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDiscountDayOfMonth(v float32) {
	o.DiscountDayOfMonth = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountPercentage() string {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret string
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given string and assigns it to the DiscountPercentage field.
func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDiscountPercentage(v string) {
	o.DiscountPercentage = &v
}

func (o QuickbooksDesktopDateDrivenTermsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopDateDrivenTermsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	toSerialize["dueDayOfMonth"] = o.DueDayOfMonth
	if !IsNil(o.GracePeriodDays) {
		toSerialize["gracePeriodDays"] = o.GracePeriodDays
	}
	if !IsNil(o.DiscountDayOfMonth) {
		toSerialize["discountDayOfMonth"] = o.DiscountDayOfMonth
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopDateDrivenTermsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"dueDayOfMonth",
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

	varQuickbooksDesktopDateDrivenTermsPostRequest := _QuickbooksDesktopDateDrivenTermsPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopDateDrivenTermsPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopDateDrivenTermsPostRequest(varQuickbooksDesktopDateDrivenTermsPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "dueDayOfMonth")
		delete(additionalProperties, "gracePeriodDays")
		delete(additionalProperties, "discountDayOfMonth")
		delete(additionalProperties, "discountPercentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopDateDrivenTermsPostRequest struct {
	value *QuickbooksDesktopDateDrivenTermsPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopDateDrivenTermsPostRequest) Get() *QuickbooksDesktopDateDrivenTermsPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopDateDrivenTermsPostRequest) Set(val *QuickbooksDesktopDateDrivenTermsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopDateDrivenTermsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopDateDrivenTermsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopDateDrivenTermsPostRequest(val *QuickbooksDesktopDateDrivenTermsPostRequest) *NullableQuickbooksDesktopDateDrivenTermsPostRequest {
	return &NullableQuickbooksDesktopDateDrivenTermsPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopDateDrivenTermsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopDateDrivenTermsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


