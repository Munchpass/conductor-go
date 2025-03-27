# QuickbooksDesktopPriceLevelsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this price level, unique across all price levels.  **NOTE**: Price levels do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this price level is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**FixedPercentage** | Pointer to **string** | The fixed percentage adjustment applied to all items for this price level (instead of a per-item price level). Once you create the price level, you cannot change this.  When this price level is applied to a customer, it automatically adjusts the &#x60;rate&#x60; and &#x60;amount&#x60; columns for applicable line items in sales orders and invoices for that customer. This value supports both positive and negative values - a value of \&quot;20\&quot; increases prices by 20%, while \&quot;-10\&quot; decreases prices by 10%. | [optional] 
**PerItemPriceLevels** | Pointer to [**[]QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner**](QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner.md) | The per-item price level configurations for this price level. | [optional] 
**CurrencyId** | Pointer to **string** | The price level&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 

## Methods

### NewQuickbooksDesktopPriceLevelsPostRequest

`func NewQuickbooksDesktopPriceLevelsPostRequest(name string, ) *QuickbooksDesktopPriceLevelsPostRequest`

NewQuickbooksDesktopPriceLevelsPostRequest instantiates a new QuickbooksDesktopPriceLevelsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPriceLevelsPostRequestWithDefaults

`func NewQuickbooksDesktopPriceLevelsPostRequestWithDefaults() *QuickbooksDesktopPriceLevelsPostRequest`

NewQuickbooksDesktopPriceLevelsPostRequestWithDefaults instantiates a new QuickbooksDesktopPriceLevelsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopPriceLevelsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopPriceLevelsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopPriceLevelsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetFixedPercentage

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetFixedPercentage() string`

GetFixedPercentage returns the FixedPercentage field if non-nil, zero value otherwise.

### GetFixedPercentageOk

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetFixedPercentageOk() (*string, bool)`

GetFixedPercentageOk returns a tuple with the FixedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPercentage

`func (o *QuickbooksDesktopPriceLevelsPostRequest) SetFixedPercentage(v string)`

SetFixedPercentage sets FixedPercentage field to given value.

### HasFixedPercentage

`func (o *QuickbooksDesktopPriceLevelsPostRequest) HasFixedPercentage() bool`

HasFixedPercentage returns a boolean if a field has been set.

### GetPerItemPriceLevels

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetPerItemPriceLevels() []QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner`

GetPerItemPriceLevels returns the PerItemPriceLevels field if non-nil, zero value otherwise.

### GetPerItemPriceLevelsOk

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetPerItemPriceLevelsOk() (*[]QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner, bool)`

GetPerItemPriceLevelsOk returns a tuple with the PerItemPriceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerItemPriceLevels

`func (o *QuickbooksDesktopPriceLevelsPostRequest) SetPerItemPriceLevels(v []QuickbooksDesktopPriceLevelsPostRequestPerItemPriceLevelsInner)`

SetPerItemPriceLevels sets PerItemPriceLevels field to given value.

### HasPerItemPriceLevels

`func (o *QuickbooksDesktopPriceLevelsPostRequest) HasPerItemPriceLevels() bool`

HasPerItemPriceLevels returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopPriceLevelsPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopPriceLevelsPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopPriceLevelsPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


