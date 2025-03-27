# QbdStandardTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this standard term. This ID is unique across all standard terms but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_standard_term\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this standard term was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this standard term was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this standard term object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this standard term, unique across all standard terms.  **NOTE**: Standard terms do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this standard term is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**DueDays** | **NullableFloat32** | The number of days until payment is due. | 
**DiscountDays** | **NullableFloat32** | The number of days within which payment must be received to qualify for the discount specified by &#x60;discountPercentage&#x60;. | 
**DiscountPercentage** | **NullableString** | The discount percentage applied to the payment if received within the number of days specified by &#x60;discountDays&#x60;. The value is between 0 and 100. | 

## Methods

### NewQbdStandardTerm

`func NewQbdStandardTerm(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, dueDays NullableFloat32, discountDays NullableFloat32, discountPercentage NullableString, ) *QbdStandardTerm`

NewQbdStandardTerm instantiates a new QbdStandardTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdStandardTermWithDefaults

`func NewQbdStandardTermWithDefaults() *QbdStandardTerm`

NewQbdStandardTermWithDefaults instantiates a new QbdStandardTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdStandardTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdStandardTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdStandardTerm) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdStandardTerm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdStandardTerm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdStandardTerm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdStandardTerm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdStandardTerm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdStandardTerm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdStandardTerm) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdStandardTerm) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdStandardTerm) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdStandardTerm) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdStandardTerm) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdStandardTerm) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdStandardTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdStandardTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdStandardTerm) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdStandardTerm) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdStandardTerm) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdStandardTerm) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetDueDays

`func (o *QbdStandardTerm) GetDueDays() float32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *QbdStandardTerm) GetDueDaysOk() (*float32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *QbdStandardTerm) SetDueDays(v float32)`

SetDueDays sets DueDays field to given value.


### SetDueDaysNil

`func (o *QbdStandardTerm) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *QbdStandardTerm) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetDiscountDays

`func (o *QbdStandardTerm) GetDiscountDays() float32`

GetDiscountDays returns the DiscountDays field if non-nil, zero value otherwise.

### GetDiscountDaysOk

`func (o *QbdStandardTerm) GetDiscountDaysOk() (*float32, bool)`

GetDiscountDaysOk returns a tuple with the DiscountDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDays

`func (o *QbdStandardTerm) SetDiscountDays(v float32)`

SetDiscountDays sets DiscountDays field to given value.


### SetDiscountDaysNil

`func (o *QbdStandardTerm) SetDiscountDaysNil(b bool)`

 SetDiscountDaysNil sets the value for DiscountDays to be an explicit nil

### UnsetDiscountDays
`func (o *QbdStandardTerm) UnsetDiscountDays()`

UnsetDiscountDays ensures that no value is present for DiscountDays, not even an explicit nil
### GetDiscountPercentage

`func (o *QbdStandardTerm) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *QbdStandardTerm) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *QbdStandardTerm) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.


### SetDiscountPercentageNil

`func (o *QbdStandardTerm) SetDiscountPercentageNil(b bool)`

 SetDiscountPercentageNil sets the value for DiscountPercentage to be an explicit nil

### UnsetDiscountPercentage
`func (o *QbdStandardTerm) UnsetDiscountPercentage()`

UnsetDiscountPercentage ensures that no value is present for DiscountPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


