# QuickbooksDesktopClassesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this class. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two classes could both have the &#x60;name&#x60; \&quot;Marketing\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Department:Marketing\&quot; and \&quot;Internal:Marketing\&quot;.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this class is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ParentId** | Pointer to **string** | The parent class one level above this one in the hierarchy. For example, if this class has a &#x60;fullName&#x60; of \&quot;Department:Marketing\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Department\&quot;. If this class is at the top level, this field will be &#x60;null&#x60;. | [optional] 

## Methods

### NewQuickbooksDesktopClassesPostRequest

`func NewQuickbooksDesktopClassesPostRequest(name string, ) *QuickbooksDesktopClassesPostRequest`

NewQuickbooksDesktopClassesPostRequest instantiates a new QuickbooksDesktopClassesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopClassesPostRequestWithDefaults

`func NewQuickbooksDesktopClassesPostRequestWithDefaults() *QuickbooksDesktopClassesPostRequest`

NewQuickbooksDesktopClassesPostRequestWithDefaults instantiates a new QuickbooksDesktopClassesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopClassesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopClassesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopClassesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopClassesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopClassesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopClassesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopClassesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopClassesPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopClassesPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopClassesPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopClassesPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


