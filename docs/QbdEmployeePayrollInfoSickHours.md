# QbdEmployeePayrollInfoSickHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | **string** | The total number of sick hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | 
**AccrualPeriod** | **string** | How frequently the employee&#39;s sick hours are accrued. | 
**HoursAccruedPerPeriod** | **string** | The number of sick hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**MaximumHours** | **string** | The maximum number of sick hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**ResetsHoursEachYear** | **bool** | Indicates whether the employee&#39;s sick hours reset to zero at the beginning of the new year. | 
**HoursUsed** | **string** | The number of sick hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**AccrualStartDate** | **string** | The date the employee&#39;s sick hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | 

## Methods

### NewQbdEmployeePayrollInfoSickHours

`func NewQbdEmployeePayrollInfoSickHours(hoursAvailable string, accrualPeriod string, hoursAccruedPerPeriod string, maximumHours string, resetsHoursEachYear bool, hoursUsed string, accrualStartDate string, ) *QbdEmployeePayrollInfoSickHours`

NewQbdEmployeePayrollInfoSickHours instantiates a new QbdEmployeePayrollInfoSickHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeePayrollInfoSickHoursWithDefaults

`func NewQbdEmployeePayrollInfoSickHoursWithDefaults() *QbdEmployeePayrollInfoSickHours`

NewQbdEmployeePayrollInfoSickHoursWithDefaults instantiates a new QbdEmployeePayrollInfoSickHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QbdEmployeePayrollInfoSickHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.


### GetAccrualPeriod

`func (o *QbdEmployeePayrollInfoSickHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QbdEmployeePayrollInfoSickHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QbdEmployeePayrollInfoSickHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.


### GetHoursAccruedPerPeriod

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QbdEmployeePayrollInfoSickHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.


### GetMaximumHours

`func (o *QbdEmployeePayrollInfoSickHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QbdEmployeePayrollInfoSickHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QbdEmployeePayrollInfoSickHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.


### GetResetsHoursEachYear

`func (o *QbdEmployeePayrollInfoSickHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QbdEmployeePayrollInfoSickHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QbdEmployeePayrollInfoSickHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.


### GetHoursUsed

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QbdEmployeePayrollInfoSickHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QbdEmployeePayrollInfoSickHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.


### GetAccrualStartDate

`func (o *QbdEmployeePayrollInfoSickHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QbdEmployeePayrollInfoSickHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QbdEmployeePayrollInfoSickHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


