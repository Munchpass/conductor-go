# QbdPurchaseOrderLineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this purchase order line group. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_purchase_order_line_group\&quot;&#x60;. | 
**ItemGroup** | [**QbdPurchaseOrderLineGroupItemGroup**](QbdPurchaseOrderLineGroupItemGroup.md) |  | 
**Description** | **string** | A description of this purchase order line group. | 
**Quantity** | **float32** | The quantity of the item group associated with this purchase order line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this purchase order line group. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet**](QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet.md) |  | 
**ShouldPrintItemsInGroup** | **bool** | Indicates whether the individual items in this purchase order line group and their separate amounts appear on printed forms. | 
**TotalAmount** | **string** | The total monetary amount of this purchase order line group, equivalent to the sum of the amounts in &#x60;lines&#x60;, represented as a decimal string. | 
**Lines** | [**[]QbdPurchaseOrderLine**](QbdPurchaseOrderLine.md) | The purchase order line group&#39;s line items, each representing a single product or service ordered. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the purchase order line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdPurchaseOrderLineGroup

`func NewQbdPurchaseOrderLineGroup(id string, objectType string, itemGroup QbdPurchaseOrderLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet, shouldPrintItemsInGroup bool, totalAmount string, lines []QbdPurchaseOrderLine, customFields []QbdCustomField, ) *QbdPurchaseOrderLineGroup`

NewQbdPurchaseOrderLineGroup instantiates a new QbdPurchaseOrderLineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPurchaseOrderLineGroupWithDefaults

`func NewQbdPurchaseOrderLineGroupWithDefaults() *QbdPurchaseOrderLineGroup`

NewQbdPurchaseOrderLineGroupWithDefaults instantiates a new QbdPurchaseOrderLineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdPurchaseOrderLineGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdPurchaseOrderLineGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdPurchaseOrderLineGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdPurchaseOrderLineGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdPurchaseOrderLineGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdPurchaseOrderLineGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemGroup

`func (o *QbdPurchaseOrderLineGroup) GetItemGroup() QbdPurchaseOrderLineGroupItemGroup`

GetItemGroup returns the ItemGroup field if non-nil, zero value otherwise.

### GetItemGroupOk

`func (o *QbdPurchaseOrderLineGroup) GetItemGroupOk() (*QbdPurchaseOrderLineGroupItemGroup, bool)`

GetItemGroupOk returns a tuple with the ItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroup

`func (o *QbdPurchaseOrderLineGroup) SetItemGroup(v QbdPurchaseOrderLineGroupItemGroup)`

SetItemGroup sets ItemGroup field to given value.


### GetDescription

`func (o *QbdPurchaseOrderLineGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdPurchaseOrderLineGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdPurchaseOrderLineGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdPurchaseOrderLineGroup) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdPurchaseOrderLineGroup) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdPurchaseOrderLineGroup) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdPurchaseOrderLineGroup) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdPurchaseOrderLineGroup) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdPurchaseOrderLineGroup) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdPurchaseOrderLineGroup) GetOverrideUnitOfMeasureSet() QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdPurchaseOrderLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdPurchaseOrderLineGroup) SetOverrideUnitOfMeasureSet(v QbdPurchaseOrderLineGroupOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetShouldPrintItemsInGroup

`func (o *QbdPurchaseOrderLineGroup) GetShouldPrintItemsInGroup() bool`

GetShouldPrintItemsInGroup returns the ShouldPrintItemsInGroup field if non-nil, zero value otherwise.

### GetShouldPrintItemsInGroupOk

`func (o *QbdPurchaseOrderLineGroup) GetShouldPrintItemsInGroupOk() (*bool, bool)`

GetShouldPrintItemsInGroupOk returns a tuple with the ShouldPrintItemsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrintItemsInGroup

`func (o *QbdPurchaseOrderLineGroup) SetShouldPrintItemsInGroup(v bool)`

SetShouldPrintItemsInGroup sets ShouldPrintItemsInGroup field to given value.


### GetTotalAmount

`func (o *QbdPurchaseOrderLineGroup) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdPurchaseOrderLineGroup) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdPurchaseOrderLineGroup) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetLines

`func (o *QbdPurchaseOrderLineGroup) GetLines() []QbdPurchaseOrderLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdPurchaseOrderLineGroup) GetLinesOk() (*[]QbdPurchaseOrderLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdPurchaseOrderLineGroup) SetLines(v []QbdPurchaseOrderLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdPurchaseOrderLineGroup) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdPurchaseOrderLineGroup) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdPurchaseOrderLineGroup) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


