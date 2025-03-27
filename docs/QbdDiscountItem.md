# QbdDiscountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this discount item. This ID is unique across all discount items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_discount_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this discount item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this discount item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this discount item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this discount item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two discount items could both have the &#x60;name&#x60; \&quot;10% labor discount\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Discounts:10% labor discount\&quot; and \&quot;Promotions:10% labor discount\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this discount item, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if a discount item is under \&quot;Discounts\&quot; and has the &#x60;name&#x60; \&quot;10% labor discount\&quot;, its &#x60;fullName&#x60; would be \&quot;Discounts:10% labor discount\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all discount item objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**Barcode** | **string** | The discount item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this discount item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdDiscountItemClass**](QbdDiscountItemClass.md) |  | 
**Parent** | [**QbdDiscountItemParent**](QbdDiscountItemParent.md) |  | 
**Sublevel** | **float32** | The depth level of this discount item in the hierarchy. A top-level discount item has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, a discount item with a &#x60;fullName&#x60; of \&quot;Discounts:10% labor discount\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**Description** | **string** | The discount item&#39;s description that will appear on sales forms that include this item. | 
**SalesTaxCode** | [**QbdDiscountItemSalesTaxCode**](QbdDiscountItemSalesTaxCode.md) |  | 
**DiscountRate** | **string** | The monetary amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A flat rate discount applies to ALL lines recorded above it and distributes the discount amount equally across those lines, which affects tax calculations. For example, a $10 discount applied to a $100 taxable item and $100 non-taxable item would result in a $5 taxable discount and $5 non-taxable discount. | 
**DiscountRatePercent** | **string** | The percentage amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A percentage discount only applies to the line immediately above it, so tax implications only affect that specific line. | 
**Account** | [**QbdDiscountItemAccount**](QbdDiscountItemAccount.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the discount item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdDiscountItem

`func NewQbdDiscountItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, barcode string, isActive bool, class QbdDiscountItemClass, parent QbdDiscountItemParent, sublevel float32, description string, salesTaxCode QbdDiscountItemSalesTaxCode, discountRate string, discountRatePercent string, account QbdDiscountItemAccount, externalId string, customFields []QbdCustomField, ) *QbdDiscountItem`

NewQbdDiscountItem instantiates a new QbdDiscountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdDiscountItemWithDefaults

`func NewQbdDiscountItemWithDefaults() *QbdDiscountItem`

NewQbdDiscountItemWithDefaults instantiates a new QbdDiscountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdDiscountItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdDiscountItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdDiscountItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdDiscountItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdDiscountItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdDiscountItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdDiscountItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdDiscountItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdDiscountItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdDiscountItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdDiscountItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdDiscountItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdDiscountItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdDiscountItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdDiscountItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdDiscountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdDiscountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdDiscountItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdDiscountItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdDiscountItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdDiscountItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBarcode

`func (o *QbdDiscountItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdDiscountItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdDiscountItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdDiscountItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdDiscountItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdDiscountItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdDiscountItem) GetClass() QbdDiscountItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdDiscountItem) GetClassOk() (*QbdDiscountItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdDiscountItem) SetClass(v QbdDiscountItemClass)`

SetClass sets Class field to given value.


### GetParent

`func (o *QbdDiscountItem) GetParent() QbdDiscountItemParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdDiscountItem) GetParentOk() (*QbdDiscountItemParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdDiscountItem) SetParent(v QbdDiscountItemParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdDiscountItem) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdDiscountItem) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdDiscountItem) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetDescription

`func (o *QbdDiscountItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdDiscountItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdDiscountItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSalesTaxCode

`func (o *QbdDiscountItem) GetSalesTaxCode() QbdDiscountItemSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdDiscountItem) GetSalesTaxCodeOk() (*QbdDiscountItemSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdDiscountItem) SetSalesTaxCode(v QbdDiscountItemSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetDiscountRate

`func (o *QbdDiscountItem) GetDiscountRate() string`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *QbdDiscountItem) GetDiscountRateOk() (*string, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *QbdDiscountItem) SetDiscountRate(v string)`

SetDiscountRate sets DiscountRate field to given value.


### GetDiscountRatePercent

`func (o *QbdDiscountItem) GetDiscountRatePercent() string`

GetDiscountRatePercent returns the DiscountRatePercent field if non-nil, zero value otherwise.

### GetDiscountRatePercentOk

`func (o *QbdDiscountItem) GetDiscountRatePercentOk() (*string, bool)`

GetDiscountRatePercentOk returns a tuple with the DiscountRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRatePercent

`func (o *QbdDiscountItem) SetDiscountRatePercent(v string)`

SetDiscountRatePercent sets DiscountRatePercent field to given value.


### GetAccount

`func (o *QbdDiscountItem) GetAccount() QbdDiscountItemAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdDiscountItem) GetAccountOk() (*QbdDiscountItemAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdDiscountItem) SetAccount(v QbdDiscountItemAccount)`

SetAccount sets Account field to given value.


### GetExternalId

`func (o *QbdDiscountItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdDiscountItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdDiscountItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdDiscountItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdDiscountItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdDiscountItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


