# QbdCurrentAppAccessRights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutomaticLoginAllowed** | **bool** | Indicates whether applications can use auto-login to access this company file. | 
**AutomaticLoginUserName** | **string** | If auto-login is allowed for this company file, specifies the user name that is allowed to use auto-login. | 
**IsPersonalDataAccessAllowed** | **bool** | Indicates whether access is allowed to personal (sensitive) data in this company file. | 

## Methods

### NewQbdCurrentAppAccessRights

`func NewQbdCurrentAppAccessRights(isAutomaticLoginAllowed bool, automaticLoginUserName string, isPersonalDataAccessAllowed bool, ) *QbdCurrentAppAccessRights`

NewQbdCurrentAppAccessRights instantiates a new QbdCurrentAppAccessRights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCurrentAppAccessRightsWithDefaults

`func NewQbdCurrentAppAccessRightsWithDefaults() *QbdCurrentAppAccessRights`

NewQbdCurrentAppAccessRightsWithDefaults instantiates a new QbdCurrentAppAccessRights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutomaticLoginAllowed

`func (o *QbdCurrentAppAccessRights) GetIsAutomaticLoginAllowed() bool`

GetIsAutomaticLoginAllowed returns the IsAutomaticLoginAllowed field if non-nil, zero value otherwise.

### GetIsAutomaticLoginAllowedOk

`func (o *QbdCurrentAppAccessRights) GetIsAutomaticLoginAllowedOk() (*bool, bool)`

GetIsAutomaticLoginAllowedOk returns a tuple with the IsAutomaticLoginAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticLoginAllowed

`func (o *QbdCurrentAppAccessRights) SetIsAutomaticLoginAllowed(v bool)`

SetIsAutomaticLoginAllowed sets IsAutomaticLoginAllowed field to given value.


### GetAutomaticLoginUserName

`func (o *QbdCurrentAppAccessRights) GetAutomaticLoginUserName() string`

GetAutomaticLoginUserName returns the AutomaticLoginUserName field if non-nil, zero value otherwise.

### GetAutomaticLoginUserNameOk

`func (o *QbdCurrentAppAccessRights) GetAutomaticLoginUserNameOk() (*string, bool)`

GetAutomaticLoginUserNameOk returns a tuple with the AutomaticLoginUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticLoginUserName

`func (o *QbdCurrentAppAccessRights) SetAutomaticLoginUserName(v string)`

SetAutomaticLoginUserName sets AutomaticLoginUserName field to given value.


### GetIsPersonalDataAccessAllowed

`func (o *QbdCurrentAppAccessRights) GetIsPersonalDataAccessAllowed() bool`

GetIsPersonalDataAccessAllowed returns the IsPersonalDataAccessAllowed field if non-nil, zero value otherwise.

### GetIsPersonalDataAccessAllowedOk

`func (o *QbdCurrentAppAccessRights) GetIsPersonalDataAccessAllowedOk() (*bool, bool)`

GetIsPersonalDataAccessAllowedOk returns a tuple with the IsPersonalDataAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalDataAccessAllowed

`func (o *QbdCurrentAppAccessRights) SetIsPersonalDataAccessAllowed(v bool)`

SetIsPersonalDataAccessAllowed sets IsPersonalDataAccessAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


