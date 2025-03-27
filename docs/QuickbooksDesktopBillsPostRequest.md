# QuickbooksDesktopBillsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | The vendor who sent this bill for goods or services purchased. | 
**VendorAddress** | Pointer to [**QuickbooksDesktopBillsPostRequestVendorAddress**](QuickbooksDesktopBillsPostRequestVendorAddress.md) |  | [optional] 
**PayablesAccountId** | Pointer to **string** | The Accounts-Payable (A/P) account to which this bill is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this bill is linked to other transactions, this A/P account must match the &#x60;payablesAccount&#x60; used in those other transactions. | [optional] 
**TransactionDate** | **string** | The date of this bill, in ISO 8601 format (YYYY-MM-DD). | 
**DueDate** | Pointer to **string** | The date by which this bill must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this bill, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**TermsId** | Pointer to **string** | The bill&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this bill that appears in the Accounts-Payable register and in reports that include this bill. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this bill, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the bill&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this bill&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**LinkToTransactionIds** | Pointer to **[]string** | IDs of existing purchase orders that you wish to link to this bill. Note that this links entire transactions, not individual transaction lines. If you want to link individual lines in a transaction, instead use the field &#x60;linkToTransactionLine&#x60; on this bill&#39;s lines, if available.  Transactions can only be linked when creating this bill and cannot be unlinked later.  You can use both &#x60;linkToTransactionIds&#x60; (on this bill) and &#x60;linkToTransactionLine&#x60; (on its transaction lines) as long as they do NOT link to the same transaction (otherwise, QuickBooks will return an error). QuickBooks will also return an error if you attempt to link a transaction that is empty or already closed.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the bill and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of bills, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInner.md) | The bill&#39;s expense lines, each representing one line in this expense. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLinesInner**](QuickbooksDesktopBillsPostRequestItemLinesInner.md) | The bill&#39;s item lines, each representing the purchase of a specific item or service. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsPostRequestItemLineGroupsInner.md) | The bill&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | [optional] 

## Methods

### NewQuickbooksDesktopBillsPostRequest

`func NewQuickbooksDesktopBillsPostRequest(vendorId string, transactionDate string, ) *QuickbooksDesktopBillsPostRequest`

NewQuickbooksDesktopBillsPostRequest instantiates a new QuickbooksDesktopBillsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsPostRequestWithDefaults

`func NewQuickbooksDesktopBillsPostRequestWithDefaults() *QuickbooksDesktopBillsPostRequest`

NewQuickbooksDesktopBillsPostRequestWithDefaults instantiates a new QuickbooksDesktopBillsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *QuickbooksDesktopBillsPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopBillsPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopBillsPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetVendorAddress

`func (o *QuickbooksDesktopBillsPostRequest) GetVendorAddress() QuickbooksDesktopBillsPostRequestVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QuickbooksDesktopBillsPostRequest) GetVendorAddressOk() (*QuickbooksDesktopBillsPostRequestVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QuickbooksDesktopBillsPostRequest) SetVendorAddress(v QuickbooksDesktopBillsPostRequestVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.

### HasVendorAddress

`func (o *QuickbooksDesktopBillsPostRequest) HasVendorAddress() bool`

HasVendorAddress returns a boolean if a field has been set.

### GetPayablesAccountId

`func (o *QuickbooksDesktopBillsPostRequest) GetPayablesAccountId() string`

GetPayablesAccountId returns the PayablesAccountId field if non-nil, zero value otherwise.

### GetPayablesAccountIdOk

`func (o *QuickbooksDesktopBillsPostRequest) GetPayablesAccountIdOk() (*string, bool)`

GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccountId

`func (o *QuickbooksDesktopBillsPostRequest) SetPayablesAccountId(v string)`

SetPayablesAccountId sets PayablesAccountId field to given value.

### HasPayablesAccountId

`func (o *QuickbooksDesktopBillsPostRequest) HasPayablesAccountId() bool`

HasPayablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopBillsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopBillsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopBillsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetDueDate

`func (o *QuickbooksDesktopBillsPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopBillsPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopBillsPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopBillsPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopBillsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopBillsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopBillsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopBillsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopBillsPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopBillsPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopBillsPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopBillsPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopBillsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopBillsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopBillsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopBillsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopBillsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopBillsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopBillsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopBillsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopBillsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopBillsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopBillsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopBillsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopBillsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLinkToTransactionIds

`func (o *QuickbooksDesktopBillsPostRequest) GetLinkToTransactionIds() []string`

GetLinkToTransactionIds returns the LinkToTransactionIds field if non-nil, zero value otherwise.

### GetLinkToTransactionIdsOk

`func (o *QuickbooksDesktopBillsPostRequest) GetLinkToTransactionIdsOk() (*[]string, bool)`

GetLinkToTransactionIdsOk returns a tuple with the LinkToTransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToTransactionIds

`func (o *QuickbooksDesktopBillsPostRequest) SetLinkToTransactionIds(v []string)`

SetLinkToTransactionIds sets LinkToTransactionIds field to given value.

### HasLinkToTransactionIds

`func (o *QuickbooksDesktopBillsPostRequest) HasLinkToTransactionIds() bool`

HasLinkToTransactionIds returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopBillsPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopBillsPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopBillsPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopBillsPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopBillsPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopBillsPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopBillsPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopBillsPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopBillsPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopBillsPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopBillsPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopBillsPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


