# QbdEmployeeAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **NullableString** | The first line of the employee address (e.g., street, PO Box, or company name). | 
**Line2** | **NullableString** | The second line of the employee address, if needed (e.g., apartment, suite, unit, or building). | 
**Line3** | **NullableString** | The third line of the employee address, if needed. | 
**Line4** | **NullableString** | The fourth line of the employee address, if needed. | 
**City** | **NullableString** | The city, district, suburb, town, or village name of the employee address. | 
**State** | **NullableString** | The U.S. state of the employee address. QuickBooks requires this field to be a U.S. state abbreviation (e.g., \&quot;CA\&quot; for California). See enum for all possible values. | 
**PostalCode** | **NullableString** | The postal code or ZIP code of the employee address. | 
**Country** | **NullableString** | The country name of the employee address. | 

## Methods

### NewQbdEmployeeAddress

`func NewQbdEmployeeAddress(line1 NullableString, line2 NullableString, line3 NullableString, line4 NullableString, city NullableString, state NullableString, postalCode NullableString, country NullableString, ) *QbdEmployeeAddress`

NewQbdEmployeeAddress instantiates a new QbdEmployeeAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeeAddressWithDefaults

`func NewQbdEmployeeAddressWithDefaults() *QbdEmployeeAddress`

NewQbdEmployeeAddressWithDefaults instantiates a new QbdEmployeeAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *QbdEmployeeAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QbdEmployeeAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QbdEmployeeAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *QbdEmployeeAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *QbdEmployeeAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *QbdEmployeeAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QbdEmployeeAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QbdEmployeeAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### SetLine2Nil

`func (o *QbdEmployeeAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *QbdEmployeeAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *QbdEmployeeAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QbdEmployeeAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QbdEmployeeAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### SetLine3Nil

`func (o *QbdEmployeeAddress) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *QbdEmployeeAddress) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *QbdEmployeeAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QbdEmployeeAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QbdEmployeeAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### SetLine4Nil

`func (o *QbdEmployeeAddress) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *QbdEmployeeAddress) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetCity

`func (o *QbdEmployeeAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QbdEmployeeAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QbdEmployeeAddress) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *QbdEmployeeAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *QbdEmployeeAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *QbdEmployeeAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QbdEmployeeAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QbdEmployeeAddress) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *QbdEmployeeAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *QbdEmployeeAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *QbdEmployeeAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdEmployeeAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdEmployeeAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *QbdEmployeeAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *QbdEmployeeAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *QbdEmployeeAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QbdEmployeeAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QbdEmployeeAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *QbdEmployeeAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *QbdEmployeeAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


