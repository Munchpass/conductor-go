# QuickbooksDesktopStandardTermsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this standard term, unique across all standard terms.  **NOTE**: Standard terms do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this standard term is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**DueDays** | Pointer to **float32** | The number of days until payment is due. | [optional] 
**DiscountDays** | Pointer to **float32** | The number of days within which payment must be received to qualify for the discount specified by &#x60;discountPercentage&#x60;. | [optional] 
**DiscountPercentage** | Pointer to **string** | The discount percentage applied to the payment if received within the number of days specified by &#x60;discountDays&#x60;. The value is between 0 and 100. | [optional] 

## Methods

### NewQuickbooksDesktopStandardTermsPostRequest

`func NewQuickbooksDesktopStandardTermsPostRequest(name string, ) *QuickbooksDesktopStandardTermsPostRequest`

NewQuickbooksDesktopStandardTermsPostRequest instantiates a new QuickbooksDesktopStandardTermsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopStandardTermsPostRequestWithDefaults

`func NewQuickbooksDesktopStandardTermsPostRequestWithDefaults() *QuickbooksDesktopStandardTermsPostRequest`

NewQuickbooksDesktopStandardTermsPostRequestWithDefaults instantiates a new QuickbooksDesktopStandardTermsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopStandardTermsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopStandardTermsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopStandardTermsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDueDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDueDays() float32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDueDaysOk() (*float32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) SetDueDays(v float32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetDiscountDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDiscountDays() float32`

GetDiscountDays returns the DiscountDays field if non-nil, zero value otherwise.

### GetDiscountDaysOk

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDiscountDaysOk() (*float32, bool)`

GetDiscountDaysOk returns a tuple with the DiscountDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) SetDiscountDays(v float32)`

SetDiscountDays sets DiscountDays field to given value.

### HasDiscountDays

`func (o *QuickbooksDesktopStandardTermsPostRequest) HasDiscountDays() bool`

HasDiscountDays returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *QuickbooksDesktopStandardTermsPostRequest) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *QuickbooksDesktopStandardTermsPostRequest) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *QuickbooksDesktopStandardTermsPostRequest) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


