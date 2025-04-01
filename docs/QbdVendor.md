# QbdVendor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this vendor. This ID is unique across all vendors but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_vendor\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this vendor was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this vendor was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this vendor object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this vendor, unique across all vendors.  **NOTE**: Vendors do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this vendor is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | Pointer to [**QbdVendorClass**](QbdVendorClass.md) |  | [optional] 
**CompanyName** | Pointer to **string** | The name of the company associated with this vendor. This name is used on invoices, checks, and other forms. | [optional] 
**Salutation** | Pointer to **string** | The formal salutation title that precedes the name of the contact person for this vendor, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact person for this vendor. | [optional] 
**MiddleName** | Pointer to **string** | The middle name of the contact person for this vendor. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact person for this vendor. | [optional] 
**JobTitle** | Pointer to **string** | The job title of the contact person for this vendor. | [optional] 
**BillingAddress** | Pointer to [**QbdVendorBillingAddress**](QbdVendorBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QbdVendorShippingAddress**](QbdVendorShippingAddress.md) |  | [optional] 
**Phone** | Pointer to **string** | The vendor&#39;s primary telephone number. | [optional] 
**AlternatePhone** | Pointer to **string** | The vendor&#39;s alternate telephone number. | [optional] 
**Fax** | Pointer to **string** | The vendor&#39;s fax number. | [optional] 
**Email** | Pointer to **string** | The vendor&#39;s email address. | [optional] 
**CcEmail** | Pointer to **string** | An email address to carbon copy (CC) on communications with this vendor. | [optional] 
**Contact** | Pointer to **string** | The name of the primary contact person for this vendor. | [optional] 
**AlternateContact** | Pointer to **string** | The name of a alternate contact person for this vendor. | [optional] 
**CustomContactFields** | [**[]QbdCustomContactField**](QbdCustomContactField.md) | Additional custom contact fields for this vendor, such as phone numbers or email addresses. | 
**AdditionalContacts** | [**[]QbdContact**](QbdContact.md) | Additional alternate contacts for this vendor. | 
**NameOnCheck** | Pointer to **string** | The vendor&#39;s name as it should appear on checks issued to this vendor. | [optional] 
**AccountNumber** | Pointer to **string** | The vendor&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**Note** | Pointer to **string** | A note or comment about this vendor. | [optional] 
**AdditionalNotes** | [**[]QbdNote**](QbdNote.md) | Additional notes about this vendor. | 
**VendorType** | Pointer to [**QbdVendorVendorType**](QbdVendorVendorType.md) |  | [optional] 
**Terms** | Pointer to [**QbdVendorTerms**](QbdVendorTerms.md) |  | [optional] 
**CreditLimit** | Pointer to **string** | The vendor&#39;s credit limit, represented as a decimal string. This is the maximum amount of money that can be spent being before billed by this vendor. If &#x60;null&#x60;, there is no credit limit. | [optional] 
**TaxIdentificationNumber** | Pointer to **string** | The vendor&#39;s tax identification number (e.g., EIN or SSN). | [optional] 
**IsEligibleFor1099** | Pointer to **bool** | Indicates whether this vendor is eligible to receive a 1099 form for tax reporting purposes. When &#x60;true&#x60;, then the fields &#x60;taxId&#x60; and &#x60;billingAddress&#x60; are required. | [optional] 
**Balance** | Pointer to **string** | The current balance owed to this vendor, represented as a decimal string. A positive number indicates money owed to the vendor. | [optional] 
**BillingRate** | Pointer to [**QbdVendorBillingRate**](QbdVendorBillingRate.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | [optional] 
**SalesTaxCode** | Pointer to [**QbdVendorSalesTaxCode**](QbdVendorSalesTaxCode.md) |  | [optional] 
**SalesTaxCountry** | Pointer to **string** | The country for which sales tax is collected for this vendor. | [optional] 
**IsSalesTaxAgency** | Pointer to **bool** | Indicates whether this vendor is a sales tax agency. | [optional] 
**SalesTaxReturn** | Pointer to [**QbdVendorSalesTaxReturn**](QbdVendorSalesTaxReturn.md) |  | [optional] 
**TaxRegistrationNumber** | Pointer to **string** | The vendor&#39;s tax registration number, for use in Canada or the UK. | [optional] 
**ReportingPeriod** | Pointer to **string** | The vendor&#39;s tax reporting period, for use in Canada or the UK. | [optional] 
**IsTrackingPurchaseTax** | Pointer to **bool** | Indicates whether tax is tracked on purchases for this vendor, for use in Canada or the UK. | [optional] 
**PurchaseTaxAccount** | Pointer to [**QbdVendorPurchaseTaxAccount**](QbdVendorPurchaseTaxAccount.md) |  | [optional] 
**IsTrackingSalesTax** | Pointer to **bool** | Indicates whether tax is tracked on sales for this vendor, for use in Canada or the UK. | [optional] 
**SalesTaxAccount** | Pointer to [**QbdVendorSalesTaxAccount**](QbdVendorSalesTaxAccount.md) |  | [optional] 
**IsCompoundingTax** | Pointer to **bool** | Indicates whether tax is charged on top of tax for this vendor, for use in Canada or the UK. | [optional] 
**DefaultExpenseAccounts** | [**[]QbdVendorDefaultExpenseAccountsInner**](QbdVendorDefaultExpenseAccountsInner.md) | The expense accounts to prefill when entering bills for this vendor. | 
**Currency** | Pointer to [**QbdVendorCurrency**](QbdVendorCurrency.md) |  | [optional] 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the vendor object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdVendor

`func NewQbdVendor(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, customContactFields []QbdCustomContactField, additionalContacts []QbdContact, additionalNotes []QbdNote, defaultExpenseAccounts []QbdVendorDefaultExpenseAccountsInner, customFields []QbdCustomField, ) *QbdVendor`

NewQbdVendor instantiates a new QbdVendor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdVendorWithDefaults

`func NewQbdVendorWithDefaults() *QbdVendor`

NewQbdVendorWithDefaults instantiates a new QbdVendor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdVendor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdVendor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdVendor) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdVendor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdVendor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdVendor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdVendor) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdVendor) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdVendor) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdVendor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdVendor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdVendor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdVendor) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdVendor) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdVendor) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdVendor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdVendor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdVendor) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdVendor) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdVendor) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdVendor) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdVendor) GetClass() QbdVendorClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdVendor) GetClassOk() (*QbdVendorClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdVendor) SetClass(v QbdVendorClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *QbdVendor) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCompanyName

`func (o *QbdVendor) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QbdVendor) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QbdVendor) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QbdVendor) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSalutation

`func (o *QbdVendor) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QbdVendor) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QbdVendor) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QbdVendor) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QbdVendor) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QbdVendor) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QbdVendor) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QbdVendor) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QbdVendor) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QbdVendor) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QbdVendor) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QbdVendor) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QbdVendor) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QbdVendor) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QbdVendor) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QbdVendor) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QbdVendor) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QbdVendor) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QbdVendor) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QbdVendor) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QbdVendor) GetBillingAddress() QbdVendorBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdVendor) GetBillingAddressOk() (*QbdVendorBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdVendor) SetBillingAddress(v QbdVendorBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QbdVendor) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QbdVendor) GetShippingAddress() QbdVendorShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdVendor) GetShippingAddressOk() (*QbdVendorShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdVendor) SetShippingAddress(v QbdVendorShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QbdVendor) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetPhone

`func (o *QbdVendor) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QbdVendor) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QbdVendor) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QbdVendor) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QbdVendor) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QbdVendor) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QbdVendor) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QbdVendor) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QbdVendor) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QbdVendor) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QbdVendor) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QbdVendor) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QbdVendor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QbdVendor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QbdVendor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QbdVendor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCcEmail

`func (o *QbdVendor) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *QbdVendor) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *QbdVendor) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *QbdVendor) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetContact

`func (o *QbdVendor) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QbdVendor) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QbdVendor) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QbdVendor) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAlternateContact

`func (o *QbdVendor) GetAlternateContact() string`

GetAlternateContact returns the AlternateContact field if non-nil, zero value otherwise.

### GetAlternateContactOk

`func (o *QbdVendor) GetAlternateContactOk() (*string, bool)`

GetAlternateContactOk returns a tuple with the AlternateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateContact

`func (o *QbdVendor) SetAlternateContact(v string)`

SetAlternateContact sets AlternateContact field to given value.

### HasAlternateContact

`func (o *QbdVendor) HasAlternateContact() bool`

HasAlternateContact returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QbdVendor) GetCustomContactFields() []QbdCustomContactField`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QbdVendor) GetCustomContactFieldsOk() (*[]QbdCustomContactField, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QbdVendor) SetCustomContactFields(v []QbdCustomContactField)`

SetCustomContactFields sets CustomContactFields field to given value.


### GetAdditionalContacts

`func (o *QbdVendor) GetAdditionalContacts() []QbdContact`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *QbdVendor) GetAdditionalContactsOk() (*[]QbdContact, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *QbdVendor) SetAdditionalContacts(v []QbdContact)`

SetAdditionalContacts sets AdditionalContacts field to given value.


### GetNameOnCheck

`func (o *QbdVendor) GetNameOnCheck() string`

GetNameOnCheck returns the NameOnCheck field if non-nil, zero value otherwise.

### GetNameOnCheckOk

`func (o *QbdVendor) GetNameOnCheckOk() (*string, bool)`

GetNameOnCheckOk returns a tuple with the NameOnCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCheck

`func (o *QbdVendor) SetNameOnCheck(v string)`

SetNameOnCheck sets NameOnCheck field to given value.

### HasNameOnCheck

`func (o *QbdVendor) HasNameOnCheck() bool`

HasNameOnCheck returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QbdVendor) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QbdVendor) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QbdVendor) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QbdVendor) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetNote

`func (o *QbdVendor) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdVendor) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdVendor) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QbdVendor) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QbdVendor) GetAdditionalNotes() []QbdNote`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QbdVendor) GetAdditionalNotesOk() (*[]QbdNote, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QbdVendor) SetAdditionalNotes(v []QbdNote)`

SetAdditionalNotes sets AdditionalNotes field to given value.


### GetVendorType

`func (o *QbdVendor) GetVendorType() QbdVendorVendorType`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *QbdVendor) GetVendorTypeOk() (*QbdVendorVendorType, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *QbdVendor) SetVendorType(v QbdVendorVendorType)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *QbdVendor) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.

### GetTerms

`func (o *QbdVendor) GetTerms() QbdVendorTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdVendor) GetTermsOk() (*QbdVendorTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdVendor) SetTerms(v QbdVendorTerms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *QbdVendor) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetCreditLimit

`func (o *QbdVendor) GetCreditLimit() string`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *QbdVendor) GetCreditLimitOk() (*string, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *QbdVendor) SetCreditLimit(v string)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *QbdVendor) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *QbdVendor) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *QbdVendor) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *QbdVendor) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *QbdVendor) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### GetIsEligibleFor1099

`func (o *QbdVendor) GetIsEligibleFor1099() bool`

GetIsEligibleFor1099 returns the IsEligibleFor1099 field if non-nil, zero value otherwise.

### GetIsEligibleFor1099Ok

`func (o *QbdVendor) GetIsEligibleFor1099Ok() (*bool, bool)`

GetIsEligibleFor1099Ok returns a tuple with the IsEligibleFor1099 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleFor1099

`func (o *QbdVendor) SetIsEligibleFor1099(v bool)`

SetIsEligibleFor1099 sets IsEligibleFor1099 field to given value.

### HasIsEligibleFor1099

`func (o *QbdVendor) HasIsEligibleFor1099() bool`

HasIsEligibleFor1099 returns a boolean if a field has been set.

### GetBalance

`func (o *QbdVendor) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *QbdVendor) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *QbdVendor) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *QbdVendor) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBillingRate

`func (o *QbdVendor) GetBillingRate() QbdVendorBillingRate`

GetBillingRate returns the BillingRate field if non-nil, zero value otherwise.

### GetBillingRateOk

`func (o *QbdVendor) GetBillingRateOk() (*QbdVendorBillingRate, bool)`

GetBillingRateOk returns a tuple with the BillingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRate

`func (o *QbdVendor) SetBillingRate(v QbdVendorBillingRate)`

SetBillingRate sets BillingRate field to given value.

### HasBillingRate

`func (o *QbdVendor) HasBillingRate() bool`

HasBillingRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QbdVendor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdVendor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdVendor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QbdVendor) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSalesTaxCode

`func (o *QbdVendor) GetSalesTaxCode() QbdVendorSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdVendor) GetSalesTaxCodeOk() (*QbdVendorSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdVendor) SetSalesTaxCode(v QbdVendorSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.

### HasSalesTaxCode

`func (o *QbdVendor) HasSalesTaxCode() bool`

HasSalesTaxCode returns a boolean if a field has been set.

### GetSalesTaxCountry

`func (o *QbdVendor) GetSalesTaxCountry() string`

GetSalesTaxCountry returns the SalesTaxCountry field if non-nil, zero value otherwise.

### GetSalesTaxCountryOk

`func (o *QbdVendor) GetSalesTaxCountryOk() (*string, bool)`

GetSalesTaxCountryOk returns a tuple with the SalesTaxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCountry

`func (o *QbdVendor) SetSalesTaxCountry(v string)`

SetSalesTaxCountry sets SalesTaxCountry field to given value.

### HasSalesTaxCountry

`func (o *QbdVendor) HasSalesTaxCountry() bool`

HasSalesTaxCountry returns a boolean if a field has been set.

### GetIsSalesTaxAgency

`func (o *QbdVendor) GetIsSalesTaxAgency() bool`

GetIsSalesTaxAgency returns the IsSalesTaxAgency field if non-nil, zero value otherwise.

### GetIsSalesTaxAgencyOk

`func (o *QbdVendor) GetIsSalesTaxAgencyOk() (*bool, bool)`

GetIsSalesTaxAgencyOk returns a tuple with the IsSalesTaxAgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSalesTaxAgency

`func (o *QbdVendor) SetIsSalesTaxAgency(v bool)`

SetIsSalesTaxAgency sets IsSalesTaxAgency field to given value.

### HasIsSalesTaxAgency

`func (o *QbdVendor) HasIsSalesTaxAgency() bool`

HasIsSalesTaxAgency returns a boolean if a field has been set.

### GetSalesTaxReturn

`func (o *QbdVendor) GetSalesTaxReturn() QbdVendorSalesTaxReturn`

GetSalesTaxReturn returns the SalesTaxReturn field if non-nil, zero value otherwise.

### GetSalesTaxReturnOk

`func (o *QbdVendor) GetSalesTaxReturnOk() (*QbdVendorSalesTaxReturn, bool)`

GetSalesTaxReturnOk returns a tuple with the SalesTaxReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxReturn

`func (o *QbdVendor) SetSalesTaxReturn(v QbdVendorSalesTaxReturn)`

SetSalesTaxReturn sets SalesTaxReturn field to given value.

### HasSalesTaxReturn

`func (o *QbdVendor) HasSalesTaxReturn() bool`

HasSalesTaxReturn returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *QbdVendor) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *QbdVendor) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *QbdVendor) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *QbdVendor) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetReportingPeriod

`func (o *QbdVendor) GetReportingPeriod() string`

GetReportingPeriod returns the ReportingPeriod field if non-nil, zero value otherwise.

### GetReportingPeriodOk

`func (o *QbdVendor) GetReportingPeriodOk() (*string, bool)`

GetReportingPeriodOk returns a tuple with the ReportingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPeriod

`func (o *QbdVendor) SetReportingPeriod(v string)`

SetReportingPeriod sets ReportingPeriod field to given value.

### HasReportingPeriod

`func (o *QbdVendor) HasReportingPeriod() bool`

HasReportingPeriod returns a boolean if a field has been set.

### GetIsTrackingPurchaseTax

`func (o *QbdVendor) GetIsTrackingPurchaseTax() bool`

GetIsTrackingPurchaseTax returns the IsTrackingPurchaseTax field if non-nil, zero value otherwise.

### GetIsTrackingPurchaseTaxOk

`func (o *QbdVendor) GetIsTrackingPurchaseTaxOk() (*bool, bool)`

GetIsTrackingPurchaseTaxOk returns a tuple with the IsTrackingPurchaseTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingPurchaseTax

`func (o *QbdVendor) SetIsTrackingPurchaseTax(v bool)`

SetIsTrackingPurchaseTax sets IsTrackingPurchaseTax field to given value.

### HasIsTrackingPurchaseTax

`func (o *QbdVendor) HasIsTrackingPurchaseTax() bool`

HasIsTrackingPurchaseTax returns a boolean if a field has been set.

### GetPurchaseTaxAccount

`func (o *QbdVendor) GetPurchaseTaxAccount() QbdVendorPurchaseTaxAccount`

GetPurchaseTaxAccount returns the PurchaseTaxAccount field if non-nil, zero value otherwise.

### GetPurchaseTaxAccountOk

`func (o *QbdVendor) GetPurchaseTaxAccountOk() (*QbdVendorPurchaseTaxAccount, bool)`

GetPurchaseTaxAccountOk returns a tuple with the PurchaseTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTaxAccount

`func (o *QbdVendor) SetPurchaseTaxAccount(v QbdVendorPurchaseTaxAccount)`

SetPurchaseTaxAccount sets PurchaseTaxAccount field to given value.

### HasPurchaseTaxAccount

`func (o *QbdVendor) HasPurchaseTaxAccount() bool`

HasPurchaseTaxAccount returns a boolean if a field has been set.

### GetIsTrackingSalesTax

`func (o *QbdVendor) GetIsTrackingSalesTax() bool`

GetIsTrackingSalesTax returns the IsTrackingSalesTax field if non-nil, zero value otherwise.

### GetIsTrackingSalesTaxOk

`func (o *QbdVendor) GetIsTrackingSalesTaxOk() (*bool, bool)`

GetIsTrackingSalesTaxOk returns a tuple with the IsTrackingSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingSalesTax

`func (o *QbdVendor) SetIsTrackingSalesTax(v bool)`

SetIsTrackingSalesTax sets IsTrackingSalesTax field to given value.

### HasIsTrackingSalesTax

`func (o *QbdVendor) HasIsTrackingSalesTax() bool`

HasIsTrackingSalesTax returns a boolean if a field has been set.

### GetSalesTaxAccount

`func (o *QbdVendor) GetSalesTaxAccount() QbdVendorSalesTaxAccount`

GetSalesTaxAccount returns the SalesTaxAccount field if non-nil, zero value otherwise.

### GetSalesTaxAccountOk

`func (o *QbdVendor) GetSalesTaxAccountOk() (*QbdVendorSalesTaxAccount, bool)`

GetSalesTaxAccountOk returns a tuple with the SalesTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAccount

`func (o *QbdVendor) SetSalesTaxAccount(v QbdVendorSalesTaxAccount)`

SetSalesTaxAccount sets SalesTaxAccount field to given value.

### HasSalesTaxAccount

`func (o *QbdVendor) HasSalesTaxAccount() bool`

HasSalesTaxAccount returns a boolean if a field has been set.

### GetIsCompoundingTax

`func (o *QbdVendor) GetIsCompoundingTax() bool`

GetIsCompoundingTax returns the IsCompoundingTax field if non-nil, zero value otherwise.

### GetIsCompoundingTaxOk

`func (o *QbdVendor) GetIsCompoundingTaxOk() (*bool, bool)`

GetIsCompoundingTaxOk returns a tuple with the IsCompoundingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompoundingTax

`func (o *QbdVendor) SetIsCompoundingTax(v bool)`

SetIsCompoundingTax sets IsCompoundingTax field to given value.

### HasIsCompoundingTax

`func (o *QbdVendor) HasIsCompoundingTax() bool`

HasIsCompoundingTax returns a boolean if a field has been set.

### GetDefaultExpenseAccounts

`func (o *QbdVendor) GetDefaultExpenseAccounts() []QbdVendorDefaultExpenseAccountsInner`

GetDefaultExpenseAccounts returns the DefaultExpenseAccounts field if non-nil, zero value otherwise.

### GetDefaultExpenseAccountsOk

`func (o *QbdVendor) GetDefaultExpenseAccountsOk() (*[]QbdVendorDefaultExpenseAccountsInner, bool)`

GetDefaultExpenseAccountsOk returns a tuple with the DefaultExpenseAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpenseAccounts

`func (o *QbdVendor) SetDefaultExpenseAccounts(v []QbdVendorDefaultExpenseAccountsInner)`

SetDefaultExpenseAccounts sets DefaultExpenseAccounts field to given value.


### GetCurrency

`func (o *QbdVendor) GetCurrency() QbdVendorCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdVendor) GetCurrencyOk() (*QbdVendorCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdVendor) SetCurrency(v QbdVendorCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QbdVendor) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomFields

`func (o *QbdVendor) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdVendor) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdVendor) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


