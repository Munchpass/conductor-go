# QbdInventoryAssemblyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this inventory assembly item. This ID is unique across all inventory assembly items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_inventory_assembly_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this inventory assembly item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this inventory assembly item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this inventory assembly item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this inventory assembly item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory assembly items could both have the &#x60;name&#x60; \&quot;Deluxe Kit\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Assemblies:Deluxe Kit\&quot; and \&quot;Inventory:Deluxe Kit\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this inventory assembly item, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if an inventory assembly item is under \&quot;Assemblies\&quot; and has the &#x60;name&#x60; \&quot;Deluxe Kit\&quot;, its &#x60;fullName&#x60; would be \&quot;Assemblies:Deluxe Kit\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all inventory assembly item objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**Barcode** | **string** | The inventory assembly item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this inventory assembly item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdInventoryAssemblyItemClass**](QbdInventoryAssemblyItemClass.md) |  | 
**Parent** | [**QbdInventoryAssemblyItemParent**](QbdInventoryAssemblyItemParent.md) |  | 
**Sublevel** | **float32** | The depth level of this inventory assembly item in the hierarchy. A top-level inventory assembly item has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, an inventory assembly item with a &#x60;fullName&#x60; of \&quot;Assemblies:Deluxe Kit\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**Sku** | **string** | The inventory assembly item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | 
**UnitOfMeasureSet** | [**QbdInventoryAssemblyItemUnitOfMeasureSet**](QbdInventoryAssemblyItemUnitOfMeasureSet.md) |  | 
**SalesTaxCode** | [**QbdInventoryAssemblyItemSalesTaxCode**](QbdInventoryAssemblyItemSalesTaxCode.md) |  | 
**SalesDescription** | **string** | The description of this inventory assembly item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | 
**SalesPrice** | **string** | The price at which this inventory assembly item is sold to customers, represented as a decimal string. | 
**IncomeAccount** | [**QbdInventoryAssemblyItemIncomeAccount**](QbdInventoryAssemblyItemIncomeAccount.md) |  | 
**PurchaseDescription** | **string** | The description of this inventory assembly item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | 
**PurchaseCost** | **string** | The cost at which this inventory assembly item is purchased from vendors, represented as a decimal string. | 
**PurchaseTaxCode** | [**QbdInventoryAssemblyItemPurchaseTaxCode**](QbdInventoryAssemblyItemPurchaseTaxCode.md) |  | 
**CogsAccount** | [**QbdInventoryAssemblyItemCogsAccount**](QbdInventoryAssemblyItemCogsAccount.md) |  | 
**PreferredVendor** | [**QbdInventoryAssemblyItemPreferredVendor**](QbdInventoryAssemblyItemPreferredVendor.md) |  | 
**AssetAccount** | [**QbdInventoryAssemblyItemAssetAccount**](QbdInventoryAssemblyItemAssetAccount.md) |  | 
**BuildNotificationThreshold** | **float32** | The inventory assembly item&#39;s minimum quantity threshold that triggers a build notification in QuickBooks. When the sum of &#x60;quantityOnHand&#x60; (current inventory) and &#x60;quantityOnOrder&#x60; (pending purchase orders) drops below this threshold, QuickBooks will notify users that more units need to be built or assembled. This helps ensure adequate inventory levels for inventory assembly items. | 
**MaximumQuantityOnHand** | **float32** | The maximum quantity of this inventory assembly item desired in inventory. | 
**QuantityOnHand** | **float32** | The current quantity of this inventory assembly item available in inventory. To change the &#x60;quantityOnHand&#x60; for an inventory assembly item, you must create an inventory-adjustment instead of updating this inventory assembly item directly. | 
**AverageCost** | **string** | The average cost per unit of this inventory assembly item, represented as a decimal string. | 
**QuantityOnPurchaseOrder** | **float32** | The number of units of this inventory assembly item that have been ordered from vendors (as recorded in purchase orders) but not yet received. | 
**QuantityOnSalesOrder** | **float32** | The number of units of this inventory assembly item that have been sold (as recorded in sales orders) but not yet fulfilled or delivered to customers. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**Lines** | [**[]QbdInventoryAssemblyItemLine**](QbdInventoryAssemblyItemLine.md) | The inventory assembly item&#39;s lines. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the inventory assembly item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdInventoryAssemblyItem

`func NewQbdInventoryAssemblyItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, barcode string, isActive bool, class QbdInventoryAssemblyItemClass, parent QbdInventoryAssemblyItemParent, sublevel float32, sku string, unitOfMeasureSet QbdInventoryAssemblyItemUnitOfMeasureSet, salesTaxCode QbdInventoryAssemblyItemSalesTaxCode, salesDescription string, salesPrice string, incomeAccount QbdInventoryAssemblyItemIncomeAccount, purchaseDescription string, purchaseCost string, purchaseTaxCode QbdInventoryAssemblyItemPurchaseTaxCode, cogsAccount QbdInventoryAssemblyItemCogsAccount, preferredVendor QbdInventoryAssemblyItemPreferredVendor, assetAccount QbdInventoryAssemblyItemAssetAccount, buildNotificationThreshold float32, maximumQuantityOnHand float32, quantityOnHand float32, averageCost string, quantityOnPurchaseOrder float32, quantityOnSalesOrder float32, externalId string, lines []QbdInventoryAssemblyItemLine, customFields []QbdCustomField, ) *QbdInventoryAssemblyItem`

NewQbdInventoryAssemblyItem instantiates a new QbdInventoryAssemblyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventoryAssemblyItemWithDefaults

`func NewQbdInventoryAssemblyItemWithDefaults() *QbdInventoryAssemblyItem`

NewQbdInventoryAssemblyItemWithDefaults instantiates a new QbdInventoryAssemblyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInventoryAssemblyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInventoryAssemblyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInventoryAssemblyItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInventoryAssemblyItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInventoryAssemblyItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInventoryAssemblyItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdInventoryAssemblyItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdInventoryAssemblyItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdInventoryAssemblyItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdInventoryAssemblyItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdInventoryAssemblyItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdInventoryAssemblyItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdInventoryAssemblyItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdInventoryAssemblyItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdInventoryAssemblyItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdInventoryAssemblyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdInventoryAssemblyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdInventoryAssemblyItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdInventoryAssemblyItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdInventoryAssemblyItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdInventoryAssemblyItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBarcode

`func (o *QbdInventoryAssemblyItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdInventoryAssemblyItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdInventoryAssemblyItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdInventoryAssemblyItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdInventoryAssemblyItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdInventoryAssemblyItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdInventoryAssemblyItem) GetClass() QbdInventoryAssemblyItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdInventoryAssemblyItem) GetClassOk() (*QbdInventoryAssemblyItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdInventoryAssemblyItem) SetClass(v QbdInventoryAssemblyItemClass)`

SetClass sets Class field to given value.


### GetParent

`func (o *QbdInventoryAssemblyItem) GetParent() QbdInventoryAssemblyItemParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdInventoryAssemblyItem) GetParentOk() (*QbdInventoryAssemblyItemParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdInventoryAssemblyItem) SetParent(v QbdInventoryAssemblyItemParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdInventoryAssemblyItem) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdInventoryAssemblyItem) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdInventoryAssemblyItem) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetSku

`func (o *QbdInventoryAssemblyItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QbdInventoryAssemblyItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QbdInventoryAssemblyItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetUnitOfMeasureSet

`func (o *QbdInventoryAssemblyItem) GetUnitOfMeasureSet() QbdInventoryAssemblyItemUnitOfMeasureSet`

GetUnitOfMeasureSet returns the UnitOfMeasureSet field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetOk

`func (o *QbdInventoryAssemblyItem) GetUnitOfMeasureSetOk() (*QbdInventoryAssemblyItemUnitOfMeasureSet, bool)`

GetUnitOfMeasureSetOk returns a tuple with the UnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSet

`func (o *QbdInventoryAssemblyItem) SetUnitOfMeasureSet(v QbdInventoryAssemblyItemUnitOfMeasureSet)`

SetUnitOfMeasureSet sets UnitOfMeasureSet field to given value.


### GetSalesTaxCode

`func (o *QbdInventoryAssemblyItem) GetSalesTaxCode() QbdInventoryAssemblyItemSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdInventoryAssemblyItem) GetSalesTaxCodeOk() (*QbdInventoryAssemblyItemSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdInventoryAssemblyItem) SetSalesTaxCode(v QbdInventoryAssemblyItemSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetSalesDescription

`func (o *QbdInventoryAssemblyItem) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QbdInventoryAssemblyItem) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QbdInventoryAssemblyItem) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.


### GetSalesPrice

`func (o *QbdInventoryAssemblyItem) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QbdInventoryAssemblyItem) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QbdInventoryAssemblyItem) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.


### GetIncomeAccount

`func (o *QbdInventoryAssemblyItem) GetIncomeAccount() QbdInventoryAssemblyItemIncomeAccount`

GetIncomeAccount returns the IncomeAccount field if non-nil, zero value otherwise.

### GetIncomeAccountOk

`func (o *QbdInventoryAssemblyItem) GetIncomeAccountOk() (*QbdInventoryAssemblyItemIncomeAccount, bool)`

GetIncomeAccountOk returns a tuple with the IncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccount

`func (o *QbdInventoryAssemblyItem) SetIncomeAccount(v QbdInventoryAssemblyItemIncomeAccount)`

SetIncomeAccount sets IncomeAccount field to given value.


### GetPurchaseDescription

`func (o *QbdInventoryAssemblyItem) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QbdInventoryAssemblyItem) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QbdInventoryAssemblyItem) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.


### GetPurchaseCost

`func (o *QbdInventoryAssemblyItem) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QbdInventoryAssemblyItem) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QbdInventoryAssemblyItem) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.


### GetPurchaseTaxCode

`func (o *QbdInventoryAssemblyItem) GetPurchaseTaxCode() QbdInventoryAssemblyItemPurchaseTaxCode`

GetPurchaseTaxCode returns the PurchaseTaxCode field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeOk

`func (o *QbdInventoryAssemblyItem) GetPurchaseTaxCodeOk() (*QbdInventoryAssemblyItemPurchaseTaxCode, bool)`

GetPurchaseTaxCodeOk returns a tuple with the PurchaseTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCode

`func (o *QbdInventoryAssemblyItem) SetPurchaseTaxCode(v QbdInventoryAssemblyItemPurchaseTaxCode)`

SetPurchaseTaxCode sets PurchaseTaxCode field to given value.


### GetCogsAccount

`func (o *QbdInventoryAssemblyItem) GetCogsAccount() QbdInventoryAssemblyItemCogsAccount`

GetCogsAccount returns the CogsAccount field if non-nil, zero value otherwise.

### GetCogsAccountOk

`func (o *QbdInventoryAssemblyItem) GetCogsAccountOk() (*QbdInventoryAssemblyItemCogsAccount, bool)`

GetCogsAccountOk returns a tuple with the CogsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccount

`func (o *QbdInventoryAssemblyItem) SetCogsAccount(v QbdInventoryAssemblyItemCogsAccount)`

SetCogsAccount sets CogsAccount field to given value.


### GetPreferredVendor

`func (o *QbdInventoryAssemblyItem) GetPreferredVendor() QbdInventoryAssemblyItemPreferredVendor`

GetPreferredVendor returns the PreferredVendor field if non-nil, zero value otherwise.

### GetPreferredVendorOk

`func (o *QbdInventoryAssemblyItem) GetPreferredVendorOk() (*QbdInventoryAssemblyItemPreferredVendor, bool)`

GetPreferredVendorOk returns a tuple with the PreferredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendor

`func (o *QbdInventoryAssemblyItem) SetPreferredVendor(v QbdInventoryAssemblyItemPreferredVendor)`

SetPreferredVendor sets PreferredVendor field to given value.


### GetAssetAccount

`func (o *QbdInventoryAssemblyItem) GetAssetAccount() QbdInventoryAssemblyItemAssetAccount`

GetAssetAccount returns the AssetAccount field if non-nil, zero value otherwise.

### GetAssetAccountOk

`func (o *QbdInventoryAssemblyItem) GetAssetAccountOk() (*QbdInventoryAssemblyItemAssetAccount, bool)`

GetAssetAccountOk returns a tuple with the AssetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccount

`func (o *QbdInventoryAssemblyItem) SetAssetAccount(v QbdInventoryAssemblyItemAssetAccount)`

SetAssetAccount sets AssetAccount field to given value.


### GetBuildNotificationThreshold

`func (o *QbdInventoryAssemblyItem) GetBuildNotificationThreshold() float32`

GetBuildNotificationThreshold returns the BuildNotificationThreshold field if non-nil, zero value otherwise.

### GetBuildNotificationThresholdOk

`func (o *QbdInventoryAssemblyItem) GetBuildNotificationThresholdOk() (*float32, bool)`

GetBuildNotificationThresholdOk returns a tuple with the BuildNotificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNotificationThreshold

`func (o *QbdInventoryAssemblyItem) SetBuildNotificationThreshold(v float32)`

SetBuildNotificationThreshold sets BuildNotificationThreshold field to given value.


### GetMaximumQuantityOnHand

`func (o *QbdInventoryAssemblyItem) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QbdInventoryAssemblyItem) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QbdInventoryAssemblyItem) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.


### GetQuantityOnHand

`func (o *QbdInventoryAssemblyItem) GetQuantityOnHand() float32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *QbdInventoryAssemblyItem) GetQuantityOnHandOk() (*float32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *QbdInventoryAssemblyItem) SetQuantityOnHand(v float32)`

SetQuantityOnHand sets QuantityOnHand field to given value.


### GetAverageCost

`func (o *QbdInventoryAssemblyItem) GetAverageCost() string`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *QbdInventoryAssemblyItem) GetAverageCostOk() (*string, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *QbdInventoryAssemblyItem) SetAverageCost(v string)`

SetAverageCost sets AverageCost field to given value.


### GetQuantityOnPurchaseOrder

`func (o *QbdInventoryAssemblyItem) GetQuantityOnPurchaseOrder() float32`

GetQuantityOnPurchaseOrder returns the QuantityOnPurchaseOrder field if non-nil, zero value otherwise.

### GetQuantityOnPurchaseOrderOk

`func (o *QbdInventoryAssemblyItem) GetQuantityOnPurchaseOrderOk() (*float32, bool)`

GetQuantityOnPurchaseOrderOk returns a tuple with the QuantityOnPurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnPurchaseOrder

`func (o *QbdInventoryAssemblyItem) SetQuantityOnPurchaseOrder(v float32)`

SetQuantityOnPurchaseOrder sets QuantityOnPurchaseOrder field to given value.


### GetQuantityOnSalesOrder

`func (o *QbdInventoryAssemblyItem) GetQuantityOnSalesOrder() float32`

GetQuantityOnSalesOrder returns the QuantityOnSalesOrder field if non-nil, zero value otherwise.

### GetQuantityOnSalesOrderOk

`func (o *QbdInventoryAssemblyItem) GetQuantityOnSalesOrderOk() (*float32, bool)`

GetQuantityOnSalesOrderOk returns a tuple with the QuantityOnSalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnSalesOrder

`func (o *QbdInventoryAssemblyItem) SetQuantityOnSalesOrder(v float32)`

SetQuantityOnSalesOrder sets QuantityOnSalesOrder field to given value.


### GetExternalId

`func (o *QbdInventoryAssemblyItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdInventoryAssemblyItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdInventoryAssemblyItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLines

`func (o *QbdInventoryAssemblyItem) GetLines() []QbdInventoryAssemblyItemLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdInventoryAssemblyItem) GetLinesOk() (*[]QbdInventoryAssemblyItemLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdInventoryAssemblyItem) SetLines(v []QbdInventoryAssemblyItemLine)`

SetLines sets Lines field to given value.


### GetCustomFields

`func (o *QbdInventoryAssemblyItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdInventoryAssemblyItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdInventoryAssemblyItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


