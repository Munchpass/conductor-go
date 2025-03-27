# QuickbooksDesktopNonInventoryItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this non-inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two non-inventory items could both have the &#x60;name&#x60; \&quot;Printer Ink Cartridge\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Office Supplies:Printer Ink Cartridge\&quot; and \&quot;Miscellaneous:Printer Ink Cartridge\&quot;.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopNonInventoryItemsPostRequestBarcode**](QuickbooksDesktopNonInventoryItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this non-inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ParentId** | Pointer to **string** | The parent non-inventory item one level above this one in the hierarchy. For example, if this non-inventory item has a &#x60;fullName&#x60; of \&quot;Office Supplies:Printer Ink Cartridge\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Office Supplies\&quot;. If this non-inventory item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The non-inventory item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Sku** | Pointer to **string** | The non-inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this non-inventory item, which consists of a base unit and related units. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this non-inventory item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesOrPurchaseDetails** | Pointer to [**QuickbooksDesktopNonInventoryItemsPostRequestSalesOrPurchaseDetails**](QuickbooksDesktopNonInventoryItemsPostRequestSalesOrPurchaseDetails.md) |  | [optional] 
**SalesAndPurchaseDetails** | Pointer to [**QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails**](QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopNonInventoryItemsPostRequest

`func NewQuickbooksDesktopNonInventoryItemsPostRequest(name string, ) *QuickbooksDesktopNonInventoryItemsPostRequest`

NewQuickbooksDesktopNonInventoryItemsPostRequest instantiates a new QuickbooksDesktopNonInventoryItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopNonInventoryItemsPostRequestWithDefaults

`func NewQuickbooksDesktopNonInventoryItemsPostRequestWithDefaults() *QuickbooksDesktopNonInventoryItemsPostRequest`

NewQuickbooksDesktopNonInventoryItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopNonInventoryItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetBarcode() QuickbooksDesktopNonInventoryItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopNonInventoryItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetBarcode(v QuickbooksDesktopNonInventoryItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetSku

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesOrPurchaseDetails() QuickbooksDesktopNonInventoryItemsPostRequestSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesOrPurchaseDetailsOk() (*QuickbooksDesktopNonInventoryItemsPostRequestSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetSalesOrPurchaseDetails(v QuickbooksDesktopNonInventoryItemsPostRequestSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.

### HasSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasSalesOrPurchaseDetails() bool`

HasSalesOrPurchaseDetails returns a boolean if a field has been set.

### GetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesAndPurchaseDetails() QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetSalesAndPurchaseDetailsOk() (*QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetSalesAndPurchaseDetails(v QuickbooksDesktopNonInventoryItemsPostRequestSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.

### HasSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasSalesAndPurchaseDetails() bool`

HasSalesAndPurchaseDetails returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopNonInventoryItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


