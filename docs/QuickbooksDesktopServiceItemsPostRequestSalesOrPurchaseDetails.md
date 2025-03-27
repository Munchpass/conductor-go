# QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of this item. | [optional] 
**Price** | Pointer to **string** | The price at which this item is purchased or sold, represented as a decimal string. | [optional] 
**PricePercentage** | Pointer to **string** | The price of this item expressed as a percentage, used instead of &#x60;price&#x60; when the item&#39;s cost is calculated as a percentage of another amount. For example, a service item that costs a percentage of another item&#39;s price. | [optional] 
**PostingAccountId** | **string** | The posting account to which transactions involving this item are posted. This could be an income account when selling or an expense account when purchasing. | 

## Methods

### NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails

`func NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails(postingAccountId string, ) *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails`

NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails instantiates a new QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetailsWithDefaults

`func NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetailsWithDefaults() *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails`

NewQuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetailsWithDefaults instantiates a new QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPricePercentage

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPricePercentage() string`

GetPricePercentage returns the PricePercentage field if non-nil, zero value otherwise.

### GetPricePercentageOk

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPricePercentageOk() (*string, bool)`

GetPricePercentageOk returns a tuple with the PricePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePercentage

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) SetPricePercentage(v string)`

SetPricePercentage sets PricePercentage field to given value.

### HasPricePercentage

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) HasPricePercentage() bool`

HasPricePercentage returns a boolean if a field has been set.

### GetPostingAccountId

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPostingAccountId() string`

GetPostingAccountId returns the PostingAccountId field if non-nil, zero value otherwise.

### GetPostingAccountIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) GetPostingAccountIdOk() (*string, bool)`

GetPostingAccountIdOk returns a tuple with the PostingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAccountId

`func (o *QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails) SetPostingAccountId(v string)`

SetPostingAccountId sets PostingAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


