# QbdInvoiceLineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this invoice line group. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_invoice_line_group\&quot;&#x60;. | 
**ItemGroup** | [**QbdInvoiceLineGroupItemGroup**](QbdInvoiceLineGroupItemGroup.md) |  | 
**Description** | **string** | A description of this invoice line group. | 
**Quantity** | **float32** | The quantity of the item group associated with this invoice line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this invoice line group. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdInvoiceLineGroupOverrideUnitOfMeasureSet**](QbdInvoiceLineGroupOverrideUnitOfMeasureSet.md) |  | 
**ShouldPrintItemsInGroup** | **bool** | Indicates whether the individual items in this invoice line group and their separate amounts appear on printed forms. | 
**TotalAmount** | **string** | The total monetary amount of this invoice line group, equivalent to the sum of the amounts in &#x60;lines&#x60;, represented as a decimal string. | 
**Lines** | [**[]QbdInvoiceLine**](QbdInvoiceLine.md) | The invoice line group&#39;s line items, each representing a single product or service sold. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the invoice line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdInvoiceLineGroup

`func NewQbdInvoiceLineGroup(id string, objectType string, itemGroup QbdInvoiceLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdInvoiceLineGroupOverrideUnitOfMeasureSet, shouldPrintItemsInGroup bool, totalAmount string, lines []QbdInvoiceLine, customFields []QbdCustomField, ) *QbdInvoiceLineGroup`

NewQbdInvoiceLineGroup instantiates a new QbdInvoiceLineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInvoiceLineGroupWithDefaults

`func NewQbdInvoiceLineGroupWithDefaults() *QbdInvoiceLineGroup`

NewQbdInvoiceLineGroupWithDefaults instantiates a new QbdInvoiceLineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInvoiceLineGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInvoiceLineGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInvoiceLineGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInvoiceLineGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInvoiceLineGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInvoiceLineGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemGroup

`func (o *QbdInvoiceLineGroup) GetItemGroup() QbdInvoiceLineGroupItemGroup`

GetItemGroup returns the ItemGroup field if non-nil, zero value otherwise.

### GetItemGroupOk

`func (o *QbdInvoiceLineGroup) GetItemGroupOk() (*QbdInvoiceLineGroupItemGroup, bool)`

GetItemGroupOk returns a tuple with the ItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroup

`func (o *QbdInvoiceLineGroup) SetItemGroup(v QbdInvoiceLineGroupItemGroup)`

SetItemGroup sets ItemGroup field to given value.


### GetDescription

`func (o *QbdInvoiceLineGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdInvoiceLineGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdInvoiceLineGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdInvoiceLineGroup) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdInvoiceLineGroup) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdInvoiceLineGroup) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdInvoiceLineGroup) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdInvoiceLineGroup) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdInvoiceLineGroup) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdInvoiceLineGroup) GetOverrideUnitOfMeasureSet() QbdInvoiceLineGroupOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdInvoiceLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdInvoiceLineGroupOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdInvoiceLineGroup) SetOverrideUnitOfMeasureSet(v QbdInvoiceLineGroupOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetShouldPrintItemsInGroup

`func (o *QbdInvoiceLineGroup) GetShouldPrintItemsInGroup() bool`

GetShouldPrintItemsInGroup returns the ShouldPrintItemsInGroup field if non-nil, zero value otherwise.

### GetShouldPrintItemsInGroupOk

`func (o *QbdInvoiceLineGroup) GetShouldPrintItemsInGroupOk() (*bool, bool)`

GetShouldPrintItemsInGroupOk returns a tuple with the ShouldPrintItemsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrintItemsInGroup

`func (o *QbdInvoiceLineGroup) SetShouldPrintItemsInGroup(v bool)`

SetShouldPrintItemsInGroup sets ShouldPrintItemsInGroup field to given value.


### GetTotalAmount

`func (o *QbdInvoiceLineGroup) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdInvoiceLineGroup) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdInvoiceLineGroup) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetLines

`func (o *QbdInvoiceLineGroup) GetLines() []QbdInvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdInvoiceLineGroup) GetLinesOk() (*[]QbdInvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdInvoiceLineGroup) SetLines(v []QbdInvoiceLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdInvoiceLineGroup) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdInvoiceLineGroup) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdInvoiceLineGroup) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


