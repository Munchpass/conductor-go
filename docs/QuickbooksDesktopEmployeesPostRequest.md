# QuickbooksDesktopEmployeesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Indicates whether this employee is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**Salutation** | Pointer to **string** | The employee&#39;s formal salutation title that precedes their name, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | [optional] 
**FirstName** | Pointer to **string** | The employee&#39;s first name.  Maximum length: 25 characters. | [optional] 
**MiddleName** | Pointer to **string** | The employee&#39;s middle name.  Maximum length: 5 characters. | [optional] 
**LastName** | Pointer to **string** | The employee&#39;s last name.  Maximum length: 25 characters. | [optional] 
**JobTitle** | Pointer to **string** | The employee&#39;s job title. | [optional] 
**SupervisorId** | Pointer to **string** | The employee&#39;s supervisor. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | [optional] 
**Department** | Pointer to **string** | The employee&#39;s department. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | [optional] 
**Description** | Pointer to **string** | A description of this employee. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | [optional] 
**TargetBonus** | Pointer to **string** | The target bonus for this employee, represented as a decimal string. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | [optional] 
**Address** | Pointer to [**QuickbooksDesktopEmployeesPostRequestAddress**](QuickbooksDesktopEmployeesPostRequestAddress.md) |  | [optional] 
**PrintAs** | Pointer to **string** | The name to use when printing this employee from QuickBooks. By default, this is the same as the &#x60;name&#x60; field. | [optional] 
**Phone** | Pointer to **string** | The employee&#39;s primary telephone number. | [optional] 
**Mobile** | Pointer to **string** | The employee&#39;s mobile phone number. | [optional] 
**Pager** | Pointer to **string** | The employee&#39;s pager number. | [optional] 
**PagerPin** | Pointer to **string** | The employee&#39;s pager PIN. | [optional] 
**AlternatePhone** | Pointer to **string** | The employee&#39;s alternate telephone number. | [optional] 
**Fax** | Pointer to **string** | The employee&#39;s fax number. | [optional] 
**Ssn** | Pointer to **string** | The employee&#39;s Social Security Number. The value can be with or without dashes.  **NOTE**: This field cannot be changed after the employee is created. | [optional] 
**Email** | Pointer to **string** | The employee&#39;s email address. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this employee, such as phone numbers or email addresses. | [optional] 
**EmergencyContact** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmergencyContact**](QuickbooksDesktopEmployeesPostRequestEmergencyContact.md) |  | [optional] 
**EmployeeType** | Pointer to **string** | The employee type. This affects payroll taxes - a statutory employee is defined as an employee by statute. Note that owners/partners are typically on the \&quot;Other Names\&quot; list in QuickBooks, but if listed as an employee their type will be &#x60;owner&#x60;. | [optional] [default to "regular"]
**EmploymentStatus** | Pointer to **string** | Indicates whether this employee is a part-time or full-time employee. | [optional] 
**OvertimeExemptStatus** | Pointer to **string** | Indicates whether this employee is exempt from overtime pay. This classification is based on U.S. labor laws (FLSA).  | [optional] 
**KeyEmployeeStatus** | Pointer to **string** | Indicates whether this employee is a key employee. | [optional] 
**Gender** | Pointer to **string** | This employee&#39;s gender. | [optional] 
**HiredDate** | Pointer to **string** | The date this employee was hired, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**OriginalHireDate** | Pointer to **string** | The original hire date for this employee, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**AdjustedServiceDate** | Pointer to **string** | The adjusted service date for this employee, in ISO 8601 format (YYYY-MM-DD). This date accounts for previous employment periods or leaves that affect seniority. | [optional] 
**TerminationDate** | Pointer to **string** | The date this employee&#39;s employment ended with the company, in ISO 8601 format (YYYY-MM-DD). This is also known as the released date or separation date. | [optional] 
**BirthDate** | Pointer to **string** | This employee&#39;s date of birth, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**UsCitizenshipStatus** | Pointer to **string** | Indicates whether this employee is a U.S. citizen. | [optional] 
**Ethnicity** | Pointer to **string** | This employee&#39;s ethnicity. | [optional] 
**DisabilityStatus** | Pointer to **string** | Indicates whether this employee is disabled. | [optional] 
**DisabilityDescription** | Pointer to **string** | A description of this employee&#39;s disability. | [optional] 
**I9OnFileStatus** | Pointer to **string** | Indicates whether this employee&#39;s I-9 is on file. | [optional] 
**WorkAuthorizationExpirationDate** | Pointer to **string** | The date this employee&#39;s work authorization expires, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**UsVeteranStatus** | Pointer to **string** | Indicates whether this employee is a U.S. veteran. | [optional] 
**MilitaryStatus** | Pointer to **string** | This employee&#39;s military status if they are a U.S. veteran. | [optional] 
**AccountNumber** | Pointer to **string** | The employee&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | [optional] 
**Note** | Pointer to **string** | A note or comment about this employee. | [optional] 
**AdditionalNotes** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner**](QuickbooksDesktopCustomersPostRequestAdditionalNotesInner.md) | Additional notes about this employee. | [optional] 
**BillingRateId** | Pointer to **string** | The employee&#39;s billing rate, used to override service item rates in time tracking activities. | [optional] 
**EmployeePayroll** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmployeePayroll**](QuickbooksDesktopEmployeesPostRequestEmployeePayroll.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesPostRequest

`func NewQuickbooksDesktopEmployeesPostRequest() *QuickbooksDesktopEmployeesPostRequest`

NewQuickbooksDesktopEmployeesPostRequest instantiates a new QuickbooksDesktopEmployeesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesPostRequestWithDefaults

`func NewQuickbooksDesktopEmployeesPostRequestWithDefaults() *QuickbooksDesktopEmployeesPostRequest`

NewQuickbooksDesktopEmployeesPostRequestWithDefaults instantiates a new QuickbooksDesktopEmployeesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *QuickbooksDesktopEmployeesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopEmployeesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopEmployeesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSalutation

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopEmployeesPostRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopEmployeesPostRequest) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopEmployeesPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopEmployeesPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopEmployeesPostRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopEmployeesPostRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopEmployeesPostRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopEmployeesPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopEmployeesPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopEmployeesPostRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopEmployeesPostRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopEmployeesPostRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopEmployeesPostRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetSupervisorId

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSupervisorId() string`

GetSupervisorId returns the SupervisorId field if non-nil, zero value otherwise.

### GetSupervisorIdOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSupervisorIdOk() (*string, bool)`

GetSupervisorIdOk returns a tuple with the SupervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorId

`func (o *QuickbooksDesktopEmployeesPostRequest) SetSupervisorId(v string)`

SetSupervisorId sets SupervisorId field to given value.

### HasSupervisorId

`func (o *QuickbooksDesktopEmployeesPostRequest) HasSupervisorId() bool`

HasSupervisorId returns a boolean if a field has been set.

### GetDepartment

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *QuickbooksDesktopEmployeesPostRequest) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *QuickbooksDesktopEmployeesPostRequest) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetBonus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetTargetBonus() string`

GetTargetBonus returns the TargetBonus field if non-nil, zero value otherwise.

### GetTargetBonusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetTargetBonusOk() (*string, bool)`

GetTargetBonusOk returns a tuple with the TargetBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBonus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetTargetBonus(v string)`

SetTargetBonus sets TargetBonus field to given value.

### HasTargetBonus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasTargetBonus() bool`

HasTargetBonus returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAddress() QuickbooksDesktopEmployeesPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAddressOk() (*QuickbooksDesktopEmployeesPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopEmployeesPostRequest) SetAddress(v QuickbooksDesktopEmployeesPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopEmployeesPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrintAs

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPrintAs() string`

GetPrintAs returns the PrintAs field if non-nil, zero value otherwise.

### GetPrintAsOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPrintAsOk() (*string, bool)`

GetPrintAsOk returns a tuple with the PrintAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintAs

`func (o *QuickbooksDesktopEmployeesPostRequest) SetPrintAs(v string)`

SetPrintAs sets PrintAs field to given value.

### HasPrintAs

`func (o *QuickbooksDesktopEmployeesPostRequest) HasPrintAs() bool`

HasPrintAs returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopEmployeesPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopEmployeesPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMobile

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *QuickbooksDesktopEmployeesPostRequest) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *QuickbooksDesktopEmployeesPostRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPager

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *QuickbooksDesktopEmployeesPostRequest) SetPager(v string)`

SetPager sets Pager field to given value.

### HasPager

`func (o *QuickbooksDesktopEmployeesPostRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPagerPin

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPagerPin() string`

GetPagerPin returns the PagerPin field if non-nil, zero value otherwise.

### GetPagerPinOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetPagerPinOk() (*string, bool)`

GetPagerPinOk returns a tuple with the PagerPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerPin

`func (o *QuickbooksDesktopEmployeesPostRequest) SetPagerPin(v string)`

SetPagerPin sets PagerPin field to given value.

### HasPagerPin

`func (o *QuickbooksDesktopEmployeesPostRequest) HasPagerPin() bool`

HasPagerPin returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QuickbooksDesktopEmployeesPostRequest) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QuickbooksDesktopEmployeesPostRequest) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopEmployeesPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopEmployeesPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopEmployeesPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetSsn

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *QuickbooksDesktopEmployeesPostRequest) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *QuickbooksDesktopEmployeesPostRequest) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopEmployeesPostRequest) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopEmployeesPostRequest) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopEmployeesPostRequest) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetEmergencyContact

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmergencyContact() QuickbooksDesktopEmployeesPostRequestEmergencyContact`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmergencyContactOk() (*QuickbooksDesktopEmployeesPostRequestEmergencyContact, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEmergencyContact(v QuickbooksDesktopEmployeesPostRequestEmergencyContact)`

SetEmergencyContact sets EmergencyContact field to given value.

### HasEmergencyContact

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEmergencyContact() bool`

HasEmergencyContact returns a boolean if a field has been set.

### GetEmployeeType

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetEmploymentStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### GetOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetOvertimeExemptStatus() string`

GetOvertimeExemptStatus returns the OvertimeExemptStatus field if non-nil, zero value otherwise.

### GetOvertimeExemptStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetOvertimeExemptStatusOk() (*string, bool)`

GetOvertimeExemptStatusOk returns a tuple with the OvertimeExemptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetOvertimeExemptStatus(v string)`

SetOvertimeExemptStatus sets OvertimeExemptStatus field to given value.

### HasOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasOvertimeExemptStatus() bool`

HasOvertimeExemptStatus returns a boolean if a field has been set.

### GetKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetKeyEmployeeStatus() string`

GetKeyEmployeeStatus returns the KeyEmployeeStatus field if non-nil, zero value otherwise.

### GetKeyEmployeeStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetKeyEmployeeStatusOk() (*string, bool)`

GetKeyEmployeeStatusOk returns a tuple with the KeyEmployeeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetKeyEmployeeStatus(v string)`

SetKeyEmployeeStatus sets KeyEmployeeStatus field to given value.

### HasKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasKeyEmployeeStatus() bool`

HasKeyEmployeeStatus returns a boolean if a field has been set.

### GetGender

`func (o *QuickbooksDesktopEmployeesPostRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *QuickbooksDesktopEmployeesPostRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *QuickbooksDesktopEmployeesPostRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHiredDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetOriginalHireDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetOriginalHireDate() string`

GetOriginalHireDate returns the OriginalHireDate field if non-nil, zero value otherwise.

### GetOriginalHireDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetOriginalHireDateOk() (*string, bool)`

GetOriginalHireDateOk returns a tuple with the OriginalHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalHireDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetOriginalHireDate(v string)`

SetOriginalHireDate sets OriginalHireDate field to given value.

### HasOriginalHireDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasOriginalHireDate() bool`

HasOriginalHireDate returns a boolean if a field has been set.

### GetAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAdjustedServiceDate() string`

GetAdjustedServiceDate returns the AdjustedServiceDate field if non-nil, zero value otherwise.

### GetAdjustedServiceDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAdjustedServiceDateOk() (*string, bool)`

GetAdjustedServiceDateOk returns a tuple with the AdjustedServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetAdjustedServiceDate(v string)`

SetAdjustedServiceDate sets AdjustedServiceDate field to given value.

### HasAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasAdjustedServiceDate() bool`

HasAdjustedServiceDate returns a boolean if a field has been set.

### GetTerminationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetUsCitizenshipStatus() string`

GetUsCitizenshipStatus returns the UsCitizenshipStatus field if non-nil, zero value otherwise.

### GetUsCitizenshipStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetUsCitizenshipStatusOk() (*string, bool)`

GetUsCitizenshipStatusOk returns a tuple with the UsCitizenshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetUsCitizenshipStatus(v string)`

SetUsCitizenshipStatus sets UsCitizenshipStatus field to given value.

### HasUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasUsCitizenshipStatus() bool`

HasUsCitizenshipStatus returns a boolean if a field has been set.

### GetEthnicity

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.

### HasEthnicity

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### GetDisabilityStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDisabilityStatus() string`

GetDisabilityStatus returns the DisabilityStatus field if non-nil, zero value otherwise.

### GetDisabilityStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDisabilityStatusOk() (*string, bool)`

GetDisabilityStatusOk returns a tuple with the DisabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetDisabilityStatus(v string)`

SetDisabilityStatus sets DisabilityStatus field to given value.

### HasDisabilityStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasDisabilityStatus() bool`

HasDisabilityStatus returns a boolean if a field has been set.

### GetDisabilityDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDisabilityDescription() string`

GetDisabilityDescription returns the DisabilityDescription field if non-nil, zero value otherwise.

### GetDisabilityDescriptionOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetDisabilityDescriptionOk() (*string, bool)`

GetDisabilityDescriptionOk returns a tuple with the DisabilityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) SetDisabilityDescription(v string)`

SetDisabilityDescription sets DisabilityDescription field to given value.

### HasDisabilityDescription

`func (o *QuickbooksDesktopEmployeesPostRequest) HasDisabilityDescription() bool`

HasDisabilityDescription returns a boolean if a field has been set.

### GetI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetI9OnFileStatus() string`

GetI9OnFileStatus returns the I9OnFileStatus field if non-nil, zero value otherwise.

### GetI9OnFileStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetI9OnFileStatusOk() (*string, bool)`

GetI9OnFileStatusOk returns a tuple with the I9OnFileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetI9OnFileStatus(v string)`

SetI9OnFileStatus sets I9OnFileStatus field to given value.

### HasI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasI9OnFileStatus() bool`

HasI9OnFileStatus returns a boolean if a field has been set.

### GetWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) GetWorkAuthorizationExpirationDate() string`

GetWorkAuthorizationExpirationDate returns the WorkAuthorizationExpirationDate field if non-nil, zero value otherwise.

### GetWorkAuthorizationExpirationDateOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetWorkAuthorizationExpirationDateOk() (*string, bool)`

GetWorkAuthorizationExpirationDateOk returns a tuple with the WorkAuthorizationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) SetWorkAuthorizationExpirationDate(v string)`

SetWorkAuthorizationExpirationDate sets WorkAuthorizationExpirationDate field to given value.

### HasWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesPostRequest) HasWorkAuthorizationExpirationDate() bool`

HasWorkAuthorizationExpirationDate returns a boolean if a field has been set.

### GetUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetUsVeteranStatus() string`

GetUsVeteranStatus returns the UsVeteranStatus field if non-nil, zero value otherwise.

### GetUsVeteranStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetUsVeteranStatusOk() (*string, bool)`

GetUsVeteranStatusOk returns a tuple with the UsVeteranStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetUsVeteranStatus(v string)`

SetUsVeteranStatus sets UsVeteranStatus field to given value.

### HasUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasUsVeteranStatus() bool`

HasUsVeteranStatus returns a boolean if a field has been set.

### GetMilitaryStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMilitaryStatus() string`

GetMilitaryStatus returns the MilitaryStatus field if non-nil, zero value otherwise.

### GetMilitaryStatusOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetMilitaryStatusOk() (*string, bool)`

GetMilitaryStatusOk returns a tuple with the MilitaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilitaryStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) SetMilitaryStatus(v string)`

SetMilitaryStatus sets MilitaryStatus field to given value.

### HasMilitaryStatus

`func (o *QuickbooksDesktopEmployeesPostRequest) HasMilitaryStatus() bool`

HasMilitaryStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopEmployeesPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopEmployeesPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopEmployeesPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopEmployeesPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopEmployeesPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAdditionalNotes() []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetAdditionalNotesOk() (*[]QuickbooksDesktopCustomersPostRequestAdditionalNotesInner, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QuickbooksDesktopEmployeesPostRequest) SetAdditionalNotes(v []QuickbooksDesktopCustomersPostRequestAdditionalNotesInner)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QuickbooksDesktopEmployeesPostRequest) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetBillingRateId

`func (o *QuickbooksDesktopEmployeesPostRequest) GetBillingRateId() string`

GetBillingRateId returns the BillingRateId field if non-nil, zero value otherwise.

### GetBillingRateIdOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetBillingRateIdOk() (*string, bool)`

GetBillingRateIdOk returns a tuple with the BillingRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRateId

`func (o *QuickbooksDesktopEmployeesPostRequest) SetBillingRateId(v string)`

SetBillingRateId sets BillingRateId field to given value.

### HasBillingRateId

`func (o *QuickbooksDesktopEmployeesPostRequest) HasBillingRateId() bool`

HasBillingRateId returns a boolean if a field has been set.

### GetEmployeePayroll

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmployeePayroll() QuickbooksDesktopEmployeesPostRequestEmployeePayroll`

GetEmployeePayroll returns the EmployeePayroll field if non-nil, zero value otherwise.

### GetEmployeePayrollOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetEmployeePayrollOk() (*QuickbooksDesktopEmployeesPostRequestEmployeePayroll, bool)`

GetEmployeePayrollOk returns a tuple with the EmployeePayroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayroll

`func (o *QuickbooksDesktopEmployeesPostRequest) SetEmployeePayroll(v QuickbooksDesktopEmployeesPostRequestEmployeePayroll)`

SetEmployeePayroll sets EmployeePayroll field to given value.

### HasEmployeePayroll

`func (o *QuickbooksDesktopEmployeesPostRequest) HasEmployeePayroll() bool`

HasEmployeePayroll returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopEmployeesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopEmployeesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopEmployeesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopEmployeesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


