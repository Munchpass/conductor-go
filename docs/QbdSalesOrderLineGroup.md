# QbdSalesOrderLineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales order line group. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_order_line_group\&quot;&#x60;. | 
**ItemGroup** | [**QbdSalesOrderLineGroupItemGroup**](QbdSalesOrderLineGroupItemGroup.md) |  | 
**Description** | **string** | A description of this sales order line group. | 
**Quantity** | **float32** | The quantity of the item group associated with this sales order line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales order line group. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdSalesOrderLineGroupOverrideUnitOfMeasureSet**](QbdSalesOrderLineGroupOverrideUnitOfMeasureSet.md) |  | 
**ShouldPrintItemsInGroup** | **bool** | Indicates whether the individual items in this sales order line group and their separate amounts appear on printed forms. | 
**TotalAmount** | **string** | The total monetary amount of this sales order line group, equivalent to the sum of the amounts in &#x60;lines&#x60;, represented as a decimal string. | 
**Lines** | [**[]QbdSalesOrderLine**](QbdSalesOrderLine.md) | The sales order line group&#39;s line items, each representing a single product or service ordered. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales order line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesOrderLineGroup

`func NewQbdSalesOrderLineGroup(id string, objectType string, itemGroup QbdSalesOrderLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdSalesOrderLineGroupOverrideUnitOfMeasureSet, shouldPrintItemsInGroup bool, totalAmount string, lines []QbdSalesOrderLine, customFields []QbdCustomField, ) *QbdSalesOrderLineGroup`

NewQbdSalesOrderLineGroup instantiates a new QbdSalesOrderLineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesOrderLineGroupWithDefaults

`func NewQbdSalesOrderLineGroupWithDefaults() *QbdSalesOrderLineGroup`

NewQbdSalesOrderLineGroupWithDefaults instantiates a new QbdSalesOrderLineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesOrderLineGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesOrderLineGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesOrderLineGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesOrderLineGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesOrderLineGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesOrderLineGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemGroup

`func (o *QbdSalesOrderLineGroup) GetItemGroup() QbdSalesOrderLineGroupItemGroup`

GetItemGroup returns the ItemGroup field if non-nil, zero value otherwise.

### GetItemGroupOk

`func (o *QbdSalesOrderLineGroup) GetItemGroupOk() (*QbdSalesOrderLineGroupItemGroup, bool)`

GetItemGroupOk returns a tuple with the ItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroup

`func (o *QbdSalesOrderLineGroup) SetItemGroup(v QbdSalesOrderLineGroupItemGroup)`

SetItemGroup sets ItemGroup field to given value.


### GetDescription

`func (o *QbdSalesOrderLineGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesOrderLineGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesOrderLineGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdSalesOrderLineGroup) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdSalesOrderLineGroup) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdSalesOrderLineGroup) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdSalesOrderLineGroup) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdSalesOrderLineGroup) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdSalesOrderLineGroup) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdSalesOrderLineGroup) GetOverrideUnitOfMeasureSet() QbdSalesOrderLineGroupOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdSalesOrderLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdSalesOrderLineGroupOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdSalesOrderLineGroup) SetOverrideUnitOfMeasureSet(v QbdSalesOrderLineGroupOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetShouldPrintItemsInGroup

`func (o *QbdSalesOrderLineGroup) GetShouldPrintItemsInGroup() bool`

GetShouldPrintItemsInGroup returns the ShouldPrintItemsInGroup field if non-nil, zero value otherwise.

### GetShouldPrintItemsInGroupOk

`func (o *QbdSalesOrderLineGroup) GetShouldPrintItemsInGroupOk() (*bool, bool)`

GetShouldPrintItemsInGroupOk returns a tuple with the ShouldPrintItemsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrintItemsInGroup

`func (o *QbdSalesOrderLineGroup) SetShouldPrintItemsInGroup(v bool)`

SetShouldPrintItemsInGroup sets ShouldPrintItemsInGroup field to given value.


### GetTotalAmount

`func (o *QbdSalesOrderLineGroup) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdSalesOrderLineGroup) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdSalesOrderLineGroup) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetLines

`func (o *QbdSalesOrderLineGroup) GetLines() []QbdSalesOrderLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdSalesOrderLineGroup) GetLinesOk() (*[]QbdSalesOrderLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdSalesOrderLineGroup) SetLines(v []QbdSalesOrderLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdSalesOrderLineGroup) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesOrderLineGroup) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesOrderLineGroup) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


