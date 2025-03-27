# QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemGroupId** | **string** | The sales order line group&#39;s item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry. | 
**Quantity** | Pointer to **float32** | The quantity of the item group associated with this sales order line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales order line group. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item group associated with this sales order line group is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item group associated with this sales order line group is stored. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the sales order line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner

`func NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner(itemGroupId string, ) *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner`

NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInner instantiates a new QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInnerWithDefaults

`func NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner`

NewQuickbooksDesktopSalesOrdersPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemGroupId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.


### GetQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


