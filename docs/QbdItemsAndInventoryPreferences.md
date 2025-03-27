# QbdItemsAndInventoryPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnhancedInventoryReceivingEnabled** | **NullableBool** | Indicates whether enhanced inventory receiving is enabled for this company file. | 
**InventoryTrackingMethod** | **NullableString** | Specifies the type of inventory tracking that this company file uses. | 
**IsInventoryExpirationDateEnabled** | **NullableBool** | Indicates whether expiration dates for inventory serial/lot numbers are enabled for this company file. This feature is supported from QuickBooks Desktop 2023. | 
**IsTrackingOnSalesTransactionsEnabled** | **NullableBool** | Indicates whether serial/lot number tracking is enabled for sales transactions in this company file. | 
**IsTrackingOnPurchaseTransactionsEnabled** | **NullableBool** | Indicates whether serial/lot number tracking is enabled for purchase transactions in this company file. | 
**IsTrackingOnInventoryAdjustmentEnabled** | **NullableBool** | Indicates whether serial/lot number tracking is enabled for inventory adjustments in this company file. | 
**IsTrackingOnBuildAssemblyEnabled** | **NullableBool** | Indicates whether serial/lot number tracking is enabled for build assemblies in this company file. | 
**IsFifoEnabled** | **NullableBool** | Indicates whether this company file is configured to use FIFO (First In, First Out) to calculate the value of inventory sold and on-hand. | 
**FifoEffectiveDate** | **NullableString** | The date from which FIFO (First In, First Out) is used to calculate the value of inventory sold and on-hand for this company file, in ISO 8601 format (YYYY-MM-DD). | 
**IsBinTrackingEnabled** | **NullableBool** | Indicates whether bin tracking is enabled for this company file. When &#x60;true&#x60;, inventory can be tracked by bin locations within sites. | 
**IsBarcodeEnabled** | **NullableBool** | Indicates whether barcode functionality is enabled for this company file. | 

## Methods

### NewQbdItemsAndInventoryPreferences

`func NewQbdItemsAndInventoryPreferences(isEnhancedInventoryReceivingEnabled NullableBool, inventoryTrackingMethod NullableString, isInventoryExpirationDateEnabled NullableBool, isTrackingOnSalesTransactionsEnabled NullableBool, isTrackingOnPurchaseTransactionsEnabled NullableBool, isTrackingOnInventoryAdjustmentEnabled NullableBool, isTrackingOnBuildAssemblyEnabled NullableBool, isFifoEnabled NullableBool, fifoEffectiveDate NullableString, isBinTrackingEnabled NullableBool, isBarcodeEnabled NullableBool, ) *QbdItemsAndInventoryPreferences`

NewQbdItemsAndInventoryPreferences instantiates a new QbdItemsAndInventoryPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdItemsAndInventoryPreferencesWithDefaults

`func NewQbdItemsAndInventoryPreferencesWithDefaults() *QbdItemsAndInventoryPreferences`

NewQbdItemsAndInventoryPreferencesWithDefaults instantiates a new QbdItemsAndInventoryPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnhancedInventoryReceivingEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsEnhancedInventoryReceivingEnabled() bool`

GetIsEnhancedInventoryReceivingEnabled returns the IsEnhancedInventoryReceivingEnabled field if non-nil, zero value otherwise.

### GetIsEnhancedInventoryReceivingEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsEnhancedInventoryReceivingEnabledOk() (*bool, bool)`

GetIsEnhancedInventoryReceivingEnabledOk returns a tuple with the IsEnhancedInventoryReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnhancedInventoryReceivingEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsEnhancedInventoryReceivingEnabled(v bool)`

SetIsEnhancedInventoryReceivingEnabled sets IsEnhancedInventoryReceivingEnabled field to given value.


### SetIsEnhancedInventoryReceivingEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsEnhancedInventoryReceivingEnabledNil(b bool)`

 SetIsEnhancedInventoryReceivingEnabledNil sets the value for IsEnhancedInventoryReceivingEnabled to be an explicit nil

### UnsetIsEnhancedInventoryReceivingEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsEnhancedInventoryReceivingEnabled()`

UnsetIsEnhancedInventoryReceivingEnabled ensures that no value is present for IsEnhancedInventoryReceivingEnabled, not even an explicit nil
### GetInventoryTrackingMethod

`func (o *QbdItemsAndInventoryPreferences) GetInventoryTrackingMethod() string`

GetInventoryTrackingMethod returns the InventoryTrackingMethod field if non-nil, zero value otherwise.

### GetInventoryTrackingMethodOk

`func (o *QbdItemsAndInventoryPreferences) GetInventoryTrackingMethodOk() (*string, bool)`

GetInventoryTrackingMethodOk returns a tuple with the InventoryTrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTrackingMethod

`func (o *QbdItemsAndInventoryPreferences) SetInventoryTrackingMethod(v string)`

SetInventoryTrackingMethod sets InventoryTrackingMethod field to given value.


### SetInventoryTrackingMethodNil

`func (o *QbdItemsAndInventoryPreferences) SetInventoryTrackingMethodNil(b bool)`

 SetInventoryTrackingMethodNil sets the value for InventoryTrackingMethod to be an explicit nil

### UnsetInventoryTrackingMethod
`func (o *QbdItemsAndInventoryPreferences) UnsetInventoryTrackingMethod()`

UnsetInventoryTrackingMethod ensures that no value is present for InventoryTrackingMethod, not even an explicit nil
### GetIsInventoryExpirationDateEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsInventoryExpirationDateEnabled() bool`

GetIsInventoryExpirationDateEnabled returns the IsInventoryExpirationDateEnabled field if non-nil, zero value otherwise.

### GetIsInventoryExpirationDateEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsInventoryExpirationDateEnabledOk() (*bool, bool)`

GetIsInventoryExpirationDateEnabledOk returns a tuple with the IsInventoryExpirationDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInventoryExpirationDateEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsInventoryExpirationDateEnabled(v bool)`

SetIsInventoryExpirationDateEnabled sets IsInventoryExpirationDateEnabled field to given value.


### SetIsInventoryExpirationDateEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsInventoryExpirationDateEnabledNil(b bool)`

 SetIsInventoryExpirationDateEnabledNil sets the value for IsInventoryExpirationDateEnabled to be an explicit nil

### UnsetIsInventoryExpirationDateEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsInventoryExpirationDateEnabled()`

UnsetIsInventoryExpirationDateEnabled ensures that no value is present for IsInventoryExpirationDateEnabled, not even an explicit nil
### GetIsTrackingOnSalesTransactionsEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnSalesTransactionsEnabled() bool`

GetIsTrackingOnSalesTransactionsEnabled returns the IsTrackingOnSalesTransactionsEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnSalesTransactionsEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnSalesTransactionsEnabledOk() (*bool, bool)`

GetIsTrackingOnSalesTransactionsEnabledOk returns a tuple with the IsTrackingOnSalesTransactionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnSalesTransactionsEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnSalesTransactionsEnabled(v bool)`

SetIsTrackingOnSalesTransactionsEnabled sets IsTrackingOnSalesTransactionsEnabled field to given value.


### SetIsTrackingOnSalesTransactionsEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnSalesTransactionsEnabledNil(b bool)`

 SetIsTrackingOnSalesTransactionsEnabledNil sets the value for IsTrackingOnSalesTransactionsEnabled to be an explicit nil

### UnsetIsTrackingOnSalesTransactionsEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsTrackingOnSalesTransactionsEnabled()`

UnsetIsTrackingOnSalesTransactionsEnabled ensures that no value is present for IsTrackingOnSalesTransactionsEnabled, not even an explicit nil
### GetIsTrackingOnPurchaseTransactionsEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnPurchaseTransactionsEnabled() bool`

GetIsTrackingOnPurchaseTransactionsEnabled returns the IsTrackingOnPurchaseTransactionsEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnPurchaseTransactionsEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnPurchaseTransactionsEnabledOk() (*bool, bool)`

GetIsTrackingOnPurchaseTransactionsEnabledOk returns a tuple with the IsTrackingOnPurchaseTransactionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnPurchaseTransactionsEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnPurchaseTransactionsEnabled(v bool)`

SetIsTrackingOnPurchaseTransactionsEnabled sets IsTrackingOnPurchaseTransactionsEnabled field to given value.


### SetIsTrackingOnPurchaseTransactionsEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnPurchaseTransactionsEnabledNil(b bool)`

 SetIsTrackingOnPurchaseTransactionsEnabledNil sets the value for IsTrackingOnPurchaseTransactionsEnabled to be an explicit nil

### UnsetIsTrackingOnPurchaseTransactionsEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsTrackingOnPurchaseTransactionsEnabled()`

UnsetIsTrackingOnPurchaseTransactionsEnabled ensures that no value is present for IsTrackingOnPurchaseTransactionsEnabled, not even an explicit nil
### GetIsTrackingOnInventoryAdjustmentEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnInventoryAdjustmentEnabled() bool`

GetIsTrackingOnInventoryAdjustmentEnabled returns the IsTrackingOnInventoryAdjustmentEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnInventoryAdjustmentEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnInventoryAdjustmentEnabledOk() (*bool, bool)`

GetIsTrackingOnInventoryAdjustmentEnabledOk returns a tuple with the IsTrackingOnInventoryAdjustmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnInventoryAdjustmentEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnInventoryAdjustmentEnabled(v bool)`

SetIsTrackingOnInventoryAdjustmentEnabled sets IsTrackingOnInventoryAdjustmentEnabled field to given value.


### SetIsTrackingOnInventoryAdjustmentEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnInventoryAdjustmentEnabledNil(b bool)`

 SetIsTrackingOnInventoryAdjustmentEnabledNil sets the value for IsTrackingOnInventoryAdjustmentEnabled to be an explicit nil

### UnsetIsTrackingOnInventoryAdjustmentEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsTrackingOnInventoryAdjustmentEnabled()`

UnsetIsTrackingOnInventoryAdjustmentEnabled ensures that no value is present for IsTrackingOnInventoryAdjustmentEnabled, not even an explicit nil
### GetIsTrackingOnBuildAssemblyEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnBuildAssemblyEnabled() bool`

GetIsTrackingOnBuildAssemblyEnabled returns the IsTrackingOnBuildAssemblyEnabled field if non-nil, zero value otherwise.

### GetIsTrackingOnBuildAssemblyEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnBuildAssemblyEnabledOk() (*bool, bool)`

GetIsTrackingOnBuildAssemblyEnabledOk returns a tuple with the IsTrackingOnBuildAssemblyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOnBuildAssemblyEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnBuildAssemblyEnabled(v bool)`

SetIsTrackingOnBuildAssemblyEnabled sets IsTrackingOnBuildAssemblyEnabled field to given value.


### SetIsTrackingOnBuildAssemblyEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnBuildAssemblyEnabledNil(b bool)`

 SetIsTrackingOnBuildAssemblyEnabledNil sets the value for IsTrackingOnBuildAssemblyEnabled to be an explicit nil

### UnsetIsTrackingOnBuildAssemblyEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsTrackingOnBuildAssemblyEnabled()`

UnsetIsTrackingOnBuildAssemblyEnabled ensures that no value is present for IsTrackingOnBuildAssemblyEnabled, not even an explicit nil
### GetIsFifoEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsFifoEnabled() bool`

GetIsFifoEnabled returns the IsFifoEnabled field if non-nil, zero value otherwise.

### GetIsFifoEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsFifoEnabledOk() (*bool, bool)`

GetIsFifoEnabledOk returns a tuple with the IsFifoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFifoEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsFifoEnabled(v bool)`

SetIsFifoEnabled sets IsFifoEnabled field to given value.


### SetIsFifoEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsFifoEnabledNil(b bool)`

 SetIsFifoEnabledNil sets the value for IsFifoEnabled to be an explicit nil

### UnsetIsFifoEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsFifoEnabled()`

UnsetIsFifoEnabled ensures that no value is present for IsFifoEnabled, not even an explicit nil
### GetFifoEffectiveDate

`func (o *QbdItemsAndInventoryPreferences) GetFifoEffectiveDate() string`

GetFifoEffectiveDate returns the FifoEffectiveDate field if non-nil, zero value otherwise.

### GetFifoEffectiveDateOk

`func (o *QbdItemsAndInventoryPreferences) GetFifoEffectiveDateOk() (*string, bool)`

GetFifoEffectiveDateOk returns a tuple with the FifoEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifoEffectiveDate

`func (o *QbdItemsAndInventoryPreferences) SetFifoEffectiveDate(v string)`

SetFifoEffectiveDate sets FifoEffectiveDate field to given value.


### SetFifoEffectiveDateNil

`func (o *QbdItemsAndInventoryPreferences) SetFifoEffectiveDateNil(b bool)`

 SetFifoEffectiveDateNil sets the value for FifoEffectiveDate to be an explicit nil

### UnsetFifoEffectiveDate
`func (o *QbdItemsAndInventoryPreferences) UnsetFifoEffectiveDate()`

UnsetFifoEffectiveDate ensures that no value is present for FifoEffectiveDate, not even an explicit nil
### GetIsBinTrackingEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsBinTrackingEnabled() bool`

GetIsBinTrackingEnabled returns the IsBinTrackingEnabled field if non-nil, zero value otherwise.

### GetIsBinTrackingEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsBinTrackingEnabledOk() (*bool, bool)`

GetIsBinTrackingEnabledOk returns a tuple with the IsBinTrackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBinTrackingEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsBinTrackingEnabled(v bool)`

SetIsBinTrackingEnabled sets IsBinTrackingEnabled field to given value.


### SetIsBinTrackingEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsBinTrackingEnabledNil(b bool)`

 SetIsBinTrackingEnabledNil sets the value for IsBinTrackingEnabled to be an explicit nil

### UnsetIsBinTrackingEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsBinTrackingEnabled()`

UnsetIsBinTrackingEnabled ensures that no value is present for IsBinTrackingEnabled, not even an explicit nil
### GetIsBarcodeEnabled

`func (o *QbdItemsAndInventoryPreferences) GetIsBarcodeEnabled() bool`

GetIsBarcodeEnabled returns the IsBarcodeEnabled field if non-nil, zero value otherwise.

### GetIsBarcodeEnabledOk

`func (o *QbdItemsAndInventoryPreferences) GetIsBarcodeEnabledOk() (*bool, bool)`

GetIsBarcodeEnabledOk returns a tuple with the IsBarcodeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBarcodeEnabled

`func (o *QbdItemsAndInventoryPreferences) SetIsBarcodeEnabled(v bool)`

SetIsBarcodeEnabled sets IsBarcodeEnabled field to given value.


### SetIsBarcodeEnabledNil

`func (o *QbdItemsAndInventoryPreferences) SetIsBarcodeEnabledNil(b bool)`

 SetIsBarcodeEnabledNil sets the value for IsBarcodeEnabled to be an explicit nil

### UnsetIsBarcodeEnabled
`func (o *QbdItemsAndInventoryPreferences) UnsetIsBarcodeEnabled()`

UnsetIsBarcodeEnabled ensures that no value is present for IsBarcodeEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


