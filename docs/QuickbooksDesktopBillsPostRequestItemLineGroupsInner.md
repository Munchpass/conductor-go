# QuickbooksDesktopBillsPostRequestItemLineGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemGroupId** | **string** | The item line group&#39;s item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry. | 
**Quantity** | Pointer to **float32** | The quantity of the item group associated with this item line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this item line group. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item group associated with this item line group is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item group associated with this item line group is stored. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the item line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopBillsPostRequestItemLineGroupsInner

`func NewQuickbooksDesktopBillsPostRequestItemLineGroupsInner(itemGroupId string, ) *QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

NewQuickbooksDesktopBillsPostRequestItemLineGroupsInner instantiates a new QuickbooksDesktopBillsPostRequestItemLineGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsPostRequestItemLineGroupsInnerWithDefaults

`func NewQuickbooksDesktopBillsPostRequestItemLineGroupsInnerWithDefaults() *QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

NewQuickbooksDesktopBillsPostRequestItemLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopBillsPostRequestItemLineGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemGroupId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.


### GetQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLineGroupsInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


