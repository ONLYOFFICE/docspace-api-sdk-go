# WizardSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Specifies if the Wizard settings are completed or not | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewWizardSettingsDto

`func NewWizardSettingsDto() *WizardSettingsDto`

NewWizardSettingsDto instantiates a new WizardSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardSettingsDtoWithDefaults

`func NewWizardSettingsDtoWithDefaults() *WizardSettingsDto`

NewWizardSettingsDtoWithDefaults instantiates a new WizardSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *WizardSettingsDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *WizardSettingsDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *WizardSettingsDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *WizardSettingsDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetLastModified

`func (o *WizardSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WizardSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WizardSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *WizardSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


