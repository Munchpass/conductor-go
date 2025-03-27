# QbdVacationHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | **NullableString** | The total number of vacation hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | 
**AccrualPeriod** | **NullableString** | How frequently the employee&#39;s vacation hours are accrued. | 
**HoursAccruedPerPeriod** | **NullableString** | The number of vacation hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**MaximumHours** | **NullableString** | The maximum number of vacation hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**ResetsHoursEachYear** | **NullableBool** | Indicates whether the employee&#39;s vacation hours reset to zero at the beginning of the new year. | 
**HoursUsed** | **NullableString** | The number of vacation hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**AccrualStartDate** | **NullableString** | The date the employee&#39;s vacation hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | 

## Methods

### NewQbdVacationHours

`func NewQbdVacationHours(hoursAvailable NullableString, accrualPeriod NullableString, hoursAccruedPerPeriod NullableString, maximumHours NullableString, resetsHoursEachYear NullableBool, hoursUsed NullableString, accrualStartDate NullableString, ) *QbdVacationHours`

NewQbdVacationHours instantiates a new QbdVacationHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdVacationHoursWithDefaults

`func NewQbdVacationHoursWithDefaults() *QbdVacationHours`

NewQbdVacationHoursWithDefaults instantiates a new QbdVacationHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QbdVacationHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QbdVacationHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QbdVacationHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.


### SetHoursAvailableNil

`func (o *QbdVacationHours) SetHoursAvailableNil(b bool)`

 SetHoursAvailableNil sets the value for HoursAvailable to be an explicit nil

### UnsetHoursAvailable
`func (o *QbdVacationHours) UnsetHoursAvailable()`

UnsetHoursAvailable ensures that no value is present for HoursAvailable, not even an explicit nil
### GetAccrualPeriod

`func (o *QbdVacationHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QbdVacationHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QbdVacationHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.


### SetAccrualPeriodNil

`func (o *QbdVacationHours) SetAccrualPeriodNil(b bool)`

 SetAccrualPeriodNil sets the value for AccrualPeriod to be an explicit nil

### UnsetAccrualPeriod
`func (o *QbdVacationHours) UnsetAccrualPeriod()`

UnsetAccrualPeriod ensures that no value is present for AccrualPeriod, not even an explicit nil
### GetHoursAccruedPerPeriod

`func (o *QbdVacationHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QbdVacationHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QbdVacationHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.


### SetHoursAccruedPerPeriodNil

`func (o *QbdVacationHours) SetHoursAccruedPerPeriodNil(b bool)`

 SetHoursAccruedPerPeriodNil sets the value for HoursAccruedPerPeriod to be an explicit nil

### UnsetHoursAccruedPerPeriod
`func (o *QbdVacationHours) UnsetHoursAccruedPerPeriod()`

UnsetHoursAccruedPerPeriod ensures that no value is present for HoursAccruedPerPeriod, not even an explicit nil
### GetMaximumHours

`func (o *QbdVacationHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QbdVacationHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QbdVacationHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.


### SetMaximumHoursNil

`func (o *QbdVacationHours) SetMaximumHoursNil(b bool)`

 SetMaximumHoursNil sets the value for MaximumHours to be an explicit nil

### UnsetMaximumHours
`func (o *QbdVacationHours) UnsetMaximumHours()`

UnsetMaximumHours ensures that no value is present for MaximumHours, not even an explicit nil
### GetResetsHoursEachYear

`func (o *QbdVacationHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QbdVacationHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QbdVacationHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.


### SetResetsHoursEachYearNil

`func (o *QbdVacationHours) SetResetsHoursEachYearNil(b bool)`

 SetResetsHoursEachYearNil sets the value for ResetsHoursEachYear to be an explicit nil

### UnsetResetsHoursEachYear
`func (o *QbdVacationHours) UnsetResetsHoursEachYear()`

UnsetResetsHoursEachYear ensures that no value is present for ResetsHoursEachYear, not even an explicit nil
### GetHoursUsed

`func (o *QbdVacationHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QbdVacationHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QbdVacationHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.


### SetHoursUsedNil

`func (o *QbdVacationHours) SetHoursUsedNil(b bool)`

 SetHoursUsedNil sets the value for HoursUsed to be an explicit nil

### UnsetHoursUsed
`func (o *QbdVacationHours) UnsetHoursUsed()`

UnsetHoursUsed ensures that no value is present for HoursUsed, not even an explicit nil
### GetAccrualStartDate

`func (o *QbdVacationHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QbdVacationHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QbdVacationHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.


### SetAccrualStartDateNil

`func (o *QbdVacationHours) SetAccrualStartDateNil(b bool)`

 SetAccrualStartDateNil sets the value for AccrualStartDate to be an explicit nil

### UnsetAccrualStartDate
`func (o *QbdVacationHours) UnsetAccrualStartDate()`

UnsetAccrualStartDate ensures that no value is present for AccrualStartDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


