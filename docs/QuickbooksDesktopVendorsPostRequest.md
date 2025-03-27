# QuickbooksDesktopVendorsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this vendor, unique across all vendors.  **NOTE**: Vendors do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 41 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this vendor is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The vendor&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company associated with this vendor. This name is used on invoices, checks, and other forms. | [optional] 
**Salutation** | Pointer to **string** | The formal salutation title that precedes the name of the contact person for this vendor, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact person for this vendor.  Maximum length: 25 characters. | [optional] 
**MiddleName** | Pointer to **string** | The middle name of the contact person for this vendor.  Maximum length: 5 characters. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact person for this vendor.  Maximum length: 25 characters. | [optional] 
**JobTitle** | Pointer to **string** | The job title of the contact person for this vendor. | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopVendorsPostRequestBillingAddress**](QuickbooksDesktopVendorsPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopVendorsPostRequestShippingAddress**](QuickbooksDesktopVendorsPostRequestShippingAddress.md) |  | [optional] 
**Phone** | Pointer to **string** | The vendor&#39;s primary telephone number. | [optional] 
**AlternatePhone** | Pointer to **string** | The vendor&#39;s alternate telephone number. | [optional] 
**Fax** | Pointer to **string** | The vendor&#39;s fax number. | [optional] 
**Email** | Pointer to **string** | The vendor&#39;s email address. | [optional] 
**CcEmail** | Pointer to **string** | An email address to carbon copy (CC) on communications with this vendor. | [optional] 
**Contact** | Pointer to **string** | The name of the primary contact person for this vendor. | [optional] 
**AlternateContact** | Pointer to **string** | The name of a alternate contact person for this vendor. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this vendor, such as phone numbers or email addresses. | [optional] 
**AdditionalContacts** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAdditionalContactsInner**](QuickbooksDesktopCustomersPostRequestAdditionalContactsInner.md) | Additional alternate contacts for this vendor. | [optional] 
**NameOnCheck** | Pointer to **string** | The vendor&#39;s name as it should appear on checks issued to this vendor. | [optional] 
**AccountNumber** | Pointer to **string** | The vendor&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**Note** | Pointer to **string** | A note or comment about this vendor. | [optional] 
**AdditionalNotes** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner**](QuickbooksDesktopCustomersPostRequestAdditionalNotesInner.md) | Additional notes about this vendor. | [optional] 
**VendorTypeId** | Pointer to **string** | The vendor&#39;s type, used for categorizing vendors into meaningful segments, such as industry or region. | [optional] 
**TermsId** | Pointer to **string** | The vendor&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**CreditLimit** | Pointer to **string** | The vendor&#39;s credit limit, represented as a decimal string. This is the maximum amount of money that can be spent being before billed by this vendor. If &#x60;null&#x60;, there is no credit limit. | [optional] 
**TaxIdentificationNumber** | Pointer to **string** | The vendor&#39;s tax identification number (e.g., EIN or SSN). | [optional] 
**IsEligibleFor1099** | Pointer to **bool** | Indicates whether this vendor is eligible to receive a 1099 form for tax reporting purposes. When &#x60;true&#x60;, then the fields &#x60;taxId&#x60; and &#x60;billingAddress&#x60; are required. | [optional] 
**OpeningBalance** | Pointer to **string** | The opening balance of this vendor&#39;s account, indicating the amount owed to this vendor, represented as a decimal string. | [optional] 
**OpeningBalanceDate** | Pointer to **string** | The date of the opening balance of this vendor, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**BillingRateId** | Pointer to **string** | The vendor&#39;s billing rate, used to override service item rates in time tracking activities. | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for transactions with this vendor, determining whether the transactions are taxable or non-taxable. This can be overridden at the transaction or transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesTaxCountry** | Pointer to **string** | The country for which sales tax is collected for this vendor. | [optional] 
**IsSalesTaxAgency** | Pointer to **bool** | Indicates whether this vendor is a sales tax agency. | [optional] 
**SalesTaxReturnId** | Pointer to **string** | The vendor&#39;s sales tax return information, used for tracking and reporting sales tax liabilities. | [optional] 
**TaxRegistrationNumber** | Pointer to **string** | The vendor&#39;s tax registration number, for use in Canada or the UK. | [optional] 
**ReportingPeriod** | Pointer to **string** | The vendor&#39;s tax reporting period, for use in Canada or the UK. | [optional] 
**IsTrackingPurchaseTax** | Pointer to **bool** | Indicates whether tax is tracked on purchases for this vendor, for use in Canada or the UK. | [optional] 
**PurchaseTaxAccountId** | Pointer to **string** | The account used for tracking taxes on purchases for this vendor, for use in Canada or the UK. | [optional] 
**IsTrackingSalesTax** | Pointer to **bool** | Indicates whether tax is tracked on sales for this vendor, for use in Canada or the UK. | [optional] 
**SalesTaxAccountId** | Pointer to **string** | The account used for tracking taxes on sales for this vendor, for use in Canada or the UK. | [optional] 
**IsCompoundingTax** | Pointer to **bool** | Indicates whether tax is charged on top of tax for this vendor, for use in Canada or the UK. | [optional] 
**DefaultExpenseAccountIds** | Pointer to **[]string** | The expense accounts to prefill when entering bills for this vendor. | [optional] 
**CurrencyId** | Pointer to **string** | The vendor&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 

## Methods

### NewQuickbooksDesktopVendorsPostRequest

`func NewQuickbooksDesktopVendorsPostRequest(name string, ) *QuickbooksDesktopVendorsPostRequest`

NewQuickbooksDesktopVendorsPostRequest instantiates a new QuickbooksDesktopVendorsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopVendorsPostRequestWithDefaults

`func NewQuickbooksDesktopVendorsPostRequestWithDefaults() *QuickbooksDesktopVendorsPostRequest`

NewQuickbooksDesktopVendorsPostRequestWithDefaults instantiates a new QuickbooksDesktopVendorsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopVendorsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopVendorsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopVendorsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopVendorsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopVendorsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetCompanyName

`func (o *QuickbooksDesktopVendorsPostRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuickbooksDesktopVendorsPostRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuickbooksDesktopVendorsPostRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSalutation

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopVendorsPostRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopVendorsPostRequest) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopVendorsPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopVendorsPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopVendorsPostRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopVendorsPostRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopVendorsPostRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopVendorsPostRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopVendorsPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopVendorsPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopVendorsPostRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopVendorsPostRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopVendorsPostRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopVendorsPostRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) GetBillingAddress() QuickbooksDesktopVendorsPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetBillingAddressOk() (*QuickbooksDesktopVendorsPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) SetBillingAddress(v QuickbooksDesktopVendorsPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) GetShippingAddress() QuickbooksDesktopVendorsPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetShippingAddressOk() (*QuickbooksDesktopVendorsPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) SetShippingAddress(v QuickbooksDesktopVendorsPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopVendorsPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopVendorsPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopVendorsPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopVendorsPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QuickbooksDesktopVendorsPostRequest) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QuickbooksDesktopVendorsPostRequest) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QuickbooksDesktopVendorsPostRequest) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopVendorsPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopVendorsPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopVendorsPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopVendorsPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopVendorsPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopVendorsPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCcEmail

`func (o *QuickbooksDesktopVendorsPostRequest) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *QuickbooksDesktopVendorsPostRequest) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *QuickbooksDesktopVendorsPostRequest) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetContact

`func (o *QuickbooksDesktopVendorsPostRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuickbooksDesktopVendorsPostRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuickbooksDesktopVendorsPostRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAlternateContact

`func (o *QuickbooksDesktopVendorsPostRequest) GetAlternateContact() string`

GetAlternateContact returns the AlternateContact field if non-nil, zero value otherwise.

### GetAlternateContactOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetAlternateContactOk() (*string, bool)`

GetAlternateContactOk returns a tuple with the AlternateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateContact

`func (o *QuickbooksDesktopVendorsPostRequest) SetAlternateContact(v string)`

SetAlternateContact sets AlternateContact field to given value.

### HasAlternateContact

`func (o *QuickbooksDesktopVendorsPostRequest) HasAlternateContact() bool`

HasAlternateContact returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopVendorsPostRequest) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopVendorsPostRequest) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopVendorsPostRequest) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *QuickbooksDesktopVendorsPostRequest) GetAdditionalContacts() []QuickbooksDesktopCustomersPostRequestAdditionalContactsInner`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetAdditionalContactsOk() (*[]QuickbooksDesktopCustomersPostRequestAdditionalContactsInner, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *QuickbooksDesktopVendorsPostRequest) SetAdditionalContacts(v []QuickbooksDesktopCustomersPostRequestAdditionalContactsInner)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *QuickbooksDesktopVendorsPostRequest) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetNameOnCheck

`func (o *QuickbooksDesktopVendorsPostRequest) GetNameOnCheck() string`

GetNameOnCheck returns the NameOnCheck field if non-nil, zero value otherwise.

### GetNameOnCheckOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetNameOnCheckOk() (*string, bool)`

GetNameOnCheckOk returns a tuple with the NameOnCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCheck

`func (o *QuickbooksDesktopVendorsPostRequest) SetNameOnCheck(v string)`

SetNameOnCheck sets NameOnCheck field to given value.

### HasNameOnCheck

`func (o *QuickbooksDesktopVendorsPostRequest) HasNameOnCheck() bool`

HasNameOnCheck returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QuickbooksDesktopVendorsPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopVendorsPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopVendorsPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopVendorsPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopVendorsPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopVendorsPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QuickbooksDesktopVendorsPostRequest) GetAdditionalNotes() []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetAdditionalNotesOk() (*[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QuickbooksDesktopVendorsPostRequest) SetAdditionalNotes(v []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QuickbooksDesktopVendorsPostRequest) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetVendorTypeId

`func (o *QuickbooksDesktopVendorsPostRequest) GetVendorTypeId() string`

GetVendorTypeId returns the VendorTypeId field if non-nil, zero value otherwise.

### GetVendorTypeIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetVendorTypeIdOk() (*string, bool)`

GetVendorTypeIdOk returns a tuple with the VendorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorTypeId

`func (o *QuickbooksDesktopVendorsPostRequest) SetVendorTypeId(v string)`

SetVendorTypeId sets VendorTypeId field to given value.

### HasVendorTypeId

`func (o *QuickbooksDesktopVendorsPostRequest) HasVendorTypeId() bool`

HasVendorTypeId returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopVendorsPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopVendorsPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopVendorsPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetCreditLimit

`func (o *QuickbooksDesktopVendorsPostRequest) GetCreditLimit() string`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetCreditLimitOk() (*string, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *QuickbooksDesktopVendorsPostRequest) SetCreditLimit(v string)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *QuickbooksDesktopVendorsPostRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### GetIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsEligibleFor1099() bool`

GetIsEligibleFor1099 returns the IsEligibleFor1099 field if non-nil, zero value otherwise.

### GetIsEligibleFor1099Ok

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsEligibleFor1099Ok() (*bool, bool)`

GetIsEligibleFor1099Ok returns a tuple with the IsEligibleFor1099 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsEligibleFor1099(v bool)`

SetIsEligibleFor1099 sets IsEligibleFor1099 field to given value.

### HasIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsEligibleFor1099() bool`

HasIsEligibleFor1099 returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *QuickbooksDesktopVendorsPostRequest) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *QuickbooksDesktopVendorsPostRequest) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *QuickbooksDesktopVendorsPostRequest) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *QuickbooksDesktopVendorsPostRequest) GetOpeningBalanceDate() string`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetOpeningBalanceDateOk() (*string, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *QuickbooksDesktopVendorsPostRequest) SetOpeningBalanceDate(v string)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *QuickbooksDesktopVendorsPostRequest) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### GetBillingRateId

`func (o *QuickbooksDesktopVendorsPostRequest) GetBillingRateId() string`

GetBillingRateId returns the BillingRateId field if non-nil, zero value otherwise.

### GetBillingRateIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetBillingRateIdOk() (*string, bool)`

GetBillingRateIdOk returns a tuple with the BillingRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRateId

`func (o *QuickbooksDesktopVendorsPostRequest) SetBillingRateId(v string)`

SetBillingRateId sets BillingRateId field to given value.

### HasBillingRateId

`func (o *QuickbooksDesktopVendorsPostRequest) HasBillingRateId() bool`

HasBillingRateId returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopVendorsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopVendorsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopVendorsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesTaxCountry

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxCountry() string`

GetSalesTaxCountry returns the SalesTaxCountry field if non-nil, zero value otherwise.

### GetSalesTaxCountryOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxCountryOk() (*string, bool)`

GetSalesTaxCountryOk returns a tuple with the SalesTaxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCountry

`func (o *QuickbooksDesktopVendorsPostRequest) SetSalesTaxCountry(v string)`

SetSalesTaxCountry sets SalesTaxCountry field to given value.

### HasSalesTaxCountry

`func (o *QuickbooksDesktopVendorsPostRequest) HasSalesTaxCountry() bool`

HasSalesTaxCountry returns a boolean if a field has been set.

### GetIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsSalesTaxAgency() bool`

GetIsSalesTaxAgency returns the IsSalesTaxAgency field if non-nil, zero value otherwise.

### GetIsSalesTaxAgencyOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsSalesTaxAgencyOk() (*bool, bool)`

GetIsSalesTaxAgencyOk returns a tuple with the IsSalesTaxAgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsSalesTaxAgency(v bool)`

SetIsSalesTaxAgency sets IsSalesTaxAgency field to given value.

### HasIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsSalesTaxAgency() bool`

HasIsSalesTaxAgency returns a boolean if a field has been set.

### GetSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxReturnId() string`

GetSalesTaxReturnId returns the SalesTaxReturnId field if non-nil, zero value otherwise.

### GetSalesTaxReturnIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxReturnIdOk() (*string, bool)`

GetSalesTaxReturnIdOk returns a tuple with the SalesTaxReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsPostRequest) SetSalesTaxReturnId(v string)`

SetSalesTaxReturnId sets SalesTaxReturnId field to given value.

### HasSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsPostRequest) HasSalesTaxReturnId() bool`

HasSalesTaxReturnId returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsPostRequest) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetReportingPeriod

`func (o *QuickbooksDesktopVendorsPostRequest) GetReportingPeriod() string`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetReportingPeriodOk() (*string, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *QuickbooksDesktopVendorsPostRequest) SetReportingPeriod(v string)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *QuickbooksDesktopVendorsPostRequest) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsTrackingPurchaseTax() bool`

GetIsTrackingPurchaseTax returns the IsTrackingPurchaseTax field if non-nil, zero value otherwise.

### GetIsTrackingPurchaseTaxOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsTrackingPurchaseTaxOk() (*bool, bool)`

GetIsTrackingPurchaseTaxOk returns a tuple with the IsTrackingPurchaseTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsTrackingPurchaseTax(v bool)`

SetIsTrackingPurchaseTax sets IsTrackingPurchaseTax field to given value.

### HasIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsTrackingPurchaseTax() bool`

HasIsTrackingPurchaseTax returns a boolean if a field has been set.

### GetPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) GetPurchaseTaxAccountId() string`

GetPurchaseTaxAccountId returns the PurchaseTaxAccountId field if non-nil, zero value otherwise.

### GetPurchaseTaxAccountIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetPurchaseTaxAccountIdOk() (*string, bool)`

GetPurchaseTaxAccountIdOk returns a tuple with the PurchaseTaxAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) SetPurchaseTaxAccountId(v string)`

SetPurchaseTaxAccountId sets PurchaseTaxAccountId field to given value.

### HasPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) HasPurchaseTaxAccountId() bool`

HasPurchaseTaxAccountId returns a boolean if a field has been set.

### GetIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsTrackingSalesTax() bool`

GetIsTrackingSalesTax returns the IsTrackingSalesTax field if non-nil, zero value otherwise.

### GetIsTrackingSalesTaxOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsTrackingSalesTaxOk() (*bool, bool)`

GetIsTrackingSalesTaxOk returns a tuple with the IsTrackingSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsTrackingSalesTax(v bool)`

SetIsTrackingSalesTax sets IsTrackingSalesTax field to given value.

### HasIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsTrackingSalesTax() bool`

HasIsTrackingSalesTax returns a boolean if a field has been set.

### GetSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxAccountId() string`

GetSalesTaxAccountId returns the SalesTaxAccountId field if non-nil, zero value otherwise.

### GetSalesTaxAccountIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetSalesTaxAccountIdOk() (*string, bool)`

GetSalesTaxAccountIdOk returns a tuple with the SalesTaxAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) SetSalesTaxAccountId(v string)`

SetSalesTaxAccountId sets SalesTaxAccountId field to given value.

### HasSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsPostRequest) HasSalesTaxAccountId() bool`

HasSalesTaxAccountId returns a boolean if a field has been set.

### GetIsCompoundingTax

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsCompoundingTax() bool`

GetIsCompoundingTax returns the IsCompoundingTax field if non-nil, zero value otherwise.

### GetIsCompoundingTaxOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetIsCompoundingTaxOk() (*bool, bool)`

GetIsCompoundingTaxOk returns a tuple with the IsCompoundingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompoundingTax

`func (o *QuickbooksDesktopVendorsPostRequest) SetIsCompoundingTax(v bool)`

SetIsCompoundingTax sets IsCompoundingTax field to given value.

### HasIsCompoundingTax

`func (o *QuickbooksDesktopVendorsPostRequest) HasIsCompoundingTax() bool`

HasIsCompoundingTax returns a boolean if a field has been set.

### GetDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsPostRequest) GetDefaultExpenseAccountIds() []string`

GetDefaultExpenseAccountIds returns the DefaultExpenseAccountIds field if non-nil, zero value otherwise.

### GetDefaultExpenseAccountIdsOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetDefaultExpenseAccountIdsOk() (*[]string, bool)`

GetDefaultExpenseAccountIdsOk returns a tuple with the DefaultExpenseAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsPostRequest) SetDefaultExpenseAccountIds(v []string)`

SetDefaultExpenseAccountIds sets DefaultExpenseAccountIds field to given value.

### HasDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsPostRequest) HasDefaultExpenseAccountIds() bool`

HasDefaultExpenseAccountIds returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopVendorsPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopVendorsPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopVendorsPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopVendorsPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


