# IntegrationConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this integration connection. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;integration_connection\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this integration connection record was created. | 
**IntegrationSlug** | **string** | The identifier of the third-party platform to integrate. | 
**LastRequestAt** | **NullableString** | The date and time of your last API request to this integration connection. | 
**LastSuccessfulRequestAt** | **NullableString** | The date and time of your last *successful* API request to this integration connection. A successful request means the integration fully processed and returned a response without any errors end-to-end. | 

## Methods

### NewIntegrationConnection

`func NewIntegrationConnection(id string, objectType string, createdAt string, integrationSlug string, lastRequestAt NullableString, lastSuccessfulRequestAt NullableString, ) *IntegrationConnection`

NewIntegrationConnection instantiates a new IntegrationConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConnectionWithDefaults

`func NewIntegrationConnectionWithDefaults() *IntegrationConnection`

NewIntegrationConnectionWithDefaults instantiates a new IntegrationConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationConnection) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *IntegrationConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IntegrationConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IntegrationConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *IntegrationConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetIntegrationSlug

`func (o *IntegrationConnection) GetIntegrationSlug() string`

GetIntegrationSlug returns the IntegrationSlug field if non-nil, zero value otherwise.

### GetIntegrationSlugOk

`func (o *IntegrationConnection) GetIntegrationSlugOk() (*string, bool)`

GetIntegrationSlugOk returns a tuple with the IntegrationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationSlug

`func (o *IntegrationConnection) SetIntegrationSlug(v string)`

SetIntegrationSlug sets IntegrationSlug field to given value.


### GetLastRequestAt

`func (o *IntegrationConnection) GetLastRequestAt() string`

GetLastRequestAt returns the LastRequestAt field if non-nil, zero value otherwise.

### GetLastRequestAtOk

`func (o *IntegrationConnection) GetLastRequestAtOk() (*string, bool)`

GetLastRequestAtOk returns a tuple with the LastRequestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequestAt

`func (o *IntegrationConnection) SetLastRequestAt(v string)`

SetLastRequestAt sets LastRequestAt field to given value.


### SetLastRequestAtNil

`func (o *IntegrationConnection) SetLastRequestAtNil(b bool)`

 SetLastRequestAtNil sets the value for LastRequestAt to be an explicit nil

### UnsetLastRequestAt
`func (o *IntegrationConnection) UnsetLastRequestAt()`

UnsetLastRequestAt ensures that no value is present for LastRequestAt, not even an explicit nil
### GetLastSuccessfulRequestAt

`func (o *IntegrationConnection) GetLastSuccessfulRequestAt() string`

GetLastSuccessfulRequestAt returns the LastSuccessfulRequestAt field if non-nil, zero value otherwise.

### GetLastSuccessfulRequestAtOk

`func (o *IntegrationConnection) GetLastSuccessfulRequestAtOk() (*string, bool)`

GetLastSuccessfulRequestAtOk returns a tuple with the LastSuccessfulRequestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRequestAt

`func (o *IntegrationConnection) SetLastSuccessfulRequestAt(v string)`

SetLastSuccessfulRequestAt sets LastSuccessfulRequestAt field to given value.


### SetLastSuccessfulRequestAtNil

`func (o *IntegrationConnection) SetLastSuccessfulRequestAtNil(b bool)`

 SetLastSuccessfulRequestAtNil sets the value for LastSuccessfulRequestAt to be an explicit nil

### UnsetLastSuccessfulRequestAt
`func (o *IntegrationConnection) UnsetLastSuccessfulRequestAt()`

UnsetLastSuccessfulRequestAt ensures that no value is present for LastSuccessfulRequestAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


