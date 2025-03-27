/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
)

// checks if the QuickbooksDesktopSalesReceiptsPostRequestShippingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopSalesReceiptsPostRequestShippingAddress{}

// QuickbooksDesktopSalesReceiptsPostRequestShippingAddress The sales receipt's shipping address.
type QuickbooksDesktopSalesReceiptsPostRequestShippingAddress struct {
	// The first line of the address (e.g., street, PO Box, or company name).  Maximum length: 41 characters.
	Line1 *string `json:"line1,omitempty"`
	// The second line of the address, if needed (e.g., apartment, suite, unit, or building).  Maximum length: 41 characters.
	Line2 *string `json:"line2,omitempty"`
	// The third line of the address, if needed.  Maximum length: 41 characters.
	Line3 *string `json:"line3,omitempty"`
	// The fourth line of the address, if needed.  Maximum length: 41 characters.
	Line4 *string `json:"line4,omitempty"`
	// The fifth line of the address, if needed.  Maximum length: 41 characters.
	Line5 *string `json:"line5,omitempty"`
	// The city, district, suburb, town, or village name of the address.  Maximum length: 31 characters.
	City *string `json:"city,omitempty"`
	// The state, county, province, or region name of the address.  Maximum length: 21 characters.
	State *string `json:"state,omitempty"`
	// The postal code or ZIP code of the address.  Maximum length: 13 characters.
	PostalCode *string `json:"postalCode,omitempty"`
	// The country name of the address.
	Country *string `json:"country,omitempty"`
	// A note written at the bottom of the address in the form in which it appears, such as the invoice form.
	Note *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopSalesReceiptsPostRequestShippingAddress QuickbooksDesktopSalesReceiptsPostRequestShippingAddress

// NewQuickbooksDesktopSalesReceiptsPostRequestShippingAddress instantiates a new QuickbooksDesktopSalesReceiptsPostRequestShippingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopSalesReceiptsPostRequestShippingAddress() *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress {
	this := QuickbooksDesktopSalesReceiptsPostRequestShippingAddress{}
	return &this
}

// NewQuickbooksDesktopSalesReceiptsPostRequestShippingAddressWithDefaults instantiates a new QuickbooksDesktopSalesReceiptsPostRequestShippingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopSalesReceiptsPostRequestShippingAddressWithDefaults() *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress {
	this := QuickbooksDesktopSalesReceiptsPostRequestShippingAddress{}
	return &this
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasLine1() bool {
	if o != nil && !IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasLine2() bool {
	if o != nil && !IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine3() string {
	if o == nil || IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasLine3() bool {
	if o != nil && !IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetLine3(v string) {
	o.Line3 = &v
}

// GetLine4 returns the Line4 field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine4() string {
	if o == nil || IsNil(o.Line4) {
		var ret string
		return ret
	}
	return *o.Line4
}

// GetLine4Ok returns a tuple with the Line4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.Line4) {
		return nil, false
	}
	return o.Line4, true
}

// HasLine4 returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasLine4() bool {
	if o != nil && !IsNil(o.Line4) {
		return true
	}

	return false
}

// SetLine4 gets a reference to the given string and assigns it to the Line4 field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetLine4(v string) {
	o.Line4 = &v
}

// GetLine5 returns the Line5 field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine5() string {
	if o == nil || IsNil(o.Line5) {
		var ret string
		return ret
	}
	return *o.Line5
}

// GetLine5Ok returns a tuple with the Line5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetLine5Ok() (*string, bool) {
	if o == nil || IsNil(o.Line5) {
		return nil, false
	}
	return o.Line5, true
}

// HasLine5 returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasLine5() bool {
	if o != nil && !IsNil(o.Line5) {
		return true
	}

	return false
}

// SetLine5 gets a reference to the given string and assigns it to the Line5 field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetLine5(v string) {
	o.Line5 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetCountry(v string) {
	o.Country = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) SetNote(v string) {
	o.Note = &v
}

func (o QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !IsNil(o.Line3) {
		toSerialize["line3"] = o.Line3
	}
	if !IsNil(o.Line4) {
		toSerialize["line4"] = o.Line4
	}
	if !IsNil(o.Line5) {
		toSerialize["line5"] = o.Line5
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopSalesReceiptsPostRequestShippingAddress := _QuickbooksDesktopSalesReceiptsPostRequestShippingAddress{}

	err = json.Unmarshal(data, &varQuickbooksDesktopSalesReceiptsPostRequestShippingAddress)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopSalesReceiptsPostRequestShippingAddress(varQuickbooksDesktopSalesReceiptsPostRequestShippingAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "line1")
		delete(additionalProperties, "line2")
		delete(additionalProperties, "line3")
		delete(additionalProperties, "line4")
		delete(additionalProperties, "line5")
		delete(additionalProperties, "city")
		delete(additionalProperties, "state")
		delete(additionalProperties, "postalCode")
		delete(additionalProperties, "country")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress struct {
	value *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress
	isSet bool
}

func (v NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) Get() *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress {
	return v.value
}

func (v *NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) Set(val *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress(val *QuickbooksDesktopSalesReceiptsPostRequestShippingAddress) *NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress {
	return &NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopSalesReceiptsPostRequestShippingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


