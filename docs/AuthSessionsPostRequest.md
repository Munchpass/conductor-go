# AuthSessionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublishableKey** | **string** | Your Conductor publishable key, which we use to create the auth session&#39;s &#x60;authFlowUrl&#x60;. | 
**EndUserId** | **string** | The ID of the end-user for whom to create the integration connection. | 
**LinkExpiryMins** | Pointer to **float32** | The number of minutes after which the auth session will expire. Must be at least 15 minutes and no more than 7 days. If not provided, defaults to 30 minutes. | [optional] [default to 30]
**RedirectUrl** | Pointer to **string** | The URL to which Conductor will redirect the end-user to return to your app after they complete the authentication flow. If not provided, their browser tab will close instead. | [optional] 

## Methods

### NewAuthSessionsPostRequest

`func NewAuthSessionsPostRequest(publishableKey string, endUserId string, ) *AuthSessionsPostRequest`

NewAuthSessionsPostRequest instantiates a new AuthSessionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSessionsPostRequestWithDefaults

`func NewAuthSessionsPostRequestWithDefaults() *AuthSessionsPostRequest`

NewAuthSessionsPostRequestWithDefaults instantiates a new AuthSessionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishableKey

`func (o *AuthSessionsPostRequest) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *AuthSessionsPostRequest) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *AuthSessionsPostRequest) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.


### GetEndUserId

`func (o *AuthSessionsPostRequest) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *AuthSessionsPostRequest) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *AuthSessionsPostRequest) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.


### GetLinkExpiryMins

`func (o *AuthSessionsPostRequest) GetLinkExpiryMins() float32`

GetLinkExpiryMins returns the LinkExpiryMins field if non-nil, zero value otherwise.

### GetLinkExpiryMinsOk

`func (o *AuthSessionsPostRequest) GetLinkExpiryMinsOk() (*float32, bool)`

GetLinkExpiryMinsOk returns a tuple with the LinkExpiryMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryMins

`func (o *AuthSessionsPostRequest) SetLinkExpiryMins(v float32)`

SetLinkExpiryMins sets LinkExpiryMins field to given value.

### HasLinkExpiryMins

`func (o *AuthSessionsPostRequest) HasLinkExpiryMins() bool`

HasLinkExpiryMins returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *AuthSessionsPostRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *AuthSessionsPostRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *AuthSessionsPostRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *AuthSessionsPostRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


