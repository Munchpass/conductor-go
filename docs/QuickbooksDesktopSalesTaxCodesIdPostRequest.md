# QuickbooksDesktopSalesTaxCodesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the sales-tax code object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive unique name of this sales-tax code, unique across all sales-tax codes. This short name will appear on sales forms to identify the tax status of an item.  **NOTE**: Sales-tax codes do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 3 characters. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this sales-tax code is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**IsTaxable** | Pointer to **bool** | Indicates whether this sales-tax code is tracking taxable sales. This field cannot be modified once the sales-tax code has been used in a transaction. | [optional] 
**Description** | Pointer to **string** | A description of this sales-tax code. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this sales-tax code&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 

## Methods

### NewQuickbooksDesktopSalesTaxCodesIdPostRequest

`func NewQuickbooksDesktopSalesTaxCodesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopSalesTaxCodesIdPostRequest`

NewQuickbooksDesktopSalesTaxCodesIdPostRequest instantiates a new QuickbooksDesktopSalesTaxCodesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesTaxCodesIdPostRequestWithDefaults

`func NewQuickbooksDesktopSalesTaxCodesIdPostRequestWithDefaults() *QuickbooksDesktopSalesTaxCodesIdPostRequest`

NewQuickbooksDesktopSalesTaxCodesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxCodesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsTaxable

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.

### HasIsTaxable

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) HasIsTaxable() bool`

HasIsTaxable returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesIdPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


