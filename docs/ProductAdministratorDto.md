# ProductAdministratorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | The module the verdict is about, echoed from the request. The all-zero GUID stands for the portal as a  whole rather than for any single module. | 
**UserId** | **string** | The user the verdict is about, echoed from the request unchanged - it is not checked for existing. | 
**Administrator** | **bool** | Whether that user administers that module. It is `true` for a DocSpace administrator whatever the module,  since the portal-wide role covers every one of them. A `false` can also mean the identifiers name no user  or no module at all, so it is not proof that the user exists, and it says nothing about whether the module  is enabled for the portal - `GET api/2.0/settings/security/{id}` reports that. | 

## Methods

### NewProductAdministratorDto

`func NewProductAdministratorDto(productId string, userId string, administrator bool, ) *ProductAdministratorDto`

NewProductAdministratorDto instantiates a new ProductAdministratorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAdministratorDtoWithDefaults

`func NewProductAdministratorDtoWithDefaults() *ProductAdministratorDto`

NewProductAdministratorDtoWithDefaults instantiates a new ProductAdministratorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductAdministratorDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductAdministratorDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductAdministratorDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUserId

`func (o *ProductAdministratorDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProductAdministratorDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProductAdministratorDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAdministrator

`func (o *ProductAdministratorDto) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *ProductAdministratorDto) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *ProductAdministratorDto) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


