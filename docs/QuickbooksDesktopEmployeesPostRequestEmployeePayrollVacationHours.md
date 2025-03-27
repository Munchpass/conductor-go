# QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | Pointer to **string** | The total number of vacation hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | [optional] 
**AccrualPeriod** | Pointer to **string** | How frequently the employee&#39;s vacation hours are accrued. | [optional] 
**HoursAccruedPerPeriod** | Pointer to **string** | The number of vacation hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**MaximumHours** | Pointer to **string** | The maximum number of vacation hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**ResetsHoursEachYear** | Pointer to **bool** | Indicates whether the employee&#39;s vacation hours reset to zero at the beginning of the new year. | [optional] 
**HoursUsed** | Pointer to **string** | The number of vacation hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**AccrualStartDate** | Pointer to **string** | The date the employee&#39;s vacation hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours() *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHoursWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHoursWithDefaults() *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHoursWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.

### HasHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasHoursAvailable() bool`

HasHoursAvailable returns a boolean if a field has been set.

### GetAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.

### HasAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasAccrualPeriod() bool`

HasAccrualPeriod returns a boolean if a field has been set.

### GetHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.

### HasHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasHoursAccruedPerPeriod() bool`

HasHoursAccruedPerPeriod returns a boolean if a field has been set.

### GetMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.

### HasMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasMaximumHours() bool`

HasMaximumHours returns a boolean if a field has been set.

### GetResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.

### HasResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasResetsHoursEachYear() bool`

HasResetsHoursEachYear returns a boolean if a field has been set.

### GetHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.

### HasHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasHoursUsed() bool`

HasHoursUsed returns a boolean if a field has been set.

### GetAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.

### HasAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours) HasAccrualStartDate() bool`

HasAccrualStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


