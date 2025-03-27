# QuickbooksDesktopTimeTrackingActivitiesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | **string** | The date of this time tracking activity, in ISO 8601 format (YYYY-MM-DD). | 
**EntityId** | **string** | The employee, vendor, or person on QuickBooks&#39;s \&quot;Other Names\&quot; list whose time is being tracked in this time tracking activity. This cannot refer to a customer - use the &#x60;customer&#x60; field to associate a customer or customer-job with this time tracking activity. | 
**CustomerId** | Pointer to **string** | The customer or customer-job to which this time tracking activity could be billed. If &#x60;billingStatus&#x60; is set to \&quot;billable\&quot;, this field is required. | [optional] 
**ServiceItemId** | Pointer to **string** | The type of service performed during this time tracking activity, referring to billable or purchasable services such as specialized labor, consulting hours, and professional fees.  **NOTE**: This field is not required if no &#x60;customer&#x60; is specified. However, if &#x60;billingStatus&#x60; is set to \&quot;billable\&quot;, both this field and &#x60;customer&#x60; are required. | [optional] 
**Duration** | **string** | The time spent performing the service during this time tracking activity, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M.  **NOTE**: Although seconds can be specified when creating a time tracking activity, they are not returned in responses since QuickBooks Desktop&#39;s UI does not display seconds.  **IMPORTANT**: This field is required for updating time tracking activities, even if the field is not being modified, because of a bug in QuickBooks itself. | 
**ClassId** | Pointer to **string** | The time tracking activity&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**PayrollWageItemId** | Pointer to **string** | The payroll wage item (e.g., Regular Pay, Overtime Pay) to use for this time tracking activity. This field can only be used for time tracking if: (1) the person specified in &#x60;entity&#x60; is an employee in QuickBooks, and (2) the \&quot;Use time data to create paychecks\&quot; preference is enabled in their payroll settings. | [optional] 
**Note** | Pointer to **string** | A note or comment about this time tracking activity. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this time tracking activity.  **IMPORTANT**: When this field is set to \&quot;billable\&quot; for time tracking activities, both &#x60;customer&#x60; and &#x60;serviceItem&#x60; are required so that an invoice can be created. | [optional] [default to "billable"]
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopTimeTrackingActivitiesPostRequest

`func NewQuickbooksDesktopTimeTrackingActivitiesPostRequest(transactionDate string, entityId string, duration string, ) *QuickbooksDesktopTimeTrackingActivitiesPostRequest`

NewQuickbooksDesktopTimeTrackingActivitiesPostRequest instantiates a new QuickbooksDesktopTimeTrackingActivitiesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopTimeTrackingActivitiesPostRequestWithDefaults

`func NewQuickbooksDesktopTimeTrackingActivitiesPostRequestWithDefaults() *QuickbooksDesktopTimeTrackingActivitiesPostRequest`

NewQuickbooksDesktopTimeTrackingActivitiesPostRequestWithDefaults instantiates a new QuickbooksDesktopTimeTrackingActivitiesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetEntityId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetServiceItemId() string`

GetServiceItemId returns the ServiceItemId field if non-nil, zero value otherwise.

### GetServiceItemIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetServiceItemIdOk() (*string, bool)`

GetServiceItemIdOk returns a tuple with the ServiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetServiceItemId(v string)`

SetServiceItemId sets ServiceItemId field to given value.

### HasServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasServiceItemId() bool`

HasServiceItemId returns a boolean if a field has been set.

### GetDuration

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetPayrollWageItemId() string`

GetPayrollWageItemId returns the PayrollWageItemId field if non-nil, zero value otherwise.

### GetPayrollWageItemIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetPayrollWageItemIdOk() (*string, bool)`

GetPayrollWageItemIdOk returns a tuple with the PayrollWageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetPayrollWageItemId(v string)`

SetPayrollWageItemId sets PayrollWageItemId field to given value.

### HasPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasPayrollWageItemId() bool`

HasPayrollWageItemId returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopTimeTrackingActivitiesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


