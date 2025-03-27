# QbdReceivePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this receive-payment. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_receive_payment\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this receive-payment was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this receive-payment was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this receive-payment object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdReceivePaymentCustomer**](QbdReceivePaymentCustomer.md) |  | 
**ReceivablesAccount** | [**QbdReceivePaymentReceivablesAccount**](QbdReceivePaymentReceivablesAccount.md) |  | 
**TransactionDate** | **string** | The date of this receive-payment, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this receive-payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**TotalAmount** | **string** | The total monetary amount of this receive-payment, represented as a decimal string. | 
**Currency** | [**QbdReceivePaymentCurrency**](QbdReceivePaymentCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this receive-payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this receive-payment converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**PaymentMethod** | [**QbdReceivePaymentPaymentMethod**](QbdReceivePaymentPaymentMethod.md) |  | 
**Memo** | **string** | A memo or note for this receive-payment that will be displayed at the beginning of reports containing details about this receive-payment. | 
**DepositToAccount** | [**QbdReceivePaymentDepositToAccount**](QbdReceivePaymentDepositToAccount.md) |  | 
**CreditCardTransaction** | [**QbdReceivePaymentCreditCardTransaction**](QbdReceivePaymentCreditCardTransaction.md) |  | 
**UnusedPayment** | **string** | The amount of this receive-payment that remains unapplied to any transactions. This occurs in two cases: (1) When the sum of &#x60;paymentAmount&#x60; amounts in &#x60;applyToTransactions&#x60; is less than &#x60;totalAmount&#x60;, leaving a portion of the payment unused, or (2) When a payment is received that equals the exact amount of an invoice, but credits or discounts are also applied, resulting in excess payment. | 
**UnusedCredits** | **string** | The amount of credit that remains unused after applying credits to this receive-payment. This occurs when the &#x60;applyCredit.appliedAmount&#x60; specified for a credit memo (&#x60;applyCredit.creditMemoId&#x60;) in the &#x60;applyToTransactions&#x60; array is less than the total available credit amount for that credit memo. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**AppliedToTransactions** | [**[]QbdReceivableTransaction**](QbdReceivableTransaction.md) | The invoice(s) paid by this receive-payment. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the receive-payment object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdReceivePayment

`func NewQbdReceivePayment(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdReceivePaymentCustomer, receivablesAccount QbdReceivePaymentReceivablesAccount, transactionDate string, refNumber string, totalAmount string, currency QbdReceivePaymentCurrency, exchangeRate float32, totalAmountInHomeCurrency string, paymentMethod QbdReceivePaymentPaymentMethod, memo string, depositToAccount QbdReceivePaymentDepositToAccount, creditCardTransaction QbdReceivePaymentCreditCardTransaction, unusedPayment string, unusedCredits string, externalId string, appliedToTransactions []QbdReceivableTransaction, customFields []QbdCustomField, ) *QbdReceivePayment`

NewQbdReceivePayment instantiates a new QbdReceivePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdReceivePaymentWithDefaults

`func NewQbdReceivePaymentWithDefaults() *QbdReceivePayment`

NewQbdReceivePaymentWithDefaults instantiates a new QbdReceivePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdReceivePayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdReceivePayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdReceivePayment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdReceivePayment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdReceivePayment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdReceivePayment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdReceivePayment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdReceivePayment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdReceivePayment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdReceivePayment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdReceivePayment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdReceivePayment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdReceivePayment) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdReceivePayment) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdReceivePayment) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdReceivePayment) GetCustomer() QbdReceivePaymentCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdReceivePayment) GetCustomerOk() (*QbdReceivePaymentCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdReceivePayment) SetCustomer(v QbdReceivePaymentCustomer)`

SetCustomer sets Customer field to given value.


### GetReceivablesAccount

`func (o *QbdReceivePayment) GetReceivablesAccount() QbdReceivePaymentReceivablesAccount`

GetReceivablesAccount returns the ReceivablesAccount field if non-nil, zero value otherwise.

### GetReceivablesAccountOk

`func (o *QbdReceivePayment) GetReceivablesAccountOk() (*QbdReceivePaymentReceivablesAccount, bool)`

GetReceivablesAccountOk returns a tuple with the ReceivablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccount

`func (o *QbdReceivePayment) SetReceivablesAccount(v QbdReceivePaymentReceivablesAccount)`

SetReceivablesAccount sets ReceivablesAccount field to given value.


### GetTransactionDate

`func (o *QbdReceivePayment) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdReceivePayment) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdReceivePayment) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdReceivePayment) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdReceivePayment) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdReceivePayment) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetTotalAmount

`func (o *QbdReceivePayment) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdReceivePayment) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdReceivePayment) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdReceivePayment) GetCurrency() QbdReceivePaymentCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdReceivePayment) GetCurrencyOk() (*QbdReceivePaymentCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdReceivePayment) SetCurrency(v QbdReceivePaymentCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdReceivePayment) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdReceivePayment) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdReceivePayment) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdReceivePayment) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdReceivePayment) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdReceivePayment) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetPaymentMethod

`func (o *QbdReceivePayment) GetPaymentMethod() QbdReceivePaymentPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *QbdReceivePayment) GetPaymentMethodOk() (*QbdReceivePaymentPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *QbdReceivePayment) SetPaymentMethod(v QbdReceivePaymentPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetMemo

`func (o *QbdReceivePayment) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdReceivePayment) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdReceivePayment) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetDepositToAccount

`func (o *QbdReceivePayment) GetDepositToAccount() QbdReceivePaymentDepositToAccount`

GetDepositToAccount returns the DepositToAccount field if non-nil, zero value otherwise.

### GetDepositToAccountOk

`func (o *QbdReceivePayment) GetDepositToAccountOk() (*QbdReceivePaymentDepositToAccount, bool)`

GetDepositToAccountOk returns a tuple with the DepositToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccount

`func (o *QbdReceivePayment) SetDepositToAccount(v QbdReceivePaymentDepositToAccount)`

SetDepositToAccount sets DepositToAccount field to given value.


### GetCreditCardTransaction

`func (o *QbdReceivePayment) GetCreditCardTransaction() QbdReceivePaymentCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QbdReceivePayment) GetCreditCardTransactionOk() (*QbdReceivePaymentCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QbdReceivePayment) SetCreditCardTransaction(v QbdReceivePaymentCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.


### GetUnusedPayment

`func (o *QbdReceivePayment) GetUnusedPayment() string`

GetUnusedPayment returns the UnusedPayment field if non-nil, zero value otherwise.

### GetUnusedPaymentOk

`func (o *QbdReceivePayment) GetUnusedPaymentOk() (*string, bool)`

GetUnusedPaymentOk returns a tuple with the UnusedPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedPayment

`func (o *QbdReceivePayment) SetUnusedPayment(v string)`

SetUnusedPayment sets UnusedPayment field to given value.


### GetUnusedCredits

`func (o *QbdReceivePayment) GetUnusedCredits() string`

GetUnusedCredits returns the UnusedCredits field if non-nil, zero value otherwise.

### GetUnusedCreditsOk

`func (o *QbdReceivePayment) GetUnusedCreditsOk() (*string, bool)`

GetUnusedCreditsOk returns a tuple with the UnusedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedCredits

`func (o *QbdReceivePayment) SetUnusedCredits(v string)`

SetUnusedCredits sets UnusedCredits field to given value.


### GetExternalId

`func (o *QbdReceivePayment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdReceivePayment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdReceivePayment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAppliedToTransactions

`func (o *QbdReceivePayment) GetAppliedToTransactions() []QbdReceivableTransaction`

GetAppliedToTransactions returns the AppliedToTransactions field if non-nil, zero value otherwise.

### GetAppliedToTransactionsOk

`func (o *QbdReceivePayment) GetAppliedToTransactionsOk() (*[]QbdReceivableTransaction, bool)`

GetAppliedToTransactionsOk returns a tuple with the AppliedToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToTransactions

`func (o *QbdReceivePayment) SetAppliedToTransactions(v []QbdReceivableTransaction)`

SetAppliedToTransactions sets AppliedToTransactions field to given value.


### GetCustomFields

`func (o *QbdReceivePayment) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdReceivePayment) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdReceivePayment) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


