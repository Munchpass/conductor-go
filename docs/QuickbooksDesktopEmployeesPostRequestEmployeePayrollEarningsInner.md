# QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayrollWageItemId** | **string** | The payroll wage item that defines how this employee is paid (e.g., Regular Pay, Overtime Pay). This determines the payment scheme used for payroll calculations. | 
**Rate** | Pointer to **string** | The hourly rate for this employee, represented as a decimal string. | [optional] 
**RatePercent** | Pointer to **string** | The hourly rate for this employee expressed as a percentage. | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner(payrollWageItemId string, ) *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInnerWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInnerWithDefaults() *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner`

NewQuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInnerWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayrollWageItemId

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetPayrollWageItemId() string`

GetPayrollWageItemId returns the PayrollWageItemId field if non-nil, zero value otherwise.

### GetPayrollWageItemIdOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetPayrollWageItemIdOk() (*string, bool)`

GetPayrollWageItemIdOk returns a tuple with the PayrollWageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollWageItemId

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) SetPayrollWageItemId(v string)`

SetPayrollWageItemId sets PayrollWageItemId field to given value.


### GetRate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRatePercent

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.

### HasRatePercent

`func (o *QuickbooksDesktopEmployeesPostRequestEmployeePayrollEarningsInner) HasRatePercent() bool`

HasRatePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


