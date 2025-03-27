# QuickbooksDesktopBillsPostRequestItemLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The item associated with this item line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this item line is stored. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this item line is stored. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the item associated with this item line. This is used for tracking individual units of serialized inventory items. | [optional] 
**LotNumber** | Pointer to **string** | The lot number of the item associated with this item line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for the serial number or lot number of the item associated with this item line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | [optional] 
**Description** | Pointer to **string** | A description of this item line. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the item associated with this item line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this item line. Must be a valid unit within the item&#39;s available units of measure. | [optional] 
**Cost** | Pointer to **string** | The cost of this item line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;cost&#x60;, QuickBooks will use them to calculate &#x60;cost&#x60;. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this item line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;cost&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;cost&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;cost&#x60;. This field cannot be cleared. | [optional] 
**CustomerId** | Pointer to **string** | The customer or customer-job associated with this item line. | [optional] 
**ClassId** | Pointer to **string** | The item line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all item lines unless overridden here, at the transaction line level. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this item line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this item line. | [optional] [default to "billable"]
**OverrideItemAccountId** | Pointer to **string** | The account to use for this item line, overriding the default account associated with the item. | [optional] 
**LinkToTransactionLine** | Pointer to [**QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine**](QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine.md) |  | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The item line&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**CustomFields** | Pointer to [**[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner**](QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner.md) | The custom fields for the item line object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQuickbooksDesktopBillsPostRequestItemLinesInner

`func NewQuickbooksDesktopBillsPostRequestItemLinesInner() *QuickbooksDesktopBillsPostRequestItemLinesInner`

NewQuickbooksDesktopBillsPostRequestItemLinesInner instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsPostRequestItemLinesInnerWithDefaults

`func NewQuickbooksDesktopBillsPostRequestItemLinesInnerWithDefaults() *QuickbooksDesktopBillsPostRequestItemLinesInner`

NewQuickbooksDesktopBillsPostRequestItemLinesInnerWithDefaults instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetCost

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCustomerId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.

### GetOverrideItemAccountId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetOverrideItemAccountId() string`

GetOverrideItemAccountId returns the OverrideItemAccountId field if non-nil, zero value otherwise.

### GetOverrideItemAccountIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetOverrideItemAccountIdOk() (*string, bool)`

GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideItemAccountId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetOverrideItemAccountId(v string)`

SetOverrideItemAccountId sets OverrideItemAccountId field to given value.

### HasOverrideItemAccountId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasOverrideItemAccountId() bool`

HasOverrideItemAccountId returns a boolean if a field has been set.

### GetLinkToTransactionLine

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLinkToTransactionLine() QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine`

GetLinkToTransactionLine returns the LinkToTransactionLine field if non-nil, zero value otherwise.

### GetLinkToTransactionLineOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLinkToTransactionLineOk() (*QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine, bool)`

GetLinkToTransactionLineOk returns a tuple with the LinkToTransactionLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToTransactionLine

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetLinkToTransactionLine(v QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine)`

SetLinkToTransactionLine sets LinkToTransactionLine field to given value.

### HasLinkToTransactionLine

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasLinkToTransactionLine() bool`

HasLinkToTransactionLine returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomFieldsOk() (*[]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


