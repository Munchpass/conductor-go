# QuickbooksDesktopBillCheckPaymentsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the bill check payment object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | Pointer to **string** | The date of this bill check payment, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**BankAccountId** | Pointer to **string** | The bank account from which the funds are being drawn for this bill check payment; e.g., Checking or Savings. This bill check payment will decrease the balance of this account. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this bill check payment, represented as a decimal string. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this bill check payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this bill check payment is included in the queue of documents for QuickBooks to print. | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this bill check payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.  **IMPORTANT**: For checks, this field is the check number. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this bill check payment. | [optional] 
**ApplyToTransactions** | Pointer to [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner.md) | The bills to be paid by this bill check payment. This will create a link between this bill check payment and the specified bills.  **IMPORTANT**: In each &#x60;applyToTransactions&#x60; object, you must specify either &#x60;paymentAmount&#x60;, &#x60;applyCredits&#x60;, &#x60;discountAmount&#x60;, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target bill must have &#x60;isPaid&#x3D;false&#x60;, otherwise, QuickBooks will report this object as \&quot;cannot be found\&quot;. | [optional] 

## Methods

### NewQuickbooksDesktopBillCheckPaymentsIdPostRequest

`func NewQuickbooksDesktopBillCheckPaymentsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopBillCheckPaymentsIdPostRequest`

NewQuickbooksDesktopBillCheckPaymentsIdPostRequest instantiates a new QuickbooksDesktopBillCheckPaymentsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillCheckPaymentsIdPostRequestWithDefaults

`func NewQuickbooksDesktopBillCheckPaymentsIdPostRequestWithDefaults() *QuickbooksDesktopBillCheckPaymentsIdPostRequest`

NewQuickbooksDesktopBillCheckPaymentsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopBillCheckPaymentsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetBankAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.

### HasApplyToTransactions

`func (o *QuickbooksDesktopBillCheckPaymentsIdPostRequest) HasApplyToTransactions() bool`

HasApplyToTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


