# ExternalSharingSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalShare** | Pointer to **bool** | Whether links that open a file or a room without a portal account may be created. While it is false the portal  also reports sharing on social networks as off and the default link type as internal, whatever was asked for. | [optional] 
**DefaultShareLinkInternal** | Pointer to **bool** | The kind of link the portal offers first: true means a link only accounts of this portal can open, false one  that anyone holding it can open. | [optional] 
**ExternalShareApplyToDocuments** | Pointer to **bool** | Whether the restriction covers personal documents. It only has an effect while external sharing is off, so a  true here with sharing allowed restricts nothing. | [optional] 
**ExternalShareApplyToRooms** | Pointer to **bool** | Whether the restriction covers rooms, including the creation of new public ones. It only has an effect while  external sharing is off. | [optional] 
**BlockExistingLinksOnRestrict** | Pointer to **bool** | Whether links created before the restriction stop opening as well. With false they keep working and only new  ones are refused. | [optional] 

## Methods

### NewExternalSharingSettingsDto

`func NewExternalSharingSettingsDto() *ExternalSharingSettingsDto`

NewExternalSharingSettingsDto instantiates a new ExternalSharingSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSharingSettingsDtoWithDefaults

`func NewExternalSharingSettingsDtoWithDefaults() *ExternalSharingSettingsDto`

NewExternalSharingSettingsDtoWithDefaults instantiates a new ExternalSharingSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalShare

`func (o *ExternalSharingSettingsDto) GetExternalShare() bool`

GetExternalShare returns the ExternalShare field if non-nil, zero value otherwise.

### GetExternalShareOk

`func (o *ExternalSharingSettingsDto) GetExternalShareOk() (*bool, bool)`

GetExternalShareOk returns a tuple with the ExternalShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShare

`func (o *ExternalSharingSettingsDto) SetExternalShare(v bool)`

SetExternalShare sets ExternalShare field to given value.

### HasExternalShare

`func (o *ExternalSharingSettingsDto) HasExternalShare() bool`

HasExternalShare returns a boolean if a field has been set.

### GetDefaultShareLinkInternal

`func (o *ExternalSharingSettingsDto) GetDefaultShareLinkInternal() bool`

GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field if non-nil, zero value otherwise.

### GetDefaultShareLinkInternalOk

`func (o *ExternalSharingSettingsDto) GetDefaultShareLinkInternalOk() (*bool, bool)`

GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareLinkInternal

`func (o *ExternalSharingSettingsDto) SetDefaultShareLinkInternal(v bool)`

SetDefaultShareLinkInternal sets DefaultShareLinkInternal field to given value.

### HasDefaultShareLinkInternal

`func (o *ExternalSharingSettingsDto) HasDefaultShareLinkInternal() bool`

HasDefaultShareLinkInternal returns a boolean if a field has been set.

### GetExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsDto) GetExternalShareApplyToDocuments() bool`

GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field if non-nil, zero value otherwise.

### GetExternalShareApplyToDocumentsOk

`func (o *ExternalSharingSettingsDto) GetExternalShareApplyToDocumentsOk() (*bool, bool)`

GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsDto) SetExternalShareApplyToDocuments(v bool)`

SetExternalShareApplyToDocuments sets ExternalShareApplyToDocuments field to given value.

### HasExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsDto) HasExternalShareApplyToDocuments() bool`

HasExternalShareApplyToDocuments returns a boolean if a field has been set.

### GetExternalShareApplyToRooms

`func (o *ExternalSharingSettingsDto) GetExternalShareApplyToRooms() bool`

GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field if non-nil, zero value otherwise.

### GetExternalShareApplyToRoomsOk

`func (o *ExternalSharingSettingsDto) GetExternalShareApplyToRoomsOk() (*bool, bool)`

GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToRooms

`func (o *ExternalSharingSettingsDto) SetExternalShareApplyToRooms(v bool)`

SetExternalShareApplyToRooms sets ExternalShareApplyToRooms field to given value.

### HasExternalShareApplyToRooms

`func (o *ExternalSharingSettingsDto) HasExternalShareApplyToRooms() bool`

HasExternalShareApplyToRooms returns a boolean if a field has been set.

### GetBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsDto) GetBlockExistingLinksOnRestrict() bool`

GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field if non-nil, zero value otherwise.

### GetBlockExistingLinksOnRestrictOk

`func (o *ExternalSharingSettingsDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool)`

GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsDto) SetBlockExistingLinksOnRestrict(v bool)`

SetBlockExistingLinksOnRestrict sets BlockExistingLinksOnRestrict field to given value.

### HasBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsDto) HasBlockExistingLinksOnRestrict() bool`

HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


