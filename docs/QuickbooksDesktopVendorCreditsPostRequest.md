# QuickbooksDesktopVendorCreditsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | The vendor who sent this vendor credit for goods or services purchased. | 
**PayablesAccountId** | Pointer to **string** | The Accounts-Payable (A/P) account to which this vendor credit is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this vendor credit is linked to other transactions, this A/P account must match the &#x60;payablesAccount&#x60; used in those other transactions. | [optional] 
**TransactionDate** | **string** | The date of this vendor credit, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this vendor credit, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this vendor credit. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this vendor credit, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the vendor credit&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this vendor credit&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInner.md) | The vendor credit&#39;s expense lines, each representing one line in this expense. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLinesInner**](QuickbooksDesktopBillsPostRequestItemLinesInner.md) | The vendor credit&#39;s item lines, each representing the purchase of a specific item or service. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsPostRequestItemLineGroupsInner.md) | The vendor credit&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | [optional] 

## Methods

### NewQuickbooksDesktopVendorCreditsPostRequest

`func NewQuickbooksDesktopVendorCreditsPostRequest(vendorId string, transactionDate string, ) *QuickbooksDesktopVendorCreditsPostRequest`

NewQuickbooksDesktopVendorCreditsPostRequest instantiates a new QuickbooksDesktopVendorCreditsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopVendorCreditsPostRequestWithDefaults

`func NewQuickbooksDesktopVendorCreditsPostRequestWithDefaults() *QuickbooksDesktopVendorCreditsPostRequest`

NewQuickbooksDesktopVendorCreditsPostRequestWithDefaults instantiates a new QuickbooksDesktopVendorCreditsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetPayablesAccountId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetPayablesAccountId() string`

GetPayablesAccountId returns the PayablesAccountId field if non-nil, zero value otherwise.

### GetPayablesAccountIdOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetPayablesAccountIdOk() (*string, bool)`

GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccountId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetPayablesAccountId(v string)`

SetPayablesAccountId sets PayablesAccountId field to given value.

### HasPayablesAccountId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasPayablesAccountId() bool`

HasPayablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopVendorCreditsPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopVendorCreditsPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopVendorCreditsPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


