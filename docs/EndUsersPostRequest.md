# EndUsersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **string** | The end-user&#39;s company name that will be shown elsewhere in Conductor. | 
**SourceId** | **string** | The end-user&#39;s unique identifier from your system. Maps users between your database and Conductor. Must be unique for each user. If you have only one user, you may use any string value. | 
**Email** | **string** | The end-user&#39;s email address for identification purposes. Setting this field will not cause any emails to be sent. | 

## Methods

### NewEndUsersPostRequest

`func NewEndUsersPostRequest(companyName string, sourceId string, email string, ) *EndUsersPostRequest`

NewEndUsersPostRequest instantiates a new EndUsersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUsersPostRequestWithDefaults

`func NewEndUsersPostRequestWithDefaults() *EndUsersPostRequest`

NewEndUsersPostRequestWithDefaults instantiates a new EndUsersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *EndUsersPostRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EndUsersPostRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EndUsersPostRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetSourceId

`func (o *EndUsersPostRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EndUsersPostRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EndUsersPostRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetEmail

`func (o *EndUsersPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EndUsersPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EndUsersPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


