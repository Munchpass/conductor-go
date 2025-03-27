# QbdCustomerCreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The credit card number. Must be masked with lower case \&quot;x\&quot; and no dashes. | 
**ExpirationMonth** | **float32** | The month when the credit card expires. | 
**ExpirationYear** | **float32** | The year when the credit card expires. | 
**Name** | **string** | The cardholder&#39;s name on the card. | 
**Address** | **string** | The card&#39;s billing address. | 
**PostalCode** | **string** | The card&#39;s billing address ZIP or postal code. | 

## Methods

### NewQbdCustomerCreditCard

`func NewQbdCustomerCreditCard(number string, expirationMonth float32, expirationYear float32, name string, address string, postalCode string, ) *QbdCustomerCreditCard`

NewQbdCustomerCreditCard instantiates a new QbdCustomerCreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCustomerCreditCardWithDefaults

`func NewQbdCustomerCreditCardWithDefaults() *QbdCustomerCreditCard`

NewQbdCustomerCreditCardWithDefaults instantiates a new QbdCustomerCreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *QbdCustomerCreditCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *QbdCustomerCreditCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *QbdCustomerCreditCard) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetExpirationMonth

`func (o *QbdCustomerCreditCard) GetExpirationMonth() float32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *QbdCustomerCreditCard) GetExpirationMonthOk() (*float32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *QbdCustomerCreditCard) SetExpirationMonth(v float32)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationYear

`func (o *QbdCustomerCreditCard) GetExpirationYear() float32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *QbdCustomerCreditCard) GetExpirationYearOk() (*float32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *QbdCustomerCreditCard) SetExpirationYear(v float32)`

SetExpirationYear sets ExpirationYear field to given value.


### GetName

`func (o *QbdCustomerCreditCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdCustomerCreditCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdCustomerCreditCard) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *QbdCustomerCreditCard) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdCustomerCreditCard) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdCustomerCreditCard) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPostalCode

`func (o *QbdCustomerCreditCard) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdCustomerCreditCard) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdCustomerCreditCard) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


