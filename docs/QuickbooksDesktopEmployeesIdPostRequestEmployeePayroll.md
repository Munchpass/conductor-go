# QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPeriod** | Pointer to **string** | How frequently this employee is paid (e.g., weekly, biweekly, monthly). This determines the schedule for generating paychecks. | [optional] 
**ClassId** | Pointer to **string** | The employee&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**DeleteAllEarnings** | Pointer to **bool** | When &#x60;true&#x60;, deletes all earnings records for this employee. | [optional] 
**Earnings** | Pointer to [**[]QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner.md) | The employee&#39;s earnings.  **IMPORTANT**: When updating employees, if you include any earnings records in your update request, QuickBooks will delete all existing earnings records for this employee and replace them with the new records you provide. If you do not include any earnings records, the existing earnings records will remain unchanged. To delete all earnings records without adding new ones, set the &#x60;deleteAllEarnings&#x60; field to &#x60;true&#x60;. | [optional] 
**UseTimeDataToCreatePaychecks** | Pointer to **string** | Indicates whether this employee is using time-tracking data to create paychecks. | [optional] 
**SickHours** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours.md) |  | [optional] 
**VacationHours** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours**](QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayroll

`func NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayroll() *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll`

NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayroll instantiates a new QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayrollWithDefaults

`func NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayrollWithDefaults() *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll`

NewQuickbooksDesktopEmployeesIdPostRequestEmployeePayrollWithDefaults instantiates a new QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPeriod

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDeleteAllEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetDeleteAllEarnings() bool`

GetDeleteAllEarnings returns the DeleteAllEarnings field if non-nil, zero value otherwise.

### GetDeleteAllEarningsOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetDeleteAllEarningsOk() (*bool, bool)`

GetDeleteAllEarningsOk returns a tuple with the DeleteAllEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetDeleteAllEarnings(v bool)`

SetDeleteAllEarnings sets DeleteAllEarnings field to given value.

### HasDeleteAllEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasDeleteAllEarnings() bool`

HasDeleteAllEarnings returns a boolean if a field has been set.

### GetEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetEarnings() []QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetEarningsOk() (*[]QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetEarnings(v []QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetUseTimeDataToCreatePaychecks() string`

GetUseTimeDataToCreatePaychecks returns the UseTimeDataToCreatePaychecks field if non-nil, zero value otherwise.

### GetUseTimeDataToCreatePaychecksOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetUseTimeDataToCreatePaychecksOk() (*string, bool)`

GetUseTimeDataToCreatePaychecksOk returns a tuple with the UseTimeDataToCreatePaychecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetUseTimeDataToCreatePaychecks(v string)`

SetUseTimeDataToCreatePaychecks sets UseTimeDataToCreatePaychecks field to given value.

### HasUseTimeDataToCreatePaychecks

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasUseTimeDataToCreatePaychecks() bool`

HasUseTimeDataToCreatePaychecks returns a boolean if a field has been set.

### GetSickHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetSickHours() QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours`

GetSickHours returns the SickHours field if non-nil, zero value otherwise.

### GetSickHoursOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetSickHoursOk() (*QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours, bool)`

GetSickHoursOk returns a tuple with the SickHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetSickHours(v QuickbooksDesktopEmployeesPostRequestEmployeePayrollSickHours)`

SetSickHours sets SickHours field to given value.

### HasSickHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasSickHours() bool`

HasSickHours returns a boolean if a field has been set.

### GetVacationHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetVacationHours() QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours`

GetVacationHours returns the VacationHours field if non-nil, zero value otherwise.

### GetVacationHoursOk

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) GetVacationHoursOk() (*QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours, bool)`

GetVacationHoursOk returns a tuple with the VacationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) SetVacationHours(v QuickbooksDesktopEmployeesPostRequestEmployeePayrollVacationHours)`

SetVacationHours sets VacationHours field to given value.

### HasVacationHours

`func (o *QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll) HasVacationHours() bool`

HasVacationHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


