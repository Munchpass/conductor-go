# QuickbooksDesktopSalesReceiptsPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The item associated with this sales receipt line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**Description** | Pointer to **string** | A description of this sales receipt line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this sales receipt line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this sales receipt line. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**Rate** | Pointer to **string** | The price per unit for this sales receipt line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | [optional] 
**RatePercent** | Pointer to **string** | The price of this sales receipt line expressed as a percentage. Typically used for discount or markup items. | [optional] 
**PriceLevelId** | Pointer to **string** | The price level applied to this sales receipt line. This overrides any price level set on the corresponding customer. The resulting sales receipt line will not show this price level, only the final &#x60;rate&#x60; calculated from it. | [optional] 
**ClassId** | Pointer to **string** | The sales receipt line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all sales receipt lines unless overridden here, at the transaction line level. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this sales receipt line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | [optional] 
**PriceRuleConflictStrategy** | Pointer to **string** | Specifies how to resolve price rule conflicts when adding or modifying this sales receipt line. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this sales receipt line is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this sales receipt line is stored. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the item associated with this sales receipt line. This is used for tracking individual units of serialized inventory items. | [optional] 
**LotNumber** | Pointer to **string** | The lot number of the item associated with this sales receipt line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**ServiceDate** | Pointer to **string** | The date on which the service for this sales receipt line was or will be performed, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for service items. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this sales receipt line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OverrideItemAccountId** | Pointer to **string** | The account to use for this sales receipt line, overriding the default account associated with the item. | [optional] 
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this sales receipt line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipt lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this sales receipt line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipt lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**CreditCardTransaction** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestLinesInnerCreditCardTransaction**](QuickbooksDesktopSalesReceiptsPostRequestLinesInnerCreditCardTransaction.md) |  | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the sales receipt line object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopSalesReceiptsPostRequestLinesInner

`func NewQuickbooksDesktopSalesReceiptsPostRequestLinesInner() *QuickbooksDesktopSalesReceiptsPostRequestLinesInner`

NewQuickbooksDesktopSalesReceiptsPostRequestLinesInner instantiates a new QuickbooksDesktopSalesReceiptsPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesReceiptsPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopSalesReceiptsPostRequestLinesInnerWithDefaults() *QuickbooksDesktopSalesReceiptsPostRequestLinesInner`

NewQuickbooksDesktopSalesReceiptsPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopSalesReceiptsPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRatePercent

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.

### HasRatePercent

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasRatePercent() bool`

HasRatePercent returns a boolean if a field has been set.

### GetPriceLevelId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetPriceLevelId() string`

GetPriceLevelId returns the PriceLevelId field if non-nil, zero value otherwise.

### GetPriceLevelIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool)`

GetPriceLevelIdOk returns a tuple with the PriceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetPriceLevelId(v string)`

SetPriceLevelId sets PriceLevelId field to given value.

### HasPriceLevelId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasPriceLevelId() bool`

HasPriceLevelId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetPriceRuleConflictStrategy() string`

GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field if non-nil, zero value otherwise.

### GetPriceRuleConflictStrategyOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool)`

GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetPriceRuleConflictStrategy(v string)`

SetPriceRuleConflictStrategy sets PriceRuleConflictStrategy field to given value.

### HasPriceRuleConflictStrategy

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasPriceRuleConflictStrategy() bool`

HasPriceRuleConflictStrategy returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetServiceDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.

### HasServiceDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasServiceDate() bool`

HasServiceDate returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOverrideItemAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOverrideItemAccountId() string`

GetOverrideItemAccountId returns the OverrideItemAccountId field if non-nil, zero value otherwise.

### GetOverrideItemAccountIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOverrideItemAccountIdOk() (*string, bool)`

GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideItemAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetOverrideItemAccountId(v string)`

SetOverrideItemAccountId sets OverrideItemAccountId field to given value.

### HasOverrideItemAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasOverrideItemAccountId() bool`

HasOverrideItemAccountId returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.

### GetCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetCreditCardTransaction() QuickbooksDesktopSalesReceiptsPostRequestLinesInnerCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetCreditCardTransactionOk() (*QuickbooksDesktopSalesReceiptsPostRequestLinesInnerCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetCreditCardTransaction(v QuickbooksDesktopSalesReceiptsPostRequestLinesInnerCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.

### HasCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasCreditCardTransaction() bool`

HasCreditCardTransaction returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopSalesReceiptsPostRequestLinesInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


