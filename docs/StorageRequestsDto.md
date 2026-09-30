# StorageRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **NullableString** | The storage provider to switch to, by the identifier the matching listing operation reports - `default` for  the built-in local storage. The provider has to be available on the server, which that listing reports as  `isSet`, otherwise the request is refused with 400; sending the module already in use changes nothing. | 
**Props** | Pointer to [**[]ItemKeyValuePairStringString**](ItemKeyValuePairStringString.md) | The credentials the provider expects, as the name and value pairs it defines - a bucket, a region and an  access key for an Amazon S3 storage, for instance. Read the expected names from the entry of that provider in  the listing operation; they differ per provider, so there is no fixed set. | [optional] 

## Methods

### NewStorageRequestsDto

`func NewStorageRequestsDto(module NullableString, ) *StorageRequestsDto`

NewStorageRequestsDto instantiates a new StorageRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageRequestsDtoWithDefaults

`func NewStorageRequestsDtoWithDefaults() *StorageRequestsDto`

NewStorageRequestsDtoWithDefaults instantiates a new StorageRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *StorageRequestsDto) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *StorageRequestsDto) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *StorageRequestsDto) SetModule(v string)`

SetModule sets Module field to given value.


### SetModuleNil

`func (o *StorageRequestsDto) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *StorageRequestsDto) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetProps

`func (o *StorageRequestsDto) GetProps() []ItemKeyValuePairStringString`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *StorageRequestsDto) GetPropsOk() (*[]ItemKeyValuePairStringString, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *StorageRequestsDto) SetProps(v []ItemKeyValuePairStringString)`

SetProps sets Props field to given value.

### HasProps

`func (o *StorageRequestsDto) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *StorageRequestsDto) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *StorageRequestsDto) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


