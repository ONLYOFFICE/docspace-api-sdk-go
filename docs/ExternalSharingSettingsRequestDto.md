# ExternalSharingSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalShare** | Pointer to **bool** | Specifies whether external (public) link creation is allowed. | [optional] 
**DefaultShareLinkInternal** | Pointer to **bool** | Specifies the default sharing link type: true = DocSpace users only, false = Anyone with the link.  Relevant only when ExternalShare is true. | [optional] 
**ExternalShareApplyToDocuments** | Pointer to **bool** | When external sharing is restricted, specifies whether to apply the restriction to the My Documents section.  Relevant only when ExternalShare is false. | [optional] 
**ExternalShareApplyToRooms** | Pointer to **bool** | When external sharing is restricted, specifies whether to apply the restriction to the Rooms section.  Relevant only when ExternalShare is false. | [optional] 
**BlockExistingLinksOnRestrict** | Pointer to **bool** | When external sharing is restricted, specifies whether to block existing public links immediately.  Relevant only when ExternalShare is false. | [optional] 

## Methods

### NewExternalSharingSettingsRequestDto

`func NewExternalSharingSettingsRequestDto() *ExternalSharingSettingsRequestDto`

NewExternalSharingSettingsRequestDto instantiates a new ExternalSharingSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSharingSettingsRequestDtoWithDefaults

`func NewExternalSharingSettingsRequestDtoWithDefaults() *ExternalSharingSettingsRequestDto`

NewExternalSharingSettingsRequestDtoWithDefaults instantiates a new ExternalSharingSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalShare

`func (o *ExternalSharingSettingsRequestDto) GetExternalShare() bool`

GetExternalShare returns the ExternalShare field if non-nil, zero value otherwise.

### GetExternalShareOk

`func (o *ExternalSharingSettingsRequestDto) GetExternalShareOk() (*bool, bool)`

GetExternalShareOk returns a tuple with the ExternalShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShare

`func (o *ExternalSharingSettingsRequestDto) SetExternalShare(v bool)`

SetExternalShare sets ExternalShare field to given value.

### HasExternalShare

`func (o *ExternalSharingSettingsRequestDto) HasExternalShare() bool`

HasExternalShare returns a boolean if a field has been set.

### GetDefaultShareLinkInternal

`func (o *ExternalSharingSettingsRequestDto) GetDefaultShareLinkInternal() bool`

GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field if non-nil, zero value otherwise.

### GetDefaultShareLinkInternalOk

`func (o *ExternalSharingSettingsRequestDto) GetDefaultShareLinkInternalOk() (*bool, bool)`

GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareLinkInternal

`func (o *ExternalSharingSettingsRequestDto) SetDefaultShareLinkInternal(v bool)`

SetDefaultShareLinkInternal sets DefaultShareLinkInternal field to given value.

### HasDefaultShareLinkInternal

`func (o *ExternalSharingSettingsRequestDto) HasDefaultShareLinkInternal() bool`

HasDefaultShareLinkInternal returns a boolean if a field has been set.

### GetExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToDocuments() bool`

GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field if non-nil, zero value otherwise.

### GetExternalShareApplyToDocumentsOk

`func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToDocumentsOk() (*bool, bool)`

GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsRequestDto) SetExternalShareApplyToDocuments(v bool)`

SetExternalShareApplyToDocuments sets ExternalShareApplyToDocuments field to given value.

### HasExternalShareApplyToDocuments

`func (o *ExternalSharingSettingsRequestDto) HasExternalShareApplyToDocuments() bool`

HasExternalShareApplyToDocuments returns a boolean if a field has been set.

### GetExternalShareApplyToRooms

`func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToRooms() bool`

GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field if non-nil, zero value otherwise.

### GetExternalShareApplyToRoomsOk

`func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToRoomsOk() (*bool, bool)`

GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToRooms

`func (o *ExternalSharingSettingsRequestDto) SetExternalShareApplyToRooms(v bool)`

SetExternalShareApplyToRooms sets ExternalShareApplyToRooms field to given value.

### HasExternalShareApplyToRooms

`func (o *ExternalSharingSettingsRequestDto) HasExternalShareApplyToRooms() bool`

HasExternalShareApplyToRooms returns a boolean if a field has been set.

### GetBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsRequestDto) GetBlockExistingLinksOnRestrict() bool`

GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field if non-nil, zero value otherwise.

### GetBlockExistingLinksOnRestrictOk

`func (o *ExternalSharingSettingsRequestDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool)`

GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsRequestDto) SetBlockExistingLinksOnRestrict(v bool)`

SetBlockExistingLinksOnRestrict sets BlockExistingLinksOnRestrict field to given value.

### HasBlockExistingLinksOnRestrict

`func (o *ExternalSharingSettingsRequestDto) HasBlockExistingLinksOnRestrict() bool`

HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


