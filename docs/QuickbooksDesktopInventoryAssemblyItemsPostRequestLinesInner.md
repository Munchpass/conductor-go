# QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryItemId** | Pointer to **string** | The inventory item associated with this inventory assembly item line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this inventory assembly item line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner

`func NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner() *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner`

NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner instantiates a new QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInnerWithDefaults() *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner`

NewQuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryItemId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


