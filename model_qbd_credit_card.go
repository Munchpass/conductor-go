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

// checks if the QbdCreditCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdCreditCard{}

// QbdCreditCard struct for QbdCreditCard
type QbdCreditCard struct {
	// The credit card number. Must be masked with lower case \"x\" and no dashes.
	Number NullableString `json:"number"`
	// The month when the credit card expires.
	ExpirationMonth NullableFloat32 `json:"expirationMonth"`
	// The year when the credit card expires.
	ExpirationYear NullableFloat32 `json:"expirationYear"`
	// The cardholder's name on the card.
	Name NullableString `json:"name"`
	// The card's billing address.
	Address NullableString `json:"address"`
	// The card's billing address ZIP or postal code.
	PostalCode NullableString `json:"postalCode"`
}

type _QbdCreditCard QbdCreditCard

// NewQbdCreditCard instantiates a new QbdCreditCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdCreditCard(number NullableString, expirationMonth NullableFloat32, expirationYear NullableFloat32, name NullableString, address NullableString, postalCode NullableString) *QbdCreditCard {
	this := QbdCreditCard{}
	this.Number = number
	this.ExpirationMonth = expirationMonth
	this.ExpirationYear = expirationYear
	this.Name = name
	this.Address = address
	this.PostalCode = postalCode
	return &this
}

// NewQbdCreditCardWithDefaults instantiates a new QbdCreditCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdCreditCardWithDefaults() *QbdCreditCard {
	this := QbdCreditCard{}
	return &this
}

// GetNumber returns the Number field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdCreditCard) GetNumber() string {
	if o == nil || o.Number.Get() == nil {
		var ret string
		return ret
	}

	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// SetNumber sets field value
func (o *QbdCreditCard) SetNumber(v string) {
	o.Number.Set(&v)
}

// GetExpirationMonth returns the ExpirationMonth field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *QbdCreditCard) GetExpirationMonth() float32 {
	if o == nil || o.ExpirationMonth.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ExpirationMonth.Get()
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetExpirationMonthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationMonth.Get(), o.ExpirationMonth.IsSet()
}

// SetExpirationMonth sets field value
func (o *QbdCreditCard) SetExpirationMonth(v float32) {
	o.ExpirationMonth.Set(&v)
}

// GetExpirationYear returns the ExpirationYear field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *QbdCreditCard) GetExpirationYear() float32 {
	if o == nil || o.ExpirationYear.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ExpirationYear.Get()
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetExpirationYearOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationYear.Get(), o.ExpirationYear.IsSet()
}

// SetExpirationYear sets field value
func (o *QbdCreditCard) SetExpirationYear(v float32) {
	o.ExpirationYear.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdCreditCard) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *QbdCreditCard) SetName(v string) {
	o.Name.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdCreditCard) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *QbdCreditCard) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdCreditCard) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdCreditCard) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *QbdCreditCard) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

func (o QbdCreditCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdCreditCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number.Get()
	toSerialize["expirationMonth"] = o.ExpirationMonth.Get()
	toSerialize["expirationYear"] = o.ExpirationYear.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["address"] = o.Address.Get()
	toSerialize["postalCode"] = o.PostalCode.Get()
	return toSerialize, nil
}

func (o *QbdCreditCard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
		"expirationMonth",
		"expirationYear",
		"name",
		"address",
		"postalCode",
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

	varQbdCreditCard := _QbdCreditCard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdCreditCard)

	if err != nil {
		return err
	}

	*o = QbdCreditCard(varQbdCreditCard)

	return err
}

type NullableQbdCreditCard struct {
	value *QbdCreditCard
	isSet bool
}

func (v NullableQbdCreditCard) Get() *QbdCreditCard {
	return v.value
}

func (v *NullableQbdCreditCard) Set(val *QbdCreditCard) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdCreditCard) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdCreditCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdCreditCard(val *QbdCreditCard) *NullableQbdCreditCard {
	return &NullableQbdCreditCard{value: val, isSet: true}
}

func (v NullableQbdCreditCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdCreditCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


