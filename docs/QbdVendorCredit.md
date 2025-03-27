# QbdVendorCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this vendor credit. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_vendor_credit\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this vendor credit was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this vendor credit was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this vendor credit object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdVendorCreditVendor**](QbdVendorCreditVendor.md) |  | 
**PayablesAccount** | [**QbdVendorCreditPayablesAccount**](QbdVendorCreditPayablesAccount.md) |  | 
**TransactionDate** | **string** | The date of this vendor credit, in ISO 8601 format (YYYY-MM-DD). | 
**CreditAmount** | **string** | The monetary amount of the vendor credit, represented as a decimal string. When applied to a vendor bill, this amount reduces the outstanding balance owed to the vendor. | 
**Currency** | [**QbdVendorCreditCurrency**](QbdVendorCreditCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this vendor credit&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**CreditAmountInHomeCurrency** | **string** | The monetary amount of the vendor credit, converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this vendor credit, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Memo** | **string** | A memo or note for this vendor credit. | 
**SalesTaxCode** | [**QbdVendorCreditSalesTaxCode**](QbdVendorCreditSalesTaxCode.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The vendor credit&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of vendor credits to receive this field because it is not returned by default. | 
**ExpenseLines** | [**[]QbdExpenseLine**](QbdExpenseLine.md) | The vendor credit&#39;s expense lines, each representing one line in this expense. | 
**ItemLines** | [**[]QbdItemLine**](QbdItemLine.md) | The vendor credit&#39;s item lines, each representing the purchase of a specific item or service. | 
**ItemLineGroups** | [**[]QbdItemLineGroup**](QbdItemLineGroup.md) | The vendor credit&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | 
**OpenAmount** | **string** | The remaining amount still owed on this vendor credit, represented as a decimal string. This equals the vendor credit&#39;s amount minus any credits or discounts. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the vendor credit object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdVendorCredit

`func NewQbdVendorCredit(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdVendorCreditVendor, payablesAccount QbdVendorCreditPayablesAccount, transactionDate string, creditAmount string, currency QbdVendorCreditCurrency, exchangeRate float32, creditAmountInHomeCurrency string, refNumber string, memo string, salesTaxCode QbdVendorCreditSalesTaxCode, externalId string, linkedTransactions []QbdLinkedTransaction, expenseLines []QbdExpenseLine, itemLines []QbdItemLine, itemLineGroups []QbdItemLineGroup, openAmount string, customFields []QbdCustomField, ) *QbdVendorCredit`

NewQbdVendorCredit instantiates a new QbdVendorCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdVendorCreditWithDefaults

`func NewQbdVendorCreditWithDefaults() *QbdVendorCredit`

NewQbdVendorCreditWithDefaults instantiates a new QbdVendorCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdVendorCredit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdVendorCredit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdVendorCredit) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdVendorCredit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdVendorCredit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdVendorCredit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdVendorCredit) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdVendorCredit) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdVendorCredit) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdVendorCredit) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdVendorCredit) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdVendorCredit) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdVendorCredit) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdVendorCredit) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdVendorCredit) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdVendorCredit) GetVendor() QbdVendorCreditVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdVendorCredit) GetVendorOk() (*QbdVendorCreditVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdVendorCredit) SetVendor(v QbdVendorCreditVendor)`

SetVendor sets Vendor field to given value.


### GetPayablesAccount

`func (o *QbdVendorCredit) GetPayablesAccount() QbdVendorCreditPayablesAccount`

GetPayablesAccount returns the PayablesAccount field if non-nil, zero value otherwise.

### GetPayablesAccountOk

`func (o *QbdVendorCredit) GetPayablesAccountOk() (*QbdVendorCreditPayablesAccount, bool)`

GetPayablesAccountOk returns a tuple with the PayablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccount

`func (o *QbdVendorCredit) SetPayablesAccount(v QbdVendorCreditPayablesAccount)`

SetPayablesAccount sets PayablesAccount field to given value.


### GetTransactionDate

`func (o *QbdVendorCredit) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdVendorCredit) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdVendorCredit) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetCreditAmount

`func (o *QbdVendorCredit) GetCreditAmount() string`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *QbdVendorCredit) GetCreditAmountOk() (*string, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *QbdVendorCredit) SetCreditAmount(v string)`

SetCreditAmount sets CreditAmount field to given value.


### GetCurrency

`func (o *QbdVendorCredit) GetCurrency() QbdVendorCreditCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdVendorCredit) GetCurrencyOk() (*QbdVendorCreditCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdVendorCredit) SetCurrency(v QbdVendorCreditCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdVendorCredit) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdVendorCredit) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdVendorCredit) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetCreditAmountInHomeCurrency

`func (o *QbdVendorCredit) GetCreditAmountInHomeCurrency() string`

GetCreditAmountInHomeCurrency returns the CreditAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetCreditAmountInHomeCurrencyOk

`func (o *QbdVendorCredit) GetCreditAmountInHomeCurrencyOk() (*string, bool)`

GetCreditAmountInHomeCurrencyOk returns a tuple with the CreditAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountInHomeCurrency

`func (o *QbdVendorCredit) SetCreditAmountInHomeCurrency(v string)`

SetCreditAmountInHomeCurrency sets CreditAmountInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdVendorCredit) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdVendorCredit) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdVendorCredit) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetMemo

`func (o *QbdVendorCredit) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdVendorCredit) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdVendorCredit) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetSalesTaxCode

`func (o *QbdVendorCredit) GetSalesTaxCode() QbdVendorCreditSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdVendorCredit) GetSalesTaxCodeOk() (*QbdVendorCreditSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdVendorCredit) SetSalesTaxCode(v QbdVendorCreditSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetExternalId

`func (o *QbdVendorCredit) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdVendorCredit) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdVendorCredit) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdVendorCredit) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdVendorCredit) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdVendorCredit) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetExpenseLines

`func (o *QbdVendorCredit) GetExpenseLines() []QbdExpenseLine`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QbdVendorCredit) GetExpenseLinesOk() (*[]QbdExpenseLine, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QbdVendorCredit) SetExpenseLines(v []QbdExpenseLine)`

SetExpenseLines sets ExpenseLines field to given value.


### GetItemLines

`func (o *QbdVendorCredit) GetItemLines() []QbdItemLine`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QbdVendorCredit) GetItemLinesOk() (*[]QbdItemLine, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QbdVendorCredit) SetItemLines(v []QbdItemLine)`

SetItemLines sets ItemLines field to given value.


### GetItemLineGroups

`func (o *QbdVendorCredit) GetItemLineGroups() []QbdItemLineGroup`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QbdVendorCredit) GetItemLineGroupsOk() (*[]QbdItemLineGroup, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QbdVendorCredit) SetItemLineGroups(v []QbdItemLineGroup)`

SetItemLineGroups sets ItemLineGroups field to given value.


### GetOpenAmount

`func (o *QbdVendorCredit) GetOpenAmount() string`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *QbdVendorCredit) GetOpenAmountOk() (*string, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *QbdVendorCredit) SetOpenAmount(v string)`

SetOpenAmount sets OpenAmount field to given value.


### GetCustomFields

`func (o *QbdVendorCredit) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdVendorCredit) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdVendorCredit) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


