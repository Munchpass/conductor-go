# QuickbooksDesktopCustomersPostRequestAdditionalContactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Salutation** | Pointer to **string** | The contact&#39;s formal salutation title that precedes their name, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | **string** | The contact&#39;s first name.  Maximum length: 25 characters. | 
**MiddleName** | Pointer to **string** | The contact&#39;s middle name.  Maximum length: 5 characters. | [optional] 
**LastName** | Pointer to **string** | The contact&#39;s last name.  Maximum length: 25 characters. | [optional] 
**JobTitle** | Pointer to **string** | The contact&#39;s job title. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this contact, such as phone numbers or email addresses. | [optional] 

## Methods

### NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInner

`func NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInner(firstName string, ) *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner`

NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInner instantiates a new QuickbooksDesktopCustomersPostRequestAdditionalContactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInnerWithDefaults

`func NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInnerWithDefaults() *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner`

NewQuickbooksDesktopCustomersPostRequestAdditionalContactsInnerWithDefaults instantiates a new QuickbooksDesktopCustomersPostRequestAdditionalContactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalutation

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequestAdditionalContactsInner) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


