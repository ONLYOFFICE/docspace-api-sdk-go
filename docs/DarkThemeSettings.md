# DarkThemeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Theme** | Pointer to [**DarkThemeSettingsType**](DarkThemeSettingsType.md) | The theme type. | [optional] 
**LastModified** | Pointer to **time.Time** | The last modified date. | [optional] 

## Methods

### NewDarkThemeSettings

`func NewDarkThemeSettings() *DarkThemeSettings`

NewDarkThemeSettings instantiates a new DarkThemeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDarkThemeSettingsWithDefaults

`func NewDarkThemeSettingsWithDefaults() *DarkThemeSettings`

NewDarkThemeSettingsWithDefaults instantiates a new DarkThemeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTheme

`func (o *DarkThemeSettings) GetTheme() DarkThemeSettingsType`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *DarkThemeSettings) GetThemeOk() (*DarkThemeSettingsType, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *DarkThemeSettings) SetTheme(v DarkThemeSettingsType)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *DarkThemeSettings) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetLastModified

`func (o *DarkThemeSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DarkThemeSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DarkThemeSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DarkThemeSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


