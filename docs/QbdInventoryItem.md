# QbdInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this inventory item. This ID is unique across all inventory items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_inventory_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this inventory item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this inventory item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this inventory item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this inventory item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two inventory items could both have the &#x60;name&#x60; \&quot;Cabinet\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Kitchen:Cabinet\&quot; and \&quot;Inventory:Cabinet\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this inventory item, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if an inventory item is under \&quot;Products:Electronics\&quot; and has the &#x60;name&#x60; \&quot;Widgets\&quot;, its &#x60;fullName&#x60; would be \&quot;Products:Electronics:Widgets\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all inventory item objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field. | 
**Barcode** | **string** | The inventory item&#39;s barcode. | 
**IsActive** | **bool** | Indicates whether this inventory item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | [**QbdInventoryItemClass**](QbdInventoryItemClass.md) |  | 
**Parent** | [**QbdInventoryItemParent**](QbdInventoryItemParent.md) |  | 
**Sublevel** | **float32** | The depth level of this inventory item in the hierarchy. A top-level inventory item has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, an inventory item with a &#x60;fullName&#x60; of \&quot;Kitchen:Cabinet\&quot; would have a &#x60;sublevel&#x60; of 1. | 
**Sku** | **string** | The inventory item&#39;s stock keeping unit (SKU), which is sometimes the manufacturer&#39;s part number. | 
**UnitOfMeasureSet** | [**QbdInventoryItemUnitOfMeasureSet**](QbdInventoryItemUnitOfMeasureSet.md) |  | 
**SalesTaxCode** | [**QbdInventoryItemSalesTaxCode**](QbdInventoryItemSalesTaxCode.md) |  | 
**SalesDescription** | **string** | The description of this inventory item that appears on sales forms (e.g., invoices, sales receipts) when sold to customers. | 
**SalesPrice** | **string** | The price at which this inventory item is sold to customers, represented as a decimal string. | 
**IncomeAccount** | [**QbdInventoryItemIncomeAccount**](QbdInventoryItemIncomeAccount.md) |  | 
**PurchaseDescription** | **string** | The description of this inventory item that appears on purchase forms (e.g., checks, bills, item receipts) when it is ordered or bought from vendors. | 
**PurchaseCost** | **string** | The cost at which this inventory item is purchased from vendors, represented as a decimal string. | 
**PurchaseTaxCode** | [**QbdInventoryItemPurchaseTaxCode**](QbdInventoryItemPurchaseTaxCode.md) |  | 
**CogsAccount** | [**QbdInventoryItemCogsAccount**](QbdInventoryItemCogsAccount.md) |  | 
**PreferredVendor** | [**QbdInventoryItemPreferredVendor**](QbdInventoryItemPreferredVendor.md) |  | 
**AssetAccount** | [**QbdInventoryItemAssetAccount**](QbdInventoryItemAssetAccount.md) |  | 
**ReorderPoint** | **float32** | The minimum quantity of this inventory item at which QuickBooks prompts for reordering. | 
**MaximumQuantityOnHand** | **float32** | The maximum quantity of this inventory item desired in inventory. | 
**QuantityOnHand** | **float32** | The current quantity of this inventory item available in inventory. To change the &#x60;quantityOnHand&#x60; for an inventory item, you must create an inventory-adjustment instead of updating this inventory item directly. | 
**AverageCost** | **string** | The average cost per unit of this inventory item, represented as a decimal string. | 
**QuantityOnPurchaseOrder** | **float32** | The number of units of this inventory item that have been ordered from vendors (as recorded in purchase orders) but not yet received. | 
**QuantityOnSalesOrder** | **float32** | The number of units of this inventory item that have been sold (as recorded in sales orders) but not yet fulfilled or delivered to customers. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the inventory item object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdInventoryItem

`func NewQbdInventoryItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, barcode string, isActive bool, class QbdInventoryItemClass, parent QbdInventoryItemParent, sublevel float32, sku string, unitOfMeasureSet QbdInventoryItemUnitOfMeasureSet, salesTaxCode QbdInventoryItemSalesTaxCode, salesDescription string, salesPrice string, incomeAccount QbdInventoryItemIncomeAccount, purchaseDescription string, purchaseCost string, purchaseTaxCode QbdInventoryItemPurchaseTaxCode, cogsAccount QbdInventoryItemCogsAccount, preferredVendor QbdInventoryItemPreferredVendor, assetAccount QbdInventoryItemAssetAccount, reorderPoint float32, maximumQuantityOnHand float32, quantityOnHand float32, averageCost string, quantityOnPurchaseOrder float32, quantityOnSalesOrder float32, externalId string, customFields []QbdCustomField, ) *QbdInventoryItem`

NewQbdInventoryItem instantiates a new QbdInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventoryItemWithDefaults

`func NewQbdInventoryItemWithDefaults() *QbdInventoryItem`

NewQbdInventoryItemWithDefaults instantiates a new QbdInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInventoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInventoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInventoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInventoryItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInventoryItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInventoryItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdInventoryItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdInventoryItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdInventoryItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdInventoryItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdInventoryItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdInventoryItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdInventoryItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdInventoryItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdInventoryItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdInventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdInventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdInventoryItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdInventoryItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdInventoryItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdInventoryItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBarcode

`func (o *QbdInventoryItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QbdInventoryItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QbdInventoryItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetIsActive

`func (o *QbdInventoryItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdInventoryItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdInventoryItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdInventoryItem) GetClass() QbdInventoryItemClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdInventoryItem) GetClassOk() (*QbdInventoryItemClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdInventoryItem) SetClass(v QbdInventoryItemClass)`

SetClass sets Class field to given value.


### GetParent

`func (o *QbdInventoryItem) GetParent() QbdInventoryItemParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdInventoryItem) GetParentOk() (*QbdInventoryItemParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdInventoryItem) SetParent(v QbdInventoryItemParent)`

SetParent sets Parent field to given value.


### GetSublevel

`func (o *QbdInventoryItem) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdInventoryItem) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdInventoryItem) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetSku

`func (o *QbdInventoryItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QbdInventoryItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QbdInventoryItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetUnitOfMeasureSet

`func (o *QbdInventoryItem) GetUnitOfMeasureSet() QbdInventoryItemUnitOfMeasureSet`

GetUnitOfMeasureSet returns the UnitOfMeasureSet field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetOk

`func (o *QbdInventoryItem) GetUnitOfMeasureSetOk() (*QbdInventoryItemUnitOfMeasureSet, bool)`

GetUnitOfMeasureSetOk returns a tuple with the UnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSet

`func (o *QbdInventoryItem) SetUnitOfMeasureSet(v QbdInventoryItemUnitOfMeasureSet)`

SetUnitOfMeasureSet sets UnitOfMeasureSet field to given value.


### GetSalesTaxCode

`func (o *QbdInventoryItem) GetSalesTaxCode() QbdInventoryItemSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdInventoryItem) GetSalesTaxCodeOk() (*QbdInventoryItemSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdInventoryItem) SetSalesTaxCode(v QbdInventoryItemSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetSalesDescription

`func (o *QbdInventoryItem) GetSalesDescription() string`

GetSalesDescription returns the SalesDescription field if non-nil, zero value otherwise.

### GetSalesDescriptionOk

`func (o *QbdInventoryItem) GetSalesDescriptionOk() (*string, bool)`

GetSalesDescriptionOk returns a tuple with the SalesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDescription

`func (o *QbdInventoryItem) SetSalesDescription(v string)`

SetSalesDescription sets SalesDescription field to given value.


### GetSalesPrice

`func (o *QbdInventoryItem) GetSalesPrice() string`

GetSalesPrice returns the SalesPrice field if non-nil, zero value otherwise.

### GetSalesPriceOk

`func (o *QbdInventoryItem) GetSalesPriceOk() (*string, bool)`

GetSalesPriceOk returns a tuple with the SalesPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrice

`func (o *QbdInventoryItem) SetSalesPrice(v string)`

SetSalesPrice sets SalesPrice field to given value.


### GetIncomeAccount

`func (o *QbdInventoryItem) GetIncomeAccount() QbdInventoryItemIncomeAccount`

GetIncomeAccount returns the IncomeAccount field if non-nil, zero value otherwise.

### GetIncomeAccountOk

`func (o *QbdInventoryItem) GetIncomeAccountOk() (*QbdInventoryItemIncomeAccount, bool)`

GetIncomeAccountOk returns a tuple with the IncomeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAccount

`func (o *QbdInventoryItem) SetIncomeAccount(v QbdInventoryItemIncomeAccount)`

SetIncomeAccount sets IncomeAccount field to given value.


### GetPurchaseDescription

`func (o *QbdInventoryItem) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *QbdInventoryItem) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *QbdInventoryItem) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.


### GetPurchaseCost

`func (o *QbdInventoryItem) GetPurchaseCost() string`

GetPurchaseCost returns the PurchaseCost field if non-nil, zero value otherwise.

### GetPurchaseCostOk

`func (o *QbdInventoryItem) GetPurchaseCostOk() (*string, bool)`

GetPurchaseCostOk returns a tuple with the PurchaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCost

`func (o *QbdInventoryItem) SetPurchaseCost(v string)`

SetPurchaseCost sets PurchaseCost field to given value.


### GetPurchaseTaxCode

`func (o *QbdInventoryItem) GetPurchaseTaxCode() QbdInventoryItemPurchaseTaxCode`

GetPurchaseTaxCode returns the PurchaseTaxCode field if non-nil, zero value otherwise.

### GetPurchaseTaxCodeOk

`func (o *QbdInventoryItem) GetPurchaseTaxCodeOk() (*QbdInventoryItemPurchaseTaxCode, bool)`

GetPurchaseTaxCodeOk returns a tuple with the PurchaseTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxCode

`func (o *QbdInventoryItem) SetPurchaseTaxCode(v QbdInventoryItemPurchaseTaxCode)`

SetPurchaseTaxCode sets PurchaseTaxCode field to given value.


### GetCogsAccount

`func (o *QbdInventoryItem) GetCogsAccount() QbdInventoryItemCogsAccount`

GetCogsAccount returns the CogsAccount field if non-nil, zero value otherwise.

### GetCogsAccountOk

`func (o *QbdInventoryItem) GetCogsAccountOk() (*QbdInventoryItemCogsAccount, bool)`

GetCogsAccountOk returns a tuple with the CogsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogsAccount

`func (o *QbdInventoryItem) SetCogsAccount(v QbdInventoryItemCogsAccount)`

SetCogsAccount sets CogsAccount field to given value.


### GetPreferredVendor

`func (o *QbdInventoryItem) GetPreferredVendor() QbdInventoryItemPreferredVendor`

GetPreferredVendor returns the PreferredVendor field if non-nil, zero value otherwise.

### GetPreferredVendorOk

`func (o *QbdInventoryItem) GetPreferredVendorOk() (*QbdInventoryItemPreferredVendor, bool)`

GetPreferredVendorOk returns a tuple with the PreferredVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVendor

`func (o *QbdInventoryItem) SetPreferredVendor(v QbdInventoryItemPreferredVendor)`

SetPreferredVendor sets PreferredVendor field to given value.


### GetAssetAccount

`func (o *QbdInventoryItem) GetAssetAccount() QbdInventoryItemAssetAccount`

GetAssetAccount returns the AssetAccount field if non-nil, zero value otherwise.

### GetAssetAccountOk

`func (o *QbdInventoryItem) GetAssetAccountOk() (*QbdInventoryItemAssetAccount, bool)`

GetAssetAccountOk returns a tuple with the AssetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAccount

`func (o *QbdInventoryItem) SetAssetAccount(v QbdInventoryItemAssetAccount)`

SetAssetAccount sets AssetAccount field to given value.


### GetReorderPoint

`func (o *QbdInventoryItem) GetReorderPoint() float32`

GetReorderPoint returns the ReorderPoint field if non-nil, zero value otherwise.

### GetReorderPointOk

`func (o *QbdInventoryItem) GetReorderPointOk() (*float32, bool)`

GetReorderPointOk returns a tuple with the ReorderPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderPoint

`func (o *QbdInventoryItem) SetReorderPoint(v float32)`

SetReorderPoint sets ReorderPoint field to given value.


### GetMaximumQuantityOnHand

`func (o *QbdInventoryItem) GetMaximumQuantityOnHand() float32`

GetMaximumQuantityOnHand returns the MaximumQuantityOnHand field if non-nil, zero value otherwise.

### GetMaximumQuantityOnHandOk

`func (o *QbdInventoryItem) GetMaximumQuantityOnHandOk() (*float32, bool)`

GetMaximumQuantityOnHandOk returns a tuple with the MaximumQuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumQuantityOnHand

`func (o *QbdInventoryItem) SetMaximumQuantityOnHand(v float32)`

SetMaximumQuantityOnHand sets MaximumQuantityOnHand field to given value.


### GetQuantityOnHand

`func (o *QbdInventoryItem) GetQuantityOnHand() float32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *QbdInventoryItem) GetQuantityOnHandOk() (*float32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *QbdInventoryItem) SetQuantityOnHand(v float32)`

SetQuantityOnHand sets QuantityOnHand field to given value.


### GetAverageCost

`func (o *QbdInventoryItem) GetAverageCost() string`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *QbdInventoryItem) GetAverageCostOk() (*string, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *QbdInventoryItem) SetAverageCost(v string)`

SetAverageCost sets AverageCost field to given value.


### GetQuantityOnPurchaseOrder

`func (o *QbdInventoryItem) GetQuantityOnPurchaseOrder() float32`

GetQuantityOnPurchaseOrder returns the QuantityOnPurchaseOrder field if non-nil, zero value otherwise.

### GetQuantityOnPurchaseOrderOk

`func (o *QbdInventoryItem) GetQuantityOnPurchaseOrderOk() (*float32, bool)`

GetQuantityOnPurchaseOrderOk returns a tuple with the QuantityOnPurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnPurchaseOrder

`func (o *QbdInventoryItem) SetQuantityOnPurchaseOrder(v float32)`

SetQuantityOnPurchaseOrder sets QuantityOnPurchaseOrder field to given value.


### GetQuantityOnSalesOrder

`func (o *QbdInventoryItem) GetQuantityOnSalesOrder() float32`

GetQuantityOnSalesOrder returns the QuantityOnSalesOrder field if non-nil, zero value otherwise.

### GetQuantityOnSalesOrderOk

`func (o *QbdInventoryItem) GetQuantityOnSalesOrderOk() (*float32, bool)`

GetQuantityOnSalesOrderOk returns a tuple with the QuantityOnSalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnSalesOrder

`func (o *QbdInventoryItem) SetQuantityOnSalesOrder(v float32)`

SetQuantityOnSalesOrder sets QuantityOnSalesOrder field to given value.


### GetExternalId

`func (o *QbdInventoryItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdInventoryItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdInventoryItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdInventoryItem) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdInventoryItem) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdInventoryItem) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


