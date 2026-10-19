# EmailActivationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Show** | Pointer to **bool** | Specifies whether the email activation settings are shown or hidden. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewEmailActivationSettingsDto

`func NewEmailActivationSettingsDto() *EmailActivationSettingsDto`

NewEmailActivationSettingsDto instantiates a new EmailActivationSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActivationSettingsDtoWithDefaults

`func NewEmailActivationSettingsDtoWithDefaults() *EmailActivationSettingsDto`

NewEmailActivationSettingsDtoWithDefaults instantiates a new EmailActivationSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShow

`func (o *EmailActivationSettingsDto) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *EmailActivationSettingsDto) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *EmailActivationSettingsDto) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *EmailActivationSettingsDto) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetLastModified

`func (o *EmailActivationSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EmailActivationSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EmailActivationSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *EmailActivationSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


