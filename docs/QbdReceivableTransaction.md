# QbdReceivableTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The ID of the receivable transaction to which this payment is applied. | 
**TransactionType** | **string** | The type of transaction for this receivable transaction. | 
**TransactionDate** | **string** | The date of this receivable transaction, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this receivable transaction, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BalanceRemaining** | **string** | The outstanding balance of this receivable transaction after applying any credits or payments. Represented as a decimal string. | 
**Amount** | **string** | The monetary amount of this receivable transaction, represented as a decimal string. | 
**DiscountAmount** | **string** | The monetary amount by which to reduce the receivable transaction&#39;s receivable amount, represented as a decimal string. | 
**DiscountAccount** | [**QbdReceivableTransactionDiscountAccount**](QbdReceivableTransactionDiscountAccount.md) |  | 
**DiscountClass** | [**QbdReceivableTransactionDiscountClass**](QbdReceivableTransactionDiscountClass.md) |  | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The receivable transaction&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of receivable transactions to receive this field because it is not returned by default. | 

## Methods

### NewQbdReceivableTransaction

`func NewQbdReceivableTransaction(transactionId string, transactionType string, transactionDate string, refNumber string, balanceRemaining string, amount string, discountAmount string, discountAccount QbdReceivableTransactionDiscountAccount, discountClass QbdReceivableTransactionDiscountClass, linkedTransactions []QbdLinkedTransaction, ) *QbdReceivableTransaction`

NewQbdReceivableTransaction instantiates a new QbdReceivableTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdReceivableTransactionWithDefaults

`func NewQbdReceivableTransactionWithDefaults() *QbdReceivableTransaction`

NewQbdReceivableTransactionWithDefaults instantiates a new QbdReceivableTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *QbdReceivableTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *QbdReceivableTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *QbdReceivableTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionType

`func (o *QbdReceivableTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *QbdReceivableTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *QbdReceivableTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetTransactionDate

`func (o *QbdReceivableTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdReceivableTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdReceivableTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdReceivableTransaction) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdReceivableTransaction) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdReceivableTransaction) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBalanceRemaining

`func (o *QbdReceivableTransaction) GetBalanceRemaining() string`

GetBalanceRemaining returns the BalanceRemaining field if non-nil, zero value otherwise.

### GetBalanceRemainingOk

`func (o *QbdReceivableTransaction) GetBalanceRemainingOk() (*string, bool)`

GetBalanceRemainingOk returns a tuple with the BalanceRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceRemaining

`func (o *QbdReceivableTransaction) SetBalanceRemaining(v string)`

SetBalanceRemaining sets BalanceRemaining field to given value.


### GetAmount

`func (o *QbdReceivableTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdReceivableTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdReceivableTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDiscountAmount

`func (o *QbdReceivableTransaction) GetDiscountAmount() string`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *QbdReceivableTransaction) GetDiscountAmountOk() (*string, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *QbdReceivableTransaction) SetDiscountAmount(v string)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetDiscountAccount

`func (o *QbdReceivableTransaction) GetDiscountAccount() QbdReceivableTransactionDiscountAccount`

GetDiscountAccount returns the DiscountAccount field if non-nil, zero value otherwise.

### GetDiscountAccountOk

`func (o *QbdReceivableTransaction) GetDiscountAccountOk() (*QbdReceivableTransactionDiscountAccount, bool)`

GetDiscountAccountOk returns a tuple with the DiscountAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAccount

`func (o *QbdReceivableTransaction) SetDiscountAccount(v QbdReceivableTransactionDiscountAccount)`

SetDiscountAccount sets DiscountAccount field to given value.


### GetDiscountClass

`func (o *QbdReceivableTransaction) GetDiscountClass() QbdReceivableTransactionDiscountClass`

GetDiscountClass returns the DiscountClass field if non-nil, zero value otherwise.

### GetDiscountClassOk

`func (o *QbdReceivableTransaction) GetDiscountClassOk() (*QbdReceivableTransactionDiscountClass, bool)`

GetDiscountClassOk returns a tuple with the DiscountClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountClass

`func (o *QbdReceivableTransaction) SetDiscountClass(v QbdReceivableTransactionDiscountClass)`

SetDiscountClass sets DiscountClass field to given value.


### GetLinkedTransactions

`func (o *QbdReceivableTransaction) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdReceivableTransaction) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdReceivableTransaction) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


