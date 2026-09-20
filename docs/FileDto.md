# FileDto

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
**Id** | Pointer to **int32** | The identifier to pass back to the other operations of this entry. It is a number for storage on the portal  and a string for a connected third-party account, and it is unique only within its own kind, so files and  folders may carry the same value. | [optional] 
**RootFolderId** | Pointer to **int32** | The section the entry ultimately lies in, as an identifier that can be listed like any other folder. For an  entry inside a room this is the rooms section, not the room. | [optional] 
**OriginId** | Pointer to **int32** | The folder the entry was deleted from, which is where restoring it puts it back. It is left out of the answer  unless the entry is in the trash. | [optional] 
**OriginRoomId** | Pointer to **int32** | The room the entry was deleted from, left out of the answer for anything that was not deleted out of a room. | [optional] 
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
**FolderId** | Pointer to **int32** | The folder the file is stored in. When the file was reached through a share and the caller cannot open its  real parent, the identifier of the Shared with me section is reported instead, so this is where the file is  visible rather than where it physically sits. | [optional] 
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
**IsForm** | Pointer to **NullableBool** | Whether the PDF is a fillable form rather than a plain document. When the stored classification does not say,  the portal opens the file to find out, so the answer is reliable for a PDF and null for anything else. | [optional] 
**CustomFilterEnabled** | Pointer to **NullableBool** | True while a spreadsheet is in the mode where each person sorts and filters their own view without changing  what the others see, and null rather than false when it is not. | [optional] 
**CustomFilterEnabledBy** | Pointer to **NullableString** | The display name of the account that turned that mode on, and null when the caller turned it on themselves. | [optional] 
**StartFilling** | Pointer to **NullableBool** | For a form in a room for filling, whether it has been released for filling; until then it is still being  prepared and only the people running the room work with it. Null for a file this does not apply to. | [optional] 
**IsFillingPreparing** | Pointer to **NullableBool** | True during the short window in which a released form is still being written out by the editor. Neither  filling nor editing is accepted while it lasts, so a client should wait and read the file again. | [optional] 
**InProcessFolderId** | Pointer to **NullableInt32** | Left empty by the portal: the folder holding the caller's draft is reported in `draftLocation` instead. | [optional] 
**InProcessFolderTitle** | Pointer to **NullableString** | Left empty by the portal, like the identifier beside it; the draft's folder is named in `draftLocation`. | [optional] 
**ResultsFolderId** | Pointer to **NullableInt32** | The folder that collects the completed copies of this form. It is filled in only for the original form of a  room for filling, and only for a caller allowed to work with that form; null everywhere else. | [optional] 
**DraftLocation** | Pointer to [**DraftLocation**](DraftLocation.md) | Where the caller's own filling draft of this form is kept. Null when there is no draft yet, which is the same  thing `hasDraft` reports. | [optional] 
**ViewAccessibility** | Pointer to [**NullableFileDtoAllOfViewAccessibility**](FileDtoAllOfViewAccessibility.md) |  | [optional] 
**LastOpened** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the caller last opened the file. It is kept per account and is what orders the Recent section, so  it is null for a file this account has never opened. Written with the offset of the portal's time zone. | [optional] 
**Expired** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the file falls under the lifetime rule of the room holding it and is removed. It is counted from  the first revision rather than the latest one, so editing a file does not postpone it, and it is null when the  room sets no lifetime. Written with the offset of the portal's time zone. | [optional] 
**VectorizationStatus** | Pointer to [**VectorizationStatus**](VectorizationStatus.md) | How far the indexing of the file's content for AI search has got. It is null for a file that has never been  queued for indexing, which is every file while the feature is off for the portal. | [optional] 
**ExternalDbTableName** | Pointer to **NullableString** | The table collecting the submitted values of this form in the external database configured for its room. The  field is left out of the answer entirely when the form has no such table. | [optional] 
**Dimensions** | Pointer to [**Size**](Size.md) | The pixel size of the picture, measured by reading the stored file rather than taken from any stored metadata.  Null for anything that is not a picture the portal can show, and also when the file could not be read. | [optional] 

## Methods

### NewFileDto

`func NewFileDto() *FileDto`

NewFileDto instantiates a new FileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDtoWithDefaults

`func NewFileDtoWithDefaults() *FileDto`

NewFileDtoWithDefaults instantiates a new FileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FileDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccess

`func (o *FileDto) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileDto) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileDto) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileDto) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSharedBy

`func (o *FileDto) GetSharedBy() EmployeeDto`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *FileDto) GetSharedByOk() (*EmployeeDto, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *FileDto) SetSharedBy(v EmployeeDto)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *FileDto) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetOwnedBy

`func (o *FileDto) GetOwnedBy() EmployeeDto`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FileDto) GetOwnedByOk() (*EmployeeDto, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FileDto) SetOwnedBy(v EmployeeDto)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *FileDto) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetShared

`func (o *FileDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *FileDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *FileDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *FileDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSharedForUser

`func (o *FileDto) GetSharedForUser() bool`

GetSharedForUser returns the SharedForUser field if non-nil, zero value otherwise.

### GetSharedForUserOk

`func (o *FileDto) GetSharedForUserOk() (*bool, bool)`

GetSharedForUserOk returns a tuple with the SharedForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedForUser

`func (o *FileDto) SetSharedForUser(v bool)`

SetSharedForUser sets SharedForUser field to given value.

### HasSharedForUser

`func (o *FileDto) HasSharedForUser() bool`

HasSharedForUser returns a boolean if a field has been set.

### GetSharedExternal

`func (o *FileDto) GetSharedExternal() bool`

GetSharedExternal returns the SharedExternal field if non-nil, zero value otherwise.

### GetSharedExternalOk

`func (o *FileDto) GetSharedExternalOk() (*bool, bool)`

GetSharedExternalOk returns a tuple with the SharedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedExternal

`func (o *FileDto) SetSharedExternal(v bool)`

SetSharedExternal sets SharedExternal field to given value.

### HasSharedExternal

`func (o *FileDto) HasSharedExternal() bool`

HasSharedExternal returns a boolean if a field has been set.

### GetParentShared

`func (o *FileDto) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *FileDto) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *FileDto) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *FileDto) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetShortWebUrl

`func (o *FileDto) GetShortWebUrl() string`

GetShortWebUrl returns the ShortWebUrl field if non-nil, zero value otherwise.

### GetShortWebUrlOk

`func (o *FileDto) GetShortWebUrlOk() (*string, bool)`

GetShortWebUrlOk returns a tuple with the ShortWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortWebUrl

`func (o *FileDto) SetShortWebUrl(v string)`

SetShortWebUrl sets ShortWebUrl field to given value.

### HasShortWebUrl

`func (o *FileDto) HasShortWebUrl() bool`

HasShortWebUrl returns a boolean if a field has been set.

### GetCreated

`func (o *FileDto) GetCreated() ApiDateTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileDto) GetCreatedOk() (*ApiDateTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileDto) SetCreated(v ApiDateTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *FileDto) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FileDto) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FileDto) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FileDto) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FileDto) GetAutoDelete() ApiDateTime`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FileDto) GetAutoDeleteOk() (*ApiDateTime, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FileDto) SetAutoDelete(v ApiDateTime)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FileDto) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetRootFolderType

`func (o *FileDto) GetRootFolderType() FolderType`

GetRootFolderType returns the RootFolderType field if non-nil, zero value otherwise.

### GetRootFolderTypeOk

`func (o *FileDto) GetRootFolderTypeOk() (*FolderType, bool)`

GetRootFolderTypeOk returns a tuple with the RootFolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderType

`func (o *FileDto) SetRootFolderType(v FolderType)`

SetRootFolderType sets RootFolderType field to given value.

### HasRootFolderType

`func (o *FileDto) HasRootFolderType() bool`

HasRootFolderType returns a boolean if a field has been set.

### GetParentRoomType

`func (o *FileDto) GetParentRoomType() FolderType`

GetParentRoomType returns the ParentRoomType field if non-nil, zero value otherwise.

### GetParentRoomTypeOk

`func (o *FileDto) GetParentRoomTypeOk() (*FolderType, bool)`

GetParentRoomTypeOk returns a tuple with the ParentRoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoomType

`func (o *FileDto) SetParentRoomType(v FolderType)`

SetParentRoomType sets ParentRoomType field to given value.

### HasParentRoomType

`func (o *FileDto) HasParentRoomType() bool`

HasParentRoomType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileDto) GetUpdatedBy() EmployeeDto`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileDto) GetUpdatedByOk() (*EmployeeDto, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileDto) SetUpdatedBy(v EmployeeDto)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProviderItem

`func (o *FileDto) GetProviderItem() bool`

GetProviderItem returns the ProviderItem field if non-nil, zero value otherwise.

### GetProviderItemOk

`func (o *FileDto) GetProviderItemOk() (*bool, bool)`

GetProviderItemOk returns a tuple with the ProviderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderItem

`func (o *FileDto) SetProviderItem(v bool)`

SetProviderItem sets ProviderItem field to given value.

### HasProviderItem

`func (o *FileDto) HasProviderItem() bool`

HasProviderItem returns a boolean if a field has been set.

### GetProviderKey

`func (o *FileDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *FileDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *FileDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *FileDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderId

`func (o *FileDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *FileDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *FileDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *FileDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrder

`func (o *FileDto) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FileDto) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FileDto) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FileDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsFavorite

`func (o *FileDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *FileDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *FileDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *FileDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetFileEntryType

`func (o *FileDto) GetFileEntryType() FileEntryType`

GetFileEntryType returns the FileEntryType field if non-nil, zero value otherwise.

### GetFileEntryTypeOk

`func (o *FileDto) GetFileEntryTypeOk() (*FileEntryType, bool)`

GetFileEntryTypeOk returns a tuple with the FileEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEntryType

`func (o *FileDto) SetFileEntryType(v FileEntryType)`

SetFileEntryType sets FileEntryType field to given value.

### HasFileEntryType

`func (o *FileDto) HasFileEntryType() bool`

HasFileEntryType returns a boolean if a field has been set.

### GetId

`func (o *FileDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootFolderId

`func (o *FileDto) GetRootFolderId() int32`

GetRootFolderId returns the RootFolderId field if non-nil, zero value otherwise.

### GetRootFolderIdOk

`func (o *FileDto) GetRootFolderIdOk() (*int32, bool)`

GetRootFolderIdOk returns a tuple with the RootFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderId

`func (o *FileDto) SetRootFolderId(v int32)`

SetRootFolderId sets RootFolderId field to given value.

### HasRootFolderId

`func (o *FileDto) HasRootFolderId() bool`

HasRootFolderId returns a boolean if a field has been set.

### GetOriginId

`func (o *FileDto) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *FileDto) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *FileDto) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *FileDto) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginRoomId

`func (o *FileDto) GetOriginRoomId() int32`

GetOriginRoomId returns the OriginRoomId field if non-nil, zero value otherwise.

### GetOriginRoomIdOk

`func (o *FileDto) GetOriginRoomIdOk() (*int32, bool)`

GetOriginRoomIdOk returns a tuple with the OriginRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomId

`func (o *FileDto) SetOriginRoomId(v int32)`

SetOriginRoomId sets OriginRoomId field to given value.

### HasOriginRoomId

`func (o *FileDto) HasOriginRoomId() bool`

HasOriginRoomId returns a boolean if a field has been set.

### GetOriginTitle

`func (o *FileDto) GetOriginTitle() string`

GetOriginTitle returns the OriginTitle field if non-nil, zero value otherwise.

### GetOriginTitleOk

`func (o *FileDto) GetOriginTitleOk() (*string, bool)`

GetOriginTitleOk returns a tuple with the OriginTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTitle

`func (o *FileDto) SetOriginTitle(v string)`

SetOriginTitle sets OriginTitle field to given value.

### HasOriginTitle

`func (o *FileDto) HasOriginTitle() bool`

HasOriginTitle returns a boolean if a field has been set.

### GetOriginRoomTitle

`func (o *FileDto) GetOriginRoomTitle() string`

GetOriginRoomTitle returns the OriginRoomTitle field if non-nil, zero value otherwise.

### GetOriginRoomTitleOk

`func (o *FileDto) GetOriginRoomTitleOk() (*string, bool)`

GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRoomTitle

`func (o *FileDto) SetOriginRoomTitle(v string)`

SetOriginRoomTitle sets OriginRoomTitle field to given value.

### HasOriginRoomTitle

`func (o *FileDto) HasOriginRoomTitle() bool`

HasOriginRoomTitle returns a boolean if a field has been set.

### GetCanShare

`func (o *FileDto) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *FileDto) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *FileDto) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *FileDto) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### GetShareSettings

`func (o *FileDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings`

GetShareSettings returns the ShareSettings field if non-nil, zero value otherwise.

### GetShareSettingsOk

`func (o *FileDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool)`

GetShareSettingsOk returns a tuple with the ShareSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSettings

`func (o *FileDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings)`

SetShareSettings sets ShareSettings field to given value.

### HasShareSettings

`func (o *FileDto) HasShareSettings() bool`

HasShareSettings returns a boolean if a field has been set.

### SetShareSettingsNil

`func (o *FileDto) SetShareSettingsNil(b bool)`

 SetShareSettingsNil sets the value for ShareSettings to be an explicit nil

### UnsetShareSettings
`func (o *FileDto) UnsetShareSettings()`

UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
### GetSecurity

`func (o *FileDto) GetSecurity() AiFileEntryDtoAllOfSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FileDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FileDto) SetSecurity(v AiFileEntryDtoAllOfSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *FileDto) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *FileDto) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *FileDto) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetAvailableShareRights

`func (o *FileDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights`

GetAvailableShareRights returns the AvailableShareRights field if non-nil, zero value otherwise.

### GetAvailableShareRightsOk

`func (o *FileDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool)`

GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShareRights

`func (o *FileDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights)`

SetAvailableShareRights sets AvailableShareRights field to given value.

### HasAvailableShareRights

`func (o *FileDto) HasAvailableShareRights() bool`

HasAvailableShareRights returns a boolean if a field has been set.

### SetAvailableShareRightsNil

`func (o *FileDto) SetAvailableShareRightsNil(b bool)`

 SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil

### UnsetAvailableShareRights
`func (o *FileDto) UnsetAvailableShareRights()`

UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
### GetRequestToken

`func (o *FileDto) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FileDto) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FileDto) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FileDto) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetExternal

`func (o *FileDto) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *FileDto) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *FileDto) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *FileDto) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExpirationDate

`func (o *FileDto) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileDto) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileDto) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsLinkExpired

`func (o *FileDto) GetIsLinkExpired() bool`

GetIsLinkExpired returns the IsLinkExpired field if non-nil, zero value otherwise.

### GetIsLinkExpiredOk

`func (o *FileDto) GetIsLinkExpiredOk() (*bool, bool)`

GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkExpired

`func (o *FileDto) SetIsLinkExpired(v bool)`

SetIsLinkExpired sets IsLinkExpired field to given value.

### HasIsLinkExpired

`func (o *FileDto) HasIsLinkExpired() bool`

HasIsLinkExpired returns a boolean if a field has been set.

### GetFolderId

`func (o *FileDto) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *FileDto) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *FileDto) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *FileDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVersion

`func (o *FileDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FileDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *FileDto) GetVersionGroup() int32`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *FileDto) GetVersionGroupOk() (*int32, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *FileDto) SetVersionGroup(v int32)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *FileDto) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetContentLength

`func (o *FileDto) GetContentLength() string`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *FileDto) GetContentLengthOk() (*string, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *FileDto) SetContentLength(v string)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *FileDto) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *FileDto) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *FileDto) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetPureContentLength

`func (o *FileDto) GetPureContentLength() int64`

GetPureContentLength returns the PureContentLength field if non-nil, zero value otherwise.

### GetPureContentLengthOk

`func (o *FileDto) GetPureContentLengthOk() (*int64, bool)`

GetPureContentLengthOk returns a tuple with the PureContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureContentLength

`func (o *FileDto) SetPureContentLength(v int64)`

SetPureContentLength sets PureContentLength field to given value.

### HasPureContentLength

`func (o *FileDto) HasPureContentLength() bool`

HasPureContentLength returns a boolean if a field has been set.

### SetPureContentLengthNil

`func (o *FileDto) SetPureContentLengthNil(b bool)`

 SetPureContentLengthNil sets the value for PureContentLength to be an explicit nil

### UnsetPureContentLength
`func (o *FileDto) UnsetPureContentLength()`

UnsetPureContentLength ensures that no value is present for PureContentLength, not even an explicit nil
### GetFileStatus

`func (o *FileDto) GetFileStatus() FileStatus`

GetFileStatus returns the FileStatus field if non-nil, zero value otherwise.

### GetFileStatusOk

`func (o *FileDto) GetFileStatusOk() (*FileStatus, bool)`

GetFileStatusOk returns a tuple with the FileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStatus

`func (o *FileDto) SetFileStatus(v FileStatus)`

SetFileStatus sets FileStatus field to given value.

### HasFileStatus

`func (o *FileDto) HasFileStatus() bool`

HasFileStatus returns a boolean if a field has been set.

### GetEditingBy

`func (o *FileDto) GetEditingBy() map[string]*string`

GetEditingBy returns the EditingBy field if non-nil, zero value otherwise.

### GetEditingByOk

`func (o *FileDto) GetEditingByOk() (*map[string]*string, bool)`

GetEditingByOk returns a tuple with the EditingBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditingBy

`func (o *FileDto) SetEditingBy(v map[string]*string)`

SetEditingBy sets EditingBy field to given value.

### HasEditingBy

`func (o *FileDto) HasEditingBy() bool`

HasEditingBy returns a boolean if a field has been set.

### GetMute

`func (o *FileDto) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *FileDto) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *FileDto) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *FileDto) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetViewUrl

`func (o *FileDto) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *FileDto) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *FileDto) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *FileDto) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### SetViewUrlNil

`func (o *FileDto) SetViewUrlNil(b bool)`

 SetViewUrlNil sets the value for ViewUrl to be an explicit nil

### UnsetViewUrl
`func (o *FileDto) UnsetViewUrl()`

UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
### GetWebUrl

`func (o *FileDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *FileDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *FileDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *FileDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *FileDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *FileDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFileType

`func (o *FileDto) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileDto) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileDto) SetFileType(v FileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FileDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileExst

`func (o *FileDto) GetFileExst() string`

GetFileExst returns the FileExst field if non-nil, zero value otherwise.

### GetFileExstOk

`func (o *FileDto) GetFileExstOk() (*string, bool)`

GetFileExstOk returns a tuple with the FileExst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExst

`func (o *FileDto) SetFileExst(v string)`

SetFileExst sets FileExst field to given value.

### HasFileExst

`func (o *FileDto) HasFileExst() bool`

HasFileExst returns a boolean if a field has been set.

### SetFileExstNil

`func (o *FileDto) SetFileExstNil(b bool)`

 SetFileExstNil sets the value for FileExst to be an explicit nil

### UnsetFileExst
`func (o *FileDto) UnsetFileExst()`

UnsetFileExst ensures that no value is present for FileExst, not even an explicit nil
### GetComment

`func (o *FileDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FileDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FileDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FileDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *FileDto) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FileDto) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEncrypted

`func (o *FileDto) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FileDto) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FileDto) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *FileDto) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### SetEncryptedNil

`func (o *FileDto) SetEncryptedNil(b bool)`

 SetEncryptedNil sets the value for Encrypted to be an explicit nil

### UnsetEncrypted
`func (o *FileDto) UnsetEncrypted()`

UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
### GetThumbnailUrl

`func (o *FileDto) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *FileDto) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *FileDto) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *FileDto) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *FileDto) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *FileDto) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailStatus

`func (o *FileDto) GetThumbnailStatus() Thumbnail`

GetThumbnailStatus returns the ThumbnailStatus field if non-nil, zero value otherwise.

### GetThumbnailStatusOk

`func (o *FileDto) GetThumbnailStatusOk() (*Thumbnail, bool)`

GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailStatus

`func (o *FileDto) SetThumbnailStatus(v Thumbnail)`

SetThumbnailStatus sets ThumbnailStatus field to given value.

### HasThumbnailStatus

`func (o *FileDto) HasThumbnailStatus() bool`

HasThumbnailStatus returns a boolean if a field has been set.

### GetLocked

`func (o *FileDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FileDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FileDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FileDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *FileDto) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *FileDto) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetLockedBy

`func (o *FileDto) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *FileDto) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *FileDto) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *FileDto) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### SetLockedByNil

`func (o *FileDto) SetLockedByNil(b bool)`

 SetLockedByNil sets the value for LockedBy to be an explicit nil

### UnsetLockedBy
`func (o *FileDto) UnsetLockedBy()`

UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
### GetHasDraft

`func (o *FileDto) GetHasDraft() bool`

GetHasDraft returns the HasDraft field if non-nil, zero value otherwise.

### GetHasDraftOk

`func (o *FileDto) GetHasDraftOk() (*bool, bool)`

GetHasDraftOk returns a tuple with the HasDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDraft

`func (o *FileDto) SetHasDraft(v bool)`

SetHasDraft sets HasDraft field to given value.

### HasHasDraft

`func (o *FileDto) HasHasDraft() bool`

HasHasDraft returns a boolean if a field has been set.

### SetHasDraftNil

`func (o *FileDto) SetHasDraftNil(b bool)`

 SetHasDraftNil sets the value for HasDraft to be an explicit nil

### UnsetHasDraft
`func (o *FileDto) UnsetHasDraft()`

UnsetHasDraft ensures that no value is present for HasDraft, not even an explicit nil
### GetFormFillingStatus

`func (o *FileDto) GetFormFillingStatus() FormFillingStatus`

GetFormFillingStatus returns the FormFillingStatus field if non-nil, zero value otherwise.

### GetFormFillingStatusOk

`func (o *FileDto) GetFormFillingStatusOk() (*FormFillingStatus, bool)`

GetFormFillingStatusOk returns a tuple with the FormFillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFillingStatus

`func (o *FileDto) SetFormFillingStatus(v FormFillingStatus)`

SetFormFillingStatus sets FormFillingStatus field to given value.

### HasFormFillingStatus

`func (o *FileDto) HasFormFillingStatus() bool`

HasFormFillingStatus returns a boolean if a field has been set.

### GetIsForm

`func (o *FileDto) GetIsForm() bool`

GetIsForm returns the IsForm field if non-nil, zero value otherwise.

### GetIsFormOk

`func (o *FileDto) GetIsFormOk() (*bool, bool)`

GetIsFormOk returns a tuple with the IsForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForm

`func (o *FileDto) SetIsForm(v bool)`

SetIsForm sets IsForm field to given value.

### HasIsForm

`func (o *FileDto) HasIsForm() bool`

HasIsForm returns a boolean if a field has been set.

### SetIsFormNil

`func (o *FileDto) SetIsFormNil(b bool)`

 SetIsFormNil sets the value for IsForm to be an explicit nil

### UnsetIsForm
`func (o *FileDto) UnsetIsForm()`

UnsetIsForm ensures that no value is present for IsForm, not even an explicit nil
### GetCustomFilterEnabled

`func (o *FileDto) GetCustomFilterEnabled() bool`

GetCustomFilterEnabled returns the CustomFilterEnabled field if non-nil, zero value otherwise.

### GetCustomFilterEnabledOk

`func (o *FileDto) GetCustomFilterEnabledOk() (*bool, bool)`

GetCustomFilterEnabledOk returns a tuple with the CustomFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabled

`func (o *FileDto) SetCustomFilterEnabled(v bool)`

SetCustomFilterEnabled sets CustomFilterEnabled field to given value.

### HasCustomFilterEnabled

`func (o *FileDto) HasCustomFilterEnabled() bool`

HasCustomFilterEnabled returns a boolean if a field has been set.

### SetCustomFilterEnabledNil

`func (o *FileDto) SetCustomFilterEnabledNil(b bool)`

 SetCustomFilterEnabledNil sets the value for CustomFilterEnabled to be an explicit nil

### UnsetCustomFilterEnabled
`func (o *FileDto) UnsetCustomFilterEnabled()`

UnsetCustomFilterEnabled ensures that no value is present for CustomFilterEnabled, not even an explicit nil
### GetCustomFilterEnabledBy

`func (o *FileDto) GetCustomFilterEnabledBy() string`

GetCustomFilterEnabledBy returns the CustomFilterEnabledBy field if non-nil, zero value otherwise.

### GetCustomFilterEnabledByOk

`func (o *FileDto) GetCustomFilterEnabledByOk() (*string, bool)`

GetCustomFilterEnabledByOk returns a tuple with the CustomFilterEnabledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterEnabledBy

`func (o *FileDto) SetCustomFilterEnabledBy(v string)`

SetCustomFilterEnabledBy sets CustomFilterEnabledBy field to given value.

### HasCustomFilterEnabledBy

`func (o *FileDto) HasCustomFilterEnabledBy() bool`

HasCustomFilterEnabledBy returns a boolean if a field has been set.

### SetCustomFilterEnabledByNil

`func (o *FileDto) SetCustomFilterEnabledByNil(b bool)`

 SetCustomFilterEnabledByNil sets the value for CustomFilterEnabledBy to be an explicit nil

### UnsetCustomFilterEnabledBy
`func (o *FileDto) UnsetCustomFilterEnabledBy()`

UnsetCustomFilterEnabledBy ensures that no value is present for CustomFilterEnabledBy, not even an explicit nil
### GetStartFilling

`func (o *FileDto) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *FileDto) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *FileDto) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *FileDto) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *FileDto) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *FileDto) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetIsFillingPreparing

`func (o *FileDto) GetIsFillingPreparing() bool`

GetIsFillingPreparing returns the IsFillingPreparing field if non-nil, zero value otherwise.

### GetIsFillingPreparingOk

`func (o *FileDto) GetIsFillingPreparingOk() (*bool, bool)`

GetIsFillingPreparingOk returns a tuple with the IsFillingPreparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFillingPreparing

`func (o *FileDto) SetIsFillingPreparing(v bool)`

SetIsFillingPreparing sets IsFillingPreparing field to given value.

### HasIsFillingPreparing

`func (o *FileDto) HasIsFillingPreparing() bool`

HasIsFillingPreparing returns a boolean if a field has been set.

### SetIsFillingPreparingNil

`func (o *FileDto) SetIsFillingPreparingNil(b bool)`

 SetIsFillingPreparingNil sets the value for IsFillingPreparing to be an explicit nil

### UnsetIsFillingPreparing
`func (o *FileDto) UnsetIsFillingPreparing()`

UnsetIsFillingPreparing ensures that no value is present for IsFillingPreparing, not even an explicit nil
### GetInProcessFolderId

`func (o *FileDto) GetInProcessFolderId() int32`

GetInProcessFolderId returns the InProcessFolderId field if non-nil, zero value otherwise.

### GetInProcessFolderIdOk

`func (o *FileDto) GetInProcessFolderIdOk() (*int32, bool)`

GetInProcessFolderIdOk returns a tuple with the InProcessFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderId

`func (o *FileDto) SetInProcessFolderId(v int32)`

SetInProcessFolderId sets InProcessFolderId field to given value.

### HasInProcessFolderId

`func (o *FileDto) HasInProcessFolderId() bool`

HasInProcessFolderId returns a boolean if a field has been set.

### SetInProcessFolderIdNil

`func (o *FileDto) SetInProcessFolderIdNil(b bool)`

 SetInProcessFolderIdNil sets the value for InProcessFolderId to be an explicit nil

### UnsetInProcessFolderId
`func (o *FileDto) UnsetInProcessFolderId()`

UnsetInProcessFolderId ensures that no value is present for InProcessFolderId, not even an explicit nil
### GetInProcessFolderTitle

`func (o *FileDto) GetInProcessFolderTitle() string`

GetInProcessFolderTitle returns the InProcessFolderTitle field if non-nil, zero value otherwise.

### GetInProcessFolderTitleOk

`func (o *FileDto) GetInProcessFolderTitleOk() (*string, bool)`

GetInProcessFolderTitleOk returns a tuple with the InProcessFolderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProcessFolderTitle

`func (o *FileDto) SetInProcessFolderTitle(v string)`

SetInProcessFolderTitle sets InProcessFolderTitle field to given value.

### HasInProcessFolderTitle

`func (o *FileDto) HasInProcessFolderTitle() bool`

HasInProcessFolderTitle returns a boolean if a field has been set.

### SetInProcessFolderTitleNil

`func (o *FileDto) SetInProcessFolderTitleNil(b bool)`

 SetInProcessFolderTitleNil sets the value for InProcessFolderTitle to be an explicit nil

### UnsetInProcessFolderTitle
`func (o *FileDto) UnsetInProcessFolderTitle()`

UnsetInProcessFolderTitle ensures that no value is present for InProcessFolderTitle, not even an explicit nil
### GetResultsFolderId

`func (o *FileDto) GetResultsFolderId() int32`

GetResultsFolderId returns the ResultsFolderId field if non-nil, zero value otherwise.

### GetResultsFolderIdOk

`func (o *FileDto) GetResultsFolderIdOk() (*int32, bool)`

GetResultsFolderIdOk returns a tuple with the ResultsFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFolderId

`func (o *FileDto) SetResultsFolderId(v int32)`

SetResultsFolderId sets ResultsFolderId field to given value.

### HasResultsFolderId

`func (o *FileDto) HasResultsFolderId() bool`

HasResultsFolderId returns a boolean if a field has been set.

### SetResultsFolderIdNil

`func (o *FileDto) SetResultsFolderIdNil(b bool)`

 SetResultsFolderIdNil sets the value for ResultsFolderId to be an explicit nil

### UnsetResultsFolderId
`func (o *FileDto) UnsetResultsFolderId()`

UnsetResultsFolderId ensures that no value is present for ResultsFolderId, not even an explicit nil
### GetDraftLocation

`func (o *FileDto) GetDraftLocation() DraftLocation`

GetDraftLocation returns the DraftLocation field if non-nil, zero value otherwise.

### GetDraftLocationOk

`func (o *FileDto) GetDraftLocationOk() (*DraftLocation, bool)`

GetDraftLocationOk returns a tuple with the DraftLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftLocation

`func (o *FileDto) SetDraftLocation(v DraftLocation)`

SetDraftLocation sets DraftLocation field to given value.

### HasDraftLocation

`func (o *FileDto) HasDraftLocation() bool`

HasDraftLocation returns a boolean if a field has been set.

### GetViewAccessibility

`func (o *FileDto) GetViewAccessibility() FileDtoAllOfViewAccessibility`

GetViewAccessibility returns the ViewAccessibility field if non-nil, zero value otherwise.

### GetViewAccessibilityOk

`func (o *FileDto) GetViewAccessibilityOk() (*FileDtoAllOfViewAccessibility, bool)`

GetViewAccessibilityOk returns a tuple with the ViewAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAccessibility

`func (o *FileDto) SetViewAccessibility(v FileDtoAllOfViewAccessibility)`

SetViewAccessibility sets ViewAccessibility field to given value.

### HasViewAccessibility

`func (o *FileDto) HasViewAccessibility() bool`

HasViewAccessibility returns a boolean if a field has been set.

### SetViewAccessibilityNil

`func (o *FileDto) SetViewAccessibilityNil(b bool)`

 SetViewAccessibilityNil sets the value for ViewAccessibility to be an explicit nil

### UnsetViewAccessibility
`func (o *FileDto) UnsetViewAccessibility()`

UnsetViewAccessibility ensures that no value is present for ViewAccessibility, not even an explicit nil
### GetLastOpened

`func (o *FileDto) GetLastOpened() ApiDateTime`

GetLastOpened returns the LastOpened field if non-nil, zero value otherwise.

### GetLastOpenedOk

`func (o *FileDto) GetLastOpenedOk() (*ApiDateTime, bool)`

GetLastOpenedOk returns a tuple with the LastOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpened

`func (o *FileDto) SetLastOpened(v ApiDateTime)`

SetLastOpened sets LastOpened field to given value.

### HasLastOpened

`func (o *FileDto) HasLastOpened() bool`

HasLastOpened returns a boolean if a field has been set.

### GetExpired

`func (o *FileDto) GetExpired() ApiDateTime`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *FileDto) GetExpiredOk() (*ApiDateTime, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *FileDto) SetExpired(v ApiDateTime)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *FileDto) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetVectorizationStatus

`func (o *FileDto) GetVectorizationStatus() VectorizationStatus`

GetVectorizationStatus returns the VectorizationStatus field if non-nil, zero value otherwise.

### GetVectorizationStatusOk

`func (o *FileDto) GetVectorizationStatusOk() (*VectorizationStatus, bool)`

GetVectorizationStatusOk returns a tuple with the VectorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationStatus

`func (o *FileDto) SetVectorizationStatus(v VectorizationStatus)`

SetVectorizationStatus sets VectorizationStatus field to given value.

### HasVectorizationStatus

`func (o *FileDto) HasVectorizationStatus() bool`

HasVectorizationStatus returns a boolean if a field has been set.

### GetExternalDbTableName

`func (o *FileDto) GetExternalDbTableName() string`

GetExternalDbTableName returns the ExternalDbTableName field if non-nil, zero value otherwise.

### GetExternalDbTableNameOk

`func (o *FileDto) GetExternalDbTableNameOk() (*string, bool)`

GetExternalDbTableNameOk returns a tuple with the ExternalDbTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDbTableName

`func (o *FileDto) SetExternalDbTableName(v string)`

SetExternalDbTableName sets ExternalDbTableName field to given value.

### HasExternalDbTableName

`func (o *FileDto) HasExternalDbTableName() bool`

HasExternalDbTableName returns a boolean if a field has been set.

### SetExternalDbTableNameNil

`func (o *FileDto) SetExternalDbTableNameNil(b bool)`

 SetExternalDbTableNameNil sets the value for ExternalDbTableName to be an explicit nil

### UnsetExternalDbTableName
`func (o *FileDto) UnsetExternalDbTableName()`

UnsetExternalDbTableName ensures that no value is present for ExternalDbTableName, not even an explicit nil
### GetDimensions

`func (o *FileDto) GetDimensions() Size`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *FileDto) GetDimensionsOk() (*Size, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *FileDto) SetDimensions(v Size)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *FileDto) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


