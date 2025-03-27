# QbdSalesTaxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales-tax item. This ID is unique across all sales-tax items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_tax_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this sales-tax item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this sales-tax item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this sales-tax item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this sales-tax item, unique across all sales-tax items.  **NOTE**: Sales-tax items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**Barcode** | **string** | The sales-tax item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this sales-tax item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdSalesTaxItemClass**](QbdSalesTaxItemClass.md) |  | 
**Description** | **string** | The sales-tax item&#39;s description that will appear on sales forms that include this item. | 
**TaxRate** | **string** | The tax rate defined by this sales-tax item, represented as a decimal string. For example, \&quot;7.5\&quot; represents a 7.5% tax rate. This rate determines the amount of sales tax applied when this item is used in transactions. If a non-zero &#x60;taxRate&#x60; is specified, then the &#x60;taxVendor&#x60; field is required. | 
**TaxVendor** | [**QbdSalesTaxItemTaxVendor**](QbdSalesTaxItemTaxVendor.md) |  | 
**SalesTaxReturnLine** | [**QbdSalesTaxItemSalesTaxReturnLine**](QbdSalesTaxItemSalesTaxReturnLine.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales-tax item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesTaxItem

`func NewQbdSalesTaxItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, barcode string, isActive bool, class QbdSalesTaxItemClass, description string, taxRate string, taxVendor QbdSalesTaxItemTaxVendor, salesTaxReturnLine QbdSalesTaxItemSalesTaxReturnLine, externalId string, customFields []QbdCustomField, ) *QbdSalesTaxItem`

NewQbdSalesTaxItem instantiates a new QbdSalesTaxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesTaxItemWithDefaults

`func NewQbdSalesTaxItemWithDefaults() *QbdSalesTaxItem`

NewQbdSalesTaxItemWithDefaults instantiates a new QbdSalesTaxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesTaxItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesTaxItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesTaxItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesTaxItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesTaxItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesTaxItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSalesTaxItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSalesTaxItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSalesTaxItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSalesTaxItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSalesTaxItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSalesTaxItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSalesTaxItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSalesTaxItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSalesTaxItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdSalesTaxItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdSalesTaxItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdSalesTaxItem) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QbdSalesTaxItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdSalesTaxItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdSalesTaxItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdSalesTaxItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdSalesTaxItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdSalesTaxItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdSalesTaxItem) GetClass() QbdSalesTaxItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdSalesTaxItem) GetClassOk() (*QbdSalesTaxItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdSalesTaxItem) SetClass(v QbdSalesTaxItemClass)`

SetClass sets Class field to given value.


### GetDescription

`func (o *QbdSalesTaxItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesTaxItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesTaxItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTaxRate

`func (o *QbdSalesTaxItem) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *QbdSalesTaxItem) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *QbdSalesTaxItem) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.


### GetTaxVendor

`func (o *QbdSalesTaxItem) GetTaxVendor() QbdSalesTaxItemTaxVendor`

GetTaxVendor returns the TaxVendor field if non-nil, zero value otherwise.

### GetTaxVendorOk

`func (o *QbdSalesTaxItem) GetTaxVendorOk() (*QbdSalesTaxItemTaxVendor, bool)`

GetTaxVendorOk returns a tuple with the TaxVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxVendor

`func (o *QbdSalesTaxItem) SetTaxVendor(v QbdSalesTaxItemTaxVendor)`

SetTaxVendor sets TaxVendor field to given value.


### GetSalesTaxReturnLine

`func (o *QbdSalesTaxItem) GetSalesTaxReturnLine() QbdSalesTaxItemSalesTaxReturnLine`

GetSalesTaxReturnLine returns the SalesTaxReturnLine field if non-nil, zero value otherwise.

### GetSalesTaxReturnLineOk

`func (o *QbdSalesTaxItem) GetSalesTaxReturnLineOk() (*QbdSalesTaxItemSalesTaxReturnLine, bool)`

GetSalesTaxReturnLineOk returns a tuple with the SalesTaxReturnLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturnLine

`func (o *QbdSalesTaxItem) SetSalesTaxReturnLine(v QbdSalesTaxItemSalesTaxReturnLine)`

SetSalesTaxReturnLine sets SalesTaxReturnLine field to given value.


### GetExternalId

`func (o *QbdSalesTaxItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdSalesTaxItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdSalesTaxItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdSalesTaxItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesTaxItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesTaxItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


