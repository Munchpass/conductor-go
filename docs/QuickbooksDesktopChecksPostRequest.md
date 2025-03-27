# QuickbooksDesktopChecksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountId** | **string** | The bank account from which the funds are being drawn for this check; e.g., Checking or Savings. This check will decrease the balance of this account. | 
**PayeeId** | Pointer to **string** | The person or company who will receive this check. | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this check, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment).  **IMPORTANT**: For checks, this field is the check number. | [optional] 
**TransactionDate** | **string** | The date written on this check, in ISO 8601 format (YYYY-MM-DD). | 
**Memo** | Pointer to **string** | The memo that is printed on this check. | [optional] 
**Address** | Pointer to [**QuickbooksDesktopChecksPostRequestAddress**](QuickbooksDesktopChecksPostRequestAddress.md) |  | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this check is included in the queue of documents for QuickBooks to print. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this check, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the payee. This can be overridden on the check&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this check&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**ApplyToTransactions** | Pointer to [**[]QuickbooksDesktopChecksPostRequestApplyToTransactionsInner**](QuickbooksDesktopChecksPostRequestApplyToTransactionsInner.md) | Transactions to be paid by this check. This will create a link between this check and the specified transactions.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the check and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of checks, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInner.md) | The check&#39;s expense lines, each representing one line in this expense. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLinesInner**](QuickbooksDesktopBillsPostRequestItemLinesInner.md) | The check&#39;s item lines, each representing the purchase of a specific item or service. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsPostRequestItemLineGroupsInner.md) | The check&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | [optional] 

## Methods

### NewQuickbooksDesktopChecksPostRequest

`func NewQuickbooksDesktopChecksPostRequest(bankAccountId string, transactionDate string, ) *QuickbooksDesktopChecksPostRequest`

NewQuickbooksDesktopChecksPostRequest instantiates a new QuickbooksDesktopChecksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopChecksPostRequestWithDefaults

`func NewQuickbooksDesktopChecksPostRequestWithDefaults() *QuickbooksDesktopChecksPostRequest`

NewQuickbooksDesktopChecksPostRequestWithDefaults instantiates a new QuickbooksDesktopChecksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountId

`func (o *QuickbooksDesktopChecksPostRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *QuickbooksDesktopChecksPostRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *QuickbooksDesktopChecksPostRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetPayeeId

`func (o *QuickbooksDesktopChecksPostRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *QuickbooksDesktopChecksPostRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *QuickbooksDesktopChecksPostRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *QuickbooksDesktopChecksPostRequest) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopChecksPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopChecksPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopChecksPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopChecksPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopChecksPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopChecksPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopChecksPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetMemo

`func (o *QuickbooksDesktopChecksPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopChecksPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopChecksPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopChecksPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopChecksPostRequest) GetAddress() QuickbooksDesktopChecksPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopChecksPostRequest) GetAddressOk() (*QuickbooksDesktopChecksPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopChecksPostRequest) SetAddress(v QuickbooksDesktopChecksPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopChecksPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopChecksPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopChecksPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopChecksPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopChecksPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopChecksPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopChecksPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopChecksPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopChecksPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopChecksPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopChecksPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopChecksPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopChecksPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopChecksPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopChecksPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopChecksPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopChecksPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopChecksPostRequest) GetApplyToTransactions() []QuickbooksDesktopChecksPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopChecksPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopChecksPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopChecksPostRequest) SetApplyToTransactions(v []QuickbooksDesktopChecksPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.

### HasApplyToTransactions

`func (o *QuickbooksDesktopChecksPostRequest) HasApplyToTransactions() bool`

HasApplyToTransactions returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopChecksPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopChecksPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopChecksPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopChecksPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopChecksPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopChecksPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopChecksPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopChecksPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopChecksPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopChecksPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopChecksPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopChecksPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


