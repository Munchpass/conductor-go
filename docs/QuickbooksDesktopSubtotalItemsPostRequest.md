# QuickbooksDesktopSubtotalItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this subtotal item, unique across all subtotal items.  **NOTE**: Subtotal items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopSubtotalItemsPostRequestBarcode**](QuickbooksDesktopSubtotalItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this subtotal item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**Description** | Pointer to **string** | The subtotal item&#39;s description that will appear on sales forms that include this item. | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopSubtotalItemsPostRequest

`func NewQuickbooksDesktopSubtotalItemsPostRequest(name string, ) *QuickbooksDesktopSubtotalItemsPostRequest`

NewQuickbooksDesktopSubtotalItemsPostRequest instantiates a new QuickbooksDesktopSubtotalItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSubtotalItemsPostRequestWithDefaults

`func NewQuickbooksDesktopSubtotalItemsPostRequestWithDefaults() *QuickbooksDesktopSubtotalItemsPostRequest`

NewQuickbooksDesktopSubtotalItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopSubtotalItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetBarcode() QuickbooksDesktopSubtotalItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopSubtotalItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) SetBarcode(v QuickbooksDesktopSubtotalItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopSubtotalItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


