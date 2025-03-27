# TheFinanceChargePreferencesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualInterestRate** | **float32** | The interest rate that QuickBooks will use to calculate finance charges for this company file. Default is &#x60;0&#x60;. | [default to 0]
**MinimumFinanceCharge** | **float32** | The minimum finance charge that will be applied regardless of the amount overdue for this company file. Default is &#x60;0&#x60;. | [default to 0]
**GracePeriod** | **float32** | The number of days before finance charges apply to customers&#39; overdue invoices for this company file. Default is &#x60;0&#x60;. | [default to 0]
**FinanceChargeAccount** | [**TheFinanceChargePreferencesObjectFinanceChargeAccount**](TheFinanceChargePreferencesObjectFinanceChargeAccount.md) |  | 
**IsAssessingForOverdueCharges** | **bool** | Indicates whether this company file is configured to assess finance charges for overdue invoices. Default is &#x60;false&#x60;. (Note that laws vary about whether a company can charge interest on overdue interest payments.) | 
**CalculateChargesFrom** | **string** | The date from which finance charges are calculated for this company file. Default is &#x60;due_date&#x60;. | 
**IsMarkedToBePrinted** | **bool** | Indicates whether this company file is configured to mark all newly created finance-charge invoices as \&quot;to be printed\&quot;. Default is &#x60;false&#x60;. The user can still change this preference for each individual invoice. | 

## Methods

### NewTheFinanceChargePreferencesObject

`func NewTheFinanceChargePreferencesObject(annualInterestRate float32, minimumFinanceCharge float32, gracePeriod float32, financeChargeAccount TheFinanceChargePreferencesObjectFinanceChargeAccount, isAssessingForOverdueCharges bool, calculateChargesFrom string, isMarkedToBePrinted bool, ) *TheFinanceChargePreferencesObject`

NewTheFinanceChargePreferencesObject instantiates a new TheFinanceChargePreferencesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTheFinanceChargePreferencesObjectWithDefaults

`func NewTheFinanceChargePreferencesObjectWithDefaults() *TheFinanceChargePreferencesObject`

NewTheFinanceChargePreferencesObjectWithDefaults instantiates a new TheFinanceChargePreferencesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualInterestRate

`func (o *TheFinanceChargePreferencesObject) GetAnnualInterestRate() float32`

GetAnnualInterestRate returns the AnnualInterestRate field if non-nil, zero value otherwise.

### GetAnnualInterestRateOk

`func (o *TheFinanceChargePreferencesObject) GetAnnualInterestRateOk() (*float32, bool)`

GetAnnualInterestRateOk returns a tuple with the AnnualInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualInterestRate

`func (o *TheFinanceChargePreferencesObject) SetAnnualInterestRate(v float32)`

SetAnnualInterestRate sets AnnualInterestRate field to given value.


### GetMinimumFinanceCharge

`func (o *TheFinanceChargePreferencesObject) GetMinimumFinanceCharge() float32`

GetMinimumFinanceCharge returns the MinimumFinanceCharge field if non-nil, zero value otherwise.

### GetMinimumFinanceChargeOk

`func (o *TheFinanceChargePreferencesObject) GetMinimumFinanceChargeOk() (*float32, bool)`

GetMinimumFinanceChargeOk returns a tuple with the MinimumFinanceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFinanceCharge

`func (o *TheFinanceChargePreferencesObject) SetMinimumFinanceCharge(v float32)`

SetMinimumFinanceCharge sets MinimumFinanceCharge field to given value.


### GetGracePeriod

`func (o *TheFinanceChargePreferencesObject) GetGracePeriod() float32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TheFinanceChargePreferencesObject) GetGracePeriodOk() (*float32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TheFinanceChargePreferencesObject) SetGracePeriod(v float32)`

SetGracePeriod sets GracePeriod field to given value.


### GetFinanceChargeAccount

`func (o *TheFinanceChargePreferencesObject) GetFinanceChargeAccount() TheFinanceChargePreferencesObjectFinanceChargeAccount`

GetFinanceChargeAccount returns the FinanceChargeAccount field if non-nil, zero value otherwise.

### GetFinanceChargeAccountOk

`func (o *TheFinanceChargePreferencesObject) GetFinanceChargeAccountOk() (*TheFinanceChargePreferencesObjectFinanceChargeAccount, bool)`

GetFinanceChargeAccountOk returns a tuple with the FinanceChargeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceChargeAccount

`func (o *TheFinanceChargePreferencesObject) SetFinanceChargeAccount(v TheFinanceChargePreferencesObjectFinanceChargeAccount)`

SetFinanceChargeAccount sets FinanceChargeAccount field to given value.


### GetIsAssessingForOverdueCharges

`func (o *TheFinanceChargePreferencesObject) GetIsAssessingForOverdueCharges() bool`

GetIsAssessingForOverdueCharges returns the IsAssessingForOverdueCharges field if non-nil, zero value otherwise.

### GetIsAssessingForOverdueChargesOk

`func (o *TheFinanceChargePreferencesObject) GetIsAssessingForOverdueChargesOk() (*bool, bool)`

GetIsAssessingForOverdueChargesOk returns a tuple with the IsAssessingForOverdueCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssessingForOverdueCharges

`func (o *TheFinanceChargePreferencesObject) SetIsAssessingForOverdueCharges(v bool)`

SetIsAssessingForOverdueCharges sets IsAssessingForOverdueCharges field to given value.


### GetCalculateChargesFrom

`func (o *TheFinanceChargePreferencesObject) GetCalculateChargesFrom() string`

GetCalculateChargesFrom returns the CalculateChargesFrom field if non-nil, zero value otherwise.

### GetCalculateChargesFromOk

`func (o *TheFinanceChargePreferencesObject) GetCalculateChargesFromOk() (*string, bool)`

GetCalculateChargesFromOk returns a tuple with the CalculateChargesFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateChargesFrom

`func (o *TheFinanceChargePreferencesObject) SetCalculateChargesFrom(v string)`

SetCalculateChargesFrom sets CalculateChargesFrom field to given value.


### GetIsMarkedToBePrinted

`func (o *TheFinanceChargePreferencesObject) GetIsMarkedToBePrinted() bool`

GetIsMarkedToBePrinted returns the IsMarkedToBePrinted field if non-nil, zero value otherwise.

### GetIsMarkedToBePrintedOk

`func (o *TheFinanceChargePreferencesObject) GetIsMarkedToBePrintedOk() (*bool, bool)`

GetIsMarkedToBePrintedOk returns a tuple with the IsMarkedToBePrinted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedToBePrinted

`func (o *TheFinanceChargePreferencesObject) SetIsMarkedToBePrinted(v bool)`

SetIsMarkedToBePrinted sets IsMarkedToBePrinted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


