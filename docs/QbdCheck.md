# QbdCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this check. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_check\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this check was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this check was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this check object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**BankAccount** | [**QbdCheckBankAccount**](QbdCheckBankAccount.md) |  | 
**Payee** | [**QbdCheckPayee**](QbdCheckPayee.md) |  | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this check, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.  **IMPORTANT**: For checks, this field is the check number. | 
**TransactionDate** | **string** | The date written on this check, in ISO 8601 format (YYYY-MM-DD). | 
**Amount** | **string** | The total monetary amount of this check, represented as a decimal string. This equals the sum of the amounts in the check&#39;s expense lines, item lines, and item group lines. | 
**Currency** | [**QbdCheckCurrency**](QbdCheckCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this check&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountInHomeCurrency** | **string** | The monetary amount of this check converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | The memo that is printed on this check. | 
**Address** | [**QbdCheckAddress**](QbdCheckAddress.md) |  | 
**IsPending** | **bool** | Indicates whether this check has not been completed. | 
**IsQueuedForPrint** | **bool** | Indicates whether this check is included in the queue of documents for QuickBooks to print. | 
**SalesTaxCode** | [**QbdCheckSalesTaxCode**](QbdCheckSalesTaxCode.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The check&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of checks to receive this field because it is not returned by default. | 
**ExpenseLines** | [**[]QbdExpenseLine**](QbdExpenseLine.md) | The check&#39;s expense lines, each representing one line in this expense. | 
**ItemLines** | [**[]QbdItemLine**](QbdItemLine.md) | The check&#39;s item lines, each representing the purchase of a specific item or service. | 
**ItemLineGroups** | [**[]QbdItemLineGroup**](QbdItemLineGroup.md) | The check&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the check object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdCheck

`func NewQbdCheck(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, bankAccount QbdCheckBankAccount, payee QbdCheckPayee, refNumber string, transactionDate string, amount string, currency QbdCheckCurrency, exchangeRate float32, amountInHomeCurrency string, memo string, address QbdCheckAddress, isPending bool, isQueuedForPrint bool, salesTaxCode QbdCheckSalesTaxCode, externalId string, linkedTransactions []QbdLinkedTransaction, expenseLines []QbdExpenseLine, itemLines []QbdItemLine, itemLineGroups []QbdItemLineGroup, customFields []QbdCustomField, ) *QbdCheck`

NewQbdCheck instantiates a new QbdCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCheckWithDefaults

`func NewQbdCheckWithDefaults() *QbdCheck`

NewQbdCheckWithDefaults instantiates a new QbdCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCheck) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdCheck) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdCheck) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdCheck) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdCheck) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdCheck) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdCheck) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdCheck) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdCheck) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdCheck) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetBankAccount

`func (o *QbdCheck) GetBankAccount() QbdCheckBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *QbdCheck) GetBankAccountOk() (*QbdCheckBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *QbdCheck) SetBankAccount(v QbdCheckBankAccount)`

SetBankAccount sets BankAccount field to given value.


### GetPayee

`func (o *QbdCheck) GetPayee() QbdCheckPayee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *QbdCheck) GetPayeeOk() (*QbdCheckPayee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *QbdCheck) SetPayee(v QbdCheckPayee)`

SetPayee sets Payee field to given value.


### GetRefNumber

`func (o *QbdCheck) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdCheck) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdCheck) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetTransactionDate

`func (o *QbdCheck) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdCheck) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdCheck) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetAmount

`func (o *QbdCheck) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdCheck) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdCheck) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *QbdCheck) GetCurrency() QbdCheckCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdCheck) GetCurrencyOk() (*QbdCheckCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdCheck) SetCurrency(v QbdCheckCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdCheck) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdCheck) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdCheck) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountInHomeCurrency

`func (o *QbdCheck) GetAmountInHomeCurrency() string`

GetAmountInHomeCurrency returns the AmountInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountInHomeCurrencyOk

`func (o *QbdCheck) GetAmountInHomeCurrencyOk() (*string, bool)`

GetAmountInHomeCurrencyOk returns a tuple with the AmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInHomeCurrency

`func (o *QbdCheck) SetAmountInHomeCurrency(v string)`

SetAmountInHomeCurrency sets AmountInHomeCurrency field to given value.


### GetMemo

`func (o *QbdCheck) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdCheck) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdCheck) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetAddress

`func (o *QbdCheck) GetAddress() QbdCheckAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdCheck) GetAddressOk() (*QbdCheckAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdCheck) SetAddress(v QbdCheckAddress)`

SetAddress sets Address field to given value.


### GetIsPending

`func (o *QbdCheck) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QbdCheck) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QbdCheck) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetIsQueuedForPrint

`func (o *QbdCheck) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdCheck) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdCheck) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetSalesTaxCode

`func (o *QbdCheck) GetSalesTaxCode() QbdCheckSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdCheck) GetSalesTaxCodeOk() (*QbdCheckSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdCheck) SetSalesTaxCode(v QbdCheckSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetExternalId

`func (o *QbdCheck) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdCheck) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdCheck) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdCheck) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdCheck) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdCheck) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetExpenseLines

`func (o *QbdCheck) GetExpenseLines() []QbdExpenseLine`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QbdCheck) GetExpenseLinesOk() (*[]QbdExpenseLine, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QbdCheck) SetExpenseLines(v []QbdExpenseLine)`

SetExpenseLines sets ExpenseLines field to given value.


### GetItemLines

`func (o *QbdCheck) GetItemLines() []QbdItemLine`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QbdCheck) GetItemLinesOk() (*[]QbdItemLine, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QbdCheck) SetItemLines(v []QbdItemLine)`

SetItemLines sets ItemLines field to given value.


### GetItemLineGroups

`func (o *QbdCheck) GetItemLineGroups() []QbdItemLineGroup`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QbdCheck) GetItemLineGroupsOk() (*[]QbdItemLineGroup, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QbdCheck) SetItemLineGroups(v []QbdItemLineGroup)`

SetItemLineGroups sets ItemLineGroups field to given value.


### GetCustomFields

`func (o *QbdCheck) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCheck) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCheck) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


