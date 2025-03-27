# QuickbooksDesktopInvoicesPostRequestLineGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemGroupId** | **string** | The invoice line group&#39;s item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry. | 
**Quantity** | Pointer to **float32** | The quantity of the item group associated with this invoice line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this invoice line group. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item group associated with this invoice line group is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item group associated with this invoice line group is stored. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the invoice line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopInvoicesPostRequestLineGroupsInner

`func NewQuickbooksDesktopInvoicesPostRequestLineGroupsInner(itemGroupId string, ) *QuickbooksDesktopInvoicesPostRequestLineGroupsInner`

NewQuickbooksDesktopInvoicesPostRequestLineGroupsInner instantiates a new QuickbooksDesktopInvoicesPostRequestLineGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInvoicesPostRequestLineGroupsInnerWithDefaults

`func NewQuickbooksDesktopInvoicesPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopInvoicesPostRequestLineGroupsInner`

NewQuickbooksDesktopInvoicesPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopInvoicesPostRequestLineGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemGroupId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.


### GetQuantity

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopInvoicesPostRequestLineGroupsInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


