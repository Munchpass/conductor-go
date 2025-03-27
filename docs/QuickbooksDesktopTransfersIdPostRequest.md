# QuickbooksDesktopTransfersIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the transfer object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | Pointer to **string** | The date of this transfer, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SourceAccountId** | Pointer to **string** | The account from which money will be transferred. | [optional] 
**TargetAccountId** | Pointer to **string** | The account to which money will be transferred. | [optional] 
**ClassId** | Pointer to **string** | The transfer&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this transfer, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this transfer. | [optional] 

## Methods

### NewQuickbooksDesktopTransfersIdPostRequest

`func NewQuickbooksDesktopTransfersIdPostRequest(revisionNumber string, ) *QuickbooksDesktopTransfersIdPostRequest`

NewQuickbooksDesktopTransfersIdPostRequest instantiates a new QuickbooksDesktopTransfersIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopTransfersIdPostRequestWithDefaults

`func NewQuickbooksDesktopTransfersIdPostRequestWithDefaults() *QuickbooksDesktopTransfersIdPostRequest`

NewQuickbooksDesktopTransfersIdPostRequestWithDefaults instantiates a new QuickbooksDesktopTransfersIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetTargetAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.

### HasTargetAccountId

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasTargetAccountId() bool`

HasTargetAccountId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopTransfersIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopTransfersIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopTransfersIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


