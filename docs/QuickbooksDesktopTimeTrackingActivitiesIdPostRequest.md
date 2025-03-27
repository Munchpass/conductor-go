# QuickbooksDesktopTimeTrackingActivitiesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the time tracking activity object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | Pointer to **string** | The date of this time tracking activity, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**EntityId** | **string** | The employee, vendor, or person on QuickBooks&#39;s \&quot;Other Names\&quot; list whose time is being tracked in this time tracking activity. This cannot refer to a customer - use the &#x60;customer&#x60; field to associate a customer or customer-job with this time tracking activity.  **IMPORTANT**: This field is required for updating time tracking activities, even if the field is not being modified, because of a bug in QuickBooks itself. | 
**CustomerId** | Pointer to **string** | The customer or customer-job to which this time tracking activity could be billed. If &#x60;billingStatus&#x60; is set to \&quot;billable\&quot;, this field is required. | [optional] 
**ServiceItemId** | Pointer to **string** | The type of service performed during this time tracking activity, referring to billable or purchasable services such as specialized labor, consulting hours, and professional fees.  **NOTE**: This field is not required if no &#x60;customer&#x60; is specified. However, if &#x60;billingStatus&#x60; is set to \&quot;billable\&quot;, both this field and &#x60;customer&#x60; are required. | [optional] 
**Duration** | **string** | The time spent performing the service during this time tracking activity, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M.  **NOTE**: Although seconds can be specified when creating a time tracking activity, they are not returned in responses since QuickBooks Desktop&#39;s UI does not display seconds.  **IMPORTANT**: This field is required for updating time tracking activities, even if the field is not being modified, because of a bug in QuickBooks itself. | 
**ClassId** | Pointer to **string** | The time tracking activity&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**PayrollWageItemId** | Pointer to **string** | The payroll wage item (e.g., Regular Pay, Overtime Pay) to use for this time tracking activity. This field can only be used for time tracking if: (1) the person specified in &#x60;entity&#x60; is an employee in QuickBooks, and (2) the \&quot;Use time data to create paychecks\&quot; preference is enabled in their payroll settings. | [optional] 
**Note** | Pointer to **string** | A note or comment about this time tracking activity. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this time tracking activity.  **IMPORTANT**: When this field is set to \&quot;billable\&quot; for time tracking activities, both &#x60;customer&#x60; and &#x60;serviceItem&#x60; are required so that an invoice can be created. | [optional] 

## Methods

### NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequest

`func NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequest(revisionNumber string, entityId string, duration string, ) *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest`

NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequest instantiates a new QuickbooksDesktopTimeTrackingActivitiesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequestWithDefaults

`func NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequestWithDefaults() *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest`

NewQuickbooksDesktopTimeTrackingActivitiesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopTimeTrackingActivitiesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetEntityId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetServiceItemId() string`

GetServiceItemId returns the ServiceItemId field if non-nil, zero value otherwise.

### GetServiceItemIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetServiceItemIdOk() (*string, bool)`

GetServiceItemIdOk returns a tuple with the ServiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetServiceItemId(v string)`

SetServiceItemId sets ServiceItemId field to given value.

### HasServiceItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasServiceItemId() bool`

HasServiceItemId returns a boolean if a field has been set.

### GetDuration

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetPayrollWageItemId() string`

GetPayrollWageItemId returns the PayrollWageItemId field if non-nil, zero value otherwise.

### GetPayrollWageItemIdOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetPayrollWageItemIdOk() (*string, bool)`

GetPayrollWageItemIdOk returns a tuple with the PayrollWageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetPayrollWageItemId(v string)`

SetPayrollWageItemId sets PayrollWageItemId field to given value.

### HasPayrollWageItemId

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasPayrollWageItemId() bool`

HasPayrollWageItemId returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopTimeTrackingActivitiesIdPostRequest) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


