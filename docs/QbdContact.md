# QbdContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this contact. This ID is unique across all contacts but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_contact\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this contact was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this contact was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this contact object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **NullableString** | The contact&#39;s full name. | 
**Salutation** | **NullableString** | The contact&#39;s formal salutation title that precedes their name, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | 
**FirstName** | **string** | The contact&#39;s first name. | 
**MiddleName** | **NullableString** | The contact&#39;s middle name. | 
**LastName** | **NullableString** | The contact&#39;s last name. | 
**JobTitle** | **NullableString** | The contact&#39;s job title. | 
**CustomContactFields** | [**[]QbdCustomContactField**](QbdCustomContactField.md) | Additional custom contact fields for this contact, such as phone numbers or email addresses. | 

## Methods

### NewQbdContact

`func NewQbdContact(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name NullableString, salutation NullableString, firstName string, middleName NullableString, lastName NullableString, jobTitle NullableString, customContactFields []QbdCustomContactField, ) *QbdContact`

NewQbdContact instantiates a new QbdContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdContactWithDefaults

`func NewQbdContactWithDefaults() *QbdContact`

NewQbdContactWithDefaults instantiates a new QbdContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdContact) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdContact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdContact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdContact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdContact) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdContact) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdContact) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdContact) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdContact) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdContact) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdContact) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdContact) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdContact) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdContact) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *QbdContact) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QbdContact) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSalutation

`func (o *QbdContact) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QbdContact) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QbdContact) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.


### SetSalutationNil

`func (o *QbdContact) SetSalutationNil(b bool)`

 SetSalutationNil sets the value for Salutation to be an explicit nil

### UnsetSalutation
`func (o *QbdContact) UnsetSalutation()`

UnsetSalutation ensures that no value is present for Salutation, not even an explicit nil
### GetFirstName

`func (o *QbdContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QbdContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QbdContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *QbdContact) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QbdContact) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QbdContact) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *QbdContact) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *QbdContact) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *QbdContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QbdContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QbdContact) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *QbdContact) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *QbdContact) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetJobTitle

`func (o *QbdContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QbdContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QbdContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### SetJobTitleNil

`func (o *QbdContact) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *QbdContact) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCustomContactFields

`func (o *QbdContact) GetCustomContactFields() []QbdCustomContactField`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QbdContact) GetCustomContactFieldsOk() (*[]QbdCustomContactField, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QbdContact) SetCustomContactFields(v []QbdCustomContactField)`

SetCustomContactFields sets CustomContactFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


