# EndUsersIdDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the deleted end-user. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;end_user\&quot;&#x60;. | 
**Deleted** | **bool** | Indicates whether the end-user was deleted. | 

## Methods

### NewEndUsersIdDelete200Response

`func NewEndUsersIdDelete200Response(id string, objectType string, deleted bool, ) *EndUsersIdDelete200Response`

NewEndUsersIdDelete200Response instantiates a new EndUsersIdDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUsersIdDelete200ResponseWithDefaults

`func NewEndUsersIdDelete200ResponseWithDefaults() *EndUsersIdDelete200Response`

NewEndUsersIdDelete200ResponseWithDefaults instantiates a new EndUsersIdDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndUsersIdDelete200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndUsersIdDelete200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndUsersIdDelete200Response) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *EndUsersIdDelete200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EndUsersIdDelete200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EndUsersIdDelete200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeleted

`func (o *EndUsersIdDelete200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *EndUsersIdDelete200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *EndUsersIdDelete200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


