# QuickbooksDesktopCustomersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this customer. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two customers could both have the &#x60;name&#x60; \&quot;Website Redesign Project\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;ABC Corporation:Website Redesign Project\&quot; and \&quot;Baker:Website Redesign Project\&quot;.  Maximum length: 41 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this customer is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The customer&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent customer one level above this one in the hierarchy. For example, if this customer has a &#x60;fullName&#x60; of \&quot;ABC Corporation:Website Redesign Project\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;ABC Corporation\&quot;. If this customer is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company associated with this customer. This name is used on invoices, checks, and other forms. | [optional] 
**Salutation** | Pointer to **string** | The formal salutation title that precedes the name of the contact person for this customer, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The first name of the contact person for this customer.  Maximum length: 25 characters. | [optional] 
**MiddleName** | Pointer to **string** | The middle name of the contact person for this customer.  Maximum length: 5 characters. | [optional] 
**LastName** | Pointer to **string** | The last name of the contact person for this customer.  Maximum length: 25 characters. | [optional] 
**JobTitle** | Pointer to **string** | The job title of the contact person for this customer. | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopCustomersPostRequestBillingAddress**](QuickbooksDesktopCustomersPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopCustomersPostRequestShippingAddress**](QuickbooksDesktopCustomersPostRequestShippingAddress.md) |  | [optional] 
**AlternateShippingAddresses** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner**](QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner.md) | A list of additional shipping addresses for this customer. Useful when the customer has multiple shipping locations. | [optional] 
**Phone** | Pointer to **string** | The customer&#39;s primary telephone number. | [optional] 
**AlternatePhone** | Pointer to **string** | The customer&#39;s alternate telephone number. | [optional] 
**Fax** | Pointer to **string** | The customer&#39;s fax number. | [optional] 
**Email** | Pointer to **string** | The customer&#39;s email address. | [optional] 
**CcEmail** | Pointer to **string** | An email address to carbon copy (CC) on communications with this customer. | [optional] 
**Contact** | Pointer to **string** | The name of the primary contact person for this customer. | [optional] 
**AlternateContact** | Pointer to **string** | The name of a alternate contact person for this customer. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this customer, such as phone numbers or email addresses. | [optional] 
**AdditionalContacts** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAdditionalContactsInner**](QuickbooksDesktopCustomersPostRequestAdditionalContactsInner.md) | Additional alternate contacts for this customer. | [optional] 
**CustomerTypeId** | Pointer to **string** | The customer&#39;s type, used for categorizing customers into meaningful segments, such as industry or region. | [optional] 
**TermsId** | Pointer to **string** | The customer&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The customer&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**OpeningBalance** | Pointer to **string** | The opening balance of this customer&#39;s account, indicating the amount owed by this customer, represented as a decimal string. | [optional] 
**OpeningBalanceDate** | Pointer to **string** | The date of the opening balance of this customer, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for transactions with this customer, determining whether the transactions are taxable or non-taxable. This can be overridden at the transaction or transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this customer&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**SalesTaxCountry** | Pointer to **string** | The country for which sales tax is collected for this customer. | [optional] 
**ResaleNumber** | Pointer to **string** | The customer&#39;s resale number, used if the customer is purchasing items for resale. This number does not affect sales tax calculations or reports in QuickBooks. | [optional] 
**AccountNumber** | Pointer to **string** | The customer&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**CreditLimit** | Pointer to **string** | The customer&#39;s credit limit, represented as a decimal string. This is the maximum amount of money this customer can spend before being billed. If &#x60;null&#x60;, there is no credit limit. | [optional] 
**PreferredPaymentMethodId** | Pointer to **string** | The customer&#39;s preferred payment method (e.g., cash, check, credit card). | [optional] 
**CreditCard** | Pointer to [**QuickbooksDesktopCustomersPostRequestCreditCard**](QuickbooksDesktopCustomersPostRequestCreditCard.md) |  | [optional] 
**JobStatus** | Pointer to **string** | The status of this customer&#39;s job, if this object is a job (i.e., sub-customer). | [optional] [default to "none"]
**JobStartDate** | Pointer to **string** | The date when work on this customer&#39;s job began, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobProjectedEndDate** | Pointer to **string** | The projected completion date for this customer&#39;s job, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobEndDate** | Pointer to **string** | The actual completion date of this customer&#39;s job, if applicable, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**JobDescription** | Pointer to **string** | A brief description of this customer&#39;s job, if this object is a job (i.e., sub-customer). | [optional] 
**JobTypeId** | Pointer to **string** | The type or category of this customer&#39;s job, if this object is a job (i.e., sub-customer). Useful for classifying into meaningful segments (e.g., repair, installation, consulting). | [optional] 
**Note** | Pointer to **string** | A note or comment about this customer. | [optional] 
**AdditionalNotes** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner**](QuickbooksDesktopCustomersPostRequestAdditionalNotesInner.md) | Additional notes about this customer. | [optional] 
**PreferredDeliveryMethod** | Pointer to **string** | The preferred method for delivering invoices and other documents to this customer. | [optional] [default to "none"]
**PriceLevelId** | Pointer to **string** | The customer&#39;s custom price level that QuickBooks automatically applies to calculate item rates in new transactions (e.g., invoices, sales receipts, sales orders, and credit memos) for this customer. While applied automatically, this can be overridden when creating individual transactions. Note that transactions will not show the price level itself, only the final &#x60;rate&#x60; calculated from it. | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**TaxRegistrationNumber** | Pointer to **string** | The customer&#39;s tax registration number, for use in Canada or the UK. | [optional] 
**CurrencyId** | Pointer to **string** | The customer&#39;s currency. For built-in currencies, the name and code are standard international values. For user-defined currencies, all values are editable. | [optional] 

## Methods

### NewQuickbooksDesktopCustomersPostRequest

`func NewQuickbooksDesktopCustomersPostRequest(name string, ) *QuickbooksDesktopCustomersPostRequest`

NewQuickbooksDesktopCustomersPostRequest instantiates a new QuickbooksDesktopCustomersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCustomersPostRequestWithDefaults

`func NewQuickbooksDesktopCustomersPostRequestWithDefaults() *QuickbooksDesktopCustomersPostRequest`

NewQuickbooksDesktopCustomersPostRequestWithDefaults instantiates a new QuickbooksDesktopCustomersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopCustomersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopCustomersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopCustomersPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopCustomersPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopCustomersPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopCustomersPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopCustomersPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopCustomersPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopCustomersPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopCustomersPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopCustomersPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetCompanyName

`func (o *QuickbooksDesktopCustomersPostRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuickbooksDesktopCustomersPostRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuickbooksDesktopCustomersPostRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSalutation

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopCustomersPostRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopCustomersPostRequest) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopCustomersPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopCustomersPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopCustomersPostRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopCustomersPostRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopCustomersPostRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopCustomersPostRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopCustomersPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopCustomersPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopCustomersPostRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) GetBillingAddress() QuickbooksDesktopCustomersPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetBillingAddressOk() (*QuickbooksDesktopCustomersPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) SetBillingAddress(v QuickbooksDesktopCustomersPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) GetShippingAddress() QuickbooksDesktopCustomersPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetShippingAddressOk() (*QuickbooksDesktopCustomersPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) SetShippingAddress(v QuickbooksDesktopCustomersPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopCustomersPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetAlternateShippingAddresses

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternateShippingAddresses() []QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner`

GetAlternateShippingAddresses returns the AlternateShippingAddresses field if non-nil, zero value otherwise.

### GetAlternateShippingAddressesOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternateShippingAddressesOk() (*[]QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner, bool)`

GetAlternateShippingAddressesOk returns a tuple with the AlternateShippingAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateShippingAddresses

`func (o *QuickbooksDesktopCustomersPostRequest) SetAlternateShippingAddresses(v []QuickbooksDesktopCustomersPostRequestAlternateShippingAddressesInner)`

SetAlternateShippingAddresses sets AlternateShippingAddresses field to given value.

### HasAlternateShippingAddresses

`func (o *QuickbooksDesktopCustomersPostRequest) HasAlternateShippingAddresses() bool`

HasAlternateShippingAddresses returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopCustomersPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopCustomersPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopCustomersPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QuickbooksDesktopCustomersPostRequest) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QuickbooksDesktopCustomersPostRequest) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopCustomersPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopCustomersPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopCustomersPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopCustomersPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopCustomersPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopCustomersPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCcEmail

`func (o *QuickbooksDesktopCustomersPostRequest) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *QuickbooksDesktopCustomersPostRequest) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *QuickbooksDesktopCustomersPostRequest) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetContact

`func (o *QuickbooksDesktopCustomersPostRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuickbooksDesktopCustomersPostRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuickbooksDesktopCustomersPostRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAlternateContact

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternateContact() string`

GetAlternateContact returns the AlternateContact field if non-nil, zero value otherwise.

### GetAlternateContactOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAlternateContactOk() (*string, bool)`

GetAlternateContactOk returns a tuple with the AlternateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateContact

`func (o *QuickbooksDesktopCustomersPostRequest) SetAlternateContact(v string)`

SetAlternateContact sets AlternateContact field to given value.

### HasAlternateContact

`func (o *QuickbooksDesktopCustomersPostRequest) HasAlternateContact() bool`

HasAlternateContact returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequest) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequest) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopCustomersPostRequest) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *QuickbooksDesktopCustomersPostRequest) GetAdditionalContacts() []QuickbooksDesktopCustomersPostRequestAdditionalContactsInner`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAdditionalContactsOk() (*[]QuickbooksDesktopCustomersPostRequestAdditionalContactsInner, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *QuickbooksDesktopCustomersPostRequest) SetAdditionalContacts(v []QuickbooksDesktopCustomersPostRequestAdditionalContactsInner)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *QuickbooksDesktopCustomersPostRequest) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetCustomerTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) GetCustomerTypeId() string`

GetCustomerTypeId returns the CustomerTypeId field if non-nil, zero value otherwise.

### GetCustomerTypeIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCustomerTypeIdOk() (*string, bool)`

GetCustomerTypeIdOk returns a tuple with the CustomerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) SetCustomerTypeId(v string)`

SetCustomerTypeId sets CustomerTypeId field to given value.

### HasCustomerTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) HasCustomerTypeId() bool`

HasCustomerTypeId returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopCustomersPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopCustomersPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopCustomersPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopCustomersPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopCustomersPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *QuickbooksDesktopCustomersPostRequest) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *QuickbooksDesktopCustomersPostRequest) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *QuickbooksDesktopCustomersPostRequest) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *QuickbooksDesktopCustomersPostRequest) GetOpeningBalanceDate() string`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetOpeningBalanceDateOk() (*string, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *QuickbooksDesktopCustomersPostRequest) SetOpeningBalanceDate(v string)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *QuickbooksDesktopCustomersPostRequest) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopCustomersPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopCustomersPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopCustomersPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopCustomersPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetSalesTaxCountry

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxCountry() string`

GetSalesTaxCountry returns the SalesTaxCountry field if non-nil, zero value otherwise.

### GetSalesTaxCountryOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetSalesTaxCountryOk() (*string, bool)`

GetSalesTaxCountryOk returns a tuple with the SalesTaxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCountry

`func (o *QuickbooksDesktopCustomersPostRequest) SetSalesTaxCountry(v string)`

SetSalesTaxCountry sets SalesTaxCountry field to given value.

### HasSalesTaxCountry

`func (o *QuickbooksDesktopCustomersPostRequest) HasSalesTaxCountry() bool`

HasSalesTaxCountry returns a boolean if a field has been set.

### GetResaleNumber

`func (o *QuickbooksDesktopCustomersPostRequest) GetResaleNumber() string`

GetResaleNumber returns the ResaleNumber field if non-nil, zero value otherwise.

### GetResaleNumberOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetResaleNumberOk() (*string, bool)`

GetResaleNumberOk returns a tuple with the ResaleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResaleNumber

`func (o *QuickbooksDesktopCustomersPostRequest) SetResaleNumber(v string)`

SetResaleNumber sets ResaleNumber field to given value.

### HasResaleNumber

`func (o *QuickbooksDesktopCustomersPostRequest) HasResaleNumber() bool`

HasResaleNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QuickbooksDesktopCustomersPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopCustomersPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopCustomersPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCreditLimit

`func (o *QuickbooksDesktopCustomersPostRequest) GetCreditLimit() string`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCreditLimitOk() (*string, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *QuickbooksDesktopCustomersPostRequest) SetCreditLimit(v string)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *QuickbooksDesktopCustomersPostRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetPreferredPaymentMethodId

`func (o *QuickbooksDesktopCustomersPostRequest) GetPreferredPaymentMethodId() string`

GetPreferredPaymentMethodId returns the PreferredPaymentMethodId field if non-nil, zero value otherwise.

### GetPreferredPaymentMethodIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetPreferredPaymentMethodIdOk() (*string, bool)`

GetPreferredPaymentMethodIdOk returns a tuple with the PreferredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredPaymentMethodId

`func (o *QuickbooksDesktopCustomersPostRequest) SetPreferredPaymentMethodId(v string)`

SetPreferredPaymentMethodId sets PreferredPaymentMethodId field to given value.

### HasPreferredPaymentMethodId

`func (o *QuickbooksDesktopCustomersPostRequest) HasPreferredPaymentMethodId() bool`

HasPreferredPaymentMethodId returns a boolean if a field has been set.

### GetCreditCard

`func (o *QuickbooksDesktopCustomersPostRequest) GetCreditCard() QuickbooksDesktopCustomersPostRequestCreditCard`

GetCreditCard returns the CreditCard field if non-nil, zero value otherwise.

### GetCreditCardOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCreditCardOk() (*QuickbooksDesktopCustomersPostRequestCreditCard, bool)`

GetCreditCardOk returns a tuple with the CreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCard

`func (o *QuickbooksDesktopCustomersPostRequest) SetCreditCard(v QuickbooksDesktopCustomersPostRequestCreditCard)`

SetCreditCard sets CreditCard field to given value.

### HasCreditCard

`func (o *QuickbooksDesktopCustomersPostRequest) HasCreditCard() bool`

HasCreditCard returns a boolean if a field has been set.

### GetJobStatus

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobStatus() string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobStatusOk() (*string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobStatus(v string)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetJobStartDate

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobStartDate() string`

GetJobStartDate returns the JobStartDate field if non-nil, zero value otherwise.

### GetJobStartDateOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobStartDateOk() (*string, bool)`

GetJobStartDateOk returns a tuple with the JobStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStartDate

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobStartDate(v string)`

SetJobStartDate sets JobStartDate field to given value.

### HasJobStartDate

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobStartDate() bool`

HasJobStartDate returns a boolean if a field has been set.

### GetJobProjectedEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobProjectedEndDate() string`

GetJobProjectedEndDate returns the JobProjectedEndDate field if non-nil, zero value otherwise.

### GetJobProjectedEndDateOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobProjectedEndDateOk() (*string, bool)`

GetJobProjectedEndDateOk returns a tuple with the JobProjectedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProjectedEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobProjectedEndDate(v string)`

SetJobProjectedEndDate sets JobProjectedEndDate field to given value.

### HasJobProjectedEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobProjectedEndDate() bool`

HasJobProjectedEndDate returns a boolean if a field has been set.

### GetJobEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobEndDate() string`

GetJobEndDate returns the JobEndDate field if non-nil, zero value otherwise.

### GetJobEndDateOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobEndDateOk() (*string, bool)`

GetJobEndDateOk returns a tuple with the JobEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobEndDate(v string)`

SetJobEndDate sets JobEndDate field to given value.

### HasJobEndDate

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobEndDate() bool`

HasJobEndDate returns a boolean if a field has been set.

### GetJobDescription

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobDescription() string`

GetJobDescription returns the JobDescription field if non-nil, zero value otherwise.

### GetJobDescriptionOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobDescriptionOk() (*string, bool)`

GetJobDescriptionOk returns a tuple with the JobDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDescription

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobDescription(v string)`

SetJobDescription sets JobDescription field to given value.

### HasJobDescription

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobDescription() bool`

HasJobDescription returns a boolean if a field has been set.

### GetJobTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobTypeId() string`

GetJobTypeId returns the JobTypeId field if non-nil, zero value otherwise.

### GetJobTypeIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetJobTypeIdOk() (*string, bool)`

GetJobTypeIdOk returns a tuple with the JobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) SetJobTypeId(v string)`

SetJobTypeId sets JobTypeId field to given value.

### HasJobTypeId

`func (o *QuickbooksDesktopCustomersPostRequest) HasJobTypeId() bool`

HasJobTypeId returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopCustomersPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopCustomersPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopCustomersPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QuickbooksDesktopCustomersPostRequest) GetAdditionalNotes() []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetAdditionalNotesOk() (*[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QuickbooksDesktopCustomersPostRequest) SetAdditionalNotes(v []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QuickbooksDesktopCustomersPostRequest) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetPreferredDeliveryMethod

`func (o *QuickbooksDesktopCustomersPostRequest) GetPreferredDeliveryMethod() string`

GetPreferredDeliveryMethod returns the PreferredDeliveryMethod field if non-nil, zero value otherwise.

### GetPreferredDeliveryMethodOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetPreferredDeliveryMethodOk() (*string, bool)`

GetPreferredDeliveryMethodOk returns a tuple with the PreferredDeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDeliveryMethod

`func (o *QuickbooksDesktopCustomersPostRequest) SetPreferredDeliveryMethod(v string)`

SetPreferredDeliveryMethod sets PreferredDeliveryMethod field to given value.

### HasPreferredDeliveryMethod

`func (o *QuickbooksDesktopCustomersPostRequest) HasPreferredDeliveryMethod() bool`

HasPreferredDeliveryMethod returns a boolean if a field has been set.

### GetPriceLevelId

`func (o *QuickbooksDesktopCustomersPostRequest) GetPriceLevelId() string`

GetPriceLevelId returns the PriceLevelId field if non-nil, zero value otherwise.

### GetPriceLevelIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetPriceLevelIdOk() (*string, bool)`

GetPriceLevelIdOk returns a tuple with the PriceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevelId

`func (o *QuickbooksDesktopCustomersPostRequest) SetPriceLevelId(v string)`

SetPriceLevelId sets PriceLevelId field to given value.

### HasPriceLevelId

`func (o *QuickbooksDesktopCustomersPostRequest) HasPriceLevelId() bool`

HasPriceLevelId returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopCustomersPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopCustomersPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopCustomersPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *QuickbooksDesktopCustomersPostRequest) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *QuickbooksDesktopCustomersPostRequest) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *QuickbooksDesktopCustomersPostRequest) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetCurrencyId

`func (o *QuickbooksDesktopCustomersPostRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuickbooksDesktopCustomersPostRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuickbooksDesktopCustomersPostRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuickbooksDesktopCustomersPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


