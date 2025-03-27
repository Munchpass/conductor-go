# QuickbooksDesktopSalesOrdersIdDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of the deleted sales order. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_order\&quot;&#x60;. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number of the deleted sales order. | 
**Deleted** | **bool** | Indicates whether the sales order was deleted. | 

## Methods

### NewQuickbooksDesktopSalesOrdersIdDelete200Response

`func NewQuickbooksDesktopSalesOrdersIdDelete200Response(id string, objectType string, refNumber string, deleted bool, ) *QuickbooksDesktopSalesOrdersIdDelete200Response`

NewQuickbooksDesktopSalesOrdersIdDelete200Response instantiates a new QuickbooksDesktopSalesOrdersIdDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesOrdersIdDelete200ResponseWithDefaults

`func NewQuickbooksDesktopSalesOrdersIdDelete200ResponseWithDefaults() *QuickbooksDesktopSalesOrdersIdDelete200Response`

NewQuickbooksDesktopSalesOrdersIdDelete200ResponseWithDefaults instantiates a new QuickbooksDesktopSalesOrdersIdDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetDeleted

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *QuickbooksDesktopSalesOrdersIdDelete200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


