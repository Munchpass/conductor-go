# QbdClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this class. This ID is unique across all classes but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_class\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this class was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this class was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this class object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this class. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two classes could both have the &#x60;name&#x60; \&quot;Marketing\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Department:Marketing\&quot; and \&quot;Internal:Marketing\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this class, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if a class is under \&quot;Department\&quot; and has the &#x60;name&#x60; \&quot;Marketing\&quot;, its &#x60;fullName&#x60; would be \&quot;Department:Marketing\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all class objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**IsActive** | **bool** | Indicates whether this class is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Parent** | [**QbdClassParent**](QbdClassParent.md) |  | 
**Sublevel** | **float32** | The depth level of this class in the hierarchy. A top-level class has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, a class with a &#x60;fullName&#x60; of \&quot;Department:Marketing\&quot; would have a &#x60;sublevel&#x60; of 1. | 

## Methods

### NewQbdClass

`func NewQbdClass(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, isActive bool, parent QbdClassParent, sublevel float32, ) *QbdClass`

NewQbdClass instantiates a new QbdClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdClassWithDefaults

`func NewQbdClassWithDefaults() *QbdClass`

NewQbdClassWithDefaults instantiates a new QbdClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdClass) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdClass) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdClass) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdClass) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdClass) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdClass) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdClass) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdClass) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdClass) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdClass) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdClass) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdClass) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdClass) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdClass) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdClass) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdClass) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdClass) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsActive

`func (o *QbdClass) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdClass) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdClass) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetParent

`func (o *QbdClass) GetParent() QbdClassParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdClass) GetParentOk() (*QbdClassParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdClass) SetParent(v QbdClassParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdClass) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdClass) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdClass) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


