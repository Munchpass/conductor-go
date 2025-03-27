# QbdEmployeeAddress1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **string** | The first line of the employee address (e.g., street, PO Box, or company name). | 
**Line2** | **string** | The second line of the employee address, if needed (e.g., apartment, suite, unit, or building). | 
**Line3** | **string** | The third line of the employee address, if needed. | 
**Line4** | **string** | The fourth line of the employee address, if needed. | 
**City** | **string** | The city, district, suburb, town, or village name of the employee address. | 
**State** | **string** | The U.S. state of the employee address. QuickBooks requires this field to be a U.S. state abbreviation (e.g., \&quot;CA\&quot; for California). See enum for all possible values. | 
**PostalCode** | **string** | The postal code or ZIP code of the employee address. | 
**Country** | **string** | The country name of the employee address. | 

## Methods

### NewQbdEmployeeAddress1

`func NewQbdEmployeeAddress1(line1 string, line2 string, line3 string, line4 string, city string, state string, postalCode string, country string, ) *QbdEmployeeAddress1`

NewQbdEmployeeAddress1 instantiates a new QbdEmployeeAddress1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeeAddress1WithDefaults

`func NewQbdEmployeeAddress1WithDefaults() *QbdEmployeeAddress1`

NewQbdEmployeeAddress1WithDefaults instantiates a new QbdEmployeeAddress1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *QbdEmployeeAddress1) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QbdEmployeeAddress1) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QbdEmployeeAddress1) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *QbdEmployeeAddress1) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QbdEmployeeAddress1) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QbdEmployeeAddress1) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### GetLine3

`func (o *QbdEmployeeAddress1) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QbdEmployeeAddress1) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QbdEmployeeAddress1) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### GetLine4

`func (o *QbdEmployeeAddress1) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QbdEmployeeAddress1) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QbdEmployeeAddress1) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### GetCity

`func (o *QbdEmployeeAddress1) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QbdEmployeeAddress1) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QbdEmployeeAddress1) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *QbdEmployeeAddress1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QbdEmployeeAddress1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QbdEmployeeAddress1) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *QbdEmployeeAddress1) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdEmployeeAddress1) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdEmployeeAddress1) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *QbdEmployeeAddress1) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QbdEmployeeAddress1) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QbdEmployeeAddress1) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


