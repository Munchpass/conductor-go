# QbdBillCreditCardPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this bill credit card payment. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_bill_credit_card_payment\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this bill credit card payment was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this bill credit card payment was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this bill credit card payment object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdBillCreditCardPaymentVendor**](QbdBillCreditCardPaymentVendor.md) |  | 
**PayablesAccount** | [**QbdBillCreditCardPaymentPayablesAccount**](QbdBillCreditCardPaymentPayablesAccount.md) |  | 
**TransactionDate** | **string** | The date of this bill credit card payment, in ISO 8601 format (YYYY-MM-DD). | 
**CreditCardAccount** | [**QbdBillCreditCardPaymentCreditCardAccount**](QbdBillCreditCardPaymentCreditCardAccount.md) |  | 
**Amount** | **string** | The monetary amount of this bill credit card payment, represented as a decimal string. | 
**Currency** | [**QbdBillCreditCardPaymentCurrency**](QbdBillCreditCardPaymentCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this bill credit card payment&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountInHomeCurrency** | **string** | The monetary amount of this bill credit card payment converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this bill credit card payment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Memo** | **string** | A memo or note for this bill credit card payment. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**AppliedToTransactions** | [**[]QbdReceivableTransaction**](QbdReceivableTransaction.md) | The bill(s) paid by this bill credit card payment. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the bill credit card payment object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdBillCreditCardPayment

`func NewQbdBillCreditCardPayment(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdBillCreditCardPaymentVendor, payablesAccount QbdBillCreditCardPaymentPayablesAccount, transactionDate string, creditCardAccount QbdBillCreditCardPaymentCreditCardAccount, amount string, currency QbdBillCreditCardPaymentCurrency, exchangeRate float32, amountInHomeCurrency string, refNumber string, memo string, externalId string, appliedToTransactions []QbdReceivableTransaction, customFields []QbdCustomField, ) *QbdBillCreditCardPayment`

NewQbdBillCreditCardPayment instantiates a new QbdBillCreditCardPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdBillCreditCardPaymentWithDefaults

`func NewQbdBillCreditCardPaymentWithDefaults() *QbdBillCreditCardPayment`

NewQbdBillCreditCardPaymentWithDefaults instantiates a new QbdBillCreditCardPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdBillCreditCardPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdBillCreditCardPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdBillCreditCardPayment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdBillCreditCardPayment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdBillCreditCardPayment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdBillCreditCardPayment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdBillCreditCardPayment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdBillCreditCardPayment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdBillCreditCardPayment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdBillCreditCardPayment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdBillCreditCardPayment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdBillCreditCardPayment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdBillCreditCardPayment) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdBillCreditCardPayment) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdBillCreditCardPayment) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdBillCreditCardPayment) GetVendor() QbdBillCreditCardPaymentVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdBillCreditCardPayment) GetVendorOk() (*QbdBillCreditCardPaymentVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdBillCreditCardPayment) SetVendor(v QbdBillCreditCardPaymentVendor)`

SetVendor sets Vendor field to given value.


### GetPayablesAccount

`func (o *QbdBillCreditCardPayment) GetPayablesAccount() QbdBillCreditCardPaymentPayablesAccount`

GetPayablesAccount returns the PayablesAccount field if non-nil, zero value otherwise.

### GetPayablesAccountOk

`func (o *QbdBillCreditCardPayment) GetPayablesAccountOk() (*QbdBillCreditCardPaymentPayablesAccount, bool)`

GetPayablesAccountOk returns a tuple with the PayablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccount

`func (o *QbdBillCreditCardPayment) SetPayablesAccount(v QbdBillCreditCardPaymentPayablesAccount)`

SetPayablesAccount sets PayablesAccount field to given value.


### GetTransactionDate

`func (o *QbdBillCreditCardPayment) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdBillCreditCardPayment) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdBillCreditCardPayment) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetCreditCardAccount

`func (o *QbdBillCreditCardPayment) GetCreditCardAccount() QbdBillCreditCardPaymentCreditCardAccount`

GetCreditCardAccount returns the CreditCardAccount field if non-nil, zero value otherwise.

### GetCreditCardAccountOk

`func (o *QbdBillCreditCardPayment) GetCreditCardAccountOk() (*QbdBillCreditCardPaymentCreditCardAccount, bool)`

GetCreditCardAccountOk returns a tuple with the CreditCardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardAccount

`func (o *QbdBillCreditCardPayment) SetCreditCardAccount(v QbdBillCreditCardPaymentCreditCardAccount)`

SetCreditCardAccount sets CreditCardAccount field to given value.


### GetAmount

`func (o *QbdBillCreditCardPayment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdBillCreditCardPayment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdBillCreditCardPayment) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *QbdBillCreditCardPayment) GetCurrency() QbdBillCreditCardPaymentCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdBillCreditCardPayment) GetCurrencyOk() (*QbdBillCreditCardPaymentCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdBillCreditCardPayment) SetCurrency(v QbdBillCreditCardPaymentCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdBillCreditCardPayment) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdBillCreditCardPayment) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdBillCreditCardPayment) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountInHomeCurrency

`func (o *QbdBillCreditCardPayment) GetAmountInHomeCurrency() string`

GetAmountInHomeCurrency returns the AmountInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountInHomeCurrencyOk

`func (o *QbdBillCreditCardPayment) GetAmountInHomeCurrencyOk() (*string, bool)`

GetAmountInHomeCurrencyOk returns a tuple with the AmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInHomeCurrency

`func (o *QbdBillCreditCardPayment) SetAmountInHomeCurrency(v string)`

SetAmountInHomeCurrency sets AmountInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdBillCreditCardPayment) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdBillCreditCardPayment) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdBillCreditCardPayment) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetMemo

`func (o *QbdBillCreditCardPayment) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdBillCreditCardPayment) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdBillCreditCardPayment) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetExternalId

`func (o *QbdBillCreditCardPayment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdBillCreditCardPayment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdBillCreditCardPayment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAppliedToTransactions

`func (o *QbdBillCreditCardPayment) GetAppliedToTransactions() []QbdReceivableTransaction`

GetAppliedToTransactions returns the AppliedToTransactions field if non-nil, zero value otherwise.

### GetAppliedToTransactionsOk

`func (o *QbdBillCreditCardPayment) GetAppliedToTransactionsOk() (*[]QbdReceivableTransaction, bool)`

GetAppliedToTransactionsOk returns a tuple with the AppliedToTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToTransactions

`func (o *QbdBillCreditCardPayment) SetAppliedToTransactions(v []QbdReceivableTransaction)`

SetAppliedToTransactions sets AppliedToTransactions field to given value.


### GetCustomFields

`func (o *QbdBillCreditCardPayment) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdBillCreditCardPayment) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdBillCreditCardPayment) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


