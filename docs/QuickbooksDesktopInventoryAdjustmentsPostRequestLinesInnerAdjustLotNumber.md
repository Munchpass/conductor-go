# QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LotNumber** | Pointer to **string** | The lot number of the item associated with this inventory adjustment line. Used for tracking groups of inventory items that are purchased or manufactured together. | [optional] 
**AdjustCount** | Pointer to **float32** | The amount to adjust the count of the inventory item associated with this inventory adjustment line. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for the serial number or lot number of the item associated with this inventory adjustment line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this inventory adjustment line is stored. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumberWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumberWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumberWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetAdjustCount() float32`

GetAdjustCount returns the AdjustCount field if non-nil, zero value otherwise.

### GetAdjustCountOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetAdjustCountOk() (*float32, bool)`

GetAdjustCountOk returns a tuple with the AdjustCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) SetAdjustCount(v float32)`

SetAdjustCount sets AdjustCount field to given value.

### HasAdjustCount

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) HasAdjustCount() bool`

HasAdjustCount returns a boolean if a field has been set.

### GetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustLotNumber) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


