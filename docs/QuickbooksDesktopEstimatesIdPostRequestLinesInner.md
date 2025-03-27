# QuickbooksDesktopEstimatesIdPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing estimate line you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new estimate lines you wish to add. | 
**ItemId** | Pointer to **string** | The item associated with this estimate line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**Description** | Pointer to **string** | A description of this estimate line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this estimate line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this estimate line. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**OverrideUnitOfMeasureSetId** | Pointer to **string** | Specifies an alternative unit-of-measure set when updating this estimate line&#39;s &#x60;unitOfMeasure&#x60; field (e.g., \&quot;pound\&quot; or \&quot;kilogram\&quot;). This allows you to select units from a different set than the item&#39;s default unit-of-measure set, which remains unchanged on the item itself. The override applies only to this specific line. For example, you can sell an item typically measured in volume units using weight units in a specific transaction by specifying a different unit-of-measure set with this field. | [optional] 
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
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this estimate line. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimate lines for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 

## Methods

### NewQuickbooksDesktopEstimatesIdPostRequestLinesInner

`func NewQuickbooksDesktopEstimatesIdPostRequestLinesInner(id string, ) *QuickbooksDesktopEstimatesIdPostRequestLinesInner`

NewQuickbooksDesktopEstimatesIdPostRequestLinesInner instantiates a new QuickbooksDesktopEstimatesIdPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEstimatesIdPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopEstimatesIdPostRequestLinesInnerWithDefaults() *QuickbooksDesktopEstimatesIdPostRequestLinesInner`

NewQuickbooksDesktopEstimatesIdPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopEstimatesIdPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOverrideUnitOfMeasureSetId() string`

GetOverrideUnitOfMeasureSetId returns the OverrideUnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOverrideUnitOfMeasureSetIdOk() (*string, bool)`

GetOverrideUnitOfMeasureSetIdOk returns a tuple with the OverrideUnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetOverrideUnitOfMeasureSetId(v string)`

SetOverrideUnitOfMeasureSetId sets OverrideUnitOfMeasureSetId field to given value.

### HasOverrideUnitOfMeasureSetId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasOverrideUnitOfMeasureSetId() bool`

HasOverrideUnitOfMeasureSetId returns a boolean if a field has been set.

### GetRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetRatePercent() string`

GetRatePercent returns the RatePercent field if non-nil, zero value otherwise.

### GetRatePercentOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetRatePercentOk() (*string, bool)`

GetRatePercentOk returns a tuple with the RatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetRatePercent(v string)`

SetRatePercent sets RatePercent field to given value.

### HasRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasRatePercent() bool`

HasRatePercent returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetPriceRuleConflictStrategy() string`

GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field if non-nil, zero value otherwise.

### GetPriceRuleConflictStrategyOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool)`

GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetPriceRuleConflictStrategy(v string)`

SetPriceRuleConflictStrategy sets PriceRuleConflictStrategy field to given value.

### HasPriceRuleConflictStrategy

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasPriceRuleConflictStrategy() bool`

HasPriceRuleConflictStrategy returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetMarkupRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetMarkupRate() string`

GetMarkupRate returns the MarkupRate field if non-nil, zero value otherwise.

### GetMarkupRateOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetMarkupRateOk() (*string, bool)`

GetMarkupRateOk returns a tuple with the MarkupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetMarkupRate(v string)`

SetMarkupRate sets MarkupRate field to given value.

### HasMarkupRate

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasMarkupRate() bool`

HasMarkupRate returns a boolean if a field has been set.

### GetMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetMarkupRatePercent() string`

GetMarkupRatePercent returns the MarkupRatePercent field if non-nil, zero value otherwise.

### GetMarkupRatePercentOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetMarkupRatePercentOk() (*string, bool)`

GetMarkupRatePercentOk returns a tuple with the MarkupRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetMarkupRatePercent(v string)`

SetMarkupRatePercent sets MarkupRatePercent field to given value.

### HasMarkupRatePercent

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasMarkupRatePercent() bool`

HasMarkupRatePercent returns a boolean if a field has been set.

### GetPriceLevelId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetPriceLevelId() string`

GetPriceLevelId returns the PriceLevelId field if non-nil, zero value otherwise.

### GetPriceLevelIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool)`

GetPriceLevelIdOk returns a tuple with the PriceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetPriceLevelId(v string)`

SetPriceLevelId sets PriceLevelId field to given value.

### HasPriceLevelId

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasPriceLevelId() bool`

HasPriceLevelId returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopEstimatesIdPostRequestLinesInner) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


