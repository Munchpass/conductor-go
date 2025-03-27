# QbdCreditCardTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The credit card number. Must be masked with lower case \&quot;x\&quot; and no dashes. | 
**ExpirationMonth** | **float32** | The month when the credit card expires. | 
**ExpirationYear** | **float32** | The year when the credit card expires. | 
**Name** | **string** | The cardholder&#39;s name on the card. | 
**Address** | **NullableString** | The card&#39;s billing address. | 
**PostalCode** | **NullableString** | The card&#39;s billing address ZIP or postal code. | 
**CommercialCardCode** | **NullableString** | The commercial card code identifies the type of business credit card being used (purchase, corporate, or business) for Visa and Mastercard transactions only. When provided, this code may qualify the transaction for lower processing fees compared to the standard rates that apply when no code is specified. | 
**TransactionMode** | **NullableString** | Indicates whether this credit card transaction came from a card swipe (&#x60;card_present&#x60;) or not (&#x60;card_not_present&#x60;). | 
**TransactionType** | **NullableString** | The QBMS transaction type from which the current transaction data originated. | 

## Methods

### NewQbdCreditCardTransactionRequest

`func NewQbdCreditCardTransactionRequest(number string, expirationMonth float32, expirationYear float32, name string, address NullableString, postalCode NullableString, commercialCardCode NullableString, transactionMode NullableString, transactionType NullableString, ) *QbdCreditCardTransactionRequest`

NewQbdCreditCardTransactionRequest instantiates a new QbdCreditCardTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditCardTransactionRequestWithDefaults

`func NewQbdCreditCardTransactionRequestWithDefaults() *QbdCreditCardTransactionRequest`

NewQbdCreditCardTransactionRequestWithDefaults instantiates a new QbdCreditCardTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *QbdCreditCardTransactionRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *QbdCreditCardTransactionRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *QbdCreditCardTransactionRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetExpirationMonth

`func (o *QbdCreditCardTransactionRequest) GetExpirationMonth() float32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *QbdCreditCardTransactionRequest) GetExpirationMonthOk() (*float32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *QbdCreditCardTransactionRequest) SetExpirationMonth(v float32)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationYear

`func (o *QbdCreditCardTransactionRequest) GetExpirationYear() float32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *QbdCreditCardTransactionRequest) GetExpirationYearOk() (*float32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *QbdCreditCardTransactionRequest) SetExpirationYear(v float32)`

SetExpirationYear sets ExpirationYear field to given value.


### GetName

`func (o *QbdCreditCardTransactionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdCreditCardTransactionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdCreditCardTransactionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *QbdCreditCardTransactionRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdCreditCardTransactionRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdCreditCardTransactionRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *QbdCreditCardTransactionRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *QbdCreditCardTransactionRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPostalCode

`func (o *QbdCreditCardTransactionRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdCreditCardTransactionRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdCreditCardTransactionRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *QbdCreditCardTransactionRequest) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *QbdCreditCardTransactionRequest) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCommercialCardCode

`func (o *QbdCreditCardTransactionRequest) GetCommercialCardCode() string`

GetCommercialCardCode returns the CommercialCardCode field if non-nil, zero value otherwise.

### GetCommercialCardCodeOk

`func (o *QbdCreditCardTransactionRequest) GetCommercialCardCodeOk() (*string, bool)`

GetCommercialCardCodeOk returns a tuple with the CommercialCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialCardCode

`func (o *QbdCreditCardTransactionRequest) SetCommercialCardCode(v string)`

SetCommercialCardCode sets CommercialCardCode field to given value.


### SetCommercialCardCodeNil

`func (o *QbdCreditCardTransactionRequest) SetCommercialCardCodeNil(b bool)`

 SetCommercialCardCodeNil sets the value for CommercialCardCode to be an explicit nil

### UnsetCommercialCardCode
`func (o *QbdCreditCardTransactionRequest) UnsetCommercialCardCode()`

UnsetCommercialCardCode ensures that no value is present for CommercialCardCode, not even an explicit nil
### GetTransactionMode

`func (o *QbdCreditCardTransactionRequest) GetTransactionMode() string`

GetTransactionMode returns the TransactionMode field if non-nil, zero value otherwise.

### GetTransactionModeOk

`func (o *QbdCreditCardTransactionRequest) GetTransactionModeOk() (*string, bool)`

GetTransactionModeOk returns a tuple with the TransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMode

`func (o *QbdCreditCardTransactionRequest) SetTransactionMode(v string)`

SetTransactionMode sets TransactionMode field to given value.


### SetTransactionModeNil

`func (o *QbdCreditCardTransactionRequest) SetTransactionModeNil(b bool)`

 SetTransactionModeNil sets the value for TransactionMode to be an explicit nil

### UnsetTransactionMode
`func (o *QbdCreditCardTransactionRequest) UnsetTransactionMode()`

UnsetTransactionMode ensures that no value is present for TransactionMode, not even an explicit nil
### GetTransactionType

`func (o *QbdCreditCardTransactionRequest) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *QbdCreditCardTransactionRequest) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *QbdCreditCardTransactionRequest) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### SetTransactionTypeNil

`func (o *QbdCreditCardTransactionRequest) SetTransactionTypeNil(b bool)`

 SetTransactionTypeNil sets the value for TransactionType to be an explicit nil

### UnsetTransactionType
`func (o *QbdCreditCardTransactionRequest) UnsetTransactionType()`

UnsetTransactionType ensures that no value is present for TransactionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


