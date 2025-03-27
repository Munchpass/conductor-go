# QbdDateDrivenTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this date-driven term. This ID is unique across all date-driven terms but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_date_driven_term\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this date-driven term was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this date-driven term was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this date-driven term object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this date-driven term, unique across all date-driven terms.  **NOTE**: Date-driven terms do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this date-driven term is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**DueDayOfMonth** | **float32** | The day of the month when full payment is due without discount. | 
**GracePeriodDays** | **NullableFloat32** | The number of days before &#x60;dueDayOfMonth&#x60; when an invoice or bill issued within this threshold is considered due the following month. For example, with &#x60;dueDayOfMonth&#x60; set to 15 and &#x60;gracePeriodDays&#x60; set to 2, an invoice issued on the 13th would be due on the 15th of the next month, while an invoice issued on the 12th would be due on the 15th of the current month. | 
**DiscountDayOfMonth** | **NullableFloat32** | The day of the month within which payment must be received to qualify for the discount specified by &#x60;discountPercentage&#x60;. | 
**DiscountPercentage** | **NullableString** | The discount percentage applied to the payment if received on or before the specified &#x60;discountDayOfMonth&#x60;. The value is between 0 and 100. | 

## Methods

### NewQbdDateDrivenTerm

`func NewQbdDateDrivenTerm(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, dueDayOfMonth float32, gracePeriodDays NullableFloat32, discountDayOfMonth NullableFloat32, discountPercentage NullableString, ) *QbdDateDrivenTerm`

NewQbdDateDrivenTerm instantiates a new QbdDateDrivenTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdDateDrivenTermWithDefaults

`func NewQbdDateDrivenTermWithDefaults() *QbdDateDrivenTerm`

NewQbdDateDrivenTermWithDefaults instantiates a new QbdDateDrivenTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdDateDrivenTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdDateDrivenTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdDateDrivenTerm) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdDateDrivenTerm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdDateDrivenTerm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdDateDrivenTerm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdDateDrivenTerm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdDateDrivenTerm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdDateDrivenTerm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdDateDrivenTerm) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdDateDrivenTerm) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdDateDrivenTerm) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdDateDrivenTerm) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdDateDrivenTerm) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdDateDrivenTerm) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdDateDrivenTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdDateDrivenTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdDateDrivenTerm) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdDateDrivenTerm) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdDateDrivenTerm) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdDateDrivenTerm) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetDueDayOfMonth

`func (o *QbdDateDrivenTerm) GetDueDayOfMonth() float32`

GetDueDayOfMonth returns the DueDayOfMonth field if non-nil, zero value otherwise.

### GetDueDayOfMonthOk

`func (o *QbdDateDrivenTerm) GetDueDayOfMonthOk() (*float32, bool)`

GetDueDayOfMonthOk returns a tuple with the DueDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDayOfMonth

`func (o *QbdDateDrivenTerm) SetDueDayOfMonth(v float32)`

SetDueDayOfMonth sets DueDayOfMonth field to given value.


### GetGracePeriodDays

`func (o *QbdDateDrivenTerm) GetGracePeriodDays() float32`

GetGracePeriodDays returns the GracePeriodDays field if non-nil, zero value otherwise.

### GetGracePeriodDaysOk

`func (o *QbdDateDrivenTerm) GetGracePeriodDaysOk() (*float32, bool)`

GetGracePeriodDaysOk returns a tuple with the GracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDays

`func (o *QbdDateDrivenTerm) SetGracePeriodDays(v float32)`

SetGracePeriodDays sets GracePeriodDays field to given value.


### SetGracePeriodDaysNil

`func (o *QbdDateDrivenTerm) SetGracePeriodDaysNil(b bool)`

 SetGracePeriodDaysNil sets the value for GracePeriodDays to be an explicit nil

### UnsetGracePeriodDays
`func (o *QbdDateDrivenTerm) UnsetGracePeriodDays()`

UnsetGracePeriodDays ensures that no value is present for GracePeriodDays, not even an explicit nil
### GetDiscountDayOfMonth

`func (o *QbdDateDrivenTerm) GetDiscountDayOfMonth() float32`

GetDiscountDayOfMonth returns the DiscountDayOfMonth field if non-nil, zero value otherwise.

### GetDiscountDayOfMonthOk

`func (o *QbdDateDrivenTerm) GetDiscountDayOfMonthOk() (*float32, bool)`

GetDiscountDayOfMonthOk returns a tuple with the DiscountDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDayOfMonth

`func (o *QbdDateDrivenTerm) SetDiscountDayOfMonth(v float32)`

SetDiscountDayOfMonth sets DiscountDayOfMonth field to given value.


### SetDiscountDayOfMonthNil

`func (o *QbdDateDrivenTerm) SetDiscountDayOfMonthNil(b bool)`

 SetDiscountDayOfMonthNil sets the value for DiscountDayOfMonth to be an explicit nil

### UnsetDiscountDayOfMonth
`func (o *QbdDateDrivenTerm) UnsetDiscountDayOfMonth()`

UnsetDiscountDayOfMonth ensures that no value is present for DiscountDayOfMonth, not even an explicit nil
### GetDiscountPercentage

`func (o *QbdDateDrivenTerm) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *QbdDateDrivenTerm) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *QbdDateDrivenTerm) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.


### SetDiscountPercentageNil

`func (o *QbdDateDrivenTerm) SetDiscountPercentageNil(b bool)`

 SetDiscountPercentageNil sets the value for DiscountPercentage to be an explicit nil

### UnsetDiscountPercentage
`func (o *QbdDateDrivenTerm) UnsetDiscountPercentage()`

UnsetDiscountPercentage ensures that no value is present for DiscountPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


