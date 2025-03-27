# QuickbooksDesktopPriceLevelsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;list\&quot;&#x60;. | 
**Url** | **string** | The endpoint URL where this list can be accessed. | 
**Data** | [**[]QbdPriceLevel**](QbdPriceLevel.md) | The array of price levels. | 

## Methods

### NewQuickbooksDesktopPriceLevelsGet200Response

`func NewQuickbooksDesktopPriceLevelsGet200Response(objectType string, url string, data []QbdPriceLevel, ) *QuickbooksDesktopPriceLevelsGet200Response`

NewQuickbooksDesktopPriceLevelsGet200Response instantiates a new QuickbooksDesktopPriceLevelsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPriceLevelsGet200ResponseWithDefaults

`func NewQuickbooksDesktopPriceLevelsGet200ResponseWithDefaults() *QuickbooksDesktopPriceLevelsGet200Response`

NewQuickbooksDesktopPriceLevelsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopPriceLevelsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopPriceLevelsGet200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUrl

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *QuickbooksDesktopPriceLevelsGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetData

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetData() []QbdPriceLevel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuickbooksDesktopPriceLevelsGet200Response) GetDataOk() (*[]QbdPriceLevel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuickbooksDesktopPriceLevelsGet200Response) SetData(v []QbdPriceLevel)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


