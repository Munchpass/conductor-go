# QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewQuantity** | Pointer to **float32** | The new quantity for the inventory item associated with this inventory adjustment line. | [optional] 
**QuantityDifference** | Pointer to **float32** | Either a positive or negative number that shows the change in quantity for the inventory item associated with this inventory adjustment line. A positive number increases the quantity, while a negative number decreases it. | [optional] 
**NewValue** | Pointer to **string** | The new total value of the entire stock of the inventory item associated with this inventory adjustment line.  **NOTE**: The new value does _not_ have to equal &#x60;quantityOnHand&#x60; times &#x60;purchaseCost&#x60;. | [optional] 
**ValueDifference** | Pointer to **float32** | Either a positive or negative number that shows the change in the total value of the entire stock of the inventory item associated with this inventory adjustment line. A positive number increases the value, while a negative number decreases it. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValueWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValueWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValueWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetNewQuantity() float32`

GetNewQuantity returns the NewQuantity field if non-nil, zero value otherwise.

### GetNewQuantityOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetNewQuantityOk() (*float32, bool)`

GetNewQuantityOk returns a tuple with the NewQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) SetNewQuantity(v float32)`

SetNewQuantity sets NewQuantity field to given value.

### HasNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) HasNewQuantity() bool`

HasNewQuantity returns a boolean if a field has been set.

### GetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetQuantityDifference() float32`

GetQuantityDifference returns the QuantityDifference field if non-nil, zero value otherwise.

### GetQuantityDifferenceOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetQuantityDifferenceOk() (*float32, bool)`

GetQuantityDifferenceOk returns a tuple with the QuantityDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) SetQuantityDifference(v float32)`

SetQuantityDifference sets QuantityDifference field to given value.

### HasQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) HasQuantityDifference() bool`

HasQuantityDifference returns a boolean if a field has been set.

### GetNewValue

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetValueDifference() float32`

GetValueDifference returns the ValueDifference field if non-nil, zero value otherwise.

### GetValueDifferenceOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) GetValueDifferenceOk() (*float32, bool)`

GetValueDifferenceOk returns a tuple with the ValueDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) SetValueDifference(v float32)`

SetValueDifference sets ValueDifference field to given value.

### HasValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustValue) HasValueDifference() bool`

HasValueDifference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


