# QbdCreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **NullableString** | The credit card number. Must be masked with lower case \&quot;x\&quot; and no dashes. | 
**ExpirationMonth** | **NullableFloat32** | The month when the credit card expires. | 
**ExpirationYear** | **NullableFloat32** | The year when the credit card expires. | 
**Name** | **NullableString** | The cardholder&#39;s name on the card. | 
**Address** | **NullableString** | The card&#39;s billing address. | 
**PostalCode** | **NullableString** | The card&#39;s billing address ZIP or postal code. | 

## Methods

### NewQbdCreditCard

`func NewQbdCreditCard(number NullableString, expirationMonth NullableFloat32, expirationYear NullableFloat32, name NullableString, address NullableString, postalCode NullableString, ) *QbdCreditCard`

NewQbdCreditCard instantiates a new QbdCreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditCardWithDefaults

`func NewQbdCreditCardWithDefaults() *QbdCreditCard`

NewQbdCreditCardWithDefaults instantiates a new QbdCreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *QbdCreditCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *QbdCreditCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *QbdCreditCard) SetNumber(v string)`

SetNumber sets Number field to given value.


### SetNumberNil

`func (o *QbdCreditCard) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *QbdCreditCard) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetExpirationMonth

`func (o *QbdCreditCard) GetExpirationMonth() float32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *QbdCreditCard) GetExpirationMonthOk() (*float32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *QbdCreditCard) SetExpirationMonth(v float32)`

SetExpirationMonth sets ExpirationMonth field to given value.


### SetExpirationMonthNil

`func (o *QbdCreditCard) SetExpirationMonthNil(b bool)`

 SetExpirationMonthNil sets the value for ExpirationMonth to be an explicit nil

### UnsetExpirationMonth
`func (o *QbdCreditCard) UnsetExpirationMonth()`

UnsetExpirationMonth ensures that no value is present for ExpirationMonth, not even an explicit nil
### GetExpirationYear

`func (o *QbdCreditCard) GetExpirationYear() float32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *QbdCreditCard) GetExpirationYearOk() (*float32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *QbdCreditCard) SetExpirationYear(v float32)`

SetExpirationYear sets ExpirationYear field to given value.


### SetExpirationYearNil

`func (o *QbdCreditCard) SetExpirationYearNil(b bool)`

 SetExpirationYearNil sets the value for ExpirationYear to be an explicit nil

### UnsetExpirationYear
`func (o *QbdCreditCard) UnsetExpirationYear()`

UnsetExpirationYear ensures that no value is present for ExpirationYear, not even an explicit nil
### GetName

`func (o *QbdCreditCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdCreditCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdCreditCard) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *QbdCreditCard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QbdCreditCard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *QbdCreditCard) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdCreditCard) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdCreditCard) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *QbdCreditCard) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *QbdCreditCard) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPostalCode

`func (o *QbdCreditCard) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdCreditCard) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdCreditCard) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *QbdCreditCard) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *QbdCreditCard) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


