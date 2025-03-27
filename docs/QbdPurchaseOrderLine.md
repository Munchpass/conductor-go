# QbdPurchaseOrderLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this purchase order line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_purchase_order_line\&quot;&#x60;. | 
**Item** | [**QbdPurchaseOrderLineItem**](QbdPurchaseOrderLineItem.md) |  | 
**Sku** | **string** | The purchase order line&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | 
**Description** | **string** | A description of this purchase order line. | 
**Quantity** | **float32** | The quantity of the item associated with this purchase order line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this purchase order line. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdPurchaseOrderLineOverrideUnitOfMeasureSet**](QbdPurchaseOrderLineOverrideUnitOfMeasureSet.md) |  | 
**Rate** | **string** | The price per unit for this purchase order line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | 
**Class** | [**QbdPurchaseOrderLineClass**](QbdPurchaseOrderLineClass.md) |  | 
**Amount** | **string** | The monetary amount of this purchase order line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | 
**InventorySiteLocation** | [**QbdPurchaseOrderLineInventorySiteLocation**](QbdPurchaseOrderLineInventorySiteLocation.md) |  | 
**Payee** | [**QbdExpenseLinePayee**](QbdExpenseLinePayee.md) |  | 
**ServiceDate** | **string** | The date on which the service for this purchase order line was or will be performed, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for service items. | 
**SalesTaxCode** | [**QbdPurchaseOrderLineSalesTaxCode**](QbdPurchaseOrderLineSalesTaxCode.md) |  | 
**ReceivedQuantity** | **float32** | The quantity that has been received against this purchase order line. | 
**UnbilledQuantity** | **float32** | The quantity that has not been billed for this purchase order line. | 
**IsBilled** | **bool** | Indicates whether this purchase order line has been billed. | 
**IsManuallyClosed** | **bool** | Indicates whether this purchase order line has been manually marked as closed, even if this item has not been received or its sale has not been cancelled. If all the purchase order lines are marked as closed, the purchase order itself is marked as closed as well. You cannot change &#x60;isManuallyClosed&#x60; to &#x60;false&#x60; after the purchase order line has been fully received. | 
**OtherCustomField1** | **string** | A built-in custom field for additional information specific to this purchase order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase order lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**OtherCustomField2** | **string** | A second built-in custom field for additional information specific to this purchase order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase order lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the purchase order line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdPurchaseOrderLine

`func NewQbdPurchaseOrderLine(id string, objectType string, item QbdPurchaseOrderLineItem, sku string, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdPurchaseOrderLineOverrideUnitOfMeasureSet, rate string, class QbdPurchaseOrderLineClass, amount string, inventorySiteLocation QbdPurchaseOrderLineInventorySiteLocation, payee QbdExpenseLinePayee, serviceDate string, salesTaxCode QbdPurchaseOrderLineSalesTaxCode, receivedQuantity float32, unbilledQuantity float32, isBilled bool, isManuallyClosed bool, otherCustomField1 string, otherCustomField2 string, customFields []QbdCustomField, ) *QbdPurchaseOrderLine`

NewQbdPurchaseOrderLine instantiates a new QbdPurchaseOrderLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPurchaseOrderLineWithDefaults

`func NewQbdPurchaseOrderLineWithDefaults() *QbdPurchaseOrderLine`

NewQbdPurchaseOrderLineWithDefaults instantiates a new QbdPurchaseOrderLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdPurchaseOrderLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdPurchaseOrderLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdPurchaseOrderLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdPurchaseOrderLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdPurchaseOrderLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdPurchaseOrderLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdPurchaseOrderLine) GetItem() QbdPurchaseOrderLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdPurchaseOrderLine) GetItemOk() (*QbdPurchaseOrderLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdPurchaseOrderLine) SetItem(v QbdPurchaseOrderLineItem)`

SetItem sets Item field to given value.


### GetSku

`func (o *QbdPurchaseOrderLine) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QbdPurchaseOrderLine) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QbdPurchaseOrderLine) SetSku(v string)`

SetSku sets Sku field to given value.


### GetDescription

`func (o *QbdPurchaseOrderLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdPurchaseOrderLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdPurchaseOrderLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdPurchaseOrderLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdPurchaseOrderLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdPurchaseOrderLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdPurchaseOrderLine) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdPurchaseOrderLine) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdPurchaseOrderLine) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdPurchaseOrderLine) GetOverrideUnitOfMeasureSet() QbdPurchaseOrderLineOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdPurchaseOrderLine) GetOverrideUnitOfMeasureSetOk() (*QbdPurchaseOrderLineOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdPurchaseOrderLine) SetOverrideUnitOfMeasureSet(v QbdPurchaseOrderLineOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetRate

`func (o *QbdPurchaseOrderLine) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QbdPurchaseOrderLine) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QbdPurchaseOrderLine) SetRate(v string)`

SetRate sets Rate field to given value.


### GetClass

`func (o *QbdPurchaseOrderLine) GetClass() QbdPurchaseOrderLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdPurchaseOrderLine) GetClassOk() (*QbdPurchaseOrderLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdPurchaseOrderLine) SetClass(v QbdPurchaseOrderLineClass)`

SetClass sets Class field to given value.


### GetAmount

`func (o *QbdPurchaseOrderLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdPurchaseOrderLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdPurchaseOrderLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetInventorySiteLocation

`func (o *QbdPurchaseOrderLine) GetInventorySiteLocation() QbdPurchaseOrderLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdPurchaseOrderLine) GetInventorySiteLocationOk() (*QbdPurchaseOrderLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdPurchaseOrderLine) SetInventorySiteLocation(v QbdPurchaseOrderLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetPayee

`func (o *QbdPurchaseOrderLine) GetPayee() QbdExpenseLinePayee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *QbdPurchaseOrderLine) GetPayeeOk() (*QbdExpenseLinePayee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *QbdPurchaseOrderLine) SetPayee(v QbdExpenseLinePayee)`

SetPayee sets Payee field to given value.


### GetServiceDate

`func (o *QbdPurchaseOrderLine) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *QbdPurchaseOrderLine) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *QbdPurchaseOrderLine) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetSalesTaxCode

`func (o *QbdPurchaseOrderLine) GetSalesTaxCode() QbdPurchaseOrderLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdPurchaseOrderLine) GetSalesTaxCodeOk() (*QbdPurchaseOrderLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdPurchaseOrderLine) SetSalesTaxCode(v QbdPurchaseOrderLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetReceivedQuantity

`func (o *QbdPurchaseOrderLine) GetReceivedQuantity() float32`

GetReceivedQuantity returns the ReceivedQuantity field if non-nil, zero value otherwise.

### GetReceivedQuantityOk

`func (o *QbdPurchaseOrderLine) GetReceivedQuantityOk() (*float32, bool)`

GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQuantity

`func (o *QbdPurchaseOrderLine) SetReceivedQuantity(v float32)`

SetReceivedQuantity sets ReceivedQuantity field to given value.


### GetUnbilledQuantity

`func (o *QbdPurchaseOrderLine) GetUnbilledQuantity() float32`

GetUnbilledQuantity returns the UnbilledQuantity field if non-nil, zero value otherwise.

### GetUnbilledQuantityOk

`func (o *QbdPurchaseOrderLine) GetUnbilledQuantityOk() (*float32, bool)`

GetUnbilledQuantityOk returns a tuple with the UnbilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledQuantity

`func (o *QbdPurchaseOrderLine) SetUnbilledQuantity(v float32)`

SetUnbilledQuantity sets UnbilledQuantity field to given value.


### GetIsBilled

`func (o *QbdPurchaseOrderLine) GetIsBilled() bool`

GetIsBilled returns the IsBilled field if non-nil, zero value otherwise.

### GetIsBilledOk

`func (o *QbdPurchaseOrderLine) GetIsBilledOk() (*bool, bool)`

GetIsBilledOk returns a tuple with the IsBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBilled

`func (o *QbdPurchaseOrderLine) SetIsBilled(v bool)`

SetIsBilled sets IsBilled field to given value.


### GetIsManuallyClosed

`func (o *QbdPurchaseOrderLine) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QbdPurchaseOrderLine) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QbdPurchaseOrderLine) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.


### GetOtherCustomField1

`func (o *QbdPurchaseOrderLine) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QbdPurchaseOrderLine) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QbdPurchaseOrderLine) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.


### GetOtherCustomField2

`func (o *QbdPurchaseOrderLine) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QbdPurchaseOrderLine) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QbdPurchaseOrderLine) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.


### GetCustomFields

`func (o *QbdPurchaseOrderLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdPurchaseOrderLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdPurchaseOrderLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


