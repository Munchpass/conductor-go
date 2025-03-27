# QbdItemLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this item line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_item_line\&quot;&#x60;. | 
**Item** | [**QbdItemLineItem**](QbdItemLineItem.md) |  | 
**InventorySite** | [**QbdItemLineInventorySite**](QbdItemLineInventorySite.md) |  | 
**InventorySiteLocation** | [**QbdItemLineInventorySiteLocation**](QbdItemLineInventorySiteLocation.md) |  | 
**SerialNumber** | **string** | The serial number of the item associated with this item line. This is used for tracking individual units of serialized inventory items. | 
**LotNumber** | **string** | The lot number of the item associated with this item line. Used for tracking groups of inventory items that are purchased or manufactured together. | 
**ExpirationDate** | **string** | The expiration date for the serial number or lot number of the item associated with this item line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | 
**Description** | **string** | A description of this item line. | 
**Quantity** | **float32** | The quantity of the item associated with this item line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item. | 
**UnitOfMeasure** | **string** | The unit-of-measure used for the &#x60;quantity&#x60; in this item line. Must be a valid unit within the item&#39;s available units of measure. | 
**OverrideUnitOfMeasureSet** | [**QbdItemLineOverrideUnitOfMeasureSet**](QbdItemLineOverrideUnitOfMeasureSet.md) |  | 
**Cost** | **string** | The cost of this item line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;amount&#x60; are specified but not &#x60;cost&#x60;, QuickBooks will use them to calculate &#x60;cost&#x60;. | 
**Amount** | **string** | The monetary amount of this item line, represented as a decimal string. If both &#x60;quantity&#x60; and &#x60;cost&#x60; are specified but not &#x60;amount&#x60;, QuickBooks will use them to calculate &#x60;amount&#x60;. If &#x60;amount&#x60;, &#x60;cost&#x60;, and &#x60;quantity&#x60; are all unspecified, then QuickBooks will calculate &#x60;amount&#x60; based on a &#x60;quantity&#x60; of &#x60;1&#x60; and the suggested &#x60;cost&#x60;. This field cannot be cleared. | 
**Customer** | [**QbdItemLineCustomer**](QbdItemLineCustomer.md) |  | 
**Class** | [**QbdItemLineClass**](QbdItemLineClass.md) |  | 
**SalesTaxCode** | [**QbdItemLineSalesTaxCode**](QbdItemLineSalesTaxCode.md) |  | 
**BillingStatus** | **string** | The billing status of this item line. | 
**SalesRepresentative** | [**QbdItemLineSalesRepresentative**](QbdItemLineSalesRepresentative.md) |  | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the item line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdItemLine

`func NewQbdItemLine(id string, objectType string, item QbdItemLineItem, inventorySite QbdItemLineInventorySite, inventorySiteLocation QbdItemLineInventorySiteLocation, serialNumber string, lotNumber string, expirationDate string, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdItemLineOverrideUnitOfMeasureSet, cost string, amount string, customer QbdItemLineCustomer, class QbdItemLineClass, salesTaxCode QbdItemLineSalesTaxCode, billingStatus string, salesRepresentative QbdItemLineSalesRepresentative, customFields []QbdCustomField, ) *QbdItemLine`

NewQbdItemLine instantiates a new QbdItemLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdItemLineWithDefaults

`func NewQbdItemLineWithDefaults() *QbdItemLine`

NewQbdItemLineWithDefaults instantiates a new QbdItemLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdItemLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdItemLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdItemLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdItemLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdItemLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdItemLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdItemLine) GetItem() QbdItemLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdItemLine) GetItemOk() (*QbdItemLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdItemLine) SetItem(v QbdItemLineItem)`

SetItem sets Item field to given value.


### GetInventorySite

`func (o *QbdItemLine) GetInventorySite() QbdItemLineInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdItemLine) GetInventorySiteOk() (*QbdItemLineInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdItemLine) SetInventorySite(v QbdItemLineInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetInventorySiteLocation

`func (o *QbdItemLine) GetInventorySiteLocation() QbdItemLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdItemLine) GetInventorySiteLocationOk() (*QbdItemLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdItemLine) SetInventorySiteLocation(v QbdItemLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetSerialNumber

`func (o *QbdItemLine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QbdItemLine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QbdItemLine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetLotNumber

`func (o *QbdItemLine) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QbdItemLine) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QbdItemLine) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.


### GetExpirationDate

`func (o *QbdItemLine) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QbdItemLine) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QbdItemLine) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetDescription

`func (o *QbdItemLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdItemLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdItemLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *QbdItemLine) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QbdItemLine) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QbdItemLine) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitOfMeasure

`func (o *QbdItemLine) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QbdItemLine) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QbdItemLine) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetOverrideUnitOfMeasureSet

`func (o *QbdItemLine) GetOverrideUnitOfMeasureSet() QbdItemLineOverrideUnitOfMeasureSet`

GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field if non-nil, zero value otherwise.

### GetOverrideUnitOfMeasureSetOk

`func (o *QbdItemLine) GetOverrideUnitOfMeasureSetOk() (*QbdItemLineOverrideUnitOfMeasureSet, bool)`

GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnitOfMeasureSet

`func (o *QbdItemLine) SetOverrideUnitOfMeasureSet(v QbdItemLineOverrideUnitOfMeasureSet)`

SetOverrideUnitOfMeasureSet sets OverrideUnitOfMeasureSet field to given value.


### GetCost

`func (o *QbdItemLine) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *QbdItemLine) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *QbdItemLine) SetCost(v string)`

SetCost sets Cost field to given value.


### GetAmount

`func (o *QbdItemLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdItemLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdItemLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCustomer

`func (o *QbdItemLine) GetCustomer() QbdItemLineCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdItemLine) GetCustomerOk() (*QbdItemLineCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdItemLine) SetCustomer(v QbdItemLineCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdItemLine) GetClass() QbdItemLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdItemLine) GetClassOk() (*QbdItemLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdItemLine) SetClass(v QbdItemLineClass)`

SetClass sets Class field to given value.


### GetSalesTaxCode

`func (o *QbdItemLine) GetSalesTaxCode() QbdItemLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdItemLine) GetSalesTaxCodeOk() (*QbdItemLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdItemLine) SetSalesTaxCode(v QbdItemLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetBillingStatus

`func (o *QbdItemLine) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QbdItemLine) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QbdItemLine) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.


### GetSalesRepresentative

`func (o *QbdItemLine) GetSalesRepresentative() QbdItemLineSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdItemLine) GetSalesRepresentativeOk() (*QbdItemLineSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdItemLine) SetSalesRepresentative(v QbdItemLineSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetCustomFields

`func (o *QbdItemLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdItemLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdItemLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


