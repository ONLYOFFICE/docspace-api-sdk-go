# CreateApiKeyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The API key name. | 
**Permissions** | Pointer to **[]string** | The list of permissions granted to the API key. | [optional] 
**ExpiresInDays** | Pointer to **NullableInt32** | The number of days until the API key expires (null for no expiration). | [optional] 

## Methods

### NewCreateApiKeyRequestDto

`func NewCreateApiKeyRequestDto(name string, ) *CreateApiKeyRequestDto`

NewCreateApiKeyRequestDto instantiates a new CreateApiKeyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestDtoWithDefaults

`func NewCreateApiKeyRequestDtoWithDefaults() *CreateApiKeyRequestDto`

NewCreateApiKeyRequestDtoWithDefaults instantiates a new CreateApiKeyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApiKeyRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKeyRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKeyRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CreateApiKeyRequestDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateApiKeyRequestDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateApiKeyRequestDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateApiKeyRequestDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreateApiKeyRequestDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreateApiKeyRequestDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetExpiresInDays

`func (o *CreateApiKeyRequestDto) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *CreateApiKeyRequestDto) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *CreateApiKeyRequestDto) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.

### HasExpiresInDays

`func (o *CreateApiKeyRequestDto) HasExpiresInDays() bool`

HasExpiresInDays returns a boolean if a field has been set.

### SetExpiresInDaysNil

`func (o *CreateApiKeyRequestDto) SetExpiresInDaysNil(b bool)`

 SetExpiresInDaysNil sets the value for ExpiresInDays to be an explicit nil

### UnsetExpiresInDays
`func (o *CreateApiKeyRequestDto) UnsetExpiresInDays()`

UnsetExpiresInDays ensures that no value is present for ExpiresInDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


