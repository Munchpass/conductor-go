# QbdCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this customer. This ID is unique across all customers but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_customer\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this customer was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this customer was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this customer object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive name of this customer. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two customers could both have the &#x60;name&#x60; \&quot;Website Redesign Project\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;ABC Corporation:Website Redesign Project\&quot; and \&quot;Baker:Website Redesign Project\&quot;. | 
**FullName** | **string** | The case-insensitive fully-qualified unique name of this customer, formed by combining the names of its hierarchical parent objects with its own &#x60;name&#x60;, separated by colons. For example, if a customer is under \&quot;ABC Corporation\&quot; and has the &#x60;name&#x60; \&quot;Website Redesign Project\&quot;, its &#x60;fullName&#x60; would be \&quot;ABC Corporation:Website Redesign Project\&quot;.  **NOTE**: Unlike &#x60;name&#x60;, &#x60;fullName&#x60; is guaranteed to be unique across all customer objects. However, &#x60;fullName&#x60; can still be arbitrarily changed by the QuickBooks user when they modify the underlying &#x60;name&#x60; field.  **IMPORTANT**: If this object is a job (i.e., a sub-customer), this value would likely be the job&#39;s &#x60;name&#x60; prefixed by the customer&#39;s &#x60;name&#x60;. | 
**IsActive** | **bool** | Indicates whether this customer is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Class** | Pointer to [**QbdCustomerClass**](QbdCustomerClass.md) |  | [optional] 
**Parent** | Pointer to [**QbdCustomerParent**](QbdCustomerParent.md) |  | [optional] 
**Sublevel** | **float32** | The depth level of this customer in the hierarchy. A top-level customer has a &#x60;sublevel&#x60; of 0; each subsequent sublevel increases this number by 1. For example, a customer with a &#x60;fullName&#x60; of \&quot;ABC Corporation:Website Redesign Project\&quot; would have a &#x60;sublevel&#x60; of 1. When &#x60;sublevel&#x60; is 0, this object is a customer; when &#x60;sublevel&#x60; is greater than 0, this object is typically a job (i.e., a sub-customer). | 
**CompanyName** | Pointer to **string** | The name of the company associated with this customer. This name is used on invoices, checks, and other forms. | [optional] 
**Salutation** | Pointer to **string** | The formal salutation title that precedes the name of the contact person for this customer, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact person for this customer. | [optional] 
**MiddleName** | Pointer to **string** | The middle name of the contact person for this customer. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact person for this customer. | [optional] 
**JobTitle** | Pointer to **string** | The job title of the contact person for this customer. | [optional] 
**BillingAddress** | Pointer to [**QbdCustomerBillingAddress**](QbdCustomerBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QbdCustomerShippingAddress**](QbdCustomerShippingAddress.md) |  | [optional] 
**AlternateShippingAddresses** | Pointer to [**[]QbdShippingAddress**](QbdShippingAddress.md) | A list of additional shipping addresses for this customer. Useful when the customer has multiple shipping locations. | [optional] 
**Phone** | Pointer to **string** | The customer&#39;s primary telephone number. | [optional] 
**AlternatePhone** | Pointer to **string** | The customer&#39;s alternate telephone number. | [optional] 
**Fax** | Pointer to **string** | The customer&#39;s fax number. | [optional] 
**Email** | Pointer to **string** | The customer&#39;s email address. | [optional] 
**CcEmail** | Pointer to **string** | An email address to carbon copy (CC) on communications with this customer. | [optional] 
**Contact** | Pointer to **string** | The name of the primary contact person for this customer. | [optional] 
**AlternateContact** | Pointer to **string** | The name of a alternate contact person for this customer. | [optional] 
**CustomContactFields** | Pointer to [**[]QbdCustomContactField**](QbdCustomContactField.md) | Additional custom contact fields for this customer, such as phone numbers or email addresses. | [optional] 
**AdditionalContacts** | Pointer to [**[]QbdContact**](QbdContact.md) | Additional alternate contacts for this customer. | [optional] 
**CustomerType** | Pointer to [**QbdCustomerCustomerType**](QbdCustomerCustomerType.md) |  | [optional] 
**Terms** | Pointer to [**QbdCustomerTerms**](QbdCustomerTerms.md) |  | [optional] 
**SalesRepresentative** | Pointer to [**QbdCustomerSalesRepresentative**](QbdCustomerSalesRepresentative.md) |  | [optional] 
**Balance** | Pointer to **string** | The current balance owed by this customer, excluding balances from any jobs (i.e., sub-customers), represented as a decimal string. Compare with &#x60;totalBalance&#x60;. A positive number indicates money owed by the customer. | [optional] 
**TotalBalance** | Pointer to **string** | The combined balance of this customer and all of this customer&#39;s jobs (i.e., sub-customers), represented as a decimal string. If there are no sub-customers, &#x60;totalBalance&#x60; and &#x60;balance&#x60; are equal. A positive number indicates money owed by the customer. | [optional] 
**SalesTaxCode** | Pointer to [**QbdCustomerSalesTaxCode**](QbdCustomerSalesTaxCode.md) |  | [optional] 
**SalesTaxItem** | Pointer to [**QbdCustomerSalesTaxItem**](QbdCustomerSalesTaxItem.md) |  | [optional] 
**SalesTaxCountry** | Pointer to **string** | The country for which sales tax is collected for this customer. | [optional] 
**ResaleNumber** | Pointer to **string** | The customer&#39;s resale number, used if the customer is purchasing items for resale. This number does not affect sales tax calculations or reports in QuickBooks. | [optional] 
**AccountNumber** | Pointer to **string** | The customer&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**CreditLimit** | Pointer to **string** | The customer&#39;s credit limit, represented as a decimal string. This is the maximum amount of money this customer can spend before being billed. If &#x60;null&#x60;, there is no credit limit. | [optional] 
**PreferredPaymentMethod** | Pointer to [**QbdCustomerPreferredPaymentMethod**](QbdCustomerPreferredPaymentMethod.md) |  | [optional] 
**CreditCard** | Pointer to [**QbdCustomerCreditCard**](QbdCustomerCreditCard.md) |  | [optional] 
**JobStatus** | Pointer to **string** | The status of this customer&#39;s job, if this object is a job (i.e., sub-customer). | [optional] 
**JobStartDate** | Pointer to **string** | The date when work on this customer&#39;s job began, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobProjectedEndDate** | Pointer to **string** | The projected completion date for this customer&#39;s job, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobEndDate** | Pointer to **string** | The actual completion date of this customer&#39;s job, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobDescription** | Pointer to **string** | A brief description of this customer&#39;s job, if this object is a job (i.e., sub-customer). | [optional] 
**JobType** | Pointer to [**QbdCustomerJobType**](QbdCustomerJobType.md) |  | [optional] 
**Note** | Pointer to **string** | A note or comment about this customer. | [optional] 
**AdditionalNotes** | Pointer to [**[]QbdNote**](QbdNote.md) | Additional notes about this customer. | [optional] 
**PreferredDeliveryMethod** | Pointer to **string** | The preferred method for delivering invoices and other documents to this customer. | [optional] 
**PriceLevel** | Pointer to [**QbdCustomerPriceLevel**](QbdCustomerPriceLevel.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | [optional] 
**TaxRegistrationNumber** | Pointer to **string** | The customer&#39;s tax registration number, for use in Canada or the UK. | [optional] 
**Currency** | Pointer to [**QbdCustomerCurrency**](QbdCustomerCurrency.md) |  | [optional] 
**CustomFields** | Pointer to [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the customer object, added as user-defined data extensions, not included in the standard QuickBooks object. | [optional] 

## Methods

### NewQbdCustomer

`func NewQbdCustomer(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, fullName string, isActive bool, sublevel float32, ) *QbdCustomer`

NewQbdCustomer instantiates a new QbdCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCustomerWithDefaults

`func NewQbdCustomerWithDefaults() *QbdCustomer`

NewQbdCustomerWithDefaults instantiates a new QbdCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCustomer) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdCustomer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdCustomer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdCustomer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdCustomer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdCustomer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdCustomer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdCustomer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdCustomer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdCustomer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdCustomer) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdCustomer) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdCustomer) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdCustomer) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *QbdCustomer) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *QbdCustomer) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *QbdCustomer) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsActive

`func (o *QbdCustomer) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdCustomer) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdCustomer) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetClass

`func (o *QbdCustomer) GetClass() QbdCustomerClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdCustomer) GetClassOk() (*QbdCustomerClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdCustomer) SetClass(v QbdCustomerClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *QbdCustomer) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetParent

`func (o *QbdCustomer) GetParent() QbdCustomerParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *QbdCustomer) GetParentOk() (*QbdCustomerParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *QbdCustomer) SetParent(v QbdCustomerParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *QbdCustomer) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSublevel

`func (o *QbdCustomer) GetSublevel() float32`

GetSublevel returns the Sublevel field if non-nil, zero value otherwise.

### GetSublevelOk

`func (o *QbdCustomer) GetSublevelOk() (*float32, bool)`

GetSublevelOk returns a tuple with the Sublevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublevel

`func (o *QbdCustomer) SetSublevel(v float32)`

SetSublevel sets Sublevel field to given value.


### GetCompanyName

`func (o *QbdCustomer) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QbdCustomer) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QbdCustomer) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QbdCustomer) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSalutation

`func (o *QbdCustomer) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QbdCustomer) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QbdCustomer) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QbdCustomer) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QbdCustomer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QbdCustomer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QbdCustomer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QbdCustomer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QbdCustomer) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QbdCustomer) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QbdCustomer) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QbdCustomer) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QbdCustomer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QbdCustomer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QbdCustomer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QbdCustomer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QbdCustomer) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QbdCustomer) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QbdCustomer) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QbdCustomer) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QbdCustomer) GetBillingAddress() QbdCustomerBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdCustomer) GetBillingAddressOk() (*QbdCustomerBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdCustomer) SetBillingAddress(v QbdCustomerBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QbdCustomer) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QbdCustomer) GetShippingAddress() QbdCustomerShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdCustomer) GetShippingAddressOk() (*QbdCustomerShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdCustomer) SetShippingAddress(v QbdCustomerShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QbdCustomer) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetAlternateShippingAddresses

`func (o *QbdCustomer) GetAlternateShippingAddresses() []QbdShippingAddress`

GetAlternateShippingAddresses returns the AlternateShippingAddresses field if non-nil, zero value otherwise.

### GetAlternateShippingAddressesOk

`func (o *QbdCustomer) GetAlternateShippingAddressesOk() (*[]QbdShippingAddress, bool)`

GetAlternateShippingAddressesOk returns a tuple with the AlternateShippingAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateShippingAddresses

`func (o *QbdCustomer) SetAlternateShippingAddresses(v []QbdShippingAddress)`

SetAlternateShippingAddresses sets AlternateShippingAddresses field to given value.

### HasAlternateShippingAddresses

`func (o *QbdCustomer) HasAlternateShippingAddresses() bool`

HasAlternateShippingAddresses returns a boolean if a field has been set.

### GetPhone

`func (o *QbdCustomer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QbdCustomer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QbdCustomer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QbdCustomer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QbdCustomer) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QbdCustomer) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QbdCustomer) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QbdCustomer) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QbdCustomer) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QbdCustomer) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QbdCustomer) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QbdCustomer) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QbdCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QbdCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QbdCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QbdCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCcEmail

`func (o *QbdCustomer) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *QbdCustomer) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *QbdCustomer) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *QbdCustomer) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetContact

`func (o *QbdCustomer) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QbdCustomer) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QbdCustomer) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QbdCustomer) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAlternateContact

`func (o *QbdCustomer) GetAlternateContact() string`

GetAlternateContact returns the AlternateContact field if non-nil, zero value otherwise.

### GetAlternateContactOk

`func (o *QbdCustomer) GetAlternateContactOk() (*string, bool)`

GetAlternateContactOk returns a tuple with the AlternateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateContact

`func (o *QbdCustomer) SetAlternateContact(v string)`

SetAlternateContact sets AlternateContact field to given value.

### HasAlternateContact

`func (o *QbdCustomer) HasAlternateContact() bool`

HasAlternateContact returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QbdCustomer) GetCustomContactFields() []QbdCustomContactField`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QbdCustomer) GetCustomContactFieldsOk() (*[]QbdCustomContactField, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QbdCustomer) SetCustomContactFields(v []QbdCustomContactField)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QbdCustomer) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *QbdCustomer) GetAdditionalContacts() []QbdContact`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *QbdCustomer) GetAdditionalContactsOk() (*[]QbdContact, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *QbdCustomer) SetAdditionalContacts(v []QbdContact)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *QbdCustomer) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetCustomerType

`func (o *QbdCustomer) GetCustomerType() QbdCustomerCustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *QbdCustomer) GetCustomerTypeOk() (*QbdCustomerCustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *QbdCustomer) SetCustomerType(v QbdCustomerCustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *QbdCustomer) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetTerms

`func (o *QbdCustomer) GetTerms() QbdCustomerTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdCustomer) GetTermsOk() (*QbdCustomerTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdCustomer) SetTerms(v QbdCustomerTerms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *QbdCustomer) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetSalesRepresentative

`func (o *QbdCustomer) GetSalesRepresentative() QbdCustomerSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdCustomer) GetSalesRepresentativeOk() (*QbdCustomerSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdCustomer) SetSalesRepresentative(v QbdCustomerSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.

### HasSalesRepresentative

`func (o *QbdCustomer) HasSalesRepresentative() bool`

HasSalesRepresentative returns a boolean if a field has been set.

### GetBalance

`func (o *QbdCustomer) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *QbdCustomer) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *QbdCustomer) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *QbdCustomer) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetTotalBalance

`func (o *QbdCustomer) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *QbdCustomer) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *QbdCustomer) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *QbdCustomer) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetSalesTaxCode

`func (o *QbdCustomer) GetSalesTaxCode() QbdCustomerSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdCustomer) GetSalesTaxCodeOk() (*QbdCustomerSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdCustomer) SetSalesTaxCode(v QbdCustomerSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.

### HasSalesTaxCode

`func (o *QbdCustomer) HasSalesTaxCode() bool`

HasSalesTaxCode returns a boolean if a field has been set.

### GetSalesTaxItem

`func (o *QbdCustomer) GetSalesTaxItem() QbdCustomerSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdCustomer) GetSalesTaxItemOk() (*QbdCustomerSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdCustomer) SetSalesTaxItem(v QbdCustomerSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.

### HasSalesTaxItem

`func (o *QbdCustomer) HasSalesTaxItem() bool`

HasSalesTaxItem returns a boolean if a field has been set.

### GetSalesTaxCountry

`func (o *QbdCustomer) GetSalesTaxCountry() string`

GetSalesTaxCountry returns the SalesTaxCountry field if non-nil, zero value otherwise.

### GetSalesTaxCountryOk

`func (o *QbdCustomer) GetSalesTaxCountryOk() (*string, bool)`

GetSalesTaxCountryOk returns a tuple with the SalesTaxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCountry

`func (o *QbdCustomer) SetSalesTaxCountry(v string)`

SetSalesTaxCountry sets SalesTaxCountry field to given value.

### HasSalesTaxCountry

`func (o *QbdCustomer) HasSalesTaxCountry() bool`

HasSalesTaxCountry returns a boolean if a field has been set.

### GetResaleNumber

`func (o *QbdCustomer) GetResaleNumber() string`

GetResaleNumber returns the ResaleNumber field if non-nil, zero value otherwise.

### GetResaleNumberOk

`func (o *QbdCustomer) GetResaleNumberOk() (*string, bool)`

GetResaleNumberOk returns a tuple with the ResaleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResaleNumber

`func (o *QbdCustomer) SetResaleNumber(v string)`

SetResaleNumber sets ResaleNumber field to given value.

### HasResaleNumber

`func (o *QbdCustomer) HasResaleNumber() bool`

HasResaleNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QbdCustomer) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QbdCustomer) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QbdCustomer) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QbdCustomer) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCreditLimit

`func (o *QbdCustomer) GetCreditLimit() string`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *QbdCustomer) GetCreditLimitOk() (*string, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *QbdCustomer) SetCreditLimit(v string)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *QbdCustomer) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetPreferredPaymentMethod

`func (o *QbdCustomer) GetPreferredPaymentMethod() QbdCustomerPreferredPaymentMethod`

GetPreferredPaymentMethod returns the PreferredPaymentMethod field if non-nil, zero value otherwise.

### GetPreferredPaymentMethodOk

`func (o *QbdCustomer) GetPreferredPaymentMethodOk() (*QbdCustomerPreferredPaymentMethod, bool)`

GetPreferredPaymentMethodOk returns a tuple with the PreferredPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredPaymentMethod

`func (o *QbdCustomer) SetPreferredPaymentMethod(v QbdCustomerPreferredPaymentMethod)`

SetPreferredPaymentMethod sets PreferredPaymentMethod field to given value.

### HasPreferredPaymentMethod

`func (o *QbdCustomer) HasPreferredPaymentMethod() bool`

HasPreferredPaymentMethod returns a boolean if a field has been set.

### GetCreditCard

`func (o *QbdCustomer) GetCreditCard() QbdCustomerCreditCard`

GetCreditCard returns the CreditCard field if non-nil, zero value otherwise.

### GetCreditCardOk

`func (o *QbdCustomer) GetCreditCardOk() (*QbdCustomerCreditCard, bool)`

GetCreditCardOk returns a tuple with the CreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCard

`func (o *QbdCustomer) SetCreditCard(v QbdCustomerCreditCard)`

SetCreditCard sets CreditCard field to given value.

### HasCreditCard

`func (o *QbdCustomer) HasCreditCard() bool`

HasCreditCard returns a boolean if a field has been set.

### GetJobStatus

`func (o *QbdCustomer) GetJobStatus() string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *QbdCustomer) GetJobStatusOk() (*string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *QbdCustomer) SetJobStatus(v string)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *QbdCustomer) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetJobStartDate

`func (o *QbdCustomer) GetJobStartDate() string`

GetJobStartDate returns the JobStartDate field if non-nil, zero value otherwise.

### GetJobStartDateOk

`func (o *QbdCustomer) GetJobStartDateOk() (*string, bool)`

GetJobStartDateOk returns a tuple with the JobStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStartDate

`func (o *QbdCustomer) SetJobStartDate(v string)`

SetJobStartDate sets JobStartDate field to given value.

### HasJobStartDate

`func (o *QbdCustomer) HasJobStartDate() bool`

HasJobStartDate returns a boolean if a field has been set.

### GetJobProjectedEndDate

`func (o *QbdCustomer) GetJobProjectedEndDate() string`

GetJobProjectedEndDate returns the JobProjectedEndDate field if non-nil, zero value otherwise.

### GetJobProjectedEndDateOk

`func (o *QbdCustomer) GetJobProjectedEndDateOk() (*string, bool)`

GetJobProjectedEndDateOk returns a tuple with the JobProjectedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProjectedEndDate

`func (o *QbdCustomer) SetJobProjectedEndDate(v string)`

SetJobProjectedEndDate sets JobProjectedEndDate field to given value.

### HasJobProjectedEndDate

`func (o *QbdCustomer) HasJobProjectedEndDate() bool`

HasJobProjectedEndDate returns a boolean if a field has been set.

### GetJobEndDate

`func (o *QbdCustomer) GetJobEndDate() string`

GetJobEndDate returns the JobEndDate field if non-nil, zero value otherwise.

### GetJobEndDateOk

`func (o *QbdCustomer) GetJobEndDateOk() (*string, bool)`

GetJobEndDateOk returns a tuple with the JobEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobEndDate

`func (o *QbdCustomer) SetJobEndDate(v string)`

SetJobEndDate sets JobEndDate field to given value.

### HasJobEndDate

`func (o *QbdCustomer) HasJobEndDate() bool`

HasJobEndDate returns a boolean if a field has been set.

### GetJobDescription

`func (o *QbdCustomer) GetJobDescription() string`

GetJobDescription returns the JobDescription field if non-nil, zero value otherwise.

### GetJobDescriptionOk

`func (o *QbdCustomer) GetJobDescriptionOk() (*string, bool)`

GetJobDescriptionOk returns a tuple with the JobDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDescription

`func (o *QbdCustomer) SetJobDescription(v string)`

SetJobDescription sets JobDescription field to given value.

### HasJobDescription

`func (o *QbdCustomer) HasJobDescription() bool`

HasJobDescription returns a boolean if a field has been set.

### GetJobType

`func (o *QbdCustomer) GetJobType() QbdCustomerJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *QbdCustomer) GetJobTypeOk() (*QbdCustomerJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *QbdCustomer) SetJobType(v QbdCustomerJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *QbdCustomer) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetNote

`func (o *QbdCustomer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdCustomer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdCustomer) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QbdCustomer) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QbdCustomer) GetAdditionalNotes() []QbdNote`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QbdCustomer) GetAdditionalNotesOk() (*[]QbdNote, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QbdCustomer) SetAdditionalNotes(v []QbdNote)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QbdCustomer) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetPreferredDeliveryMethod

`func (o *QbdCustomer) GetPreferredDeliveryMethod() string`

GetPreferredDeliveryMethod returns the PreferredDeliveryMethod field if non-nil, zero value otherwise.

### GetPreferredDeliveryMethodOk

`func (o *QbdCustomer) GetPreferredDeliveryMethodOk() (*string, bool)`

GetPreferredDeliveryMethodOk returns a tuple with the PreferredDeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeliveryMethod

`func (o *QbdCustomer) SetPreferredDeliveryMethod(v string)`

SetPreferredDeliveryMethod sets PreferredDeliveryMethod field to given value.

### HasPreferredDeliveryMethod

`func (o *QbdCustomer) HasPreferredDeliveryMethod() bool`

HasPreferredDeliveryMethod returns a boolean if a field has been set.

### GetPriceLevel

`func (o *QbdCustomer) GetPriceLevel() QbdCustomerPriceLevel`

GetPriceLevel returns the PriceLevel field if non-nil, zero value otherwise.

### GetPriceLevelOk

`func (o *QbdCustomer) GetPriceLevelOk() (*QbdCustomerPriceLevel, bool)`

GetPriceLevelOk returns a tuple with the PriceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevel

`func (o *QbdCustomer) SetPriceLevel(v QbdCustomerPriceLevel)`

SetPriceLevel sets PriceLevel field to given value.

### HasPriceLevel

`func (o *QbdCustomer) HasPriceLevel() bool`

HasPriceLevel returns a boolean if a field has been set.

### GetExternalId

`func (o *QbdCustomer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdCustomer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdCustomer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QbdCustomer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *QbdCustomer) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *QbdCustomer) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *QbdCustomer) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *QbdCustomer) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *QbdCustomer) GetCurrency() QbdCustomerCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdCustomer) GetCurrencyOk() (*QbdCustomerCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdCustomer) SetCurrency(v QbdCustomerCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QbdCustomer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomFields

`func (o *QbdCustomer) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCustomer) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCustomer) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QbdCustomer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


