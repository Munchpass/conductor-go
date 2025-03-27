# QbdSalesReceiptLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales receipt line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_receipt_line\&quot;&#x60;. | 
**Item** | [**QbdSalesReceiptLineItem**](QbdSalesReceiptLineItem.md) |  | 
**Description** | **string** | A description of this sales receipt line. | 
**Quantity** | **float32** | The quantity of the item associated with this sales receipt line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales receipt line. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdSalesReceiptLineOverrideUnitOfMeasureSet**](QbdSalesReceiptLineOverrideUnitOfMeasureSet.md) |  | 
**Rate** | **string** | The price per unit for this sales receipt line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | 
**RatePercent** | **string** | The price of this sales receipt line expressed as a percentage. Typically used for discount or markup items. | 
**Class** | [**QbdSalesReceiptLineClass**](QbdSalesReceiptLineClass.md) |  | 
**Amount** | **string** | The monetary amount of this sales receipt line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | 
**InventorySite** | [**QbdSalesReceiptLineInventorySite**](QbdSalesReceiptLineInventorySite.md) |  | 
**InventorySiteLocation** | [**QbdSalesReceiptLineInventorySiteLocation**](QbdSalesReceiptLineInventorySiteLocation.md) |  | 
**SerialNumber** | **string** | The serial number of the item associated with this sales receipt line. This is used for tracking individual units of serialized inventory items. | 
**LotNumber** | **string** | The lot number of the item associated with this sales receipt line. Used for tracking groups of inventory items that are purchased or manufactured together. | 
**ExpirationDate** | **string** | The expiration date for the serial number or lot number of the item associated with this sales receipt line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | 
**ServiceDate** | **string** | The date on which the service for this sales receipt line was or will be performed, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for service items. | 
**SalesTaxCode** | [**QbdSalesReceiptLineSalesTaxCode**](QbdSalesReceiptLineSalesTaxCode.md) |  | 
**OtherCustomField1** | **string** | A built-in custom field for additional information specific to this sales receipt line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipt lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**OtherCustomField2** | **string** | A second built-in custom field for additional information specific to this sales receipt line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipt lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**CreditCardTransaction** | [**QbdSalesReceiptLineCreditCardTransaction**](QbdSalesReceiptLineCreditCardTransaction.md) |  | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales receipt line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesReceiptLine

`func NewQbdSalesReceiptLine(id string, objectType string, item QbdSalesReceiptLineItem, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdSalesReceiptLineOverrideUnitOfMeasureSet, rate string, ratePercent string, class QbdSalesReceiptLineClass, amount string, inventorySite QbdSalesReceiptLineInventorySite, inventorySiteLocation QbdSalesReceiptLineInventorySiteLocation, serialNumber string, lotNumber string, expirationDate string, serviceDate string, salesTaxCode QbdSalesReceiptLineSalesTaxCode, otherCustomField1 string, otherCustomField2 string, creditCardTransaction QbdSalesReceiptLineCreditCardTransaction, customFields []QbdCustomField, ) *QbdSalesReceiptLine`

NewQbdSalesReceiptLine instantiates a new QbdSalesReceiptLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesReceiptLineWithDefaults

`func NewQbdSalesReceiptLineWithDefaults() *QbdSalesReceiptLine`

NewQbdSalesReceiptLineWithDefaults instantiates a new QbdSalesReceiptLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesReceiptLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesReceiptLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesReceiptLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesReceiptLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesReceiptLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesReceiptLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdSalesReceiptLine) GetItem() QbdSalesReceiptLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdSalesReceiptLine) GetItemOk() (*QbdSalesReceiptLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdSalesReceiptLine) SetItem(v QbdSalesReceiptLineItem)`

SetItem sets Item field to given value.


### GetDescription

`func (o *QbdSalesReceiptLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesReceiptLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesReceiptLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdSalesReceiptLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdSalesReceiptLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdSalesReceiptLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdSalesReceiptLine) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdSalesReceiptLine) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdSalesReceiptLine) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdSalesReceiptLine) GetOverrideUnitOfMeasureSet() QbdSalesReceiptLineOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdSalesReceiptLine) GetOverrideUnitOfMeasureSetOk() (*QbdSalesReceiptLineOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdSalesReceiptLine) SetOverrideUnitOfMeasureSet(v QbdSalesReceiptLineOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetRate

`func (o *QbdSalesReceiptLine) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QbdSalesReceiptLine) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QbdSalesReceiptLine) SetRate(v string)`

SetRate sets Rate field to given value.


### GetRatePercent

`func (o *QbdSalesReceiptLine) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QbdSalesReceiptLine) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QbdSalesReceiptLine) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.


### GetClass

`func (o *QbdSalesReceiptLine) GetClass() QbdSalesReceiptLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdSalesReceiptLine) GetClassOk() (*QbdSalesReceiptLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdSalesReceiptLine) SetClass(v QbdSalesReceiptLineClass)`

SetClass sets Class field to given value.


### GetAmount

`func (o *QbdSalesReceiptLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdSalesReceiptLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdSalesReceiptLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetInventorySite

`func (o *QbdSalesReceiptLine) GetInventorySite() QbdSalesReceiptLineInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdSalesReceiptLine) GetInventorySiteOk() (*QbdSalesReceiptLineInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdSalesReceiptLine) SetInventorySite(v QbdSalesReceiptLineInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetInventorySiteLocation

`func (o *QbdSalesReceiptLine) GetInventorySiteLocation() QbdSalesReceiptLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdSalesReceiptLine) GetInventorySiteLocationOk() (*QbdSalesReceiptLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdSalesReceiptLine) SetInventorySiteLocation(v QbdSalesReceiptLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetSerialNumber

`func (o *QbdSalesReceiptLine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QbdSalesReceiptLine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QbdSalesReceiptLine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetLotNumber

`func (o *QbdSalesReceiptLine) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QbdSalesReceiptLine) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QbdSalesReceiptLine) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.


### GetExpirationDate

`func (o *QbdSalesReceiptLine) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QbdSalesReceiptLine) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QbdSalesReceiptLine) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetServiceDate

`func (o *QbdSalesReceiptLine) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *QbdSalesReceiptLine) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *QbdSalesReceiptLine) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetSalesTaxCode

`func (o *QbdSalesReceiptLine) GetSalesTaxCode() QbdSalesReceiptLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdSalesReceiptLine) GetSalesTaxCodeOk() (*QbdSalesReceiptLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdSalesReceiptLine) SetSalesTaxCode(v QbdSalesReceiptLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetOtherCustomField1

`func (o *QbdSalesReceiptLine) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QbdSalesReceiptLine) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QbdSalesReceiptLine) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.


### GetOtherCustomField2

`func (o *QbdSalesReceiptLine) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QbdSalesReceiptLine) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QbdSalesReceiptLine) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.


### GetCreditCardTransaction

`func (o *QbdSalesReceiptLine) GetCreditCardTransaction() QbdSalesReceiptLineCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QbdSalesReceiptLine) GetCreditCardTransactionOk() (*QbdSalesReceiptLineCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QbdSalesReceiptLine) SetCreditCardTransaction(v QbdSalesReceiptLineCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.


### GetCustomFields

`func (o *QbdSalesReceiptLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesReceiptLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesReceiptLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


