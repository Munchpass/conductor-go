# QuickbooksDesktopInventoryAdjustmentsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the inventory adjustment object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**AccountId** | Pointer to **string** | The account to which this inventory adjustment is posted for tracking inventory value changes. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this inventory adjustment is stored. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this inventory adjustment, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this inventory adjustment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**CustomerId** | Pointer to **string** | The customer or customer-job associated with this inventory adjustment. | [optional] 
**ClassId** | Pointer to **string** | The inventory adjustment&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this inventory adjustment&#39;s line items unless overridden at the line item level. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this inventory adjustment. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner**](QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner.md) | The inventory adjustment&#39;s item lines, each representing the adjustment of an inventory item&#39;s quantity, value, serial number, or lot number.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing item lines for the inventory adjustment with this array. To keep any existing item lines, you must include them in this array even if they have not changed. **Any item lines not included will be removed.**  2. To add a new item line, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any item lines, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopInventoryAdjustmentsIdPostRequest

`func NewQuickbooksDesktopInventoryAdjustmentsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopInventoryAdjustmentsIdPostRequest`

NewQuickbooksDesktopInventoryAdjustmentsIdPostRequest instantiates a new QuickbooksDesktopInventoryAdjustmentsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestWithDefaults

`func NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestWithDefaults() *QuickbooksDesktopInventoryAdjustmentsIdPostRequest`

NewQuickbooksDesktopInventoryAdjustmentsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopInventoryAdjustmentsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetAccountId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetLines() []QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) SetLines(v []QuickbooksDesktopInventoryAdjustmentsIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInventoryAdjustmentsIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


