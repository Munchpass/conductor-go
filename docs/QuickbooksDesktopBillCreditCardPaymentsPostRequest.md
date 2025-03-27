# QuickbooksDesktopBillCreditCardPaymentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | The vendor who sent the bill(s) that this bill credit card payment is paying and who will receive this payment.  **IMPORTANT**: This vendor must match the &#x60;vendor&#x60; on the bill(s) specified in &#x60;applyToTransactions&#x60;; otherwise, QuickBooks will say the &#x60;transactionId&#x60; in &#x60;applyToTransactions&#x60; \&quot;does not exist\&quot;. | 
**PayablesAccountId** | Pointer to **string** | The Accounts-Payable (A/P) account to which this bill credit card payment is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/P account.  **IMPORTANT**: If this bill credit card payment is linked to other transactions, this A/P account must match the &#x60;payablesAccount&#x60; used in those other transactions. | [optional] 
**TransactionDate** | **string** | The date of this bill credit card payment, in ISO 8601 format (YYYY-MM-DD). | 
**CreditCardAccountId** | **string** | The credit card account to which this bill credit card payment is being charged. This bill credit card payment will decrease the balance of this account. | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this bill credit card payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this bill credit card payment. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this bill credit card payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**ApplyToTransactions** | [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner.md) | The bills to be paid by this bill credit card payment. This will create a link between this bill credit card payment and the specified bills.  **IMPORTANT**: In each &#x60;applyToTransactions&#x60; object, you must specify either &#x60;paymentAmount&#x60;, &#x60;applyCredits&#x60;, &#x60;discountAmount&#x60;, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target bill must have &#x60;isPaid&#x3D;false&#x60;, otherwise, QuickBooks will report this object as \&quot;cannot be found\&quot;. | 

## Methods

### NewQuickbooksDesktopBillCreditCardPaymentsPostRequest

`func NewQuickbooksDesktopBillCreditCardPaymentsPostRequest(vendorId string, transactionDate string, creditCardAccountId string, applyToTransactions []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, ) *QuickbooksDesktopBillCreditCardPaymentsPostRequest`

NewQuickbooksDesktopBillCreditCardPaymentsPostRequest instantiates a new QuickbooksDesktopBillCreditCardPaymentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillCreditCardPaymentsPostRequestWithDefaults

`func NewQuickbooksDesktopBillCreditCardPaymentsPostRequestWithDefaults() *QuickbooksDesktopBillCreditCardPaymentsPostRequest`

NewQuickbooksDesktopBillCreditCardPaymentsPostRequestWithDefaults instantiates a new QuickbooksDesktopBillCreditCardPaymentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetPayablesAccountId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetPayablesAccountId() string`

GetPayablesAccountId returns the PayablesAccountId field if non-nil, zero value otherwise.

### GetPayablesAccountIdOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetPayablesAccountIdOk() (*string, bool)`

GetPayablesAccountIdOk returns a tuple with the PayablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccountId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetPayablesAccountId(v string)`

SetPayablesAccountId sets PayablesAccountId field to given value.

### HasPayablesAccountId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) HasPayablesAccountId() bool`

HasPayablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetCreditCardAccountId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetCreditCardAccountId() string`

GetCreditCardAccountId returns the CreditCardAccountId field if non-nil, zero value otherwise.

### GetCreditCardAccountIdOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetCreditCardAccountIdOk() (*string, bool)`

GetCreditCardAccountIdOk returns a tuple with the CreditCardAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardAccountId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetCreditCardAccountId(v string)`

SetCreditCardAccountId sets CreditCardAccountId field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopBillCreditCardPaymentsPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


