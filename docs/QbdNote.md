# QbdNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The auto-incrementing identifier assigned by QuickBooks to this note. | 
**Date** | **NullableString** | The date this note was last updated, in ISO 8601 format (YYYY-MM-DD). | 
**Note** | **string** | The text of this note. | 

## Methods

### NewQbdNote

`func NewQbdNote(id float32, date NullableString, note string, ) *QbdNote`

NewQbdNote instantiates a new QbdNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdNoteWithDefaults

`func NewQbdNoteWithDefaults() *QbdNote`

NewQbdNoteWithDefaults instantiates a new QbdNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdNote) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdNote) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdNote) SetId(v float32)`

SetId sets Id field to given value.


### GetDate

`func (o *QbdNote) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *QbdNote) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *QbdNote) SetDate(v string)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *QbdNote) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *QbdNote) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetNote

`func (o *QbdNote) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *QbdNote) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *QbdNote) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


