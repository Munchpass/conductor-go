# QbdTimeTrackingActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this time tracking activity. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_time_tracking_activity\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this time tracking activity was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this time tracking activity was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this time tracking activity object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**TransactionDate** | **string** | The date of this time tracking activity, in ISO 8601 format (YYYY-MM-DD). | 
**Entity** | [**QbdTimeTrackingActivityEntity**](QbdTimeTrackingActivityEntity.md) |  | 
**Customer** | [**QbdTimeTrackingActivityCustomer**](QbdTimeTrackingActivityCustomer.md) |  | 
**ServiceItem** | [**QbdTimeTrackingActivityServiceItem**](QbdTimeTrackingActivityServiceItem.md) |  | 
**Duration** | **string** | The time spent performing the service during this time tracking activity, in ISO 8601 format for time intervals (PTnHnMnS). For example, 1 hour and 30 minutes is represented as PT1H30M.  **NOTE**: Although seconds can be specified when creating a time tracking activity, they are not returned in responses since QuickBooks Desktop&#39;s UI does not display seconds.  **IMPORTANT**: This field is required for updating time tracking activities, even if the field is not being modified, because of a bug in QuickBooks itself. | 
**Class** | [**QbdTimeTrackingActivityClass**](QbdTimeTrackingActivityClass.md) |  | 
**PayrollWageItem** | [**QbdTimeTrackingActivityPayrollWageItem**](QbdTimeTrackingActivityPayrollWageItem.md) |  | 
**Note** | **string** | A note or comment about this time tracking activity. | 
**BillingStatus** | **string** | The billing status of this time tracking activity.  **IMPORTANT**: When this field is set to \&quot;billable\&quot; for time tracking activities, both &#x60;customer&#x60; and &#x60;serviceItem&#x60; are required so that an invoice can be created. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**IsBilled** | **bool** | Indicates whether this time tracking activity has been billed. | 

## Methods

### NewQbdTimeTrackingActivity

`func NewQbdTimeTrackingActivity(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, transactionDate string, entity QbdTimeTrackingActivityEntity, customer QbdTimeTrackingActivityCustomer, serviceItem QbdTimeTrackingActivityServiceItem, duration string, class QbdTimeTrackingActivityClass, payrollWageItem QbdTimeTrackingActivityPayrollWageItem, note string, billingStatus string, externalId string, isBilled bool, ) *QbdTimeTrackingActivity`

NewQbdTimeTrackingActivity instantiates a new QbdTimeTrackingActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdTimeTrackingActivityWithDefaults

`func NewQbdTimeTrackingActivityWithDefaults() *QbdTimeTrackingActivity`

NewQbdTimeTrackingActivityWithDefaults instantiates a new QbdTimeTrackingActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdTimeTrackingActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdTimeTrackingActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdTimeTrackingActivity) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdTimeTrackingActivity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdTimeTrackingActivity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdTimeTrackingActivity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdTimeTrackingActivity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdTimeTrackingActivity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdTimeTrackingActivity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdTimeTrackingActivity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdTimeTrackingActivity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdTimeTrackingActivity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdTimeTrackingActivity) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdTimeTrackingActivity) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdTimeTrackingActivity) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetTransactionDate

`func (o *QbdTimeTrackingActivity) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdTimeTrackingActivity) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdTimeTrackingActivity) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetEntity

`func (o *QbdTimeTrackingActivity) GetEntity() QbdTimeTrackingActivityEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *QbdTimeTrackingActivity) GetEntityOk() (*QbdTimeTrackingActivityEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *QbdTimeTrackingActivity) SetEntity(v QbdTimeTrackingActivityEntity)`

SetEntity sets Entity field to given value.


### GetCustomer

`func (o *QbdTimeTrackingActivity) GetCustomer() QbdTimeTrackingActivityCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdTimeTrackingActivity) GetCustomerOk() (*QbdTimeTrackingActivityCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdTimeTrackingActivity) SetCustomer(v QbdTimeTrackingActivityCustomer)`

SetCustomer sets Customer field to given value.


### GetServiceItem

`func (o *QbdTimeTrackingActivity) GetServiceItem() QbdTimeTrackingActivityServiceItem`

GetServiceItem returns the ServiceItem field if non-nil, zero value otherwise.

### GetServiceItemOk

`func (o *QbdTimeTrackingActivity) GetServiceItemOk() (*QbdTimeTrackingActivityServiceItem, bool)`

GetServiceItemOk returns a tuple with the ServiceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItem

`func (o *QbdTimeTrackingActivity) SetServiceItem(v QbdTimeTrackingActivityServiceItem)`

SetServiceItem sets ServiceItem field to given value.


### GetDuration

`func (o *QbdTimeTrackingActivity) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QbdTimeTrackingActivity) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QbdTimeTrackingActivity) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetClass

`func (o *QbdTimeTrackingActivity) GetClass() QbdTimeTrackingActivityClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdTimeTrackingActivity) GetClassOk() (*QbdTimeTrackingActivityClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdTimeTrackingActivity) SetClass(v QbdTimeTrackingActivityClass)`

SetClass sets Class field to given value.


### GetPayrollWageItem

`func (o *QbdTimeTrackingActivity) GetPayrollWageItem() QbdTimeTrackingActivityPayrollWageItem`

GetPayrollWageItem returns the PayrollWageItem field if non-nil, zero value otherwise.

### GetPayrollWageItemOk

`func (o *QbdTimeTrackingActivity) GetPayrollWageItemOk() (*QbdTimeTrackingActivityPayrollWageItem, bool)`

GetPayrollWageItemOk returns a tuple with the PayrollWageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollWageItem

`func (o *QbdTimeTrackingActivity) SetPayrollWageItem(v QbdTimeTrackingActivityPayrollWageItem)`

SetPayrollWageItem sets PayrollWageItem field to given value.


### GetNote

`func (o *QbdTimeTrackingActivity) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdTimeTrackingActivity) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdTimeTrackingActivity) SetNote(v string)`

SetNote sets Note field to given value.


### GetBillingStatus

`func (o *QbdTimeTrackingActivity) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QbdTimeTrackingActivity) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QbdTimeTrackingActivity) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.


### GetExternalId

`func (o *QbdTimeTrackingActivity) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdTimeTrackingActivity) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdTimeTrackingActivity) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetIsBilled

`func (o *QbdTimeTrackingActivity) GetIsBilled() bool`

GetIsBilled returns the IsBilled field if non-nil, zero value otherwise.

### GetIsBilledOk

`func (o *QbdTimeTrackingActivity) GetIsBilledOk() (*bool, bool)`

GetIsBilledOk returns a tuple with the IsBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBilled

`func (o *QbdTimeTrackingActivity) SetIsBilled(v bool)`

SetIsBilled sets IsBilled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


