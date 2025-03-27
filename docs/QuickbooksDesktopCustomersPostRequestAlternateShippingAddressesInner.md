# QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this shipping address, unique across all shipping addresses.  **NOTE**: Shipping addresses do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 41 characters. | 
**Line1** | Pointer to **string** | The first line of the shipping address (e.g., street, PO Box, or company name).  Maximum length: 41 characters. | [optional] 
**Line2** | Pointer to **string** | The second line of the shipping address, if needed (e.g., apartment, suite, unit, or building).  Maximum length: 41 characters. | [optional] 
**Line3** | Pointer to **string** | The third line of the shipping address, if needed.  Maximum length: 41 characters. | [optional] 
**Line4** | Pointer to **string** | The fourth line of the shipping address, if needed.  Maximum length: 41 characters. | [optional] 
**Line5** | Pointer to **string** | The fifth line of the shipping address, if needed.  Maximum length: 41 characters. | [optional] 
**City** | Pointer to **string** | The city, district, suburb, town, or village name of the shipping address.  Maximum length: 31 characters. | [optional] 
**State** | Pointer to **string** | The state, county, province, or region name of the shipping address.  Maximum length: 21 characters. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the shipping address.  Maximum length: 13 characters. | [optional] 
**Country** | Pointer to **string** | The country name of the shipping address. | [optional] 
**Note** | Pointer to **string** | A note written at the bottom of the shipping address in the form in which it appears, such as the invoice form. | [optional] 
**IsDefaultShippingAddress** | Pointer to **bool** | Indicates whether this shipping address is the default shipping address. | [optional] 

## Methods

### NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner

`func NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner(name string, ) *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner`

NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner instantiates a new QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInnerWithDefaults

`func NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInnerWithDefaults() *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner`

NewQuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInnerWithDefaults instantiates a new QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetName(v string)`

SetName sets Name field to given value.


### GetLine1

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### GetLine4

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetLine4(v string)`

SetLine4 sets Line4 field to given value.

### HasLine4

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasLine4() bool`

HasLine4 returns a boolean if a field has been set.

### GetLine5

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine5() string`

GetLine5 returns the Line5 field if non-nil, zero value otherwise.

### GetLine5Ok

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetLine5Ok() (*string, bool)`

GetLine5Ok returns a tuple with the Line5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine5

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetLine5(v string)`

SetLine5 sets Line5 field to given value.

### HasLine5

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasLine5() bool`

HasLine5 returns a boolean if a field has been set.

### GetCity

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetIsDefaultShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetIsDefaultShippingAddress() bool`

GetIsDefaultShippingAddress returns the IsDefaultShippingAddress field if non-nil, zero value otherwise.

### GetIsDefaultShippingAddressOk

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) GetIsDefaultShippingAddressOk() (*bool, bool)`

GetIsDefaultShippingAddressOk returns a tuple with the IsDefaultShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) SetIsDefaultShippingAddress(v bool)`

SetIsDefaultShippingAddress sets IsDefaultShippingAddress field to given value.

### HasIsDefaultShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner) HasIsDefaultShippingAddress() bool`

HasIsDefaultShippingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


