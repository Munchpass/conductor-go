# QuickbooksDesktopEmployeesPostRequestEmployeePayroll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPeriod** | Pointer to **string** | How frequently this employee is paid (e.g., weekly, biweekly, monthly). This determines the schedule for generating paychecks. | [optional] 
**ClassId** | Pointer to **string** | The employee&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Earnings** | Pointer to [**[]QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner.md) | The employee&#39;s earnings. | [optional] 
**UseTimeDataToCreatePaychecks** | Pointer to **string** | Indicates whether this employee is using time-tracking data to create paychecks. | [optional] 
**SickHours** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours.md) |  | [optional] 
**VacationHours** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayroll

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayroll() *QuickbooksDesktopEmployeesPostRequestEmployeePayroll`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayroll instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayroll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollWithDefaults() *QuickbooksDesktopEmployeesPostRequestEmployeePayroll`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayroll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetEarnings

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetEarnings() []QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetEarningsOk() (*[]QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetEarnings(v []QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetUseTimeDataToCreatePaychecks() string`

GetUseTimeDataToCreatePaychecks returns the UseTimeDataToCreatePaychecks field if non-nil, zero value otherwise.

### GetUseTimeDataToCreatePaychecksOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetUseTimeDataToCreatePaychecksOk() (*string, bool)`

GetUseTimeDataToCreatePaychecksOk returns a tuple with the UseTimeDataToCreatePaychecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetUseTimeDataToCreatePaychecks(v string)`

SetUseTimeDataToCreatePaychecks sets UseTimeDataToCreatePaychecks field to given value.

### HasUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasUseTimeDataToCreatePaychecks() bool`

HasUseTimeDataToCreatePaychecks returns a boolean if a field has been set.

### GetSickHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetSickHours() QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours`

GetSickHours returns the SickHours field if non-nil, zero value otherwise.

### GetSickHoursOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetSickHoursOk() (*QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours, bool)`

GetSickHoursOk returns a tuple with the SickHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetSickHours(v QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours)`

SetSickHours sets SickHours field to given value.

### HasSickHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasSickHours() bool`

HasSickHours returns a boolean if a field has been set.

### GetVacationHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetVacationHours() QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours`

GetVacationHours returns the VacationHours field if non-nil, zero value otherwise.

### GetVacationHoursOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) GetVacationHoursOk() (*QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours, bool)`

GetVacationHoursOk returns a tuple with the VacationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) SetVacationHours(v QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours)`

SetVacationHours sets VacationHours field to given value.

### HasVacationHours

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayroll) HasVacationHours() bool`

HasVacationHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


