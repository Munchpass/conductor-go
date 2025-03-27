# QuickbooksDesktopJournalEntriesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | **string** | The date of this journal entry, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this journal entry, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**IsAdjustment** | Pointer to **bool** | Indicates whether this journal entry is an adjustment entry. When &#x60;true&#x60;, QuickBooks retains the original entry information to maintain an audit trail of the adjustments. | [optional] 
**IsHomeCurrencyAdjustment** | Pointer to **bool** | Indicates whether this journal entry is an adjustment made in the company&#39;s home currency for a transaction that was originally recorded in a foreign currency. | [optional] 
**AreAmountsEnteredInHomeCurrency** | Pointer to **bool** | Indicates whether the amounts in this journal entry were entered in the company&#39;s home currency rather than a foreign currency. When &#x60;true&#x60;, amounts are in the home currency regardless of the &#x60;currency&#x60; field. | [optional] 
**CurrencyId** | Pointer to **string** | The journal entry&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this journal entry&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**DebitLines** | Pointer to [**[]QuickbooksDesktopJournalEntriesPostRequestDebitLinesInner**](QuickbooksDesktopJournalEntriesPostRequestDebitLinesInner.md) | The journal entry&#39;s debit lines. | [optional] 
**CreditLines** | Pointer to [**[]QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner**](QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner.md) | The journal entry&#39;s credit lines. | [optional] 

## Methods

### NewQuickbooksDesktopJournalEntriesPostRequest

`func NewQuickbooksDesktopJournalEntriesPostRequest(transactionDate string, ) *QuickbooksDesktopJournalEntriesPostRequest`

NewQuickbooksDesktopJournalEntriesPostRequest instantiates a new QuickbooksDesktopJournalEntriesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopJournalEntriesPostRequestWithDefaults

`func NewQuickbooksDesktopJournalEntriesPostRequestWithDefaults() *QuickbooksDesktopJournalEntriesPostRequest`

NewQuickbooksDesktopJournalEntriesPostRequestWithDefaults instantiates a new QuickbooksDesktopJournalEntriesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetIsAdjustment() bool`

GetIsAdjustment returns the IsAdjustment field if non-nil, zero value otherwise.

### GetIsAdjustmentOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetIsAdjustmentOk() (*bool, bool)`

GetIsAdjustmentOk returns a tuple with the IsAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetIsAdjustment(v bool)`

SetIsAdjustment sets IsAdjustment field to given value.

### HasIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasIsAdjustment() bool`

HasIsAdjustment returns a boolean if a field has been set.

### GetIsHomeCurrencyAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetIsHomeCurrencyAdjustment() bool`

GetIsHomeCurrencyAdjustment returns the IsHomeCurrencyAdjustment field if non-nil, zero value otherwise.

### GetIsHomeCurrencyAdjustmentOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetIsHomeCurrencyAdjustmentOk() (*bool, bool)`

GetIsHomeCurrencyAdjustmentOk returns a tuple with the IsHomeCurrencyAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHomeCurrencyAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetIsHomeCurrencyAdjustment(v bool)`

SetIsHomeCurrencyAdjustment sets IsHomeCurrencyAdjustment field to given value.

### HasIsHomeCurrencyAdjustment

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasIsHomeCurrencyAdjustment() bool`

HasIsHomeCurrencyAdjustment returns a boolean if a field has been set.

### GetAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetAreAmountsEnteredInHomeCurrency() bool`

GetAreAmountsEnteredInHomeCurrency returns the AreAmountsEnteredInHomeCurrency field if non-nil, zero value otherwise.

### GetAreAmountsEnteredInHomeCurrencyOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetAreAmountsEnteredInHomeCurrencyOk() (*bool, bool)`

GetAreAmountsEnteredInHomeCurrencyOk returns a tuple with the AreAmountsEnteredInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetAreAmountsEnteredInHomeCurrency(v bool)`

SetAreAmountsEnteredInHomeCurrency sets AreAmountsEnteredInHomeCurrency field to given value.

### HasAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasAreAmountsEnteredInHomeCurrency() bool`

HasAreAmountsEnteredInHomeCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDebitLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetDebitLines() []QuickbooksDesktopJournalEntriesPostRequestDebitLinesInner`

GetDebitLines returns the DebitLines field if non-nil, zero value otherwise.

### GetDebitLinesOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetDebitLinesOk() (*[]QuickbooksDesktopJournalEntriesPostRequestDebitLinesInner, bool)`

GetDebitLinesOk returns a tuple with the DebitLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetDebitLines(v []QuickbooksDesktopJournalEntriesPostRequestDebitLinesInner)`

SetDebitLines sets DebitLines field to given value.

### HasDebitLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasDebitLines() bool`

HasDebitLines returns a boolean if a field has been set.

### GetCreditLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetCreditLines() []QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner`

GetCreditLines returns the CreditLines field if non-nil, zero value otherwise.

### GetCreditLinesOk

`func (o *QuickbooksDesktopJournalEntriesPostRequest) GetCreditLinesOk() (*[]QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner, bool)`

GetCreditLinesOk returns a tuple with the CreditLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) SetCreditLines(v []QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner)`

SetCreditLines sets CreditLines field to given value.

### HasCreditLines

`func (o *QuickbooksDesktopJournalEntriesPostRequest) HasCreditLines() bool`

HasCreditLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


