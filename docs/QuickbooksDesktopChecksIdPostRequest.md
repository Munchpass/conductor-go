# QuickbooksDesktopChecksIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the check object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**BankAccountId** | Pointer to **string** | The bank account from which the funds are being drawn for this check; e.g., Checking or Savings. This check will decrease the balance of this account. | [optional] 
**PayeeId** | Pointer to **string** | The person or company who will receive this check. | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this check, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.  **IMPORTANT**: For checks, this field is the check number. | [optional] 
**TransactionDate** | Pointer to **string** | The date written on this check, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**Memo** | Pointer to **string** | The memo that is printed on this check. | [optional] 
**Address** | Pointer to [**QuickbooksDesktopChecksPostRequestAddress**](QuickbooksDesktopChecksPostRequestAddress.md) |  | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this check is included in the queue of documents for QuickBooks to print. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this check, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the payee. This can be overridden on the check&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this check&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ApplyToTransactions** | Pointer to [**[]QuickbooksDesktopChecksPostRequestApplyToTransactionsInner**](QuickbooksDesktopChecksPostRequestApplyToTransactionsInner.md) | Transactions to be paid by this check. This will create a link between this check and the specified transactions.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the check and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of checks, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**ClearExpenseLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing expense lines associated with this check. To modify or add individual expense lines, use the field &#x60;expenseLines&#x60; instead. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner**](QuickbooksDesktopBillsIdPostRequestExpenseLinesInner.md) | The check&#39;s expense lines, each representing one line in this expense.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing expense lines for the check with this array. To keep any existing expense lines, you must include them in this array even if they have not changed. **Any expense lines not included will be removed.**  2. To add a new expense line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any expense lines, omit this field entirely to keep them unchanged. | [optional] 
**ClearItemLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing item lines associated with this check. To modify or add individual item lines, use the field &#x60;itemLines&#x60; instead. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLinesInner**](QuickbooksDesktopBillsIdPostRequestItemLinesInner.md) | The check&#39;s item lines, each representing the purchase of a specific item or service.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item lines for the check with this array. To keep any existing item lines, you must include them in this array even if they have not changed. **Any item lines not included will be removed.**  2. To add a new item line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item lines, omit this field entirely to keep them unchanged. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner.md) | The check&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item group lines for the check with this array. To keep any existing item group lines, you must include them in this array even if they have not changed. **Any item group lines not included will be removed.**  2. To add a new item group line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item group lines, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopChecksIdPostRequest

`func NewQuickbooksDesktopChecksIdPostRequest(revisionNumber string, ) *QuickbooksDesktopChecksIdPostRequest`

NewQuickbooksDesktopChecksIdPostRequest instantiates a new QuickbooksDesktopChecksIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopChecksIdPostRequestWithDefaults

`func NewQuickbooksDesktopChecksIdPostRequestWithDefaults() *QuickbooksDesktopChecksIdPostRequest`

NewQuickbooksDesktopChecksIdPostRequestWithDefaults instantiates a new QuickbooksDesktopChecksIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopChecksIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopChecksIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetBankAccountId

`func (o *QuickbooksDesktopChecksIdPostRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *QuickbooksDesktopChecksIdPostRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *QuickbooksDesktopChecksIdPostRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetPayeeId

`func (o *QuickbooksDesktopChecksIdPostRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *QuickbooksDesktopChecksIdPostRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *QuickbooksDesktopChecksIdPostRequest) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopChecksIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopChecksIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopChecksIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopChecksIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopChecksIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopChecksIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopChecksIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopChecksIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopChecksIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopChecksIdPostRequest) GetAddress() QuickbooksDesktopChecksPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetAddressOk() (*QuickbooksDesktopChecksPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopChecksIdPostRequest) SetAddress(v QuickbooksDesktopChecksPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopChecksIdPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopChecksIdPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopChecksIdPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopChecksIdPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopChecksIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopChecksIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopChecksIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopChecksIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopChecksIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopChecksIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopChecksIdPostRequest) GetApplyToTransactions() []QuickbooksDesktopChecksPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopChecksPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopChecksIdPostRequest) SetApplyToTransactions(v []QuickbooksDesktopChecksPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.

### HasApplyToTransactions

`func (o *QuickbooksDesktopChecksIdPostRequest) HasApplyToTransactions() bool`

HasApplyToTransactions returns a boolean if a field has been set.

### GetClearExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) GetClearExpenseLines() bool`

GetClearExpenseLines returns the ClearExpenseLines field if non-nil, zero value otherwise.

### GetClearExpenseLinesOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetClearExpenseLinesOk() (*bool, bool)`

GetClearExpenseLinesOk returns a tuple with the ClearExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) SetClearExpenseLines(v bool)`

SetClearExpenseLines sets ClearExpenseLines field to given value.

### HasClearExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) HasClearExpenseLines() bool`

HasClearExpenseLines returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) GetExpenseLines() []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopChecksIdPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetClearItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) GetClearItemLines() bool`

GetClearItemLines returns the ClearItemLines field if non-nil, zero value otherwise.

### GetClearItemLinesOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetClearItemLinesOk() (*bool, bool)`

GetClearItemLinesOk returns a tuple with the ClearItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) SetClearItemLines(v bool)`

SetClearItemLines sets ClearItemLines field to given value.

### HasClearItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) HasClearItemLines() bool`

HasClearItemLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) GetItemLines() []QuickbooksDesktopBillsIdPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) SetItemLines(v []QuickbooksDesktopBillsIdPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopChecksIdPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopChecksIdPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopChecksIdPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopChecksIdPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopChecksIdPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


