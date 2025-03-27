# QbdCreditCardTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | The status code returned in the original QBMS transaction response for this credit card transaction. | 
**StatusMessage** | **string** | The status message returned in the original QBMS transaction response for this credit card transaction. | 
**CreditCardTransactionId** | **string** | The ID returned from the credit card processor for this credit card transaction. | 
**MerchantAccountNumber** | **string** | The QBMS account number of the merchant who is running this transaction using the customer&#39;s credit card. | 
**AuthorizationCode** | **NullableString** | The authorization code returned from the credit card processor to indicate that this charge will be paid by the card issuer. | 
**AvsStreetStatus** | **NullableString** | Indicates whether the street address supplied in the transaction request matches the customer&#39;s address on file at the card issuer. | 
**AvsZipStatus** | **NullableString** | Indicates whether the customer postal ZIP code supplied in the transaction request matches the customer&#39;s postal code recognized at the card issuer. | 
**CardSecurityCodeMatch** | **NullableString** | Indicates whether the card security code supplied in the transaction request matches the card security code recognized for that credit card number at the card issuer. | 
**ReconBatchId** | **NullableString** | An internal ID returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | 
**PaymentGroupingCode** | **NullableFloat32** | An internal code returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | 
**PaymentStatus** | **string** | Indicates whether this credit card transaction is known to have been successfully processed by the card issuer. | 
**TransactionAuthorizedAt** | **string** | The date and time when the credit card processor authorized this credit card transaction. | 
**TransactionAuthorizationStamp** | **NullableFloat32** | An internal value for this credit card transaction, needed for the QuickBooks reconciliation feature. | 
**ClientTransactionId** | **NullableString** | A value returned from QBMS transactions for future use by the QuickBooks Reconciliation feature. | 

## Methods

### NewQbdCreditCardTransactionResponse

`func NewQbdCreditCardTransactionResponse(statusCode float32, statusMessage string, creditCardTransactionId string, merchantAccountNumber string, authorizationCode NullableString, avsStreetStatus NullableString, avsZipStatus NullableString, cardSecurityCodeMatch NullableString, reconBatchId NullableString, paymentGroupingCode NullableFloat32, paymentStatus string, transactionAuthorizedAt string, transactionAuthorizationStamp NullableFloat32, clientTransactionId NullableString, ) *QbdCreditCardTransactionResponse`

NewQbdCreditCardTransactionResponse instantiates a new QbdCreditCardTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditCardTransactionResponseWithDefaults

`func NewQbdCreditCardTransactionResponseWithDefaults() *QbdCreditCardTransactionResponse`

NewQbdCreditCardTransactionResponseWithDefaults instantiates a new QbdCreditCardTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *QbdCreditCardTransactionResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *QbdCreditCardTransactionResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *QbdCreditCardTransactionResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *QbdCreditCardTransactionResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *QbdCreditCardTransactionResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *QbdCreditCardTransactionResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetCreditCardTransactionId

`func (o *QbdCreditCardTransactionResponse) GetCreditCardTransactionId() string`

GetCreditCardTransactionId returns the CreditCardTransactionId field if non-nil, zero value otherwise.

### GetCreditCardTransactionIdOk

`func (o *QbdCreditCardTransactionResponse) GetCreditCardTransactionIdOk() (*string, bool)`

GetCreditCardTransactionIdOk returns a tuple with the CreditCardTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransactionId

`func (o *QbdCreditCardTransactionResponse) SetCreditCardTransactionId(v string)`

SetCreditCardTransactionId sets CreditCardTransactionId field to given value.


### GetMerchantAccountNumber

`func (o *QbdCreditCardTransactionResponse) GetMerchantAccountNumber() string`

GetMerchantAccountNumber returns the MerchantAccountNumber field if non-nil, zero value otherwise.

### GetMerchantAccountNumberOk

`func (o *QbdCreditCardTransactionResponse) GetMerchantAccountNumberOk() (*string, bool)`

GetMerchantAccountNumberOk returns a tuple with the MerchantAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountNumber

`func (o *QbdCreditCardTransactionResponse) SetMerchantAccountNumber(v string)`

SetMerchantAccountNumber sets MerchantAccountNumber field to given value.


### GetAuthorizationCode

`func (o *QbdCreditCardTransactionResponse) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *QbdCreditCardTransactionResponse) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *QbdCreditCardTransactionResponse) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.


### SetAuthorizationCodeNil

`func (o *QbdCreditCardTransactionResponse) SetAuthorizationCodeNil(b bool)`

 SetAuthorizationCodeNil sets the value for AuthorizationCode to be an explicit nil

### UnsetAuthorizationCode
`func (o *QbdCreditCardTransactionResponse) UnsetAuthorizationCode()`

UnsetAuthorizationCode ensures that no value is present for AuthorizationCode, not even an explicit nil
### GetAvsStreetStatus

`func (o *QbdCreditCardTransactionResponse) GetAvsStreetStatus() string`

GetAvsStreetStatus returns the AvsStreetStatus field if non-nil, zero value otherwise.

### GetAvsStreetStatusOk

`func (o *QbdCreditCardTransactionResponse) GetAvsStreetStatusOk() (*string, bool)`

GetAvsStreetStatusOk returns a tuple with the AvsStreetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsStreetStatus

`func (o *QbdCreditCardTransactionResponse) SetAvsStreetStatus(v string)`

SetAvsStreetStatus sets AvsStreetStatus field to given value.


### SetAvsStreetStatusNil

`func (o *QbdCreditCardTransactionResponse) SetAvsStreetStatusNil(b bool)`

 SetAvsStreetStatusNil sets the value for AvsStreetStatus to be an explicit nil

### UnsetAvsStreetStatus
`func (o *QbdCreditCardTransactionResponse) UnsetAvsStreetStatus()`

UnsetAvsStreetStatus ensures that no value is present for AvsStreetStatus, not even an explicit nil
### GetAvsZipStatus

`func (o *QbdCreditCardTransactionResponse) GetAvsZipStatus() string`

GetAvsZipStatus returns the AvsZipStatus field if non-nil, zero value otherwise.

### GetAvsZipStatusOk

`func (o *QbdCreditCardTransactionResponse) GetAvsZipStatusOk() (*string, bool)`

GetAvsZipStatusOk returns a tuple with the AvsZipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsZipStatus

`func (o *QbdCreditCardTransactionResponse) SetAvsZipStatus(v string)`

SetAvsZipStatus sets AvsZipStatus field to given value.


### SetAvsZipStatusNil

`func (o *QbdCreditCardTransactionResponse) SetAvsZipStatusNil(b bool)`

 SetAvsZipStatusNil sets the value for AvsZipStatus to be an explicit nil

### UnsetAvsZipStatus
`func (o *QbdCreditCardTransactionResponse) UnsetAvsZipStatus()`

UnsetAvsZipStatus ensures that no value is present for AvsZipStatus, not even an explicit nil
### GetCardSecurityCodeMatch

`func (o *QbdCreditCardTransactionResponse) GetCardSecurityCodeMatch() string`

GetCardSecurityCodeMatch returns the CardSecurityCodeMatch field if non-nil, zero value otherwise.

### GetCardSecurityCodeMatchOk

`func (o *QbdCreditCardTransactionResponse) GetCardSecurityCodeMatchOk() (*string, bool)`

GetCardSecurityCodeMatchOk returns a tuple with the CardSecurityCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSecurityCodeMatch

`func (o *QbdCreditCardTransactionResponse) SetCardSecurityCodeMatch(v string)`

SetCardSecurityCodeMatch sets CardSecurityCodeMatch field to given value.


### SetCardSecurityCodeMatchNil

`func (o *QbdCreditCardTransactionResponse) SetCardSecurityCodeMatchNil(b bool)`

 SetCardSecurityCodeMatchNil sets the value for CardSecurityCodeMatch to be an explicit nil

### UnsetCardSecurityCodeMatch
`func (o *QbdCreditCardTransactionResponse) UnsetCardSecurityCodeMatch()`

UnsetCardSecurityCodeMatch ensures that no value is present for CardSecurityCodeMatch, not even an explicit nil
### GetReconBatchId

`func (o *QbdCreditCardTransactionResponse) GetReconBatchId() string`

GetReconBatchId returns the ReconBatchId field if non-nil, zero value otherwise.

### GetReconBatchIdOk

`func (o *QbdCreditCardTransactionResponse) GetReconBatchIdOk() (*string, bool)`

GetReconBatchIdOk returns a tuple with the ReconBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconBatchId

`func (o *QbdCreditCardTransactionResponse) SetReconBatchId(v string)`

SetReconBatchId sets ReconBatchId field to given value.


### SetReconBatchIdNil

`func (o *QbdCreditCardTransactionResponse) SetReconBatchIdNil(b bool)`

 SetReconBatchIdNil sets the value for ReconBatchId to be an explicit nil

### UnsetReconBatchId
`func (o *QbdCreditCardTransactionResponse) UnsetReconBatchId()`

UnsetReconBatchId ensures that no value is present for ReconBatchId, not even an explicit nil
### GetPaymentGroupingCode

`func (o *QbdCreditCardTransactionResponse) GetPaymentGroupingCode() float32`

GetPaymentGroupingCode returns the PaymentGroupingCode field if non-nil, zero value otherwise.

### GetPaymentGroupingCodeOk

`func (o *QbdCreditCardTransactionResponse) GetPaymentGroupingCodeOk() (*float32, bool)`

GetPaymentGroupingCodeOk returns a tuple with the PaymentGroupingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroupingCode

`func (o *QbdCreditCardTransactionResponse) SetPaymentGroupingCode(v float32)`

SetPaymentGroupingCode sets PaymentGroupingCode field to given value.


### SetPaymentGroupingCodeNil

`func (o *QbdCreditCardTransactionResponse) SetPaymentGroupingCodeNil(b bool)`

 SetPaymentGroupingCodeNil sets the value for PaymentGroupingCode to be an explicit nil

### UnsetPaymentGroupingCode
`func (o *QbdCreditCardTransactionResponse) UnsetPaymentGroupingCode()`

UnsetPaymentGroupingCode ensures that no value is present for PaymentGroupingCode, not even an explicit nil
### GetPaymentStatus

`func (o *QbdCreditCardTransactionResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *QbdCreditCardTransactionResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *QbdCreditCardTransactionResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetTransactionAuthorizedAt

`func (o *QbdCreditCardTransactionResponse) GetTransactionAuthorizedAt() string`

GetTransactionAuthorizedAt returns the TransactionAuthorizedAt field if non-nil, zero value otherwise.

### GetTransactionAuthorizedAtOk

`func (o *QbdCreditCardTransactionResponse) GetTransactionAuthorizedAtOk() (*string, bool)`

GetTransactionAuthorizedAtOk returns a tuple with the TransactionAuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizedAt

`func (o *QbdCreditCardTransactionResponse) SetTransactionAuthorizedAt(v string)`

SetTransactionAuthorizedAt sets TransactionAuthorizedAt field to given value.


### GetTransactionAuthorizationStamp

`func (o *QbdCreditCardTransactionResponse) GetTransactionAuthorizationStamp() float32`

GetTransactionAuthorizationStamp returns the TransactionAuthorizationStamp field if non-nil, zero value otherwise.

### GetTransactionAuthorizationStampOk

`func (o *QbdCreditCardTransactionResponse) GetTransactionAuthorizationStampOk() (*float32, bool)`

GetTransactionAuthorizationStampOk returns a tuple with the TransactionAuthorizationStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizationStamp

`func (o *QbdCreditCardTransactionResponse) SetTransactionAuthorizationStamp(v float32)`

SetTransactionAuthorizationStamp sets TransactionAuthorizationStamp field to given value.


### SetTransactionAuthorizationStampNil

`func (o *QbdCreditCardTransactionResponse) SetTransactionAuthorizationStampNil(b bool)`

 SetTransactionAuthorizationStampNil sets the value for TransactionAuthorizationStamp to be an explicit nil

### UnsetTransactionAuthorizationStamp
`func (o *QbdCreditCardTransactionResponse) UnsetTransactionAuthorizationStamp()`

UnsetTransactionAuthorizationStamp ensures that no value is present for TransactionAuthorizationStamp, not even an explicit nil
### GetClientTransactionId

`func (o *QbdCreditCardTransactionResponse) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *QbdCreditCardTransactionResponse) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *QbdCreditCardTransactionResponse) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.


### SetClientTransactionIdNil

`func (o *QbdCreditCardTransactionResponse) SetClientTransactionIdNil(b bool)`

 SetClientTransactionIdNil sets the value for ClientTransactionId to be an explicit nil

### UnsetClientTransactionId
`func (o *QbdCreditCardTransactionResponse) UnsetClientTransactionId()`

UnsetClientTransactionId ensures that no value is present for ClientTransactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


