# QuickbooksDesktopInventoryItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the inventory item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory items could both have the &#x60;name&#x60; \&quot;Cabinet\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Kitchen:Cabinet\&quot; and \&quot;Inventory:Cabinet\&quot;.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopInventoryItemsPostRequestBarcode**](QuickbooksDesktopInventoryItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The inventory item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent inventory item one level above this one in the hierarchy. For example, if this inventory item has a &#x60;fullName&#x60; of \&quot;Kitchen:Cabinet\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Kitchen\&quot;. If this inventory item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**Sku** | Pointer to **string** | The inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this inventory item, which consists of a base unit and related units. | [optional] 
**ForceUnitOfMeasureChange** | Pointer to **bool** | Indicates whether to allow changing the inventory item&#39;s unit-of-measure set (using the &#x60;unitOfMeasureSetId&#x60; field) when the base unit of the new unit-of-measure set does not match that of the currently assigned set. Without setting this field to &#x60;true&#x60; in this scenario, the request will fail with an error; hence, this field is equivalent to accepting the warning prompt in the QuickBooks UI.  NOTE: Changing the base unit requires you to update the item&#39;s quantities-on-hand and cost to reflect the new unit; otherwise, these values will be inaccurate. Alternatively, consider creating a new item with the desired unit-of-measure set and deactivating the old item. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this inventory item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesDescription** | Pointer to **string** | The description of this inventory item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this inventory item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | Pointer to **string** | The income account used to track revenue from sales of this inventory item. | [optional] 
**UpdateExistingTransactionsIncomeAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new income account (specified by the &#x60;incomeAccountId&#x60; field) to all existing transactions that use this inventory item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 
**PurchaseDescription** | Pointer to **string** | The description of this inventory item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this inventory item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this inventory item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**CogsAccountId** | Pointer to **string** | The Cost of Goods Sold (COGS) account for this inventory item, tracking the original direct costs of producing goods sold. | [optional] 
**UpdateExistingTransactionsCogsAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new COGS account (specified by the &#x60;cogsAccountId&#x60; field) to all existing transactions that use this inventory item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this inventory item is typically purchased. | [optional] 
**AssetAccountId** | Pointer to **string** | The asset account used to track the current value of this inventory item in inventory. | [optional] 
**ReorderPoint** | Pointer to **float32** | The minimum quantity of this inventory item at which QuickBooks prompts for reordering. | [optional] 
**MaximumQuantityOnHand** | Pointer to **float32** | The maximum quantity of this inventory item desired in inventory. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryItemsIdPostRequest

`func NewQuickbooksDesktopInventoryItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopInventoryItemsIdPostRequest`

NewQuickbooksDesktopInventoryItemsIdPostRequest instantiates a new QuickbooksDesktopInventoryItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryItemsIdPostRequestWithDefaults() *QuickbooksDesktopInventoryItemsIdPostRequest`

NewQuickbooksDesktopInventoryItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetBarcode() QuickbooksDesktopInventoryItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopInventoryItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetBarcode(v QuickbooksDesktopInventoryItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSku

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetForceUnitOfMeasureChange() bool`

GetForceUnitOfMeasureChange returns the ForceUnitOfMeasureChange field if non-nil, zero value otherwise.

### GetForceUnitOfMeasureChangeOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetForceUnitOfMeasureChangeOk() (*bool, bool)`

GetForceUnitOfMeasureChangeOk returns a tuple with the ForceUnitOfMeasureChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetForceUnitOfMeasureChange(v bool)`

SetForceUnitOfMeasureChange sets ForceUnitOfMeasureChange field to given value.

### HasForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasForceUnitOfMeasureChange() bool`

HasForceUnitOfMeasureChange returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.

### HasIncomeAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasIncomeAccountId() bool`

HasIncomeAccountId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUpdateExistingTransactionsIncomeAccount() bool`

GetUpdateExistingTransactionsIncomeAccount returns the UpdateExistingTransactionsIncomeAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsIncomeAccountOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUpdateExistingTransactionsIncomeAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsIncomeAccountOk returns a tuple with the UpdateExistingTransactionsIncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetUpdateExistingTransactionsIncomeAccount(v bool)`

SetUpdateExistingTransactionsIncomeAccount sets UpdateExistingTransactionsIncomeAccount field to given value.

### HasUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasUpdateExistingTransactionsIncomeAccount() bool`

HasUpdateExistingTransactionsIncomeAccount returns a boolean if a field has been set.

### GetPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetCogsAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetCogsAccountId() string`

GetCogsAccountId returns the CogsAccountId field if non-nil, zero value otherwise.

### GetCogsAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetCogsAccountIdOk() (*string, bool)`

GetCogsAccountIdOk returns a tuple with the CogsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetCogsAccountId(v string)`

SetCogsAccountId sets CogsAccountId field to given value.

### HasCogsAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasCogsAccountId() bool`

HasCogsAccountId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsCogsAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUpdateExistingTransactionsCogsAccount() bool`

GetUpdateExistingTransactionsCogsAccount returns the UpdateExistingTransactionsCogsAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsCogsAccountOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetUpdateExistingTransactionsCogsAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsCogsAccountOk returns a tuple with the UpdateExistingTransactionsCogsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsCogsAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetUpdateExistingTransactionsCogsAccount(v bool)`

SetUpdateExistingTransactionsCogsAccount sets UpdateExistingTransactionsCogsAccount field to given value.

### HasUpdateExistingTransactionsCogsAccount

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasUpdateExistingTransactionsCogsAccount() bool`

HasUpdateExistingTransactionsCogsAccount returns a boolean if a field has been set.

### GetPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.

### GetAssetAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetAssetAccountId() string`

GetAssetAccountId returns the AssetAccountId field if non-nil, zero value otherwise.

### GetAssetAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetAssetAccountIdOk() (*string, bool)`

GetAssetAccountIdOk returns a tuple with the AssetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetAssetAccountId(v string)`

SetAssetAccountId sets AssetAccountId field to given value.

### HasAssetAccountId

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasAssetAccountId() bool`

HasAssetAccountId returns a boolean if a field has been set.

### GetReorderPoint

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetReorderPoint() float32`

GetReorderPoint returns the ReorderPoint field if non-nil, zero value otherwise.

### GetReorderPointOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetReorderPointOk() (*float32, bool)`

GetReorderPointOk returns a tuple with the ReorderPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderPoint

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetReorderPoint(v float32)`

SetReorderPoint sets ReorderPoint field to given value.

### HasReorderPoint

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasReorderPoint() bool`

HasReorderPoint returns a boolean if a field has been set.

### GetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.

### HasMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsIdPostRequest) HasMaximumQuantityOnHand() bool`

HasMaximumQuantityOnHand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


