# StorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The storage ID. | 
**Title** | **NullableString** | The storage title. | 
**Properties** | Pointer to [**[]AuthKey**](AuthKey.md) | The list of storage authentication keys. | [optional] 
**Current** | **bool** | Specifies if this is the current portal storage or not. | 
**IsSet** | **bool** | Specifies if this storage can be set or not. | 

## Methods

### NewStorageDto

`func NewStorageDto(id NullableString, title NullableString, current bool, isSet bool, ) *StorageDto`

NewStorageDto instantiates a new StorageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDtoWithDefaults

`func NewStorageDtoWithDefaults() *StorageDto`

NewStorageDtoWithDefaults instantiates a new StorageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *StorageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StorageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *StorageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StorageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StorageDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *StorageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *StorageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetProperties

`func (o *StorageDto) GetProperties() []AuthKey`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StorageDto) GetPropertiesOk() (*[]AuthKey, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StorageDto) SetProperties(v []AuthKey)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *StorageDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *StorageDto) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *StorageDto) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetCurrent

`func (o *StorageDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *StorageDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *StorageDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetIsSet

`func (o *StorageDto) GetIsSet() bool`

GetIsSet returns the IsSet field if non-nil, zero value otherwise.

### GetIsSetOk

`func (o *StorageDto) GetIsSetOk() (*bool, bool)`

GetIsSetOk returns a tuple with the IsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSet

`func (o *StorageDto) SetIsSet(v bool)`

SetIsSet sets IsSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


