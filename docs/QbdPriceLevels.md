# QbdPriceLevels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsingPriceLevels** | **bool** | Indicates whether this company file has price levels enabled. When &#x60;true&#x60;, price levels can be created and used to automatically calculate custom pricing for different customers. | 
**IsRoundingSalesPriceUp** | **NullableBool** | Indicates whether this company file is configured to round amounts up to the nearest whole dollar for fixed percentage price levels. This setting does not affect per-item price levels. | 

## Methods

### NewQbdPriceLevels

`func NewQbdPriceLevels(isUsingPriceLevels bool, isRoundingSalesPriceUp NullableBool, ) *QbdPriceLevels`

NewQbdPriceLevels instantiates a new QbdPriceLevels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPriceLevelsWithDefaults

`func NewQbdPriceLevelsWithDefaults() *QbdPriceLevels`

NewQbdPriceLevelsWithDefaults instantiates a new QbdPriceLevels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsingPriceLevels

`func (o *QbdPriceLevels) GetIsUsingPriceLevels() bool`

GetIsUsingPriceLevels returns the IsUsingPriceLevels field if non-nil, zero value otherwise.

### GetIsUsingPriceLevelsOk

`func (o *QbdPriceLevels) GetIsUsingPriceLevelsOk() (*bool, bool)`

GetIsUsingPriceLevelsOk returns a tuple with the IsUsingPriceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingPriceLevels

`func (o *QbdPriceLevels) SetIsUsingPriceLevels(v bool)`

SetIsUsingPriceLevels sets IsUsingPriceLevels field to given value.


### GetIsRoundingSalesPriceUp

`func (o *QbdPriceLevels) GetIsRoundingSalesPriceUp() bool`

GetIsRoundingSalesPriceUp returns the IsRoundingSalesPriceUp field if non-nil, zero value otherwise.

### GetIsRoundingSalesPriceUpOk

`func (o *QbdPriceLevels) GetIsRoundingSalesPriceUpOk() (*bool, bool)`

GetIsRoundingSalesPriceUpOk returns a tuple with the IsRoundingSalesPriceUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundingSalesPriceUp

`func (o *QbdPriceLevels) SetIsRoundingSalesPriceUp(v bool)`

SetIsRoundingSalesPriceUp sets IsRoundingSalesPriceUp field to given value.


### SetIsRoundingSalesPriceUpNil

`func (o *QbdPriceLevels) SetIsRoundingSalesPriceUpNil(b bool)`

 SetIsRoundingSalesPriceUpNil sets the value for IsRoundingSalesPriceUp to be an explicit nil

### UnsetIsRoundingSalesPriceUp
`func (o *QbdPriceLevels) UnsetIsRoundingSalesPriceUp()`

UnsetIsRoundingSalesPriceUp ensures that no value is present for IsRoundingSalesPriceUp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


