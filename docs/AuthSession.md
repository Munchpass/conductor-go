# AuthSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this auth session. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;auth_session\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this auth session record was created. | 
**EndUserId** | **string** | The ID of the end-user for whom to create an integration connection. | 
**ClientSecret** | **string** | The secret used in &#x60;authFlowUrl&#x60; to securely access the authentication flow. | 
**AuthFlowUrl** | **string** | The URL of the authentication flow that you will pass to your client for your user to set up their integration connection. | 
**ExpiresAt** | **string** | The date and time when this auth session expires. By default, this value is 30 minutes from creation. You can extend this time by setting &#x60;linkExpiryMins&#x60; when creating the auth session. | 
**RedirectUrl** | **NullableString** | The URL to which Conductor will redirect your user to return to your app after they complete the authentication flow. If &#x60;null&#x60;, their browser tab will close instead. | 

## Methods

### NewAuthSession

`func NewAuthSession(id string, objectType string, createdAt string, endUserId string, clientSecret string, authFlowUrl string, expiresAt string, redirectUrl NullableString, ) *AuthSession`

NewAuthSession instantiates a new AuthSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSessionWithDefaults

`func NewAuthSessionWithDefaults() *AuthSession`

NewAuthSessionWithDefaults instantiates a new AuthSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthSession) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *AuthSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AuthSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AuthSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *AuthSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndUserId

`func (o *AuthSession) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *AuthSession) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *AuthSession) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.


### GetClientSecret

`func (o *AuthSession) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthSession) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthSession) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetAuthFlowUrl

`func (o *AuthSession) GetAuthFlowUrl() string`

GetAuthFlowUrl returns the AuthFlowUrl field if non-nil, zero value otherwise.

### GetAuthFlowUrlOk

`func (o *AuthSession) GetAuthFlowUrlOk() (*string, bool)`

GetAuthFlowUrlOk returns a tuple with the AuthFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlowUrl

`func (o *AuthSession) SetAuthFlowUrl(v string)`

SetAuthFlowUrl sets AuthFlowUrl field to given value.


### GetExpiresAt

`func (o *AuthSession) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthSession) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthSession) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetRedirectUrl

`func (o *AuthSession) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *AuthSession) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *AuthSession) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### SetRedirectUrlNil

`func (o *AuthSession) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *AuthSession) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


