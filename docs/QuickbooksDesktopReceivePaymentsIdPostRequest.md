# QuickbooksDesktopReceivePaymentsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the receive-payment object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**CustomerId** | Pointer to **string** | The customer or customer-job to which the payment for this receive-payment is credited. | [optional] 
**ReceivablesAccountId** | Pointer to **string** | The Accounts-Receivable (A/R) account to which this receive-payment is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/R account.  **IMPORTANT**: If this receive-payment is linked to other transactions, this A/R account must match the &#x60;receivablesAccount&#x60; used in all linked transactions. For example, when refunding a credit card payment, the A/R account must match the one used in the original credit transactions being refunded. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this receive-payment, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this receive-payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**TotalAmount** | Pointer to **string** | The total monetary amount of this receive-payment, represented as a decimal string.  **NOTE**: The sum of the &#x60;paymentAmount&#x60; amounts in the &#x60;applyToTransactions&#x60; array cannot exceed the &#x60;totalAmount&#x60;, or you will receive an error. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this receive-payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**PaymentMethodId** | Pointer to **string** | The receive-payment&#39;s payment method (e.g., cash, check, credit card). | [optional] 
**Memo** | Pointer to **string** | A memo or note for this receive-payment that will be displayed at the beginning of reports containing details about this receive-payment. | [optional] 
**DepositToAccountId** | Pointer to **string** | The account where the funds for this receive-payment will be or have been deposited. | [optional] 
**CreditCardTransaction** | Pointer to [**QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction**](QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction.md) |  | [optional] 
**ApplyToTransactions** | Pointer to [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner.md) | The invoices to be paid by this receive-payment. This will create a link between this receive-payment and the specified invoices.  **IMPORTANT**: In each &#x60;applyToTransactions&#x60; object, you must specify either &#x60;paymentAmount&#x60;, &#x60;applyCredits&#x60;, &#x60;discountAmount&#x60;, or any combination of these; if none of these are specified, you will receive an error for an empty transaction.  **IMPORTANT**: The target invoice must have &#x60;isPaid&#x3D;false&#x60;, otherwise, QuickBooks will report this object as \&quot;cannot be found\&quot;. | [optional] 

## Methods

### NewQuickbooksDesktopReceivePaymentsIdPostRequest

`func NewQuickbooksDesktopReceivePaymentsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopReceivePaymentsIdPostRequest`

NewQuickbooksDesktopReceivePaymentsIdPostRequest instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopReceivePaymentsIdPostRequestWithDefaults

`func NewQuickbooksDesktopReceivePaymentsIdPostRequestWithDefaults() *QuickbooksDesktopReceivePaymentsIdPostRequest`

NewQuickbooksDesktopReceivePaymentsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetReceivablesAccountId() string`

GetReceivablesAccountId returns the ReceivablesAccountId field if non-nil, zero value otherwise.

### GetReceivablesAccountIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetReceivablesAccountIdOk() (*string, bool)`

GetReceivablesAccountIdOk returns a tuple with the ReceivablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetReceivablesAccountId(v string)`

SetReceivablesAccountId sets ReceivablesAccountId field to given value.

### HasReceivablesAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasReceivablesAccountId() bool`

HasReceivablesAccountId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetTotalAmount

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetDepositToAccountId() string`

GetDepositToAccountId returns the DepositToAccountId field if non-nil, zero value otherwise.

### GetDepositToAccountIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetDepositToAccountIdOk() (*string, bool)`

GetDepositToAccountIdOk returns a tuple with the DepositToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetDepositToAccountId(v string)`

SetDepositToAccountId sets DepositToAccountId field to given value.

### HasDepositToAccountId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasDepositToAccountId() bool`

HasDepositToAccountId returns a boolean if a field has been set.

### GetCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCreditCardTransaction() QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetCreditCardTransactionOk() (*QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetCreditCardTransaction(v QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.

### HasCreditCardTransaction

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasCreditCardTransaction() bool`

HasCreditCardTransaction returns a boolean if a field has been set.

### GetApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetApplyToTransactions() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

GetApplyToTransactions returns the ApplyToTransactions field if non-nil, zero value otherwise.

### GetApplyToTransactionsOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) GetApplyToTransactionsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner, bool)`

GetApplyToTransactionsOk returns a tuple with the ApplyToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) SetApplyToTransactions(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner)`

SetApplyToTransactions sets ApplyToTransactions field to given value.

### HasApplyToTransactions

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequest) HasApplyToTransactions() bool`

HasApplyToTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


