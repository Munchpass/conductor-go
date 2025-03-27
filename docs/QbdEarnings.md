# QbdEarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayrollWageItem** | [**QbdEarningsPayrollWageItem**](QbdEarningsPayrollWageItem.md) |  | 
**Rate** | **string** | The hourly rate for this employee, represented as a decimal string. | 
**RatePercent** | **string** | The hourly rate for this employee expressed as a percentage. | 

## Methods

### NewQbdEarnings

`func NewQbdEarnings(payrollWageItem QbdEarningsPayrollWageItem, rate string, ratePercent string, ) *QbdEarnings`

NewQbdEarnings instantiates a new QbdEarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEarningsWithDefaults

`func NewQbdEarningsWithDefaults() *QbdEarnings`

NewQbdEarningsWithDefaults instantiates a new QbdEarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayrollWageItem

`func (o *QbdEarnings) GetPayrollWageItem() QbdEarningsPayrollWageItem`

GetPayrollWageItem returns the PayrollWageItem field if non-nil, zero value otherwise.

### GetPayrollWageItemOk

`func (o *QbdEarnings) GetPayrollWageItemOk() (*QbdEarningsPayrollWageItem, bool)`

GetPayrollWageItemOk returns a tuple with the PayrollWageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollWageItem

`func (o *QbdEarnings) SetPayrollWageItem(v QbdEarningsPayrollWageItem)`

SetPayrollWageItem sets PayrollWageItem field to given value.


### GetRate

`func (o *QbdEarnings) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QbdEarnings) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QbdEarnings) SetRate(v string)`

SetRate sets Rate field to given value.


### GetRatePercent

`func (o *QbdEarnings) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QbdEarnings) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QbdEarnings) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


