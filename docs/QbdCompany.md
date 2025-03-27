# QbdCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSampleCompanyFile** | **bool** | Indicates whether the connected QuickBooks company file is a \&quot;sample file\&quot;, which is a mock company file used for testing. | 
**CompanyName** | **NullableString** | The name of the QuickBooks user&#39;s business associated with this company. This name is used on invoices, checks, and other forms, while &#x60;legalCompanyName&#x60; is used on tax forms and pay stubs. | 
**LegalCompanyName** | **NullableString** | The legal name of this company&#39;s business, as specified in QuickBooks. This value is used on tax forms and pay stubs, while &#x60;companyName&#x60; is used on invoices, checks, and other forms. | 
**Address** | [**NullableQbdAddress**](QbdAddress.md) |  | 
**LegalAddress** | [**NullableQbdAddress**](QbdAddress.md) |  | 
**AddressForCustomer** | [**NullableQbdAddress**](QbdAddress.md) |  | 
**Phone** | **NullableString** | The company&#39;s primary telephone number. | 
**Fax** | **NullableString** | The company&#39;s fax number. | 
**Email** | **NullableString** | The company&#39;s email address. | 
**Website** | **NullableString** | The company&#39;s public website. | 
**FiscalYearStartMonth** | **NullableString** | The first month of this company&#39;s fiscal year, which determines the default date range for financial reports. | 
**IncomeTaxYearStartMonth** | **NullableString** | The first month of this company&#39;s income tax year, which determines the default date range for financial reports. | 
**CompanyType** | **NullableString** | The company type, which the QuickBooks user selected from a list when creating the company file. | 
**Ein** | **NullableString** | The company&#39;s Employer Identification Number. | 
**Ssn** | **NullableString** | The company&#39;s Social Security Number. The value can be with or without dashes.  **NOTE**: This field cannot be changed after the company is created. | 
**TaxForm** | **NullableString** | The tax form that the QuickBooks user expects to file for this company&#39;s taxes. When a specific tax form is selected (any value other than &#x60;other_or_none&#x60;), QuickBooks allows associating each account with a specific tax form line. This association appears in account query responses. | 
**SubscribedServices** | [**NullableQbdSubscribedService**](QbdSubscribedService.md) |  | 
**AccountantCopy** | [**NullableQbdAccountantCopy**](QbdAccountantCopy.md) |  | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the company object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdCompany

`func NewQbdCompany(isSampleCompanyFile bool, companyName NullableString, legalCompanyName NullableString, address NullableQbdAddress, legalAddress NullableQbdAddress, addressForCustomer NullableQbdAddress, phone NullableString, fax NullableString, email NullableString, website NullableString, fiscalYearStartMonth NullableString, incomeTaxYearStartMonth NullableString, companyType NullableString, ein NullableString, ssn NullableString, taxForm NullableString, subscribedServices NullableQbdSubscribedService, accountantCopy NullableQbdAccountantCopy, customFields []QbdCustomField, ) *QbdCompany`

NewQbdCompany instantiates a new QbdCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCompanyWithDefaults

`func NewQbdCompanyWithDefaults() *QbdCompany`

NewQbdCompanyWithDefaults instantiates a new QbdCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSampleCompanyFile

`func (o *QbdCompany) GetIsSampleCompanyFile() bool`

GetIsSampleCompanyFile returns the IsSampleCompanyFile field if non-nil, zero value otherwise.

### GetIsSampleCompanyFileOk

`func (o *QbdCompany) GetIsSampleCompanyFileOk() (*bool, bool)`

GetIsSampleCompanyFileOk returns a tuple with the IsSampleCompanyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSampleCompanyFile

`func (o *QbdCompany) SetIsSampleCompanyFile(v bool)`

SetIsSampleCompanyFile sets IsSampleCompanyFile field to given value.


### GetCompanyName

`func (o *QbdCompany) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QbdCompany) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QbdCompany) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### SetCompanyNameNil

`func (o *QbdCompany) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *QbdCompany) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetLegalCompanyName

`func (o *QbdCompany) GetLegalCompanyName() string`

GetLegalCompanyName returns the LegalCompanyName field if non-nil, zero value otherwise.

### GetLegalCompanyNameOk

`func (o *QbdCompany) GetLegalCompanyNameOk() (*string, bool)`

GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalCompanyName

`func (o *QbdCompany) SetLegalCompanyName(v string)`

SetLegalCompanyName sets LegalCompanyName field to given value.


### SetLegalCompanyNameNil

`func (o *QbdCompany) SetLegalCompanyNameNil(b bool)`

 SetLegalCompanyNameNil sets the value for LegalCompanyName to be an explicit nil

### UnsetLegalCompanyName
`func (o *QbdCompany) UnsetLegalCompanyName()`

UnsetLegalCompanyName ensures that no value is present for LegalCompanyName, not even an explicit nil
### GetAddress

`func (o *QbdCompany) GetAddress() QbdAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdCompany) GetAddressOk() (*QbdAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdCompany) SetAddress(v QbdAddress)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *QbdCompany) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *QbdCompany) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetLegalAddress

`func (o *QbdCompany) GetLegalAddress() QbdAddress`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *QbdCompany) GetLegalAddressOk() (*QbdAddress, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *QbdCompany) SetLegalAddress(v QbdAddress)`

SetLegalAddress sets LegalAddress field to given value.


### SetLegalAddressNil

`func (o *QbdCompany) SetLegalAddressNil(b bool)`

 SetLegalAddressNil sets the value for LegalAddress to be an explicit nil

### UnsetLegalAddress
`func (o *QbdCompany) UnsetLegalAddress()`

UnsetLegalAddress ensures that no value is present for LegalAddress, not even an explicit nil
### GetAddressForCustomer

`func (o *QbdCompany) GetAddressForCustomer() QbdAddress`

GetAddressForCustomer returns the AddressForCustomer field if non-nil, zero value otherwise.

### GetAddressForCustomerOk

`func (o *QbdCompany) GetAddressForCustomerOk() (*QbdAddress, bool)`

GetAddressForCustomerOk returns a tuple with the AddressForCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressForCustomer

`func (o *QbdCompany) SetAddressForCustomer(v QbdAddress)`

SetAddressForCustomer sets AddressForCustomer field to given value.


### SetAddressForCustomerNil

`func (o *QbdCompany) SetAddressForCustomerNil(b bool)`

 SetAddressForCustomerNil sets the value for AddressForCustomer to be an explicit nil

### UnsetAddressForCustomer
`func (o *QbdCompany) UnsetAddressForCustomer()`

UnsetAddressForCustomer ensures that no value is present for AddressForCustomer, not even an explicit nil
### GetPhone

`func (o *QbdCompany) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QbdCompany) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QbdCompany) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *QbdCompany) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *QbdCompany) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *QbdCompany) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QbdCompany) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QbdCompany) SetFax(v string)`

SetFax sets Fax field to given value.


### SetFaxNil

`func (o *QbdCompany) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *QbdCompany) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetEmail

`func (o *QbdCompany) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QbdCompany) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QbdCompany) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *QbdCompany) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *QbdCompany) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetWebsite

`func (o *QbdCompany) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *QbdCompany) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *QbdCompany) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *QbdCompany) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *QbdCompany) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetFiscalYearStartMonth

`func (o *QbdCompany) GetFiscalYearStartMonth() string`

GetFiscalYearStartMonth returns the FiscalYearStartMonth field if non-nil, zero value otherwise.

### GetFiscalYearStartMonthOk

`func (o *QbdCompany) GetFiscalYearStartMonthOk() (*string, bool)`

GetFiscalYearStartMonthOk returns a tuple with the FiscalYearStartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStartMonth

`func (o *QbdCompany) SetFiscalYearStartMonth(v string)`

SetFiscalYearStartMonth sets FiscalYearStartMonth field to given value.


### SetFiscalYearStartMonthNil

`func (o *QbdCompany) SetFiscalYearStartMonthNil(b bool)`

 SetFiscalYearStartMonthNil sets the value for FiscalYearStartMonth to be an explicit nil

### UnsetFiscalYearStartMonth
`func (o *QbdCompany) UnsetFiscalYearStartMonth()`

UnsetFiscalYearStartMonth ensures that no value is present for FiscalYearStartMonth, not even an explicit nil
### GetIncomeTaxYearStartMonth

`func (o *QbdCompany) GetIncomeTaxYearStartMonth() string`

GetIncomeTaxYearStartMonth returns the IncomeTaxYearStartMonth field if non-nil, zero value otherwise.

### GetIncomeTaxYearStartMonthOk

`func (o *QbdCompany) GetIncomeTaxYearStartMonthOk() (*string, bool)`

GetIncomeTaxYearStartMonthOk returns a tuple with the IncomeTaxYearStartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTaxYearStartMonth

`func (o *QbdCompany) SetIncomeTaxYearStartMonth(v string)`

SetIncomeTaxYearStartMonth sets IncomeTaxYearStartMonth field to given value.


### SetIncomeTaxYearStartMonthNil

`func (o *QbdCompany) SetIncomeTaxYearStartMonthNil(b bool)`

 SetIncomeTaxYearStartMonthNil sets the value for IncomeTaxYearStartMonth to be an explicit nil

### UnsetIncomeTaxYearStartMonth
`func (o *QbdCompany) UnsetIncomeTaxYearStartMonth()`

UnsetIncomeTaxYearStartMonth ensures that no value is present for IncomeTaxYearStartMonth, not even an explicit nil
### GetCompanyType

`func (o *QbdCompany) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *QbdCompany) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *QbdCompany) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.


### SetCompanyTypeNil

`func (o *QbdCompany) SetCompanyTypeNil(b bool)`

 SetCompanyTypeNil sets the value for CompanyType to be an explicit nil

### UnsetCompanyType
`func (o *QbdCompany) UnsetCompanyType()`

UnsetCompanyType ensures that no value is present for CompanyType, not even an explicit nil
### GetEin

`func (o *QbdCompany) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *QbdCompany) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *QbdCompany) SetEin(v string)`

SetEin sets Ein field to given value.


### SetEinNil

`func (o *QbdCompany) SetEinNil(b bool)`

 SetEinNil sets the value for Ein to be an explicit nil

### UnsetEin
`func (o *QbdCompany) UnsetEin()`

UnsetEin ensures that no value is present for Ein, not even an explicit nil
### GetSsn

`func (o *QbdCompany) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *QbdCompany) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *QbdCompany) SetSsn(v string)`

SetSsn sets Ssn field to given value.


### SetSsnNil

`func (o *QbdCompany) SetSsnNil(b bool)`

 SetSsnNil sets the value for Ssn to be an explicit nil

### UnsetSsn
`func (o *QbdCompany) UnsetSsn()`

UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
### GetTaxForm

`func (o *QbdCompany) GetTaxForm() string`

GetTaxForm returns the TaxForm field if non-nil, zero value otherwise.

### GetTaxFormOk

`func (o *QbdCompany) GetTaxFormOk() (*string, bool)`

GetTaxFormOk returns a tuple with the TaxForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxForm

`func (o *QbdCompany) SetTaxForm(v string)`

SetTaxForm sets TaxForm field to given value.


### SetTaxFormNil

`func (o *QbdCompany) SetTaxFormNil(b bool)`

 SetTaxFormNil sets the value for TaxForm to be an explicit nil

### UnsetTaxForm
`func (o *QbdCompany) UnsetTaxForm()`

UnsetTaxForm ensures that no value is present for TaxForm, not even an explicit nil
### GetSubscribedServices

`func (o *QbdCompany) GetSubscribedServices() QbdSubscribedService`

GetSubscribedServices returns the SubscribedServices field if non-nil, zero value otherwise.

### GetSubscribedServicesOk

`func (o *QbdCompany) GetSubscribedServicesOk() (*QbdSubscribedService, bool)`

GetSubscribedServicesOk returns a tuple with the SubscribedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedServices

`func (o *QbdCompany) SetSubscribedServices(v QbdSubscribedService)`

SetSubscribedServices sets SubscribedServices field to given value.


### SetSubscribedServicesNil

`func (o *QbdCompany) SetSubscribedServicesNil(b bool)`

 SetSubscribedServicesNil sets the value for SubscribedServices to be an explicit nil

### UnsetSubscribedServices
`func (o *QbdCompany) UnsetSubscribedServices()`

UnsetSubscribedServices ensures that no value is present for SubscribedServices, not even an explicit nil
### GetAccountantCopy

`func (o *QbdCompany) GetAccountantCopy() QbdAccountantCopy`

GetAccountantCopy returns the AccountantCopy field if non-nil, zero value otherwise.

### GetAccountantCopyOk

`func (o *QbdCompany) GetAccountantCopyOk() (*QbdAccountantCopy, bool)`

GetAccountantCopyOk returns a tuple with the AccountantCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountantCopy

`func (o *QbdCompany) SetAccountantCopy(v QbdAccountantCopy)`

SetAccountantCopy sets AccountantCopy field to given value.


### SetAccountantCopyNil

`func (o *QbdCompany) SetAccountantCopyNil(b bool)`

 SetAccountantCopyNil sets the value for AccountantCopy to be an explicit nil

### UnsetAccountantCopy
`func (o *QbdCompany) UnsetAccountantCopy()`

UnsetAccountantCopy ensures that no value is present for AccountantCopy, not even an explicit nil
### GetCustomFields

`func (o *QbdCompany) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCompany) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCompany) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


