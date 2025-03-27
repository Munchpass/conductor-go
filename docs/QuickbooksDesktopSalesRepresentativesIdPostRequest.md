# QuickbooksDesktopSalesRepresentativesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the sales representative object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Initial** | Pointer to **string** | The initials of this sales representative&#39;s name. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this sales representative is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**EntityId** | Pointer to **string** | The sales representative&#39;s corresponding person entity in QuickBooks, stored as either an employee, vendor, or other-name entry. | [optional] 

## Methods

### NewQuickbooksDesktopSalesRepresentativesIdPostRequest

`func NewQuickbooksDesktopSalesRepresentativesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopSalesRepresentativesIdPostRequest`

NewQuickbooksDesktopSalesRepresentativesIdPostRequest instantiates a new QuickbooksDesktopSalesRepresentativesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesRepresentativesIdPostRequestWithDefaults

`func NewQuickbooksDesktopSalesRepresentativesIdPostRequestWithDefaults() *QuickbooksDesktopSalesRepresentativesIdPostRequest`

NewQuickbooksDesktopSalesRepresentativesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesRepresentativesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetInitial

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) SetInitial(v string)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEntityId

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QuickbooksDesktopSalesRepresentativesIdPostRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


