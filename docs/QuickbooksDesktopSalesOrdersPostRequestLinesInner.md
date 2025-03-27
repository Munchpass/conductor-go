# QuickbooksDesktopSalesOrdersPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The item associated with this sales order line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**Description** | Pointer to **string** | A description of this sales order line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this sales order line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales order line. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**Rate** | Pointer to **string** | The price per unit for this sales order line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | [optional] 
**RatePercent** | Pointer to **string** | The price of this sales order line expressed as a percentage. Typically used for discount or markup items. | [optional] 
**PriceLevelId** | Pointer to **string** | The price level applied to this sales order line. This overrides any price level set on the corresponding customer. The resulting sales order line will not show this price level, only the final &#x60;rate&#x60; calculated from it. | [optional] 
**ClassId** | Pointer to **string** | The sales order line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all sales order lines unless overridden here, at the transaction line level. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this sales order line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | [optional] 
**PriceRuleConflictStrategy** | Pointer to **string** | Specifies how to resolve price rule conflicts when adding or modifying this sales order line. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this sales order line is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this sales order line is stored. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the item associated with this sales order line. This is used for tracking individual units of serialized inventory items. | [optional] 
**LotNumber** | Pointer to **string** | The lot number of the item associated with this sales order line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this sales order line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**IsManuallyClosed** | Pointer to **bool** | Indicates whether this sales order line has been manually marked as closed, even if it has not been invoiced. | [optional] [default to false]
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the sales order line object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopSalesOrdersPostRequestLinesInner

`func NewQuickbooksDesktopSalesOrdersPostRequestLinesInner() *QuickbooksDesktopSalesOrdersPostRequestLinesInner`

NewQuickbooksDesktopSalesOrdersPostRequestLinesInner instantiates a new QuickbooksDesktopSalesOrdersPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesOrdersPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopSalesOrdersPostRequestLinesInnerWithDefaults() *QuickbooksDesktopSalesOrdersPostRequestLinesInner`

NewQuickbooksDesktopSalesOrdersPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopSalesOrdersPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetRate

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRatePercent

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.

### HasRatePercent

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasRatePercent() bool`

HasRatePercent returns a boolean if a field has been set.

### GetPriceLevelId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetPriceLevelId() string`

GetPriceLevelId returns the PriceLevelId field if non-nil, zero value otherwise.

### GetPriceLevelIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool)`

GetPriceLevelIdOk returns a tuple with the PriceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetPriceLevelId(v string)`

SetPriceLevelId sets PriceLevelId field to given value.

### HasPriceLevelId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasPriceLevelId() bool`

HasPriceLevelId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetPriceRuleConflictStrategy() string`

GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field if non-nil, zero value otherwise.

### GetPriceRuleConflictStrategyOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool)`

GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetPriceRuleConflictStrategy(v string)`

SetPriceRuleConflictStrategy sets PriceRuleConflictStrategy field to given value.

### HasPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasPriceRuleConflictStrategy() bool`

HasPriceRuleConflictStrategy returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.

### HasIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasIsManuallyClosed() bool`

HasIsManuallyClosed returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopSalesOrdersPostRequestLinesInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


