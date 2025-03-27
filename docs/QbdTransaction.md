# QbdTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | **string** | The type of transaction. | 
**TransactionId** | **string** | The QuickBooks-assigned unique identifier of this transaction. If &#x60;transactionLineId&#x60; is also defined, this is the identifier of the line&#39;s parent transaction object. | 
**TransactionLineId** | **string** | The QuickBooks-assigned unique identifier of this transaction line. If &#x60;null&#x60;, this result is a transaction object. | 
**CreatedAt** | **string** | The date and time when this transaction was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this transaction was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**Entity** | [**QbdTransactionEntity**](QbdTransactionEntity.md) |  | 
**Account** | [**QbdTransactionAccount**](QbdTransactionAccount.md) |  | 
**TransactionDate** | **string** | The date of this transaction, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this transaction, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Amount** | **string** | The monetary amount of this transaction, represented as a decimal string. | 
**Currency** | [**QbdTransactionCurrency**](QbdTransactionCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this transaction&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountInHomeCurrency** | **string** | The monetary amount of this transaction converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | A memo or note for this transaction. | 

## Methods

### NewQbdTransaction

`func NewQbdTransaction(transactionType string, transactionId string, transactionLineId string, createdAt string, updatedAt string, entity QbdTransactionEntity, account QbdTransactionAccount, transactionDate string, refNumber string, amount string, currency QbdTransactionCurrency, exchangeRate float32, amountInHomeCurrency string, memo string, ) *QbdTransaction`

NewQbdTransaction instantiates a new QbdTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdTransactionWithDefaults

`func NewQbdTransactionWithDefaults() *QbdTransaction`

NewQbdTransactionWithDefaults instantiates a new QbdTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *QbdTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *QbdTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *QbdTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetTransactionId

`func (o *QbdTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *QbdTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *QbdTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionLineId

`func (o *QbdTransaction) GetTransactionLineId() string`

GetTransactionLineId returns the TransactionLineId field if non-nil, zero value otherwise.

### GetTransactionLineIdOk

`func (o *QbdTransaction) GetTransactionLineIdOk() (*string, bool)`

GetTransactionLineIdOk returns a tuple with the TransactionLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLineId

`func (o *QbdTransaction) SetTransactionLineId(v string)`

SetTransactionLineId sets TransactionLineId field to given value.


### GetCreatedAt

`func (o *QbdTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdTransaction) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdTransaction) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdTransaction) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEntity

`func (o *QbdTransaction) GetEntity() QbdTransactionEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *QbdTransaction) GetEntityOk() (*QbdTransactionEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *QbdTransaction) SetEntity(v QbdTransactionEntity)`

SetEntity sets Entity field to given value.


### GetAccount

`func (o *QbdTransaction) GetAccount() QbdTransactionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdTransaction) GetAccountOk() (*QbdTransactionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdTransaction) SetAccount(v QbdTransactionAccount)`

SetAccount sets Account field to given value.


### GetTransactionDate

`func (o *QbdTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdTransaction) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdTransaction) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdTransaction) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetAmount

`func (o *QbdTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *QbdTransaction) GetCurrency() QbdTransactionCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdTransaction) GetCurrencyOk() (*QbdTransactionCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdTransaction) SetCurrency(v QbdTransactionCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdTransaction) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdTransaction) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdTransaction) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountInHomeCurrency

`func (o *QbdTransaction) GetAmountInHomeCurrency() string`

GetAmountInHomeCurrency returns the AmountInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountInHomeCurrencyOk

`func (o *QbdTransaction) GetAmountInHomeCurrencyOk() (*string, bool)`

GetAmountInHomeCurrencyOk returns a tuple with the AmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInHomeCurrency

`func (o *QbdTransaction) SetAmountInHomeCurrency(v string)`

SetAmountInHomeCurrency sets AmountInHomeCurrency field to given value.


### GetMemo

`func (o *QbdTransaction) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdTransaction) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdTransaction) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


