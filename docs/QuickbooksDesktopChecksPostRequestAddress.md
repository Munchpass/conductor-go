# QuickbooksDesktopChecksPostRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **string** | The first line of the address (e.g., street, PO Box, or company name).  Maximum length: 41 characters. | [optional] 
**Line2** | Pointer to **string** | The second line of the address, if needed (e.g., apartment, suite, unit, or building).  Maximum length: 41 characters. | [optional] 
**Line3** | Pointer to **string** | The third line of the address, if needed.  Maximum length: 41 characters. | [optional] 
**Line4** | Pointer to **string** | The fourth line of the address, if needed.  Maximum length: 41 characters. | [optional] 
**Line5** | Pointer to **string** | The fifth line of the address, if needed.  Maximum length: 41 characters. | [optional] 
**City** | Pointer to **string** | The city, district, suburb, town, or village name of the address.  Maximum length: 31 characters. | [optional] 
**State** | Pointer to **string** | The state, county, province, or region name of the address.  Maximum length: 21 characters. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the address.  Maximum length: 13 characters. | [optional] 
**Country** | Pointer to **string** | The country name of the address. | [optional] 
**Note** | Pointer to **string** | A note written at the bottom of the address in the form in which it appears, such as the invoice form. | [optional] 

## Methods

### NewQuickbooksDesktopChecksPostRequestAddress

`func NewQuickbooksDesktopChecksPostRequestAddress() *QuickbooksDesktopChecksPostRequestAddress`

NewQuickbooksDesktopChecksPostRequestAddress instantiates a new QuickbooksDesktopChecksPostRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopChecksPostRequestAddressWithDefaults

`func NewQuickbooksDesktopChecksPostRequestAddressWithDefaults() *QuickbooksDesktopChecksPostRequestAddress`

NewQuickbooksDesktopChecksPostRequestAddressWithDefaults instantiates a new QuickbooksDesktopChecksPostRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### GetLine4

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.

### HasLine4

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasLine4() bool`

HasLine4 returns a boolean if a field has been set.

### GetLine5

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine5() string`

GetLine5 returns the Line5 field if non-nil, zero value otherwise.

### GetLine5Ok

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetLine5Ok() (*string, bool)`

GetLine5Ok returns a tuple with the Line5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine5

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetLine5(v string)`

SetLine5 sets Line5 field to given value.

### HasLine5

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasLine5() bool`

HasLine5 returns a boolean if a field has been set.

### GetCity

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopChecksPostRequestAddress) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopChecksPostRequestAddress) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopChecksPostRequestAddress) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


