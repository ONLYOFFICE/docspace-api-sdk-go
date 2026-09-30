# ActiveConnectionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginEvent** | **int32** | The `id` of the item in `items` that the current request is authenticated by. It is `0` when the request  carried a token in the `Authorization` header instead of the portal cookie, and in that case none of the  items is the current connection. | 
**Items** | Pointer to [**[]ActiveConnectionsItemDto**](ActiveConnectionsItemDto.md) | One item per sign-in of the caller that is still active, ordered newest sign-in first, with the connection  the request itself uses moved to the front. Sign-ins older than a year are left out, and a caller with no  stored connection gets a single item describing the current request rather than an empty list. | [optional] 

## Methods

### NewActiveConnectionsDto

`func NewActiveConnectionsDto(loginEvent int32, ) *ActiveConnectionsDto`

NewActiveConnectionsDto instantiates a new ActiveConnectionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveConnectionsDtoWithDefaults

`func NewActiveConnectionsDtoWithDefaults() *ActiveConnectionsDto`

NewActiveConnectionsDtoWithDefaults instantiates a new ActiveConnectionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginEvent

`func (o *ActiveConnectionsDto) GetLoginEvent() int32`

GetLoginEvent returns the LoginEvent field if non-nil, zero value otherwise.

### GetLoginEventOk

`func (o *ActiveConnectionsDto) GetLoginEventOk() (*int32, bool)`

GetLoginEventOk returns a tuple with the LoginEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEvent

`func (o *ActiveConnectionsDto) SetLoginEvent(v int32)`

SetLoginEvent sets LoginEvent field to given value.


### GetItems

`func (o *ActiveConnectionsDto) GetItems() []ActiveConnectionsItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ActiveConnectionsDto) GetItemsOk() (*[]ActiveConnectionsItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ActiveConnectionsDto) SetItems(v []ActiveConnectionsItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *ActiveConnectionsDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *ActiveConnectionsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ActiveConnectionsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


