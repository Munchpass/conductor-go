# QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditTransactionId** | **string** | The unique identifier of the credit transaction (credit memo or vendor credit) to apply to this transaction. | 
**AppliedAmount** | **string** | The amount of credit applied to this transaction. This could include customer deposits, payments, or credits. Represented as a decimal string. | 
**OverrideCreditApplication** | Pointer to **bool** | Indicates whether to override the credit. | [optional] [default to false]

## Methods

### NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner

`func NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner(creditTransactionId string, appliedAmount string, ) *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner`

NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner instantiates a new QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInnerWithDefaults

`func NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInnerWithDefaults() *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner`

NewQuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInnerWithDefaults instantiates a new QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditTransactionId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetCreditTransactionId() string`

GetCreditTransactionId returns the CreditTransactionId field if non-nil, zero value otherwise.

### GetCreditTransactionIdOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetCreditTransactionIdOk() (*string, bool)`

GetCreditTransactionIdOk returns a tuple with the CreditTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTransactionId

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) SetCreditTransactionId(v string)`

SetCreditTransactionId sets CreditTransactionId field to given value.


### GetAppliedAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetAppliedAmount() string`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetAppliedAmountOk() (*string, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) SetAppliedAmount(v string)`

SetAppliedAmount sets AppliedAmount field to given value.


### GetOverrideCreditApplication

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetOverrideCreditApplication() bool`

GetOverrideCreditApplication returns the OverrideCreditApplication field if non-nil, zero value otherwise.

### GetOverrideCreditApplicationOk

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) GetOverrideCreditApplicationOk() (*bool, bool)`

GetOverrideCreditApplicationOk returns a tuple with the OverrideCreditApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCreditApplication

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) SetOverrideCreditApplication(v bool)`

SetOverrideCreditApplication sets OverrideCreditApplication field to given value.

### HasOverrideCreditApplication

`func (o *QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner) HasOverrideCreditApplication() bool`

HasOverrideCreditApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


