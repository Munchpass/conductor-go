# QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing invoice line group you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new invoice line groups you wish to add. | 
**ItemGroupId** | Pointer to **string** | The invoice line group&#39;s item group, representing a predefined set of items bundled because they are commonly purchased together or grouped for faster entry. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item group associated with this invoice line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this invoice line group. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**OverrideUnitOfMeasureSetId** | Pointer to **string** | Specifies an alternative unit-of-measure set when updating this invoice line group&#39;s &#x60;unitOfMeasure&#x60; field (e.g., \&quot;pound\&quot; or \&quot;kilogram\&quot;). This allows you to select units from a different set than the item&#39;s default unit-of-measure set, which remains unchanged on the item itself. The override applies only to this specific line. For example, you can sell an item typically measured in volume units using weight units in a specific transaction by specifying a different unit-of-measure set with this field. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInvoicesIdPostRequestLinesInner**](QuickbooksDesktopInvoicesIdPostRequestLinesInner.md) | The invoice line group&#39;s line items, each representing a single product or service sold.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the invoice line group with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInner

`func NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInner(id string, ) *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner`

NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInner instantiates a new QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInnerWithDefaults

`func NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInnerWithDefaults() *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner`

NewQuickbooksDesktopInvoicesIdPostRequestLineGroupsInnerWithDefaults instantiates a new QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemGroupId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.

### HasItemGroupId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) HasItemGroupId() bool`

HasItemGroupId returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetOverrideUnitOfMeasureSetId() string`

GetOverrideUnitOfMeasureSetId returns the OverrideUnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetOverrideUnitOfMeasureSetIdOk() (*string, bool)`

GetOverrideUnitOfMeasureSetIdOk returns a tuple with the OverrideUnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetOverrideUnitOfMeasureSetId(v string)`

SetOverrideUnitOfMeasureSetId sets OverrideUnitOfMeasureSetId field to given value.

### HasOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) HasOverrideUnitOfMeasureSetId() bool`

HasOverrideUnitOfMeasureSetId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetLines() []QuickbooksDesktopInvoicesIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) GetLinesOk() (*[]QuickbooksDesktopInvoicesIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) SetLines(v []QuickbooksDesktopInvoicesIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInvoicesIdPostRequestLineGroupsInner) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


