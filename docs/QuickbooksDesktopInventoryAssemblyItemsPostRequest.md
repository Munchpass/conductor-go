# QuickbooksDesktopInventoryAssemblyItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this inventory assembly item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory assembly items could both have the &#x60;name&#x60; \&quot;Deluxe Kit\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Assemblies:Deluxe Kit\&quot; and \&quot;Inventory:Deluxe Kit\&quot;.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode**](QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this inventory assembly item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The inventory assembly item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent inventory assembly item one level above this one in the hierarchy. For example, if this inventory assembly item has a &#x60;fullName&#x60; of \&quot;Assemblies:Deluxe Kit\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Assemblies\&quot;. If this inventory assembly item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this inventory assembly item, which consists of a base unit and related units. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this inventory assembly item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesDescription** | Pointer to **string** | The description of this inventory assembly item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | [optional] 
**SalesPrice** | Pointer to **string** | The price at which this inventory assembly item is sold to customers, represented as a decimal string. | [optional] 
**IncomeAccountId** | **string** | The income account used to track revenue from sales of this inventory assembly item. | 
**PurchaseDescription** | Pointer to **string** | The description of this inventory assembly item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | [optional] 
**PurchaseCost** | Pointer to **string** | The cost at which this inventory assembly item is purchased from vendors, represented as a decimal string. | [optional] 
**PurchaseTaxCodeId** | Pointer to **string** | The tax code applied to purchases of this inventory assembly item. Applicable in regions where purchase taxes are used, such as Canada or the UK. | [optional] 
**CogsAccountId** | **string** | The Cost of Goods Sold (COGS) account for this inventory assembly item, tracking the original direct costs of producing goods sold. | 
**PreferredVendorId** | Pointer to **string** | The preferred vendor from whom this inventory assembly item is typically purchased. | [optional] 
**AssetAccountId** | **string** | The asset account used to track the current value of this inventory assembly item in inventory. | 
**BuildNotificationThreshold** | Pointer to **float32** | The inventory assembly item&#39;s minimum quantity threshold that triggers a build notification in QuickBooks. When the sum of &#x60;quantityOnHand&#x60; (current inventory) and &#x60;quantityOnOrder&#x60; (pending purchase orders) drops below this threshold, QuickBooks will notify users that more units need to be built or assembled. This helps ensure adequate inventory levels for inventory assembly items. | [optional] 
**MaximumQuantityOnHand** | Pointer to **float32** | The maximum quantity of this inventory assembly item desired in inventory. | [optional] 
**QuantityOnHand** | Pointer to **float32** | The current quantity of this inventory assembly item available in inventory. To change the &#x60;quantityOnHand&#x60; for an inventory assembly item, you must create an inventory-adjustment instead of updating this inventory assembly item directly. | [optional] 
**TotalValue** | Pointer to **string** | The total value of this inventory assembly item. If &#x60;totalValue&#x60; is provided, &#x60;quantityOnHand&#x60; must also be provided and must be greater than zero. If both &#x60;quantityOnHand&#x60; and &#x60;purchaseCost&#x60; are provided, then &#x60;totalValue&#x60; will be set to &#x60;quantityOnHand&#x60; times &#x60;purchaseCost&#x60;, regardless of what &#x60;totalValue&#x60; is explicitly set to. | [optional] 
**InventoryDate** | Pointer to **string** | The date when this inventory assembly item was converted into an inventory item from some other type of item, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner**](QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner.md) | The inventory assembly item&#39;s lines. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAssemblyItemsPostRequest

`func NewQuickbooksDesktopInventoryAssemblyItemsPostRequest(name string, incomeAccountId string, cogsAccountId string, assetAccountId string, ) *QuickbooksDesktopInventoryAssemblyItemsPostRequest`

NewQuickbooksDesktopInventoryAssemblyItemsPostRequest instantiates a new QuickbooksDesktopInventoryAssemblyItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAssemblyItemsPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryAssemblyItemsPostRequestWithDefaults() *QuickbooksDesktopInventoryAssemblyItemsPostRequest`

NewQuickbooksDesktopInventoryAssemblyItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryAssemblyItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetBarcode() QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetBarcode(v QuickbooksDesktopInventoryAssemblyItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.

### HasSalesDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasSalesDescription() bool`

HasSalesDescription returns a boolean if a field has been set.

### GetSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.

### HasSalesPrice

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasSalesPrice() bool`

HasSalesPrice returns a boolean if a field has been set.

### GetIncomeAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetIncomeAccountId() string`

GetIncomeAccountId returns the IncomeAccountId field if non-nil, zero value otherwise.

### GetIncomeAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetIncomeAccountIdOk() (*string, bool)`

GetIncomeAccountIdOk returns a tuple with the IncomeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetIncomeAccountId(v string)`

SetIncomeAccountId sets IncomeAccountId field to given value.


### GetPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescription

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasPurchaseDescription() bool`

HasPurchaseDescription returns a boolean if a field has been set.

### GetPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.

### HasPurchaseCost

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasPurchaseCost() bool`

HasPurchaseCost returns a boolean if a field has been set.

### GetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseTaxCodeId() string`

GetPurchaseTaxCodeId returns the PurchaseTaxCodeId field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPurchaseTaxCodeIdOk() (*string, bool)`

GetPurchaseTaxCodeIdOk returns a tuple with the PurchaseTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetPurchaseTaxCodeId(v string)`

SetPurchaseTaxCodeId sets PurchaseTaxCodeId field to given value.

### HasPurchaseTaxCodeId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasPurchaseTaxCodeId() bool`

HasPurchaseTaxCodeId returns a boolean if a field has been set.

### GetCogsAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetCogsAccountId() string`

GetCogsAccountId returns the CogsAccountId field if non-nil, zero value otherwise.

### GetCogsAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetCogsAccountIdOk() (*string, bool)`

GetCogsAccountIdOk returns a tuple with the CogsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetCogsAccountId(v string)`

SetCogsAccountId sets CogsAccountId field to given value.


### GetPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPreferredVendorId() string`

GetPreferredVendorId returns the PreferredVendorId field if non-nil, zero value otherwise.

### GetPreferredVendorIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetPreferredVendorIdOk() (*string, bool)`

GetPreferredVendorIdOk returns a tuple with the PreferredVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetPreferredVendorId(v string)`

SetPreferredVendorId sets PreferredVendorId field to given value.

### HasPreferredVendorId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasPreferredVendorId() bool`

HasPreferredVendorId returns a boolean if a field has been set.

### GetAssetAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetAssetAccountId() string`

GetAssetAccountId returns the AssetAccountId field if non-nil, zero value otherwise.

### GetAssetAccountIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetAssetAccountIdOk() (*string, bool)`

GetAssetAccountIdOk returns a tuple with the AssetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccountId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetAssetAccountId(v string)`

SetAssetAccountId sets AssetAccountId field to given value.


### GetBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetBuildNotificationThreshold() float32`

GetBuildNotificationThreshold returns the BuildNotificationThreshold field if non-nil, zero value otherwise.

### GetBuildNotificationThresholdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetBuildNotificationThresholdOk() (*float32, bool)`

GetBuildNotificationThresholdOk returns a tuple with the BuildNotificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetBuildNotificationThreshold(v float32)`

SetBuildNotificationThreshold sets BuildNotificationThreshold field to given value.

### HasBuildNotificationThreshold

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasBuildNotificationThreshold() bool`

HasBuildNotificationThreshold returns a boolean if a field has been set.

### GetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.

### HasMaximumQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasMaximumQuantityOnHand() bool`

HasMaximumQuantityOnHand returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetQuantityOnHand() float32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetQuantityOnHandOk() (*float32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetQuantityOnHand(v float32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### GetTotalValue

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetTotalValue() string`

GetTotalValue returns the TotalValue field if non-nil, zero value otherwise.

### GetTotalValueOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetTotalValueOk() (*string, bool)`

GetTotalValueOk returns a tuple with the TotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValue

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetTotalValue(v string)`

SetTotalValue sets TotalValue field to given value.

### HasTotalValue

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasTotalValue() bool`

HasTotalValue returns a boolean if a field has been set.

### GetInventoryDate

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetInventoryDate() string`

GetInventoryDate returns the InventoryDate field if non-nil, zero value otherwise.

### GetInventoryDateOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetInventoryDateOk() (*string, bool)`

GetInventoryDateOk returns a tuple with the InventoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDate

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetInventoryDate(v string)`

SetInventoryDate sets InventoryDate field to given value.

### HasInventoryDate

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasInventoryDate() bool`

HasInventoryDate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetLines() []QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) GetLinesOk() (*[]QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) SetLines(v []QuickbooksDesktopInventoryAssemblyItemsPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInventoryAssemblyItemsPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


