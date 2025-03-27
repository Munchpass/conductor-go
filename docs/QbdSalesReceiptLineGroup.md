# QbdSalesReceiptLineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales receipt line group. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_receipt_line_group\&quot;&#x60;. | 
**ItemGroup** | [**QbdSalesReceiptLineGroupItemGroup**](QbdSalesReceiptLineGroupItemGroup.md) |  | 
**Description** | **string** | A description of this sales receipt line group. | 
**Quantity** | **float32** | The quantity of the item group associated with this sales receipt line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales receipt line group. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet**](QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet.md) |  | 
**ShouldPrintItemsInGroup** | **bool** | Indicates whether the individual items in this sales receipt line group and their separate amounts appear on printed forms. | 
**TotalAmount** | **string** | The total monetary amount of this sales receipt line group, equivalent to the sum of the amounts in &#x60;lines&#x60;, represented as a decimal string. | 
**Lines** | [**[]QbdSalesReceiptLine**](QbdSalesReceiptLine.md) | The sales receipt line group&#39;s line items, each representing a single product or service sold. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales receipt line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesReceiptLineGroup

`func NewQbdSalesReceiptLineGroup(id string, objectType string, itemGroup QbdSalesReceiptLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet, shouldPrintItemsInGroup bool, totalAmount string, lines []QbdSalesReceiptLine, customFields []QbdCustomField, ) *QbdSalesReceiptLineGroup`

NewQbdSalesReceiptLineGroup instantiates a new QbdSalesReceiptLineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesReceiptLineGroupWithDefaults

`func NewQbdSalesReceiptLineGroupWithDefaults() *QbdSalesReceiptLineGroup`

NewQbdSalesReceiptLineGroupWithDefaults instantiates a new QbdSalesReceiptLineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesReceiptLineGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesReceiptLineGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesReceiptLineGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesReceiptLineGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesReceiptLineGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesReceiptLineGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemGroup

`func (o *QbdSalesReceiptLineGroup) GetItemGroup() QbdSalesReceiptLineGroupItemGroup`

GetItemGroup returns the ItemGroup field if non-nil, zero value otherwise.

### GetItemGroupOk

`func (o *QbdSalesReceiptLineGroup) GetItemGroupOk() (*QbdSalesReceiptLineGroupItemGroup, bool)`

GetItemGroupOk returns a tuple with the ItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroup

`func (o *QbdSalesReceiptLineGroup) SetItemGroup(v QbdSalesReceiptLineGroupItemGroup)`

SetItemGroup sets ItemGroup field to given value.


### GetDescription

`func (o *QbdSalesReceiptLineGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesReceiptLineGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesReceiptLineGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdSalesReceiptLineGroup) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdSalesReceiptLineGroup) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdSalesReceiptLineGroup) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdSalesReceiptLineGroup) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdSalesReceiptLineGroup) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdSalesReceiptLineGroup) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdSalesReceiptLineGroup) GetOverrideUnitOfMeasureSet() QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdSalesReceiptLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdSalesReceiptLineGroup) SetOverrideUnitOfMeasureSet(v QbdSalesReceiptLineGroupOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetShouldPrintItemsInGroup

`func (o *QbdSalesReceiptLineGroup) GetShouldPrintItemsInGroup() bool`

GetShouldPrintItemsInGroup returns the ShouldPrintItemsInGroup field if non-nil, zero value otherwise.

### GetShouldPrintItemsInGroupOk

`func (o *QbdSalesReceiptLineGroup) GetShouldPrintItemsInGroupOk() (*bool, bool)`

GetShouldPrintItemsInGroupOk returns a tuple with the ShouldPrintItemsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrintItemsInGroup

`func (o *QbdSalesReceiptLineGroup) SetShouldPrintItemsInGroup(v bool)`

SetShouldPrintItemsInGroup sets ShouldPrintItemsInGroup field to given value.


### GetTotalAmount

`func (o *QbdSalesReceiptLineGroup) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdSalesReceiptLineGroup) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdSalesReceiptLineGroup) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetLines

`func (o *QbdSalesReceiptLineGroup) GetLines() []QbdSalesReceiptLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdSalesReceiptLineGroup) GetLinesOk() (*[]QbdSalesReceiptLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdSalesReceiptLineGroup) SetLines(v []QbdSalesReceiptLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdSalesReceiptLineGroup) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesReceiptLineGroup) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesReceiptLineGroup) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


