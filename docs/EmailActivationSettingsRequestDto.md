# EmailActivationSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Show** | Pointer to **bool** | Whether the reminder is shown; send false to dismiss it for the calling user. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewEmailActivationSettingsRequestDto

`func NewEmailActivationSettingsRequestDto() *EmailActivationSettingsRequestDto`

NewEmailActivationSettingsRequestDto instantiates a new EmailActivationSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActivationSettingsRequestDtoWithDefaults

`func NewEmailActivationSettingsRequestDtoWithDefaults() *EmailActivationSettingsRequestDto`

NewEmailActivationSettingsRequestDtoWithDefaults instantiates a new EmailActivationSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShow

`func (o *EmailActivationSettingsRequestDto) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *EmailActivationSettingsRequestDto) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *EmailActivationSettingsRequestDto) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *EmailActivationSettingsRequestDto) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetLastModified

`func (o *EmailActivationSettingsRequestDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EmailActivationSettingsRequestDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EmailActivationSettingsRequestDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *EmailActivationSettingsRequestDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


