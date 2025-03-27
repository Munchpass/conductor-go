# QbdPreferencesMultiCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMultiCurrencyEnabled** | **bool** | Indicates whether the multicurrency feature is enabled for this company file. Once multicurrency is enabled for a company file, it cannot be disabled. | 
**HomeCurrency** | [**QbdMultiCurrencyPreferencesHomeCurrency**](QbdMultiCurrencyPreferencesHomeCurrency.md) |  | 

## Methods

### NewQbdPreferencesMultiCurrency

`func NewQbdPreferencesMultiCurrency(isMultiCurrencyEnabled bool, homeCurrency QbdMultiCurrencyPreferencesHomeCurrency, ) *QbdPreferencesMultiCurrency`

NewQbdPreferencesMultiCurrency instantiates a new QbdPreferencesMultiCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPreferencesMultiCurrencyWithDefaults

`func NewQbdPreferencesMultiCurrencyWithDefaults() *QbdPreferencesMultiCurrency`

NewQbdPreferencesMultiCurrencyWithDefaults instantiates a new QbdPreferencesMultiCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMultiCurrencyEnabled

`func (o *QbdPreferencesMultiCurrency) GetIsMultiCurrencyEnabled() bool`

GetIsMultiCurrencyEnabled returns the IsMultiCurrencyEnabled field if non-nil, zero value otherwise.

### GetIsMultiCurrencyEnabledOk

`func (o *QbdPreferencesMultiCurrency) GetIsMultiCurrencyEnabledOk() (*bool, bool)`

GetIsMultiCurrencyEnabledOk returns a tuple with the IsMultiCurrencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiCurrencyEnabled

`func (o *QbdPreferencesMultiCurrency) SetIsMultiCurrencyEnabled(v bool)`

SetIsMultiCurrencyEnabled sets IsMultiCurrencyEnabled field to given value.


### GetHomeCurrency

`func (o *QbdPreferencesMultiCurrency) GetHomeCurrency() QbdMultiCurrencyPreferencesHomeCurrency`

GetHomeCurrency returns the HomeCurrency field if non-nil, zero value otherwise.

### GetHomeCurrencyOk

`func (o *QbdPreferencesMultiCurrency) GetHomeCurrencyOk() (*QbdMultiCurrencyPreferencesHomeCurrency, bool)`

GetHomeCurrencyOk returns a tuple with the HomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeCurrency

`func (o *QbdPreferencesMultiCurrency) SetHomeCurrency(v QbdMultiCurrencyPreferencesHomeCurrency)`

SetHomeCurrency sets HomeCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


