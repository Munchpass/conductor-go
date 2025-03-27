# QbdEmployee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this employee. This ID is unique across all employees but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_employee\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this employee was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this employee was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this employee object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this employee, unique across all employees. A concatenation of the employee&#39;s &#x60;firstName&#x60;, &#x60;middleName&#x60;, and &#x60;lastName&#x60; fields.  **NOTE**: Employees do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this employee is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**Salutation** | **string** | The employee&#39;s formal salutation title that precedes their name, such as \&quot;Mr.\&quot;, \&quot;Ms.\&quot;, or \&quot;Dr.\&quot;. | 
**FirstName** | **string** | The employee&#39;s first name. | 
**MiddleName** | **string** | The employee&#39;s middle name. | 
**LastName** | **string** | The employee&#39;s last name. | 
**JobTitle** | **string** | The employee&#39;s job title. | 
**Supervisor** | [**QbdEmployeeSupervisor**](QbdEmployeeSupervisor.md) |  | 
**Department** | **string** | The employee&#39;s department. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | 
**Description** | **string** | A description of this employee. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | 
**TargetBonus** | **string** | The target bonus for this employee, represented as a decimal string. Found in the \&quot;employment job details\&quot; section of the employee&#39;s record in QuickBooks. | 
**Address** | [**QbdEmployeeAddress1**](QbdEmployeeAddress1.md) |  | 
**PrintAs** | **string** | The name to use when printing this employee from QuickBooks. By default, this is the same as the &#x60;name&#x60; field. | 
**Phone** | **string** | The employee&#39;s primary telephone number. | 
**Mobile** | **string** | The employee&#39;s mobile phone number. | 
**Pager** | **string** | The employee&#39;s pager number. | 
**PagerPin** | **string** | The employee&#39;s pager PIN. | 
**AlternatePhone** | **string** | The employee&#39;s alternate telephone number. | 
**Fax** | **string** | The employee&#39;s fax number. | 
**Ssn** | **string** | The employee&#39;s Social Security Number. The value can be with or without dashes.  **NOTE**: This field cannot be changed after the employee is created. | 
**Email** | **string** | The employee&#39;s email address. | 
**CustomContactFields** | [**[]QbdCustomContactField**](QbdCustomContactField.md) | Additional custom contact fields for this employee, such as phone numbers or email addresses. | 
**EmergencyContact** | [**QbdEmployeeEmergencyContact**](QbdEmployeeEmergencyContact.md) |  | 
**EmployeeType** | **string** | The employee type. This affects payroll taxes - a statutory employee is defined as an employee by statute. Note that owners/partners are typically on the \&quot;Other Names\&quot; list in QuickBooks, but if listed as an employee their type will be &#x60;owner&#x60;. | 
**EmploymentStatus** | **string** | Indicates whether this employee is a part-time or full-time employee. | 
**OvertimeExemptStatus** | **string** | Indicates whether this employee is exempt from overtime pay. This classification is based on U.S. labor laws (FLSA).  | 
**KeyEmployeeStatus** | **string** | Indicates whether this employee is a key employee. | 
**Gender** | **string** | This employee&#39;s gender. | 
**HiredDate** | **string** | The date this employee was hired, in ISO 8601 format (YYYY-MM-DD). | 
**OriginalHireDate** | **string** | The original hire date for this employee, in ISO 8601 format (YYYY-MM-DD). | 
**AdjustedServiceDate** | **string** | The adjusted service date for this employee, in ISO 8601 format (YYYY-MM-DD). This date accounts for previous employment periods or leaves that affect seniority. | 
**TerminationDate** | **string** | The date this employee&#39;s employment ended with the company, in ISO 8601 format (YYYY-MM-DD). This is also known as the released date or separation date. | 
**BirthDate** | **string** | This employee&#39;s date of birth, in ISO 8601 format (YYYY-MM-DD). | 
**UsCitizenshipStatus** | **string** | Indicates whether this employee is a U.S. citizen. | 
**Ethnicity** | **string** | This employee&#39;s ethnicity. | 
**DisabilityStatus** | **string** | Indicates whether this employee is disabled. | 
**DisabilityDescription** | **string** | A description of this employee&#39;s disability. | 
**I9OnFileStatus** | **string** | Indicates whether this employee&#39;s I-9 is on file. | 
**WorkAuthorizationExpirationDate** | **string** | The date this employee&#39;s work authorization expires, in ISO 8601 format (YYYY-MM-DD). | 
**UsVeteranStatus** | **string** | Indicates whether this employee is a U.S. veteran. | 
**MilitaryStatus** | **string** | This employee&#39;s military status if they are a U.S. veteran. | 
**AccountNumber** | **string** | The employee&#39;s account number, which appears in the QuickBooks chart of accounts, reports, and graphs.  Note that if the \&quot;Use Account Numbers\&quot; preference is turned off in QuickBooks, the account number may not be visible in the user interface, but it can still be set and retrieved through the API. | 
**Note** | **string** | A note or comment about this employee. | 
**AdditionalNotes** | [**[]QbdNote**](QbdNote.md) | Additional notes about this employee. | 
**BillingRate** | [**QbdEmployeeBillingRate**](QbdEmployeeBillingRate.md) |  | 
**EmployeePayroll** | [**QbdEmployeeEmployeePayroll**](QbdEmployeeEmployeePayroll.md) |  | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the employee object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdEmployee

`func NewQbdEmployee(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, salutation string, firstName string, middleName string, lastName string, jobTitle string, supervisor QbdEmployeeSupervisor, department string, description string, targetBonus string, address QbdEmployeeAddress1, printAs string, phone string, mobile string, pager string, pagerPin string, alternatePhone string, fax string, ssn string, email string, customContactFields []QbdCustomContactField, emergencyContact QbdEmployeeEmergencyContact, employeeType string, employmentStatus string, overtimeExemptStatus string, keyEmployeeStatus string, gender string, hiredDate string, originalHireDate string, adjustedServiceDate string, terminationDate string, birthDate string, usCitizenshipStatus string, ethnicity string, disabilityStatus string, disabilityDescription string, i9OnFileStatus string, workAuthorizationExpirationDate string, usVeteranStatus string, militaryStatus string, accountNumber string, note string, additionalNotes []QbdNote, billingRate QbdEmployeeBillingRate, employeePayroll QbdEmployeeEmployeePayroll, externalId string, customFields []QbdCustomField, ) *QbdEmployee`

NewQbdEmployee instantiates a new QbdEmployee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEmployeeWithDefaults

`func NewQbdEmployeeWithDefaults() *QbdEmployee`

NewQbdEmployeeWithDefaults instantiates a new QbdEmployee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdEmployee) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdEmployee) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdEmployee) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdEmployee) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdEmployee) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdEmployee) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdEmployee) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdEmployee) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdEmployee) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdEmployee) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdEmployee) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdEmployee) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdEmployee) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdEmployee) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdEmployee) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdEmployee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdEmployee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdEmployee) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdEmployee) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdEmployee) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdEmployee) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSalutation

`func (o *QbdEmployee) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *QbdEmployee) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *QbdEmployee) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.


### GetFirstName

`func (o *QbdEmployee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *QbdEmployee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *QbdEmployee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *QbdEmployee) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *QbdEmployee) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *QbdEmployee) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### GetLastName

`func (o *QbdEmployee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *QbdEmployee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *QbdEmployee) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetJobTitle

`func (o *QbdEmployee) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *QbdEmployee) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *QbdEmployee) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### GetSupervisor

`func (o *QbdEmployee) GetSupervisor() QbdEmployeeSupervisor`

GetSupervisor returns the Supervisor field if non-nil, zero value otherwise.

### GetSupervisorOk

`func (o *QbdEmployee) GetSupervisorOk() (*QbdEmployeeSupervisor, bool)`

GetSupervisorOk returns a tuple with the Supervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisor

`func (o *QbdEmployee) SetSupervisor(v QbdEmployeeSupervisor)`

SetSupervisor sets Supervisor field to given value.


### GetDepartment

`func (o *QbdEmployee) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *QbdEmployee) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *QbdEmployee) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDescription

`func (o *QbdEmployee) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QbdEmployee) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QbdEmployee) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTargetBonus

`func (o *QbdEmployee) GetTargetBonus() string`

GetTargetBonus returns the TargetBonus field if non-nil, zero value otherwise.

### GetTargetBonusOk

`func (o *QbdEmployee) GetTargetBonusOk() (*string, bool)`

GetTargetBonusOk returns a tuple with the TargetBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBonus

`func (o *QbdEmployee) SetTargetBonus(v string)`

SetTargetBonus sets TargetBonus field to given value.


### GetAddress

`func (o *QbdEmployee) GetAddress() QbdEmployeeAddress1`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QbdEmployee) GetAddressOk() (*QbdEmployeeAddress1, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QbdEmployee) SetAddress(v QbdEmployeeAddress1)`

SetAddress sets Address field to given value.


### GetPrintAs

`func (o *QbdEmployee) GetPrintAs() string`

GetPrintAs returns the PrintAs field if non-nil, zero value otherwise.

### GetPrintAsOk

`func (o *QbdEmployee) GetPrintAsOk() (*string, bool)`

GetPrintAsOk returns a tuple with the PrintAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintAs

`func (o *QbdEmployee) SetPrintAs(v string)`

SetPrintAs sets PrintAs field to given value.


### GetPhone

`func (o *QbdEmployee) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QbdEmployee) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QbdEmployee) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetMobile

`func (o *QbdEmployee) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *QbdEmployee) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *QbdEmployee) SetMobile(v string)`

SetMobile sets Mobile field to given value.


### GetPager

`func (o *QbdEmployee) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *QbdEmployee) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *QbdEmployee) SetPager(v string)`

SetPager sets Pager field to given value.


### GetPagerPin

`func (o *QbdEmployee) GetPagerPin() string`

GetPagerPin returns the PagerPin field if non-nil, zero value otherwise.

### GetPagerPinOk

`func (o *QbdEmployee) GetPagerPinOk() (*string, bool)`

GetPagerPinOk returns a tuple with the PagerPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerPin

`func (o *QbdEmployee) SetPagerPin(v string)`

SetPagerPin sets PagerPin field to given value.


### GetAlternatePhone

`func (o *QbdEmployee) GetAlternatePhone() string`

GetAlternatePhone returns the AlternatePhone field if non-nil, zero value otherwise.

### GetAlternatePhoneOk

`func (o *QbdEmployee) GetAlternatePhoneOk() (*string, bool)`

GetAlternatePhoneOk returns a tuple with the AlternatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePhone

`func (o *QbdEmployee) SetAlternatePhone(v string)`

SetAlternatePhone sets AlternatePhone field to given value.


### GetFax

`func (o *QbdEmployee) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *QbdEmployee) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *QbdEmployee) SetFax(v string)`

SetFax sets Fax field to given value.


### GetSsn

`func (o *QbdEmployee) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *QbdEmployee) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *QbdEmployee) SetSsn(v string)`

SetSsn sets Ssn field to given value.


### GetEmail

`func (o *QbdEmployee) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QbdEmployee) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QbdEmployee) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCustomContactFields

`func (o *QbdEmployee) GetCustomContactFields() []QbdCustomContactField`

GetCustomContactFields returns the CustomContactFields field if non-nil, zero value otherwise.

### GetCustomContactFieldsOk

`func (o *QbdEmployee) GetCustomContactFieldsOk() (*[]QbdCustomContactField, bool)`

GetCustomContactFieldsOk returns a tuple with the CustomContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContactFields

`func (o *QbdEmployee) SetCustomContactFields(v []QbdCustomContactField)`

SetCustomContactFields sets CustomContactFields field to given value.


### GetEmergencyContact

`func (o *QbdEmployee) GetEmergencyContact() QbdEmployeeEmergencyContact`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *QbdEmployee) GetEmergencyContactOk() (*QbdEmployeeEmergencyContact, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *QbdEmployee) SetEmergencyContact(v QbdEmployeeEmergencyContact)`

SetEmergencyContact sets EmergencyContact field to given value.


### GetEmployeeType

`func (o *QbdEmployee) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *QbdEmployee) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *QbdEmployee) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.


### GetEmploymentStatus

`func (o *QbdEmployee) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *QbdEmployee) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *QbdEmployee) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.


### GetOvertimeExemptStatus

`func (o *QbdEmployee) GetOvertimeExemptStatus() string`

GetOvertimeExemptStatus returns the OvertimeExemptStatus field if non-nil, zero value otherwise.

### GetOvertimeExemptStatusOk

`func (o *QbdEmployee) GetOvertimeExemptStatusOk() (*string, bool)`

GetOvertimeExemptStatusOk returns a tuple with the OvertimeExemptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimeExemptStatus

`func (o *QbdEmployee) SetOvertimeExemptStatus(v string)`

SetOvertimeExemptStatus sets OvertimeExemptStatus field to given value.


### GetKeyEmployeeStatus

`func (o *QbdEmployee) GetKeyEmployeeStatus() string`

GetKeyEmployeeStatus returns the KeyEmployeeStatus field if non-nil, zero value otherwise.

### GetKeyEmployeeStatusOk

`func (o *QbdEmployee) GetKeyEmployeeStatusOk() (*string, bool)`

GetKeyEmployeeStatusOk returns a tuple with the KeyEmployeeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEmployeeStatus

`func (o *QbdEmployee) SetKeyEmployeeStatus(v string)`

SetKeyEmployeeStatus sets KeyEmployeeStatus field to given value.


### GetGender

`func (o *QbdEmployee) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *QbdEmployee) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *QbdEmployee) SetGender(v string)`

SetGender sets Gender field to given value.


### GetHiredDate

`func (o *QbdEmployee) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *QbdEmployee) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *QbdEmployee) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.


### GetOriginalHireDate

`func (o *QbdEmployee) GetOriginalHireDate() string`

GetOriginalHireDate returns the OriginalHireDate field if non-nil, zero value otherwise.

### GetOriginalHireDateOk

`func (o *QbdEmployee) GetOriginalHireDateOk() (*string, bool)`

GetOriginalHireDateOk returns a tuple with the OriginalHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalHireDate

`func (o *QbdEmployee) SetOriginalHireDate(v string)`

SetOriginalHireDate sets OriginalHireDate field to given value.


### GetAdjustedServiceDate

`func (o *QbdEmployee) GetAdjustedServiceDate() string`

GetAdjustedServiceDate returns the AdjustedServiceDate field if non-nil, zero value otherwise.

### GetAdjustedServiceDateOk

`func (o *QbdEmployee) GetAdjustedServiceDateOk() (*string, bool)`

GetAdjustedServiceDateOk returns a tuple with the AdjustedServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedServiceDate

`func (o *QbdEmployee) SetAdjustedServiceDate(v string)`

SetAdjustedServiceDate sets AdjustedServiceDate field to given value.


### GetTerminationDate

`func (o *QbdEmployee) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *QbdEmployee) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *QbdEmployee) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.


### GetBirthDate

`func (o *QbdEmployee) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *QbdEmployee) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *QbdEmployee) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.


### GetUsCitizenshipStatus

`func (o *QbdEmployee) GetUsCitizenshipStatus() string`

GetUsCitizenshipStatus returns the UsCitizenshipStatus field if non-nil, zero value otherwise.

### GetUsCitizenshipStatusOk

`func (o *QbdEmployee) GetUsCitizenshipStatusOk() (*string, bool)`

GetUsCitizenshipStatusOk returns a tuple with the UsCitizenshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsCitizenshipStatus

`func (o *QbdEmployee) SetUsCitizenshipStatus(v string)`

SetUsCitizenshipStatus sets UsCitizenshipStatus field to given value.


### GetEthnicity

`func (o *QbdEmployee) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *QbdEmployee) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *QbdEmployee) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.


### GetDisabilityStatus

`func (o *QbdEmployee) GetDisabilityStatus() string`

GetDisabilityStatus returns the DisabilityStatus field if non-nil, zero value otherwise.

### GetDisabilityStatusOk

`func (o *QbdEmployee) GetDisabilityStatusOk() (*string, bool)`

GetDisabilityStatusOk returns a tuple with the DisabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityStatus

`func (o *QbdEmployee) SetDisabilityStatus(v string)`

SetDisabilityStatus sets DisabilityStatus field to given value.


### GetDisabilityDescription

`func (o *QbdEmployee) GetDisabilityDescription() string`

GetDisabilityDescription returns the DisabilityDescription field if non-nil, zero value otherwise.

### GetDisabilityDescriptionOk

`func (o *QbdEmployee) GetDisabilityDescriptionOk() (*string, bool)`

GetDisabilityDescriptionOk returns a tuple with the DisabilityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityDescription

`func (o *QbdEmployee) SetDisabilityDescription(v string)`

SetDisabilityDescription sets DisabilityDescription field to given value.


### GetI9OnFileStatus

`func (o *QbdEmployee) GetI9OnFileStatus() string`

GetI9OnFileStatus returns the I9OnFileStatus field if non-nil, zero value otherwise.

### GetI9OnFileStatusOk

`func (o *QbdEmployee) GetI9OnFileStatusOk() (*string, bool)`

GetI9OnFileStatusOk returns a tuple with the I9OnFileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI9OnFileStatus

`func (o *QbdEmployee) SetI9OnFileStatus(v string)`

SetI9OnFileStatus sets I9OnFileStatus field to given value.


### GetWorkAuthorizationExpirationDate

`func (o *QbdEmployee) GetWorkAuthorizationExpirationDate() string`

GetWorkAuthorizationExpirationDate returns the WorkAuthorizationExpirationDate field if non-nil, zero value otherwise.

### GetWorkAuthorizationExpirationDateOk

`func (o *QbdEmployee) GetWorkAuthorizationExpirationDateOk() (*string, bool)`

GetWorkAuthorizationExpirationDateOk returns a tuple with the WorkAuthorizationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkAuthorizationExpirationDate

`func (o *QbdEmployee) SetWorkAuthorizationExpirationDate(v string)`

SetWorkAuthorizationExpirationDate sets WorkAuthorizationExpirationDate field to given value.


### GetUsVeteranStatus

`func (o *QbdEmployee) GetUsVeteranStatus() string`

GetUsVeteranStatus returns the UsVeteranStatus field if non-nil, zero value otherwise.

### GetUsVeteranStatusOk

`func (o *QbdEmployee) GetUsVeteranStatusOk() (*string, bool)`

GetUsVeteranStatusOk returns a tuple with the UsVeteranStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVeteranStatus

`func (o *QbdEmployee) SetUsVeteranStatus(v string)`

SetUsVeteranStatus sets UsVeteranStatus field to given value.


### GetMilitaryStatus

`func (o *QbdEmployee) GetMilitaryStatus() string`

GetMilitaryStatus returns the MilitaryStatus field if non-nil, zero value otherwise.

### GetMilitaryStatusOk

`func (o *QbdEmployee) GetMilitaryStatusOk() (*string, bool)`

GetMilitaryStatusOk returns a tuple with the MilitaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilitaryStatus

`func (o *QbdEmployee) SetMilitaryStatus(v string)`

SetMilitaryStatus sets MilitaryStatus field to given value.


### GetAccountNumber

`func (o *QbdEmployee) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *QbdEmployee) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *QbdEmployee) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetNote

`func (o *QbdEmployee) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdEmployee) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdEmployee) SetNote(v string)`

SetNote sets Note field to given value.


### GetAdditionalNotes

`func (o *QbdEmployee) GetAdditionalNotes() []QbdNote`

GetAdditionalNotes returns the AdditionalNotes field if non-nil, zero value otherwise.

### GetAdditionalNotesOk

`func (o *QbdEmployee) GetAdditionalNotesOk() (*[]QbdNote, bool)`

GetAdditionalNotesOk returns a tuple with the AdditionalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotes

`func (o *QbdEmployee) SetAdditionalNotes(v []QbdNote)`

SetAdditionalNotes sets AdditionalNotes field to given value.


### GetBillingRate

`func (o *QbdEmployee) GetBillingRate() QbdEmployeeBillingRate`

GetBillingRate returns the BillingRate field if non-nil, zero value otherwise.

### GetBillingRateOk

`func (o *QbdEmployee) GetBillingRateOk() (*QbdEmployeeBillingRate, bool)`

GetBillingRateOk returns a tuple with the BillingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRate

`func (o *QbdEmployee) SetBillingRate(v QbdEmployeeBillingRate)`

SetBillingRate sets BillingRate field to given value.


### GetEmployeePayroll

`func (o *QbdEmployee) GetEmployeePayroll() QbdEmployeeEmployeePayroll`

GetEmployeePayroll returns the EmployeePayroll field if non-nil, zero value otherwise.

### GetEmployeePayrollOk

`func (o *QbdEmployee) GetEmployeePayrollOk() (*QbdEmployeeEmployeePayroll, bool)`

GetEmployeePayrollOk returns a tuple with the EmployeePayroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeePayroll

`func (o *QbdEmployee) SetEmployeePayroll(v QbdEmployeeEmployeePayroll)`

SetEmployeePayroll sets EmployeePayroll field to given value.


### GetExternalId

`func (o *QbdEmployee) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdEmployee) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdEmployee) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCustomFields

`func (o *QbdEmployee) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdEmployee) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdEmployee) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


