# QbdPreferencesItemsAndInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnhancedInventoryReceivingEnabled** | **bool** | Indicates whether enhanced inventory receiving is enabled for this company file. | 
**InventoryTrackingMethod** | **string** | Specifies the type of inventory tracking that this company file uses. | 
**IsInventoryExpirationDateEnabled** | **bool** | Indicates whether expiration dates for inventory serial/lot numbers are enabled for this company file. This feature is supported from QuickBooks Desktop 2023. | 
**IsTrackingOnSalesTransactionsEnabled** | **bool** | Indicates whether serial/lot number tracking is enabled for sales transactions in this company file. | 
**IsTrackingOnPurchaseTransactionsEnabled** | **bool** | Indicates whether serial/lot number tracking is enabled for purchase transactions in this company file. | 
**IsTrackingOnInventoryAdjustmentEnabled** | **bool** | Indicates whether serial/lot number tracking is enabled for inventory adjustments in this company file. | 
**IsTrackingOnBuildAssemblyEnabled** | **bool** | Indicates whether serial/lot number tracking is enabled for build assemblies in this company file. | 
**IsFifoEnabled** | **bool** | Indicates whether this company file is configured to use FIFO (First In, First Out) to calculate the value of inventory sold and on-hand. | 
**FifoEffectiveDate** | **string** | The date from which FIFO (First In, First Out) is used to calculate the value of inventory sold and on-hand for this company file, in ISO 8601 format (YYYY-MM-DD). | 
**IsBinTrackingEnabled** | **bool** | Indicates whether bin tracking is enabled for this company file. When &#x60;true&#x60;, inventory can be tracked by bin locations within sites. | 
**IsBarcodeEnabled** | **bool** | Indicates whether barcode functionality is enabled for this company file. | 

## Methods

### NewQbdPreferencesItemsAndInventory

`func NewQbdPreferencesItemsAndInventory(isEnhancedInventoryReceivingEnabled bool, inventoryTrackingMethod string, isInventoryExpirationDateEnabled bool, isTrackingOnSalesTransactionsEnabled bool, isTrackingOnPurchaseTransactionsEnabled bool, isTrackingOnInventoryAdjustmentEnabled bool, isTrackingOnBuildAssemblyEnabled bool, isFifoEnabled bool, fifoEffectiveDate string, isBinTrackingEnabled bool, isBarcodeEnabled bool, ) *QbdPreferencesItemsAndInventory`

NewQbdPreferencesItemsAndInventory instantiates a new QbdPreferencesItemsAndInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPreferencesItemsAndInventoryWithDefaults

`func NewQbdPreferencesItemsAndInventoryWithDefaults() *QbdPreferencesItemsAndInventory`

NewQbdPreferencesItemsAndInventoryWithDefaults instantiates a new QbdPreferencesItemsAndInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnhancedInventoryReceivingEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsEnhancedInventoryReceivingEnabled() bool`

GetIsEnhancedInventoryReceivingEnabled returns the IsEnhancedInventoryReceivingEnabled field if non-nil, zero value otherwise.

### GetIsEnhancedInventoryReceivingEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsEnhancedInventoryReceivingEnabledOk() (*bool, bool)`

GetIsEnhancedInventoryReceivingEnabledOk returns a tuple with the IsEnhancedInventoryReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnhancedInventoryReceivingEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsEnhancedInventoryReceivingEnabled(v bool)`

SetIsEnhancedInventoryReceivingEnabled sets IsEnhancedInventoryReceivingEnabled field to given value.


### GetInventoryTrackingMethod

`func (o *QbdPreferencesItemsAndInventory) GetInventoryTrackingMethod() string`

GetInventoryTrackingMethod returns the InventoryTrackingMethod field if non-nil, zero value otherwise.

### GetInventoryTrackingMethodOk

`func (o *QbdPreferencesItemsAndInventory) GetInventoryTrackingMethodOk() (*string, bool)`

GetInventoryTrackingMethodOk returns a tuple with the InventoryTrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTrackingMethod

`func (o *QbdPreferencesItemsAndInventory) SetInventoryTrackingMethod(v string)`

SetInventoryTrackingMethod sets InventoryTrackingMethod field to given value.


### GetIsInventoryExpirationDateEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsInventoryExpirationDateEnabled() bool`

GetIsInventoryExpirationDateEnabled returns the IsInventoryExpirationDateEnabled field if non-nil, zero value otherwise.

### GetIsInventoryExpirationDateEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsInventoryExpirationDateEnabledOk() (*bool, bool)`

GetIsInventoryExpirationDateEnabledOk returns a tuple with the IsInventoryExpirationDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInventoryExpirationDateEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsInventoryExpirationDateEnabled(v bool)`

SetIsInventoryExpirationDateEnabled sets IsInventoryExpirationDateEnabled field to given value.


### GetIsTrackingOnSalesTransactionsEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnSalesTransactionsEnabled() bool`

GetIsTrackingOnSalesTransactionsEnabled returns the IsTrackingOnSalesTransactionsEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnSalesTransactionsEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnSalesTransactionsEnabledOk() (*bool, bool)`

GetIsTrackingOnSalesTransactionsEnabledOk returns a tuple with the IsTrackingOnSalesTransactionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnSalesTransactionsEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsTrackingOnSalesTransactionsEnabled(v bool)`

SetIsTrackingOnSalesTransactionsEnabled sets IsTrackingOnSalesTransactionsEnabled field to given value.


### GetIsTrackingOnPurchaseTransactionsEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnPurchaseTransactionsEnabled() bool`

GetIsTrackingOnPurchaseTransactionsEnabled returns the IsTrackingOnPurchaseTransactionsEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnPurchaseTransactionsEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnPurchaseTransactionsEnabledOk() (*bool, bool)`

GetIsTrackingOnPurchaseTransactionsEnabledOk returns a tuple with the IsTrackingOnPurchaseTransactionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnPurchaseTransactionsEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsTrackingOnPurchaseTransactionsEnabled(v bool)`

SetIsTrackingOnPurchaseTransactionsEnabled sets IsTrackingOnPurchaseTransactionsEnabled field to given value.


### GetIsTrackingOnInventoryAdjustmentEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnInventoryAdjustmentEnabled() bool`

GetIsTrackingOnInventoryAdjustmentEnabled returns the IsTrackingOnInventoryAdjustmentEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnInventoryAdjustmentEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnInventoryAdjustmentEnabledOk() (*bool, bool)`

GetIsTrackingOnInventoryAdjustmentEnabledOk returns a tuple with the IsTrackingOnInventoryAdjustmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnInventoryAdjustmentEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsTrackingOnInventoryAdjustmentEnabled(v bool)`

SetIsTrackingOnInventoryAdjustmentEnabled sets IsTrackingOnInventoryAdjustmentEnabled field to given value.


### GetIsTrackingOnBuildAssemblyEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnBuildAssemblyEnabled() bool`

GetIsTrackingOnBuildAssemblyEnabled returns the IsTrackingOnBuildAssemblyEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnBuildAssemblyEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsTrackingOnBuildAssemblyEnabledOk() (*bool, bool)`

GetIsTrackingOnBuildAssemblyEnabledOk returns a tuple with the IsTrackingOnBuildAssemblyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnBuildAssemblyEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsTrackingOnBuildAssemblyEnabled(v bool)`

SetIsTrackingOnBuildAssemblyEnabled sets IsTrackingOnBuildAssemblyEnabled field to given value.


### GetIsFifoEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsFifoEnabled() bool`

GetIsFifoEnabled returns the IsFifoEnabled field if non-nil, zero value otherwise.

### GetIsFifoEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsFifoEnabledOk() (*bool, bool)`

GetIsFifoEnabledOk returns a tuple with the IsFifoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFifoEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsFifoEnabled(v bool)`

SetIsFifoEnabled sets IsFifoEnabled field to given value.


### GetFifoEffectiveDate

`func (o *QbdPreferencesItemsAndInventory) GetFifoEffectiveDate() string`

GetFifoEffectiveDate returns the FifoEffectiveDate field if non-nil, zero value otherwise.

### GetFifoEffectiveDateOk

`func (o *QbdPreferencesItemsAndInventory) GetFifoEffectiveDateOk() (*string, bool)`

GetFifoEffectiveDateOk returns a tuple with the FifoEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifoEffectiveDate

`func (o *QbdPreferencesItemsAndInventory) SetFifoEffectiveDate(v string)`

SetFifoEffectiveDate sets FifoEffectiveDate field to given value.


### GetIsBinTrackingEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsBinTrackingEnabled() bool`

GetIsBinTrackingEnabled returns the IsBinTrackingEnabled field if non-nil, zero value otherwise.

### GetIsBinTrackingEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsBinTrackingEnabledOk() (*bool, bool)`

GetIsBinTrackingEnabledOk returns a tuple with the IsBinTrackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBinTrackingEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsBinTrackingEnabled(v bool)`

SetIsBinTrackingEnabled sets IsBinTrackingEnabled field to given value.


### GetIsBarcodeEnabled

`func (o *QbdPreferencesItemsAndInventory) GetIsBarcodeEnabled() bool`

GetIsBarcodeEnabled returns the IsBarcodeEnabled field if non-nil, zero value otherwise.

### GetIsBarcodeEnabledOk

`func (o *QbdPreferencesItemsAndInventory) GetIsBarcodeEnabledOk() (*bool, bool)`

GetIsBarcodeEnabledOk returns a tuple with the IsBarcodeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBarcodeEnabled

`func (o *QbdPreferencesItemsAndInventory) SetIsBarcodeEnabled(v bool)`

SetIsBarcodeEnabled sets IsBarcodeEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


