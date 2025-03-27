# QuickbooksDesktopJournalEntriesIdPostRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing journal line you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new journal lines you wish to add. | 
**JournalLineType** | Pointer to **string** | The type of journal line (debit or credit). | [optional] 
**AccountId** | Pointer to **string** | The account to which this journal line is being credited or debited. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this journal line, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this journal line. | [optional] 
**EntityId** | Pointer to **string** | The customer, vendor, employee, or other entity associated with this journal line.  **IMPORTANT**: If the journal line&#39;s &#x60;account&#x60; is an Accounts Receivable (A/R) account, this field must refer to a customer. If the journal line&#39;s &#x60;account&#x60; is an Accounts Payable (A/P) account, this field must refer to a vendor. If these requirements are not met, QuickBooks Desktop will not record the transaction. | [optional] 
**ClassId** | Pointer to **string** | The journal line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all journal lines unless overridden here, at the transaction line level. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this journal line&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this journal line. | [optional] 

## Methods

### NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInner

`func NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInner(id string, ) *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner`

NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInner instantiates a new QuickbooksDesktopJournalEntriesIdPostRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInnerWithDefaults

`func NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInnerWithDefaults() *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner`

NewQuickbooksDesktopJournalEntriesIdPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopJournalEntriesIdPostRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetId(v string)`

SetId sets Id field to given value.


### GetJournalLineType

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetJournalLineType() string`

GetJournalLineType returns the JournalLineType field if non-nil, zero value otherwise.

### GetJournalLineTypeOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetJournalLineTypeOk() (*string, bool)`

GetJournalLineTypeOk returns a tuple with the JournalLineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalLineType

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetJournalLineType(v string)`

SetJournalLineType sets JournalLineType field to given value.

### HasJournalLineType

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasJournalLineType() bool`

HasJournalLineType returns a boolean if a field has been set.

### GetAccountId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetEntityId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopJournalEntriesIdPostRequestLinesInner) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


