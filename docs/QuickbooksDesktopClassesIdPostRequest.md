# QuickbooksDesktopClassesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the class object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this class. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two classes could both have the &#x60;name&#x60; \&quot;Marketing\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Department:Marketing\&quot; and \&quot;Internal:Marketing\&quot;.  Maximum length: 31 characters. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this class is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ParentId** | Pointer to **string** | The parent class one level above this one in the hierarchy. For example, if this class has a &#x60;fullName&#x60; of \&quot;Department:Marketing\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Department\&quot;. If this class is at the top level, this field will be &#x60;null&#x60;. | [optional] 

## Methods

### NewQuickbooksDesktopClassesIdPostRequest

`func NewQuickbooksDesktopClassesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopClassesIdPostRequest`

NewQuickbooksDesktopClassesIdPostRequest instantiates a new QuickbooksDesktopClassesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopClassesIdPostRequestWithDefaults

`func NewQuickbooksDesktopClassesIdPostRequestWithDefaults() *QuickbooksDesktopClassesIdPostRequest`

NewQuickbooksDesktopClassesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopClassesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopClassesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopClassesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopClassesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopClassesIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopClassesIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopClassesIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopClassesIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopClassesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopClassesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopClassesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopClassesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopClassesIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopClassesIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopClassesIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopClassesIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


