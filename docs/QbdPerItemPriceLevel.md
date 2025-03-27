# QbdPerItemPriceLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**QbdPerItemPriceLevelItem**](QbdPerItemPriceLevelItem.md) |  | 
**CustomPrice** | **string** | The fixed amount custom price for this per-item price level that overrides the standard price for the specified item. Used when setting an absolute price value for the item in this price level. | 
**CustomPricePercent** | **string** | The fixed discount percentage for this per-item price level that modifies the specified item&#39;s standard price. Used to create a fixed percentage markup or discount specific to this item within this price level. | 

## Methods

### NewQbdPerItemPriceLevel

`func NewQbdPerItemPriceLevel(item QbdPerItemPriceLevelItem, customPrice string, customPricePercent string, ) *QbdPerItemPriceLevel`

NewQbdPerItemPriceLevel instantiates a new QbdPerItemPriceLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPerItemPriceLevelWithDefaults

`func NewQbdPerItemPriceLevelWithDefaults() *QbdPerItemPriceLevel`

NewQbdPerItemPriceLevelWithDefaults instantiates a new QbdPerItemPriceLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *QbdPerItemPriceLevel) GetItem() QbdPerItemPriceLevelItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdPerItemPriceLevel) GetItemOk() (*QbdPerItemPriceLevelItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdPerItemPriceLevel) SetItem(v QbdPerItemPriceLevelItem)`

SetItem sets Item field to given value.


### GetCustomPrice

`func (o *QbdPerItemPriceLevel) GetCustomPrice() string`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *QbdPerItemPriceLevel) GetCustomPriceOk() (*string, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *QbdPerItemPriceLevel) SetCustomPrice(v string)`

SetCustomPrice sets CustomPrice field to given value.


### GetCustomPricePercent

`func (o *QbdPerItemPriceLevel) GetCustomPricePercent() string`

GetCustomPricePercent returns the CustomPricePercent field if non-nil, zero value otherwise.

### GetCustomPricePercentOk

`func (o *QbdPerItemPriceLevel) GetCustomPricePercentOk() (*string, bool)`

GetCustomPricePercentOk returns a tuple with the CustomPricePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePercent

`func (o *QbdPerItemPriceLevel) SetCustomPricePercent(v string)`

SetCustomPricePercent sets CustomPricePercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


