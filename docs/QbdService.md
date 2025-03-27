# QbdService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this service, unique across all services.  **NOTE**: Services do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**Domain** | **NullableString** | The domain of this subscribed service | 
**ServiceStatus** | **NullableString** | The status of this service&#39;s subscription. | 

## Methods

### NewQbdService

`func NewQbdService(name string, domain NullableString, serviceStatus NullableString, ) *QbdService`

NewQbdService instantiates a new QbdService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdServiceWithDefaults

`func NewQbdServiceWithDefaults() *QbdService`

NewQbdServiceWithDefaults instantiates a new QbdService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QbdService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdService) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *QbdService) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *QbdService) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *QbdService) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *QbdService) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *QbdService) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetServiceStatus

`func (o *QbdService) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *QbdService) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *QbdService) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### SetServiceStatusNil

`func (o *QbdService) SetServiceStatusNil(b bool)`

 SetServiceStatusNil sets the value for ServiceStatus to be an explicit nil

### UnsetServiceStatus
`func (o *QbdService) UnsetServiceStatus()`

UnsetServiceStatus ensures that no value is present for ServiceStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


