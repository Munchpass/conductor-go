# QbdSalesOrderBillingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **string** | The first line of the address (e.g., street, PO Box, or company name). | 
**Line2** | **string** | The second line of the address, if needed (e.g., apartment, suite, unit, or building). | 
**Line3** | **string** | The third line of the address, if needed. | 
**Line4** | **string** | The fourth line of the address, if needed. | 
**Line5** | **string** | The fifth line of the address, if needed. | 
**City** | **string** | The city, district, suburb, town, or village name of the address. | 
**State** | **string** | The state, county, province, or region name of the address. | 
**PostalCode** | **string** | The postal code or ZIP code of the address. | 
**Country** | **string** | The country name of the address. | 
**Note** | **string** | A note written at the bottom of the address in the form in which it appears, such as the invoice form. | 

## Methods

### NewQbdSalesOrderBillingAddress

`func NewQbdSalesOrderBillingAddress(line1 string, line2 string, line3 string, line4 string, line5 string, city string, state string, postalCode string, country string, note string, ) *QbdSalesOrderBillingAddress`

NewQbdSalesOrderBillingAddress instantiates a new QbdSalesOrderBillingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesOrderBillingAddressWithDefaults

`func NewQbdSalesOrderBillingAddressWithDefaults() *QbdSalesOrderBillingAddress`

NewQbdSalesOrderBillingAddressWithDefaults instantiates a new QbdSalesOrderBillingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *QbdSalesOrderBillingAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QbdSalesOrderBillingAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QbdSalesOrderBillingAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *QbdSalesOrderBillingAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QbdSalesOrderBillingAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QbdSalesOrderBillingAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### GetLine3

`func (o *QbdSalesOrderBillingAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QbdSalesOrderBillingAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QbdSalesOrderBillingAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### GetLine4

`func (o *QbdSalesOrderBillingAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QbdSalesOrderBillingAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QbdSalesOrderBillingAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### GetLine5

`func (o *QbdSalesOrderBillingAddress) GetLine5() string`

GetLine5 returns the Line5 field if non-nil, zero value otherwise.

### GetLine5Ok

`func (o *QbdSalesOrderBillingAddress) GetLine5Ok() (*string, bool)`

GetLine5Ok returns a tuple with the Line5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine5

`func (o *QbdSalesOrderBillingAddress) SetLine5(v string)`

SetLine5 sets Line5 field to given value.


### GetCity

`func (o *QbdSalesOrderBillingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QbdSalesOrderBillingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QbdSalesOrderBillingAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *QbdSalesOrderBillingAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QbdSalesOrderBillingAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QbdSalesOrderBillingAddress) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *QbdSalesOrderBillingAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QbdSalesOrderBillingAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QbdSalesOrderBillingAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *QbdSalesOrderBillingAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QbdSalesOrderBillingAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QbdSalesOrderBillingAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetNote

`func (o *QbdSalesOrderBillingAddress) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdSalesOrderBillingAddress) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdSalesOrderBillingAddress) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


