# QuickbooksDesktopVendorsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the vendor object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive unique name of this vendor, unique across all vendors.  **NOTE**: Vendors do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 41 characters. | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this vendor is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
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
**AdditionalContacts** | Pointer to [**[]QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner**](QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner.md) | Additional alternate contacts for this vendor. | [optional] 
**NameOnCheck** | Pointer to **string** | The vendor&#39;s name as it should appear on checks issued to this vendor. | [optional] 
**AccountNumber** | Pointer to **string** | The vendor&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**Note** | Pointer to **string** | A note or comment about this vendor. | [optional] 
**AdditionalNotes** | Pointer to [**[]QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner**](QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner.md) | Additional notes about this vendor. | [optional] 
**VendorTypeId** | Pointer to **string** | The vendor&#39;s type, used for categorizing vendors into meaningful segments, such as industry or region. | [optional] 
**TermsId** | Pointer to **string** | The vendor&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**CreditLimit** | Pointer to **string** | The vendor&#39;s credit limit, represented as a decimal string. This is the maximum amount of money that can be spent being before billed by this vendor. If &#x60;null&#x60;, there is no credit limit. | [optional] 
**TaxIdentificationNumber** | Pointer to **string** | The vendor&#39;s tax identification number (e.g., EIN or SSN). | [optional] 
**IsEligibleFor1099** | Pointer to **bool** | Indicates whether this vendor is eligible to receive a 1099 form for tax reporting purposes. When &#x60;true&#x60;, then the fields &#x60;taxId&#x60; and &#x60;billingAddress&#x60; are required. | [optional] 
**BillingRateId** | Pointer to **string** | The vendor&#39;s billing rate, used to override service item rates in time tracking activities. | [optional] 
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

### NewQuickbooksDesktopVendorsIdPostRequest

`func NewQuickbooksDesktopVendorsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopVendorsIdPostRequest`

NewQuickbooksDesktopVendorsIdPostRequest instantiates a new QuickbooksDesktopVendorsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopVendorsIdPostRequestWithDefaults

`func NewQuickbooksDesktopVendorsIdPostRequestWithDefaults() *QuickbooksDesktopVendorsIdPostRequest`

NewQuickbooksDesktopVendorsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopVendorsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetCompanyName

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSalutation

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetBillingAddress() QuickbooksDesktopVendorsPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetBillingAddressOk() (*QuickbooksDesktopVendorsPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetBillingAddress(v QuickbooksDesktopVendorsPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetShippingAddress() QuickbooksDesktopVendorsPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopVendorsPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetShippingAddress(v QuickbooksDesktopVendorsPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCcEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAlternateContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAlternateContact() string`

GetAlternateContact returns the AlternateContact field if non-nil, zero value otherwise.

### GetAlternateContactOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAlternateContactOk() (*string, bool)`

GetAlternateContactOk returns a tuple with the AlternateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetAlternateContact(v string)`

SetAlternateContact sets AlternateContact field to given value.

### HasAlternateContact

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasAlternateContact() bool`

HasAlternateContact returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAdditionalContacts() []QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAdditionalContactsOk() (*[]QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetAdditionalContacts(v []QuickbooksDesktopCustomersIdPostRequestAdditionalContactsInner)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetNameOnCheck

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetNameOnCheck() string`

GetNameOnCheck returns the NameOnCheck field if non-nil, zero value otherwise.

### GetNameOnCheckOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetNameOnCheckOk() (*string, bool)`

GetNameOnCheckOk returns a tuple with the NameOnCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCheck

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetNameOnCheck(v string)`

SetNameOnCheck sets NameOnCheck field to given value.

### HasNameOnCheck

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasNameOnCheck() bool`

HasNameOnCheck returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAdditionalNotes() []QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetAdditionalNotesOk() (*[]QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetAdditionalNotes(v []QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetVendorTypeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetVendorTypeId() string`

GetVendorTypeId returns the VendorTypeId field if non-nil, zero value otherwise.

### GetVendorTypeIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetVendorTypeIdOk() (*string, bool)`

GetVendorTypeIdOk returns a tuple with the VendorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorTypeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetVendorTypeId(v string)`

SetVendorTypeId sets VendorTypeId field to given value.

### HasVendorTypeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasVendorTypeId() bool`

HasVendorTypeId returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetCreditLimit

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCreditLimit() string`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCreditLimitOk() (*string, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetCreditLimit(v string)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### GetIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsEligibleFor1099() bool`

GetIsEligibleFor1099 returns the IsEligibleFor1099 field if non-nil, zero value otherwise.

### GetIsEligibleFor1099Ok

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsEligibleFor1099Ok() (*bool, bool)`

GetIsEligibleFor1099Ok returns a tuple with the IsEligibleFor1099 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsEligibleFor1099(v bool)`

SetIsEligibleFor1099 sets IsEligibleFor1099 field to given value.

### HasIsEligibleFor1099

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsEligibleFor1099() bool`

HasIsEligibleFor1099 returns a boolean if a field has been set.

### GetBillingRateId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetBillingRateId() string`

GetBillingRateId returns the BillingRateId field if non-nil, zero value otherwise.

### GetBillingRateIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetBillingRateIdOk() (*string, bool)`

GetBillingRateIdOk returns a tuple with the BillingRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRateId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetBillingRateId(v string)`

SetBillingRateId sets BillingRateId field to given value.

### HasBillingRateId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasBillingRateId() bool`

HasBillingRateId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesTaxCountry

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxCountry() string`

GetSalesTaxCountry returns the SalesTaxCountry field if non-nil, zero value otherwise.

### GetSalesTaxCountryOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxCountryOk() (*string, bool)`

GetSalesTaxCountryOk returns a tuple with the SalesTaxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCountry

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetSalesTaxCountry(v string)`

SetSalesTaxCountry sets SalesTaxCountry field to given value.

### HasSalesTaxCountry

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasSalesTaxCountry() bool`

HasSalesTaxCountry returns a boolean if a field has been set.

### GetIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsSalesTaxAgency() bool`

GetIsSalesTaxAgency returns the IsSalesTaxAgency field if non-nil, zero value otherwise.

### GetIsSalesTaxAgencyOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsSalesTaxAgencyOk() (*bool, bool)`

GetIsSalesTaxAgencyOk returns a tuple with the IsSalesTaxAgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsSalesTaxAgency(v bool)`

SetIsSalesTaxAgency sets IsSalesTaxAgency field to given value.

### HasIsSalesTaxAgency

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsSalesTaxAgency() bool`

HasIsSalesTaxAgency returns a boolean if a field has been set.

### GetSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxReturnId() string`

GetSalesTaxReturnId returns the SalesTaxReturnId field if non-nil, zero value otherwise.

### GetSalesTaxReturnIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxReturnIdOk() (*string, bool)`

GetSalesTaxReturnIdOk returns a tuple with the SalesTaxReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetSalesTaxReturnId(v string)`

SetSalesTaxReturnId sets SalesTaxReturnId field to given value.

### HasSalesTaxReturnId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasSalesTaxReturnId() bool`

HasSalesTaxReturnId returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetReportingPeriod

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetReportingPeriod() string`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetReportingPeriodOk() (*string, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetReportingPeriod(v string)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsTrackingPurchaseTax() bool`

GetIsTrackingPurchaseTax returns the IsTrackingPurchaseTax field if non-nil, zero value otherwise.

### GetIsTrackingPurchaseTaxOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsTrackingPurchaseTaxOk() (*bool, bool)`

GetIsTrackingPurchaseTaxOk returns a tuple with the IsTrackingPurchaseTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsTrackingPurchaseTax(v bool)`

SetIsTrackingPurchaseTax sets IsTrackingPurchaseTax field to given value.

### HasIsTrackingPurchaseTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsTrackingPurchaseTax() bool`

HasIsTrackingPurchaseTax returns a boolean if a field has been set.

### GetPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetPurchaseTaxAccountId() string`

GetPurchaseTaxAccountId returns the PurchaseTaxAccountId field if non-nil, zero value otherwise.

### GetPurchaseTaxAccountIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetPurchaseTaxAccountIdOk() (*string, bool)`

GetPurchaseTaxAccountIdOk returns a tuple with the PurchaseTaxAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetPurchaseTaxAccountId(v string)`

SetPurchaseTaxAccountId sets PurchaseTaxAccountId field to given value.

### HasPurchaseTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasPurchaseTaxAccountId() bool`

HasPurchaseTaxAccountId returns a boolean if a field has been set.

### GetIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsTrackingSalesTax() bool`

GetIsTrackingSalesTax returns the IsTrackingSalesTax field if non-nil, zero value otherwise.

### GetIsTrackingSalesTaxOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsTrackingSalesTaxOk() (*bool, bool)`

GetIsTrackingSalesTaxOk returns a tuple with the IsTrackingSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsTrackingSalesTax(v bool)`

SetIsTrackingSalesTax sets IsTrackingSalesTax field to given value.

### HasIsTrackingSalesTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsTrackingSalesTax() bool`

HasIsTrackingSalesTax returns a boolean if a field has been set.

### GetSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxAccountId() string`

GetSalesTaxAccountId returns the SalesTaxAccountId field if non-nil, zero value otherwise.

### GetSalesTaxAccountIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetSalesTaxAccountIdOk() (*string, bool)`

GetSalesTaxAccountIdOk returns a tuple with the SalesTaxAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetSalesTaxAccountId(v string)`

SetSalesTaxAccountId sets SalesTaxAccountId field to given value.

### HasSalesTaxAccountId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasSalesTaxAccountId() bool`

HasSalesTaxAccountId returns a boolean if a field has been set.

### GetIsCompoundingTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsCompoundingTax() bool`

GetIsCompoundingTax returns the IsCompoundingTax field if non-nil, zero value otherwise.

### GetIsCompoundingTaxOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetIsCompoundingTaxOk() (*bool, bool)`

GetIsCompoundingTaxOk returns a tuple with the IsCompoundingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompoundingTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetIsCompoundingTax(v bool)`

SetIsCompoundingTax sets IsCompoundingTax field to given value.

### HasIsCompoundingTax

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasIsCompoundingTax() bool`

HasIsCompoundingTax returns a boolean if a field has been set.

### GetDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetDefaultExpenseAccountIds() []string`

GetDefaultExpenseAccountIds returns the DefaultExpenseAccountIds field if non-nil, zero value otherwise.

### GetDefaultExpenseAccountIdsOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetDefaultExpenseAccountIdsOk() (*[]string, bool)`

GetDefaultExpenseAccountIdsOk returns a tuple with the DefaultExpenseAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetDefaultExpenseAccountIds(v []string)`

SetDefaultExpenseAccountIds sets DefaultExpenseAccountIds field to given value.

### HasDefaultExpenseAccountIds

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasDefaultExpenseAccountIds() bool`

HasDefaultExpenseAccountIds returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopVendorsIdPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopVendorsIdPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopVendorsIdPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


