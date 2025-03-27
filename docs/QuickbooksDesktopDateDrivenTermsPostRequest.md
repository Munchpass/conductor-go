# QuickbooksDesktopDateDrivenTermsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this date-driven term, unique across all date-driven terms.  **NOTE**: Date-driven terms do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this date-driven term is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**DueDayOfMonth** | **float32** | The day of the month when full payment is due without discount. | 
**GracePeriodDays** | Pointer to **float32** | The number of days before &#x60;dueDayOfMonth&#x60; when an invoice or bill issued within this threshold is considered due the following month. For example, with &#x60;dueDayOfMonth&#x60; set to 15 and &#x60;gracePeriodDays&#x60; set to 2, an invoice issued on the 13th would be due on the 15th of the next month, while an invoice issued on the 12th would be due on the 15th of the current month. | [optional] 
**DiscountDayOfMonth** | Pointer to **float32** | The day of the month within which payment must be received to qualify for the discount specified by &#x60;discountPercentage&#x60;. | [optional] 
**DiscountPercentage** | Pointer to **string** | The discount percentage applied to the payment if received on or before the specified &#x60;discountDayOfMonth&#x60;. The value is between 0 and 100. | [optional] 

## Methods

### NewQuickbooksDesktopDateDrivenTermsPostRequest

`func NewQuickbooksDesktopDateDrivenTermsPostRequest(name string, dueDayOfMonth float32, ) *QuickbooksDesktopDateDrivenTermsPostRequest`

NewQuickbooksDesktopDateDrivenTermsPostRequest instantiates a new QuickbooksDesktopDateDrivenTermsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopDateDrivenTermsPostRequestWithDefaults

`func NewQuickbooksDesktopDateDrivenTermsPostRequestWithDefaults() *QuickbooksDesktopDateDrivenTermsPostRequest`

NewQuickbooksDesktopDateDrivenTermsPostRequestWithDefaults instantiates a new QuickbooksDesktopDateDrivenTermsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDueDayOfMonth

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDueDayOfMonth() float32`

GetDueDayOfMonth returns the DueDayOfMonth field if non-nil, zero value otherwise.

### GetDueDayOfMonthOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDueDayOfMonthOk() (*float32, bool)`

GetDueDayOfMonthOk returns a tuple with the DueDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDayOfMonth

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDueDayOfMonth(v float32)`

SetDueDayOfMonth sets DueDayOfMonth field to given value.


### GetGracePeriodDays

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetGracePeriodDays() float32`

GetGracePeriodDays returns the GracePeriodDays field if non-nil, zero value otherwise.

### GetGracePeriodDaysOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetGracePeriodDaysOk() (*float32, bool)`

GetGracePeriodDaysOk returns a tuple with the GracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDays

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetGracePeriodDays(v float32)`

SetGracePeriodDays sets GracePeriodDays field to given value.

### HasGracePeriodDays

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasGracePeriodDays() bool`

HasGracePeriodDays returns a boolean if a field has been set.

### GetDiscountDayOfMonth

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountDayOfMonth() float32`

GetDiscountDayOfMonth returns the DiscountDayOfMonth field if non-nil, zero value otherwise.

### GetDiscountDayOfMonthOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountDayOfMonthOk() (*float32, bool)`

GetDiscountDayOfMonthOk returns a tuple with the DiscountDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDayOfMonth

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDiscountDayOfMonth(v float32)`

SetDiscountDayOfMonth sets DiscountDayOfMonth field to given value.

### HasDiscountDayOfMonth

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasDiscountDayOfMonth() bool`

HasDiscountDayOfMonth returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *QuickbooksDesktopDateDrivenTermsPostRequest) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


