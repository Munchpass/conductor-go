# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The developer error message for your logs. | 
**UserFacingMessage** | **string** | The user-friendly error message, written specifically for displaying to your end-users in your app&#39;s UI.  This value exists for _every_ error. E.g., a QBD connection error might recommend the end-user to confirm their QuickBooks Desktop is open and that they&#39;re logged in. But if a Conductor API key is expired, e.g., this message will be masked and just say _\&quot;An internal server error occurred. Please try again later.\&quot;_ | 
**Type** | **string** | The type of error that occurred. | 
**Code** | **string** | The unique error code from Conductor, which is useful for adding special handling for specific errors. E.g., &#x60;&#39;RESOURCE_MISSING&#39;&#x60;, &#x60;&#39;API_KEY_INVALID&#39;&#x60;, or &#x60;&#39;QBD_REQUEST_ERROR&#39;&#x60;. In contrast, the error field &#x60;type&#x60; is more general and categorizes the error. | 
**IntegrationCode** | Pointer to **string** | The unique error code supplied by the third-party integration for errors returned by the integration (e.g., QuickBooks Desktop) or integration connector (e.g., Web Connector). This is useful for adding special handling for specific errors from the third-party integration or connector.  The integration&#39;s corresponding error message for this code is in &#x60;error.message&#x60;. | [optional] 
**HttpStatusCode** | **float32** | The HTTP status code of the response that returned this error. | 
**RequestId** | **string** | The unique identifier for the request that returned this error.  If you need to contact us about a specific request, providing the request identifier will ensure the fastest possible resolution. | 

## Methods

### NewError

`func NewError(message string, userFacingMessage string, type_ string, code string, httpStatusCode float32, requestId string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserFacingMessage

`func (o *Error) GetUserFacingMessage() string`

GetUserFacingMessage returns the UserFacingMessage field if non-nil, zero value otherwise.

### GetUserFacingMessageOk

`func (o *Error) GetUserFacingMessageOk() (*string, bool)`

GetUserFacingMessageOk returns a tuple with the UserFacingMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFacingMessage

`func (o *Error) SetUserFacingMessage(v string)`

SetUserFacingMessage sets UserFacingMessage field to given value.


### GetType

`func (o *Error) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.


### GetIntegrationCode

`func (o *Error) GetIntegrationCode() string`

GetIntegrationCode returns the IntegrationCode field if non-nil, zero value otherwise.

### GetIntegrationCodeOk

`func (o *Error) GetIntegrationCodeOk() (*string, bool)`

GetIntegrationCodeOk returns a tuple with the IntegrationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationCode

`func (o *Error) SetIntegrationCode(v string)`

SetIntegrationCode sets IntegrationCode field to given value.

### HasIntegrationCode

`func (o *Error) HasIntegrationCode() bool`

HasIntegrationCode returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *Error) GetHttpStatusCode() float32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *Error) GetHttpStatusCodeOk() (*float32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *Error) SetHttpStatusCode(v float32)`

SetHttpStatusCode sets HttpStatusCode field to given value.


### GetRequestId

`func (o *Error) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


