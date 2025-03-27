# QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursAvailable** | Pointer to **string** | The total number of sick hours currently available for the employee to use, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. Defaults to 0. | [optional] 
**AccrualPeriod** | Pointer to **string** | How frequently the employee&#39;s sick hours are accrued. | [optional] 
**HoursAccruedPerPeriod** | Pointer to **string** | The number of sick hours the employee will accrue per accrual period, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**MaximumHours** | Pointer to **string** | The maximum number of sick hours the employee can accrue, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**ResetsHoursEachYear** | Pointer to **bool** | Indicates whether the employee&#39;s sick hours reset to zero at the beginning of the new year. | [optional] 
**HoursUsed** | Pointer to **string** | The number of sick hours the employee has used, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M. | [optional] 
**AccrualStartDate** | Pointer to **string** | The date the employee&#39;s sick hours began to accrue, in ISO 8601 format (YYYY-MM-DD). | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours() *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHoursWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHoursWithDefaults() *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHoursWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursAvailable() string`

GetHoursAvailable returns the HoursAvailable field if non-nil, zero value otherwise.

### GetHoursAvailableOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursAvailableOk() (*string, bool)`

GetHoursAvailableOk returns a tuple with the HoursAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetHoursAvailable(v string)`

SetHoursAvailable sets HoursAvailable field to given value.

### HasHoursAvailable

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasHoursAvailable() bool`

HasHoursAvailable returns a boolean if a field has been set.

### GetAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.

### HasAccrualPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasAccrualPeriod() bool`

HasAccrualPeriod returns a boolean if a field has been set.

### GetHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursAccruedPerPeriod() string`

GetHoursAccruedPerPeriod returns the HoursAccruedPerPeriod field if non-nil, zero value otherwise.

### GetHoursAccruedPerPeriodOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursAccruedPerPeriodOk() (*string, bool)`

GetHoursAccruedPerPeriodOk returns a tuple with the HoursAccruedPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetHoursAccruedPerPeriod(v string)`

SetHoursAccruedPerPeriod sets HoursAccruedPerPeriod field to given value.

### HasHoursAccruedPerPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasHoursAccruedPerPeriod() bool`

HasHoursAccruedPerPeriod returns a boolean if a field has been set.

### GetMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetMaximumHours() string`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetMaximumHoursOk() (*string, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetMaximumHours(v string)`

SetMaximumHours sets MaximumHours field to given value.

### HasMaximumHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasMaximumHours() bool`

HasMaximumHours returns a boolean if a field has been set.

### GetResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetResetsHoursEachYear() bool`

GetResetsHoursEachYear returns the ResetsHoursEachYear field if non-nil, zero value otherwise.

### GetResetsHoursEachYearOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetResetsHoursEachYearOk() (*bool, bool)`

GetResetsHoursEachYearOk returns a tuple with the ResetsHoursEachYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetResetsHoursEachYear(v bool)`

SetResetsHoursEachYear sets ResetsHoursEachYear field to given value.

### HasResetsHoursEachYear

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasResetsHoursEachYear() bool`

HasResetsHoursEachYear returns a boolean if a field has been set.

### GetHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursUsed() string`

GetHoursUsed returns the HoursUsed field if non-nil, zero value otherwise.

### GetHoursUsedOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetHoursUsedOk() (*string, bool)`

GetHoursUsedOk returns a tuple with the HoursUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetHoursUsed(v string)`

SetHoursUsed sets HoursUsed field to given value.

### HasHoursUsed

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasHoursUsed() bool`

HasHoursUsed returns a boolean if a field has been set.

### GetAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetAccrualStartDate() string`

GetAccrualStartDate returns the AccrualStartDate field if non-nil, zero value otherwise.

### GetAccrualStartDateOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) GetAccrualStartDateOk() (*string, bool)`

GetAccrualStartDateOk returns a tuple with the AccrualStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) SetAccrualStartDate(v string)`

SetAccrualStartDate sets AccrualStartDate field to given value.

### HasAccrualStartDate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours) HasAccrualStartDate() bool`

HasAccrualStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


