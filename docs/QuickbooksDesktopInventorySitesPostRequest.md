# QuickbooksDesktopInventorySitesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this inventory site, unique across all inventory sites.  **NOTE**: Inventory sites do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this inventory site is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ParentId** | Pointer to **string** | The parent inventory site one level above this one in the hierarchy. | [optional] 
**Description** | Pointer to **string** | A description of this inventory site. | [optional] 
**Email** | Pointer to **string** | The inventory site&#39;s email address. | [optional] 
**Address** | Pointer to [**QuickbooksDesktopInventorySitesPostRequestAddress**](QuickbooksDesktopInventorySitesPostRequestAddress.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopInventorySitesPostRequest

`func NewQuickbooksDesktopInventorySitesPostRequest(name string, ) *QuickbooksDesktopInventorySitesPostRequest`

NewQuickbooksDesktopInventorySitesPostRequest instantiates a new QuickbooksDesktopInventorySitesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventorySitesPostRequestWithDefaults

`func NewQuickbooksDesktopInventorySitesPostRequestWithDefaults() *QuickbooksDesktopInventorySitesPostRequest`

NewQuickbooksDesktopInventorySitesPostRequestWithDefaults instantiates a new QuickbooksDesktopInventorySitesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventorySitesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventorySitesPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopInventorySitesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopInventorySitesPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetAddress() QuickbooksDesktopInventorySitesPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopInventorySitesPostRequest) GetAddressOk() (*QuickbooksDesktopInventorySitesPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopInventorySitesPostRequest) SetAddress(v QuickbooksDesktopInventorySitesPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopInventorySitesPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


