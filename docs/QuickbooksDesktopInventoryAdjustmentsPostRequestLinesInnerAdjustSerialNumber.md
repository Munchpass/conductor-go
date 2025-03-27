# QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddSerialNumber** | Pointer to **string** | The serial number, which represents a unique unit of the inventory item associated with this inventory adjustment line, to add to inventory. | [optional] 
**RemoveSerialNumber** | Pointer to **string** | The serial number, which represents a unique unit of the inventory item associated with this inventory adjustment line, to remove from inventory. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for the serial number or lot number of the item associated with this inventory adjustment line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later. | [optional] 
**InventorySiteLocationId** | Pointer to **string** | The specific location (e.g., bin or shelf) within the inventory site where the item associated with this inventory adjustment line is stored. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumberWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumberWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumberWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetAddSerialNumber() string`

GetAddSerialNumber returns the AddSerialNumber field if non-nil, zero value otherwise.

### GetAddSerialNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetAddSerialNumberOk() (*string, bool)`

GetAddSerialNumberOk returns a tuple with the AddSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) SetAddSerialNumber(v string)`

SetAddSerialNumber sets AddSerialNumber field to given value.

### HasAddSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) HasAddSerialNumber() bool`

HasAddSerialNumber returns a boolean if a field has been set.

### GetRemoveSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetRemoveSerialNumber() string`

GetRemoveSerialNumber returns the RemoveSerialNumber field if non-nil, zero value otherwise.

### GetRemoveSerialNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetRemoveSerialNumberOk() (*string, bool)`

GetRemoveSerialNumberOk returns a tuple with the RemoveSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) SetRemoveSerialNumber(v string)`

SetRemoveSerialNumber sets RemoveSerialNumber field to given value.

### HasRemoveSerialNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) HasRemoveSerialNumber() bool`

HasRemoveSerialNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetInventorySiteLocationId() string`

GetInventorySiteLocationId returns the InventorySiteLocationId field if non-nil, zero value otherwise.

### GetInventorySiteLocationIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) GetInventorySiteLocationIdOk() (*string, bool)`

GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) SetInventorySiteLocationId(v string)`

SetInventorySiteLocationId sets InventorySiteLocationId field to given value.

### HasInventorySiteLocationId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInnerAdjustSerialNumber) HasInventorySiteLocationId() bool`

HasInventorySiteLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


