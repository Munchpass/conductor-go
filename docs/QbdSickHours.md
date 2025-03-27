# QbdSickHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | **NullableString** | The total number of sick hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | 
**AccrualPeriod** | **NullableString** | How frequently the employee&#39;s sick hours are accrued. | 
**HoursAccruedPerPeriod** | **NullableString** | The number of sick hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**MaximumHours** | **NullableString** | The maximum number of sick hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**ResetsHoursEachYear** | **NullableBool** | Indicates whether the employee&#39;s sick hours reset to zero at the beginning of the new year. | 
**HoursUsed** | **NullableString** | The number of sick hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**AccrualStartDate** | **NullableString** | The date the employee&#39;s sick hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | 

## Methods

### NewQbdSickHours

`func NewQbdSickHours(hoursAvailable NullableString, accrualPeriod NullableString, hoursAccruedPerPeriod NullableString, maximumHours NullableString, resetsHoursEachYear NullableBool, hoursUsed NullableString, accrualStartDate NullableString, ) *QbdSickHours`

NewQbdSickHours instantiates a new QbdSickHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSickHoursWithDefaults

`func NewQbdSickHoursWithDefaults() *QbdSickHours`

NewQbdSickHoursWithDefaults instantiates a new QbdSickHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QbdSickHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QbdSickHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QbdSickHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.


### SetHoursAvailableNil

`func (o *QbdSickHours) SetHoursAvailableNil(b bool)`

 SetHoursAvailableNil sets the value for HoursAvailable to be an explicit nil

### UnsetHoursAvailable
`func (o *QbdSickHours) UnsetHoursAvailable()`

UnsetHoursAvailable ensures that no value is present for HoursAvailable, not even an explicit nil
### GetAccrualPeriod

`func (o *QbdSickHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QbdSickHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QbdSickHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.


### SetAccrualPeriodNil

`func (o *QbdSickHours) SetAccrualPeriodNil(b bool)`

 SetAccrualPeriodNil sets the value for AccrualPeriod to be an explicit nil

### UnsetAccrualPeriod
`func (o *QbdSickHours) UnsetAccrualPeriod()`

UnsetAccrualPeriod ensures that no value is present for AccrualPeriod, not even an explicit nil
### GetHoursAccruedPerPeriod

`func (o *QbdSickHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QbdSickHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QbdSickHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.


### SetHoursAccruedPerPeriodNil

`func (o *QbdSickHours) SetHoursAccruedPerPeriodNil(b bool)`

 SetHoursAccruedPerPeriodNil sets the value for HoursAccruedPerPeriod to be an explicit nil

### UnsetHoursAccruedPerPeriod
`func (o *QbdSickHours) UnsetHoursAccruedPerPeriod()`

UnsetHoursAccruedPerPeriod ensures that no value is present for HoursAccruedPerPeriod, not even an explicit nil
### GetMaximumHours

`func (o *QbdSickHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QbdSickHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QbdSickHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.


### SetMaximumHoursNil

`func (o *QbdSickHours) SetMaximumHoursNil(b bool)`

 SetMaximumHoursNil sets the value for MaximumHours to be an explicit nil

### UnsetMaximumHours
`func (o *QbdSickHours) UnsetMaximumHours()`

UnsetMaximumHours ensures that no value is present for MaximumHours, not even an explicit nil
### GetResetsHoursEachYear

`func (o *QbdSickHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QbdSickHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QbdSickHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.


### SetResetsHoursEachYearNil

`func (o *QbdSickHours) SetResetsHoursEachYearNil(b bool)`

 SetResetsHoursEachYearNil sets the value for ResetsHoursEachYear to be an explicit nil

### UnsetResetsHoursEachYear
`func (o *QbdSickHours) UnsetResetsHoursEachYear()`

UnsetResetsHoursEachYear ensures that no value is present for ResetsHoursEachYear, not even an explicit nil
### GetHoursUsed

`func (o *QbdSickHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QbdSickHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QbdSickHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.


### SetHoursUsedNil

`func (o *QbdSickHours) SetHoursUsedNil(b bool)`

 SetHoursUsedNil sets the value for HoursUsed to be an explicit nil

### UnsetHoursUsed
`func (o *QbdSickHours) UnsetHoursUsed()`

UnsetHoursUsed ensures that no value is present for HoursUsed, not even an explicit nil
### GetAccrualStartDate

`func (o *QbdSickHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QbdSickHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QbdSickHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.


### SetAccrualStartDateNil

`func (o *QbdSickHours) SetAccrualStartDateNil(b bool)`

 SetAccrualStartDateNil sets the value for AccrualStartDate to be an explicit nil

### UnsetAccrualStartDate
`func (o *QbdSickHours) UnsetAccrualStartDate()`

UnsetAccrualStartDate ensures that no value is present for AccrualStartDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


