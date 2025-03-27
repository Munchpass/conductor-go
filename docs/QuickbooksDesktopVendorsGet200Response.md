# QuickbooksDesktopVendorsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;list\&quot;&#x60;. | 
**Url** | **string** | The endpoint URL where this list can be accessed. | 
**Data** | [**[]QbdVendor**](QbdVendor.md) | The array of vendors. | 
**NextCursor** | **string** | The &#x60;nextCursor&#x60; is a pagination token returned in the response when you use the &#x60;limit&#x60; parameter in your request. To retrieve subsequent pages of results, include this token as the value of the &#x60;cursor&#x60; request parameter in your following API calls.  **NOTE**: The &#x60;nextCursor&#x60; value remains constant throughout the pagination process for a specific list instance; continue to use the same &#x60;nextCursor&#x60; token in each request to fetch additional pages. | 
**RemainingCount** | **float32** | The number of objects remaining to be fetched. | 
**HasMore** | **bool** | Indicates whether there are more objects to be fetched. | 

## Methods

### NewQuickbooksDesktopVendorsGet200Response

`func NewQuickbooksDesktopVendorsGet200Response(objectType string, url string, data []QbdVendor, nextCursor string, remainingCount float32, hasMore bool, ) *QuickbooksDesktopVendorsGet200Response`

NewQuickbooksDesktopVendorsGet200Response instantiates a new QuickbooksDesktopVendorsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopVendorsGet200ResponseWithDefaults

`func NewQuickbooksDesktopVendorsGet200ResponseWithDefaults() *QuickbooksDesktopVendorsGet200Response`

NewQuickbooksDesktopVendorsGet200ResponseWithDefaults instantiates a new QuickbooksDesktopVendorsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *QuickbooksDesktopVendorsGet200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopVendorsGet200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUrl

`func (o *QuickbooksDesktopVendorsGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *QuickbooksDesktopVendorsGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetData

`func (o *QuickbooksDesktopVendorsGet200Response) GetData() []QbdVendor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetDataOk() (*[]QbdVendor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuickbooksDesktopVendorsGet200Response) SetData(v []QbdVendor)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *QuickbooksDesktopVendorsGet200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *QuickbooksDesktopVendorsGet200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetRemainingCount

`func (o *QuickbooksDesktopVendorsGet200Response) GetRemainingCount() float32`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetRemainingCountOk() (*float32, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *QuickbooksDesktopVendorsGet200Response) SetRemainingCount(v float32)`

SetRemainingCount sets RemainingCount field to given value.


### GetHasMore

`func (o *QuickbooksDesktopVendorsGet200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QuickbooksDesktopVendorsGet200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QuickbooksDesktopVendorsGet200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


