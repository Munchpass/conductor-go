/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QbdItemsAndInventoryPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdItemsAndInventoryPreferences{}

// QbdItemsAndInventoryPreferences struct for QbdItemsAndInventoryPreferences
type QbdItemsAndInventoryPreferences struct {
	// Indicates whether enhanced inventory receiving is enabled for this company file.
	IsEnhancedInventoryReceivingEnabled NullableBool `json:"isEnhancedInventoryReceivingEnabled"`
	// Specifies the type of inventory tracking that this company file uses.
	InventoryTrackingMethod NullableString `json:"inventoryTrackingMethod"`
	// Indicates whether expiration dates for inventory serial/lot numbers are enabled for this company file. This feature is supported from QuickBooks Desktop 2023.
	IsInventoryExpirationDateEnabled NullableBool `json:"isInventoryExpirationDateEnabled"`
	// Indicates whether serial/lot number tracking is enabled for sales transactions in this company file.
	IsTrackingOnSalesTransactionsEnabled NullableBool `json:"isTrackingOnSalesTransactionsEnabled"`
	// Indicates whether serial/lot number tracking is enabled for purchase transactions in this company file.
	IsTrackingOnPurchaseTransactionsEnabled NullableBool `json:"isTrackingOnPurchaseTransactionsEnabled"`
	// Indicates whether serial/lot number tracking is enabled for inventory adjustments in this company file.
	IsTrackingOnInventoryAdjustmentEnabled NullableBool `json:"isTrackingOnInventoryAdjustmentEnabled"`
	// Indicates whether serial/lot number tracking is enabled for build assemblies in this company file.
	IsTrackingOnBuildAssemblyEnabled NullableBool `json:"isTrackingOnBuildAssemblyEnabled"`
	// Indicates whether this company file is configured to use FIFO (First In, First Out) to calculate the value of inventory sold and on-hand.
	IsFifoEnabled NullableBool `json:"isFifoEnabled"`
	// The date from which FIFO (First In, First Out) is used to calculate the value of inventory sold and on-hand for this company file, in ISO 8601 format (YYYY-MM-DD).
	FifoEffectiveDate NullableString `json:"fifoEffectiveDate"`
	// Indicates whether bin tracking is enabled for this company file. When `true`, inventory can be tracked by bin locations within sites.
	IsBinTrackingEnabled NullableBool `json:"isBinTrackingEnabled"`
	// Indicates whether barcode functionality is enabled for this company file.
	IsBarcodeEnabled NullableBool `json:"isBarcodeEnabled"`
}

type _QbdItemsAndInventoryPreferences QbdItemsAndInventoryPreferences

// NewQbdItemsAndInventoryPreferences instantiates a new QbdItemsAndInventoryPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdItemsAndInventoryPreferences(isEnhancedInventoryReceivingEnabled NullableBool, inventoryTrackingMethod NullableString, isInventoryExpirationDateEnabled NullableBool, isTrackingOnSalesTransactionsEnabled NullableBool, isTrackingOnPurchaseTransactionsEnabled NullableBool, isTrackingOnInventoryAdjustmentEnabled NullableBool, isTrackingOnBuildAssemblyEnabled NullableBool, isFifoEnabled NullableBool, fifoEffectiveDate NullableString, isBinTrackingEnabled NullableBool, isBarcodeEnabled NullableBool) *QbdItemsAndInventoryPreferences {
	this := QbdItemsAndInventoryPreferences{}
	this.IsEnhancedInventoryReceivingEnabled = isEnhancedInventoryReceivingEnabled
	this.InventoryTrackingMethod = inventoryTrackingMethod
	this.IsInventoryExpirationDateEnabled = isInventoryExpirationDateEnabled
	this.IsTrackingOnSalesTransactionsEnabled = isTrackingOnSalesTransactionsEnabled
	this.IsTrackingOnPurchaseTransactionsEnabled = isTrackingOnPurchaseTransactionsEnabled
	this.IsTrackingOnInventoryAdjustmentEnabled = isTrackingOnInventoryAdjustmentEnabled
	this.IsTrackingOnBuildAssemblyEnabled = isTrackingOnBuildAssemblyEnabled
	this.IsFifoEnabled = isFifoEnabled
	this.FifoEffectiveDate = fifoEffectiveDate
	this.IsBinTrackingEnabled = isBinTrackingEnabled
	this.IsBarcodeEnabled = isBarcodeEnabled
	return &this
}

// NewQbdItemsAndInventoryPreferencesWithDefaults instantiates a new QbdItemsAndInventoryPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdItemsAndInventoryPreferencesWithDefaults() *QbdItemsAndInventoryPreferences {
	this := QbdItemsAndInventoryPreferences{}
	return &this
}

// GetIsEnhancedInventoryReceivingEnabled returns the IsEnhancedInventoryReceivingEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsEnhancedInventoryReceivingEnabled() bool {
	if o == nil || o.IsEnhancedInventoryReceivingEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsEnhancedInventoryReceivingEnabled.Get()
}

// GetIsEnhancedInventoryReceivingEnabledOk returns a tuple with the IsEnhancedInventoryReceivingEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsEnhancedInventoryReceivingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnhancedInventoryReceivingEnabled.Get(), o.IsEnhancedInventoryReceivingEnabled.IsSet()
}

// SetIsEnhancedInventoryReceivingEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsEnhancedInventoryReceivingEnabled(v bool) {
	o.IsEnhancedInventoryReceivingEnabled.Set(&v)
}

// GetInventoryTrackingMethod returns the InventoryTrackingMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdItemsAndInventoryPreferences) GetInventoryTrackingMethod() string {
	if o == nil || o.InventoryTrackingMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.InventoryTrackingMethod.Get()
}

// GetInventoryTrackingMethodOk returns a tuple with the InventoryTrackingMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetInventoryTrackingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryTrackingMethod.Get(), o.InventoryTrackingMethod.IsSet()
}

// SetInventoryTrackingMethod sets field value
func (o *QbdItemsAndInventoryPreferences) SetInventoryTrackingMethod(v string) {
	o.InventoryTrackingMethod.Set(&v)
}

// GetIsInventoryExpirationDateEnabled returns the IsInventoryExpirationDateEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsInventoryExpirationDateEnabled() bool {
	if o == nil || o.IsInventoryExpirationDateEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsInventoryExpirationDateEnabled.Get()
}

// GetIsInventoryExpirationDateEnabledOk returns a tuple with the IsInventoryExpirationDateEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsInventoryExpirationDateEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInventoryExpirationDateEnabled.Get(), o.IsInventoryExpirationDateEnabled.IsSet()
}

// SetIsInventoryExpirationDateEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsInventoryExpirationDateEnabled(v bool) {
	o.IsInventoryExpirationDateEnabled.Set(&v)
}

// GetIsTrackingOnSalesTransactionsEnabled returns the IsTrackingOnSalesTransactionsEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnSalesTransactionsEnabled() bool {
	if o == nil || o.IsTrackingOnSalesTransactionsEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsTrackingOnSalesTransactionsEnabled.Get()
}

// GetIsTrackingOnSalesTransactionsEnabledOk returns a tuple with the IsTrackingOnSalesTransactionsEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnSalesTransactionsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTrackingOnSalesTransactionsEnabled.Get(), o.IsTrackingOnSalesTransactionsEnabled.IsSet()
}

// SetIsTrackingOnSalesTransactionsEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnSalesTransactionsEnabled(v bool) {
	o.IsTrackingOnSalesTransactionsEnabled.Set(&v)
}

// GetIsTrackingOnPurchaseTransactionsEnabled returns the IsTrackingOnPurchaseTransactionsEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnPurchaseTransactionsEnabled() bool {
	if o == nil || o.IsTrackingOnPurchaseTransactionsEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsTrackingOnPurchaseTransactionsEnabled.Get()
}

// GetIsTrackingOnPurchaseTransactionsEnabledOk returns a tuple with the IsTrackingOnPurchaseTransactionsEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnPurchaseTransactionsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTrackingOnPurchaseTransactionsEnabled.Get(), o.IsTrackingOnPurchaseTransactionsEnabled.IsSet()
}

// SetIsTrackingOnPurchaseTransactionsEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnPurchaseTransactionsEnabled(v bool) {
	o.IsTrackingOnPurchaseTransactionsEnabled.Set(&v)
}

// GetIsTrackingOnInventoryAdjustmentEnabled returns the IsTrackingOnInventoryAdjustmentEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnInventoryAdjustmentEnabled() bool {
	if o == nil || o.IsTrackingOnInventoryAdjustmentEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsTrackingOnInventoryAdjustmentEnabled.Get()
}

// GetIsTrackingOnInventoryAdjustmentEnabledOk returns a tuple with the IsTrackingOnInventoryAdjustmentEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnInventoryAdjustmentEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTrackingOnInventoryAdjustmentEnabled.Get(), o.IsTrackingOnInventoryAdjustmentEnabled.IsSet()
}

// SetIsTrackingOnInventoryAdjustmentEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnInventoryAdjustmentEnabled(v bool) {
	o.IsTrackingOnInventoryAdjustmentEnabled.Set(&v)
}

// GetIsTrackingOnBuildAssemblyEnabled returns the IsTrackingOnBuildAssemblyEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnBuildAssemblyEnabled() bool {
	if o == nil || o.IsTrackingOnBuildAssemblyEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsTrackingOnBuildAssemblyEnabled.Get()
}

// GetIsTrackingOnBuildAssemblyEnabledOk returns a tuple with the IsTrackingOnBuildAssemblyEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsTrackingOnBuildAssemblyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTrackingOnBuildAssemblyEnabled.Get(), o.IsTrackingOnBuildAssemblyEnabled.IsSet()
}

// SetIsTrackingOnBuildAssemblyEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsTrackingOnBuildAssemblyEnabled(v bool) {
	o.IsTrackingOnBuildAssemblyEnabled.Set(&v)
}

// GetIsFifoEnabled returns the IsFifoEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsFifoEnabled() bool {
	if o == nil || o.IsFifoEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsFifoEnabled.Get()
}

// GetIsFifoEnabledOk returns a tuple with the IsFifoEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsFifoEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFifoEnabled.Get(), o.IsFifoEnabled.IsSet()
}

// SetIsFifoEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsFifoEnabled(v bool) {
	o.IsFifoEnabled.Set(&v)
}

// GetFifoEffectiveDate returns the FifoEffectiveDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QbdItemsAndInventoryPreferences) GetFifoEffectiveDate() string {
	if o == nil || o.FifoEffectiveDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.FifoEffectiveDate.Get()
}

// GetFifoEffectiveDateOk returns a tuple with the FifoEffectiveDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetFifoEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FifoEffectiveDate.Get(), o.FifoEffectiveDate.IsSet()
}

// SetFifoEffectiveDate sets field value
func (o *QbdItemsAndInventoryPreferences) SetFifoEffectiveDate(v string) {
	o.FifoEffectiveDate.Set(&v)
}

// GetIsBinTrackingEnabled returns the IsBinTrackingEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsBinTrackingEnabled() bool {
	if o == nil || o.IsBinTrackingEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsBinTrackingEnabled.Get()
}

// GetIsBinTrackingEnabledOk returns a tuple with the IsBinTrackingEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsBinTrackingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBinTrackingEnabled.Get(), o.IsBinTrackingEnabled.IsSet()
}

// SetIsBinTrackingEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsBinTrackingEnabled(v bool) {
	o.IsBinTrackingEnabled.Set(&v)
}

// GetIsBarcodeEnabled returns the IsBarcodeEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsBarcodeEnabled() bool {
	if o == nil || o.IsBarcodeEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsBarcodeEnabled.Get()
}

// GetIsBarcodeEnabledOk returns a tuple with the IsBarcodeEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QbdItemsAndInventoryPreferences) GetIsBarcodeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBarcodeEnabled.Get(), o.IsBarcodeEnabled.IsSet()
}

// SetIsBarcodeEnabled sets field value
func (o *QbdItemsAndInventoryPreferences) SetIsBarcodeEnabled(v bool) {
	o.IsBarcodeEnabled.Set(&v)
}

func (o QbdItemsAndInventoryPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdItemsAndInventoryPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isEnhancedInventoryReceivingEnabled"] = o.IsEnhancedInventoryReceivingEnabled.Get()
	toSerialize["inventoryTrackingMethod"] = o.InventoryTrackingMethod.Get()
	toSerialize["isInventoryExpirationDateEnabled"] = o.IsInventoryExpirationDateEnabled.Get()
	toSerialize["isTrackingOnSalesTransactionsEnabled"] = o.IsTrackingOnSalesTransactionsEnabled.Get()
	toSerialize["isTrackingOnPurchaseTransactionsEnabled"] = o.IsTrackingOnPurchaseTransactionsEnabled.Get()
	toSerialize["isTrackingOnInventoryAdjustmentEnabled"] = o.IsTrackingOnInventoryAdjustmentEnabled.Get()
	toSerialize["isTrackingOnBuildAssemblyEnabled"] = o.IsTrackingOnBuildAssemblyEnabled.Get()
	toSerialize["isFifoEnabled"] = o.IsFifoEnabled.Get()
	toSerialize["fifoEffectiveDate"] = o.FifoEffectiveDate.Get()
	toSerialize["isBinTrackingEnabled"] = o.IsBinTrackingEnabled.Get()
	toSerialize["isBarcodeEnabled"] = o.IsBarcodeEnabled.Get()
	return toSerialize, nil
}

func (o *QbdItemsAndInventoryPreferences) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isEnhancedInventoryReceivingEnabled",
		"inventoryTrackingMethod",
		"isInventoryExpirationDateEnabled",
		"isTrackingOnSalesTransactionsEnabled",
		"isTrackingOnPurchaseTransactionsEnabled",
		"isTrackingOnInventoryAdjustmentEnabled",
		"isTrackingOnBuildAssemblyEnabled",
		"isFifoEnabled",
		"fifoEffectiveDate",
		"isBinTrackingEnabled",
		"isBarcodeEnabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQbdItemsAndInventoryPreferences := _QbdItemsAndInventoryPreferences{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbdItemsAndInventoryPreferences)

	if err != nil {
		return err
	}

	*o = QbdItemsAndInventoryPreferences(varQbdItemsAndInventoryPreferences)

	return err
}

type NullableQbdItemsAndInventoryPreferences struct {
	value *QbdItemsAndInventoryPreferences
	isSet bool
}

func (v NullableQbdItemsAndInventoryPreferences) Get() *QbdItemsAndInventoryPreferences {
	return v.value
}

func (v *NullableQbdItemsAndInventoryPreferences) Set(val *QbdItemsAndInventoryPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdItemsAndInventoryPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdItemsAndInventoryPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdItemsAndInventoryPreferences(val *QbdItemsAndInventoryPreferences) *NullableQbdItemsAndInventoryPreferences {
	return &NullableQbdItemsAndInventoryPreferences{value: val, isSet: true}
}

func (v NullableQbdItemsAndInventoryPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdItemsAndInventoryPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


