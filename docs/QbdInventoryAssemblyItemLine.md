# QbdInventoryAssemblyItemLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryItem** | [**QbdInventoryAssemblyItemLineInventoryItem**](QbdInventoryAssemblyItemLineInventoryItem.md) |  | 
**Quantity** | **float32** | The quantity of the item associated with this inventory assembly item line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 

## Methods

### NewQbdInventoryAssemblyItemLine

`func NewQbdInventoryAssemblyItemLine(inventoryItem QbdInventoryAssemblyItemLineInventoryItem, quantity float32, ) *QbdInventoryAssemblyItemLine`

NewQbdInventoryAssemblyItemLine instantiates a new QbdInventoryAssemblyItemLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventoryAssemblyItemLineWithDefaults

`func NewQbdInventoryAssemblyItemLineWithDefaults() *QbdInventoryAssemblyItemLine`

NewQbdInventoryAssemblyItemLineWithDefaults instantiates a new QbdInventoryAssemblyItemLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryItem

`func (o *QbdInventoryAssemblyItemLine) GetInventoryItem() QbdInventoryAssemblyItemLineInventoryItem`

GetInventoryItem returns the InventoryItem field if non-nil, zero value otherwise.

### GetInventoryItemOk

`func (o *QbdInventoryAssemblyItemLine) GetInventoryItemOk() (*QbdInventoryAssemblyItemLineInventoryItem, bool)`

GetInventoryItemOk returns a tuple with the InventoryItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItem

`func (o *QbdInventoryAssemblyItemLine) SetInventoryItem(v QbdInventoryAssemblyItemLineInventoryItem)`

SetInventoryItem sets InventoryItem field to given value.


### GetQuantity

`func (o *QbdInventoryAssemblyItemLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdInventoryAssemblyItemLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdInventoryAssemblyItemLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


