# QbdCreditMemoLineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this credit memo line group. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_credit_memo_line_group\&quot;&#x60;. | 
**ItemGroup** | [**QbdCreditMemoLineGroupItemGroup**](QbdCreditMemoLineGroupItemGroup.md) |  | 
**Description** | **string** | A description of this credit memo line group. | 
**Quantity** | **float32** | The quantity of the item group associated with this credit memo line group. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item group is a discount item group. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this credit memo line group. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdCreditMemoLineGroupOverrideUnitOfMeasureSet**](QbdCreditMemoLineGroupOverrideUnitOfMeasureSet.md) |  | 
**ShouldPrintItemsInGroup** | **bool** | Indicates whether the individual items in this credit memo line group and their separate amounts appear on printed forms. | 
**TotalAmount** | **string** | The total monetary amount of this credit memo line group, equivalent to the sum of the amounts in &#x60;lines&#x60;, represented as a decimal string. | 
**Lines** | [**[]QbdCreditMemoLine**](QbdCreditMemoLine.md) | The credit memo line group&#39;s line items, each representing a single product or service sold. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the credit memo line group object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdCreditMemoLineGroup

`func NewQbdCreditMemoLineGroup(id string, objectType string, itemGroup QbdCreditMemoLineGroupItemGroup, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdCreditMemoLineGroupOverrideUnitOfMeasureSet, shouldPrintItemsInGroup bool, totalAmount string, lines []QbdCreditMemoLine, customFields []QbdCustomField, ) *QbdCreditMemoLineGroup`

NewQbdCreditMemoLineGroup instantiates a new QbdCreditMemoLineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditMemoLineGroupWithDefaults

`func NewQbdCreditMemoLineGroupWithDefaults() *QbdCreditMemoLineGroup`

NewQbdCreditMemoLineGroupWithDefaults instantiates a new QbdCreditMemoLineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCreditMemoLineGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCreditMemoLineGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCreditMemoLineGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdCreditMemoLineGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdCreditMemoLineGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdCreditMemoLineGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemGroup

`func (o *QbdCreditMemoLineGroup) GetItemGroup() QbdCreditMemoLineGroupItemGroup`

GetItemGroup returns the ItemGroup field if non-nil, zero value otherwise.

### GetItemGroupOk

`func (o *QbdCreditMemoLineGroup) GetItemGroupOk() (*QbdCreditMemoLineGroupItemGroup, bool)`

GetItemGroupOk returns a tuple with the ItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroup

`func (o *QbdCreditMemoLineGroup) SetItemGroup(v QbdCreditMemoLineGroupItemGroup)`

SetItemGroup sets ItemGroup field to given value.


### GetDescription

`func (o *QbdCreditMemoLineGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdCreditMemoLineGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdCreditMemoLineGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdCreditMemoLineGroup) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdCreditMemoLineGroup) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdCreditMemoLineGroup) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdCreditMemoLineGroup) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdCreditMemoLineGroup) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdCreditMemoLineGroup) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdCreditMemoLineGroup) GetOverrideUnitOfMeasureSet() QbdCreditMemoLineGroupOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdCreditMemoLineGroup) GetOverrideUnitOfMeasureSetOk() (*QbdCreditMemoLineGroupOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdCreditMemoLineGroup) SetOverrideUnitOfMeasureSet(v QbdCreditMemoLineGroupOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetShouldPrintItemsInGroup

`func (o *QbdCreditMemoLineGroup) GetShouldPrintItemsInGroup() bool`

GetShouldPrintItemsInGroup returns the ShouldPrintItemsInGroup field if non-nil, zero value otherwise.

### GetShouldPrintItemsInGroupOk

`func (o *QbdCreditMemoLineGroup) GetShouldPrintItemsInGroupOk() (*bool, bool)`

GetShouldPrintItemsInGroupOk returns a tuple with the ShouldPrintItemsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrintItemsInGroup

`func (o *QbdCreditMemoLineGroup) SetShouldPrintItemsInGroup(v bool)`

SetShouldPrintItemsInGroup sets ShouldPrintItemsInGroup field to given value.


### GetTotalAmount

`func (o *QbdCreditMemoLineGroup) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdCreditMemoLineGroup) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdCreditMemoLineGroup) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetLines

`func (o *QbdCreditMemoLineGroup) GetLines() []QbdCreditMemoLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdCreditMemoLineGroup) GetLinesOk() (*[]QbdCreditMemoLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdCreditMemoLineGroup) SetLines(v []QbdCreditMemoLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdCreditMemoLineGroup) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCreditMemoLineGroup) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCreditMemoLineGroup) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


