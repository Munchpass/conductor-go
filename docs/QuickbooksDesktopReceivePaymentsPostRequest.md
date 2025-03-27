# QuickbooksDesktopReceivePaymentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer or customer-job to which the payment for this receive-payment is credited. | 
**ReceivablesAccountId** | Pointer to **string** | The Accounts-Receivable (A/R) account to which this receive-payment is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/R account.  **IMPORTANT**: If this receive-payment is linked to other transactions, this A/R account must match the &#x60;receivablesAccount&#x60; used in all linked transactions. For example, when refunding a credit card payment, the A/R account must match the one used in the original credit transactions being refunded. | [optional] 
**TransactionDate** | **string** | The date of this receive-payment, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this receive-payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**TotalAmount** | **string** | The total monetary amount of this receive-payment, represented as a decimal string.  **NOTE**: The sum of the &#x60;paymentAmount&#x60; amounts in the &#x60;applyToTransactions&#x60; array cannot exceed the &#x60;totalAmount&#x60;, or you will receive an error. | 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this receive-payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**PaymentMethodId** | Pointer to **string** | The receive-payment&#39;s payment method (e.g., cash, check, credit card).  **NOTE**: If this receive-payment contains credit card transaction data supplied from QuickBooks Merchant Services (QBMS) transaction responses, you must specify a credit card payment method (e.g., \&quot;Visa\&quot;, \&quot;MasterCard\&quot;, etc.). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this receive-payment that will be displayed at the beginning of reports containing details about this receive-payment. | [optional] 
**DepositToAccountId** | Pointer to **string** | The account where the funds for this receive-payment will be or have been deposited. If omitted, QuickBooks will use the default Undeposited Funds account. | [optional] 
**CreditCardTransaction** | Pointer to [**QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction**](QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**IsAutoApply** | Pointer to **bool** | When &#x60;true&#x60;, QuickBooks applies &#x60;totalAmount&#x60; to any outstanding transaction that exactly matches &#x60;totalAmount&#x60;. If no exact match is found, this receive-payment is applied to the oldest outstanding transaction for the customer-job. When &#x60;false&#x60;, QuickBooks records the payment but does not apply it to any specific transaction, causing the amount to appear as a credit on the customer-job&#39;s next transaction.  **IMPORTANT**: You must specify either &#x60;isAutoApply&#x60; or &#x60;applyToTransactions&#x60; when creating a receive-payment, but never both. | [optional] [default to false]
**ApplyToTransactions** | Pointer to [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner.md) | The invoices to be paid by this receive-payment. This will create a link between this receive-payment and the specified invoices.  **IMPORTANT**: In each &#x60;applyToTransactions&#x60; object, you must specify either &#x60;paymentAmount&#x60;, &#x60;applyCredits&#x60;, &#x60;discountAmount&#x60;, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target invoice must have &#x60;isPaid&#x3D;false&#x60;, otherwise, QuickBooks will report this object as \&quot;cannot be found\&quot;.  **NOTE**: You must specify either &#x60;isAutoApply&#x60; or &#x60;applyToTransactions&#x60; when creating a receive-payment, but never both. | [optional] 

## Methods

### NewQuickbooksDesktopReceivePaymentsPostRequest

`func NewQuickbooksDesktopReceivePaymentsPostRequest(customerId string, transactionDate string, totalAmount string, ) *QuickbooksDesktopReceivePaymentsPostRequest`

NewQuickbooksDesktopReceivePaymentsPostRequest instantiates a new QuickbooksDesktopReceivePaymentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopReceivePaymentsPostRequestWithDefaults

`func NewQuickbooksDesktopReceivePaymentsPostRequestWithDefaults() *QuickbooksDesktopReceivePaymentsPostRequest`

NewQuickbooksDesktopReceivePaymentsPostRequestWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetReceivablesAccountId() string`

GetReceivablesAccountId returns the ReceivablesAccountId field if non-nil, zero value otherwise.

### GetReceivablesAccountIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetReceivablesAccountIdOk() (*string, bool)`

GetReceivablesAccountIdOk returns a tuple with the ReceivablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetReceivablesAccountId(v string)`

SetReceivablesAccountId sets ReceivablesAccountId field to given value.

### HasReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasReceivablesAccountId() bool`

HasReceivablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTotalAmount

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetDepositToAccountId() string`

GetDepositToAccountId returns the DepositToAccountId field if non-nil, zero value otherwise.

### GetDepositToAccountIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetDepositToAccountIdOk() (*string, bool)`

GetDepositToAccountIdOk returns a tuple with the DepositToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetDepositToAccountId(v string)`

SetDepositToAccountId sets DepositToAccountId field to given value.

### HasDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasDepositToAccountId() bool`

HasDepositToAccountId returns a boolean if a field has been set.

### GetCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetCreditCardTransaction() QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetCreditCardTransactionOk() (*QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetCreditCardTransaction(v QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.

### HasCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasCreditCardTransaction() bool`

HasCreditCardTransaction returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIsAutoApply

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetIsAutoApply() bool`

GetIsAutoApply returns the IsAutoApply field if non-nil, zero value otherwise.

### GetIsAutoApplyOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetIsAutoApplyOk() (*bool, bool)`

GetIsAutoApplyOk returns a tuple with the IsAutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoApply

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetIsAutoApply(v bool)`

SetIsAutoApply sets IsAutoApply field to given value.

### HasIsAutoApply

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasIsAutoApply() bool`

HasIsAutoApply returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.

### HasApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsPostRequest) HasApplyToTransactions() bool`

HasApplyToTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


