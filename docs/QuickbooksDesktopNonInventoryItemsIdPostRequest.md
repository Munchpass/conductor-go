# QuickbooksDesktopNonInventoryItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the non-inventory item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this non-inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two non-inventory items could both have the &#x60;name&#x60; \&quot;Printer Ink Cartridge\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Office Supplies:Printer Ink Cartridge\&quot; and \&quot;Miscellaneous:Printer Ink Cartridge\&quot;.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopNonInventoryItemsPostRequestBarcode**](QuickbooksDesktopNonInventoryItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this non-inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The non-inventory item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent non-inventory item one level above this one in the hierarchy. For example, if this non-inventory item has a &#x60;fullName&#x60; of \&quot;Office Supplies:Printer Ink Cartridge\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Office Supplies\&quot;. If this non-inventory item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**Sku** | Pointer to **string** | The non-inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this non-inventory item, which consists of a base unit and related units. | [optional] 
**ForceUnitOfMeasureChange** | Pointer to **bool** | Indicates whether to allow changing the non-inventory item&#39;s unit-of-measure set (using the &#x60;unitOfMeasureSetId&#x60; field) when the base unit of the new unit-of-measure set does not match that of the currently assigned set. Without setting this field to &#x60;true&#x60; in this scenario, the request will fail with an error; hence, this field is equivalent to accepting the warning prompt in the QuickBooks UI.  NOTE: Changing the base unit requires you to update the item&#39;s quantities-on-hand and cost to reflect the new unit; otherwise, these values will be inaccurate. Alternatively, consider creating a new item with the desired unit-of-measure set and deactivating the old item. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this non-inventory item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesOrPurchaseDetails** | Pointer to [**QuickbooksDesktopNonInventoryItemsIdPostRequestSalesOrPurchaseDetails**](QuickbooksDesktopNonInventoryItemsIdPostRequestSalesOrPurchaseDetails.md) |  | [optional] 
**SalesAndPurchaseDetails** | Pointer to [**QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails**](QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopNonInventoryItemsIdPostRequest

`func NewQuickbooksDesktopNonInventoryItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopNonInventoryItemsIdPostRequest`

NewQuickbooksDesktopNonInventoryItemsIdPostRequest instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopNonInventoryItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopNonInventoryItemsIdPostRequestWithDefaults() *QuickbooksDesktopNonInventoryItemsIdPostRequest`

NewQuickbooksDesktopNonInventoryItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopNonInventoryItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetBarcode() QuickbooksDesktopNonInventoryItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopNonInventoryItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetBarcode(v QuickbooksDesktopNonInventoryItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSku

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetForceUnitOfMeasureChange() bool`

GetForceUnitOfMeasureChange returns the ForceUnitOfMeasureChange field if non-nil, zero value otherwise.

### GetForceUnitOfMeasureChangeOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetForceUnitOfMeasureChangeOk() (*bool, bool)`

GetForceUnitOfMeasureChangeOk returns a tuple with the ForceUnitOfMeasureChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetForceUnitOfMeasureChange(v bool)`

SetForceUnitOfMeasureChange sets ForceUnitOfMeasureChange field to given value.

### HasForceUnitOfMeasureChange

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasForceUnitOfMeasureChange() bool`

HasForceUnitOfMeasureChange returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesOrPurchaseDetails() QuickbooksDesktopNonInventoryItemsIdPostRequestSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesOrPurchaseDetailsOk() (*QuickbooksDesktopNonInventoryItemsIdPostRequestSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetSalesOrPurchaseDetails(v QuickbooksDesktopNonInventoryItemsIdPostRequestSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.

### HasSalesOrPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasSalesOrPurchaseDetails() bool`

HasSalesOrPurchaseDetails returns a boolean if a field has been set.

### GetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesAndPurchaseDetails() QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) GetSalesAndPurchaseDetailsOk() (*QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) SetSalesAndPurchaseDetails(v QuickbooksDesktopNonInventoryItemsIdPostRequestSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.

### HasSalesAndPurchaseDetails

`func (o *QuickbooksDesktopNonInventoryItemsIdPostRequest) HasSalesAndPurchaseDetails() bool`

HasSalesAndPurchaseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


