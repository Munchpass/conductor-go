# QuickbooksDesktopCreditCardCreditsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The bank or credit card account to which this credit card credit is posted. | 
**PayeeId** | Pointer to **string** | The vendor or company from whom this credit card credit was received for purchased merchandise or services. | [optional] 
**TransactionDate** | **string** | The date of this credit card credit, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this credit card credit, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this credit card credit. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this credit card credit, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the payee. This can be overridden on the credit card credit&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this credit card credit&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**ExpenseLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInner.md) | The credit card credit&#39;s expense lines, each representing one line in this expense. | [optional] 
**ItemLines** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLinesInner**](QuickbooksDesktopBillsPostRequestItemLinesInner.md) | The credit card credit&#39;s item lines, each representing the purchase of a specific item or service. | [optional] 
**ItemLineGroups** | Pointer to [**[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner**](QuickbooksDesktopBillsPostRequestItemLineGroupsInner.md) | The credit card credit&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | [optional] 

## Methods

### NewQuickbooksDesktopCreditCardCreditsPostRequest

`func NewQuickbooksDesktopCreditCardCreditsPostRequest(accountId string, transactionDate string, ) *QuickbooksDesktopCreditCardCreditsPostRequest`

NewQuickbooksDesktopCreditCardCreditsPostRequest instantiates a new QuickbooksDesktopCreditCardCreditsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCreditCardCreditsPostRequestWithDefaults

`func NewQuickbooksDesktopCreditCardCreditsPostRequestWithDefaults() *QuickbooksDesktopCreditCardCreditsPostRequest`

NewQuickbooksDesktopCreditCardCreditsPostRequestWithDefaults instantiates a new QuickbooksDesktopCreditCardCreditsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPayeeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExpenseLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExpenseLines() []QuickbooksDesktopBillsPostRequestExpenseLinesInner`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetExpenseLinesOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInner, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetExpenseLines(v []QuickbooksDesktopBillsPostRequestExpenseLinesInner)`

SetExpenseLines sets ExpenseLines field to given value.

### HasExpenseLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasExpenseLines() bool`

HasExpenseLines returns a boolean if a field has been set.

### GetItemLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetItemLines() []QuickbooksDesktopBillsPostRequestItemLinesInner`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetItemLinesOk() (*[]QuickbooksDesktopBillsPostRequestItemLinesInner, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetItemLines(v []QuickbooksDesktopBillsPostRequestItemLinesInner)`

SetItemLines sets ItemLines field to given value.

### HasItemLines

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasItemLines() bool`

HasItemLines returns a boolean if a field has been set.

### GetItemLineGroups

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetItemLineGroups() []QuickbooksDesktopBillsPostRequestItemLineGroupsInner`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) GetItemLineGroupsOk() (*[]QuickbooksDesktopBillsPostRequestItemLineGroupsInner, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) SetItemLineGroups(v []QuickbooksDesktopBillsPostRequestItemLineGroupsInner)`

SetItemLineGroups sets ItemLineGroups field to given value.

### HasItemLineGroups

`func (o *QuickbooksDesktopCreditCardCreditsPostRequest) HasItemLineGroups() bool`

HasItemLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


