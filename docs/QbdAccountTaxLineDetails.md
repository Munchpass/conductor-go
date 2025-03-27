# QbdAccountTaxLineDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxLineId** | **float32** | The identifier of the tax line associated with this account. | 
**TaxLineName** | **string** | The name of the tax line associated with this account, as it appears on the tax form. | 

## Methods

### NewQbdAccountTaxLineDetails

`func NewQbdAccountTaxLineDetails(taxLineId float32, taxLineName string, ) *QbdAccountTaxLineDetails`

NewQbdAccountTaxLineDetails instantiates a new QbdAccountTaxLineDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdAccountTaxLineDetailsWithDefaults

`func NewQbdAccountTaxLineDetailsWithDefaults() *QbdAccountTaxLineDetails`

NewQbdAccountTaxLineDetailsWithDefaults instantiates a new QbdAccountTaxLineDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxLineId

`func (o *QbdAccountTaxLineDetails) GetTaxLineId() float32`

GetTaxLineId returns the TaxLineId field if non-nil, zero value otherwise.

### GetTaxLineIdOk

`func (o *QbdAccountTaxLineDetails) GetTaxLineIdOk() (*float32, bool)`

GetTaxLineIdOk returns a tuple with the TaxLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLineId

`func (o *QbdAccountTaxLineDetails) SetTaxLineId(v float32)`

SetTaxLineId sets TaxLineId field to given value.


### GetTaxLineName

`func (o *QbdAccountTaxLineDetails) GetTaxLineName() string`

GetTaxLineName returns the TaxLineName field if non-nil, zero value otherwise.

### GetTaxLineNameOk

`func (o *QbdAccountTaxLineDetails) GetTaxLineNameOk() (*string, bool)`

GetTaxLineNameOk returns a tuple with the TaxLineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLineName

`func (o *QbdAccountTaxLineDetails) SetTaxLineName(v string)`

SetTaxLineName sets TaxLineName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


