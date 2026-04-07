# SecurityRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | The product ID for which permissions are being set. | 
**UserId** | **string** | The ID of the user whose permissions are being configured. | 
**Administrator** | Pointer to **bool** | Specifies whether the user has administrative privileges. | [optional] 

## Methods

### NewSecurityRequestsDto

`func NewSecurityRequestsDto(productId string, userId string, ) *SecurityRequestsDto`

NewSecurityRequestsDto instantiates a new SecurityRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRequestsDtoWithDefaults

`func NewSecurityRequestsDtoWithDefaults() *SecurityRequestsDto`

NewSecurityRequestsDtoWithDefaults instantiates a new SecurityRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *SecurityRequestsDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SecurityRequestsDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SecurityRequestsDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUserId

`func (o *SecurityRequestsDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SecurityRequestsDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SecurityRequestsDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAdministrator

`func (o *SecurityRequestsDto) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *SecurityRequestsDto) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *SecurityRequestsDto) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *SecurityRequestsDto) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


