# QuickbooksDesktopEmployeesPostRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **string** | The first line of the employee address (e.g., street, PO Box, or company name).  Maximum length: 41 characters. | [optional] 
**Line2** | Pointer to **string** | The second line of the employee address, if needed (e.g., apartment, suite, unit, or building).  Maximum length: 41 characters. | [optional] 
**Line3** | Pointer to **string** | The third line of the employee address, if needed.  Maximum length: 41 characters. | [optional] 
**Line4** | Pointer to **string** | The fourth line of the employee address, if needed.  Maximum length: 41 characters. | [optional] 
**City** | Pointer to **string** | The city, district, suburb, town, or village name of the employee address.  Maximum length: 31 characters. | [optional] 
**State** | Pointer to **string** | The U.S. state of the employee address. QuickBooks requires this field to be a U.S. state abbreviation (e.g., \&quot;CA\&quot; for California). See enum for all possible values. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the employee address.  Maximum length: 13 characters. | [optional] 
**Country** | Pointer to **string** | The country name of the employee address. | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequestAddress

`func NewQuickbooksDesktopEmployeesPostRequestAddress() *QuickbooksDesktopEmployeesPostRequestAddress`

NewQuickbooksDesktopEmployeesPostRequestAddress instantiates a new QuickbooksDesktopEmployeesPostRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestAddressWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestAddressWithDefaults() *QuickbooksDesktopEmployeesPostRequestAddress`

NewQuickbooksDesktopEmployeesPostRequestAddressWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### GetLine4

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.

### HasLine4

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasLine4() bool`

HasLine4 returns a boolean if a field has been set.

### GetCity

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QuickbooksDesktopEmployeesPostRequestAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


