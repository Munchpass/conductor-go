# QbdEmployeePayrollInfoVacationHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | **string** | The total number of vacation hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | 
**AccrualPeriod** | **string** | How frequently the employee&#39;s vacation hours are accrued. | 
**HoursAccruedPerPeriod** | **string** | The number of vacation hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**MaximumHours** | **string** | The maximum number of vacation hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**ResetsHoursEachYear** | **bool** | Indicates whether the employee&#39;s vacation hours reset to zero at the beginning of the new year. | 
**HoursUsed** | **string** | The number of vacation hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | 
**AccrualStartDate** | **string** | The date the employee&#39;s vacation hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | 

## Methods

### NewQbdEmployeePayrollInfoVacationHours

`func NewQbdEmployeePayrollInfoVacationHours(hoursAvailable string, accrualPeriod string, hoursAccruedPerPeriod string, maximumHours string, resetsHoursEachYear bool, hoursUsed string, accrualStartDate string, ) *QbdEmployeePayrollInfoVacationHours`

NewQbdEmployeePayrollInfoVacationHours instantiates a new QbdEmployeePayrollInfoVacationHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeePayrollInfoVacationHoursWithDefaults

`func NewQbdEmployeePayrollInfoVacationHoursWithDefaults() *QbdEmployeePayrollInfoVacationHours`

NewQbdEmployeePayrollInfoVacationHoursWithDefaults instantiates a new QbdEmployeePayrollInfoVacationHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QbdEmployeePayrollInfoVacationHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.


### GetAccrualPeriod

`func (o *QbdEmployeePayrollInfoVacationHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QbdEmployeePayrollInfoVacationHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.


### GetHoursAccruedPerPeriod

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QbdEmployeePayrollInfoVacationHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.


### GetMaximumHours

`func (o *QbdEmployeePayrollInfoVacationHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QbdEmployeePayrollInfoVacationHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.


### GetResetsHoursEachYear

`func (o *QbdEmployeePayrollInfoVacationHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QbdEmployeePayrollInfoVacationHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.


### GetHoursUsed

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QbdEmployeePayrollInfoVacationHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.


### GetAccrualStartDate

`func (o *QbdEmployeePayrollInfoVacationHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QbdEmployeePayrollInfoVacationHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QbdEmployeePayrollInfoVacationHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


