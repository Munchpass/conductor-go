# QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of the contact to update. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the contact object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Salutation** | Pointer to **string** | The contact&#39;s formal salutation title that precedes their name, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The contact&#39;s first name.  Maximum length: 25 characters. | [optional] 
**MiddleName** | Pointer to **string** | The contact&#39;s middle name.  Maximum length: 5 characters. | [optional] 
**LastName** | Pointer to **string** | The contact&#39;s last name.  Maximum length: 25 characters. | [optional] 
**JobTitle** | Pointer to **string** | The contact&#39;s job title. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this contact, such as phone numbers or email addresses. | [optional] 

## Methods

### NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner

`func NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner(id string, revisionNumber string, ) *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner`

NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner instantiates a new QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInnerWithDefaults

`func NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInnerWithDefaults() *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner`

NewQuickbooksDesktopCustomersIdPostRequestAdditionalContactsInnerWithDefaults instantiates a new QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetId(v string)`

SetId sets Id field to given value.


### GetRevisionNumber

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetSalutation

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


