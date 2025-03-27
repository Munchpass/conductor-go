# QuickbooksDesktopChecksPostRequestApplyToTransactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The ID of the transaction to be paid by this check. | 
**Amount** | Pointer to **string** | The monetary amount from this check to apply to the specified transaction, represented as a decimal string. | [optional] 

## Methods

### NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInner

`func NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInner(transactionId string, ) *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner`

NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInner instantiates a new QuickbooksDesktopChecksPostRequestApplyToTransactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInnerWithDefaults

`func NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInnerWithDefaults() *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner`

NewQuickbooksDesktopChecksPostRequestApplyToTransactionsInnerWithDefaults instantiates a new QuickbooksDesktopChecksPostRequestApplyToTransactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopChecksPostRequestApplyToTransactionsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


