# SecurityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | The module the role applies to, given by its GUID. The all-zero GUID stands for the portal itself and grants  or revokes the DocSpace administrator role, which covers every module at once; a GUID that names no module  group is stored without effect rather than refused. | 
**UserId** | **string** | The portal member the role is given to or taken from, by user ID. The member has to exist already - nobody is  created here - and promoting a guest or a plain member turns them into a paid one. | 
**Administrator** | Pointer to **bool** | Which way the role goes: `true` adds the member to the module administrator group, `false` removes them from  it. Taking away the portal-wide role also drops the member from every product group. | [optional] 

## Methods

### NewSecurityRequestDto

`func NewSecurityRequestDto(productId string, userId string, ) *SecurityRequestDto`

NewSecurityRequestDto instantiates a new SecurityRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRequestDtoWithDefaults

`func NewSecurityRequestDtoWithDefaults() *SecurityRequestDto`

NewSecurityRequestDtoWithDefaults instantiates a new SecurityRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *SecurityRequestDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SecurityRequestDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SecurityRequestDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUserId

`func (o *SecurityRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SecurityRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SecurityRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAdministrator

`func (o *SecurityRequestDto) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *SecurityRequestDto) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *SecurityRequestDto) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *SecurityRequestDto) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


