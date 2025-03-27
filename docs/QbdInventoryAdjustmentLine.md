# QbdInventoryAdjustmentLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this inventory adjustment line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_inventory_adjustment_line\&quot;&#x60;. | 
**Item** | [**QbdInventoryAdjustmentLineItem**](QbdInventoryAdjustmentLineItem.md) |  | 
**SerialNumber** | **string** | The serial number of the item associated with this inventory adjustment line. This is used for tracking individual units of serialized inventory items. | 
**SerialNumberAction** | **string** | Indicates whether this inventory adjustment line&#39;s serial number was added or removed from the inventory. | 
**LotNumber** | **string** | The lot number of the item associated with this inventory adjustment line. Used for tracking groups of inventory items that are purchased or manufactured together. | 
**ExpirationDate** | **string** | The expiration date for the serial number or lot number of the item associated with this inventory adjustment line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | 
**InventorySiteLocation** | [**QbdInventoryAdjustmentLineInventorySiteLocation**](QbdInventoryAdjustmentLineInventorySiteLocation.md) |  | 
**QuantityDifference** | **float32** | Either a positive or negative number that shows the change in quantity for the inventory item associated with this inventory adjustment line. A positive number increases the quantity, while a negative number decreases it. | 
**ValueDifference** | **float32** | Either a positive or negative number that shows the change in the total value of the entire stock of the inventory item associated with this inventory adjustment line. A positive number increases the value, while a negative number decreases it. | 

## Methods

### NewQbdInventoryAdjustmentLine

`func NewQbdInventoryAdjustmentLine(id string, objectType string, item QbdInventoryAdjustmentLineItem, serialNumber string, serialNumberAction string, lotNumber string, expirationDate string, inventorySiteLocation QbdInventoryAdjustmentLineInventorySiteLocation, quantityDifference float32, valueDifference float32, ) *QbdInventoryAdjustmentLine`

NewQbdInventoryAdjustmentLine instantiates a new QbdInventoryAdjustmentLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInventoryAdjustmentLineWithDefaults

`func NewQbdInventoryAdjustmentLineWithDefaults() *QbdInventoryAdjustmentLine`

NewQbdInventoryAdjustmentLineWithDefaults instantiates a new QbdInventoryAdjustmentLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInventoryAdjustmentLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInventoryAdjustmentLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInventoryAdjustmentLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInventoryAdjustmentLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInventoryAdjustmentLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInventoryAdjustmentLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItem

`func (o *QbdInventoryAdjustmentLine) GetItem() QbdInventoryAdjustmentLineItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *QbdInventoryAdjustmentLine) GetItemOk() (*QbdInventoryAdjustmentLineItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *QbdInventoryAdjustmentLine) SetItem(v QbdInventoryAdjustmentLineItem)`

SetItem sets Item field to given value.


### GetSerialNumber

`func (o *QbdInventoryAdjustmentLine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QbdInventoryAdjustmentLine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QbdInventoryAdjustmentLine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetSerialNumberAction

`func (o *QbdInventoryAdjustmentLine) GetSerialNumberAction() string`

GetSerialNumberAction returns the SerialNumberAction field if non-nil, zero value otherwise.

### GetSerialNumberActionOk

`func (o *QbdInventoryAdjustmentLine) GetSerialNumberActionOk() (*string, bool)`

GetSerialNumberActionOk returns a tuple with the SerialNumberAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberAction

`func (o *QbdInventoryAdjustmentLine) SetSerialNumberAction(v string)`

SetSerialNumberAction sets SerialNumberAction field to given value.


### GetLotNumber

`func (o *QbdInventoryAdjustmentLine) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QbdInventoryAdjustmentLine) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QbdInventoryAdjustmentLine) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.


### GetExpirationDate

`func (o *QbdInventoryAdjustmentLine) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QbdInventoryAdjustmentLine) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QbdInventoryAdjustmentLine) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetInventorySiteLocation

`func (o *QbdInventoryAdjustmentLine) GetInventorySiteLocation() QbdInventoryAdjustmentLineInventorySiteLocation`

GetInventorySiteLocation returns the InventorySiteLocation field if non-nil, zero value otherwise.

### GetInventorySiteLocationOk

`func (o *QbdInventoryAdjustmentLine) GetInventorySiteLocationOk() (*QbdInventoryAdjustmentLineInventorySiteLocation, bool)`

GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocation

`func (o *QbdInventoryAdjustmentLine) SetInventorySiteLocation(v QbdInventoryAdjustmentLineInventorySiteLocation)`

SetInventorySiteLocation sets InventorySiteLocation field to given value.


### GetQuantityDifference

`func (o *QbdInventoryAdjustmentLine) GetQuantityDifference() float32`

GetQuantityDifference returns the QuantityDifference field if non-nil, zero value otherwise.

### GetQuantityDifferenceOk

`func (o *QbdInventoryAdjustmentLine) GetQuantityDifferenceOk() (*float32, bool)`

GetQuantityDifferenceOk returns a tuple with the QuantityDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDifference

`func (o *QbdInventoryAdjustmentLine) SetQuantityDifference(v float32)`

SetQuantityDifference sets QuantityDifference field to given value.


### GetValueDifference

`func (o *QbdInventoryAdjustmentLine) GetValueDifference() float32`

GetValueDifference returns the ValueDifference field if non-nil, zero value otherwise.

### GetValueDifferenceOk

`func (o *QbdInventoryAdjustmentLine) GetValueDifferenceOk() (*float32, bool)`

GetValueDifferenceOk returns a tuple with the ValueDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDifference

`func (o *QbdInventoryAdjustmentLine) SetValueDifference(v float32)`

SetValueDifference sets ValueDifference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


