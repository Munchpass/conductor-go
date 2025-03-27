# QbdCreditCardChargeCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this object. This ID is unique across all objects of the same type, but not across different QuickBooks object types. | 
**FullName** | **string** | The fully-qualified unique name for this object, formed by combining the names of its parent objects with its own &#x60;name&#x60;, separated by colons. Not case-sensitive. | 

## Methods

### NewQbdCreditCardChargeCurrency

`func NewQbdCreditCardChargeCurrency(id string, fullName string, ) *QbdCreditCardChargeCurrency`

NewQbdCreditCardChargeCurrency instantiates a new QbdCreditCardChargeCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditCardChargeCurrencyWithDefaults

`func NewQbdCreditCardChargeCurrencyWithDefaults() *QbdCreditCardChargeCurrency`

NewQbdCreditCardChargeCurrencyWithDefaults instantiates a new QbdCreditCardChargeCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCreditCardChargeCurrency) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCreditCardChargeCurrency) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCreditCardChargeCurrency) SetId(v string)`

SetId sets Id field to given value.


### GetFullName

`func (o *QbdCreditCardChargeCurrency) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdCreditCardChargeCurrency) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdCreditCardChargeCurrency) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


