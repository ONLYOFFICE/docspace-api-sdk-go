# ThirdPartyFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The name shown for the entry. For a file it carries the extension, which is how the format is recognised, and  for a room it is the room name. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | The level the calling account holds on this entry, resolved from its own rights, the groups it belongs to and  any link it came in through. It is the level itself, not what the account may do with it - the action flags  below answer that. | [optional] 
**SharedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | Who gave the calling account the access it is using. It is filled in only while the entry is being read  through a share, and never for a caller without an account. | [optional] 
**OwnedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | Who owns the place the entry is shared from - the creator of the room it lies in, or of the personal section  that holds it. It is filled in only while the entry is being read through a share, and never for a caller  without an account. | [optional] 
**Shared** | Pointer to **bool** | Whether at least one external link exists for the entry, whichever kind. It says nothing about accounts and  groups - those are counted by the flag for members below. | [optional] 
**SharedForUser** | Pointer to **bool** | Whether at least one account or group has been given rights on the entry directly, as opposed to reaching it  through a link or through the room around it. | [optional] 
**SharedExternal** | Pointer to **bool** | Whether one of the entry's links is open to people outside the portal, as opposed to a link that only its own  members can follow. This is the flag to watch when the concern is who can reach the content from outside. | [optional] 
**ParentShared** | Pointer to **bool** | Whether the entry is reachable because the room or folder around it is shared, rather than through rights of  its own. A copy or a move takes the entry out of that scope. | [optional] 
**ShortWebUrl** | Pointer to **string** | A shortened address that opens the entry through the link it is being read with. It is an empty string  whenever no link applies, which is the usual case for a member browsing their own rooms. | [optional] 
**Created** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the entry was created, written with the offset of the portal's time zone. For a file restored from an  older version this is still the moment the file first appeared. | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | Who created the entry. It is null for a caller without an account, who is told nothing about the portal's  members. | [optional] 
**Updated** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the entry last changed, written with the offset of the portal's time zone. It is never reported as  earlier than the creation moment, so the two can be compared safely. | [optional] 
**AutoDelete** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the entry will disappear on its own, written with the offset of the portal's time zone. It is filled in  only where a removal is actually scheduled - something in the trash while the portal cleans it up  automatically, or a guest's own documents - so a null means nothing is scheduled rather than that the entry is  permanent. | [optional] 
**RootFolderType** | Pointer to [**FolderType**](FolderType.md) | The section the entry ultimately belongs to, which is what tells a personal document from one inside a room,  from a template and from something in the trash or the archive. | [optional] 
**ParentRoomType** | Pointer to [**FolderType**](FolderType.md) | The kind of room the entry lies in, which decides what the room allows - filling forms, public links,  indexing. It is null for an entry that is not inside a room at all. | [optional] 
**UpdatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | Who changed the entry last. It is null for a caller without an account. | [optional] 
**ProviderItem** | Pointer to **bool** | Set when the entry is stored on a connected third-party account rather than on the portal, and null when it is  stored on the portal. Such an entry is identified by a string rather than a number, and some operations skip  it. | [optional] 
**ProviderKey** | Pointer to **string** | Which third-party service holds the entry, matching the keys accepted by the third-party operations. It is  null for an entry stored on the portal. | [optional] 
**ProviderId** | Pointer to **int32** | The connected account the entry comes from, for telling apart two connections to the same service. It is null  for an entry stored on the portal. | [optional] 
**Order** | Pointer to **string** | The place of the entry in a room where the members arrange the content themselves, given as the position of  the entry preceded by the positions of the folders leading to it, separated by dots. It is empty when nothing  has been arranged. | [optional] 
**IsFavorite** | Pointer to **bool** | Set when the calling account has marked the entry as a favorite, which is what puts it into the favorites  listing. For a file that is not marked it is null rather than false. | [optional] 
**FileEntryType** | Pointer to [**FileEntryType**](FileEntryType.md) | Tells a folder from a file, and so which of the two shapes the rest of the object has. A room is reported as a  folder here. | [optional] 
**Id** | Pointer to **string** | The identifier to pass back to the other operations of this entry. It is a number for storage on the portal  and a string for a connected third-party account, and it is unique only within its own kind, so files and  folders may carry the same value. | [optional] 
**RootFolderId** | Pointer to **string** | The section the entry ultimately lies in, as an identifier that can be listed like any other folder. For an  entry inside a room this is the rooms section, not the room. | [optional] 
**OriginId** | Pointer to **string** | The folder the entry was deleted from, which is where restoring it puts it back. It is left out of the answer  unless the entry is in the trash. | [optional] 
**OriginRoomId** | Pointer to **string** | The room the entry was deleted from, left out of the answer for anything that was not deleted out of a room. | [optional] 
**OriginTitle** | Pointer to **string** | The name of the folder the entry was deleted from, for showing where it would be restored to. It is null for  an entry that is not in the trash. | [optional] 
**OriginRoomTitle** | Pointer to **string** | The name of the room the entry was deleted from, null for anything that was not deleted out of a room. | [optional] 
**CanShare** | Pointer to **bool** | Whether the calling account may change who has access to the entry, and so whether offering a sharing dialog  for it makes sense. It is false in rooms whose access is fixed by the room itself, such as a private one, even  for its manager. | [optional] 
**ShareSettings** | Pointer to [**NullableAiFileEntryDtoAllOfShareSettings**](AiFileEntryDtoAllOfShareSettings.md) |  | [optional] 
**Security** | Pointer to [**NullableAiFileEntryDtoAllOfSecurity**](AiFileEntryDtoAllOfSecurity.md) |  | [optional] 
**AvailableShareRights** | Pointer to [**NullableAiFileEntryDtoAllOfAvailableShareRights**](AiFileEntryDtoAllOfAvailableShareRights.md) |  | [optional] 
**RequestToken** | Pointer to **string** | The token of the link the entry is being read through, which is the value the external-share operations expect  and which also has to be carried by the download and preview addresses. It is null whenever the entry is not  being read through a link. | [optional] 
**External** | Pointer to **bool** | Set when the link being used was made for this very entry, and false when the entry is reached through a link  to the room around it. It is null when no link is involved. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the link being used stops working, written with the offset of the portal's time zone. It is null for a  link that never expires and whenever no link is involved. | [optional] 
**IsLinkExpired** | Pointer to **bool** | Set when the link being used has already passed its expiration date, which is why the entry cannot be opened  even though it is described here. It is null when no link is involved. | [optional] 
**FolderId** | Pointer to **NullableString** | The folder the file is stored in. When the file was reached through a share and the caller cannot open its  real parent, the identifier of the Shared with me section is reported instead, so this is where the file is  visible rather than where it physically sits. | [optional] 
**Version** | Pointer to **int32** | The revision this entry describes. It starts at 1 and moves to the next number each time new content is stored  over the file, except for an editing session opened against the file itself, which replaces the content and  keeps the number. `GET api/2.0/files/file/{fileId}/history` lists them all. | [optional] 
**VersionGroup** | Pointer to **int32** | Groups revisions that belong together, which is how a history can fold a long editing session into one entry:  versions saved inside one session share this number, and an upload over the file starts a new group. | [optional] 
**ContentLength** | Pointer to **NullableString** | The size already formatted for display, with a unit and the separators of the caller's language. Read  `pureContentLength` for a number to calculate with. | [optional] 
**PureContentLength** | Pointer to **NullableInt64** | The size of the stored content in bytes, and null for an empty file. | [optional] 
**FileStatus** | Pointer to [**FileStatus**](FileStatus.md) | What the portal is currently doing with the file and how the caller stands towards it - open in the editor,  unread, being converted, and so on. The value is a bit mask that combines those states, so a file can report a  number that matches none of the published members on its own. | [optional] 
**EditingBy** | Pointer to **map[string]string** | The accounts that have the file open in the editor at this moment, as account identifier to display name, and  empty when nobody has. The all-zero identifier stands for people who came in through an external link without  signing in, and its name carries their number in brackets when there is more than one. | [optional] 
**Mute** | Pointer to **bool** | Not a property of the file at all: it repeats, inverted, the calling account's own switch for new-item badges,  so it is the same in every entry of one answer. True means that account has badges turned off. | [optional] 
**ViewUrl** | Pointer to **NullableString** | The address that returns the bytes of the file - a download, in spite of the name; `webUrl` is the address a  person opens. When the file was reached through an external link the address carries the key of that link, so  it keeps working without signing in. | [optional] 
**WebUrl** | Pointer to **NullableString** | The page that opens the file in a browser: the editor for a format the portal edits, the media viewer for  pictures, audio and video, and the download address for a format it cannot show at all. | [optional] 
**FileType** | Pointer to [**FileType**](FileType.md) | The broad kind of content, worked out from the extension, which is what a client uses to pick an icon or a  viewer without parsing `fileExst` itself. | [optional] 
**FileExst** | Pointer to **NullableString** | The extension of the stored file, leading dot included and always lower case. For a format the portal keeps in  a converted shape this is the extension it is served under, not the one it was uploaded with. | [optional] 
**Comment** | Pointer to **NullableString** | The note kept with this revision. The portal writes it itself for revisions it creates, an upload over an  existing file among them, and an editor stores the note a person typed when saving a version. | [optional] 
**Encrypted** | Pointer to **NullableBool** | True for a file in a private room, whose content the server never sees and which therefore cannot be converted  or taken over by an upload. Null, rather than false, for an ordinary file. | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | The address of the generated preview image. It is filled in only while `thumbnailStatus` says the preview has  been created, and it carries a suffix that changes with the file, so an image cached for an earlier revision  is not reused. | [optional] 
**ThumbnailStatus** | Pointer to [**Thumbnail**](Thumbnail.md) | How far the preview image has got. Only the created state means `thumbnailUrl` holds an address; the others  mean there is none, either because it is still being produced or because this format has no preview. | [optional] 
**Locked** | Pointer to **NullableBool** | True while the file is held under a lock that stops anyone but its holder from editing it, and null rather  than false when there is no lock. `lockedBy` names the holder unless the caller is the holder. | [optional] 
**LockedBy** | Pointer to **NullableString** | The display name of the account holding the lock, and null when the caller holds it - so `locked` true  together with no name here means the lock is the caller's own. | [optional] 
**HasDraft** | Pointer to **NullableBool** | For a fillable PDF form, whether the caller already has a filling draft of it, in which case `draftLocation`  says where that draft lives. Null for anything that is not a form. | [optional] 
**FormFillingStatus** | Pointer to [**FormFillingStatus**](FormFillingStatus.md) | How far the filling of this form has got for the calling account, and whose turn it is now. It is worked out  only inside a virtual data room, where filling runs in steps; everywhere else it stays at the none value. | [optional] 
**IsForm** | Pointer to **NullableBool** | Whether the file is a PDF, and so offered as a fillable form. It is null for any other file type. | [optional] 
**CustomFilterEnabled** | Pointer to **NullableBool** | True while a spreadsheet is in the mode where each person sorts and filters their own view without changing  what the others see, and null rather than false when it is not. | [optional] 
**CustomFilterEnabledBy** | Pointer to **NullableString** | The display name of the account that turned that mode on, and null when the caller turned it on themselves. | [optional] 
**StartFilling** | Pointer to **NullableBool** | For a form in a room for filling, whether it has been released for filling; until then it is still being  prepared and only the people running the room work with it. Null for a file this does not apply to. | [optional] 
**IsFillingPreparing** | Pointer to **NullableBool** | True during the short window in which a released form is still being written out by the editor. Neither  filling nor editing is accepted while it lasts, so a client should wait and read the file again. | [optional] 
**InProcessFolderId** | Pointer to **NullableInt32** | Left empty by the portal: the folder holding the caller's draft is reported in `draftLocation` instead. | [optional] 
**InProcessFolderTitle** | Pointer to **NullableString** | Left empty by the portal, like the identifier beside it; the draft's folder is named in `draftLocation`. | [optional] 
**ResultsFolderId** | Pointer to **NullableInt32** | The folder that collects the completed copies of this form. It is filled in only for the original form of a  room for filling, and only for a caller allowed to work with that form; null everywhere else. | [optional] 
**DraftLocation** | Pointer to [**ThirdPartyDraftLocation**](ThirdPartyDraftLocation.md) | Where the caller's own filling draft of this form is kept. Null when there is no draft yet, which is the same  thing `hasDraft` reports. | [optional] 
**ViewAccessibility** | Pointer to [**NullableFileDtoAllOfViewAccessibility**](FileDtoAllOfViewAccessibility.md) |  | [optional] 
**LastOpened** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the caller last opened the file. It is kept per account and is what orders the Recent section, so  it is null for a file this account has never opened. Written with the offset of the portal's time zone. | [optional] 
**Expired** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the file falls under the lifetime rule of the room holding it and is removed. It is counted from  the first revision rather than the latest one, so editing a file does not postpone it, and it is null when the  room sets no lifetime. Written with the offset of the portal's time zone. | [optional] 
**VectorizationStatus** | Pointer to [**VectorizationStatus**](VectorizationStatus.md) | How far the indexing of the file's content for AI search has got. It is null for a file that has never been  queued for indexing, which is every file while the feature is off for the portal. | [optional] 
**ExternalDbTableName** | Pointer to **NullableString** | The table collecting the submitted values of this form in the external database configured for its room. The  field is left out of the answer entirely when the form has no such table. | [optional] 
**Dimensions** | Pointer to [**Size**](Size.md) | The pixel size of the picture, measured by reading the stored file rather than taken from any stored metadata.  Null for anything that is not a picture the portal can show, and also when the file could not be read. | [optional] 

## Methods

### NewThirdPartyFileDto

`func NewThirdPartyFileDto() *ThirdPartyFileDto`

NewThirdPartyFileDto instantiates a new ThirdPartyFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyFileDtoWithDefaults

`func NewThirdPartyFileDtoWithDefaults() *ThirdPartyFileDto`

NewThirdPartyFileDtoWithDefaults instantiates a new ThirdPartyFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ThirdPartyFileDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ThirdPartyFileDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ThirdPartyFileDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ThirdPartyFileDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *ThirdPartyFileDto) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ThirdPartyFileDto) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ThirdPartyFileDto) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ThirdPartyFileDto) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *ThirdPartyFileDto) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *ThirdPartyFileDto) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *ThirdPartyFileDto) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *ThirdPartyFileDto) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ThirdPartyFileDto) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ThirdPartyFileDto) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ThirdPartyFileDto) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ThirdPartyFileDto) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *ThirdPartyFileDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ThirdPartyFileDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ThirdPartyFileDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ThirdPartyFileDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *ThirdPartyFileDto) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *ThirdPartyFileDto) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *ThirdPartyFileDto) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *ThirdPartyFileDto) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *ThirdPartyFileDto) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *ThirdPartyFileDto) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *ThirdPartyFileDto) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *ThirdPartyFileDto) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *ThirdPartyFileDto) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *ThirdPartyFileDto) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *ThirdPartyFileDto) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *ThirdPartyFileDto) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *ThirdPartyFileDto) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *ThirdPartyFileDto) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *ThirdPartyFileDto) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *ThirdPartyFileDto) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *ThirdPartyFileDto) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThirdPartyFileDto) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThirdPartyFileDto) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThirdPartyFileDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ThirdPartyFileDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ThirdPartyFileDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ThirdPartyFileDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ThirdPartyFileDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *ThirdPartyFileDto) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ThirdPartyFileDto) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ThirdPartyFileDto) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ThirdPartyFileDto) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *ThirdPartyFileDto) GetAutoDelete() ApiDateTime`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *ThirdPartyFileDto) GetAutoDeleteOk() (*ApiDateTime, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *ThirdPartyFileDto) SetAutoDelete(v ApiDateTime)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *ThirdPartyFileDto) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *ThirdPartyFileDto) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *ThirdPartyFileDto) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *ThirdPartyFileDto) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *ThirdPartyFileDto) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *ThirdPartyFileDto) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *ThirdPartyFileDto) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *ThirdPartyFileDto) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *ThirdPartyFileDto) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ThirdPartyFileDto) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ThirdPartyFileDto) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ThirdPartyFileDto) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ThirdPartyFileDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *ThirdPartyFileDto) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *ThirdPartyFileDto) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *ThirdPartyFileDto) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *ThirdPartyFileDto) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *ThirdPartyFileDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyFileDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyFileDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *ThirdPartyFileDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *ThirdPartyFileDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ThirdPartyFileDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ThirdPartyFileDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ThirdPartyFileDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *ThirdPartyFileDto) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ThirdPartyFileDto) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ThirdPartyFileDto) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ThirdPartyFileDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *ThirdPartyFileDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ThirdPartyFileDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ThirdPartyFileDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ThirdPartyFileDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *ThirdPartyFileDto) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *ThirdPartyFileDto) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *ThirdPartyFileDto) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *ThirdPartyFileDto) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *ThirdPartyFileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyFileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyFileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartyFileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *ThirdPartyFileDto) GetRootFolderId() string`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *ThirdPartyFileDto) GetRootFolderIdOk() (*string, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *ThirdPartyFileDto) SetRootFolderId(v string)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *ThirdPartyFileDto) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *ThirdPartyFileDto) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *ThirdPartyFileDto) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *ThirdPartyFileDto) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *ThirdPartyFileDto) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *ThirdPartyFileDto) GetOriginRoomId() string`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *ThirdPartyFileDto) GetOriginRoomIdOk() (*string, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *ThirdPartyFileDto) SetOriginRoomId(v string)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *ThirdPartyFileDto) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *ThirdPartyFileDto) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *ThirdPartyFileDto) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *ThirdPartyFileDto) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *ThirdPartyFileDto) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### GetOriginRoomTitle

`func (o *ThirdPartyFileDto) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *ThirdPartyFileDto) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *ThirdPartyFileDto) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *ThirdPartyFileDto) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### GetCanShare

`func (o *ThirdPartyFileDto) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *ThirdPartyFileDto) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *ThirdPartyFileDto) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *ThirdPartyFileDto) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *ThirdPartyFileDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *ThirdPartyFileDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *ThirdPartyFileDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *ThirdPartyFileDto) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *ThirdPartyFileDto) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *ThirdPartyFileDto) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *ThirdPartyFileDto) GetSecurity() AiFileEntryDtoAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ThirdPartyFileDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ThirdPartyFileDto) SetSecurity(v AiFileEntryDtoAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ThirdPartyFileDto) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *ThirdPartyFileDto) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *ThirdPartyFileDto) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *ThirdPartyFileDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *ThirdPartyFileDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *ThirdPartyFileDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *ThirdPartyFileDto) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *ThirdPartyFileDto) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *ThirdPartyFileDto) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *ThirdPartyFileDto) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *ThirdPartyFileDto) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *ThirdPartyFileDto) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *ThirdPartyFileDto) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetExternal

`func (o *ThirdPartyFileDto) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ThirdPartyFileDto) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ThirdPartyFileDto) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ThirdPartyFileDto) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ThirdPartyFileDto) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ThirdPartyFileDto) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ThirdPartyFileDto) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ThirdPartyFileDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *ThirdPartyFileDto) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *ThirdPartyFileDto) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *ThirdPartyFileDto) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *ThirdPartyFileDto) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### GetFolderId

`func (o *ThirdPartyFileDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ThirdPartyFileDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ThirdPartyFileDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ThirdPartyFileDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ThirdPartyFileDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ThirdPartyFileDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetVersion

`func (o *ThirdPartyFileDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdPartyFileDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdPartyFileDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdPartyFileDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *ThirdPartyFileDto) GetVersionGroup() int32`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *ThirdPartyFileDto) GetVersionGroupOk() (*int32, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *ThirdPartyFileDto) SetVersionGroup(v int32)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *ThirdPartyFileDto) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetContentLength

`func (o *ThirdPartyFileDto) GetContentLength() string`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *ThirdPartyFileDto) GetContentLengthOk() (*string, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *ThirdPartyFileDto) SetContentLength(v string)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *ThirdPartyFileDto) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *ThirdPartyFileDto) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *ThirdPartyFileDto) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetPureContentLength

`func (o *ThirdPartyFileDto) GetPureContentLength() int64`

GetPureContentLength returns the PureContentLength field if non-nil, zero value otherwise.

### GetPureContentLengthOk

`func (o *ThirdPartyFileDto) GetPureContentLengthOk() (*int64, bool)`

GetPureContentLengthOk returns a tuple with the PureContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureContentLength

`func (o *ThirdPartyFileDto) SetPureContentLength(v int64)`

SetPureContentLength sets PureContentLength field to given value.

### HasPureContentLength

`func (o *ThirdPartyFileDto) HasPureContentLength() bool`

HasPureContentLength returns a boolean if a field has been set.

### SetPureContentLengthNil

`func (o *ThirdPartyFileDto) SetPureContentLengthNil(b bool)`

 SetPureContentLengthNil sets the value for PureContentLength to be an explicit nil

### UnsetPureContentLength
`func (o *ThirdPartyFileDto) UnsetPureContentLength()`

UnsetPureContentLength ensures that no value is present for PureContentLength, not even an explicit nil
### GetFileStatus

`func (o *ThirdPartyFileDto) GetFileStatus() FileStatus`

GetFileStatus returns the FileStatus field if non-nil, zero value otherwise.

### GetFileStatusOk

`func (o *ThirdPartyFileDto) GetFileStatusOk() (*FileStatus, bool)`

GetFileStatusOk returns a tuple with the FileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStatus

`func (o *ThirdPartyFileDto) SetFileStatus(v FileStatus)`

SetFileStatus sets FileStatus field to given value.

### HasFileStatus

`func (o *ThirdPartyFileDto) HasFileStatus() bool`

HasFileStatus returns a boolean if a field has been set.

### GetEditingBy

`func (o *ThirdPartyFileDto) GetEditingBy() map[string]*string`

GetEditingBy returns the EditingBy field if non-nil, zero value otherwise.

### GetEditingByOk

`func (o *ThirdPartyFileDto) GetEditingByOk() (*map[string]*string, bool)`

GetEditingByOk returns a tuple with the EditingBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditingBy

`func (o *ThirdPartyFileDto) SetEditingBy(v map[string]*string)`

SetEditingBy sets EditingBy field to given value.

### HasEditingBy

`func (o *ThirdPartyFileDto) HasEditingBy() bool`

HasEditingBy returns a boolean if a field has been set.

### GetMute

`func (o *ThirdPartyFileDto) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *ThirdPartyFileDto) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *ThirdPartyFileDto) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *ThirdPartyFileDto) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetViewUrl

`func (o *ThirdPartyFileDto) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *ThirdPartyFileDto) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *ThirdPartyFileDto) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *ThirdPartyFileDto) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### SetViewUrlNil

`func (o *ThirdPartyFileDto) SetViewUrlNil(b bool)`

 SetViewUrlNil sets the value for ViewUrl to be an explicit nil

### UnsetViewUrl
`func (o *ThirdPartyFileDto) UnsetViewUrl()`

UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
### GetWebUrl

`func (o *ThirdPartyFileDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ThirdPartyFileDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ThirdPartyFileDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ThirdPartyFileDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *ThirdPartyFileDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *ThirdPartyFileDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFileType

`func (o *ThirdPartyFileDto) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ThirdPartyFileDto) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ThirdPartyFileDto) SetFileType(v FileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ThirdPartyFileDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileExst

`func (o *ThirdPartyFileDto) GetFileExst() string`

GetFileExst returns the FileExst field if non-nil, zero value otherwise.

### GetFileExstOk

`func (o *ThirdPartyFileDto) GetFileExstOk() (*string, bool)`

GetFileExstOk returns a tuple with the FileExst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExst

`func (o *ThirdPartyFileDto) SetFileExst(v string)`

SetFileExst sets FileExst field to given value.

### HasFileExst

`func (o *ThirdPartyFileDto) HasFileExst() bool`

HasFileExst returns a boolean if a field has been set.

### SetFileExstNil

`func (o *ThirdPartyFileDto) SetFileExstNil(b bool)`

 SetFileExstNil sets the value for FileExst to be an explicit nil

### UnsetFileExst
`func (o *ThirdPartyFileDto) UnsetFileExst()`

UnsetFileExst ensures that no value is present for FileExst, not even an explicit nil
### GetComment

`func (o *ThirdPartyFileDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ThirdPartyFileDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ThirdPartyFileDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ThirdPartyFileDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ThirdPartyFileDto) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ThirdPartyFileDto) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEncrypted

`func (o *ThirdPartyFileDto) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *ThirdPartyFileDto) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *ThirdPartyFileDto) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *ThirdPartyFileDto) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### SetEncryptedNil

`func (o *ThirdPartyFileDto) SetEncryptedNil(b bool)`

 SetEncryptedNil sets the value for Encrypted to be an explicit nil

### UnsetEncrypted
`func (o *ThirdPartyFileDto) UnsetEncrypted()`

UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
### GetThumbnailUrl

`func (o *ThirdPartyFileDto) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ThirdPartyFileDto) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ThirdPartyFileDto) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ThirdPartyFileDto) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *ThirdPartyFileDto) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *ThirdPartyFileDto) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailStatus

`func (o *ThirdPartyFileDto) GetThumbnailStatus() Thumbnail`

GetThumbnailStatus returns the ThumbnailStatus field if non-nil, zero value otherwise.

### GetThumbnailStatusOk

`func (o *ThirdPartyFileDto) GetThumbnailStatusOk() (*Thumbnail, bool)`

GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailStatus

`func (o *ThirdPartyFileDto) SetThumbnailStatus(v Thumbnail)`

SetThumbnailStatus sets ThumbnailStatus field to given value.

### HasThumbnailStatus

`func (o *ThirdPartyFileDto) HasThumbnailStatus() bool`

HasThumbnailStatus returns a boolean if a field has been set.

### GetLocked

`func (o *ThirdPartyFileDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ThirdPartyFileDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ThirdPartyFileDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ThirdPartyFileDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *ThirdPartyFileDto) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *ThirdPartyFileDto) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetLockedBy

`func (o *ThirdPartyFileDto) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ThirdPartyFileDto) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ThirdPartyFileDto) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ThirdPartyFileDto) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### SetLockedByNil

`func (o *ThirdPartyFileDto) SetLockedByNil(b bool)`

 SetLockedByNil sets the value for LockedBy to be an explicit nil

### UnsetLockedBy
`func (o *ThirdPartyFileDto) UnsetLockedBy()`

UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
### GetHasDraft

`func (o *ThirdPartyFileDto) GetHasDraft() bool`

GetHasDraft returns the HasDraft field if non-nil, zero value otherwise.

### GetHasDraftOk

`func (o *ThirdPartyFileDto) GetHasDraftOk() (*bool, bool)`

GetHasDraftOk returns a tuple with the HasDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDraft

`func (o *ThirdPartyFileDto) SetHasDraft(v bool)`

SetHasDraft sets HasDraft field to given value.

### HasHasDraft

`func (o *ThirdPartyFileDto) HasHasDraft() bool`

HasHasDraft returns a boolean if a field has been set.

### SetHasDraftNil

`func (o *ThirdPartyFileDto) SetHasDraftNil(b bool)`

 SetHasDraftNil sets the value for HasDraft to be an explicit nil

### UnsetHasDraft
`func (o *ThirdPartyFileDto) UnsetHasDraft()`

UnsetHasDraft ensures that no value is present for HasDraft, not even an explicit nil
### GetFormFillingStatus

`func (o *ThirdPartyFileDto) GetFormFillingStatus() FormFillingStatus`

GetFormFillingStatus returns the FormFillingStatus field if non-nil, zero value otherwise.

### GetFormFillingStatusOk

`func (o *ThirdPartyFileDto) GetFormFillingStatusOk() (*FormFillingStatus, bool)`

GetFormFillingStatusOk returns a tuple with the FormFillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFillingStatus

`func (o *ThirdPartyFileDto) SetFormFillingStatus(v FormFillingStatus)`

SetFormFillingStatus sets FormFillingStatus field to given value.

### HasFormFillingStatus

`func (o *ThirdPartyFileDto) HasFormFillingStatus() bool`

HasFormFillingStatus returns a boolean if a field has been set.

### GetIsForm

`func (o *ThirdPartyFileDto) GetIsForm() bool`

GetIsForm returns the IsForm field if non-nil, zero value otherwise.

### GetIsFormOk

`func (o *ThirdPartyFileDto) GetIsFormOk() (*bool, bool)`

GetIsFormOk returns a tuple with the IsForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForm

`func (o *ThirdPartyFileDto) SetIsForm(v bool)`

SetIsForm sets IsForm field to given value.

### HasIsForm

`func (o *ThirdPartyFileDto) HasIsForm() bool`

HasIsForm returns a boolean if a field has been set.

### SetIsFormNil

`func (o *ThirdPartyFileDto) SetIsFormNil(b bool)`

 SetIsFormNil sets the value for IsForm to be an explicit nil

### UnsetIsForm
`func (o *ThirdPartyFileDto) UnsetIsForm()`

UnsetIsForm ensures that no value is present for IsForm, not even an explicit nil
### GetCustomFilterEnabled

`func (o *ThirdPartyFileDto) GetCustomFilterEnabled() bool`

GetCustomFilterEnabled returns the CustomFilterEnabled field if non-nil, zero value otherwise.

### GetCustomFilterEnabledOk

`func (o *ThirdPartyFileDto) GetCustomFilterEnabledOk() (*bool, bool)`

GetCustomFilterEnabledOk returns a tuple with the CustomFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabled

`func (o *ThirdPartyFileDto) SetCustomFilterEnabled(v bool)`

SetCustomFilterEnabled sets CustomFilterEnabled field to given value.

### HasCustomFilterEnabled

`func (o *ThirdPartyFileDto) HasCustomFilterEnabled() bool`

HasCustomFilterEnabled returns a boolean if a field has been set.

### SetCustomFilterEnabledNil

`func (o *ThirdPartyFileDto) SetCustomFilterEnabledNil(b bool)`

 SetCustomFilterEnabledNil sets the value for CustomFilterEnabled to be an explicit nil

### UnsetCustomFilterEnabled
`func (o *ThirdPartyFileDto) UnsetCustomFilterEnabled()`

UnsetCustomFilterEnabled ensures that no value is present for CustomFilterEnabled, not even an explicit nil
### GetCustomFilterEnabledBy

`func (o *ThirdPartyFileDto) GetCustomFilterEnabledBy() string`

GetCustomFilterEnabledBy returns the CustomFilterEnabledBy field if non-nil, zero value otherwise.

### GetCustomFilterEnabledByOk

`func (o *ThirdPartyFileDto) GetCustomFilterEnabledByOk() (*string, bool)`

GetCustomFilterEnabledByOk returns a tuple with the CustomFilterEnabledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabledBy

`func (o *ThirdPartyFileDto) SetCustomFilterEnabledBy(v string)`

SetCustomFilterEnabledBy sets CustomFilterEnabledBy field to given value.

### HasCustomFilterEnabledBy

`func (o *ThirdPartyFileDto) HasCustomFilterEnabledBy() bool`

HasCustomFilterEnabledBy returns a boolean if a field has been set.

### SetCustomFilterEnabledByNil

`func (o *ThirdPartyFileDto) SetCustomFilterEnabledByNil(b bool)`

 SetCustomFilterEnabledByNil sets the value for CustomFilterEnabledBy to be an explicit nil

### UnsetCustomFilterEnabledBy
`func (o *ThirdPartyFileDto) UnsetCustomFilterEnabledBy()`

UnsetCustomFilterEnabledBy ensures that no value is present for CustomFilterEnabledBy, not even an explicit nil
### GetStartFilling

`func (o *ThirdPartyFileDto) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *ThirdPartyFileDto) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *ThirdPartyFileDto) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *ThirdPartyFileDto) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *ThirdPartyFileDto) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *ThirdPartyFileDto) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetIsFillingPreparing

`func (o *ThirdPartyFileDto) GetIsFillingPreparing() bool`

GetIsFillingPreparing returns the IsFillingPreparing field if non-nil, zero value otherwise.

### GetIsFillingPreparingOk

`func (o *ThirdPartyFileDto) GetIsFillingPreparingOk() (*bool, bool)`

GetIsFillingPreparingOk returns a tuple with the IsFillingPreparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFillingPreparing

`func (o *ThirdPartyFileDto) SetIsFillingPreparing(v bool)`

SetIsFillingPreparing sets IsFillingPreparing field to given value.

### HasIsFillingPreparing

`func (o *ThirdPartyFileDto) HasIsFillingPreparing() bool`

HasIsFillingPreparing returns a boolean if a field has been set.

### SetIsFillingPreparingNil

`func (o *ThirdPartyFileDto) SetIsFillingPreparingNil(b bool)`

 SetIsFillingPreparingNil sets the value for IsFillingPreparing to be an explicit nil

### UnsetIsFillingPreparing
`func (o *ThirdPartyFileDto) UnsetIsFillingPreparing()`

UnsetIsFillingPreparing ensures that no value is present for IsFillingPreparing, not even an explicit nil
### GetInProcessFolderId

`func (o *ThirdPartyFileDto) GetInProcessFolderId() int32`

GetInProcessFolderId returns the InProcessFolderId field if non-nil, zero value otherwise.

### GetInProcessFolderIdOk

`func (o *ThirdPartyFileDto) GetInProcessFolderIdOk() (*int32, bool)`

GetInProcessFolderIdOk returns a tuple with the InProcessFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderId

`func (o *ThirdPartyFileDto) SetInProcessFolderId(v int32)`

SetInProcessFolderId sets InProcessFolderId field to given value.

### HasInProcessFolderId

`func (o *ThirdPartyFileDto) HasInProcessFolderId() bool`

HasInProcessFolderId returns a boolean if a field has been set.

### SetInProcessFolderIdNil

`func (o *ThirdPartyFileDto) SetInProcessFolderIdNil(b bool)`

 SetInProcessFolderIdNil sets the value for InProcessFolderId to be an explicit nil

### UnsetInProcessFolderId
`func (o *ThirdPartyFileDto) UnsetInProcessFolderId()`

UnsetInProcessFolderId ensures that no value is present for InProcessFolderId, not even an explicit nil
### GetInProcessFolderTitle

`func (o *ThirdPartyFileDto) GetInProcessFolderTitle() string`

GetInProcessFolderTitle returns the InProcessFolderTitle field if non-nil, zero value otherwise.

### GetInProcessFolderTitleOk

`func (o *ThirdPartyFileDto) GetInProcessFolderTitleOk() (*string, bool)`

GetInProcessFolderTitleOk returns a tuple with the InProcessFolderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderTitle

`func (o *ThirdPartyFileDto) SetInProcessFolderTitle(v string)`

SetInProcessFolderTitle sets InProcessFolderTitle field to given value.

### HasInProcessFolderTitle

`func (o *ThirdPartyFileDto) HasInProcessFolderTitle() bool`

HasInProcessFolderTitle returns a boolean if a field has been set.

### SetInProcessFolderTitleNil

`func (o *ThirdPartyFileDto) SetInProcessFolderTitleNil(b bool)`

 SetInProcessFolderTitleNil sets the value for InProcessFolderTitle to be an explicit nil

### UnsetInProcessFolderTitle
`func (o *ThirdPartyFileDto) UnsetInProcessFolderTitle()`

UnsetInProcessFolderTitle ensures that no value is present for InProcessFolderTitle, not even an explicit nil
### GetResultsFolderId

`func (o *ThirdPartyFileDto) GetResultsFolderId() int32`

GetResultsFolderId returns the ResultsFolderId field if non-nil, zero value otherwise.

### GetResultsFolderIdOk

`func (o *ThirdPartyFileDto) GetResultsFolderIdOk() (*int32, bool)`

GetResultsFolderIdOk returns a tuple with the ResultsFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFolderId

`func (o *ThirdPartyFileDto) SetResultsFolderId(v int32)`

SetResultsFolderId sets ResultsFolderId field to given value.

### HasResultsFolderId

`func (o *ThirdPartyFileDto) HasResultsFolderId() bool`

HasResultsFolderId returns a boolean if a field has been set.

### SetResultsFolderIdNil

`func (o *ThirdPartyFileDto) SetResultsFolderIdNil(b bool)`

 SetResultsFolderIdNil sets the value for ResultsFolderId to be an explicit nil

### UnsetResultsFolderId
`func (o *ThirdPartyFileDto) UnsetResultsFolderId()`

UnsetResultsFolderId ensures that no value is present for ResultsFolderId, not even an explicit nil
### GetDraftLocation

`func (o *ThirdPartyFileDto) GetDraftLocation() ThirdPartyDraftLocation`

GetDraftLocation returns the DraftLocation field if non-nil, zero value otherwise.

### GetDraftLocationOk

`func (o *ThirdPartyFileDto) GetDraftLocationOk() (*ThirdPartyDraftLocation, bool)`

GetDraftLocationOk returns a tuple with the DraftLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftLocation

`func (o *ThirdPartyFileDto) SetDraftLocation(v ThirdPartyDraftLocation)`

SetDraftLocation sets DraftLocation field to given value.

### HasDraftLocation

`func (o *ThirdPartyFileDto) HasDraftLocation() bool`

HasDraftLocation returns a boolean if a field has been set.

### GetViewAccessibility

`func (o *ThirdPartyFileDto) GetViewAccessibility() FileDtoAllOfViewAccessibility`

GetViewAccessibility returns the ViewAccessibility field if non-nil, zero value otherwise.

### GetViewAccessibilityOk

`func (o *ThirdPartyFileDto) GetViewAccessibilityOk() (*FileDtoAllOfViewAccessibility, bool)`

GetViewAccessibilityOk returns a tuple with the ViewAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAccessibility

`func (o *ThirdPartyFileDto) SetViewAccessibility(v FileDtoAllOfViewAccessibility)`

SetViewAccessibility sets ViewAccessibility field to given value.

### HasViewAccessibility

`func (o *ThirdPartyFileDto) HasViewAccessibility() bool`

HasViewAccessibility returns a boolean if a field has been set.

### SetViewAccessibilityNil

`func (o *ThirdPartyFileDto) SetViewAccessibilityNil(b bool)`

 SetViewAccessibilityNil sets the value for ViewAccessibility to be an explicit nil

### UnsetViewAccessibility
`func (o *ThirdPartyFileDto) UnsetViewAccessibility()`

UnsetViewAccessibility ensures that no value is present for ViewAccessibility, not even an explicit nil
### GetLastOpened

`func (o *ThirdPartyFileDto) GetLastOpened() ApiDateTime`

GetLastOpened returns the LastOpened field if non-nil, zero value otherwise.

### GetLastOpenedOk

`func (o *ThirdPartyFileDto) GetLastOpenedOk() (*ApiDateTime, bool)`

GetLastOpenedOk returns a tuple with the LastOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpened

`func (o *ThirdPartyFileDto) SetLastOpened(v ApiDateTime)`

SetLastOpened sets LastOpened field to given value.

### HasLastOpened

`func (o *ThirdPartyFileDto) HasLastOpened() bool`

HasLastOpened returns a boolean if a field has been set.

### GetExpired

`func (o *ThirdPartyFileDto) GetExpired() ApiDateTime`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ThirdPartyFileDto) GetExpiredOk() (*ApiDateTime, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ThirdPartyFileDto) SetExpired(v ApiDateTime)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ThirdPartyFileDto) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetVectorizationStatus

`func (o *ThirdPartyFileDto) GetVectorizationStatus() VectorizationStatus`

GetVectorizationStatus returns the VectorizationStatus field if non-nil, zero value otherwise.

### GetVectorizationStatusOk

`func (o *ThirdPartyFileDto) GetVectorizationStatusOk() (*VectorizationStatus, bool)`

GetVectorizationStatusOk returns a tuple with the VectorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationStatus

`func (o *ThirdPartyFileDto) SetVectorizationStatus(v VectorizationStatus)`

SetVectorizationStatus sets VectorizationStatus field to given value.

### HasVectorizationStatus

`func (o *ThirdPartyFileDto) HasVectorizationStatus() bool`

HasVectorizationStatus returns a boolean if a field has been set.

### GetExternalDbTableName

`func (o *ThirdPartyFileDto) GetExternalDbTableName() string`

GetExternalDbTableName returns the ExternalDbTableName field if non-nil, zero value otherwise.

### GetExternalDbTableNameOk

`func (o *ThirdPartyFileDto) GetExternalDbTableNameOk() (*string, bool)`

GetExternalDbTableNameOk returns a tuple with the ExternalDbTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDbTableName

`func (o *ThirdPartyFileDto) SetExternalDbTableName(v string)`

SetExternalDbTableName sets ExternalDbTableName field to given value.

### HasExternalDbTableName

`func (o *ThirdPartyFileDto) HasExternalDbTableName() bool`

HasExternalDbTableName returns a boolean if a field has been set.

### SetExternalDbTableNameNil

`func (o *ThirdPartyFileDto) SetExternalDbTableNameNil(b bool)`

 SetExternalDbTableNameNil sets the value for ExternalDbTableName to be an explicit nil

### UnsetExternalDbTableName
`func (o *ThirdPartyFileDto) UnsetExternalDbTableName()`

UnsetExternalDbTableName ensures that no value is present for ExternalDbTableName, not even an explicit nil
### GetDimensions

`func (o *ThirdPartyFileDto) GetDimensions() Size`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ThirdPartyFileDto) GetDimensionsOk() (*Size, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ThirdPartyFileDto) SetDimensions(v Size)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ThirdPartyFileDto) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


