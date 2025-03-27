# QuickbooksDesktopBillCheckPaymentsIdDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of the deleted bill check payment. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_bill_check_payment\&quot;&#x60;. | 
**RefNumber** | **string** | The case-sensitive user-defined reference number of the deleted bill check payment. | 
**Deleted** | **bool** | Indicates whether the bill check payment was deleted. | 

## Methods

### NewQuickbooksDesktopBillCheckPaymentsIdDelete200Response

`func NewQuickbooksDesktopBillCheckPaymentsIdDelete200Response(id string, objectType string, refNumber string, deleted bool, ) *QuickbooksDesktopBillCheckPaymentsIdDelete200Response`

NewQuickbooksDesktopBillCheckPaymentsIdDelete200Response instantiates a new QuickbooksDesktopBillCheckPaymentsIdDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillCheckPaymentsIdDelete200ResponseWithDefaults

`func NewQuickbooksDesktopBillCheckPaymentsIdDelete200ResponseWithDefaults() *QuickbooksDesktopBillCheckPaymentsIdDelete200Response`

NewQuickbooksDesktopBillCheckPaymentsIdDelete200ResponseWithDefaults instantiates a new QuickbooksDesktopBillCheckPaymentsIdDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetDeleted

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *QuickbooksDesktopBillCheckPaymentsIdDelete200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


