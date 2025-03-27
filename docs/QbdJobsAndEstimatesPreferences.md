# QbdJobsAndEstimatesPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsingEstimates** | **bool** | Indicates whether this company file is configured to create estimates for jobs. | 
**IsUsingProgressInvoicing** | **bool** | Indicates whether this company file permits creating invoices for only a portion of an estimate. | 
**IsPrintingItemsWithZeroAmounts** | **bool** | Indicates whether this company file is configured to print line items with zero amounts on progress invoices. This preference is only relevant if &#x60;isUsingProgressInvoicing&#x60; is &#x60;true&#x60;. | 

## Methods

### NewQbdJobsAndEstimatesPreferences

`func NewQbdJobsAndEstimatesPreferences(isUsingEstimates bool, isUsingProgressInvoicing bool, isPrintingItemsWithZeroAmounts bool, ) *QbdJobsAndEstimatesPreferences`

NewQbdJobsAndEstimatesPreferences instantiates a new QbdJobsAndEstimatesPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdJobsAndEstimatesPreferencesWithDefaults

`func NewQbdJobsAndEstimatesPreferencesWithDefaults() *QbdJobsAndEstimatesPreferences`

NewQbdJobsAndEstimatesPreferencesWithDefaults instantiates a new QbdJobsAndEstimatesPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsingEstimates

`func (o *QbdJobsAndEstimatesPreferences) GetIsUsingEstimates() bool`

GetIsUsingEstimates returns the IsUsingEstimates field if non-nil, zero value otherwise.

### GetIsUsingEstimatesOk

`func (o *QbdJobsAndEstimatesPreferences) GetIsUsingEstimatesOk() (*bool, bool)`

GetIsUsingEstimatesOk returns a tuple with the IsUsingEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingEstimates

`func (o *QbdJobsAndEstimatesPreferences) SetIsUsingEstimates(v bool)`

SetIsUsingEstimates sets IsUsingEstimates field to given value.


### GetIsUsingProgressInvoicing

`func (o *QbdJobsAndEstimatesPreferences) GetIsUsingProgressInvoicing() bool`

GetIsUsingProgressInvoicing returns the IsUsingProgressInvoicing field if non-nil, zero value otherwise.

### GetIsUsingProgressInvoicingOk

`func (o *QbdJobsAndEstimatesPreferences) GetIsUsingProgressInvoicingOk() (*bool, bool)`

GetIsUsingProgressInvoicingOk returns a tuple with the IsUsingProgressInvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingProgressInvoicing

`func (o *QbdJobsAndEstimatesPreferences) SetIsUsingProgressInvoicing(v bool)`

SetIsUsingProgressInvoicing sets IsUsingProgressInvoicing field to given value.


### GetIsPrintingItemsWithZeroAmounts

`func (o *QbdJobsAndEstimatesPreferences) GetIsPrintingItemsWithZeroAmounts() bool`

GetIsPrintingItemsWithZeroAmounts returns the IsPrintingItemsWithZeroAmounts field if non-nil, zero value otherwise.

### GetIsPrintingItemsWithZeroAmountsOk

`func (o *QbdJobsAndEstimatesPreferences) GetIsPrintingItemsWithZeroAmountsOk() (*bool, bool)`

GetIsPrintingItemsWithZeroAmountsOk returns a tuple with the IsPrintingItemsWithZeroAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrintingItemsWithZeroAmounts

`func (o *QbdJobsAndEstimatesPreferences) SetIsPrintingItemsWithZeroAmounts(v bool)`

SetIsPrintingItemsWithZeroAmounts sets IsPrintingItemsWithZeroAmounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


