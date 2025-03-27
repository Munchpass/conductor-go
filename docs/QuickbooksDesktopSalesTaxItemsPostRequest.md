# QuickbooksDesktopSalesTaxItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this sales-tax item, unique across all sales-tax items.  **NOTE**: Sales-tax items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopSalesTaxItemsPostRequestBarcode**](QuickbooksDesktopSalesTaxItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this sales-tax item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The sales-tax item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Description** | Pointer to **string** | The sales-tax item&#39;s description that will appear on sales forms that include this item. | [optional] 
**TaxRate** | Pointer to **string** | The tax rate defined by this sales-tax item, represented as a decimal string. For example, \&quot;7.5\&quot; represents a 7.5% tax rate. This rate determines the amount of sales tax applied when this item is used in transactions. If a non-zero &#x60;taxRate&#x60; is specified, then the &#x60;taxVendor&#x60; field is required. | [optional] 
**TaxVendorId** | Pointer to **string** | The tax agency (vendor) to whom collected sales taxes are owed for this sales-tax item. This field refers to a vendor in QuickBooks that represents the tax authority. If a non-zero &#x60;taxRate&#x60; is specified, then &#x60;taxVendor&#x60; is required. | [optional] 
**SalesTaxReturnLineId** | Pointer to **string** | The specific line on the sales tax return form where the tax collected using this sales-tax item should be reported. | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopSalesTaxItemsPostRequest

`func NewQuickbooksDesktopSalesTaxItemsPostRequest(name string, ) *QuickbooksDesktopSalesTaxItemsPostRequest`

NewQuickbooksDesktopSalesTaxItemsPostRequest instantiates a new QuickbooksDesktopSalesTaxItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesTaxItemsPostRequestWithDefaults

`func NewQuickbooksDesktopSalesTaxItemsPostRequestWithDefaults() *QuickbooksDesktopSalesTaxItemsPostRequest`

NewQuickbooksDesktopSalesTaxItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetBarcode() QuickbooksDesktopSalesTaxItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopSalesTaxItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetBarcode(v QuickbooksDesktopSalesTaxItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxVendorId() string`

GetTaxVendorId returns the TaxVendorId field if non-nil, zero value otherwise.

### GetTaxVendorIdOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetTaxVendorIdOk() (*string, bool)`

GetTaxVendorIdOk returns a tuple with the TaxVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetTaxVendorId(v string)`

SetTaxVendorId sets TaxVendorId field to given value.

### HasTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasTaxVendorId() bool`

HasTaxVendorId returns a boolean if a field has been set.

### GetSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetSalesTaxReturnLineId() string`

GetSalesTaxReturnLineId returns the SalesTaxReturnLineId field if non-nil, zero value otherwise.

### GetSalesTaxReturnLineIdOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetSalesTaxReturnLineIdOk() (*string, bool)`

GetSalesTaxReturnLineIdOk returns a tuple with the SalesTaxReturnLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetSalesTaxReturnLineId(v string)`

SetSalesTaxReturnLineId sets SalesTaxReturnLineId field to given value.

### HasSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasSalesTaxReturnLineId() bool`

HasSalesTaxReturnLineId returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopSalesTaxItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


