# TheReportsPreferencesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingReportBasis** | **string** | Determines how the aging periods are calculated in accounts receivable and accounts payable reports for this company file. When set to &#x60;age_from_due_date&#x60;, the overdue days shown in these reports begin with the due date on the invoice. When set to &#x60;age_from_transaction_date&#x60;, the overdue days begin with the date the transaction was created. | 
**SummaryReportBasis** | **string** | Indicates whether summary reports for this company file use cash-basis or accrual-basis bookkeeping. With &#x60;accrual&#x60; basis, transactions are recorded when they occur regardless of when payment is received or made. With &#x60;cash&#x60; basis, transactions are recorded only when payment is received or made. | 

## Methods

### NewTheReportsPreferencesObject

`func NewTheReportsPreferencesObject(agingReportBasis string, summaryReportBasis string, ) *TheReportsPreferencesObject`

NewTheReportsPreferencesObject instantiates a new TheReportsPreferencesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTheReportsPreferencesObjectWithDefaults

`func NewTheReportsPreferencesObjectWithDefaults() *TheReportsPreferencesObject`

NewTheReportsPreferencesObjectWithDefaults instantiates a new TheReportsPreferencesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingReportBasis

`func (o *TheReportsPreferencesObject) GetAgingReportBasis() string`

GetAgingReportBasis returns the AgingReportBasis field if non-nil, zero value otherwise.

### GetAgingReportBasisOk

`func (o *TheReportsPreferencesObject) GetAgingReportBasisOk() (*string, bool)`

GetAgingReportBasisOk returns a tuple with the AgingReportBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingReportBasis

`func (o *TheReportsPreferencesObject) SetAgingReportBasis(v string)`

SetAgingReportBasis sets AgingReportBasis field to given value.


### GetSummaryReportBasis

`func (o *TheReportsPreferencesObject) GetSummaryReportBasis() string`

GetSummaryReportBasis returns the SummaryReportBasis field if non-nil, zero value otherwise.

### GetSummaryReportBasisOk

`func (o *TheReportsPreferencesObject) GetSummaryReportBasisOk() (*string, bool)`

GetSummaryReportBasisOk returns a tuple with the SummaryReportBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryReportBasis

`func (o *TheReportsPreferencesObject) SetSummaryReportBasis(v string)`

SetSummaryReportBasis sets SummaryReportBasis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


