# QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesDescription** | Pointer to **string** | The description of this item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | Pointer to **string** | The income account used to track revenue from sales of this item. | [optional] 
**PurchaseDescription** | Pointer to **string** | The description of this item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**ExpenseAccountId** | Pointer to **string** | The expense account used to track costs from purchases of this item. | [optional] 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this item is typically purchased. | [optional] 
**UpdateExistingTransactionsIncomeAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new income account (specified by the &#x60;incomeAccountId&#x60; field) to all existing transactions that use this item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 
**UpdateExistingTransactionsExpenseAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new expense account (specified by the &#x60;expenseAccountId&#x60; field) to all existing transactions that use this item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 

## Methods

### NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails

`func NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails() *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails`

NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetailsWithDefaults

`func NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetailsWithDefaults() *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails`

NewQuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetailsWithDefaults instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.

### HasIncomeAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasIncomeAccountId() bool`

HasIncomeAccountId returns a boolean if a field has been set.

### GetPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetExpenseAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetExpenseAccountId() string`

GetExpenseAccountId returns the ExpenseAccountId field if non-nil, zero value otherwise.

### GetExpenseAccountIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetExpenseAccountIdOk() (*string, bool)`

GetExpenseAccountIdOk returns a tuple with the ExpenseAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetExpenseAccountId(v string)`

SetExpenseAccountId sets ExpenseAccountId field to given value.

### HasExpenseAccountId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasExpenseAccountId() bool`

HasExpenseAccountId returns a boolean if a field has been set.

### GetPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsIncomeAccount() bool`

GetUpdateExistingTransactionsIncomeAccount returns the UpdateExistingTransactionsIncomeAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsIncomeAccountOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsIncomeAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsIncomeAccountOk returns a tuple with the UpdateExistingTransactionsIncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetUpdateExistingTransactionsIncomeAccount(v bool)`

SetUpdateExistingTransactionsIncomeAccount sets UpdateExistingTransactionsIncomeAccount field to given value.

### HasUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasUpdateExistingTransactionsIncomeAccount() bool`

HasUpdateExistingTransactionsIncomeAccount returns a boolean if a field has been set.

### GetUpdateExistingTransactionsExpenseAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsExpenseAccount() bool`

GetUpdateExistingTransactionsExpenseAccount returns the UpdateExistingTransactionsExpenseAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsExpenseAccountOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) GetUpdateExistingTransactionsExpenseAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsExpenseAccountOk returns a tuple with the UpdateExistingTransactionsExpenseAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsExpenseAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) SetUpdateExistingTransactionsExpenseAccount(v bool)`

SetUpdateExistingTransactionsExpenseAccount sets UpdateExistingTransactionsExpenseAccount field to given value.

### HasUpdateExistingTransactionsExpenseAccount

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails) HasUpdateExistingTransactionsExpenseAccount() bool`

HasUpdateExistingTransactionsExpenseAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


