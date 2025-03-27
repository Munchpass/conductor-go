# QuickbooksDesktopBillsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the bill object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**VendorId** | Pointer to **string** | The vendor who sent this bill for goods or services purchased. | [optional] 
**VendorAddress** | Pointer to [**QuickbooksDesktopBillsPostRequestVendorAddress**](QuickbooksDesktopBillsPostRequestVendorAddress.md) |  | [optional] 
**PayablesAccountId** | Pointer to **string** | The Accounts-Payable (A/P) account to which this bill is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this bill is linked to other transactions, this A/P account must match the &#x60;payablesAccount&#x60; used in those other transactions. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this bill, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**DueDate** | Pointer to **string** | The date by which this bill must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this bill, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**TermsId** | Pointer to **string** | The bill&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this bill that appears in the Accounts-Payable register and in reports that include this bill. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this bill, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the bill&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this bill&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ClearExpenseLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing expense lines associated with this bill. To modify or add individual expense lines, use the field &#x60;expenseLines&#x60; instead. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner**](QuickbooksDesktopBillsIdPostRequestExpenseLinesInner.md) | The bill&#39;s expense lines, each representing one line in this expense.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing expense lines for the bill with this array. To keep any existing expense lines, you must include them in this array even if they have not changed. **Any expense lines not included will be removed.**  2. To add a new expense line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any expense lines, omit this field entirely to keep them unchanged. | [optional] 
**ClearItemLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing item lines associated with this bill. To modify or add individual item lines, use the field &#x60;itemLines&#x60; instead. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLinesInner**](QuickbooksDesktopBillsIdPostRequestItemLinesInner.md) | The bill&#39;s item lines, each representing the purchase of a specific item or service.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item lines for the bill with this array. To keep any existing item lines, you must include them in this array even if they have not changed. **Any item lines not included will be removed.**  2. To add a new item line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item lines, omit this field entirely to keep them unchanged. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner.md) | The bill&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item group lines for the bill with this array. To keep any existing item group lines, you must include them in this array even if they have not changed. **Any item group lines not included will be removed.**  2. To add a new item group line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item group lines, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopBillsIdPostRequest

`func NewQuickbooksDesktopBillsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopBillsIdPostRequest`

NewQuickbooksDesktopBillsIdPostRequest instantiates a new QuickbooksDesktopBillsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsIdPostRequestWithDefaults

`func NewQuickbooksDesktopBillsIdPostRequestWithDefaults() *QuickbooksDesktopBillsIdPostRequest`

NewQuickbooksDesktopBillsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopBillsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopBillsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopBillsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendorId

`func (o *QuickbooksDesktopBillsIdPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopBillsIdPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *QuickbooksDesktopBillsIdPostRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorAddress

`func (o *QuickbooksDesktopBillsIdPostRequest) GetVendorAddress() QuickbooksDesktopBillsPostRequestVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetVendorAddressOk() (*QuickbooksDesktopBillsPostRequestVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QuickbooksDesktopBillsIdPostRequest) SetVendorAddress(v QuickbooksDesktopBillsPostRequestVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.

### HasVendorAddress

`func (o *QuickbooksDesktopBillsIdPostRequest) HasVendorAddress() bool`

HasVendorAddress returns a boolean if a field has been set.

### GetPayablesAccountId

`func (o *QuickbooksDesktopBillsIdPostRequest) GetPayablesAccountId() string`

GetPayablesAccountId returns the PayablesAccountId field if non-nil, zero value otherwise.

### GetPayablesAccountIdOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetPayablesAccountIdOk() (*string, bool)`

GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccountId

`func (o *QuickbooksDesktopBillsIdPostRequest) SetPayablesAccountId(v string)`

SetPayablesAccountId sets PayablesAccountId field to given value.

### HasPayablesAccountId

`func (o *QuickbooksDesktopBillsIdPostRequest) HasPayablesAccountId() bool`

HasPayablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopBillsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopBillsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopBillsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopBillsIdPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopBillsIdPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopBillsIdPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopBillsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopBillsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopBillsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopBillsIdPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopBillsIdPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopBillsIdPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopBillsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopBillsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopBillsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopBillsIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopBillsIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopBillsIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetClearExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) GetClearExpenseLines() bool`

GetClearExpenseLines returns the ClearExpenseLines field if non-nil, zero value otherwise.

### GetClearExpenseLinesOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetClearExpenseLinesOk() (*bool, bool)`

GetClearExpenseLinesOk returns a tuple with the ClearExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) SetClearExpenseLines(v bool)`

SetClearExpenseLines sets ClearExpenseLines field to given value.

### HasClearExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) HasClearExpenseLines() bool`

HasClearExpenseLines returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) GetExpenseLines() []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsIdPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopBillsIdPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetClearItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) GetClearItemLines() bool`

GetClearItemLines returns the ClearItemLines field if non-nil, zero value otherwise.

### GetClearItemLinesOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetClearItemLinesOk() (*bool, bool)`

GetClearItemLinesOk returns a tuple with the ClearItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) SetClearItemLines(v bool)`

SetClearItemLines sets ClearItemLines field to given value.

### HasClearItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) HasClearItemLines() bool`

HasClearItemLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) GetItemLines() []QuickbooksDesktopBillsIdPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) SetItemLines(v []QuickbooksDesktopBillsIdPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopBillsIdPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopBillsIdPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopBillsIdPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopBillsIdPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsIdPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopBillsIdPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


