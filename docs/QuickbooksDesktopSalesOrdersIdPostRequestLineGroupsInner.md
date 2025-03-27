# QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing sales order line group you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new sales order line groups you wish to add. | 
**ItemGroupId** | Pointer to **string** | The sales order line group&#39;s item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item group associated with this sales order line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales order line group. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**OverrideUnitOfMeasureSetId** | Pointer to **string** | Specifies an alternative unit-of-measure set when updating this sales order line group&#39;s &#x60;unitOfMeasure&#x60; field (e.g., \&quot;pound\&quot; or \&quot;kilogram\&quot;). This allows you to select units from a different set than the item&#39;s default unit-of-measure set, which remains unchanged on the item itself. The override applies only to this specific line. For example, you can sell an item typically measured in volume units using weight units in a specific transaction by specifying a different unit-of-measure set with this field. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopSalesOrdersIdPostRequestLinesInner**](QuickbooksDesktopSalesOrdersIdPostRequestLinesInner.md) | The sales order line group&#39;s line items, each representing a single product or service ordered.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the sales order line group with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner

`func NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner(id string, ) *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner`

NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner instantiates a new QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInnerWithDefaults

`func NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner`

NewQuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemGroupId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.

### HasItemGroupId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) HasItemGroupId() bool`

HasItemGroupId returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetOverrideUnitOfMeasureSetId() string`

GetOverrideUnitOfMeasureSetId returns the OverrideUnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetOverrideUnitOfMeasureSetIdOk() (*string, bool)`

GetOverrideUnitOfMeasureSetIdOk returns a tuple with the OverrideUnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetOverrideUnitOfMeasureSetId(v string)`

SetOverrideUnitOfMeasureSetId sets OverrideUnitOfMeasureSetId field to given value.

### HasOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) HasOverrideUnitOfMeasureSetId() bool`

HasOverrideUnitOfMeasureSetId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetLines() []QuickbooksDesktopSalesOrdersIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) GetLinesOk() (*[]QuickbooksDesktopSalesOrdersIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) SetLines(v []QuickbooksDesktopSalesOrdersIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopSalesOrdersIdPostRequestLineGroupsInner) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


