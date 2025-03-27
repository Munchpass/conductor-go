# QuickbooksDesktopSalesTaxCodesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this sales-tax code, unique across all sales-tax codes. This short name will appear on sales forms to identify the tax status of an item.  **NOTE**: Sales-tax codes do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 3 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this sales-tax code is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**IsTaxable** | **bool** | Indicates whether this sales-tax code is tracking taxable sales. This field cannot be modified once the sales-tax code has been used in a transaction. | 
**Description** | Pointer to **string** | A description of this sales-tax code. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this sales-tax code&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 

## Methods

### NewQuickbooksDesktopSalesTaxCodesPostRequest

`func NewQuickbooksDesktopSalesTaxCodesPostRequest(name string, isTaxable bool, ) *QuickbooksDesktopSalesTaxCodesPostRequest`

NewQuickbooksDesktopSalesTaxCodesPostRequest instantiates a new QuickbooksDesktopSalesTaxCodesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesTaxCodesPostRequestWithDefaults

`func NewQuickbooksDesktopSalesTaxCodesPostRequestWithDefaults() *QuickbooksDesktopSalesTaxCodesPostRequest`

NewQuickbooksDesktopSalesTaxCodesPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxCodesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsTaxable

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.


### GetDescription

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopSalesTaxCodesPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


