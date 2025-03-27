# QbdNonInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this non-inventory item. This ID is unique across all non-inventory items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_non_inventory_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this non-inventory item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this non-inventory item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this non-inventory item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this non-inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two non-inventory items could both have the &#x60;name&#x60; \&quot;Printer Ink Cartridge\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Office Supplies:Printer Ink Cartridge\&quot; and \&quot;Miscellaneous:Printer Ink Cartridge\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this non-inventory item, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if a non-inventory item is under \&quot;Office Supplies\&quot; and has the &#x60;name&#x60; \&quot;Printer Ink Cartridge\&quot;, its &#x60;fullName&#x60; would be \&quot;Office Supplies:Printer Ink Cartridge\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all non-inventory item objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**Barcode** | **string** | The non-inventory item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this non-inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdNonInventoryItemClass**](QbdNonInventoryItemClass.md) |  | 
**Parent** | [**QbdNonInventoryItemParent**](QbdNonInventoryItemParent.md) |  | 
**Sublevel** | **float32** | The depth level of this non-inventory item in the hierarchy. A top-level non-inventory item has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, a non-inventory item with a &#x60;fullName&#x60; of \&quot;Office Supplies:Printer Ink Cartridge\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**Sku** | **string** | The non-inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | 
**UnitOfMeasureSet** | [**QbdNonInventoryItemUnitOfMeasureSet**](QbdNonInventoryItemUnitOfMeasureSet.md) |  | 
**SalesTaxCode** | [**QbdNonInventoryItemSalesTaxCode**](QbdNonInventoryItemSalesTaxCode.md) |  | 
**SalesOrPurchaseDetails** | [**QbdNonInventoryItemSalesOrPurchaseDetails**](QbdNonInventoryItemSalesOrPurchaseDetails.md) |  | 
**SalesAndPurchaseDetails** | [**QbdNonInventoryItemSalesAndPurchaseDetails**](QbdNonInventoryItemSalesAndPurchaseDetails.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the non-inventory item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdNonInventoryItem

`func NewQbdNonInventoryItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, barcode string, isActive bool, class QbdNonInventoryItemClass, parent QbdNonInventoryItemParent, sublevel float32, sku string, unitOfMeasureSet QbdNonInventoryItemUnitOfMeasureSet, salesTaxCode QbdNonInventoryItemSalesTaxCode, salesOrPurchaseDetails QbdNonInventoryItemSalesOrPurchaseDetails, salesAndPurchaseDetails QbdNonInventoryItemSalesAndPurchaseDetails, externalId string, customFields []QbdCustomField, ) *QbdNonInventoryItem`

NewQbdNonInventoryItem instantiates a new QbdNonInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdNonInventoryItemWithDefaults

`func NewQbdNonInventoryItemWithDefaults() *QbdNonInventoryItem`

NewQbdNonInventoryItemWithDefaults instantiates a new QbdNonInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdNonInventoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdNonInventoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdNonInventoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdNonInventoryItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdNonInventoryItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdNonInventoryItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdNonInventoryItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdNonInventoryItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdNonInventoryItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdNonInventoryItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdNonInventoryItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdNonInventoryItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdNonInventoryItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdNonInventoryItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdNonInventoryItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdNonInventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdNonInventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdNonInventoryItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdNonInventoryItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdNonInventoryItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdNonInventoryItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBarcode

`func (o *QbdNonInventoryItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdNonInventoryItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdNonInventoryItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdNonInventoryItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdNonInventoryItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdNonInventoryItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdNonInventoryItem) GetClass() QbdNonInventoryItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdNonInventoryItem) GetClassOk() (*QbdNonInventoryItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdNonInventoryItem) SetClass(v QbdNonInventoryItemClass)`

SetClass sets Class field to given value.


### GetParent

`func (o *QbdNonInventoryItem) GetParent() QbdNonInventoryItemParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdNonInventoryItem) GetParentOk() (*QbdNonInventoryItemParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdNonInventoryItem) SetParent(v QbdNonInventoryItemParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdNonInventoryItem) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdNonInventoryItem) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdNonInventoryItem) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetSku

`func (o *QbdNonInventoryItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QbdNonInventoryItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QbdNonInventoryItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetUnitOfMeasureSet

`func (o *QbdNonInventoryItem) GetUnitOfMeasureSet() QbdNonInventoryItemUnitOfMeasureSet`

GetUnitOfMeasureSet returns the UnitOfMeasureSet field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetOk

`func (o *QbdNonInventoryItem) GetUnitOfMeasureSetOk() (*QbdNonInventoryItemUnitOfMeasureSet, bool)`

GetUnitOfMeasureSetOk returns a tuple with the UnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSet

`func (o *QbdNonInventoryItem) SetUnitOfMeasureSet(v QbdNonInventoryItemUnitOfMeasureSet)`

SetUnitOfMeasureSet sets UnitOfMeasureSet field to given value.


### GetSalesTaxCode

`func (o *QbdNonInventoryItem) GetSalesTaxCode() QbdNonInventoryItemSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdNonInventoryItem) GetSalesTaxCodeOk() (*QbdNonInventoryItemSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdNonInventoryItem) SetSalesTaxCode(v QbdNonInventoryItemSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetSalesOrPurchaseDetails

`func (o *QbdNonInventoryItem) GetSalesOrPurchaseDetails() QbdNonInventoryItemSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QbdNonInventoryItem) GetSalesOrPurchaseDetailsOk() (*QbdNonInventoryItemSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QbdNonInventoryItem) SetSalesOrPurchaseDetails(v QbdNonInventoryItemSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.


### GetSalesAndPurchaseDetails

`func (o *QbdNonInventoryItem) GetSalesAndPurchaseDetails() QbdNonInventoryItemSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QbdNonInventoryItem) GetSalesAndPurchaseDetailsOk() (*QbdNonInventoryItemSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QbdNonInventoryItem) SetSalesAndPurchaseDetails(v QbdNonInventoryItemSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.


### GetExternalId

`func (o *QbdNonInventoryItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdNonInventoryItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdNonInventoryItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdNonInventoryItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdNonInventoryItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdNonInventoryItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


