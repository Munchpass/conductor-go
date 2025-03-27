# TheSalesAndCustomersPreferencesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultShippingMethod** | [**TheSalesAndCustomersPreferencesObjectDefaultShippingMethod**](TheSalesAndCustomersPreferencesObjectDefaultShippingMethod.md) |  | 
**DefaultShipmentOrigin** | **string** | The default shipment-origin location (i.e., FOB - freight on board) from which invoiced products are shipped for this company file. This indicates the point at which ownership and liability for goods transfer from seller to buyer. | 
**DefaultMarkupPercentage** | **string** | The default percentage that an inventory item will be marked up from its cost for this company file. | 
**IsTrackingReimbursedExpensesAsIncome** | **bool** | Indicates whether this company file is configured to track an expense and the customer&#39;s reimbursement for that expense in separate accounts. When &#x60;true&#x60;, reimbursements can be tracked as income rather than as a reduction of the original expense. | 
**IsAutoApplyingPayments** | **bool** | Indicates whether this company file is configured to automatically apply a customer&#39;s payment to their outstanding invoices, beginning with the oldest invoice. | 
**PriceLevels** | [**TheSalesAndCustomersPreferencesObjectPriceLevels**](TheSalesAndCustomersPreferencesObjectPriceLevels.md) |  | 

## Methods

### NewTheSalesAndCustomersPreferencesObject

`func NewTheSalesAndCustomersPreferencesObject(defaultShippingMethod TheSalesAndCustomersPreferencesObjectDefaultShippingMethod, defaultShipmentOrigin string, defaultMarkupPercentage string, isTrackingReimbursedExpensesAsIncome bool, isAutoApplyingPayments bool, priceLevels TheSalesAndCustomersPreferencesObjectPriceLevels, ) *TheSalesAndCustomersPreferencesObject`

NewTheSalesAndCustomersPreferencesObject instantiates a new TheSalesAndCustomersPreferencesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTheSalesAndCustomersPreferencesObjectWithDefaults

`func NewTheSalesAndCustomersPreferencesObjectWithDefaults() *TheSalesAndCustomersPreferencesObject`

NewTheSalesAndCustomersPreferencesObjectWithDefaults instantiates a new TheSalesAndCustomersPreferencesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultShippingMethod

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShippingMethod() TheSalesAndCustomersPreferencesObjectDefaultShippingMethod`

GetDefaultShippingMethod returns the DefaultShippingMethod field if non-nil, zero value otherwise.

### GetDefaultShippingMethodOk

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShippingMethodOk() (*TheSalesAndCustomersPreferencesObjectDefaultShippingMethod, bool)`

GetDefaultShippingMethodOk returns a tuple with the DefaultShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingMethod

`func (o *TheSalesAndCustomersPreferencesObject) SetDefaultShippingMethod(v TheSalesAndCustomersPreferencesObjectDefaultShippingMethod)`

SetDefaultShippingMethod sets DefaultShippingMethod field to given value.


### GetDefaultShipmentOrigin

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShipmentOrigin() string`

GetDefaultShipmentOrigin returns the DefaultShipmentOrigin field if non-nil, zero value otherwise.

### GetDefaultShipmentOriginOk

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultShipmentOriginOk() (*string, bool)`

GetDefaultShipmentOriginOk returns a tuple with the DefaultShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShipmentOrigin

`func (o *TheSalesAndCustomersPreferencesObject) SetDefaultShipmentOrigin(v string)`

SetDefaultShipmentOrigin sets DefaultShipmentOrigin field to given value.


### GetDefaultMarkupPercentage

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultMarkupPercentage() string`

GetDefaultMarkupPercentage returns the DefaultMarkupPercentage field if non-nil, zero value otherwise.

### GetDefaultMarkupPercentageOk

`func (o *TheSalesAndCustomersPreferencesObject) GetDefaultMarkupPercentageOk() (*string, bool)`

GetDefaultMarkupPercentageOk returns a tuple with the DefaultMarkupPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMarkupPercentage

`func (o *TheSalesAndCustomersPreferencesObject) SetDefaultMarkupPercentage(v string)`

SetDefaultMarkupPercentage sets DefaultMarkupPercentage field to given value.


### GetIsTrackingReimbursedExpensesAsIncome

`func (o *TheSalesAndCustomersPreferencesObject) GetIsTrackingReimbursedExpensesAsIncome() bool`

GetIsTrackingReimbursedExpensesAsIncome returns the IsTrackingReimbursedExpensesAsIncome field if non-nil, zero value otherwise.

### GetIsTrackingReimbursedExpensesAsIncomeOk

`func (o *TheSalesAndCustomersPreferencesObject) GetIsTrackingReimbursedExpensesAsIncomeOk() (*bool, bool)`

GetIsTrackingReimbursedExpensesAsIncomeOk returns a tuple with the IsTrackingReimbursedExpensesAsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingReimbursedExpensesAsIncome

`func (o *TheSalesAndCustomersPreferencesObject) SetIsTrackingReimbursedExpensesAsIncome(v bool)`

SetIsTrackingReimbursedExpensesAsIncome sets IsTrackingReimbursedExpensesAsIncome field to given value.


### GetIsAutoApplyingPayments

`func (o *TheSalesAndCustomersPreferencesObject) GetIsAutoApplyingPayments() bool`

GetIsAutoApplyingPayments returns the IsAutoApplyingPayments field if non-nil, zero value otherwise.

### GetIsAutoApplyingPaymentsOk

`func (o *TheSalesAndCustomersPreferencesObject) GetIsAutoApplyingPaymentsOk() (*bool, bool)`

GetIsAutoApplyingPaymentsOk returns a tuple with the IsAutoApplyingPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoApplyingPayments

`func (o *TheSalesAndCustomersPreferencesObject) SetIsAutoApplyingPayments(v bool)`

SetIsAutoApplyingPayments sets IsAutoApplyingPayments field to given value.


### GetPriceLevels

`func (o *TheSalesAndCustomersPreferencesObject) GetPriceLevels() TheSalesAndCustomersPreferencesObjectPriceLevels`

GetPriceLevels returns the PriceLevels field if non-nil, zero value otherwise.

### GetPriceLevelsOk

`func (o *TheSalesAndCustomersPreferencesObject) GetPriceLevelsOk() (*TheSalesAndCustomersPreferencesObjectPriceLevels, bool)`

GetPriceLevelsOk returns a tuple with the PriceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevels

`func (o *TheSalesAndCustomersPreferencesObject) SetPriceLevels(v TheSalesAndCustomersPreferencesObjectPriceLevels)`

SetPriceLevels sets PriceLevels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


