# QbdSalesTaxCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales-tax code. This ID is unique across all sales-tax codes but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_tax_code\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this sales-tax code was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this sales-tax code was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this sales-tax code object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this sales-tax code, unique across all sales-tax codes. This short name will appear on sales forms to identify the tax status of an item.  **NOTE**: Sales-tax codes do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this sales-tax code is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**IsTaxable** | **bool** | Indicates whether this sales-tax code is tracking taxable sales. This field cannot be modified once the sales-tax code has been used in a transaction. | 
**Description** | **string** | A description of this sales-tax code. | 
**SalesTaxItem** | [**QbdSalesTaxCodeSalesTaxItem**](QbdSalesTaxCodeSalesTaxItem.md) |  | 

## Methods

### NewQbdSalesTaxCode

`func NewQbdSalesTaxCode(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, isTaxable bool, description string, salesTaxItem QbdSalesTaxCodeSalesTaxItem, ) *QbdSalesTaxCode`

NewQbdSalesTaxCode instantiates a new QbdSalesTaxCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesTaxCodeWithDefaults

`func NewQbdSalesTaxCodeWithDefaults() *QbdSalesTaxCode`

NewQbdSalesTaxCodeWithDefaults instantiates a new QbdSalesTaxCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesTaxCode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesTaxCode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesTaxCode) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesTaxCode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesTaxCode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesTaxCode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSalesTaxCode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSalesTaxCode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSalesTaxCode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSalesTaxCode) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSalesTaxCode) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSalesTaxCode) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSalesTaxCode) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSalesTaxCode) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSalesTaxCode) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdSalesTaxCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdSalesTaxCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdSalesTaxCode) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdSalesTaxCode) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdSalesTaxCode) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdSalesTaxCode) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsTaxable

`func (o *QbdSalesTaxCode) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *QbdSalesTaxCode) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *QbdSalesTaxCode) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.


### GetDescription

`func (o *QbdSalesTaxCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdSalesTaxCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdSalesTaxCode) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSalesTaxItem

`func (o *QbdSalesTaxCode) GetSalesTaxItem() QbdSalesTaxCodeSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdSalesTaxCode) GetSalesTaxItemOk() (*QbdSalesTaxCodeSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdSalesTaxCode) SetSalesTaxItem(v QbdSalesTaxCodeSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


