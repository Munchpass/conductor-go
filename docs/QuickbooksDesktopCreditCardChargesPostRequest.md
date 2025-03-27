# QuickbooksDesktopCreditCardChargesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The bank or credit card account to which money is owed for this credit card charge. | 
**PayeeId** | Pointer to **string** | The vendor or company from whom merchandise or services were purchased for this credit card charge. | [optional] 
**TransactionDate** | **string** | The date of this credit card charge, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this credit card charge, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this credit card charge. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this credit card charge, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the payee. This can be overridden on the credit card charge&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this credit card charge&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInner.md) | The credit card charge&#39;s expense lines, each representing one line in this expense. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLinesInner**](QuickbooksDesktopBillsPostRequestItemLinesInner.md) | The credit card charge&#39;s item lines, each representing the purchase of a specific item or service. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsPostRequestItemLineGroupsInner.md) | The credit card charge&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | [optional] 

## Methods

### NewQuickbooksDesktopCreditCardChargesPostRequest

`func NewQuickbooksDesktopCreditCardChargesPostRequest(accountId string, transactionDate string, ) *QuickbooksDesktopCreditCardChargesPostRequest`

NewQuickbooksDesktopCreditCardChargesPostRequest instantiates a new QuickbooksDesktopCreditCardChargesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCreditCardChargesPostRequestWithDefaults

`func NewQuickbooksDesktopCreditCardChargesPostRequestWithDefaults() *QuickbooksDesktopCreditCardChargesPostRequest`

NewQuickbooksDesktopCreditCardChargesPostRequestWithDefaults instantiates a new QuickbooksDesktopCreditCardChargesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPayeeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopCreditCardChargesPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


