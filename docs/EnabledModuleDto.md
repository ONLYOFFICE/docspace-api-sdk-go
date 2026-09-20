# EnabledModuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The module's product class name, HTML-escaped. It is a display-oriented identifier and not the GUID the  access-settings operations work with, so it must not be passed to `GET api/2.0/settings/security/{id}`. | [optional] 
**Title** | Pointer to **NullableString** | The module name in the portal language, HTML-escaped and ready to be rendered as text. | [optional] 

## Methods

### NewEnabledModuleDto

`func NewEnabledModuleDto() *EnabledModuleDto`

NewEnabledModuleDto instantiates a new EnabledModuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledModuleDtoWithDefaults

`func NewEnabledModuleDtoWithDefaults() *EnabledModuleDto`

NewEnabledModuleDtoWithDefaults instantiates a new EnabledModuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnabledModuleDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnabledModuleDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnabledModuleDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnabledModuleDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EnabledModuleDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EnabledModuleDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *EnabledModuleDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EnabledModuleDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EnabledModuleDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EnabledModuleDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EnabledModuleDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EnabledModuleDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


