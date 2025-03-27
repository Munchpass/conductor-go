# QbdServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this service item. This ID is unique across all service items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_service_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this service item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this service item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this service item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this service item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two service items could both have the &#x60;name&#x60; \&quot;Web-Design\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Consulting:Web-Design\&quot; and \&quot;Contracting:Web-Design\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this service item, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if a service item is under \&quot;Consulting\&quot; and has the &#x60;name&#x60; \&quot;Web-Design\&quot;, its &#x60;fullName&#x60; would be \&quot;Consulting:Web-Design\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all service item objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**Barcode** | **string** | The service item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this service item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdServiceItemClass**](QbdServiceItemClass.md) |  | 
**Parent** | [**QbdServiceItemParent**](QbdServiceItemParent.md) |  | 
**Sublevel** | **float32** | The depth level of this service item in the hierarchy. A top-level service item has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, a service item with a &#x60;fullName&#x60; of \&quot;Consulting:Web-Design\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**UnitOfMeasureSet** | [**QbdServiceItemUnitOfMeasureSet**](QbdServiceItemUnitOfMeasureSet.md) |  | 
**SalesTaxCode** | [**QbdServiceItemSalesTaxCode**](QbdServiceItemSalesTaxCode.md) |  | 
**SalesOrPurchaseDetails** | [**QbdServiceItemSalesOrPurchaseDetails**](QbdServiceItemSalesOrPurchaseDetails.md) |  | 
**SalesAndPurchaseDetails** | [**QbdServiceItemSalesAndPurchaseDetails**](QbdServiceItemSalesAndPurchaseDetails.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the service item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdServiceItem

`func NewQbdServiceItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, barcode string, isActive bool, class QbdServiceItemClass, parent QbdServiceItemParent, sublevel float32, unitOfMeasureSet QbdServiceItemUnitOfMeasureSet, salesTaxCode QbdServiceItemSalesTaxCode, salesOrPurchaseDetails QbdServiceItemSalesOrPurchaseDetails, salesAndPurchaseDetails QbdServiceItemSalesAndPurchaseDetails, externalId string, customFields []QbdCustomField, ) *QbdServiceItem`

NewQbdServiceItem instantiates a new QbdServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdServiceItemWithDefaults

`func NewQbdServiceItemWithDefaults() *QbdServiceItem`

NewQbdServiceItemWithDefaults instantiates a new QbdServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdServiceItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdServiceItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdServiceItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdServiceItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdServiceItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdServiceItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdServiceItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdServiceItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdServiceItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdServiceItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdServiceItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdServiceItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdServiceItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdServiceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdServiceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdServiceItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdServiceItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdServiceItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdServiceItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBarcode

`func (o *QbdServiceItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdServiceItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdServiceItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdServiceItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdServiceItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdServiceItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdServiceItem) GetClass() QbdServiceItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdServiceItem) GetClassOk() (*QbdServiceItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdServiceItem) SetClass(v QbdServiceItemClass)`

SetClass sets Class field to given value.


### GetParent

`func (o *QbdServiceItem) GetParent() QbdServiceItemParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdServiceItem) GetParentOk() (*QbdServiceItemParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdServiceItem) SetParent(v QbdServiceItemParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdServiceItem) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdServiceItem) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdServiceItem) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetUnitOfMeasureSet

`func (o *QbdServiceItem) GetUnitOfMeasureSet() QbdServiceItemUnitOfMeasureSet`

GetUnitOfMeasureSet returns the UnitOfMeasureSet field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetOk

`func (o *QbdServiceItem) GetUnitOfMeasureSetOk() (*QbdServiceItemUnitOfMeasureSet, bool)`

GetUnitOfMeasureSetOk returns a tuple with the UnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSet

`func (o *QbdServiceItem) SetUnitOfMeasureSet(v QbdServiceItemUnitOfMeasureSet)`

SetUnitOfMeasureSet sets UnitOfMeasureSet field to given value.


### GetSalesTaxCode

`func (o *QbdServiceItem) GetSalesTaxCode() QbdServiceItemSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdServiceItem) GetSalesTaxCodeOk() (*QbdServiceItemSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdServiceItem) SetSalesTaxCode(v QbdServiceItemSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetSalesOrPurchaseDetails

`func (o *QbdServiceItem) GetSalesOrPurchaseDetails() QbdServiceItemSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QbdServiceItem) GetSalesOrPurchaseDetailsOk() (*QbdServiceItemSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QbdServiceItem) SetSalesOrPurchaseDetails(v QbdServiceItemSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.


### GetSalesAndPurchaseDetails

`func (o *QbdServiceItem) GetSalesAndPurchaseDetails() QbdServiceItemSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QbdServiceItem) GetSalesAndPurchaseDetailsOk() (*QbdServiceItemSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QbdServiceItem) SetSalesAndPurchaseDetails(v QbdServiceItemSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.


### GetExternalId

`func (o *QbdServiceItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdServiceItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdServiceItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdServiceItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdServiceItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdServiceItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


