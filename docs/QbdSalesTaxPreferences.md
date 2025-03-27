# QbdSalesTaxPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultItemSalesTax** | [**QbdSalesTaxPreferencesDefaultItemSalesTax**](QbdSalesTaxPreferencesDefaultItemSalesTax.md) |  | 
**SalesTaxReportingFrequency** | **string** | The frequency at which sales tax reports are generated for this company file. | 
**DefaultTaxableSalesTaxCode** | [**QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode**](QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode.md) |  | 
**DefaultNonTaxableSalesTaxCode** | [**QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode**](QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode.md) |  | 
**IsUsingVendorTaxCode** | **bool** | Indicates whether this company file is configured to use tax codes for vendors. | 
**IsUsingCustomerTaxCode** | **bool** | Indicates whether this company file is configured to use tax codes for customers. | 
**IsUsingTaxInclusivePrices** | **bool** | Indicates whether this company file is configured to allow tax-inclusive prices. | 

## Methods

### NewQbdSalesTaxPreferences

`func NewQbdSalesTaxPreferences(defaultItemSalesTax QbdSalesTaxPreferencesDefaultItemSalesTax, salesTaxReportingFrequency string, defaultTaxableSalesTaxCode QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode, defaultNonTaxableSalesTaxCode QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode, isUsingVendorTaxCode bool, isUsingCustomerTaxCode bool, isUsingTaxInclusivePrices bool, ) *QbdSalesTaxPreferences`

NewQbdSalesTaxPreferences instantiates a new QbdSalesTaxPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesTaxPreferencesWithDefaults

`func NewQbdSalesTaxPreferencesWithDefaults() *QbdSalesTaxPreferences`

NewQbdSalesTaxPreferencesWithDefaults instantiates a new QbdSalesTaxPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultItemSalesTax

`func (o *QbdSalesTaxPreferences) GetDefaultItemSalesTax() QbdSalesTaxPreferencesDefaultItemSalesTax`

GetDefaultItemSalesTax returns the DefaultItemSalesTax field if non-nil, zero value otherwise.

### GetDefaultItemSalesTaxOk

`func (o *QbdSalesTaxPreferences) GetDefaultItemSalesTaxOk() (*QbdSalesTaxPreferencesDefaultItemSalesTax, bool)`

GetDefaultItemSalesTaxOk returns a tuple with the DefaultItemSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultItemSalesTax

`func (o *QbdSalesTaxPreferences) SetDefaultItemSalesTax(v QbdSalesTaxPreferencesDefaultItemSalesTax)`

SetDefaultItemSalesTax sets DefaultItemSalesTax field to given value.


### GetSalesTaxReportingFrequency

`func (o *QbdSalesTaxPreferences) GetSalesTaxReportingFrequency() string`

GetSalesTaxReportingFrequency returns the SalesTaxReportingFrequency field if non-nil, zero value otherwise.

### GetSalesTaxReportingFrequencyOk

`func (o *QbdSalesTaxPreferences) GetSalesTaxReportingFrequencyOk() (*string, bool)`

GetSalesTaxReportingFrequencyOk returns a tuple with the SalesTaxReportingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReportingFrequency

`func (o *QbdSalesTaxPreferences) SetSalesTaxReportingFrequency(v string)`

SetSalesTaxReportingFrequency sets SalesTaxReportingFrequency field to given value.


### GetDefaultTaxableSalesTaxCode

`func (o *QbdSalesTaxPreferences) GetDefaultTaxableSalesTaxCode() QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode`

GetDefaultTaxableSalesTaxCode returns the DefaultTaxableSalesTaxCode field if non-nil, zero value otherwise.

### GetDefaultTaxableSalesTaxCodeOk

`func (o *QbdSalesTaxPreferences) GetDefaultTaxableSalesTaxCodeOk() (*QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode, bool)`

GetDefaultTaxableSalesTaxCodeOk returns a tuple with the DefaultTaxableSalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxableSalesTaxCode

`func (o *QbdSalesTaxPreferences) SetDefaultTaxableSalesTaxCode(v QbdSalesTaxPreferencesDefaultTaxableSalesTaxCode)`

SetDefaultTaxableSalesTaxCode sets DefaultTaxableSalesTaxCode field to given value.


### GetDefaultNonTaxableSalesTaxCode

`func (o *QbdSalesTaxPreferences) GetDefaultNonTaxableSalesTaxCode() QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode`

GetDefaultNonTaxableSalesTaxCode returns the DefaultNonTaxableSalesTaxCode field if non-nil, zero value otherwise.

### GetDefaultNonTaxableSalesTaxCodeOk

`func (o *QbdSalesTaxPreferences) GetDefaultNonTaxableSalesTaxCodeOk() (*QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode, bool)`

GetDefaultNonTaxableSalesTaxCodeOk returns a tuple with the DefaultNonTaxableSalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNonTaxableSalesTaxCode

`func (o *QbdSalesTaxPreferences) SetDefaultNonTaxableSalesTaxCode(v QbdSalesTaxPreferencesDefaultNonTaxableSalesTaxCode)`

SetDefaultNonTaxableSalesTaxCode sets DefaultNonTaxableSalesTaxCode field to given value.


### GetIsUsingVendorTaxCode

`func (o *QbdSalesTaxPreferences) GetIsUsingVendorTaxCode() bool`

GetIsUsingVendorTaxCode returns the IsUsingVendorTaxCode field if non-nil, zero value otherwise.

### GetIsUsingVendorTaxCodeOk

`func (o *QbdSalesTaxPreferences) GetIsUsingVendorTaxCodeOk() (*bool, bool)`

GetIsUsingVendorTaxCodeOk returns a tuple with the IsUsingVendorTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingVendorTaxCode

`func (o *QbdSalesTaxPreferences) SetIsUsingVendorTaxCode(v bool)`

SetIsUsingVendorTaxCode sets IsUsingVendorTaxCode field to given value.


### GetIsUsingCustomerTaxCode

`func (o *QbdSalesTaxPreferences) GetIsUsingCustomerTaxCode() bool`

GetIsUsingCustomerTaxCode returns the IsUsingCustomerTaxCode field if non-nil, zero value otherwise.

### GetIsUsingCustomerTaxCodeOk

`func (o *QbdSalesTaxPreferences) GetIsUsingCustomerTaxCodeOk() (*bool, bool)`

GetIsUsingCustomerTaxCodeOk returns a tuple with the IsUsingCustomerTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingCustomerTaxCode

`func (o *QbdSalesTaxPreferences) SetIsUsingCustomerTaxCode(v bool)`

SetIsUsingCustomerTaxCode sets IsUsingCustomerTaxCode field to given value.


### GetIsUsingTaxInclusivePrices

`func (o *QbdSalesTaxPreferences) GetIsUsingTaxInclusivePrices() bool`

GetIsUsingTaxInclusivePrices returns the IsUsingTaxInclusivePrices field if non-nil, zero value otherwise.

### GetIsUsingTaxInclusivePricesOk

`func (o *QbdSalesTaxPreferences) GetIsUsingTaxInclusivePricesOk() (*bool, bool)`

GetIsUsingTaxInclusivePricesOk returns a tuple with the IsUsingTaxInclusivePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingTaxInclusivePrices

`func (o *QbdSalesTaxPreferences) SetIsUsingTaxInclusivePrices(v bool)`

SetIsUsingTaxInclusivePrices sets IsUsingTaxInclusivePrices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


