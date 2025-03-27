# QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The identifier of the owner of the custom field, which QuickBooks internally calls a \&quot;data extension\&quot;. For public custom fields visible in the UI, such as those added by the QuickBooks user, this is always \&quot;0\&quot;. For private custom fields that are only visible to the application that created them, this is a valid GUID identifying the owning application. Internally, Conductor always fetches all public custom fields (those with an &#x60;ownerId&#x60; of \&quot;0\&quot;) for all objects. | 
**Name** | **string** | The name of the custom field, unique for the specified &#x60;ownerId&#x60;. For public custom fields, this name is visible as a label in the QuickBooks UI. | 
**Value** | **string** | The value of this custom field. The maximum length depends on the field&#39;s data type. | 

## Methods

### NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner

`func NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner(ownerId string, name string, value string, ) *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner instantiates a new QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInnerWithDefaults

`func NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInnerWithDefaults() *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

NewQuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInnerWithDefaults instantiates a new QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


