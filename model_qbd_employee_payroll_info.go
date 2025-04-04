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

// checks if the QbdEmployeePayrollInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdEmployeePayrollInfo{}

// QbdEmployeePayrollInfo struct for QbdEmployeePayrollInfo
type QbdEmployeePayrollInfo struct {
	// How frequently this employee is paid (e.g., weekly, biweekly, monthly). This determines the schedule for generating paychecks.
	PayPeriod string `json:"payPeriod"`
	Class QbdEmployeePayrollInfoClass `json:"class"`
	// The employee's earnings.
	Earnings []QbdEarnings `json:"earnings"`
	// Indicates whether this employee is using time-tracking data to create paychecks.
	UseTimeDataToCreatePaychecks string `json:"useTimeDataToCreatePaychecks"`
	SickHours QbdEmployeePayrollInfoSickHours `json:"sickHours"`
	VacationHours QbdEmployeePayrollInfoVacationHours `json:"vacationHours"`
	AdditionalProperties map[string]interface{}
}

type _QbdEmployeePayrollInfo QbdEmployeePayrollInfo

// NewQbdEmployeePayrollInfo instantiates a new QbdEmployeePayrollInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdEmployeePayrollInfo(payPeriod string, class QbdEmployeePayrollInfoClass, earnings []QbdEarnings, useTimeDataToCreatePaychecks string, sickHours QbdEmployeePayrollInfoSickHours, vacationHours QbdEmployeePayrollInfoVacationHours) *QbdEmployeePayrollInfo {
	this := QbdEmployeePayrollInfo{}
	this.PayPeriod = payPeriod
	this.Class = class
	this.Earnings = earnings
	this.UseTimeDataToCreatePaychecks = useTimeDataToCreatePaychecks
	this.SickHours = sickHours
	this.VacationHours = vacationHours
	return &this
}

// NewQbdEmployeePayrollInfoWithDefaults instantiates a new QbdEmployeePayrollInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdEmployeePayrollInfoWithDefaults() *QbdEmployeePayrollInfo {
	this := QbdEmployeePayrollInfo{}
	return &this
}

// GetPayPeriod returns the PayPeriod field value
func (o *QbdEmployeePayrollInfo) GetPayPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayPeriod
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetPayPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayPeriod, true
}

// SetPayPeriod sets field value
func (o *QbdEmployeePayrollInfo) SetPayPeriod(v string) {
	o.PayPeriod = v
}

// GetClass returns the Class field value
func (o *QbdEmployeePayrollInfo) GetClass() QbdEmployeePayrollInfoClass {
	if o == nil {
		var ret QbdEmployeePayrollInfoClass
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetClassOk() (*QbdEmployeePayrollInfoClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *QbdEmployeePayrollInfo) SetClass(v QbdEmployeePayrollInfoClass) {
	o.Class = v
}

// GetEarnings returns the Earnings field value
func (o *QbdEmployeePayrollInfo) GetEarnings() []QbdEarnings {
	if o == nil {
		var ret []QbdEarnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetEarningsOk() ([]QbdEarnings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Earnings, true
}

// SetEarnings sets field value
func (o *QbdEmployeePayrollInfo) SetEarnings(v []QbdEarnings) {
	o.Earnings = v
}

// GetUseTimeDataToCreatePaychecks returns the UseTimeDataToCreatePaychecks field value
func (o *QbdEmployeePayrollInfo) GetUseTimeDataToCreatePaychecks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UseTimeDataToCreatePaychecks
}

// GetUseTimeDataToCreatePaychecksOk returns a tuple with the UseTimeDataToCreatePaychecks field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetUseTimeDataToCreatePaychecksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseTimeDataToCreatePaychecks, true
}

// SetUseTimeDataToCreatePaychecks sets field value
func (o *QbdEmployeePayrollInfo) SetUseTimeDataToCreatePaychecks(v string) {
	o.UseTimeDataToCreatePaychecks = v
}

// GetSickHours returns the SickHours field value
func (o *QbdEmployeePayrollInfo) GetSickHours() QbdEmployeePayrollInfoSickHours {
	if o == nil {
		var ret QbdEmployeePayrollInfoSickHours
		return ret
	}

	return o.SickHours
}

// GetSickHoursOk returns a tuple with the SickHours field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetSickHoursOk() (*QbdEmployeePayrollInfoSickHours, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SickHours, true
}

// SetSickHours sets field value
func (o *QbdEmployeePayrollInfo) SetSickHours(v QbdEmployeePayrollInfoSickHours) {
	o.SickHours = v
}

// GetVacationHours returns the VacationHours field value
func (o *QbdEmployeePayrollInfo) GetVacationHours() QbdEmployeePayrollInfoVacationHours {
	if o == nil {
		var ret QbdEmployeePayrollInfoVacationHours
		return ret
	}

	return o.VacationHours
}

// GetVacationHoursOk returns a tuple with the VacationHours field value
// and a boolean to check if the value has been set.
func (o *QbdEmployeePayrollInfo) GetVacationHoursOk() (*QbdEmployeePayrollInfoVacationHours, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VacationHours, true
}

// SetVacationHours sets field value
func (o *QbdEmployeePayrollInfo) SetVacationHours(v QbdEmployeePayrollInfoVacationHours) {
	o.VacationHours = v
}

func (o QbdEmployeePayrollInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdEmployeePayrollInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payPeriod"] = o.PayPeriod
	toSerialize["class"] = o.Class
	toSerialize["earnings"] = o.Earnings
	toSerialize["useTimeDataToCreatePaychecks"] = o.UseTimeDataToCreatePaychecks
	toSerialize["sickHours"] = o.SickHours
	toSerialize["vacationHours"] = o.VacationHours

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdEmployeePayrollInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payPeriod",
		"class",
		"earnings",
		"useTimeDataToCreatePaychecks",
		"sickHours",
		"vacationHours",
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

	varQbdEmployeePayrollInfo := _QbdEmployeePayrollInfo{}

	err = json.Unmarshal(data, &varQbdEmployeePayrollInfo)

	if err != nil {
		return err
	}

	*o = QbdEmployeePayrollInfo(varQbdEmployeePayrollInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payPeriod")
		delete(additionalProperties, "class")
		delete(additionalProperties, "earnings")
		delete(additionalProperties, "useTimeDataToCreatePaychecks")
		delete(additionalProperties, "sickHours")
		delete(additionalProperties, "vacationHours")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdEmployeePayrollInfo struct {
	value *QbdEmployeePayrollInfo
	isSet bool
}

func (v NullableQbdEmployeePayrollInfo) Get() *QbdEmployeePayrollInfo {
	return v.value
}

func (v *NullableQbdEmployeePayrollInfo) Set(val *QbdEmployeePayrollInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdEmployeePayrollInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdEmployeePayrollInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdEmployeePayrollInfo(val *QbdEmployeePayrollInfo) *NullableQbdEmployeePayrollInfo {
	return &NullableQbdEmployeePayrollInfo{value: val, isSet: true}
}

func (v NullableQbdEmployeePayrollInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdEmployeePayrollInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


