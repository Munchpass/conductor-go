# QbdJournalCreditLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this journal credit line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_journal_credit_line\&quot;&#x60;. | 
**Account** | [**QbdJournalCreditLineAccount**](QbdJournalCreditLineAccount.md) |  | 
**Amount** | Pointer to **string** | The monetary amount of this journal credit line, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this journal credit line. | [optional] 
**Entity** | Pointer to [**QbdJournalCreditLineEntity**](QbdJournalCreditLineEntity.md) |  | [optional] 
**Class** | Pointer to [**QbdJournalCreditLineClass**](QbdJournalCreditLineClass.md) |  | [optional] 
**SalesTaxItem** | Pointer to [**QbdJournalCreditLineSalesTaxItem**](QbdJournalCreditLineSalesTaxItem.md) |  | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this journal credit line. | [optional] 

## Methods

### NewQbdJournalCreditLine

`func NewQbdJournalCreditLine(id string, objectType string, account QbdJournalCreditLineAccount, ) *QbdJournalCreditLine`

NewQbdJournalCreditLine instantiates a new QbdJournalCreditLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdJournalCreditLineWithDefaults

`func NewQbdJournalCreditLineWithDefaults() *QbdJournalCreditLine`

NewQbdJournalCreditLineWithDefaults instantiates a new QbdJournalCreditLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdJournalCreditLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdJournalCreditLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdJournalCreditLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdJournalCreditLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdJournalCreditLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdJournalCreditLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *QbdJournalCreditLine) GetAccount() QbdJournalCreditLineAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdJournalCreditLine) GetAccountOk() (*QbdJournalCreditLineAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdJournalCreditLine) SetAccount(v QbdJournalCreditLineAccount)`

SetAccount sets Account field to given value.


### GetAmount

`func (o *QbdJournalCreditLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdJournalCreditLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdJournalCreditLine) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QbdJournalCreditLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QbdJournalCreditLine) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdJournalCreditLine) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdJournalCreditLine) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QbdJournalCreditLine) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetEntity

`func (o *QbdJournalCreditLine) GetEntity() QbdJournalCreditLineEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *QbdJournalCreditLine) GetEntityOk() (*QbdJournalCreditLineEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *QbdJournalCreditLine) SetEntity(v QbdJournalCreditLineEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *QbdJournalCreditLine) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetClass

`func (o *QbdJournalCreditLine) GetClass() QbdJournalCreditLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdJournalCreditLine) GetClassOk() (*QbdJournalCreditLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdJournalCreditLine) SetClass(v QbdJournalCreditLineClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *QbdJournalCreditLine) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSalesTaxItem

`func (o *QbdJournalCreditLine) GetSalesTaxItem() QbdJournalCreditLineSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdJournalCreditLine) GetSalesTaxItemOk() (*QbdJournalCreditLineSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdJournalCreditLine) SetSalesTaxItem(v QbdJournalCreditLineSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.

### HasSalesTaxItem

`func (o *QbdJournalCreditLine) HasSalesTaxItem() bool`

HasSalesTaxItem returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QbdJournalCreditLine) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QbdJournalCreditLine) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QbdJournalCreditLine) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QbdJournalCreditLine) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


