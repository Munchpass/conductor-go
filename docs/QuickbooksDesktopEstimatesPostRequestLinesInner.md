# QuickbooksDesktopEstimatesPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The item associated with this estimate line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**Description** | Pointer to **string** | A description of this estimate line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this estimate line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this estimate line. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**Rate** | Pointer to **string** | The price per unit for this estimate line. If both &#x60;rate&#x60; and &#x60;amount&#x60; are specified, &#x60;rate&#x60; will be ignored. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;rate&#x60;, QuickBooks will use them to calculate &#x60;rate&#x60;. Represented as a decimal string. This field cannot be cleared. | [optional] 
**RatePercent** | Pointer to **string** | The price of this estimate line expressed as a percentage. Typically used for discount or markup items. | [optional] 
**ClassId** | Pointer to **string** | The estimate line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all estimate lines unless overridden here, at the transaction line level. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this estimate line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;rate&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;rate&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;rate&#x60;. This field cannot be cleared. | [optional] 
**PriceRuleConflictStrategy** | Pointer to **string** | Specifies how to resolve price rule conflicts when adding or modifying this estimate line. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this estimate line is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this estimate line is stored. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this estimate line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**MarkupRate** | Pointer to **string** | The markup that will be passed on to the customer for this item on this estimate line. &#x60;amount &#x3D; (quantity * rate) * (1 + markupRate)&#x60; | [optional] 
**MarkupRatePercent** | Pointer to **string** | The markup, expressed as a percentage, that will be passed on to the customer for this item on this estimate line. &#x60;amount &#x3D; (quantity * rate) * (1 + markupRatePercent/100)&#x60; | [optional] 
**PriceLevelId** | Pointer to **string** | The price level applied to this estimate line. This overrides any price level set on the corresponding customer. The resulting estimate line will not show this price level, only the final &#x60;rate&#x60; calculated from it. | [optional] 
**OverrideItemAccountId** | Pointer to **string** | The account to use for this estimate line, overriding the default account associated with the item. | [optional] 
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the estimate line object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopEstimatesPostRequestLinesInner

`func NewQuickbooksDesktopEstimatesPostRequestLinesInner() *QuickbooksDesktopEstimatesPostRequestLinesInner`

NewQuickbooksDesktopEstimatesPostRequestLinesInner instantiates a new QuickbooksDesktopEstimatesPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEstimatesPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopEstimatesPostRequestLinesInnerWithDefaults() *QuickbooksDesktopEstimatesPostRequestLinesInner`

NewQuickbooksDesktopEstimatesPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopEstimatesPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.

### HasRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasRatePercent() bool`

HasRatePercent returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetPriceRuleConflictStrategy() string`

GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field if non-nil, zero value otherwise.

### GetPriceRuleConflictStrategyOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool)`

GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetPriceRuleConflictStrategy(v string)`

SetPriceRuleConflictStrategy sets PriceRuleConflictStrategy field to given value.

### HasPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasPriceRuleConflictStrategy() bool`

HasPriceRuleConflictStrategy returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetMarkupRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetMarkupRate() string`

GetMarkupRate returns the MarkupRate field if non-nil, zero value otherwise.

### GetMarkupRateOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetMarkupRateOk() (*string, bool)`

GetMarkupRateOk returns a tuple with the MarkupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetMarkupRate(v string)`

SetMarkupRate sets MarkupRate field to given value.

### HasMarkupRate

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasMarkupRate() bool`

HasMarkupRate returns a boolean if a field has been set.

### GetMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetMarkupRatePercent() string`

GetMarkupRatePercent returns the MarkupRatePercent field if non-nil, zero value otherwise.

### GetMarkupRatePercentOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetMarkupRatePercentOk() (*string, bool)`

GetMarkupRatePercentOk returns a tuple with the MarkupRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetMarkupRatePercent(v string)`

SetMarkupRatePercent sets MarkupRatePercent field to given value.

### HasMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasMarkupRatePercent() bool`

HasMarkupRatePercent returns a boolean if a field has been set.

### GetPriceLevelId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetPriceLevelId() string`

GetPriceLevelId returns the PriceLevelId field if non-nil, zero value otherwise.

### GetPriceLevelIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool)`

GetPriceLevelIdOk returns a tuple with the PriceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetPriceLevelId(v string)`

SetPriceLevelId sets PriceLevelId field to given value.

### HasPriceLevelId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasPriceLevelId() bool`

HasPriceLevelId returns a boolean if a field has been set.

### GetOverrideItemAccountId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOverrideItemAccountId() string`

GetOverrideItemAccountId returns the OverrideItemAccountId field if non-nil, zero value otherwise.

### GetOverrideItemAccountIdOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOverrideItemAccountIdOk() (*string, bool)`

GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideItemAccountId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetOverrideItemAccountId(v string)`

SetOverrideItemAccountId sets OverrideItemAccountId field to given value.

### HasOverrideItemAccountId

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasOverrideItemAccountId() bool`

HasOverrideItemAccountId returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopEstimatesPostRequestLinesInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


