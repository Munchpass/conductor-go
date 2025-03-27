# QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing inventory adjustment line you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new inventory adjustment lines you wish to add. | 
**ItemId** | Pointer to **string** | The inventory item associated with this inventory adjustment line. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the item associated with this inventory adjustment line. This is used for tracking individual units of serialized inventory items. | [optional] 
**LotNumber** | Pointer to **string** | The lot number of the item associated with this inventory adjustment line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**AdjustCount** | Pointer to **float32** | The amount to adjust the count of the inventory item associated with this inventory adjustment line. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for the serial number or lot number of the item associated with this inventory adjustment line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this inventory adjustment line is stored. | [optional] 
**QuantityDifference** | Pointer to **float32** | Either a positive or negative number that shows the change in quantity for the inventory item associated with this inventory adjustment line. A positive number increases the quantity, while a negative number decreases it. | [optional] 
**ValueDifference** | Pointer to **float32** | Either a positive or negative number that shows the change in the total value of the entire stock of the inventory item associated with this inventory adjustment line. A positive number increases the value, while a negative number decreases it. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner

`func NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner(id string, ) *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner`

NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner instantiates a new QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInnerWithDefaults() *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner`

NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetAdjustCount() float32`

GetAdjustCount returns the AdjustCount field if non-nil, zero value otherwise.

### GetAdjustCountOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetAdjustCountOk() (*float32, bool)`

GetAdjustCountOk returns a tuple with the AdjustCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetAdjustCount(v float32)`

SetAdjustCount sets AdjustCount field to given value.

### HasAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasAdjustCount() bool`

HasAdjustCount returns a boolean if a field has been set.

### GetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.

### GetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetQuantityDifference() float32`

GetQuantityDifference returns the QuantityDifference field if non-nil, zero value otherwise.

### GetQuantityDifferenceOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetQuantityDifferenceOk() (*float32, bool)`

GetQuantityDifferenceOk returns a tuple with the QuantityDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetQuantityDifference(v float32)`

SetQuantityDifference sets QuantityDifference field to given value.

### HasQuantityDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasQuantityDifference() bool`

HasQuantityDifference returns a boolean if a field has been set.

### GetValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetValueDifference() float32`

GetValueDifference returns the ValueDifference field if non-nil, zero value otherwise.

### GetValueDifferenceOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) GetValueDifferenceOk() (*float32, bool)`

GetValueDifferenceOk returns a tuple with the ValueDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) SetValueDifference(v float32)`

SetValueDifference sets ValueDifference field to given value.

### HasValueDifference

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner) HasValueDifference() bool`

HasValueDifference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


