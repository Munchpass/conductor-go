# QuickbooksDesktopEmployeesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the employee object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**IsActive** | Pointer to **bool** | Indicates whether this employee is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
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
**Email** | Pointer to **string** | The employee&#39;s email address. | [optional] 
**CustomContactFields** | Pointer to [**[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner**](QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner.md) | Additional custom contact fields for this employee, such as phone numbers or email addresses. | [optional] 
**EmergencyContact** | Pointer to [**QuickbooksDesktopEmployeesPostRequestEmergencyContact**](QuickbooksDesktopEmployeesPostRequestEmergencyContact.md) |  | [optional] 
**EmployeeType** | Pointer to **string** | The employee type. This affects payroll taxes - a statutory employee is defined as an employee by statute. Note that owners/partners are typically on the \&quot;Other Names\&quot; list in QuickBooks, but if listed as an employee their type will be &#x60;owner&#x60;. | [optional] 
**EmploymentStatus** | Pointer to **string** | Indicates whether this employee is a part-time or full-time employee. | [optional] 
**OvertimeExemptStatus** | Pointer to **string** | Indicates whether this employee is exempt from overtime pay. This classification is based on U.S. labor laws (FLSA).  | [optional] 
**KeyEmployeeStatus** | Pointer to **string** | Indicates whether this employee is a key employee. | [optional] 
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
**AdditionalNotes** | Pointer to [**[]QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner**](QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner.md) | Additional notes about this employee. | [optional] 
**BillingRateId** | Pointer to **string** | The employee&#39;s billing rate, used to override service item rates in time tracking activities. | [optional] 
**EmployeePayroll** | Pointer to [**QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll**](QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopEmployeesIdPostRequest

`func NewQuickbooksDesktopEmployeesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopEmployeesIdPostRequest`

NewQuickbooksDesktopEmployeesIdPostRequest instantiates a new QuickbooksDesktopEmployeesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEmployeesIdPostRequestWithDefaults

`func NewQuickbooksDesktopEmployeesIdPostRequestWithDefaults() *QuickbooksDesktopEmployeesIdPostRequest`

NewQuickbooksDesktopEmployeesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopEmployeesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetIsActive

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSalutation

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.

### HasSalutation

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasSalutation() bool`

HasSalutation returns a boolean if a field has been set.

### GetFirstName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetJobTitle

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetSupervisorId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetSupervisorId() string`

GetSupervisorId returns the SupervisorId field if non-nil, zero value otherwise.

### GetSupervisorIdOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetSupervisorIdOk() (*string, bool)`

GetSupervisorIdOk returns a tuple with the SupervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetSupervisorId(v string)`

SetSupervisorId sets SupervisorId field to given value.

### HasSupervisorId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasSupervisorId() bool`

HasSupervisorId returns a boolean if a field has been set.

### GetDepartment

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetBonus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetTargetBonus() string`

GetTargetBonus returns the TargetBonus field if non-nil, zero value otherwise.

### GetTargetBonusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetTargetBonusOk() (*string, bool)`

GetTargetBonusOk returns a tuple with the TargetBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBonus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetTargetBonus(v string)`

SetTargetBonus sets TargetBonus field to given value.

### HasTargetBonus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasTargetBonus() bool`

HasTargetBonus returns a boolean if a field has been set.

### GetAddress

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAddress() QuickbooksDesktopEmployeesPostRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAddressOk() (*QuickbooksDesktopEmployeesPostRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetAddress(v QuickbooksDesktopEmployeesPostRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrintAs

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPrintAs() string`

GetPrintAs returns the PrintAs field if non-nil, zero value otherwise.

### GetPrintAsOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPrintAsOk() (*string, bool)`

GetPrintAsOk returns a tuple with the PrintAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintAs

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetPrintAs(v string)`

SetPrintAs sets PrintAs field to given value.

### HasPrintAs

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasPrintAs() bool`

HasPrintAs returns a boolean if a field has been set.

### GetPhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMobile

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPager

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetPager(v string)`

SetPager sets Pager field to given value.

### HasPager

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPagerPin

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPagerPin() string`

GetPagerPin returns the PagerPin field if non-nil, zero value otherwise.

### GetPagerPinOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetPagerPinOk() (*string, bool)`

GetPagerPinOk returns a tuple with the PagerPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerPin

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetPagerPin(v string)`

SetPagerPin sets PagerPin field to given value.

### HasPagerPin

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasPagerPin() bool`

HasPagerPin returns a boolean if a field has been set.

### GetAlternatePhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.

### HasAlternatePhone

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasAlternatePhone() bool`

HasAlternatePhone returns a boolean if a field has been set.

### GetFax

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCustomContactFields

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetCustomContactFields() []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetCustomContactFieldsOk() (*[]QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetCustomContactFields(v []QuickbooksDesktopCustomersPostRequestCustomContactFieldsInner)`

SetCustomContactFields sets CustomContactFields field to given value.

### HasCustomContactFields

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasCustomContactFields() bool`

HasCustomContactFields returns a boolean if a field has been set.

### GetEmergencyContact

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmergencyContact() QuickbooksDesktopEmployeesPostRequestEmergencyContact`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmergencyContactOk() (*QuickbooksDesktopEmployeesPostRequestEmergencyContact, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEmergencyContact(v QuickbooksDesktopEmployeesPostRequestEmergencyContact)`

SetEmergencyContact sets EmergencyContact field to given value.

### HasEmergencyContact

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEmergencyContact() bool`

HasEmergencyContact returns a boolean if a field has been set.

### GetEmployeeType

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetEmploymentStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### GetOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetOvertimeExemptStatus() string`

GetOvertimeExemptStatus returns the OvertimeExemptStatus field if non-nil, zero value otherwise.

### GetOvertimeExemptStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetOvertimeExemptStatusOk() (*string, bool)`

GetOvertimeExemptStatusOk returns a tuple with the OvertimeExemptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetOvertimeExemptStatus(v string)`

SetOvertimeExemptStatus sets OvertimeExemptStatus field to given value.

### HasOvertimeExemptStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasOvertimeExemptStatus() bool`

HasOvertimeExemptStatus returns a boolean if a field has been set.

### GetKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetKeyEmployeeStatus() string`

GetKeyEmployeeStatus returns the KeyEmployeeStatus field if non-nil, zero value otherwise.

### GetKeyEmployeeStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetKeyEmployeeStatusOk() (*string, bool)`

GetKeyEmployeeStatusOk returns a tuple with the KeyEmployeeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetKeyEmployeeStatus(v string)`

SetKeyEmployeeStatus sets KeyEmployeeStatus field to given value.

### HasKeyEmployeeStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasKeyEmployeeStatus() bool`

HasKeyEmployeeStatus returns a boolean if a field has been set.

### GetHiredDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetOriginalHireDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetOriginalHireDate() string`

GetOriginalHireDate returns the OriginalHireDate field if non-nil, zero value otherwise.

### GetOriginalHireDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetOriginalHireDateOk() (*string, bool)`

GetOriginalHireDateOk returns a tuple with the OriginalHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalHireDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetOriginalHireDate(v string)`

SetOriginalHireDate sets OriginalHireDate field to given value.

### HasOriginalHireDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasOriginalHireDate() bool`

HasOriginalHireDate returns a boolean if a field has been set.

### GetAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAdjustedServiceDate() string`

GetAdjustedServiceDate returns the AdjustedServiceDate field if non-nil, zero value otherwise.

### GetAdjustedServiceDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAdjustedServiceDateOk() (*string, bool)`

GetAdjustedServiceDateOk returns a tuple with the AdjustedServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetAdjustedServiceDate(v string)`

SetAdjustedServiceDate sets AdjustedServiceDate field to given value.

### HasAdjustedServiceDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasAdjustedServiceDate() bool`

HasAdjustedServiceDate returns a boolean if a field has been set.

### GetTerminationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetUsCitizenshipStatus() string`

GetUsCitizenshipStatus returns the UsCitizenshipStatus field if non-nil, zero value otherwise.

### GetUsCitizenshipStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetUsCitizenshipStatusOk() (*string, bool)`

GetUsCitizenshipStatusOk returns a tuple with the UsCitizenshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetUsCitizenshipStatus(v string)`

SetUsCitizenshipStatus sets UsCitizenshipStatus field to given value.

### HasUsCitizenshipStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasUsCitizenshipStatus() bool`

HasUsCitizenshipStatus returns a boolean if a field has been set.

### GetEthnicity

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.

### HasEthnicity

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### GetDisabilityStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDisabilityStatus() string`

GetDisabilityStatus returns the DisabilityStatus field if non-nil, zero value otherwise.

### GetDisabilityStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDisabilityStatusOk() (*string, bool)`

GetDisabilityStatusOk returns a tuple with the DisabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetDisabilityStatus(v string)`

SetDisabilityStatus sets DisabilityStatus field to given value.

### HasDisabilityStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasDisabilityStatus() bool`

HasDisabilityStatus returns a boolean if a field has been set.

### GetDisabilityDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDisabilityDescription() string`

GetDisabilityDescription returns the DisabilityDescription field if non-nil, zero value otherwise.

### GetDisabilityDescriptionOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetDisabilityDescriptionOk() (*string, bool)`

GetDisabilityDescriptionOk returns a tuple with the DisabilityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetDisabilityDescription(v string)`

SetDisabilityDescription sets DisabilityDescription field to given value.

### HasDisabilityDescription

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasDisabilityDescription() bool`

HasDisabilityDescription returns a boolean if a field has been set.

### GetI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetI9OnFileStatus() string`

GetI9OnFileStatus returns the I9OnFileStatus field if non-nil, zero value otherwise.

### GetI9OnFileStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetI9OnFileStatusOk() (*string, bool)`

GetI9OnFileStatusOk returns a tuple with the I9OnFileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetI9OnFileStatus(v string)`

SetI9OnFileStatus sets I9OnFileStatus field to given value.

### HasI9OnFileStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasI9OnFileStatus() bool`

HasI9OnFileStatus returns a boolean if a field has been set.

### GetWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetWorkAuthorizationExpirationDate() string`

GetWorkAuthorizationExpirationDate returns the WorkAuthorizationExpirationDate field if non-nil, zero value otherwise.

### GetWorkAuthorizationExpirationDateOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetWorkAuthorizationExpirationDateOk() (*string, bool)`

GetWorkAuthorizationExpirationDateOk returns a tuple with the WorkAuthorizationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetWorkAuthorizationExpirationDate(v string)`

SetWorkAuthorizationExpirationDate sets WorkAuthorizationExpirationDate field to given value.

### HasWorkAuthorizationExpirationDate

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasWorkAuthorizationExpirationDate() bool`

HasWorkAuthorizationExpirationDate returns a boolean if a field has been set.

### GetUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetUsVeteranStatus() string`

GetUsVeteranStatus returns the UsVeteranStatus field if non-nil, zero value otherwise.

### GetUsVeteranStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetUsVeteranStatusOk() (*string, bool)`

GetUsVeteranStatusOk returns a tuple with the UsVeteranStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetUsVeteranStatus(v string)`

SetUsVeteranStatus sets UsVeteranStatus field to given value.

### HasUsVeteranStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasUsVeteranStatus() bool`

HasUsVeteranStatus returns a boolean if a field has been set.

### GetMilitaryStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMilitaryStatus() string`

GetMilitaryStatus returns the MilitaryStatus field if non-nil, zero value otherwise.

### GetMilitaryStatusOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetMilitaryStatusOk() (*string, bool)`

GetMilitaryStatusOk returns a tuple with the MilitaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilitaryStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetMilitaryStatus(v string)`

SetMilitaryStatus sets MilitaryStatus field to given value.

### HasMilitaryStatus

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasMilitaryStatus() bool`

HasMilitaryStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetNote

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAdditionalNotes

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAdditionalNotes() []QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetAdditionalNotesOk() (*[]QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetAdditionalNotes(v []QuickbooksDesktopCustomersIdPostRequestAdditionalNotesInner)`

SetAdditionalNotes sets AdditionalNotes field to given value.

### HasAdditionalNotes

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasAdditionalNotes() bool`

HasAdditionalNotes returns a boolean if a field has been set.

### GetBillingRateId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetBillingRateId() string`

GetBillingRateId returns the BillingRateId field if non-nil, zero value otherwise.

### GetBillingRateIdOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetBillingRateIdOk() (*string, bool)`

GetBillingRateIdOk returns a tuple with the BillingRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRateId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetBillingRateId(v string)`

SetBillingRateId sets BillingRateId field to given value.

### HasBillingRateId

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasBillingRateId() bool`

HasBillingRateId returns a boolean if a field has been set.

### GetEmployeePayroll

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmployeePayroll() QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll`

GetEmployeePayroll returns the EmployeePayroll field if non-nil, zero value otherwise.

### GetEmployeePayrollOk

`func (o *QuickbooksDesktopEmployeesIdPostRequest) GetEmployeePayrollOk() (*QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll, bool)`

GetEmployeePayrollOk returns a tuple with the EmployeePayroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayroll

`func (o *QuickbooksDesktopEmployeesIdPostRequest) SetEmployeePayroll(v QuickbooksDesktopEmployeesIdPostRequestEmployeePayroll)`

SetEmployeePayroll sets EmployeePayroll field to given value.

### HasEmployeePayroll

`func (o *QuickbooksDesktopEmployeesIdPostRequest) HasEmployeePayroll() bool`

HasEmployeePayroll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


