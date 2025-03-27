# QuickbooksDesktopInventorySitesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;list\&quot;&#x60;. | 
**Url** | **string** | The endpoint URL where this list can be accessed. | 
**Data** | [**[]QbdInventorySite**](QbdInventorySite.md) | The array of inventory sites. | 

## Methods

### NewQuickbooksDesktopInventorySitesGet200Response

`func NewQuickbooksDesktopInventorySitesGet200Response(objectType string, url string, data []QbdInventorySite, ) *QuickbooksDesktopInventorySitesGet200Response`

NewQuickbooksDesktopInventorySitesGet200Response instantiates a new QuickbooksDesktopInventorySitesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventorySitesGet200ResponseWithDefaults

`func NewQuickbooksDesktopInventorySitesGet200ResponseWithDefaults() *QuickbooksDesktopInventorySitesGet200Response`

NewQuickbooksDesktopInventorySitesGet200ResponseWithDefaults instantiates a new QuickbooksDesktopInventorySitesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopInventorySitesGet200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUrl

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *QuickbooksDesktopInventorySitesGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetData

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetData() []QbdInventorySite`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuickbooksDesktopInventorySitesGet200Response) GetDataOk() (*[]QbdInventorySite, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuickbooksDesktopInventorySitesGet200Response) SetData(v []QbdInventorySite)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


