# QuickbooksDesktopItemReceiptsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the item receipt object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**VendorId** | Pointer to **string** | The vendor who sent this item receipt for goods or services purchased. | [optional] 
**PayablesAccountId** | Pointer to **string** | The Accounts-Payable (A/P) account to which this item receipt is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this item receipt is linked to other transactions, this A/P account must match the &#x60;payablesAccount&#x60; used in those other transactions. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this item receipt, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this item receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this item receipt. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this item receipt, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the item receipt&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this item receipt&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ClearExpenseLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing expense lines associated with this item receipt. To modify or add individual expense lines, use the field &#x60;expenseLines&#x60; instead. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner**](QuickbooksDesktopBillsIdPostRequestExpenseLinesInner.md) | The item receipt&#39;s expense lines, each representing one line in this expense.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing expense lines for the item receipt with this array. To keep any existing expense lines, you must include them in this array even if they have not changed. **Any expense lines not included will be removed.**  2. To add a new expense line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any expense lines, omit this field entirely to keep them unchanged. | [optional] 
**ClearItemLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing item lines associated with this item receipt. To modify or add individual item lines, use the field &#x60;itemLines&#x60; instead. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLinesInner**](QuickbooksDesktopBillsIdPostRequestItemLinesInner.md) | The item receipt&#39;s item lines, each representing the purchase of a specific item or service.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item lines for the item receipt with this array. To keep any existing item lines, you must include them in this array even if they have not changed. **Any item lines not included will be removed.**  2. To add a new item line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item lines, omit this field entirely to keep them unchanged. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner.md) | The item receipt&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item group lines for the item receipt with this array. To keep any existing item group lines, you must include them in this array even if they have not changed. **Any item group lines not included will be removed.**  2. To add a new item group line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item group lines, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopItemReceiptsIdPostRequest

`func NewQuickbooksDesktopItemReceiptsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopItemReceiptsIdPostRequest`

NewQuickbooksDesktopItemReceiptsIdPostRequest instantiates a new QuickbooksDesktopItemReceiptsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopItemReceiptsIdPostRequestWithDefaults

`func NewQuickbooksDesktopItemReceiptsIdPostRequestWithDefaults() *QuickbooksDesktopItemReceiptsIdPostRequest`

NewQuickbooksDesktopItemReceiptsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopItemReceiptsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendorId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetPayablesAccountId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetPayablesAccountId() string`

GetPayablesAccountId returns the PayablesAccountId field if non-nil, zero value otherwise.

### GetPayablesAccountIdOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetPayablesAccountIdOk() (*string, bool)`

GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccountId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetPayablesAccountId(v string)`

SetPayablesAccountId sets PayablesAccountId field to given value.

### HasPayablesAccountId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasPayablesAccountId() bool`

HasPayablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetClearExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetClearExpenseLines() bool`

GetClearExpenseLines returns the ClearExpenseLines field if non-nil, zero value otherwise.

### GetClearExpenseLinesOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetClearExpenseLinesOk() (*bool, bool)`

GetClearExpenseLinesOk returns a tuple with the ClearExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetClearExpenseLines(v bool)`

SetClearExpenseLines sets ClearExpenseLines field to given value.

### HasClearExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasClearExpenseLines() bool`

HasClearExpenseLines returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetExpenseLines() []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetClearItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetClearItemLines() bool`

GetClearItemLines returns the ClearItemLines field if non-nil, zero value otherwise.

### GetClearItemLinesOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetClearItemLinesOk() (*bool, bool)`

GetClearItemLinesOk returns a tuple with the ClearItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetClearItemLines(v bool)`

SetClearItemLines sets ClearItemLines field to given value.

### HasClearItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasClearItemLines() bool`

HasClearItemLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetItemLines() []QuickbooksDesktopBillsIdPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetItemLines(v []QuickbooksDesktopBillsIdPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopItemReceiptsIdPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


