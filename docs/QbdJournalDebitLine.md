# QbdJournalDebitLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this journal debit line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_journal_debit_line\&quot;&#x60;. | 
**Account** | Pointer to [**QbdJournalDebitLineAccount**](QbdJournalDebitLineAccount.md) |  | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this journal debit line, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this journal debit line. | [optional] 
**Entity** | Pointer to [**QbdJournalDebitLineEntity**](QbdJournalDebitLineEntity.md) |  | [optional] 
**Class** | Pointer to [**QbdJournalDebitLineClass**](QbdJournalDebitLineClass.md) |  | [optional] 
**SalesTaxItem** | Pointer to [**QbdJournalDebitLineSalesTaxItem**](QbdJournalDebitLineSalesTaxItem.md) |  | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this journal debit line. | [optional] 

## Methods

### NewQbdJournalDebitLine

`func NewQbdJournalDebitLine(id string, objectType string, ) *QbdJournalDebitLine`

NewQbdJournalDebitLine instantiates a new QbdJournalDebitLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdJournalDebitLineWithDefaults

`func NewQbdJournalDebitLineWithDefaults() *QbdJournalDebitLine`

NewQbdJournalDebitLineWithDefaults instantiates a new QbdJournalDebitLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdJournalDebitLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdJournalDebitLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdJournalDebitLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdJournalDebitLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdJournalDebitLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdJournalDebitLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *QbdJournalDebitLine) GetAccount() QbdJournalDebitLineAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdJournalDebitLine) GetAccountOk() (*QbdJournalDebitLineAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdJournalDebitLine) SetAccount(v QbdJournalDebitLineAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *QbdJournalDebitLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *QbdJournalDebitLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdJournalDebitLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdJournalDebitLine) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QbdJournalDebitLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QbdJournalDebitLine) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdJournalDebitLine) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdJournalDebitLine) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QbdJournalDebitLine) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetEntity

`func (o *QbdJournalDebitLine) GetEntity() QbdJournalDebitLineEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *QbdJournalDebitLine) GetEntityOk() (*QbdJournalDebitLineEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *QbdJournalDebitLine) SetEntity(v QbdJournalDebitLineEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *QbdJournalDebitLine) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetClass

`func (o *QbdJournalDebitLine) GetClass() QbdJournalDebitLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdJournalDebitLine) GetClassOk() (*QbdJournalDebitLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdJournalDebitLine) SetClass(v QbdJournalDebitLineClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *QbdJournalDebitLine) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSalesTaxItem

`func (o *QbdJournalDebitLine) GetSalesTaxItem() QbdJournalDebitLineSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdJournalDebitLine) GetSalesTaxItemOk() (*QbdJournalDebitLineSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdJournalDebitLine) SetSalesTaxItem(v QbdJournalDebitLineSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.

### HasSalesTaxItem

`func (o *QbdJournalDebitLine) HasSalesTaxItem() bool`

HasSalesTaxItem returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QbdJournalDebitLine) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QbdJournalDebitLine) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QbdJournalDebitLine) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QbdJournalDebitLine) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


