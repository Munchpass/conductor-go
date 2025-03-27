# QbdSubscribedService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | [**[]QbdService**](QbdService.md) | The list of Intuit services that this company is or has been subscribed to, for example, Intuit Payroll, QBMS. | 

## Methods

### NewQbdSubscribedService

`func NewQbdSubscribedService(services []QbdService, ) *QbdSubscribedService`

NewQbdSubscribedService instantiates a new QbdSubscribedService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSubscribedServiceWithDefaults

`func NewQbdSubscribedServiceWithDefaults() *QbdSubscribedService`

NewQbdSubscribedServiceWithDefaults instantiates a new QbdSubscribedService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *QbdSubscribedService) GetServices() []QbdService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *QbdSubscribedService) GetServicesOk() (*[]QbdService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *QbdSubscribedService) SetServices(v []QbdService)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


