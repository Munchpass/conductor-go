# QbdLinkedTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this linked transaction. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_linked_transaction\&quot;&#x60;. | 
**TransactionType** | **string** | The type of transaction for this linked transaction. | 
**TransactionDate** | **string** | The date of this linked transaction, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **NullableString** | The case-sensitive user-defined reference number for this linked transaction, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**LinkType** | **NullableString** | Indicates the nature of the link between the transactions: &#x60;amount&#x60; denotes an amount-based link (e.g., an invoice linked to a payment), and &#x60;quantity&#x60; denotes a quantity-based link (e.g., an invoice created from a sales order based on the quantity of items received). | 
**Amount** | **NullableString** | The monetary amount of this linked transaction, represented as a decimal string. | 

## Methods

### NewQbdLinkedTransaction

`func NewQbdLinkedTransaction(id string, objectType string, transactionType string, transactionDate string, refNumber NullableString, linkType NullableString, amount NullableString, ) *QbdLinkedTransaction`

NewQbdLinkedTransaction instantiates a new QbdLinkedTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdLinkedTransactionWithDefaults

`func NewQbdLinkedTransactionWithDefaults() *QbdLinkedTransaction`

NewQbdLinkedTransactionWithDefaults instantiates a new QbdLinkedTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdLinkedTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdLinkedTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdLinkedTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdLinkedTransaction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdLinkedTransaction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdLinkedTransaction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTransactionType

`func (o *QbdLinkedTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *QbdLinkedTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *QbdLinkedTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetTransactionDate

`func (o *QbdLinkedTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdLinkedTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdLinkedTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdLinkedTransaction) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdLinkedTransaction) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdLinkedTransaction) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### SetRefNumberNil

`func (o *QbdLinkedTransaction) SetRefNumberNil(b bool)`

 SetRefNumberNil sets the value for RefNumber to be an explicit nil

### UnsetRefNumber
`func (o *QbdLinkedTransaction) UnsetRefNumber()`

UnsetRefNumber ensures that no value is present for RefNumber, not even an explicit nil
### GetLinkType

`func (o *QbdLinkedTransaction) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *QbdLinkedTransaction) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *QbdLinkedTransaction) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### SetLinkTypeNil

`func (o *QbdLinkedTransaction) SetLinkTypeNil(b bool)`

 SetLinkTypeNil sets the value for LinkType to be an explicit nil

### UnsetLinkType
`func (o *QbdLinkedTransaction) UnsetLinkType()`

UnsetLinkType ensures that no value is present for LinkType, not even an explicit nil
### GetAmount

`func (o *QbdLinkedTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdLinkedTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdLinkedTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *QbdLinkedTransaction) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *QbdLinkedTransaction) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


