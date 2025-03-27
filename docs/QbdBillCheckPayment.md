# QbdBillCheckPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this bill check payment. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_bill_check_payment\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this bill check payment was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this bill check payment was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this bill check payment object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdBillCheckPaymentVendor**](QbdBillCheckPaymentVendor.md) |  | 
**PayablesAccount** | [**QbdBillCheckPaymentPayablesAccount**](QbdBillCheckPaymentPayablesAccount.md) |  | 
**TransactionDate** | **string** | The date of this bill check payment, in ISO 8601 format (YYYY-MM-DD). | 
**BankAccount** | [**QbdBillCheckPaymentBankAccount**](QbdBillCheckPaymentBankAccount.md) |  | 
**Amount** | **string** | The monetary amount of this bill check payment, represented as a decimal string. | 
**Currency** | [**QbdBillCheckPaymentCurrency**](QbdBillCheckPaymentCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this bill check payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountInHomeCurrency** | **string** | The monetary amount of this bill check payment converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this bill check payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.  **IMPORTANT**: For checks, this field is the check number. | 
**Memo** | **string** | A memo or note for this bill check payment. | 
**Address** | [**QbdBillCheckPaymentAddress**](QbdBillCheckPaymentAddress.md) |  | 
**IsQueuedForPrint** | **bool** | Indicates whether this bill check payment is included in the queue of documents for QuickBooks to print. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**AppliedToTransactions** | [**[]QbdReceivableTransaction**](QbdReceivableTransaction.md) | The bill(s) paid by this bill check payment. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the bill check payment object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdBillCheckPayment

`func NewQbdBillCheckPayment(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdBillCheckPaymentVendor, payablesAccount QbdBillCheckPaymentPayablesAccount, transactionDate string, bankAccount QbdBillCheckPaymentBankAccount, amount string, currency QbdBillCheckPaymentCurrency, exchangeRate float32, amountInHomeCurrency string, refNumber string, memo string, address QbdBillCheckPaymentAddress, isQueuedForPrint bool, externalId string, appliedToTransactions []QbdReceivableTransaction, customFields []QbdCustomField, ) *QbdBillCheckPayment`

NewQbdBillCheckPayment instantiates a new QbdBillCheckPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdBillCheckPaymentWithDefaults

`func NewQbdBillCheckPaymentWithDefaults() *QbdBillCheckPayment`

NewQbdBillCheckPaymentWithDefaults instantiates a new QbdBillCheckPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdBillCheckPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdBillCheckPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdBillCheckPayment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdBillCheckPayment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdBillCheckPayment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdBillCheckPayment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdBillCheckPayment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdBillCheckPayment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdBillCheckPayment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdBillCheckPayment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdBillCheckPayment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdBillCheckPayment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdBillCheckPayment) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdBillCheckPayment) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdBillCheckPayment) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdBillCheckPayment) GetVendor() QbdBillCheckPaymentVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdBillCheckPayment) GetVendorOk() (*QbdBillCheckPaymentVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdBillCheckPayment) SetVendor(v QbdBillCheckPaymentVendor)`

SetVendor sets Vendor field to given value.


### GetPayablesAccount

`func (o *QbdBillCheckPayment) GetPayablesAccount() QbdBillCheckPaymentPayablesAccount`

GetPayablesAccount returns the PayablesAccount field if non-nil, zero value otherwise.

### GetPayablesAccountOk

`func (o *QbdBillCheckPayment) GetPayablesAccountOk() (*QbdBillCheckPaymentPayablesAccount, bool)`

GetPayablesAccountOk returns a tuple with the PayablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccount

`func (o *QbdBillCheckPayment) SetPayablesAccount(v QbdBillCheckPaymentPayablesAccount)`

SetPayablesAccount sets PayablesAccount field to given value.


### GetTransactionDate

`func (o *QbdBillCheckPayment) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdBillCheckPayment) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdBillCheckPayment) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetBankAccount

`func (o *QbdBillCheckPayment) GetBankAccount() QbdBillCheckPaymentBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *QbdBillCheckPayment) GetBankAccountOk() (*QbdBillCheckPaymentBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *QbdBillCheckPayment) SetBankAccount(v QbdBillCheckPaymentBankAccount)`

SetBankAccount sets BankAccount field to given value.


### GetAmount

`func (o *QbdBillCheckPayment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdBillCheckPayment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdBillCheckPayment) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *QbdBillCheckPayment) GetCurrency() QbdBillCheckPaymentCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdBillCheckPayment) GetCurrencyOk() (*QbdBillCheckPaymentCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdBillCheckPayment) SetCurrency(v QbdBillCheckPaymentCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdBillCheckPayment) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdBillCheckPayment) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdBillCheckPayment) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountInHomeCurrency

`func (o *QbdBillCheckPayment) GetAmountInHomeCurrency() string`

GetAmountInHomeCurrency returns the AmountInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountInHomeCurrencyOk

`func (o *QbdBillCheckPayment) GetAmountInHomeCurrencyOk() (*string, bool)`

GetAmountInHomeCurrencyOk returns a tuple with the AmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInHomeCurrency

`func (o *QbdBillCheckPayment) SetAmountInHomeCurrency(v string)`

SetAmountInHomeCurrency sets AmountInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdBillCheckPayment) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdBillCheckPayment) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdBillCheckPayment) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetMemo

`func (o *QbdBillCheckPayment) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdBillCheckPayment) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdBillCheckPayment) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetAddress

`func (o *QbdBillCheckPayment) GetAddress() QbdBillCheckPaymentAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdBillCheckPayment) GetAddressOk() (*QbdBillCheckPaymentAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdBillCheckPayment) SetAddress(v QbdBillCheckPaymentAddress)`

SetAddress sets Address field to given value.


### GetIsQueuedForPrint

`func (o *QbdBillCheckPayment) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdBillCheckPayment) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdBillCheckPayment) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetExternalId

`func (o *QbdBillCheckPayment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdBillCheckPayment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdBillCheckPayment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAppliedToTransactions

`func (o *QbdBillCheckPayment) GetAppliedToTransactions() []QbdReceivableTransaction`

GetAppliedToTransactions returns the AppliedToTransactions field if non-nil, zero value otherwise.

### GetAppliedToTransactionsOk

`func (o *QbdBillCheckPayment) GetAppliedToTransactionsOk() (*[]QbdReceivableTransaction, bool)`

GetAppliedToTransactionsOk returns a tuple with the AppliedToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToTransactions

`func (o *QbdBillCheckPayment) SetAppliedToTransactions(v []QbdReceivableTransaction)`

SetAppliedToTransactions sets AppliedToTransactions field to given value.


### GetCustomFields

`func (o *QbdBillCheckPayment) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdBillCheckPayment) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdBillCheckPayment) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


