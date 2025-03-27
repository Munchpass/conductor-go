# QbdSubtotalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this subtotal item. This ID is unique across all subtotal items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_subtotal_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this subtotal item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this subtotal item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this subtotal item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this subtotal item, unique across all subtotal items.  **NOTE**: Subtotal items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**Barcode** | **NullableString** | The subtotal item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this subtotal item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Description** | **NullableString** | The subtotal item&#39;s description that will appear on sales forms that include this item. | 
**SpecialItemType** | **NullableString** | The type of special item for this subtotal item. | 
**ExternalId** | **NullableString** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the subtotal item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSubtotalItem

`func NewQbdSubtotalItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, barcode NullableString, isActive bool, description NullableString, specialItemType NullableString, externalId NullableString, customFields []QbdCustomField, ) *QbdSubtotalItem`

NewQbdSubtotalItem instantiates a new QbdSubtotalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSubtotalItemWithDefaults

`func NewQbdSubtotalItemWithDefaults() *QbdSubtotalItem`

NewQbdSubtotalItemWithDefaults instantiates a new QbdSubtotalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSubtotalItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSubtotalItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSubtotalItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSubtotalItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSubtotalItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSubtotalItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSubtotalItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSubtotalItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSubtotalItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSubtotalItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSubtotalItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSubtotalItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSubtotalItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSubtotalItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSubtotalItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdSubtotalItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdSubtotalItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdSubtotalItem) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QbdSubtotalItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdSubtotalItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdSubtotalItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### SetBarcodeNil

`func (o *QbdSubtotalItem) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *QbdSubtotalItem) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetIsActive

`func (o *QbdSubtotalItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdSubtotalItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdSubtotalItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetDescription

`func (o *QbdSubtotalItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSubtotalItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSubtotalItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *QbdSubtotalItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QbdSubtotalItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSpecialItemType

`func (o *QbdSubtotalItem) GetSpecialItemType() string`

GetSpecialItemType returns the SpecialItemType field if non-nil, zero value otherwise.

### GetSpecialItemTypeOk

`func (o *QbdSubtotalItem) GetSpecialItemTypeOk() (*string, bool)`

GetSpecialItemTypeOk returns a tuple with the SpecialItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialItemType

`func (o *QbdSubtotalItem) SetSpecialItemType(v string)`

SetSpecialItemType sets SpecialItemType field to given value.


### SetSpecialItemTypeNil

`func (o *QbdSubtotalItem) SetSpecialItemTypeNil(b bool)`

 SetSpecialItemTypeNil sets the value for SpecialItemType to be an explicit nil

### UnsetSpecialItemType
`func (o *QbdSubtotalItem) UnsetSpecialItemType()`

UnsetSpecialItemType ensures that no value is present for SpecialItemType, not even an explicit nil
### GetExternalId

`func (o *QbdSubtotalItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdSubtotalItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdSubtotalItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *QbdSubtotalItem) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *QbdSubtotalItem) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCustomFields

`func (o *QbdSubtotalItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSubtotalItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSubtotalItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


