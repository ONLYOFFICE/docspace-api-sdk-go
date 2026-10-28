# DarkThemeSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Theme** | Pointer to [**DarkThemeSettingsType**](DarkThemeSettingsType.md) | The theme type. | [optional] 
**LastModified** | Pointer to **time.Time** | The last modified date. | [optional] 

## Methods

### NewDarkThemeSettingsDto

`func NewDarkThemeSettingsDto() *DarkThemeSettingsDto`

NewDarkThemeSettingsDto instantiates a new DarkThemeSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDarkThemeSettingsDtoWithDefaults

`func NewDarkThemeSettingsDtoWithDefaults() *DarkThemeSettingsDto`

NewDarkThemeSettingsDtoWithDefaults instantiates a new DarkThemeSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTheme

`func (o *DarkThemeSettingsDto) GetTheme() DarkThemeSettingsType`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *DarkThemeSettingsDto) GetThemeOk() (*DarkThemeSettingsType, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *DarkThemeSettingsDto) SetTheme(v DarkThemeSettingsType)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *DarkThemeSettingsDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetLastModified

`func (o *DarkThemeSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DarkThemeSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DarkThemeSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DarkThemeSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


