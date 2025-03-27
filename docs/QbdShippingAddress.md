# QbdShippingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this shipping address, unique across all shipping addresses.  **NOTE**: Shipping addresses do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**Line1** | **NullableString** | The first line of the shipping address (e.g., street, PO Box, or company name). | 
**Line2** | **NullableString** | The second line of the shipping address, if needed (e.g., apartment, suite, unit, or building). | 
**Line3** | **NullableString** | The third line of the shipping address, if needed. | 
**Line4** | **NullableString** | The fourth line of the shipping address, if needed. | 
**Line5** | **NullableString** | The fifth line of the shipping address, if needed. | 
**City** | **NullableString** | The city, district, suburb, town, or village name of the shipping address. | 
**State** | **NullableString** | The state, county, province, or region name of the shipping address. | 
**PostalCode** | **NullableString** | The postal code or ZIP code of the shipping address. | 
**Country** | **NullableString** | The country name of the shipping address. | 
**Note** | **NullableString** | A note written at the bottom of the shipping address in the form in which it appears, such as the invoice form. | 
**IsDefaultShippingAddress** | **NullableBool** | Indicates whether this shipping address is the default shipping address. | 

## Methods

### NewQbdShippingAddress

`func NewQbdShippingAddress(name string, line1 NullableString, line2 NullableString, line3 NullableString, line4 NullableString, line5 NullableString, city NullableString, state NullableString, postalCode NullableString, country NullableString, note NullableString, isDefaultShippingAddress NullableBool, ) *QbdShippingAddress`

NewQbdShippingAddress instantiates a new QbdShippingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdShippingAddressWithDefaults

`func NewQbdShippingAddressWithDefaults() *QbdShippingAddress`

NewQbdShippingAddressWithDefaults instantiates a new QbdShippingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QbdShippingAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdShippingAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdShippingAddress) SetName(v string)`

SetName sets Name field to given value.


### GetLine1

`func (o *QbdShippingAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QbdShippingAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QbdShippingAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *QbdShippingAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *QbdShippingAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *QbdShippingAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QbdShippingAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QbdShippingAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### SetLine2Nil

`func (o *QbdShippingAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *QbdShippingAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *QbdShippingAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QbdShippingAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QbdShippingAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### SetLine3Nil

`func (o *QbdShippingAddress) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *QbdShippingAddress) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *QbdShippingAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QbdShippingAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QbdShippingAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### SetLine4Nil

`func (o *QbdShippingAddress) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *QbdShippingAddress) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetLine5

`func (o *QbdShippingAddress) GetLine5() string`

GetLine5 returns the Line5 field if non-nil, zero value otherwise.

### GetLine5Ok

`func (o *QbdShippingAddress) GetLine5Ok() (*string, bool)`

GetLine5Ok returns a tuple with the Line5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine5

`func (o *QbdShippingAddress) SetLine5(v string)`

SetLine5 sets Line5 field to given value.


### SetLine5Nil

`func (o *QbdShippingAddress) SetLine5Nil(b bool)`

 SetLine5Nil sets the value for Line5 to be an explicit nil

### UnsetLine5
`func (o *QbdShippingAddress) UnsetLine5()`

UnsetLine5 ensures that no value is present for Line5, not even an explicit nil
### GetCity

`func (o *QbdShippingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QbdShippingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QbdShippingAddress) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *QbdShippingAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *QbdShippingAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *QbdShippingAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QbdShippingAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QbdShippingAddress) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *QbdShippingAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *QbdShippingAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *QbdShippingAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdShippingAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdShippingAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *QbdShippingAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *QbdShippingAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *QbdShippingAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QbdShippingAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QbdShippingAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *QbdShippingAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *QbdShippingAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetNote

`func (o *QbdShippingAddress) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdShippingAddress) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdShippingAddress) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *QbdShippingAddress) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *QbdShippingAddress) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetIsDefaultShippingAddress

`func (o *QbdShippingAddress) GetIsDefaultShippingAddress() bool`

GetIsDefaultShippingAddress returns the IsDefaultShippingAddress field if non-nil, zero value otherwise.

### GetIsDefaultShippingAddressOk

`func (o *QbdShippingAddress) GetIsDefaultShippingAddressOk() (*bool, bool)`

GetIsDefaultShippingAddressOk returns a tuple with the IsDefaultShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultShippingAddress

`func (o *QbdShippingAddress) SetIsDefaultShippingAddress(v bool)`

SetIsDefaultShippingAddress sets IsDefaultShippingAddress field to given value.


### SetIsDefaultShippingAddressNil

`func (o *QbdShippingAddress) SetIsDefaultShippingAddressNil(b bool)`

 SetIsDefaultShippingAddressNil sets the value for IsDefaultShippingAddress to be an explicit nil

### UnsetIsDefaultShippingAddress
`func (o *QbdShippingAddress) UnsetIsDefaultShippingAddress()`

UnsetIsDefaultShippingAddress ensures that no value is present for IsDefaultShippingAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


