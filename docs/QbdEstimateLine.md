# QbdEstimateLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this estimate line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_estimate_line\&quot;&#x60;. | 
**Item** | [**QbdEstimateLineItem**](QbdEstimateLineItem.md) |  | 
**Description** | **string** | A description of this estimate line. | 
**Quantity** | **float32** | The quantity of the item associated with this estimate line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this estimate line. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdEstimateLineOverrideUnitOfMeasureSet**](QbdEstimateLineOverrideUnitOfMeasureSet.md) |  | 
**Rate** | **string** | The price per unit for this estimate line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | 
**RatePercent** | **string** | The price of this estimate line expressed as a percentage. Typically used for discount or markup items. | 
**Class** | [**QbdEstimateLineClass**](QbdEstimateLineClass.md) |  | 
**Amount** | **string** | The monetary amount of this estimate line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | 
**InventorySite** | [**QbdEstimateLineInventorySite**](QbdEstimateLineInventorySite.md) |  | 
**InventorySiteLocation** | [**QbdEstimateLineInventorySiteLocation**](QbdEstimateLineInventorySiteLocation.md) |  | 
**SalesTaxCode** | [**QbdEstimateLineSalesTaxCode**](QbdEstimateLineSalesTaxCode.md) |  | 
**MarkupRate** | **string** | The markup that will be passed on to the customer for this item on this estimate line. &#x60;amount &#x3D; (quantity * rate) * (1 + markupRate)&#x60; | 
**MarkupRatePercent** | **string** | The markup, expressed as a percentage, that will be passed on to the customer for this item on this estimate line. &#x60;amount &#x3D; (quantity * rate) * (1 + markupRatePercent/100)&#x60; | 
**OtherCustomField1** | **string** | A built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**OtherCustomField2** | **string** | A second built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the estimate line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdEstimateLine

`func NewQbdEstimateLine(id string, objectType string, item QbdEstimateLineItem, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdEstimateLineOverrideUnitOfMeasureSet, rate string, ratePercent string, class QbdEstimateLineClass, amount string, inventorySite QbdEstimateLineInventorySite, inventorySiteLocation QbdEstimateLineInventorySiteLocation, salesTaxCode QbdEstimateLineSalesTaxCode, markupRate string, markupRatePercent string, otherCustomField1 string, otherCustomField2 string, customFields []QbdCustomField, ) *QbdEstimateLine`

NewQbdEstimateLine instantiates a new QbdEstimateLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEstimateLineWithDefaults

`func NewQbdEstimateLineWithDefaults() *QbdEstimateLine`

NewQbdEstimateLineWithDefaults instantiates a new QbdEstimateLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdEstimateLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdEstimateLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdEstimateLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdEstimateLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdEstimateLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdEstimateLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdEstimateLine) GetItem() QbdEstimateLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdEstimateLine) GetItemOk() (*QbdEstimateLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdEstimateLine) SetItem(v QbdEstimateLineItem)`

SetItem sets Item field to given value.


### GetDescription

`func (o *QbdEstimateLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdEstimateLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdEstimateLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdEstimateLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdEstimateLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdEstimateLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdEstimateLine) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdEstimateLine) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdEstimateLine) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdEstimateLine) GetOverrideUnitOfMeasureSet() QbdEstimateLineOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdEstimateLine) GetOverrideUnitOfMeasureSetOk() (*QbdEstimateLineOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdEstimateLine) SetOverrideUnitOfMeasureSet(v QbdEstimateLineOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetRate

`func (o *QbdEstimateLine) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QbdEstimateLine) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QbdEstimateLine) SetRate(v string)`

SetRate sets Rate field to given value.


### GetRatePercent

`func (o *QbdEstimateLine) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QbdEstimateLine) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QbdEstimateLine) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.


### GetClass

`func (o *QbdEstimateLine) GetClass() QbdEstimateLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdEstimateLine) GetClassOk() (*QbdEstimateLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdEstimateLine) SetClass(v QbdEstimateLineClass)`

SetClass sets Class field to given value.


### GetAmount

`func (o *QbdEstimateLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdEstimateLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdEstimateLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetInventorySite

`func (o *QbdEstimateLine) GetInventorySite() QbdEstimateLineInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdEstimateLine) GetInventorySiteOk() (*QbdEstimateLineInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdEstimateLine) SetInventorySite(v QbdEstimateLineInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetInventorySiteLocation

`func (o *QbdEstimateLine) GetInventorySiteLocation() QbdEstimateLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdEstimateLine) GetInventorySiteLocationOk() (*QbdEstimateLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdEstimateLine) SetInventorySiteLocation(v QbdEstimateLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetSalesTaxCode

`func (o *QbdEstimateLine) GetSalesTaxCode() QbdEstimateLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdEstimateLine) GetSalesTaxCodeOk() (*QbdEstimateLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdEstimateLine) SetSalesTaxCode(v QbdEstimateLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetMarkupRate

`func (o *QbdEstimateLine) GetMarkupRate() string`

GetMarkupRate returns the MarkupRate field if non-nil, zero value otherwise.

### GetMarkupRateOk

`func (o *QbdEstimateLine) GetMarkupRateOk() (*string, bool)`

GetMarkupRateOk returns a tuple with the MarkupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRate

`func (o *QbdEstimateLine) SetMarkupRate(v string)`

SetMarkupRate sets MarkupRate field to given value.


### GetMarkupRatePercent

`func (o *QbdEstimateLine) GetMarkupRatePercent() string`

GetMarkupRatePercent returns the MarkupRatePercent field if non-nil, zero value otherwise.

### GetMarkupRatePercentOk

`func (o *QbdEstimateLine) GetMarkupRatePercentOk() (*string, bool)`

GetMarkupRatePercentOk returns a tuple with the MarkupRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRatePercent

`func (o *QbdEstimateLine) SetMarkupRatePercent(v string)`

SetMarkupRatePercent sets MarkupRatePercent field to given value.


### GetOtherCustomField1

`func (o *QbdEstimateLine) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QbdEstimateLine) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QbdEstimateLine) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.


### GetOtherCustomField2

`func (o *QbdEstimateLine) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QbdEstimateLine) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QbdEstimateLine) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.


### GetCustomFields

`func (o *QbdEstimateLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdEstimateLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdEstimateLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


