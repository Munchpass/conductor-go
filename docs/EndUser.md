# EndUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this end-user. You must save this value to your database because it is how you identify which of your users to receive your API requests. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;end_user\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this end-user record was created. | 
**CompanyName** | **string** | The end-user&#39;s company name that will be shown elsewhere in Conductor. | 
**SourceId** | **string** | The end-user&#39;s unique identifier from your system. Maps users between your database and Conductor. | 
**Email** | **string** | The end-user&#39;s email address for identification purposes. | 
**IntegrationConnections** | [**[]IntegrationConnection**](IntegrationConnection.md) | The end-user&#39;s integration connections. | 

## Methods

### NewEndUser

`func NewEndUser(id string, objectType string, createdAt string, companyName string, sourceId string, email string, integrationConnections []IntegrationConnection, ) *EndUser`

NewEndUser instantiates a new EndUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserWithDefaults

`func NewEndUserWithDefaults() *EndUser`

NewEndUserWithDefaults instantiates a new EndUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndUser) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *EndUser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EndUser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EndUser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *EndUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EndUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EndUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompanyName

`func (o *EndUser) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EndUser) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EndUser) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetSourceId

`func (o *EndUser) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EndUser) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EndUser) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetEmail

`func (o *EndUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EndUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EndUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIntegrationConnections

`func (o *EndUser) GetIntegrationConnections() []IntegrationConnection`

GetIntegrationConnections returns the IntegrationConnections field if non-nil, zero value otherwise.

### GetIntegrationConnectionsOk

`func (o *EndUser) GetIntegrationConnectionsOk() (*[]IntegrationConnection, bool)`

GetIntegrationConnectionsOk returns a tuple with the IntegrationConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConnections

`func (o *EndUser) SetIntegrationConnections(v []IntegrationConnection)`

SetIntegrationConnections sets IntegrationConnections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


