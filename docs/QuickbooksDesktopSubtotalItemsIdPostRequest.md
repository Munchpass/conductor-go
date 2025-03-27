# QuickbooksDesktopSubtotalItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the subtotal item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive unique name of this subtotal item, unique across all subtotal items.  **NOTE**: Subtotal items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopSubtotalItemsPostRequestBarcode**](QuickbooksDesktopSubtotalItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this subtotal item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**Description** | Pointer to **string** | The subtotal item&#39;s description that will appear on sales forms that include this item. | [optional] 

## Methods

### NewQuickbooksDesktopSubtotalItemsIdPostRequest

`func NewQuickbooksDesktopSubtotalItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopSubtotalItemsIdPostRequest`

NewQuickbooksDesktopSubtotalItemsIdPostRequest instantiates a new QuickbooksDesktopSubtotalItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSubtotalItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopSubtotalItemsIdPostRequestWithDefaults() *QuickbooksDesktopSubtotalItemsIdPostRequest`

NewQuickbooksDesktopSubtotalItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopSubtotalItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetBarcode() QuickbooksDesktopSubtotalItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopSubtotalItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) SetBarcode(v QuickbooksDesktopSubtotalItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSubtotalItemsIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


