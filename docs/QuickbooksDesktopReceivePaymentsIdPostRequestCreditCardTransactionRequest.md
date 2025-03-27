# QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The credit card number. Must be masked with lower case \&quot;x\&quot; and no dashes. | [optional] 
**ExpirationMonth** | Pointer to **float32** | The month when the credit card expires. | [optional] 
**ExpirationYear** | Pointer to **float32** | The year when the credit card expires. | [optional] 
**Name** | Pointer to **string** | The cardholder&#39;s name on the card. | [optional] 
**Address** | Pointer to **string** | The card&#39;s billing address. | [optional] 
**PostalCode** | Pointer to **string** | The card&#39;s billing address ZIP or postal code. | [optional] 
**CommercialCardCode** | Pointer to **string** | The commercial card code identifies the type of business credit card being used (purchase, corporate, or business) for Visa and Mastercard transactions only. When provided, this code may qualify the transaction for lower processing fees compared to the standard rates that apply when no code is specified. | [optional] 
**TransactionMode** | Pointer to **string** | Indicates whether this credit card transaction came from a card swipe (&#x60;card_present&#x60;) or not (&#x60;card_not_present&#x60;). | [optional] 
**TransactionType** | Pointer to **string** | The QBMS transaction type from which the current transaction data originated. | [optional] 

## Methods

### NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest

`func NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest() *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest`

NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequestWithDefaults

`func NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequestWithDefaults() *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest`

NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequestWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetExpirationMonth() float32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetExpirationMonthOk() (*float32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetExpirationMonth(v float32)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationYear

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetExpirationYear() float32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetExpirationYearOk() (*float32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetExpirationYear(v float32)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetName

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCommercialCardCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetCommercialCardCode() string`

GetCommercialCardCode returns the CommercialCardCode field if non-nil, zero value otherwise.

### GetCommercialCardCodeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetCommercialCardCodeOk() (*string, bool)`

GetCommercialCardCodeOk returns a tuple with the CommercialCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialCardCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetCommercialCardCode(v string)`

SetCommercialCardCode sets CommercialCardCode field to given value.

### HasCommercialCardCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasCommercialCardCode() bool`

HasCommercialCardCode returns a boolean if a field has been set.

### GetTransactionMode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetTransactionMode() string`

GetTransactionMode returns the TransactionMode field if non-nil, zero value otherwise.

### GetTransactionModeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetTransactionModeOk() (*string, bool)`

GetTransactionModeOk returns a tuple with the TransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetTransactionMode(v string)`

SetTransactionMode sets TransactionMode field to given value.

### HasTransactionMode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasTransactionMode() bool`

HasTransactionMode returns a boolean if a field has been set.

### GetTransactionType

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionRequest) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


