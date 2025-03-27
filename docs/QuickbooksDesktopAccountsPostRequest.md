# QuickbooksDesktopAccountsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this account. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two accounts could both have the &#x60;name&#x60; \&quot;Accounts-Payable\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Corporate:Accounts-Payable\&quot; and \&quot;Finance:Accounts-Payable\&quot;.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this account is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ParentId** | Pointer to **string** | The parent account one level above this one in the hierarchy. For example, if this account has a &#x60;fullName&#x60; of \&quot;Corporate:Accounts-Payable\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Corporate\&quot;. If this account is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**AccountType** | **string** | The classification of this account, indicating its purpose within the chart of accounts.  **NOTE**: You cannot create an account of type &#x60;non_posting&#x60; through the API because QuickBooks creates these accounts behind the scenes. | 
**AccountNumber** | Pointer to **string** | The account&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**BankAccountNumber** | Pointer to **string** | The bank account number or identifying note for this account. Access to this field may be restricted based on permissions. | [optional] 
**Description** | Pointer to **string** | A description of this account. | [optional] 
**OpeningBalance** | Pointer to **string** | The amount of money in, or the value of, this account as of &#x60;openingBalanceDate&#x60;. On a bank statement, this would be the amount of money in the account at the beginning of the statement period. | [optional] 
**OpeningBalanceDate** | Pointer to **string** | The date of the opening balance of this account, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for transactions with this account, determining whether the transactions are taxable or non-taxable. This can be overridden at the transaction or transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**TaxLineId** | Pointer to **float32** | The identifier of the tax line associated with this account. | [optional] 
**CurrencyId** | Pointer to **string** | The account&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 

## Methods

### NewQuickbooksDesktopAccountsPostRequest

`func NewQuickbooksDesktopAccountsPostRequest(name string, accountType string, ) *QuickbooksDesktopAccountsPostRequest`

NewQuickbooksDesktopAccountsPostRequest instantiates a new QuickbooksDesktopAccountsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopAccountsPostRequestWithDefaults

`func NewQuickbooksDesktopAccountsPostRequestWithDefaults() *QuickbooksDesktopAccountsPostRequest`

NewQuickbooksDesktopAccountsPostRequestWithDefaults instantiates a new QuickbooksDesktopAccountsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopAccountsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopAccountsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopAccountsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopAccountsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopAccountsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopAccountsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopAccountsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopAccountsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAccountType

`func (o *QuickbooksDesktopAccountsPostRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *QuickbooksDesktopAccountsPostRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBankAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *QuickbooksDesktopAccountsPostRequest) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopAccountsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopAccountsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopAccountsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *QuickbooksDesktopAccountsPostRequest) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *QuickbooksDesktopAccountsPostRequest) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *QuickbooksDesktopAccountsPostRequest) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *QuickbooksDesktopAccountsPostRequest) GetOpeningBalanceDate() string`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetOpeningBalanceDateOk() (*string, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *QuickbooksDesktopAccountsPostRequest) SetOpeningBalanceDate(v string)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *QuickbooksDesktopAccountsPostRequest) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopAccountsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopAccountsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopAccountsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetTaxLineId

`func (o *QuickbooksDesktopAccountsPostRequest) GetTaxLineId() float32`

GetTaxLineId returns the TaxLineId field if non-nil, zero value otherwise.

### GetTaxLineIdOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetTaxLineIdOk() (*float32, bool)`

GetTaxLineIdOk returns a tuple with the TaxLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLineId

`func (o *QuickbooksDesktopAccountsPostRequest) SetTaxLineId(v float32)`

SetTaxLineId sets TaxLineId field to given value.

### HasTaxLineId

`func (o *QuickbooksDesktopAccountsPostRequest) HasTaxLineId() bool`

HasTaxLineId returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopAccountsPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopAccountsPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopAccountsPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopAccountsPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


