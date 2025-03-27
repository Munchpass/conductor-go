# QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewQuantity** | Pointer to **float32** | The new quantity for the inventory item associated with this inventory adjustment line. | [optional] 
**QuantityDifference** | Pointer to **float32** | Either a positive or negative number that shows the change in quantity for the inventory item associated with this inventory adjustment line. A positive number increases the quantity, while a negative number decreases it. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the item associated with this inventory adjustment line. This is used for tracking individual units of serialized inventory items. | [optional] 
**LotNumber** | Pointer to **string** | The lot number of the item associated with this inventory adjustment line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for the serial number or lot number of the item associated with this inventory adjustment line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this inventory adjustment line is stored. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantityWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantityWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantityWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetNewQuantity() float32`

GetNewQuantity returns the NewQuantity field if non-nil, zero value otherwise.

### GetNewQuantityOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetNewQuantityOk() (*float32, bool)`

GetNewQuantityOk returns a tuple with the NewQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetNewQuantity(v float32)`

SetNewQuantity sets NewQuantity field to given value.

### HasNewQuantity

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasNewQuantity() bool`

HasNewQuantity returns a boolean if a field has been set.

### GetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetQuantityDifference() float32`

GetQuantityDifference returns the QuantityDifference field if non-nil, zero value otherwise.

### GetQuantityDifferenceOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetQuantityDifferenceOk() (*float32, bool)`

GetQuantityDifferenceOk returns a tuple with the QuantityDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetQuantityDifference(v float32)`

SetQuantityDifference sets QuantityDifference field to given value.

### HasQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasQuantityDifference() bool`

HasQuantityDifference returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustQuantity) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


