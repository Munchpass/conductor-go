# QbdSalesRepresentative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales representative. This ID is unique across all sales representatives but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_representative\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this sales representative was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this sales representative was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this sales representative object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Initial** | **string** | The initials of this sales representative&#39;s name. | 
**IsActive** | **bool** | Indicates whether this sales representative is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Entity** | [**QbdSalesRepresentativeEntity**](QbdSalesRepresentativeEntity.md) |  | 

## Methods

### NewQbdSalesRepresentative

`func NewQbdSalesRepresentative(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, initial string, isActive bool, entity QbdSalesRepresentativeEntity, ) *QbdSalesRepresentative`

NewQbdSalesRepresentative instantiates a new QbdSalesRepresentative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesRepresentativeWithDefaults

`func NewQbdSalesRepresentativeWithDefaults() *QbdSalesRepresentative`

NewQbdSalesRepresentativeWithDefaults instantiates a new QbdSalesRepresentative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesRepresentative) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesRepresentative) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesRepresentative) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesRepresentative) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesRepresentative) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesRepresentative) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSalesRepresentative) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSalesRepresentative) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSalesRepresentative) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSalesRepresentative) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSalesRepresentative) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSalesRepresentative) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSalesRepresentative) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSalesRepresentative) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSalesRepresentative) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetInitial

`func (o *QbdSalesRepresentative) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *QbdSalesRepresentative) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *QbdSalesRepresentative) SetInitial(v string)`

SetInitial sets Initial field to given value.


### GetIsActive

`func (o *QbdSalesRepresentative) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdSalesRepresentative) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdSalesRepresentative) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetEntity

`func (o *QbdSalesRepresentative) GetEntity() QbdSalesRepresentativeEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *QbdSalesRepresentative) GetEntityOk() (*QbdSalesRepresentativeEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *QbdSalesRepresentative) SetEntity(v QbdSalesRepresentativeEntity)`

SetEntity sets Entity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


