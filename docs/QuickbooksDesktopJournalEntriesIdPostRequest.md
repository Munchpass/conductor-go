# QuickbooksDesktopJournalEntriesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the journal entry object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | Pointer to **string** | The date of this journal entry, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this journal entry, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**IsAdjustment** | Pointer to **bool** | Indicates whether this journal entry is an adjustment entry. When &#x60;true&#x60;, QuickBooks retains the original entry information to maintain an audit trail of the adjustments. | [optional] 
**AreAmountsEnteredInHomeCurrency** | Pointer to **bool** | Indicates whether the amounts in this journal entry were entered in the company&#39;s home currency rather than a foreign currency. When &#x60;true&#x60;, amounts are in the home currency regardless of the &#x60;currency&#x60; field. | [optional] 
**CurrencyId** | Pointer to **string** | The journal entry&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this journal entry&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopJournalEntriesIdPostRequestLinesInner**](QuickbooksDesktopJournalEntriesIdPostRequestLinesInner.md) | The journal entry&#39;s credit and debit lines.  **IMPORTANT**: When updating journal entries, you must include ALL existing journal lines (both credit and debit) in your update request, even if you only want to modify a single line. QuickBooks will automatically delete any existing lines that are not included in the update request, which is why all lines must be provided in a single array when updating. | [optional] 

## Methods

### NewQuickbooksDesktopJournalEntriesIdPostRequest

`func NewQuickbooksDesktopJournalEntriesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopJournalEntriesIdPostRequest`

NewQuickbooksDesktopJournalEntriesIdPostRequest instantiates a new QuickbooksDesktopJournalEntriesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopJournalEntriesIdPostRequestWithDefaults

`func NewQuickbooksDesktopJournalEntriesIdPostRequestWithDefaults() *QuickbooksDesktopJournalEntriesIdPostRequest`

NewQuickbooksDesktopJournalEntriesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopJournalEntriesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetIsAdjustment() bool`

GetIsAdjustment returns the IsAdjustment field if non-nil, zero value otherwise.

### GetIsAdjustmentOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetIsAdjustmentOk() (*bool, bool)`

GetIsAdjustmentOk returns a tuple with the IsAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetIsAdjustment(v bool)`

SetIsAdjustment sets IsAdjustment field to given value.

### HasIsAdjustment

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasIsAdjustment() bool`

HasIsAdjustment returns a boolean if a field has been set.

### GetAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetAreAmountsEnteredInHomeCurrency() bool`

GetAreAmountsEnteredInHomeCurrency returns the AreAmountsEnteredInHomeCurrency field if non-nil, zero value otherwise.

### GetAreAmountsEnteredInHomeCurrencyOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetAreAmountsEnteredInHomeCurrencyOk() (*bool, bool)`

GetAreAmountsEnteredInHomeCurrencyOk returns a tuple with the AreAmountsEnteredInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetAreAmountsEnteredInHomeCurrency(v bool)`

SetAreAmountsEnteredInHomeCurrency sets AreAmountsEnteredInHomeCurrency field to given value.

### HasAreAmountsEnteredInHomeCurrency

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasAreAmountsEnteredInHomeCurrency() bool`

HasAreAmountsEnteredInHomeCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetLines() []QuickbooksDesktopJournalEntriesIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopJournalEntriesIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) SetLines(v []QuickbooksDesktopJournalEntriesIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopJournalEntriesIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


