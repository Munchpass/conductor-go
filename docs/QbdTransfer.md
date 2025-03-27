# QbdTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this transfer. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_transfer\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this transfer was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this transfer was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this transfer object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | **string** | The date of this transfer, in ISO 8601 format (YYYY-MM-DD). | 
**SourceAccount** | [**QbdTransferSourceAccount**](QbdTransferSourceAccount.md) |  | 
**SourceAccountBalance** | **string** | The balance of the account from which money will be transferred. | 
**TargetAccount** | [**QbdTransferTargetAccount**](QbdTransferTargetAccount.md) |  | 
**TargetAccountBalance** | **string** | The balance of the account to which money will be transferred. | 
**Class** | [**QbdTransferClass**](QbdTransferClass.md) |  | 
**Amount** | **string** | The monetary amount of this transfer, represented as a decimal string. | 
**Memo** | **string** | A memo or note for this transfer. | 

## Methods

### NewQbdTransfer

`func NewQbdTransfer(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, transactionDate string, sourceAccount QbdTransferSourceAccount, sourceAccountBalance string, targetAccount QbdTransferTargetAccount, targetAccountBalance string, class QbdTransferClass, amount string, memo string, ) *QbdTransfer`

NewQbdTransfer instantiates a new QbdTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdTransferWithDefaults

`func NewQbdTransferWithDefaults() *QbdTransfer`

NewQbdTransferWithDefaults instantiates a new QbdTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdTransfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdTransfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdTransfer) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdTransfer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdTransfer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdTransfer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdTransfer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdTransfer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdTransfer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdTransfer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdTransfer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdTransfer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdTransfer) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdTransfer) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdTransfer) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QbdTransfer) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdTransfer) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdTransfer) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetSourceAccount

`func (o *QbdTransfer) GetSourceAccount() QbdTransferSourceAccount`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *QbdTransfer) GetSourceAccountOk() (*QbdTransferSourceAccount, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *QbdTransfer) SetSourceAccount(v QbdTransferSourceAccount)`

SetSourceAccount sets SourceAccount field to given value.


### GetSourceAccountBalance

`func (o *QbdTransfer) GetSourceAccountBalance() string`

GetSourceAccountBalance returns the SourceAccountBalance field if non-nil, zero value otherwise.

### GetSourceAccountBalanceOk

`func (o *QbdTransfer) GetSourceAccountBalanceOk() (*string, bool)`

GetSourceAccountBalanceOk returns a tuple with the SourceAccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountBalance

`func (o *QbdTransfer) SetSourceAccountBalance(v string)`

SetSourceAccountBalance sets SourceAccountBalance field to given value.


### GetTargetAccount

`func (o *QbdTransfer) GetTargetAccount() QbdTransferTargetAccount`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *QbdTransfer) GetTargetAccountOk() (*QbdTransferTargetAccount, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *QbdTransfer) SetTargetAccount(v QbdTransferTargetAccount)`

SetTargetAccount sets TargetAccount field to given value.


### GetTargetAccountBalance

`func (o *QbdTransfer) GetTargetAccountBalance() string`

GetTargetAccountBalance returns the TargetAccountBalance field if non-nil, zero value otherwise.

### GetTargetAccountBalanceOk

`func (o *QbdTransfer) GetTargetAccountBalanceOk() (*string, bool)`

GetTargetAccountBalanceOk returns a tuple with the TargetAccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountBalance

`func (o *QbdTransfer) SetTargetAccountBalance(v string)`

SetTargetAccountBalance sets TargetAccountBalance field to given value.


### GetClass

`func (o *QbdTransfer) GetClass() QbdTransferClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdTransfer) GetClassOk() (*QbdTransferClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdTransfer) SetClass(v QbdTransferClass)`

SetClass sets Class field to given value.


### GetAmount

`func (o *QbdTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *QbdTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


