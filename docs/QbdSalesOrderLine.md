# QbdSalesOrderLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales order line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_order_line\&quot;&#x60;. | 
**Item** | [**QbdSalesOrderLineItem**](QbdSalesOrderLineItem.md) |  | 
**Description** | **string** | A description of this sales order line. | 
**Quantity** | **float32** | The quantity of the item associated with this sales order line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales order line. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdSalesOrderLineOverrideUnitOfMeasureSet**](QbdSalesOrderLineOverrideUnitOfMeasureSet.md) |  | 
**Rate** | **string** | The price per unit for this sales order line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | 
**RatePercent** | **string** | The price of this sales order line expressed as a percentage. Typically used for discount or markup items. | 
**Class** | [**QbdSalesOrderLineClass**](QbdSalesOrderLineClass.md) |  | 
**Amount** | **string** | The monetary amount of this sales order line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | 
**InventorySite** | [**QbdSalesOrderLineInventorySite**](QbdSalesOrderLineInventorySite.md) |  | 
**InventorySiteLocation** | [**QbdSalesOrderLineInventorySiteLocation**](QbdSalesOrderLineInventorySiteLocation.md) |  | 
**SerialNumber** | **string** | The serial number of the item associated with this sales order line. This is used for tracking individual units of serialized inventory items. | 
**LotNumber** | **string** | The lot number of the item associated with this sales order line. Used for tracking groups of inventory items that are purchased or manufactured together. | 
**ExpirationDate** | **string** | The expiration date for the serial number or lot number of the item associated with this sales order line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | 
**SalesTaxCode** | [**QbdSalesOrderLineSalesTaxCode**](QbdSalesOrderLineSalesTaxCode.md) |  | 
**QuantityInvoiced** | **float32** | The number of units of this sales order line&#39;s &#x60;quantity&#x60; that have been invoiced. | 
**IsManuallyClosed** | **bool** | Indicates whether this sales order line has been manually marked as closed, even if it has not been invoiced. | 
**OtherCustomField1** | **string** | A built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**OtherCustomField2** | **string** | A second built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales order line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesOrderLine

`func NewQbdSalesOrderLine(id string, objectType string, item QbdSalesOrderLineItem, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdSalesOrderLineOverrideUnitOfMeasureSet, rate string, ratePercent string, class QbdSalesOrderLineClass, amount string, inventorySite QbdSalesOrderLineInventorySite, inventorySiteLocation QbdSalesOrderLineInventorySiteLocation, serialNumber string, lotNumber string, expirationDate string, salesTaxCode QbdSalesOrderLineSalesTaxCode, quantityInvoiced float32, isManuallyClosed bool, otherCustomField1 string, otherCustomField2 string, customFields []QbdCustomField, ) *QbdSalesOrderLine`

NewQbdSalesOrderLine instantiates a new QbdSalesOrderLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesOrderLineWithDefaults

`func NewQbdSalesOrderLineWithDefaults() *QbdSalesOrderLine`

NewQbdSalesOrderLineWithDefaults instantiates a new QbdSalesOrderLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesOrderLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesOrderLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesOrderLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesOrderLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesOrderLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesOrderLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdSalesOrderLine) GetItem() QbdSalesOrderLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdSalesOrderLine) GetItemOk() (*QbdSalesOrderLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdSalesOrderLine) SetItem(v QbdSalesOrderLineItem)`

SetItem sets Item field to given value.


### GetDescription

`func (o *QbdSalesOrderLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesOrderLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesOrderLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdSalesOrderLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdSalesOrderLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdSalesOrderLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdSalesOrderLine) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdSalesOrderLine) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdSalesOrderLine) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdSalesOrderLine) GetOverrideUnitOfMeasureSet() QbdSalesOrderLineOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdSalesOrderLine) GetOverrideUnitOfMeasureSetOk() (*QbdSalesOrderLineOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdSalesOrderLine) SetOverrideUnitOfMeasureSet(v QbdSalesOrderLineOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetRate

`func (o *QbdSalesOrderLine) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QbdSalesOrderLine) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QbdSalesOrderLine) SetRate(v string)`

SetRate sets Rate field to given value.


### GetRatePercent

`func (o *QbdSalesOrderLine) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QbdSalesOrderLine) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QbdSalesOrderLine) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.


### GetClass

`func (o *QbdSalesOrderLine) GetClass() QbdSalesOrderLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdSalesOrderLine) GetClassOk() (*QbdSalesOrderLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdSalesOrderLine) SetClass(v QbdSalesOrderLineClass)`

SetClass sets Class field to given value.


### GetAmount

`func (o *QbdSalesOrderLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdSalesOrderLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdSalesOrderLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetInventorySite

`func (o *QbdSalesOrderLine) GetInventorySite() QbdSalesOrderLineInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdSalesOrderLine) GetInventorySiteOk() (*QbdSalesOrderLineInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdSalesOrderLine) SetInventorySite(v QbdSalesOrderLineInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetInventorySiteLocation

`func (o *QbdSalesOrderLine) GetInventorySiteLocation() QbdSalesOrderLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdSalesOrderLine) GetInventorySiteLocationOk() (*QbdSalesOrderLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdSalesOrderLine) SetInventorySiteLocation(v QbdSalesOrderLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetSerialNumber

`func (o *QbdSalesOrderLine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QbdSalesOrderLine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QbdSalesOrderLine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetLotNumber

`func (o *QbdSalesOrderLine) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QbdSalesOrderLine) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QbdSalesOrderLine) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.


### GetExpirationDate

`func (o *QbdSalesOrderLine) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QbdSalesOrderLine) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QbdSalesOrderLine) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetSalesTaxCode

`func (o *QbdSalesOrderLine) GetSalesTaxCode() QbdSalesOrderLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdSalesOrderLine) GetSalesTaxCodeOk() (*QbdSalesOrderLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdSalesOrderLine) SetSalesTaxCode(v QbdSalesOrderLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetQuantityInvoiced

`func (o *QbdSalesOrderLine) GetQuantityInvoiced() float32`

GetQuantityInvoiced returns the QuantityInvoiced field if non-nil, zero value otherwise.

### GetQuantityInvoicedOk

`func (o *QbdSalesOrderLine) GetQuantityInvoicedOk() (*float32, bool)`

GetQuantityInvoicedOk returns a tuple with the QuantityInvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityInvoiced

`func (o *QbdSalesOrderLine) SetQuantityInvoiced(v float32)`

SetQuantityInvoiced sets QuantityInvoiced field to given value.


### GetIsManuallyClosed

`func (o *QbdSalesOrderLine) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QbdSalesOrderLine) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QbdSalesOrderLine) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.


### GetOtherCustomField1

`func (o *QbdSalesOrderLine) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QbdSalesOrderLine) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QbdSalesOrderLine) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.


### GetOtherCustomField2

`func (o *QbdSalesOrderLine) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QbdSalesOrderLine) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QbdSalesOrderLine) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.


### GetCustomFields

`func (o *QbdSalesOrderLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesOrderLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesOrderLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


