# QbdItemReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this item receipt. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_item_receipt\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this item receipt was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this item receipt was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this item receipt object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdItemReceiptVendor**](QbdItemReceiptVendor.md) |  | 
**PayablesAccount** | [**QbdItemReceiptPayablesAccount**](QbdItemReceiptPayablesAccount.md) |  | 
**LiabilityAccount** | [**QbdItemReceiptLiabilityAccount**](QbdItemReceiptLiabilityAccount.md) |  | 
**TransactionDate** | **string** | The date of this item receipt, in ISO 8601 format (YYYY-MM-DD). | 
**TotalAmount** | **string** | The total monetary amount of this item receipt, equivalent to the sum of the amounts in &#x60;expenseLines&#x60;, &#x60;itemLines&#x60;, and &#x60;itemGroupLines&#x60;, represented as a decimal string. | 
**Currency** | [**QbdItemReceiptCurrency**](QbdItemReceiptCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this item receipt&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this item receipt converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this item receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Memo** | **string** | A memo or note for this item receipt. | 
**SalesTaxCode** | [**QbdItemReceiptSalesTaxCode**](QbdItemReceiptSalesTaxCode.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The item receipt&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of item receipts to receive this field because it is not returned by default. | 
**ExpenseLines** | [**[]QbdExpenseLine**](QbdExpenseLine.md) | The item receipt&#39;s expense lines, each representing one line in this expense. | 
**ItemLines** | [**[]QbdItemLine**](QbdItemLine.md) | The item receipt&#39;s item lines, each representing the purchase of a specific item or service. | 
**ItemLineGroups** | [**[]QbdItemLineGroup**](QbdItemLineGroup.md) | The item receipt&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the item receipt object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdItemReceipt

`func NewQbdItemReceipt(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdItemReceiptVendor, payablesAccount QbdItemReceiptPayablesAccount, liabilityAccount QbdItemReceiptLiabilityAccount, transactionDate string, totalAmount string, currency QbdItemReceiptCurrency, exchangeRate float32, totalAmountInHomeCurrency string, refNumber string, memo string, salesTaxCode QbdItemReceiptSalesTaxCode, externalId string, linkedTransactions []QbdLinkedTransaction, expenseLines []QbdExpenseLine, itemLines []QbdItemLine, itemLineGroups []QbdItemLineGroup, customFields []QbdCustomField, ) *QbdItemReceipt`

NewQbdItemReceipt instantiates a new QbdItemReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdItemReceiptWithDefaults

`func NewQbdItemReceiptWithDefaults() *QbdItemReceipt`

NewQbdItemReceiptWithDefaults instantiates a new QbdItemReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdItemReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdItemReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdItemReceipt) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdItemReceipt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdItemReceipt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdItemReceipt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdItemReceipt) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdItemReceipt) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdItemReceipt) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdItemReceipt) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdItemReceipt) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdItemReceipt) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdItemReceipt) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdItemReceipt) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdItemReceipt) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdItemReceipt) GetVendor() QbdItemReceiptVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdItemReceipt) GetVendorOk() (*QbdItemReceiptVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdItemReceipt) SetVendor(v QbdItemReceiptVendor)`

SetVendor sets Vendor field to given value.


### GetPayablesAccount

`func (o *QbdItemReceipt) GetPayablesAccount() QbdItemReceiptPayablesAccount`

GetPayablesAccount returns the PayablesAccount field if non-nil, zero value otherwise.

### GetPayablesAccountOk

`func (o *QbdItemReceipt) GetPayablesAccountOk() (*QbdItemReceiptPayablesAccount, bool)`

GetPayablesAccountOk returns a tuple with the PayablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccount

`func (o *QbdItemReceipt) SetPayablesAccount(v QbdItemReceiptPayablesAccount)`

SetPayablesAccount sets PayablesAccount field to given value.


### GetLiabilityAccount

`func (o *QbdItemReceipt) GetLiabilityAccount() QbdItemReceiptLiabilityAccount`

GetLiabilityAccount returns the LiabilityAccount field if non-nil, zero value otherwise.

### GetLiabilityAccountOk

`func (o *QbdItemReceipt) GetLiabilityAccountOk() (*QbdItemReceiptLiabilityAccount, bool)`

GetLiabilityAccountOk returns a tuple with the LiabilityAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityAccount

`func (o *QbdItemReceipt) SetLiabilityAccount(v QbdItemReceiptLiabilityAccount)`

SetLiabilityAccount sets LiabilityAccount field to given value.


### GetTransactionDate

`func (o *QbdItemReceipt) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdItemReceipt) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdItemReceipt) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetTotalAmount

`func (o *QbdItemReceipt) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdItemReceipt) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdItemReceipt) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdItemReceipt) GetCurrency() QbdItemReceiptCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdItemReceipt) GetCurrencyOk() (*QbdItemReceiptCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdItemReceipt) SetCurrency(v QbdItemReceiptCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdItemReceipt) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdItemReceipt) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdItemReceipt) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdItemReceipt) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdItemReceipt) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdItemReceipt) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdItemReceipt) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdItemReceipt) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdItemReceipt) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetMemo

`func (o *QbdItemReceipt) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdItemReceipt) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdItemReceipt) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetSalesTaxCode

`func (o *QbdItemReceipt) GetSalesTaxCode() QbdItemReceiptSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdItemReceipt) GetSalesTaxCodeOk() (*QbdItemReceiptSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdItemReceipt) SetSalesTaxCode(v QbdItemReceiptSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetExternalId

`func (o *QbdItemReceipt) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdItemReceipt) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdItemReceipt) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdItemReceipt) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdItemReceipt) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdItemReceipt) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetExpenseLines

`func (o *QbdItemReceipt) GetExpenseLines() []QbdExpenseLine`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QbdItemReceipt) GetExpenseLinesOk() (*[]QbdExpenseLine, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QbdItemReceipt) SetExpenseLines(v []QbdExpenseLine)`

SetExpenseLines sets ExpenseLines field to given value.


### GetItemLines

`func (o *QbdItemReceipt) GetItemLines() []QbdItemLine`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QbdItemReceipt) GetItemLinesOk() (*[]QbdItemLine, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QbdItemReceipt) SetItemLines(v []QbdItemLine)`

SetItemLines sets ItemLines field to given value.


### GetItemLineGroups

`func (o *QbdItemReceipt) GetItemLineGroups() []QbdItemLineGroup`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QbdItemReceipt) GetItemLineGroupsOk() (*[]QbdItemLineGroup, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QbdItemReceipt) SetItemLineGroups(v []QbdItemLineGroup)`

SetItemLineGroups sets ItemLineGroups field to given value.


### GetCustomFields

`func (o *QbdItemReceipt) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdItemReceipt) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdItemReceipt) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


