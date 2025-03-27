# QbdAccountingPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsingAccountNumbers** | **bool** | Indicates whether this company file is configured to record an account number for new accounts. If you include an account number when creating a new account while this preference is &#x60;false&#x60;, the account number will still be set, but will not be visible in the QuickBooks user interface. | 
**IsRequiringAccounts** | **bool** | Indicates whether this company file is configured to require an account for new transactions. If &#x60;true&#x60;, a transaction cannot be recorded in the QuickBooks user interface unless it is assigned to an account. (However, transactions affected by this preference always require an account to be specified when added through the API.) | 
**IsUsingClassTracking** | **bool** | Indicates whether this company file is configured to use the &#x60;class&#x60; field on all transactions. | 
**DefaultTransactionClass** | **string** | The default class assigned to transactions for this company file. | 
**IsUsingAuditTrail** | **bool** | Indicates whether this company file is configured to log all transaction changes in the audit trail report. If &#x60;false&#x60;, QuickBooks logs only the most recent version of each transaction. | 
**IsAssigningJournalEntryNumbers** | **bool** | Indicates whether this company file is configured to automatically assign a number to each journal entry. | 
**ClosingDate** | **string** | The company closing date set within this company file. (The QuickBooks Admin can assign a password restricting access to transactions that occurred before this date.) | 

## Methods

### NewQbdAccountingPreferences

`func NewQbdAccountingPreferences(isUsingAccountNumbers bool, isRequiringAccounts bool, isUsingClassTracking bool, defaultTransactionClass string, isUsingAuditTrail bool, isAssigningJournalEntryNumbers bool, closingDate string, ) *QbdAccountingPreferences`

NewQbdAccountingPreferences instantiates a new QbdAccountingPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdAccountingPreferencesWithDefaults

`func NewQbdAccountingPreferencesWithDefaults() *QbdAccountingPreferences`

NewQbdAccountingPreferencesWithDefaults instantiates a new QbdAccountingPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsingAccountNumbers

`func (o *QbdAccountingPreferences) GetIsUsingAccountNumbers() bool`

GetIsUsingAccountNumbers returns the IsUsingAccountNumbers field if non-nil, zero value otherwise.

### GetIsUsingAccountNumbersOk

`func (o *QbdAccountingPreferences) GetIsUsingAccountNumbersOk() (*bool, bool)`

GetIsUsingAccountNumbersOk returns a tuple with the IsUsingAccountNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingAccountNumbers

`func (o *QbdAccountingPreferences) SetIsUsingAccountNumbers(v bool)`

SetIsUsingAccountNumbers sets IsUsingAccountNumbers field to given value.


### GetIsRequiringAccounts

`func (o *QbdAccountingPreferences) GetIsRequiringAccounts() bool`

GetIsRequiringAccounts returns the IsRequiringAccounts field if non-nil, zero value otherwise.

### GetIsRequiringAccountsOk

`func (o *QbdAccountingPreferences) GetIsRequiringAccountsOk() (*bool, bool)`

GetIsRequiringAccountsOk returns a tuple with the IsRequiringAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequiringAccounts

`func (o *QbdAccountingPreferences) SetIsRequiringAccounts(v bool)`

SetIsRequiringAccounts sets IsRequiringAccounts field to given value.


### GetIsUsingClassTracking

`func (o *QbdAccountingPreferences) GetIsUsingClassTracking() bool`

GetIsUsingClassTracking returns the IsUsingClassTracking field if non-nil, zero value otherwise.

### GetIsUsingClassTrackingOk

`func (o *QbdAccountingPreferences) GetIsUsingClassTrackingOk() (*bool, bool)`

GetIsUsingClassTrackingOk returns a tuple with the IsUsingClassTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingClassTracking

`func (o *QbdAccountingPreferences) SetIsUsingClassTracking(v bool)`

SetIsUsingClassTracking sets IsUsingClassTracking field to given value.


### GetDefaultTransactionClass

`func (o *QbdAccountingPreferences) GetDefaultTransactionClass() string`

GetDefaultTransactionClass returns the DefaultTransactionClass field if non-nil, zero value otherwise.

### GetDefaultTransactionClassOk

`func (o *QbdAccountingPreferences) GetDefaultTransactionClassOk() (*string, bool)`

GetDefaultTransactionClassOk returns a tuple with the DefaultTransactionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTransactionClass

`func (o *QbdAccountingPreferences) SetDefaultTransactionClass(v string)`

SetDefaultTransactionClass sets DefaultTransactionClass field to given value.


### GetIsUsingAuditTrail

`func (o *QbdAccountingPreferences) GetIsUsingAuditTrail() bool`

GetIsUsingAuditTrail returns the IsUsingAuditTrail field if non-nil, zero value otherwise.

### GetIsUsingAuditTrailOk

`func (o *QbdAccountingPreferences) GetIsUsingAuditTrailOk() (*bool, bool)`

GetIsUsingAuditTrailOk returns a tuple with the IsUsingAuditTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingAuditTrail

`func (o *QbdAccountingPreferences) SetIsUsingAuditTrail(v bool)`

SetIsUsingAuditTrail sets IsUsingAuditTrail field to given value.


### GetIsAssigningJournalEntryNumbers

`func (o *QbdAccountingPreferences) GetIsAssigningJournalEntryNumbers() bool`

GetIsAssigningJournalEntryNumbers returns the IsAssigningJournalEntryNumbers field if non-nil, zero value otherwise.

### GetIsAssigningJournalEntryNumbersOk

`func (o *QbdAccountingPreferences) GetIsAssigningJournalEntryNumbersOk() (*bool, bool)`

GetIsAssigningJournalEntryNumbersOk returns a tuple with the IsAssigningJournalEntryNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigningJournalEntryNumbers

`func (o *QbdAccountingPreferences) SetIsAssigningJournalEntryNumbers(v bool)`

SetIsAssigningJournalEntryNumbers sets IsAssigningJournalEntryNumbers field to given value.


### GetClosingDate

`func (o *QbdAccountingPreferences) GetClosingDate() string`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *QbdAccountingPreferences) GetClosingDateOk() (*string, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *QbdAccountingPreferences) SetClosingDate(v string)`

SetClosingDate sets ClosingDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


