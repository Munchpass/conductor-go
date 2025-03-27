# QbdSalesOrPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of this item. | 
**Price** | **string** | The price at which this item is purchased or sold, represented as a decimal string. | 
**PricePercentage** | **string** | The price of this item expressed as a percentage, used instead of &#x60;price&#x60; when the item&#39;s cost is calculated as a percentage of another amount. For example, a service item that costs a percentage of another item&#39;s price. | 
**PostingAccount** | [**QbdSalesOrPurchaseDetailsPostingAccount**](QbdSalesOrPurchaseDetailsPostingAccount.md) |  | 

## Methods

### NewQbdSalesOrPurchaseDetails

`func NewQbdSalesOrPurchaseDetails(description string, price string, pricePercentage string, postingAccount QbdSalesOrPurchaseDetailsPostingAccount, ) *QbdSalesOrPurchaseDetails`

NewQbdSalesOrPurchaseDetails instantiates a new QbdSalesOrPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesOrPurchaseDetailsWithDefaults

`func NewQbdSalesOrPurchaseDetailsWithDefaults() *QbdSalesOrPurchaseDetails`

NewQbdSalesOrPurchaseDetailsWithDefaults instantiates a new QbdSalesOrPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QbdSalesOrPurchaseDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesOrPurchaseDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesOrPurchaseDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *QbdSalesOrPurchaseDetails) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QbdSalesOrPurchaseDetails) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QbdSalesOrPurchaseDetails) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetPricePercentage

`func (o *QbdSalesOrPurchaseDetails) GetPricePercentage() string`

GetPricePercentage returns the PricePercentage field if non-nil, zero value otherwise.

### GetPricePercentageOk

`func (o *QbdSalesOrPurchaseDetails) GetPricePercentageOk() (*string, bool)`

GetPricePercentageOk returns a tuple with the PricePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePercentage

`func (o *QbdSalesOrPurchaseDetails) SetPricePercentage(v string)`

SetPricePercentage sets PricePercentage field to given value.


### GetPostingAccount

`func (o *QbdSalesOrPurchaseDetails) GetPostingAccount() QbdSalesOrPurchaseDetailsPostingAccount`

GetPostingAccount returns the PostingAccount field if non-nil, zero value otherwise.

### GetPostingAccountOk

`func (o *QbdSalesOrPurchaseDetails) GetPostingAccountOk() (*QbdSalesOrPurchaseDetailsPostingAccount, bool)`

GetPostingAccountOk returns a tuple with the PostingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAccount

`func (o *QbdSalesOrPurchaseDetails) SetPostingAccount(v QbdSalesOrPurchaseDetailsPostingAccount)`

SetPostingAccount sets PostingAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


