# QbdServiceItemSalesOrPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of this item. | 
**Price** | **string** | The price at which this item is purchased or sold, represented as a decimal string. | 
**PricePercentage** | **string** | The price of this item expressed as a percentage, used instead of &#x60;price&#x60; when the item&#39;s cost is calculated as a percentage of another amount. For example, a service item that costs a percentage of another item&#39;s price. | 
**PostingAccount** | [**QbdSalesOrPurchaseDetailsPostingAccount**](QbdSalesOrPurchaseDetailsPostingAccount.md) |  | 

## Methods

### NewQbdServiceItemSalesOrPurchaseDetails

`func NewQbdServiceItemSalesOrPurchaseDetails(description string, price string, pricePercentage string, postingAccount QbdSalesOrPurchaseDetailsPostingAccount, ) *QbdServiceItemSalesOrPurchaseDetails`

NewQbdServiceItemSalesOrPurchaseDetails instantiates a new QbdServiceItemSalesOrPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdServiceItemSalesOrPurchaseDetailsWithDefaults

`func NewQbdServiceItemSalesOrPurchaseDetailsWithDefaults() *QbdServiceItemSalesOrPurchaseDetails`

NewQbdServiceItemSalesOrPurchaseDetailsWithDefaults instantiates a new QbdServiceItemSalesOrPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdServiceItemSalesOrPurchaseDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QbdServiceItemSalesOrPurchaseDetails) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetPricePercentage

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPricePercentage() string`

GetPricePercentage returns the PricePercentage field if non-nil, zero value otherwise.

### GetPricePercentageOk

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPricePercentageOk() (*string, bool)`

GetPricePercentageOk returns a tuple with the PricePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePercentage

`func (o *QbdServiceItemSalesOrPurchaseDetails) SetPricePercentage(v string)`

SetPricePercentage sets PricePercentage field to given value.


### GetPostingAccount

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPostingAccount() QbdSalesOrPurchaseDetailsPostingAccount`

GetPostingAccount returns the PostingAccount field if non-nil, zero value otherwise.

### GetPostingAccountOk

`func (o *QbdServiceItemSalesOrPurchaseDetails) GetPostingAccountOk() (*QbdSalesOrPurchaseDetailsPostingAccount, bool)`

GetPostingAccountOk returns a tuple with the PostingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAccount

`func (o *QbdServiceItemSalesOrPurchaseDetails) SetPostingAccount(v QbdSalesOrPurchaseDetailsPostingAccount)`

SetPostingAccount sets PostingAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


