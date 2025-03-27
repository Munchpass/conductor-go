# QuickbooksDesktopInventoryAssemblyItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the inventory assembly item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this inventory assembly item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory assembly items could both have the &#x60;name&#x60; \&quot;Deluxe Kit\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Assemblies:Deluxe Kit\&quot; and \&quot;Inventory:Deluxe Kit\&quot;.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode**](QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this inventory assembly item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The inventory assembly item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent inventory assembly item one level above this one in the hierarchy. For example, if this inventory assembly item has a &#x60;fullName&#x60; of \&quot;Assemblies:Deluxe Kit\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Assemblies\&quot;. If this inventory assembly item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**Sku** | Pointer to **string** | The inventory assembly item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this inventory assembly item, which consists of a base unit and related units. | [optional] 
**ForceUnitOfMeasureChange** | Pointer to **bool** | Indicates whether to allow changing the inventory assembly item&#39;s unit-of-measure set (using the &#x60;unitOfMeasureSetId&#x60; field) when the base unit of the new unit-of-measure set does not match that of the currently assigned set. Without setting this field to &#x60;true&#x60; in this scenario, the request will fail with an error; hence, this field is equivalent to accepting the warning prompt in the QuickBooks UI.  NOTE: Changing the base unit requires you to update the item&#39;s quantities-on-hand and cost to reflect the new unit; otherwise, these values will be inaccurate. Alternatively, consider creating a new item with the desired unit-of-measure set and deactivating the old item. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this inventory assembly item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesDescription** | Pointer to **string** | The description of this inventory assembly item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this inventory assembly item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | Pointer to **string** | The income account used to track revenue from sales of this inventory assembly item. | [optional] 
**UpdateExistingTransactionsIncomeAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new income account (specified by the &#x60;incomeAccountId&#x60; field) to all existing transactions that use this inventory assembly item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 
**PurchaseDescription** | Pointer to **string** | The description of this inventory assembly item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this inventory assembly item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this inventory assembly item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**CogsAccountId** | Pointer to **string** | The Cost of Goods Sold (COGS) account for this inventory assembly item, tracking the original direct costs of producing goods sold. | [optional] 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this inventory assembly item is typically purchased. | [optional] 
**AssetAccountId** | Pointer to **string** | The asset account used to track the current value of this inventory assembly item in inventory. | [optional] 
**BuildNotificationThreshold** | Pointer to **float32** | The inventory assembly item&#39;s minimum quantity threshold that triggers a build notification in QuickBooks. When the sum of &#x60;quantityOnHand&#x60; (current inventory) and &#x60;quantityOnOrder&#x60; (pending purchase orders) drops below this threshold, QuickBooks will notify users that more units need to be built or assembled. This helps ensure adequate inventory levels for inventory assembly items. | [optional] 
**MaximumQuantityOnHand** | Pointer to **float32** | The maximum quantity of this inventory assembly item desired in inventory. | [optional] 
**ClearItemLines** | Pointer to **bool** | When &#x60;true&#x60;, removes all existing item lines associated with this inventory assembly item. To modify or add individual item lines, use the field &#x60;itemLines&#x60; instead. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner**](QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner.md) | The inventory assembly item&#39;s lines. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequest

`func NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest`

NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequest instantiates a new QuickbooksDesktopInventoryAssemblyItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequestWithDefaults() *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest`

NewQuickbooksDesktopInventoryAssemblyItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryAssemblyItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetBarcode() QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetBarcode(v QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSku

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetForceUnitOfMeasureChange() bool`

GetForceUnitOfMeasureChange returns the ForceUnitOfMeasureChange field if non-nil, zero value otherwise.

### GetForceUnitOfMeasureChangeOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetForceUnitOfMeasureChangeOk() (*bool, bool)`

GetForceUnitOfMeasureChangeOk returns a tuple with the ForceUnitOfMeasureChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetForceUnitOfMeasureChange(v bool)`

SetForceUnitOfMeasureChange sets ForceUnitOfMeasureChange field to given value.

### HasForceUnitOfMeasureChange

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasForceUnitOfMeasureChange() bool`

HasForceUnitOfMeasureChange returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.

### HasIncomeAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasIncomeAccountId() bool`

HasIncomeAccountId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetUpdateExistingTransactionsIncomeAccount() bool`

GetUpdateExistingTransactionsIncomeAccount returns the UpdateExistingTransactionsIncomeAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsIncomeAccountOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetUpdateExistingTransactionsIncomeAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsIncomeAccountOk returns a tuple with the UpdateExistingTransactionsIncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetUpdateExistingTransactionsIncomeAccount(v bool)`

SetUpdateExistingTransactionsIncomeAccount sets UpdateExistingTransactionsIncomeAccount field to given value.

### HasUpdateExistingTransactionsIncomeAccount

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasUpdateExistingTransactionsIncomeAccount() bool`

HasUpdateExistingTransactionsIncomeAccount returns a boolean if a field has been set.

### GetPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetCogsAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetCogsAccountId() string`

GetCogsAccountId returns the CogsAccountId field if non-nil, zero value otherwise.

### GetCogsAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetCogsAccountIdOk() (*string, bool)`

GetCogsAccountIdOk returns a tuple with the CogsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetCogsAccountId(v string)`

SetCogsAccountId sets CogsAccountId field to given value.

### HasCogsAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasCogsAccountId() bool`

HasCogsAccountId returns a boolean if a field has been set.

### GetPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.

### GetAssetAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetAssetAccountId() string`

GetAssetAccountId returns the AssetAccountId field if non-nil, zero value otherwise.

### GetAssetAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetAssetAccountIdOk() (*string, bool)`

GetAssetAccountIdOk returns a tuple with the AssetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetAssetAccountId(v string)`

SetAssetAccountId sets AssetAccountId field to given value.

### HasAssetAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasAssetAccountId() bool`

HasAssetAccountId returns a boolean if a field has been set.

### GetBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetBuildNotificationThreshold() float32`

GetBuildNotificationThreshold returns the BuildNotificationThreshold field if non-nil, zero value otherwise.

### GetBuildNotificationThresholdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetBuildNotificationThresholdOk() (*float32, bool)`

GetBuildNotificationThresholdOk returns a tuple with the BuildNotificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetBuildNotificationThreshold(v float32)`

SetBuildNotificationThreshold sets BuildNotificationThreshold field to given value.

### HasBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasBuildNotificationThreshold() bool`

HasBuildNotificationThreshold returns a boolean if a field has been set.

### GetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.

### HasMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasMaximumQuantityOnHand() bool`

HasMaximumQuantityOnHand returns a boolean if a field has been set.

### GetClearItemLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetClearItemLines() bool`

GetClearItemLines returns the ClearItemLines field if non-nil, zero value otherwise.

### GetClearItemLinesOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetClearItemLinesOk() (*bool, bool)`

GetClearItemLinesOk returns a tuple with the ClearItemLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearItemLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetClearItemLines(v bool)`

SetClearItemLines sets ClearItemLines field to given value.

### HasClearItemLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasClearItemLines() bool`

HasClearItemLines returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetLines() []QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) SetLines(v []QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


