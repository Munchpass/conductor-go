# QuickbooksDesktopInventoryItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory items could both have the &#x60;name&#x60; \&quot;Cabinet\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Kitchen:Cabinet\&quot; and \&quot;Inventory:Cabinet\&quot;.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopInventoryItemsPostRequestBarcode**](QuickbooksDesktopInventoryItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The inventory item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent inventory item one level above this one in the hierarchy. For example, if this inventory item has a &#x60;fullName&#x60; of \&quot;Kitchen:Cabinet\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Kitchen\&quot;. If this inventory item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**Sku** | Pointer to **string** | The inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this inventory item, which consists of a base unit and related units. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this inventory item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesDescription** | Pointer to **string** | The description of this inventory item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this inventory item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | **string** | The income account used to track revenue from sales of this inventory item. | 
**PurchaseDescription** | Pointer to **string** | The description of this inventory item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this inventory item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this inventory item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**CogsAccountId** | **string** | The Cost of Goods Sold (COGS) account for this inventory item, tracking the original direct costs of producing goods sold. | 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this inventory item is typically purchased. | [optional] 
**AssetAccountId** | **string** | The asset account used to track the current value of this inventory item in inventory. | 
**ReorderPoint** | Pointer to **float32** | The minimum quantity of this inventory item at which QuickBooks prompts for reordering. | [optional] 
**MaximumQuantityOnHand** | Pointer to **float32** | The maximum quantity of this inventory item desired in inventory. | [optional] 
**QuantityOnHand** | Pointer to **float32** | The current quantity of this inventory item available in inventory. To change the &#x60;quantityOnHand&#x60; for an inventory item, you must create an inventory-adjustment instead of updating this inventory item directly. | [optional] 
**TotalValue** | Pointer to **string** | The total value of this inventory item. If &#x60;totalValue&#x60; is provided, &#x60;quantityOnHand&#x60; must also be provided and must be greater than zero. If both &#x60;quantityOnHand&#x60; and &#x60;purchaseCost&#x60; are provided, then &#x60;totalValue&#x60; will be set to &#x60;quantityOnHand&#x60; times &#x60;purchaseCost&#x60;, regardless of what &#x60;totalValue&#x60; is explicitly set to. | [optional] 
**InventoryDate** | Pointer to **string** | The date when this inventory item was converted into an inventory item from some other type of item, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryItemsPostRequest

`func NewQuickbooksDesktopInventoryItemsPostRequest(name string, incomeAccountId string, cogsAccountId string, assetAccountId string, ) *QuickbooksDesktopInventoryItemsPostRequest`

NewQuickbooksDesktopInventoryItemsPostRequest instantiates a new QuickbooksDesktopInventoryItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryItemsPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryItemsPostRequestWithDefaults() *QuickbooksDesktopInventoryItemsPostRequest`

NewQuickbooksDesktopInventoryItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetBarcode() QuickbooksDesktopInventoryItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopInventoryItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetBarcode(v QuickbooksDesktopInventoryItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSku

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.


### GetPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetCogsAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetCogsAccountId() string`

GetCogsAccountId returns the CogsAccountId field if non-nil, zero value otherwise.

### GetCogsAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetCogsAccountIdOk() (*string, bool)`

GetCogsAccountIdOk returns a tuple with the CogsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetCogsAccountId(v string)`

SetCogsAccountId sets CogsAccountId field to given value.


### GetPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.

### GetAssetAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetAssetAccountId() string`

GetAssetAccountId returns the AssetAccountId field if non-nil, zero value otherwise.

### GetAssetAccountIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetAssetAccountIdOk() (*string, bool)`

GetAssetAccountIdOk returns a tuple with the AssetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccountId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetAssetAccountId(v string)`

SetAssetAccountId sets AssetAccountId field to given value.


### GetReorderPoint

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetReorderPoint() float32`

GetReorderPoint returns the ReorderPoint field if non-nil, zero value otherwise.

### GetReorderPointOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetReorderPointOk() (*float32, bool)`

GetReorderPointOk returns a tuple with the ReorderPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderPoint

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetReorderPoint(v float32)`

SetReorderPoint sets ReorderPoint field to given value.

### HasReorderPoint

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasReorderPoint() bool`

HasReorderPoint returns a boolean if a field has been set.

### GetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.

### HasMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasMaximumQuantityOnHand() bool`

HasMaximumQuantityOnHand returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetQuantityOnHand() float32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetQuantityOnHandOk() (*float32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetQuantityOnHand(v float32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### GetTotalValue

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetTotalValue() string`

GetTotalValue returns the TotalValue field if non-nil, zero value otherwise.

### GetTotalValueOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetTotalValueOk() (*string, bool)`

GetTotalValueOk returns a tuple with the TotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValue

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetTotalValue(v string)`

SetTotalValue sets TotalValue field to given value.

### HasTotalValue

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasTotalValue() bool`

HasTotalValue returns a boolean if a field has been set.

### GetInventoryDate

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetInventoryDate() string`

GetInventoryDate returns the InventoryDate field if non-nil, zero value otherwise.

### GetInventoryDateOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetInventoryDateOk() (*string, bool)`

GetInventoryDateOk returns a tuple with the InventoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDate

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetInventoryDate(v string)`

SetInventoryDate sets InventoryDate field to given value.

### HasInventoryDate

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasInventoryDate() bool`

HasInventoryDate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopInventoryItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopInventoryItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


