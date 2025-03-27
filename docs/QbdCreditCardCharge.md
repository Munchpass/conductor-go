# QbdCreditCardCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this credit card charge. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_credit_card_charge\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this credit card charge was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this credit card charge was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this credit card charge object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Account** | [**QbdCreditCardChargeAccount**](QbdCreditCardChargeAccount.md) |  | 
**Payee** | [**QbdCreditCardChargePayee**](QbdCreditCardChargePayee.md) |  | 
**TransactionDate** | **string** | The date of this credit card charge, in ISO 8601 format (YYYY-MM-DD). | 
**Amount** | **string** | The total monetary amount of this credit card charge, represented as a decimal string. This equals the sum of the amounts in the credit card charge&#39;s expense lines, item lines, and item group lines. | 
**Currency** | [**QbdCreditCardChargeCurrency**](QbdCreditCardChargeCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this credit card charge&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**AmountInHomeCurrency** | **string** | The monetary amount of this credit card charge converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this credit card charge, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**Memo** | **string** | A memo or note for this credit card charge. | 
**SalesTaxCode** | [**QbdCreditCardChargeSalesTaxCode**](QbdCreditCardChargeSalesTaxCode.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**ExpenseLines** | [**[]QbdExpenseLine**](QbdExpenseLine.md) | The credit card charge&#39;s expense lines, each representing one line in this expense. | 
**ItemLines** | [**[]QbdItemLine**](QbdItemLine.md) | The credit card charge&#39;s item lines, each representing the purchase of a specific item or service. | 
**ItemLineGroups** | [**[]QbdItemLineGroup**](QbdItemLineGroup.md) | The credit card charge&#39;s item group lines, each representing a predefined set of items bundled together because they are commonly purchased together or grouped for faster entry. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the credit card charge object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdCreditCardCharge

`func NewQbdCreditCardCharge(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, account QbdCreditCardChargeAccount, payee QbdCreditCardChargePayee, transactionDate string, amount string, currency QbdCreditCardChargeCurrency, exchangeRate float32, amountInHomeCurrency string, refNumber string, memo string, salesTaxCode QbdCreditCardChargeSalesTaxCode, externalId string, expenseLines []QbdExpenseLine, itemLines []QbdItemLine, itemLineGroups []QbdItemLineGroup, customFields []QbdCustomField, ) *QbdCreditCardCharge`

NewQbdCreditCardCharge instantiates a new QbdCreditCardCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditCardChargeWithDefaults

`func NewQbdCreditCardChargeWithDefaults() *QbdCreditCardCharge`

NewQbdCreditCardChargeWithDefaults instantiates a new QbdCreditCardCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCreditCardCharge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCreditCardCharge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCreditCardCharge) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdCreditCardCharge) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdCreditCardCharge) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdCreditCardCharge) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdCreditCardCharge) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdCreditCardCharge) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdCreditCardCharge) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdCreditCardCharge) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdCreditCardCharge) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdCreditCardCharge) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdCreditCardCharge) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdCreditCardCharge) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdCreditCardCharge) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetAccount

`func (o *QbdCreditCardCharge) GetAccount() QbdCreditCardChargeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdCreditCardCharge) GetAccountOk() (*QbdCreditCardChargeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdCreditCardCharge) SetAccount(v QbdCreditCardChargeAccount)`

SetAccount sets Account field to given value.


### GetPayee

`func (o *QbdCreditCardCharge) GetPayee() QbdCreditCardChargePayee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *QbdCreditCardCharge) GetPayeeOk() (*QbdCreditCardChargePayee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *QbdCreditCardCharge) SetPayee(v QbdCreditCardChargePayee)`

SetPayee sets Payee field to given value.


### GetTransactionDate

`func (o *QbdCreditCardCharge) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdCreditCardCharge) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdCreditCardCharge) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetAmount

`func (o *QbdCreditCardCharge) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdCreditCardCharge) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdCreditCardCharge) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *QbdCreditCardCharge) GetCurrency() QbdCreditCardChargeCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdCreditCardCharge) GetCurrencyOk() (*QbdCreditCardChargeCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdCreditCardCharge) SetCurrency(v QbdCreditCardChargeCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdCreditCardCharge) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdCreditCardCharge) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdCreditCardCharge) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountInHomeCurrency

`func (o *QbdCreditCardCharge) GetAmountInHomeCurrency() string`

GetAmountInHomeCurrency returns the AmountInHomeCurrency field if non-nil, zero value otherwise.

### GetAmountInHomeCurrencyOk

`func (o *QbdCreditCardCharge) GetAmountInHomeCurrencyOk() (*string, bool)`

GetAmountInHomeCurrencyOk returns a tuple with the AmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInHomeCurrency

`func (o *QbdCreditCardCharge) SetAmountInHomeCurrency(v string)`

SetAmountInHomeCurrency sets AmountInHomeCurrency field to given value.


### GetRefNumber

`func (o *QbdCreditCardCharge) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdCreditCardCharge) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdCreditCardCharge) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetMemo

`func (o *QbdCreditCardCharge) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdCreditCardCharge) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdCreditCardCharge) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetSalesTaxCode

`func (o *QbdCreditCardCharge) GetSalesTaxCode() QbdCreditCardChargeSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdCreditCardCharge) GetSalesTaxCodeOk() (*QbdCreditCardChargeSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdCreditCardCharge) SetSalesTaxCode(v QbdCreditCardChargeSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetExternalId

`func (o *QbdCreditCardCharge) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdCreditCardCharge) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdCreditCardCharge) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetExpenseLines

`func (o *QbdCreditCardCharge) GetExpenseLines() []QbdExpenseLine`

GetExpenseLines returns the ExpenseLines field if non-nil, zero value otherwise.

### GetExpenseLinesOk

`func (o *QbdCreditCardCharge) GetExpenseLinesOk() (*[]QbdExpenseLine, bool)`

GetExpenseLinesOk returns a tuple with the ExpenseLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseLines

`func (o *QbdCreditCardCharge) SetExpenseLines(v []QbdExpenseLine)`

SetExpenseLines sets ExpenseLines field to given value.


### GetItemLines

`func (o *QbdCreditCardCharge) GetItemLines() []QbdItemLine`

GetItemLines returns the ItemLines field if non-nil, zero value otherwise.

### GetItemLinesOk

`func (o *QbdCreditCardCharge) GetItemLinesOk() (*[]QbdItemLine, bool)`

GetItemLinesOk returns a tuple with the ItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLines

`func (o *QbdCreditCardCharge) SetItemLines(v []QbdItemLine)`

SetItemLines sets ItemLines field to given value.


### GetItemLineGroups

`func (o *QbdCreditCardCharge) GetItemLineGroups() []QbdItemLineGroup`

GetItemLineGroups returns the ItemLineGroups field if non-nil, zero value otherwise.

### GetItemLineGroupsOk

`func (o *QbdCreditCardCharge) GetItemLineGroupsOk() (*[]QbdItemLineGroup, bool)`

GetItemLineGroupsOk returns a tuple with the ItemLineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLineGroups

`func (o *QbdCreditCardCharge) SetItemLineGroups(v []QbdItemLineGroup)`

SetItemLineGroups sets ItemLineGroups field to given value.


### GetCustomFields

`func (o *QbdCreditCardCharge) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCreditCardCharge) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCreditCardCharge) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


