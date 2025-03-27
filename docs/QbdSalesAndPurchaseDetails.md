# QbdSalesAndPurchaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesDescription** | **string** | The description of this item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | 
**SalesPrice** | **string** | The price at which this item is sold to customers, represented as a decimal string. | 
**IncomeAccount** | [**QbdSalesAndPurchaseDetailsIncomeAccount**](QbdSalesAndPurchaseDetailsIncomeAccount.md) |  | 
**PurchaseDescription** | **string** | The description of this item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | 
**PurchaseCost** | **string** | The cost at which this item is purchased from vendors, represented as a decimal string. | 
**PurchaseTaxCode** | [**QbdSalesAndPurchaseDetailsPurchaseTaxCode**](QbdSalesAndPurchaseDetailsPurchaseTaxCode.md) |  | 
**ExpenseAccount** | [**QbdSalesAndPurchaseDetailsExpenseAccount**](QbdSalesAndPurchaseDetailsExpenseAccount.md) |  | 
**PreferredVendor** | [**QbdSalesAndPurchaseDetailsPreferredVendor**](QbdSalesAndPurchaseDetailsPreferredVendor.md) |  | 

## Methods

### NewQbdSalesAndPurchaseDetails

`func NewQbdSalesAndPurchaseDetails(salesDescription string, salesPrice string, incomeAccount QbdSalesAndPurchaseDetailsIncomeAccount, purchaseDescription string, purchaseCost string, purchaseTaxCode QbdSalesAndPurchaseDetailsPurchaseTaxCode, expenseAccount QbdSalesAndPurchaseDetailsExpenseAccount, preferredVendor QbdSalesAndPurchaseDetailsPreferredVendor, ) *QbdSalesAndPurchaseDetails`

NewQbdSalesAndPurchaseDetails instantiates a new QbdSalesAndPurchaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesAndPurchaseDetailsWithDefaults

`func NewQbdSalesAndPurchaseDetailsWithDefaults() *QbdSalesAndPurchaseDetails`

NewQbdSalesAndPurchaseDetailsWithDefaults instantiates a new QbdSalesAndPurchaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesDescription

`func (o *QbdSalesAndPurchaseDetails) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QbdSalesAndPurchaseDetails) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QbdSalesAndPurchaseDetails) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.


### GetSalesPrice

`func (o *QbdSalesAndPurchaseDetails) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QbdSalesAndPurchaseDetails) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QbdSalesAndPurchaseDetails) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.


### GetIncomeAccount

`func (o *QbdSalesAndPurchaseDetails) GetIncomeAccount() QbdSalesAndPurchaseDetailsIncomeAccount`

GetIncomeAccount returns the IncomeAccount field if non-nil, zero value otherwise.

### GetIncomeAccountOk

`func (o *QbdSalesAndPurchaseDetails) GetIncomeAccountOk() (*QbdSalesAndPurchaseDetailsIncomeAccount, bool)`

GetIncomeAccountOk returns a tuple with the IncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccount

`func (o *QbdSalesAndPurchaseDetails) SetIncomeAccount(v QbdSalesAndPurchaseDetailsIncomeAccount)`

SetIncomeAccount sets IncomeAccount field to given value.


### GetPurchaseDescription

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QbdSalesAndPurchaseDetails) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.


### GetPurchaseCost

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QbdSalesAndPurchaseDetails) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.


### GetPurchaseTaxCode

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseTaxCode() QbdSalesAndPurchaseDetailsPurchaseTaxCode`

GetPurchaseTaxCode returns the PurchaseTaxCode field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeOk

`func (o *QbdSalesAndPurchaseDetails) GetPurchaseTaxCodeOk() (*QbdSalesAndPurchaseDetailsPurchaseTaxCode, bool)`

GetPurchaseTaxCodeOk returns a tuple with the PurchaseTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCode

`func (o *QbdSalesAndPurchaseDetails) SetPurchaseTaxCode(v QbdSalesAndPurchaseDetailsPurchaseTaxCode)`

SetPurchaseTaxCode sets PurchaseTaxCode field to given value.


### GetExpenseAccount

`func (o *QbdSalesAndPurchaseDetails) GetExpenseAccount() QbdSalesAndPurchaseDetailsExpenseAccount`

GetExpenseAccount returns the ExpenseAccount field if non-nil, zero value otherwise.

### GetExpenseAccountOk

`func (o *QbdSalesAndPurchaseDetails) GetExpenseAccountOk() (*QbdSalesAndPurchaseDetailsExpenseAccount, bool)`

GetExpenseAccountOk returns a tuple with the ExpenseAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAccount

`func (o *QbdSalesAndPurchaseDetails) SetExpenseAccount(v QbdSalesAndPurchaseDetailsExpenseAccount)`

SetExpenseAccount sets ExpenseAccount field to given value.


### GetPreferredVendor

`func (o *QbdSalesAndPurchaseDetails) GetPreferredVendor() QbdSalesAndPurchaseDetailsPreferredVendor`

GetPreferredVendor returns the PreferredVendor field if non-nil, zero value otherwise.

### GetPreferredVendorOk

`func (o *QbdSalesAndPurchaseDetails) GetPreferredVendorOk() (*QbdSalesAndPurchaseDetailsPreferredVendor, bool)`

GetPreferredVendorOk returns a tuple with the PreferredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendor

`func (o *QbdSalesAndPurchaseDetails) SetPreferredVendor(v QbdSalesAndPurchaseDetailsPreferredVendor)`

SetPreferredVendor sets PreferredVendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


