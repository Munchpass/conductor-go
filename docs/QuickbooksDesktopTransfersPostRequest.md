# QuickbooksDesktopTransfersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | **string** | The date of this transfer, in ISO 8601 format (YYYY-MM-DD). | 
**SourceAccountId** | **string** | The account from which money will be transferred. | 
**TargetAccountId** | **string** | The account to which money will be transferred. | 
**ClassId** | Pointer to **string** | The transfer&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**Amount** | **string** | The monetary amount of this transfer, represented as a decimal string. | 
**Memo** | Pointer to **string** | A memo or note for this transfer. | [optional] 

## Methods

### NewQuickbooksDesktopTransfersPostRequest

`func NewQuickbooksDesktopTransfersPostRequest(transactionDate string, sourceAccountId string, targetAccountId string, amount string, ) *QuickbooksDesktopTransfersPostRequest`

NewQuickbooksDesktopTransfersPostRequest instantiates a new QuickbooksDesktopTransfersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopTransfersPostRequestWithDefaults

`func NewQuickbooksDesktopTransfersPostRequestWithDefaults() *QuickbooksDesktopTransfersPostRequest`

NewQuickbooksDesktopTransfersPostRequestWithDefaults instantiates a new QuickbooksDesktopTransfersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *QuickbooksDesktopTransfersPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopTransfersPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetSourceAccountId

`func (o *QuickbooksDesktopTransfersPostRequest) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *QuickbooksDesktopTransfersPostRequest) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetTargetAccountId

`func (o *QuickbooksDesktopTransfersPostRequest) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *QuickbooksDesktopTransfersPostRequest) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.


### GetClassId

`func (o *QuickbooksDesktopTransfersPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopTransfersPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopTransfersPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopTransfersPostRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopTransfersPostRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *QuickbooksDesktopTransfersPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopTransfersPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopTransfersPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopTransfersPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


