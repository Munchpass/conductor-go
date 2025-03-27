# QbdPriceLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this price level. This ID is unique across all price levels but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_price_level\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this price level was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this price level was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this price level object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this price level, unique across all price levels.  **NOTE**: Price levels do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this price level is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**PriceLevelType** | **string** | The price level&#39;s type. | 
**FixedPercentage** | **string** | The fixed percentage adjustment applied to all items for this price level (instead of a per-item price level). Once you create the price level, you cannot change this.  When this price level is applied to a customer, it automatically adjusts the &#x60;rate&#x60; and &#x60;amount&#x60; columns for applicable line items in sales orders and invoices for that customer. This value supports both positive and negative values - a value of \&quot;20\&quot; increases prices by 20%, while \&quot;-10\&quot; decreases prices by 10%. | 
**PerItemPriceLevels** | [**[]QbdPerItemPriceLevel**](QbdPerItemPriceLevel.md) | The per-item price level configurations for this price level. | 
**Currency** | [**QbdPriceLevelCurrency**](QbdPriceLevelCurrency.md) |  | 

## Methods

### NewQbdPriceLevel

`func NewQbdPriceLevel(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, priceLevelType string, fixedPercentage string, perItemPriceLevels []QbdPerItemPriceLevel, currency QbdPriceLevelCurrency, ) *QbdPriceLevel`

NewQbdPriceLevel instantiates a new QbdPriceLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPriceLevelWithDefaults

`func NewQbdPriceLevelWithDefaults() *QbdPriceLevel`

NewQbdPriceLevelWithDefaults instantiates a new QbdPriceLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdPriceLevel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdPriceLevel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdPriceLevel) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdPriceLevel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdPriceLevel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdPriceLevel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdPriceLevel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdPriceLevel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdPriceLevel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdPriceLevel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdPriceLevel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdPriceLevel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdPriceLevel) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdPriceLevel) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdPriceLevel) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdPriceLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdPriceLevel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdPriceLevel) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdPriceLevel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdPriceLevel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdPriceLevel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPriceLevelType

`func (o *QbdPriceLevel) GetPriceLevelType() string`

GetPriceLevelType returns the PriceLevelType field if non-nil, zero value otherwise.

### GetPriceLevelTypeOk

`func (o *QbdPriceLevel) GetPriceLevelTypeOk() (*string, bool)`

GetPriceLevelTypeOk returns a tuple with the PriceLevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelType

`func (o *QbdPriceLevel) SetPriceLevelType(v string)`

SetPriceLevelType sets PriceLevelType field to given value.


### GetFixedPercentage

`func (o *QbdPriceLevel) GetFixedPercentage() string`

GetFixedPercentage returns the FixedPercentage field if non-nil, zero value otherwise.

### GetFixedPercentageOk

`func (o *QbdPriceLevel) GetFixedPercentageOk() (*string, bool)`

GetFixedPercentageOk returns a tuple with the FixedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPercentage

`func (o *QbdPriceLevel) SetFixedPercentage(v string)`

SetFixedPercentage sets FixedPercentage field to given value.


### GetPerItemPriceLevels

`func (o *QbdPriceLevel) GetPerItemPriceLevels() []QbdPerItemPriceLevel`

GetPerItemPriceLevels returns the PerItemPriceLevels field if non-nil, zero value otherwise.

### GetPerItemPriceLevelsOk

`func (o *QbdPriceLevel) GetPerItemPriceLevelsOk() (*[]QbdPerItemPriceLevel, bool)`

GetPerItemPriceLevelsOk returns a tuple with the PerItemPriceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerItemPriceLevels

`func (o *QbdPriceLevel) SetPerItemPriceLevels(v []QbdPerItemPriceLevel)`

SetPerItemPriceLevels sets PerItemPriceLevels field to given value.


### GetCurrency

`func (o *QbdPriceLevel) GetCurrency() QbdPriceLevelCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdPriceLevel) GetCurrencyOk() (*QbdPriceLevelCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdPriceLevel) SetCurrency(v QbdPriceLevelCurrency)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


