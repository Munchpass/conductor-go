# QbdEmployeeEmployeePayroll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPeriod** | **string** | How frequently this employee is paid (e.g., weekly, biweekly, monthly). This determines the schedule for generating paychecks. | 
**Class** | [**QbdEmployeePayrollInfoClass**](QbdEmployeePayrollInfoClass.md) |  | 
**Earnings** | [**[]QbdEarnings**](QbdEarnings.md) | The employee&#39;s earnings. | 
**UseTimeDataToCreatePaychecks** | **string** | Indicates whether this employee is using time-tracking data to create paychecks. | 
**SickHours** | [**QbdEmployeePayrollInfoSickHours**](QbdEmployeePayrollInfoSickHours.md) |  | 
**VacationHours** | [**QbdEmployeePayrollInfoVacationHours**](QbdEmployeePayrollInfoVacationHours.md) |  | 

## Methods

### NewQbdEmployeeEmployeePayroll

`func NewQbdEmployeeEmployeePayroll(payPeriod string, class QbdEmployeePayrollInfoClass, earnings []QbdEarnings, useTimeDataToCreatePaychecks string, sickHours QbdEmployeePayrollInfoSickHours, vacationHours QbdEmployeePayrollInfoVacationHours, ) *QbdEmployeeEmployeePayroll`

NewQbdEmployeeEmployeePayroll instantiates a new QbdEmployeeEmployeePayroll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeeEmployeePayrollWithDefaults

`func NewQbdEmployeeEmployeePayrollWithDefaults() *QbdEmployeeEmployeePayroll`

NewQbdEmployeeEmployeePayrollWithDefaults instantiates a new QbdEmployeeEmployeePayroll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPeriod

`func (o *QbdEmployeeEmployeePayroll) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *QbdEmployeeEmployeePayroll) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *QbdEmployeeEmployeePayroll) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.


### GetClass

`func (o *QbdEmployeeEmployeePayroll) GetClass() QbdEmployeePayrollInfoClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdEmployeeEmployeePayroll) GetClassOk() (*QbdEmployeePayrollInfoClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdEmployeeEmployeePayroll) SetClass(v QbdEmployeePayrollInfoClass)`

SetClass sets Class field to given value.


### GetEarnings

`func (o *QbdEmployeeEmployeePayroll) GetEarnings() []QbdEarnings`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *QbdEmployeeEmployeePayroll) GetEarningsOk() (*[]QbdEarnings, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *QbdEmployeeEmployeePayroll) SetEarnings(v []QbdEarnings)`

SetEarnings sets Earnings field to given value.


### GetUseTimeDataToCreatePaychecks

`func (o *QbdEmployeeEmployeePayroll) GetUseTimeDataToCreatePaychecks() string`

GetUseTimeDataToCreatePaychecks returns the UseTimeDataToCreatePaychecks field if non-nil, zero value otherwise.

### GetUseTimeDataToCreatePaychecksOk

`func (o *QbdEmployeeEmployeePayroll) GetUseTimeDataToCreatePaychecksOk() (*string, bool)`

GetUseTimeDataToCreatePaychecksOk returns a tuple with the UseTimeDataToCreatePaychecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeDataToCreatePaychecks

`func (o *QbdEmployeeEmployeePayroll) SetUseTimeDataToCreatePaychecks(v string)`

SetUseTimeDataToCreatePaychecks sets UseTimeDataToCreatePaychecks field to given value.


### GetSickHours

`func (o *QbdEmployeeEmployeePayroll) GetSickHours() QbdEmployeePayrollInfoSickHours`

GetSickHours returns the SickHours field if non-nil, zero value otherwise.

### GetSickHoursOk

`func (o *QbdEmployeeEmployeePayroll) GetSickHoursOk() (*QbdEmployeePayrollInfoSickHours, bool)`

GetSickHoursOk returns a tuple with the SickHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickHours

`func (o *QbdEmployeeEmployeePayroll) SetSickHours(v QbdEmployeePayrollInfoSickHours)`

SetSickHours sets SickHours field to given value.


### GetVacationHours

`func (o *QbdEmployeeEmployeePayroll) GetVacationHours() QbdEmployeePayrollInfoVacationHours`

GetVacationHours returns the VacationHours field if non-nil, zero value otherwise.

### GetVacationHoursOk

`func (o *QbdEmployeeEmployeePayroll) GetVacationHoursOk() (*QbdEmployeePayrollInfoVacationHours, bool)`

GetVacationHoursOk returns a tuple with the VacationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationHours

`func (o *QbdEmployeeEmployeePayroll) SetVacationHours(v QbdEmployeePayrollInfoVacationHours)`

SetVacationHours sets VacationHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


