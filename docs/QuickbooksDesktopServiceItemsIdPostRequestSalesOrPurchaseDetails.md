# QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of this item. | [optional] 
**Price** | Pointer to **string** | The price at which this item is purchased or sold, represented as a decimal string. | [optional] 
**PricePercentage** | Pointer to **string** | The price of this item expressed as a percentage, used instead of &#x60;price&#x60; when the item&#39;s cost is calculated as a percentage of another amount. For example, a service item that costs a percentage of another item&#39;s price. | [optional] 
**PostingAccountId** | Pointer to **string** | The posting account to which transactions involving this item are posted. This could be an income account when selling or an expense account when purchasing. | [optional] 
**UpdateExistingTransactionsAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new account (specified by the &#x60;accountId&#x60; field) to all existing transactions associated with this item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 

## Methods

### NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails

`func NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails() *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails`

NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails instantiates a new QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetailsWithDefaults

`func NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetailsWithDefaults() *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails`

NewQuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetailsWithDefaults instantiates a new QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPricePercentage

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPricePercentage() string`

GetPricePercentage returns the PricePercentage field if non-nil, zero value otherwise.

### GetPricePercentageOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPricePercentageOk() (*string, bool)`

GetPricePercentageOk returns a tuple with the PricePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePercentage

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) SetPricePercentage(v string)`

SetPricePercentage sets PricePercentage field to given value.

### HasPricePercentage

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) HasPricePercentage() bool`

HasPricePercentage returns a boolean if a field has been set.

### GetPostingAccountId

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPostingAccountId() string`

GetPostingAccountId returns the PostingAccountId field if non-nil, zero value otherwise.

### GetPostingAccountIdOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetPostingAccountIdOk() (*string, bool)`

GetPostingAccountIdOk returns a tuple with the PostingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAccountId

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) SetPostingAccountId(v string)`

SetPostingAccountId sets PostingAccountId field to given value.

### HasPostingAccountId

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) HasPostingAccountId() bool`

HasPostingAccountId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetUpdateExistingTransactionsAccount() bool`

GetUpdateExistingTransactionsAccount returns the UpdateExistingTransactionsAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsAccountOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) GetUpdateExistingTransactionsAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsAccountOk returns a tuple with the UpdateExistingTransactionsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) SetUpdateExistingTransactionsAccount(v bool)`

SetUpdateExistingTransactionsAccount sets UpdateExistingTransactionsAccount field to given value.

### HasUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails) HasUpdateExistingTransactionsAccount() bool`

HasUpdateExistingTransactionsAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


