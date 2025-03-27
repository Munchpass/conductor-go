# QuickbooksDesktopSalesTaxItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the sales-tax item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive unique name of this sales-tax item, unique across all sales-tax items.  **NOTE**: Sales-tax items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopSalesTaxItemsPostRequestBarcode**](QuickbooksDesktopSalesTaxItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this sales-tax item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The sales-tax item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Description** | Pointer to **string** | The sales-tax item&#39;s description that will appear on sales forms that include this item. | [optional] 
**TaxRate** | Pointer to **string** | The tax rate defined by this sales-tax item, represented as a decimal string. For example, \&quot;7.5\&quot; represents a 7.5% tax rate. This rate determines the amount of sales tax applied when this item is used in transactions. If a non-zero &#x60;taxRate&#x60; is specified, then the &#x60;taxVendor&#x60; field is required. | [optional] 
**TaxVendorId** | Pointer to **string** | The tax agency (vendor) to whom collected sales taxes are owed for this sales-tax item. This field refers to a vendor in QuickBooks that represents the tax authority. If a non-zero &#x60;taxRate&#x60; is specified, then &#x60;taxVendor&#x60; is required. | [optional] 
**SalesTaxReturnLineId** | Pointer to **string** | The specific line on the sales tax return form where the tax collected using this sales-tax item should be reported. | [optional] 

## Methods

### NewQuickbooksDesktopSalesTaxItemsIdPostRequest

`func NewQuickbooksDesktopSalesTaxItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopSalesTaxItemsIdPostRequest`

NewQuickbooksDesktopSalesTaxItemsIdPostRequest instantiates a new QuickbooksDesktopSalesTaxItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesTaxItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopSalesTaxItemsIdPostRequestWithDefaults() *QuickbooksDesktopSalesTaxItemsIdPostRequest`

NewQuickbooksDesktopSalesTaxItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesTaxItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetBarcode() QuickbooksDesktopSalesTaxItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopSalesTaxItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetBarcode(v QuickbooksDesktopSalesTaxItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetTaxVendorId() string`

GetTaxVendorId returns the TaxVendorId field if non-nil, zero value otherwise.

### GetTaxVendorIdOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetTaxVendorIdOk() (*string, bool)`

GetTaxVendorIdOk returns a tuple with the TaxVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetTaxVendorId(v string)`

SetTaxVendorId sets TaxVendorId field to given value.

### HasTaxVendorId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasTaxVendorId() bool`

HasTaxVendorId returns a boolean if a field has been set.

### GetSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetSalesTaxReturnLineId() string`

GetSalesTaxReturnLineId returns the SalesTaxReturnLineId field if non-nil, zero value otherwise.

### GetSalesTaxReturnLineIdOk

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) GetSalesTaxReturnLineIdOk() (*string, bool)`

GetSalesTaxReturnLineIdOk returns a tuple with the SalesTaxReturnLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) SetSalesTaxReturnLineId(v string)`

SetSalesTaxReturnLineId sets SalesTaxReturnLineId field to given value.

### HasSalesTaxReturnLineId

`func (o *QuickbooksDesktopSalesTaxItemsIdPostRequest) HasSalesTaxReturnLineId() bool`

HasSalesTaxReturnLineId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


