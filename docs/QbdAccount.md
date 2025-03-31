# QbdAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this account. This ID is unique across all accounts but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_account\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this account was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this account was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this account object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this account. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two accounts could both have the &#x60;name&#x60; \&quot;Accounts-Payable\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Corporate:Accounts-Payable\&quot; and \&quot;Finance:Accounts-Payable\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this account, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if an account is under \&quot;Corporate\&quot; and has the &#x60;name&#x60; \&quot;Accounts-Payable\&quot;, its &#x60;fullName&#x60; would be \&quot;Corporate:Accounts-Payable\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all account objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**IsActive** | **bool** | Indicates whether this account is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Parent** | Pointer to [**QbdAccountParent**](QbdAccountParent.md) |  | [optional] 
**Sublevel** | **float32** | The depth level of this account in the hierarchy. A top-level account has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, an account with a &#x60;fullName&#x60; of \&quot;Corporate:Accounts-Payable\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**AccountType** | **string** | The classification of this account, indicating its purpose within the chart of accounts.  **NOTE**: You cannot create an account of type &#x60;non_posting&#x60; through the API because QuickBooks creates these accounts behind the scenes. | 
**SpecialAccountType** | **string** | Indicates if this account is a special account automatically created by QuickBooks for specific purposes. | 
**IsTaxAccount** | Pointer to **bool** | Indicates whether this account is used for tracking taxes. | [optional] 
**AccountNumber** | **string** | The account&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | 
**BankAccountNumber** | Pointer to **string** | The bank account number or identifying note for this account. Access to this field may be restricted based on permissions. | [optional] 
**Description** | **string** | A description of this account. | 
**Balance** | **string** | The current balance of this account only, excluding balances from any subordinate accounts, represented as a decimal string. Compare with &#x60;totalBalance&#x60;. Note that income accounts and balance sheet accounts may not have balances. | 
**TotalBalance** | **string** | The combined balance of this account and all its sub-accounts, represented as a decimal string. For example, the &#x60;totalBalance&#x60; for XYZ Bank would be the total of the balances of all its sub-accounts (checking, savings, and so on). If XYZ Bank did not have any sub-accounts, &#x60;totalBalance&#x60; and &#x60;balance&#x60; would be the same. | 
**SalesTaxCode** | Pointer to [**QbdAccountSalesTaxCode**](QbdAccountSalesTaxCode.md) |  | [optional] 
**TaxLineDetails** | [**QbdAccountTaxLineDetails**](QbdAccountTaxLineDetails.md) |  | 
**CashFlowClassification** | **string** | Indicates how this account is classified for cash flow reporting. If &#x60;none&#x60;, the account has not been classified. If &#x60;not_applicable&#x60;, the account does not qualify to be classified (e.g., a bank account tracking cash transactions is not part of a cash flow report). | 
**Currency** | Pointer to [**QbdAccountCurrency**](QbdAccountCurrency.md) |  | [optional] 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the account object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdAccount

`func NewQbdAccount(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, isActive bool, sublevel float32, accountType string, specialAccountType string, accountNumber string, description string, balance string, totalBalance string, taxLineDetails QbdAccountTaxLineDetails, cashFlowClassification string, customFields []QbdCustomField, ) *QbdAccount`

NewQbdAccount instantiates a new QbdAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdAccountWithDefaults

`func NewQbdAccountWithDefaults() *QbdAccount`

NewQbdAccountWithDefaults instantiates a new QbdAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdAccount) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdAccount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdAccount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdAccount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdAccount) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdAccount) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdAccount) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdAccount) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdAccount) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdAccount) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdAccount) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsActive

`func (o *QbdAccount) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdAccount) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdAccount) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetParent

`func (o *QbdAccount) GetParent() QbdAccountParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdAccount) GetParentOk() (*QbdAccountParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdAccount) SetParent(v QbdAccountParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *QbdAccount) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSublevel

`func (o *QbdAccount) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdAccount) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdAccount) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetAccountType

`func (o *QbdAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *QbdAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *QbdAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetSpecialAccountType

`func (o *QbdAccount) GetSpecialAccountType() string`

GetSpecialAccountType returns the SpecialAccountType field if non-nil, zero value otherwise.

### GetSpecialAccountTypeOk

`func (o *QbdAccount) GetSpecialAccountTypeOk() (*string, bool)`

GetSpecialAccountTypeOk returns a tuple with the SpecialAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAccountType

`func (o *QbdAccount) SetSpecialAccountType(v string)`

SetSpecialAccountType sets SpecialAccountType field to given value.


### GetIsTaxAccount

`func (o *QbdAccount) GetIsTaxAccount() bool`

GetIsTaxAccount returns the IsTaxAccount field if non-nil, zero value otherwise.

### GetIsTaxAccountOk

`func (o *QbdAccount) GetIsTaxAccountOk() (*bool, bool)`

GetIsTaxAccountOk returns a tuple with the IsTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxAccount

`func (o *QbdAccount) SetIsTaxAccount(v bool)`

SetIsTaxAccount sets IsTaxAccount field to given value.

### HasIsTaxAccount

`func (o *QbdAccount) HasIsTaxAccount() bool`

HasIsTaxAccount returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QbdAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QbdAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QbdAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBankAccountNumber

`func (o *QbdAccount) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *QbdAccount) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *QbdAccount) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *QbdAccount) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetDescription

`func (o *QbdAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdAccount) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBalance

`func (o *QbdAccount) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *QbdAccount) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *QbdAccount) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetTotalBalance

`func (o *QbdAccount) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *QbdAccount) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *QbdAccount) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.


### GetSalesTaxCode

`func (o *QbdAccount) GetSalesTaxCode() QbdAccountSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdAccount) GetSalesTaxCodeOk() (*QbdAccountSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdAccount) SetSalesTaxCode(v QbdAccountSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.

### HasSalesTaxCode

`func (o *QbdAccount) HasSalesTaxCode() bool`

HasSalesTaxCode returns a boolean if a field has been set.

### GetTaxLineDetails

`func (o *QbdAccount) GetTaxLineDetails() QbdAccountTaxLineDetails`

GetTaxLineDetails returns the TaxLineDetails field if non-nil, zero value otherwise.

### GetTaxLineDetailsOk

`func (o *QbdAccount) GetTaxLineDetailsOk() (*QbdAccountTaxLineDetails, bool)`

GetTaxLineDetailsOk returns a tuple with the TaxLineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLineDetails

`func (o *QbdAccount) SetTaxLineDetails(v QbdAccountTaxLineDetails)`

SetTaxLineDetails sets TaxLineDetails field to given value.


### GetCashFlowClassification

`func (o *QbdAccount) GetCashFlowClassification() string`

GetCashFlowClassification returns the CashFlowClassification field if non-nil, zero value otherwise.

### GetCashFlowClassificationOk

`func (o *QbdAccount) GetCashFlowClassificationOk() (*string, bool)`

GetCashFlowClassificationOk returns a tuple with the CashFlowClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowClassification

`func (o *QbdAccount) SetCashFlowClassification(v string)`

SetCashFlowClassification sets CashFlowClassification field to given value.


### GetCurrency

`func (o *QbdAccount) GetCurrency() QbdAccountCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdAccount) GetCurrencyOk() (*QbdAccountCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdAccount) SetCurrency(v QbdAccountCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QbdAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomFields

`func (o *QbdAccount) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdAccount) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdAccount) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


