# QuickbooksDesktopInventoryAdjustmentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account to which this inventory adjustment is posted for tracking inventory value changes. | 
**TransactionDate** | **string** | The date of this inventory adjustment, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this inventory adjustment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this inventory adjustment is stored. | [optional] 
**CustomerId** | Pointer to **string** | The customer or customer-job associated with this inventory adjustment. | [optional] 
**ClassId** | Pointer to **string** | The inventory adjustment&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this inventory adjustment&#39;s line items unless overridden at the line item level. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this inventory adjustment. | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner**](QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner.md) | The inventory adjustment&#39;s item lines, each representing the adjustment of an inventory item&#39;s quantity, value, serial number, or lot number. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsPostRequest

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequest(accountId string, transactionDate string, ) *QuickbooksDesktopInventoryAdjustmentsPostRequest`

NewQuickbooksDesktopInventoryAdjustmentsPostRequest instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsPostRequestWithDefaults() *QuickbooksDesktopInventoryAdjustmentsPostRequest`

NewQuickbooksDesktopInventoryAdjustmentsPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTransactionDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetLines() []QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) GetLinesOk() (*[]QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) SetLines(v []QuickbooksDesktopInventoryAdjustmentsPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInventoryAdjustmentsPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


