# QbdJournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this journal entry. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_journal_entry\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this journal entry was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this journal entry was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this journal entry object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | **string** | The date of this journal entry, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this journal entry, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**IsAdjustment** | Pointer to **bool** | Indicates whether this journal entry is an adjustment entry. When &#x60;true&#x60;, QuickBooks retains the original entry information to maintain an audit trail of the adjustments. | [optional] 
**IsHomeCurrencyAdjustment** | Pointer to **bool** | Indicates whether this journal entry is an adjustment made in the company&#39;s home currency for a transaction that was originally recorded in a foreign currency. | [optional] 
**AreAmountsEnteredInHomeCurrency** | Pointer to **bool** | Indicates whether the amounts in this journal entry were entered in the company&#39;s home currency rather than a foreign currency. When &#x60;true&#x60;, amounts are in the home currency regardless of the &#x60;currency&#x60; field. | [optional] 
**Currency** | Pointer to [**QbdJournalEntryCurrency**](QbdJournalEntryCurrency.md) |  | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this journal entry&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | [optional] 
**DebitLines** | [**[]QbdJournalDebitLine**](QbdJournalDebitLine.md) | The journal entry&#39;s debit lines. | 
**CreditLines** | [**[]QbdJournalCreditLine**](QbdJournalCreditLine.md) | The journal entry&#39;s credit lines. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the journal entry object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdJournalEntry

`func NewQbdJournalEntry(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, transactionDate string, debitLines []QbdJournalDebitLine, creditLines []QbdJournalCreditLine, customFields []QbdCustomField, ) *QbdJournalEntry`

NewQbdJournalEntry instantiates a new QbdJournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdJournalEntryWithDefaults

`func NewQbdJournalEntryWithDefaults() *QbdJournalEntry`

NewQbdJournalEntryWithDefaults instantiates a new QbdJournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdJournalEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdJournalEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdJournalEntry) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdJournalEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdJournalEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdJournalEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdJournalEntry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdJournalEntry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdJournalEntry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdJournalEntry) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdJournalEntry) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdJournalEntry) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdJournalEntry) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdJournalEntry) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdJournalEntry) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QbdJournalEntry) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdJournalEntry) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdJournalEntry) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdJournalEntry) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdJournalEntry) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdJournalEntry) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QbdJournalEntry) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetIsAdjustment

`func (o *QbdJournalEntry) GetIsAdjustment() bool`

GetIsAdjustment returns the IsAdjustment field if non-nil, zero value otherwise.

### GetIsAdjustmentOk

`func (o *QbdJournalEntry) GetIsAdjustmentOk() (*bool, bool)`

GetIsAdjustmentOk returns a tuple with the IsAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdjustment

`func (o *QbdJournalEntry) SetIsAdjustment(v bool)`

SetIsAdjustment sets IsAdjustment field to given value.

### HasIsAdjustment

`func (o *QbdJournalEntry) HasIsAdjustment() bool`

HasIsAdjustment returns a boolean if a field has been set.

### GetIsHomeCurrencyAdjustment

`func (o *QbdJournalEntry) GetIsHomeCurrencyAdjustment() bool`

GetIsHomeCurrencyAdjustment returns the IsHomeCurrencyAdjustment field if non-nil, zero value otherwise.

### GetIsHomeCurrencyAdjustmentOk

`func (o *QbdJournalEntry) GetIsHomeCurrencyAdjustmentOk() (*bool, bool)`

GetIsHomeCurrencyAdjustmentOk returns a tuple with the IsHomeCurrencyAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHomeCurrencyAdjustment

`func (o *QbdJournalEntry) SetIsHomeCurrencyAdjustment(v bool)`

SetIsHomeCurrencyAdjustment sets IsHomeCurrencyAdjustment field to given value.

### HasIsHomeCurrencyAdjustment

`func (o *QbdJournalEntry) HasIsHomeCurrencyAdjustment() bool`

HasIsHomeCurrencyAdjustment returns a boolean if a field has been set.

### GetAreAmountsEnteredInHomeCurrency

`func (o *QbdJournalEntry) GetAreAmountsEnteredInHomeCurrency() bool`

GetAreAmountsEnteredInHomeCurrency returns the AreAmountsEnteredInHomeCurrency field if non-nil, zero value otherwise.

### GetAreAmountsEnteredInHomeCurrencyOk

`func (o *QbdJournalEntry) GetAreAmountsEnteredInHomeCurrencyOk() (*bool, bool)`

GetAreAmountsEnteredInHomeCurrencyOk returns a tuple with the AreAmountsEnteredInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreAmountsEnteredInHomeCurrency

`func (o *QbdJournalEntry) SetAreAmountsEnteredInHomeCurrency(v bool)`

SetAreAmountsEnteredInHomeCurrency sets AreAmountsEnteredInHomeCurrency field to given value.

### HasAreAmountsEnteredInHomeCurrency

`func (o *QbdJournalEntry) HasAreAmountsEnteredInHomeCurrency() bool`

HasAreAmountsEnteredInHomeCurrency returns a boolean if a field has been set.

### GetCurrency

`func (o *QbdJournalEntry) GetCurrency() QbdJournalEntryCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdJournalEntry) GetCurrencyOk() (*QbdJournalEntryCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdJournalEntry) SetCurrency(v QbdJournalEntryCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QbdJournalEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QbdJournalEntry) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdJournalEntry) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdJournalEntry) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QbdJournalEntry) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QbdJournalEntry) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdJournalEntry) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdJournalEntry) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QbdJournalEntry) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDebitLines

`func (o *QbdJournalEntry) GetDebitLines() []QbdJournalDebitLine`

GetDebitLines returns the DebitLines field if non-nil, zero value otherwise.

### GetDebitLinesOk

`func (o *QbdJournalEntry) GetDebitLinesOk() (*[]QbdJournalDebitLine, bool)`

GetDebitLinesOk returns a tuple with the DebitLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitLines

`func (o *QbdJournalEntry) SetDebitLines(v []QbdJournalDebitLine)`

SetDebitLines sets DebitLines field to given value.


### GetCreditLines

`func (o *QbdJournalEntry) GetCreditLines() []QbdJournalCreditLine`

GetCreditLines returns the CreditLines field if non-nil, zero value otherwise.

### GetCreditLinesOk

`func (o *QbdJournalEntry) GetCreditLinesOk() (*[]QbdJournalCreditLine, bool)`

GetCreditLinesOk returns a tuple with the CreditLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLines

`func (o *QbdJournalEntry) SetCreditLines(v []QbdJournalCreditLine)`

SetCreditLines sets CreditLines field to given value.


### GetCustomFields

`func (o *QbdJournalEntry) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdJournalEntry) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdJournalEntry) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


