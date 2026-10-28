# EditHistoryAuthorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The account the revision or the change is attributed to, as the editing service stored it. It is normally the  identifier of a portal account; the empty identifier stands for a change nobody could be named for. | 
**Name** | Pointer to **NullableString** | The display name of that account as the portal spells it now, which need not be the name that was stored with  the revision. An account that cannot be resolved - one removed from the portal, or a change made through an  anonymous link - is reported as a guest. | [optional] 

## Methods

### NewEditHistoryAuthorDto

`func NewEditHistoryAuthorDto(id NullableString, ) *EditHistoryAuthorDto`

NewEditHistoryAuthorDto instantiates a new EditHistoryAuthorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryAuthorDtoWithDefaults

`func NewEditHistoryAuthorDtoWithDefaults() *EditHistoryAuthorDto`

NewEditHistoryAuthorDtoWithDefaults instantiates a new EditHistoryAuthorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditHistoryAuthorDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditHistoryAuthorDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditHistoryAuthorDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *EditHistoryAuthorDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EditHistoryAuthorDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *EditHistoryAuthorDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditHistoryAuthorDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditHistoryAuthorDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditHistoryAuthorDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditHistoryAuthorDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditHistoryAuthorDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


