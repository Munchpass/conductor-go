# QbdInventoryAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this inventory adjustment. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_inventory_adjustment\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this inventory adjustment was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this inventory adjustment was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this inventory adjustment object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Account** | [**QbdInventoryAdjustmentAccount**](QbdInventoryAdjustmentAccount.md) |  | 
**InventorySite** | [**QbdInventoryAdjustmentInventorySite**](QbdInventoryAdjustmentInventorySite.md) |  | 
**TransactionDate** | **string** | The date of this inventory adjustment, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this inventory adjustment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Customer** | [**QbdInventoryAdjustmentCustomer**](QbdInventoryAdjustmentCustomer.md) |  | 
**Class** | [**QbdInventoryAdjustmentClass**](QbdInventoryAdjustmentClass.md) |  | 
**Memo** | **string** | A memo or note for this inventory adjustment. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**Lines** | [**[]QbdInventoryAdjustmentLine**](QbdInventoryAdjustmentLine.md) | The inventory adjustment&#39;s item lines, each representing the adjustment of an inventory item&#39;s quantity, value, serial number, or lot number. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the inventory adjustment object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdInventoryAdjustment

`func NewQbdInventoryAdjustment(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, account QbdInventoryAdjustmentAccount, inventorySite QbdInventoryAdjustmentInventorySite, transactionDate string, refNumber string, customer QbdInventoryAdjustmentCustomer, class QbdInventoryAdjustmentClass, memo string, externalId string, lines []QbdInventoryAdjustmentLine, customFields []QbdCustomField, ) *QbdInventoryAdjustment`

NewQbdInventoryAdjustment instantiates a new QbdInventoryAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventoryAdjustmentWithDefaults

`func NewQbdInventoryAdjustmentWithDefaults() *QbdInventoryAdjustment`

NewQbdInventoryAdjustmentWithDefaults instantiates a new QbdInventoryAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInventoryAdjustment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInventoryAdjustment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInventoryAdjustment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInventoryAdjustment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInventoryAdjustment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInventoryAdjustment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdInventoryAdjustment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdInventoryAdjustment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdInventoryAdjustment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdInventoryAdjustment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdInventoryAdjustment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdInventoryAdjustment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdInventoryAdjustment) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdInventoryAdjustment) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdInventoryAdjustment) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetAccount

`func (o *QbdInventoryAdjustment) GetAccount() QbdInventoryAdjustmentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdInventoryAdjustment) GetAccountOk() (*QbdInventoryAdjustmentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdInventoryAdjustment) SetAccount(v QbdInventoryAdjustmentAccount)`

SetAccount sets Account field to given value.


### GetInventorySite

`func (o *QbdInventoryAdjustment) GetInventorySite() QbdInventoryAdjustmentInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdInventoryAdjustment) GetInventorySiteOk() (*QbdInventoryAdjustmentInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdInventoryAdjustment) SetInventorySite(v QbdInventoryAdjustmentInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetTransactionDate

`func (o *QbdInventoryAdjustment) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdInventoryAdjustment) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdInventoryAdjustment) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdInventoryAdjustment) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdInventoryAdjustment) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdInventoryAdjustment) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetCustomer

`func (o *QbdInventoryAdjustment) GetCustomer() QbdInventoryAdjustmentCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdInventoryAdjustment) GetCustomerOk() (*QbdInventoryAdjustmentCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdInventoryAdjustment) SetCustomer(v QbdInventoryAdjustmentCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdInventoryAdjustment) GetClass() QbdInventoryAdjustmentClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdInventoryAdjustment) GetClassOk() (*QbdInventoryAdjustmentClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdInventoryAdjustment) SetClass(v QbdInventoryAdjustmentClass)`

SetClass sets Class field to given value.


### GetMemo

`func (o *QbdInventoryAdjustment) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdInventoryAdjustment) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdInventoryAdjustment) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetExternalId

`func (o *QbdInventoryAdjustment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdInventoryAdjustment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdInventoryAdjustment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLines

`func (o *QbdInventoryAdjustment) GetLines() []QbdInventoryAdjustmentLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdInventoryAdjustment) GetLinesOk() (*[]QbdInventoryAdjustmentLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdInventoryAdjustment) SetLines(v []QbdInventoryAdjustmentLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdInventoryAdjustment) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdInventoryAdjustment) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdInventoryAdjustment) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


