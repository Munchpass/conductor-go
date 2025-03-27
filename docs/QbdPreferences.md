# QbdPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | [**TheAccountingPreferencesObject**](TheAccountingPreferencesObject.md) |  | 
**FinanceCharges** | [**TheFinanceChargePreferencesObject**](TheFinanceChargePreferencesObject.md) |  | 
**JobsAndEstimates** | [**TheJobsAndEstimatesPreferencesObject**](TheJobsAndEstimatesPreferencesObject.md) |  | 
**MultiCurrency** | [**QbdPreferencesMultiCurrency**](QbdPreferencesMultiCurrency.md) |  | 
**MultiLocationInventory** | [**QbdPreferencesMultiLocationInventory**](QbdPreferencesMultiLocationInventory.md) |  | 
**PurchasesAndVendors** | [**ThePurchasesAndVendorsPreferencesObject**](ThePurchasesAndVendorsPreferencesObject.md) |  | 
**Reports** | [**TheReportsPreferencesObject**](TheReportsPreferencesObject.md) |  | 
**SalesAndCustomers** | [**TheSalesAndCustomersPreferencesObject**](TheSalesAndCustomersPreferencesObject.md) |  | 
**SalesTax** | [**QbdPreferencesSalesTax**](QbdPreferencesSalesTax.md) |  | 
**TimeTracking** | [**QbdPreferencesTimeTracking**](QbdPreferencesTimeTracking.md) |  | 
**AppAccessRights** | [**TheCurrentAppAccessRightsObject**](TheCurrentAppAccessRightsObject.md) |  | 
**ItemsAndInventory** | [**QbdPreferencesItemsAndInventory**](QbdPreferencesItemsAndInventory.md) |  | 

## Methods

### NewQbdPreferences

`func NewQbdPreferences(accounting TheAccountingPreferencesObject, financeCharges TheFinanceChargePreferencesObject, jobsAndEstimates TheJobsAndEstimatesPreferencesObject, multiCurrency QbdPreferencesMultiCurrency, multiLocationInventory QbdPreferencesMultiLocationInventory, purchasesAndVendors ThePurchasesAndVendorsPreferencesObject, reports TheReportsPreferencesObject, salesAndCustomers TheSalesAndCustomersPreferencesObject, salesTax QbdPreferencesSalesTax, timeTracking QbdPreferencesTimeTracking, appAccessRights TheCurrentAppAccessRightsObject, itemsAndInventory QbdPreferencesItemsAndInventory, ) *QbdPreferences`

NewQbdPreferences instantiates a new QbdPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPreferencesWithDefaults

`func NewQbdPreferencesWithDefaults() *QbdPreferences`

NewQbdPreferencesWithDefaults instantiates a new QbdPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *QbdPreferences) GetAccounting() TheAccountingPreferencesObject`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *QbdPreferences) GetAccountingOk() (*TheAccountingPreferencesObject, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *QbdPreferences) SetAccounting(v TheAccountingPreferencesObject)`

SetAccounting sets Accounting field to given value.


### GetFinanceCharges

`func (o *QbdPreferences) GetFinanceCharges() TheFinanceChargePreferencesObject`

GetFinanceCharges returns the FinanceCharges field if non-nil, zero value otherwise.

### GetFinanceChargesOk

`func (o *QbdPreferences) GetFinanceChargesOk() (*TheFinanceChargePreferencesObject, bool)`

GetFinanceChargesOk returns a tuple with the FinanceCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceCharges

`func (o *QbdPreferences) SetFinanceCharges(v TheFinanceChargePreferencesObject)`

SetFinanceCharges sets FinanceCharges field to given value.


### GetJobsAndEstimates

`func (o *QbdPreferences) GetJobsAndEstimates() TheJobsAndEstimatesPreferencesObject`

GetJobsAndEstimates returns the JobsAndEstimates field if non-nil, zero value otherwise.

### GetJobsAndEstimatesOk

`func (o *QbdPreferences) GetJobsAndEstimatesOk() (*TheJobsAndEstimatesPreferencesObject, bool)`

GetJobsAndEstimatesOk returns a tuple with the JobsAndEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsAndEstimates

`func (o *QbdPreferences) SetJobsAndEstimates(v TheJobsAndEstimatesPreferencesObject)`

SetJobsAndEstimates sets JobsAndEstimates field to given value.


### GetMultiCurrency

`func (o *QbdPreferences) GetMultiCurrency() QbdPreferencesMultiCurrency`

GetMultiCurrency returns the MultiCurrency field if non-nil, zero value otherwise.

### GetMultiCurrencyOk

`func (o *QbdPreferences) GetMultiCurrencyOk() (*QbdPreferencesMultiCurrency, bool)`

GetMultiCurrencyOk returns a tuple with the MultiCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCurrency

`func (o *QbdPreferences) SetMultiCurrency(v QbdPreferencesMultiCurrency)`

SetMultiCurrency sets MultiCurrency field to given value.


### GetMultiLocationInventory

`func (o *QbdPreferences) GetMultiLocationInventory() QbdPreferencesMultiLocationInventory`

GetMultiLocationInventory returns the MultiLocationInventory field if non-nil, zero value otherwise.

### GetMultiLocationInventoryOk

`func (o *QbdPreferences) GetMultiLocationInventoryOk() (*QbdPreferencesMultiLocationInventory, bool)`

GetMultiLocationInventoryOk returns a tuple with the MultiLocationInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiLocationInventory

`func (o *QbdPreferences) SetMultiLocationInventory(v QbdPreferencesMultiLocationInventory)`

SetMultiLocationInventory sets MultiLocationInventory field to given value.


### GetPurchasesAndVendors

`func (o *QbdPreferences) GetPurchasesAndVendors() ThePurchasesAndVendorsPreferencesObject`

GetPurchasesAndVendors returns the PurchasesAndVendors field if non-nil, zero value otherwise.

### GetPurchasesAndVendorsOk

`func (o *QbdPreferences) GetPurchasesAndVendorsOk() (*ThePurchasesAndVendorsPreferencesObject, bool)`

GetPurchasesAndVendorsOk returns a tuple with the PurchasesAndVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesAndVendors

`func (o *QbdPreferences) SetPurchasesAndVendors(v ThePurchasesAndVendorsPreferencesObject)`

SetPurchasesAndVendors sets PurchasesAndVendors field to given value.


### GetReports

`func (o *QbdPreferences) GetReports() TheReportsPreferencesObject`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *QbdPreferences) GetReportsOk() (*TheReportsPreferencesObject, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *QbdPreferences) SetReports(v TheReportsPreferencesObject)`

SetReports sets Reports field to given value.


### GetSalesAndCustomers

`func (o *QbdPreferences) GetSalesAndCustomers() TheSalesAndCustomersPreferencesObject`

GetSalesAndCustomers returns the SalesAndCustomers field if non-nil, zero value otherwise.

### GetSalesAndCustomersOk

`func (o *QbdPreferences) GetSalesAndCustomersOk() (*TheSalesAndCustomersPreferencesObject, bool)`

GetSalesAndCustomersOk returns a tuple with the SalesAndCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndCustomers

`func (o *QbdPreferences) SetSalesAndCustomers(v TheSalesAndCustomersPreferencesObject)`

SetSalesAndCustomers sets SalesAndCustomers field to given value.


### GetSalesTax

`func (o *QbdPreferences) GetSalesTax() QbdPreferencesSalesTax`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *QbdPreferences) GetSalesTaxOk() (*QbdPreferencesSalesTax, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *QbdPreferences) SetSalesTax(v QbdPreferencesSalesTax)`

SetSalesTax sets SalesTax field to given value.


### GetTimeTracking

`func (o *QbdPreferences) GetTimeTracking() QbdPreferencesTimeTracking`

GetTimeTracking returns the TimeTracking field if non-nil, zero value otherwise.

### GetTimeTrackingOk

`func (o *QbdPreferences) GetTimeTrackingOk() (*QbdPreferencesTimeTracking, bool)`

GetTimeTrackingOk returns a tuple with the TimeTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTracking

`func (o *QbdPreferences) SetTimeTracking(v QbdPreferencesTimeTracking)`

SetTimeTracking sets TimeTracking field to given value.


### GetAppAccessRights

`func (o *QbdPreferences) GetAppAccessRights() TheCurrentAppAccessRightsObject`

GetAppAccessRights returns the AppAccessRights field if non-nil, zero value otherwise.

### GetAppAccessRightsOk

`func (o *QbdPreferences) GetAppAccessRightsOk() (*TheCurrentAppAccessRightsObject, bool)`

GetAppAccessRightsOk returns a tuple with the AppAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessRights

`func (o *QbdPreferences) SetAppAccessRights(v TheCurrentAppAccessRightsObject)`

SetAppAccessRights sets AppAccessRights field to given value.


### GetItemsAndInventory

`func (o *QbdPreferences) GetItemsAndInventory() QbdPreferencesItemsAndInventory`

GetItemsAndInventory returns the ItemsAndInventory field if non-nil, zero value otherwise.

### GetItemsAndInventoryOk

`func (o *QbdPreferences) GetItemsAndInventoryOk() (*QbdPreferencesItemsAndInventory, bool)`

GetItemsAndInventoryOk returns a tuple with the ItemsAndInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsAndInventory

`func (o *QbdPreferences) SetItemsAndInventory(v QbdPreferencesItemsAndInventory)`

SetItemsAndInventory sets ItemsAndInventory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


