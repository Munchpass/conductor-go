# QuickbooksDesktopInventorySitesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the inventory site object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive unique name of this inventory site, unique across all inventory sites.  **NOTE**: Inventory sites do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this inventory site is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ParentId** | Pointer to **string** | The parent inventory site one level above this one in the hierarchy. | [optional] 
**Description** | Pointer to **string** | A description of this inventory site. | [optional] 
**Contact** | Pointer to **string** | The name of the primary contact person for this inventory site. | [optional] 
**Phone** | Pointer to **string** | The inventory site&#39;s primary telephone number. | [optional] 
**Fax** | Pointer to **string** | The inventory site&#39;s fax number. | [optional] 
**Email** | Pointer to **string** | The inventory site&#39;s email address. | [optional] 
**Address** | Pointer to [**QuickbooksDesktopInventorySitesPostRequestAddress**](QuickbooksDesktopInventorySitesPostRequestAddress.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopInventorySitesIdPostRequest

`func NewQuickbooksDesktopInventorySitesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopInventorySitesIdPostRequest`

NewQuickbooksDesktopInventorySitesIdPostRequest instantiates a new QuickbooksDesktopInventorySitesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventorySitesIdPostRequestWithDefaults

`func NewQuickbooksDesktopInventorySitesIdPostRequestWithDefaults() *QuickbooksDesktopInventorySitesIdPostRequest`

NewQuickbooksDesktopInventorySitesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopInventorySitesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContact

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetAddress() QuickbooksDesktopInventorySitesPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) GetAddressOk() (*QuickbooksDesktopInventorySitesPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) SetAddress(v QuickbooksDesktopInventorySitesPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopInventorySitesIdPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


