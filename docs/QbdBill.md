# QbdBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this bill. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_bill\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this bill was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this bill was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this bill object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdBillVendor**](QbdBillVendor.md) |  | 
**VendorAddress** | [**QbdBillVendorAddress**](QbdBillVendorAddress.md) |  | 
**PayablesAccount** | [**QbdBillPayablesAccount**](QbdBillPayablesAccount.md) |  | 
**TransactionDate** | **string** | The date of this bill, in ISO 8601 format (YYYY-MM-DD). | 
**DueDate** | **string** | The date by which this bill must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**AmountDue** | **string** | The total monetary amount due for this bill, represented as a decimal string. This equals the sum of the amounts in the bill&#39;s expense lines, item lines, and item group lines. It also equals &#x60;openAmount&#x60; plus any credits or discounts. | 
**Currency** | [**QbdBillCurrency**](QbdBillCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this bill&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountDueInHomeCurrency** | **string** | The monetary amount due for this bill converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this bill, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**IsPending** | **bool** | Indicates whether this bill has not been completed or is in a draft version. | 
**Terms** | [**QbdBillTerms**](QbdBillTerms.md) |  | 
**Memo** | **string** | A memo or note for this bill that appears in the Accounts-Payable register and in reports that include this bill. | 
**SalesTaxCode** | [**QbdBillSalesTaxCode**](QbdBillSalesTaxCode.md) |  | 
**IsPaid** | **bool** | Indicates whether this bill has been paid in full. When &#x60;true&#x60;, &#x60;openAmount&#x60; will be 0. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The bill&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of bills to receive this field because it is not returned by default. | 
**ExpenseLines** | [**[]QbdExpenseLine**](QbdExpenseLine.md) | The bill&#39;s expense lines, each representing one line in this expense. | 
**ItemLines** | [**[]QbdItemLine**](QbdItemLine.md) | The bill&#39;s item lines, each representing the purchase of a specific item or service. | 
**ItemLineGroups** | [**[]QbdItemLineGroup**](QbdItemLineGroup.md) | The bill&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | 
**OpenAmount** | **string** | The remaining amount still owed on this bill, represented as a decimal string. This equals the bill&#39;s amount minus any credits or discounts. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the bill object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdBill

`func NewQbdBill(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdBillVendor, vendorAddress QbdBillVendorAddress, payablesAccount QbdBillPayablesAccount, transactionDate string, dueDate string, amountDue string, currency QbdBillCurrency, exchangeRate float32, amountDueInHomeCurrency string, refNumber string, isPending bool, terms QbdBillTerms, memo string, salesTaxCode QbdBillSalesTaxCode, isPaid bool, externalId string, linkedTransactions []QbdLinkedTransaction, expenseLines []QbdExpenseLine, itemLines []QbdItemLine, itemLineGroups []QbdItemLineGroup, openAmount string, customFields []QbdCustomField, ) *QbdBill`

NewQbdBill instantiates a new QbdBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdBillWithDefaults

`func NewQbdBillWithDefaults() *QbdBill`

NewQbdBillWithDefaults instantiates a new QbdBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdBill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdBill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdBill) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdBill) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdBill) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdBill) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdBill) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdBill) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdBill) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdBill) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdBill) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdBill) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdBill) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdBill) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdBill) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdBill) GetVendor() QbdBillVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdBill) GetVendorOk() (*QbdBillVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdBill) SetVendor(v QbdBillVendor)`

SetVendor sets Vendor field to given value.


### GetVendorAddress

`func (o *QbdBill) GetVendorAddress() QbdBillVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QbdBill) GetVendorAddressOk() (*QbdBillVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QbdBill) SetVendorAddress(v QbdBillVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.


### GetPayablesAccount

`func (o *QbdBill) GetPayablesAccount() QbdBillPayablesAccount`

GetPayablesAccount returns the PayablesAccount field if non-nil, zero value otherwise.

### GetPayablesAccountOk

`func (o *QbdBill) GetPayablesAccountOk() (*QbdBillPayablesAccount, bool)`

GetPayablesAccountOk returns a tuple with the PayablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayablesAccount

`func (o *QbdBill) SetPayablesAccount(v QbdBillPayablesAccount)`

SetPayablesAccount sets PayablesAccount field to given value.


### GetTransactionDate

`func (o *QbdBill) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdBill) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdBill) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetDueDate

`func (o *QbdBill) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdBill) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdBill) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetAmountDue

`func (o *QbdBill) GetAmountDue() string`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *QbdBill) GetAmountDueOk() (*string, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *QbdBill) SetAmountDue(v string)`

SetAmountDue sets AmountDue field to given value.


### GetCurrency

`func (o *QbdBill) GetCurrency() QbdBillCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdBill) GetCurrencyOk() (*QbdBillCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdBill) SetCurrency(v QbdBillCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdBill) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdBill) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdBill) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountDueInHomeCurrency

`func (o *QbdBill) GetAmountDueInHomeCurrency() string`

GetAmountDueInHomeCurrency returns the AmountDueInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountDueInHomeCurrencyOk

`func (o *QbdBill) GetAmountDueInHomeCurrencyOk() (*string, bool)`

GetAmountDueInHomeCurrencyOk returns a tuple with the AmountDueInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueInHomeCurrency

`func (o *QbdBill) SetAmountDueInHomeCurrency(v string)`

SetAmountDueInHomeCurrency sets AmountDueInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdBill) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdBill) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdBill) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetIsPending

`func (o *QbdBill) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QbdBill) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QbdBill) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetTerms

`func (o *QbdBill) GetTerms() QbdBillTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdBill) GetTermsOk() (*QbdBillTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdBill) SetTerms(v QbdBillTerms)`

SetTerms sets Terms field to given value.


### GetMemo

`func (o *QbdBill) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdBill) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdBill) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetSalesTaxCode

`func (o *QbdBill) GetSalesTaxCode() QbdBillSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdBill) GetSalesTaxCodeOk() (*QbdBillSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdBill) SetSalesTaxCode(v QbdBillSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetIsPaid

`func (o *QbdBill) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *QbdBill) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *QbdBill) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.


### GetExternalId

`func (o *QbdBill) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdBill) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdBill) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdBill) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdBill) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdBill) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetExpenseLines

`func (o *QbdBill) GetExpenseLines() []QbdExpenseLine`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QbdBill) GetExpenseLinesOk() (*[]QbdExpenseLine, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QbdBill) SetExpenseLines(v []QbdExpenseLine)`

SetExpenseLines sets ExpenseLines field to given value.


### GetItemLines

`func (o *QbdBill) GetItemLines() []QbdItemLine`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QbdBill) GetItemLinesOk() (*[]QbdItemLine, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QbdBill) SetItemLines(v []QbdItemLine)`

SetItemLines sets ItemLines field to given value.


### GetItemLineGroups

`func (o *QbdBill) GetItemLineGroups() []QbdItemLineGroup`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QbdBill) GetItemLineGroupsOk() (*[]QbdItemLineGroup, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QbdBill) SetItemLineGroups(v []QbdItemLineGroup)`

SetItemLineGroups sets ItemLineGroups field to given value.


### GetOpenAmount

`func (o *QbdBill) GetOpenAmount() string`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *QbdBill) GetOpenAmountOk() (*string, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *QbdBill) SetOpenAmount(v string)`

SetOpenAmount sets OpenAmount field to given value.


### GetCustomFields

`func (o *QbdBill) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdBill) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdBill) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


