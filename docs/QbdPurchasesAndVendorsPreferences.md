# QbdPurchasesAndVendorsPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsingInventory** | **bool** | Indicates whether this company file has inventory-related features enabled. | 
**DaysBillsAreDue** | **float32** | The default number of days after receipt when bills are due for this company file. | 
**IsAutomaticallyUsingDiscounts** | **bool** | Indicates whether this company file is configured to automatically apply available vendor discounts or credits when paying bills. | 
**DefaultDiscountAccount** | [**QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount**](QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount.md) |  | 

## Methods

### NewQbdPurchasesAndVendorsPreferences

`func NewQbdPurchasesAndVendorsPreferences(isUsingInventory bool, daysBillsAreDue float32, isAutomaticallyUsingDiscounts bool, defaultDiscountAccount QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount, ) *QbdPurchasesAndVendorsPreferences`

NewQbdPurchasesAndVendorsPreferences instantiates a new QbdPurchasesAndVendorsPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPurchasesAndVendorsPreferencesWithDefaults

`func NewQbdPurchasesAndVendorsPreferencesWithDefaults() *QbdPurchasesAndVendorsPreferences`

NewQbdPurchasesAndVendorsPreferencesWithDefaults instantiates a new QbdPurchasesAndVendorsPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsingInventory

`func (o *QbdPurchasesAndVendorsPreferences) GetIsUsingInventory() bool`

GetIsUsingInventory returns the IsUsingInventory field if non-nil, zero value otherwise.

### GetIsUsingInventoryOk

`func (o *QbdPurchasesAndVendorsPreferences) GetIsUsingInventoryOk() (*bool, bool)`

GetIsUsingInventoryOk returns a tuple with the IsUsingInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingInventory

`func (o *QbdPurchasesAndVendorsPreferences) SetIsUsingInventory(v bool)`

SetIsUsingInventory sets IsUsingInventory field to given value.


### GetDaysBillsAreDue

`func (o *QbdPurchasesAndVendorsPreferences) GetDaysBillsAreDue() float32`

GetDaysBillsAreDue returns the DaysBillsAreDue field if non-nil, zero value otherwise.

### GetDaysBillsAreDueOk

`func (o *QbdPurchasesAndVendorsPreferences) GetDaysBillsAreDueOk() (*float32, bool)`

GetDaysBillsAreDueOk returns a tuple with the DaysBillsAreDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBillsAreDue

`func (o *QbdPurchasesAndVendorsPreferences) SetDaysBillsAreDue(v float32)`

SetDaysBillsAreDue sets DaysBillsAreDue field to given value.


### GetIsAutomaticallyUsingDiscounts

`func (o *QbdPurchasesAndVendorsPreferences) GetIsAutomaticallyUsingDiscounts() bool`

GetIsAutomaticallyUsingDiscounts returns the IsAutomaticallyUsingDiscounts field if non-nil, zero value otherwise.

### GetIsAutomaticallyUsingDiscountsOk

`func (o *QbdPurchasesAndVendorsPreferences) GetIsAutomaticallyUsingDiscountsOk() (*bool, bool)`

GetIsAutomaticallyUsingDiscountsOk returns a tuple with the IsAutomaticallyUsingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticallyUsingDiscounts

`func (o *QbdPurchasesAndVendorsPreferences) SetIsAutomaticallyUsingDiscounts(v bool)`

SetIsAutomaticallyUsingDiscounts sets IsAutomaticallyUsingDiscounts field to given value.


### GetDefaultDiscountAccount

`func (o *QbdPurchasesAndVendorsPreferences) GetDefaultDiscountAccount() QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount`

GetDefaultDiscountAccount returns the DefaultDiscountAccount field if non-nil, zero value otherwise.

### GetDefaultDiscountAccountOk

`func (o *QbdPurchasesAndVendorsPreferences) GetDefaultDiscountAccountOk() (*QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount, bool)`

GetDefaultDiscountAccountOk returns a tuple with the DefaultDiscountAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAccount

`func (o *QbdPurchasesAndVendorsPreferences) SetDefaultDiscountAccount(v QbdPurchasesAndVendorsPreferencesDefaultDiscountAccount)`

SetDefaultDiscountAccount sets DefaultDiscountAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


