# QuickbooksDesktopInventoryAdjustmentsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;list\&quot;&#x60;. | 
**Url** | **string** | The endpoint URL where this list can be accessed. | 
**Data** | [**[]QbdInventoryAdjustment**](QbdInventoryAdjustment.md) | The array of inventory adjustments. | 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsGet200Response

`func NewQuickbooksDesktopInventoryAdjustmentsGet200Response(objectType string, url string, data []QbdInventoryAdjustment, ) *QuickbooksDesktopInventoryAdjustmentsGet200Response`

NewQuickbooksDesktopInventoryAdjustmentsGet200Response instantiates a new QuickbooksDesktopInventoryAdjustmentsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsGet200ResponseWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsGet200ResponseWithDefaults() *QuickbooksDesktopInventoryAdjustmentsGet200Response`

NewQuickbooksDesktopInventoryAdjustmentsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUrl

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetData

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetData() []QbdInventoryAdjustment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) GetDataOk() (*[]QbdInventoryAdjustment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuickbooksDesktopInventoryAdjustmentsGet200Response) SetData(v []QbdInventoryAdjustment)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


