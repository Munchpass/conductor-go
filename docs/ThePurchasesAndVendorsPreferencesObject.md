# ThePurchasesAndVendorsPreferencesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsingInventory** | **bool** | Indicates whether this company file has inventory-related features enabled. | 
**DaysBillsAreDue** | **float32** | The default number of days after receipt when bills are due for this company file. | 
**IsAutomaticallyUsingDiscounts** | **bool** | Indicates whether this company file is configured to automatically apply available vendor discounts or credits when paying bills. | 
**DefaultDiscountAccount** | [**ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount**](ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount.md) |  | 

## Methods

### NewThePurchasesAndVendorsPreferencesObject

`func NewThePurchasesAndVendorsPreferencesObject(isUsingInventory bool, daysBillsAreDue float32, isAutomaticallyUsingDiscounts bool, defaultDiscountAccount ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount, ) *ThePurchasesAndVendorsPreferencesObject`

NewThePurchasesAndVendorsPreferencesObject instantiates a new ThePurchasesAndVendorsPreferencesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThePurchasesAndVendorsPreferencesObjectWithDefaults

`func NewThePurchasesAndVendorsPreferencesObjectWithDefaults() *ThePurchasesAndVendorsPreferencesObject`

NewThePurchasesAndVendorsPreferencesObjectWithDefaults instantiates a new ThePurchasesAndVendorsPreferencesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsingInventory

`func (o *ThePurchasesAndVendorsPreferencesObject) GetIsUsingInventory() bool`

GetIsUsingInventory returns the IsUsingInventory field if non-nil, zero value otherwise.

### GetIsUsingInventoryOk

`func (o *ThePurchasesAndVendorsPreferencesObject) GetIsUsingInventoryOk() (*bool, bool)`

GetIsUsingInventoryOk returns a tuple with the IsUsingInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingInventory

`func (o *ThePurchasesAndVendorsPreferencesObject) SetIsUsingInventory(v bool)`

SetIsUsingInventory sets IsUsingInventory field to given value.


### GetDaysBillsAreDue

`func (o *ThePurchasesAndVendorsPreferencesObject) GetDaysBillsAreDue() float32`

GetDaysBillsAreDue returns the DaysBillsAreDue field if non-nil, zero value otherwise.

### GetDaysBillsAreDueOk

`func (o *ThePurchasesAndVendorsPreferencesObject) GetDaysBillsAreDueOk() (*float32, bool)`

GetDaysBillsAreDueOk returns a tuple with the DaysBillsAreDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBillsAreDue

`func (o *ThePurchasesAndVendorsPreferencesObject) SetDaysBillsAreDue(v float32)`

SetDaysBillsAreDue sets DaysBillsAreDue field to given value.


### GetIsAutomaticallyUsingDiscounts

`func (o *ThePurchasesAndVendorsPreferencesObject) GetIsAutomaticallyUsingDiscounts() bool`

GetIsAutomaticallyUsingDiscounts returns the IsAutomaticallyUsingDiscounts field if non-nil, zero value otherwise.

### GetIsAutomaticallyUsingDiscountsOk

`func (o *ThePurchasesAndVendorsPreferencesObject) GetIsAutomaticallyUsingDiscountsOk() (*bool, bool)`

GetIsAutomaticallyUsingDiscountsOk returns a tuple with the IsAutomaticallyUsingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticallyUsingDiscounts

`func (o *ThePurchasesAndVendorsPreferencesObject) SetIsAutomaticallyUsingDiscounts(v bool)`

SetIsAutomaticallyUsingDiscounts sets IsAutomaticallyUsingDiscounts field to given value.


### GetDefaultDiscountAccount

`func (o *ThePurchasesAndVendorsPreferencesObject) GetDefaultDiscountAccount() ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount`

GetDefaultDiscountAccount returns the DefaultDiscountAccount field if non-nil, zero value otherwise.

### GetDefaultDiscountAccountOk

`func (o *ThePurchasesAndVendorsPreferencesObject) GetDefaultDiscountAccountOk() (*ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount, bool)`

GetDefaultDiscountAccountOk returns a tuple with the DefaultDiscountAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAccount

`func (o *ThePurchasesAndVendorsPreferencesObject) SetDefaultDiscountAccount(v ThePurchasesAndVendorsPreferencesObjectDefaultDiscountAccount)`

SetDefaultDiscountAccount sets DefaultDiscountAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


