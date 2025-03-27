# QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The item associated with this per-item price level. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**CustomPrice** | Pointer to **string** | The fixed amount custom price for this per-item price level that overrides the standard price for the specified item. Used when setting an absolute price value for the item in this price level. | [optional] 
**CustomPricePercent** | Pointer to **string** | The fixed discount percentage for this per-item price level that modifies the specified item&#39;s standard price. Used to create a fixed percentage markup or discount specific to this item within this price level. | [optional] 
**AdjustPercentage** | **string** | The percentage adjustment for this per-item price level when using relative pricing. Specifies a percentage to modify pricing, using positive values (e.g., \&quot;20\&quot;) to increase prices by that percentage, or negative values (e.g., \&quot;-10\&quot;) to apply a discount. | 
**AdjustRelativeTo** | **string** | The base value reference for this per-item price level&#39;s percentage adjustment. Specifies which price to use as the starting point for the adjustment calculation in the price level.  **NOTE:** The price level must use either a fixed pricing approach (&#x60;customPrice&#x60; or &#x60;customPricePercent&#x60;) or a relative adjustment approach (&#x60;adjustPercentage&#x60; with &#x60;adjustRelativeTo&#x60;) when configuring per-item price levels. | 

## Methods

### NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner

`func NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner(adjustPercentage string, adjustRelativeTo string, ) *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner`

NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner instantiates a new QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInnerWithDefaults

`func NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInnerWithDefaults() *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner`

NewQuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInnerWithDefaults instantiates a new QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetCustomPrice

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetCustomPrice() string`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetCustomPriceOk() (*string, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) SetCustomPrice(v string)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetCustomPricePercent

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetCustomPricePercent() string`

GetCustomPricePercent returns the CustomPricePercent field if non-nil, zero value otherwise.

### GetCustomPricePercentOk

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetCustomPricePercentOk() (*string, bool)`

GetCustomPricePercentOk returns a tuple with the CustomPricePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePercent

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) SetCustomPricePercent(v string)`

SetCustomPricePercent sets CustomPricePercent field to given value.

### HasCustomPricePercent

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) HasCustomPricePercent() bool`

HasCustomPricePercent returns a boolean if a field has been set.

### GetAdjustPercentage

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetAdjustPercentage() string`

GetAdjustPercentage returns the AdjustPercentage field if non-nil, zero value otherwise.

### GetAdjustPercentageOk

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetAdjustPercentageOk() (*string, bool)`

GetAdjustPercentageOk returns a tuple with the AdjustPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustPercentage

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) SetAdjustPercentage(v string)`

SetAdjustPercentage sets AdjustPercentage field to given value.


### GetAdjustRelativeTo

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetAdjustRelativeTo() string`

GetAdjustRelativeTo returns the AdjustRelativeTo field if non-nil, zero value otherwise.

### GetAdjustRelativeToOk

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) GetAdjustRelativeToOk() (*string, bool)`

GetAdjustRelativeToOk returns a tuple with the AdjustRelativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustRelativeTo

`func (o *QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner) SetAdjustRelativeTo(v string)`

SetAdjustRelativeTo sets AdjustRelativeTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


