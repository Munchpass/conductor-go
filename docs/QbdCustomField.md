# QbdCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The identifier of the owner of the custom field, which QuickBooks internally calls a \&quot;data extension\&quot;. For public custom fields visible in the UI, such as those added by the QuickBooks user, this is always \&quot;0\&quot;. For private custom fields that are only visible to the application that created them, this is a valid GUID identifying the owning application. Internally, Conductor always fetches all public custom fields (those with an &#x60;ownerId&#x60; of \&quot;0\&quot;) for all objects. | 
**Name** | **string** | The name of the custom field, unique for the specified &#x60;ownerId&#x60;. For public custom fields, this name is visible as a label in the QuickBooks UI. | 
**Type** | **string** | The data type of this custom field. | 
**Value** | **string** | The value of this custom field. The maximum length depends on the field&#39;s data type. | 

## Methods

### NewQbdCustomField

`func NewQbdCustomField(ownerId string, name string, type_ string, value string, ) *QbdCustomField`

NewQbdCustomField instantiates a new QbdCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCustomFieldWithDefaults

`func NewQbdCustomFieldWithDefaults() *QbdCustomField`

NewQbdCustomFieldWithDefaults instantiates a new QbdCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *QbdCustomField) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *QbdCustomField) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *QbdCustomField) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *QbdCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdCustomField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *QbdCustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QbdCustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QbdCustomField) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *QbdCustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QbdCustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QbdCustomField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


