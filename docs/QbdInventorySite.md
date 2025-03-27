# QbdInventorySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this inventory site. This ID is unique across all inventory sites but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_inventory_site\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this inventory site was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this inventory site was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this inventory site object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this inventory site, unique across all inventory sites.  **NOTE**: Inventory sites do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this inventory site is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Parent** | [**QbdInventorySiteParent**](QbdInventorySiteParent.md) |  | 
**IsDefault** | **bool** | Indicates whether this inventory site is the default site used when no specific site is provided during the creation of other objects. | 
**Description** | **string** | A description of this inventory site. | 
**Contact** | **string** | The name of the primary contact person for this inventory site. | 
**Phone** | **string** | The inventory site&#39;s primary telephone number. | 
**Fax** | **string** | The inventory site&#39;s fax number. | 
**Email** | **string** | The inventory site&#39;s email address. | 
**Address** | [**QbdInventorySiteAddress**](QbdInventorySiteAddress.md) |  | 

## Methods

### NewQbdInventorySite

`func NewQbdInventorySite(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, parent QbdInventorySiteParent, isDefault bool, description string, contact string, phone string, fax string, email string, address QbdInventorySiteAddress, ) *QbdInventorySite`

NewQbdInventorySite instantiates a new QbdInventorySite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventorySiteWithDefaults

`func NewQbdInventorySiteWithDefaults() *QbdInventorySite`

NewQbdInventorySiteWithDefaults instantiates a new QbdInventorySite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInventorySite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInventorySite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInventorySite) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInventorySite) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInventorySite) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInventorySite) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdInventorySite) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdInventorySite) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdInventorySite) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdInventorySite) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdInventorySite) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdInventorySite) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdInventorySite) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdInventorySite) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdInventorySite) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdInventorySite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdInventorySite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdInventorySite) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdInventorySite) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdInventorySite) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdInventorySite) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetParent

`func (o *QbdInventorySite) GetParent() QbdInventorySiteParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdInventorySite) GetParentOk() (*QbdInventorySiteParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdInventorySite) SetParent(v QbdInventorySiteParent)`

SetParent sets Parent field to given value.


### GetIsDefault

`func (o *QbdInventorySite) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *QbdInventorySite) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *QbdInventorySite) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetDescription

`func (o *QbdInventorySite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdInventorySite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdInventorySite) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetContact

`func (o *QbdInventorySite) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QbdInventorySite) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QbdInventorySite) SetContact(v string)`

SetContact sets Contact field to given value.


### GetPhone

`func (o *QbdInventorySite) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QbdInventorySite) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QbdInventorySite) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetFax

`func (o *QbdInventorySite) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QbdInventorySite) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QbdInventorySite) SetFax(v string)`

SetFax sets Fax field to given value.


### GetEmail

`func (o *QbdInventorySite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QbdInventorySite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QbdInventorySite) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddress

`func (o *QbdInventorySite) GetAddress() QbdInventorySiteAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdInventorySite) GetAddressOk() (*QbdInventorySiteAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdInventorySite) SetAddress(v QbdInventorySiteAddress)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


