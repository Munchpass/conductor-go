# QuickbooksDesktopSalesRepresentativesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | **string** | The initials of this sales representative&#39;s name. | 
**IsActive** | Pointer to **bool** | Indicates whether this sales representative is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**EntityId** | **string** | The sales representative&#39;s corresponding person entity in QuickBooks, stored as either an employee, vendor, or other-name entry. | 

## Methods

### NewQuickbooksDesktopSalesRepresentativesPostRequest

`func NewQuickbooksDesktopSalesRepresentativesPostRequest(initial string, entityId string, ) *QuickbooksDesktopSalesRepresentativesPostRequest`

NewQuickbooksDesktopSalesRepresentativesPostRequest instantiates a new QuickbooksDesktopSalesRepresentativesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesRepresentativesPostRequestWithDefaults

`func NewQuickbooksDesktopSalesRepresentativesPostRequestWithDefaults() *QuickbooksDesktopSalesRepresentativesPostRequest`

NewQuickbooksDesktopSalesRepresentativesPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesRepresentativesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitial

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetInitial() string`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetInitialOk() (*string, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetInitial(v string)`

SetInitial sets Initial field to given value.


### GetIsActive

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEntityId

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopSalesRepresentativesPostRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


