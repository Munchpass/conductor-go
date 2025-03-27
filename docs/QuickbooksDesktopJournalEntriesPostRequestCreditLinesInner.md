# QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account to which this journal credit line is being credited. This will increase the balance of this account. | 
**Amount** | Pointer to **string** | The monetary amount of this journal credit line, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this journal credit line. | [optional] 
**EntityId** | Pointer to **string** | The customer, vendor, employee, or other entity associated with this journal credit line.  **IMPORTANT**: If the journal credit line&#39;s &#x60;account&#x60; is an Accounts Receivable (A/R) account, this field must refer to a customer. If the journal credit line&#39;s &#x60;account&#x60; is an Accounts Payable (A/P) account, this field must refer to a vendor. If these requirements are not met, QuickBooks Desktop will not record the transaction. | [optional] 
**ClassId** | Pointer to **string** | The journal credit line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all journal credit lines unless overridden here, at the transaction line level. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this journal credit line&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this journal credit line. | [optional] [default to "billable"]

## Methods

### NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInner

`func NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInner(accountId string, ) *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner`

NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInner instantiates a new QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInnerWithDefaults

`func NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInnerWithDefaults() *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner`

NewQuickbooksDesktopJournalEntriesPostRequestCreditLinesInnerWithDefaults instantiates a new QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetEntityId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopJournalEntriesPostRequestCreditLinesInner) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


