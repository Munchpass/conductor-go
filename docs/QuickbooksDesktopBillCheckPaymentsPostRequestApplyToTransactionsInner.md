# QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The ID of the receivable transaction to which this payment is applied. | 
**PaymentAmount** | Pointer to **string** | The monetary amount to apply to the receivable transaction, represented as a decimal string. | [optional] 
**ApplyCredits** | Pointer to [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner.md) | Credit memos to apply to this receivable transaction, reducing its balance. This creates a link between this receivable transaction and the specified credit memos.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the receivable transaction and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of receivable transactions, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**DiscountAmount** | Pointer to **string** | The monetary amount by which to reduce the receivable transaction&#39;s receivable amount, represented as a decimal string. | [optional] 
**DiscountAccountId** | Pointer to **string** | The financial account used to track this receivable transaction&#39;s discount. | [optional] 
**DiscountClassId** | Pointer to **string** | The class used to track this receivable transaction&#39;s discount. | [optional] 

## Methods

### NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner

`func NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner(transactionId string, ) *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner instantiates a new QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerWithDefaults

`func NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerWithDefaults() *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner`

NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerWithDefaults instantiates a new QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetPaymentAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetPaymentAmount() string`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetPaymentAmountOk() (*string, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetPaymentAmount(v string)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetApplyCredits

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetApplyCredits() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner`

GetApplyCredits returns the ApplyCredits field if non-nil, zero value otherwise.

### GetApplyCreditsOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetApplyCreditsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner, bool)`

GetApplyCreditsOk returns a tuple with the ApplyCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyCredits

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetApplyCredits(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner)`

SetApplyCredits sets ApplyCredits field to given value.

### HasApplyCredits

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) HasApplyCredits() bool`

HasApplyCredits returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountAmount() string`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountAmountOk() (*string, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetDiscountAmount(v string)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountAccountId() string`

GetDiscountAccountId returns the DiscountAccountId field if non-nil, zero value otherwise.

### GetDiscountAccountIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountAccountIdOk() (*string, bool)`

GetDiscountAccountIdOk returns a tuple with the DiscountAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetDiscountAccountId(v string)`

SetDiscountAccountId sets DiscountAccountId field to given value.

### HasDiscountAccountId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) HasDiscountAccountId() bool`

HasDiscountAccountId returns a boolean if a field has been set.

### GetDiscountClassId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountClassId() string`

GetDiscountClassId returns the DiscountClassId field if non-nil, zero value otherwise.

### GetDiscountClassIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) GetDiscountClassIdOk() (*string, bool)`

GetDiscountClassIdOk returns a tuple with the DiscountClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountClassId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) SetDiscountClassId(v string)`

SetDiscountClassId sets DiscountClassId field to given value.

### HasDiscountClassId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInner) HasDiscountClassId() bool`

HasDiscountClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


