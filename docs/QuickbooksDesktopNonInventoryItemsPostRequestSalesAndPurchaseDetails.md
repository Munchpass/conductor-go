# QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesDescription** | Pointer to **string** | The description of this item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | **string** | The income account used to track revenue from sales of this item. | 
**PurchaseDescription** | Pointer to **string** | The description of this item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**ExpenseAccountId** | **string** | The expense account used to track costs from purchases of this item. | 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this item is typically purchased. | [optional] 

## Methods

### NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails

`func NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails(incomeAccountId string, expenseAccountId string, ) *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails`

NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails instantiates a new QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetailsWithDefaults

`func NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetailsWithDefaults() *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails`

NewQuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetailsWithDefaults instantiates a new QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.


### GetPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetExpenseAccountId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetExpenseAccountId() string`

GetExpenseAccountId returns the ExpenseAccountId field if non-nil, zero value otherwise.

### GetExpenseAccountIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetExpenseAccountIdOk() (*string, bool)`

GetExpenseAccountIdOk returns a tuple with the ExpenseAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAccountId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetExpenseAccountId(v string)`

SetExpenseAccountId sets ExpenseAccountId field to given value.


### GetPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


