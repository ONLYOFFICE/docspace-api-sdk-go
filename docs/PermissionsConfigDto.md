# PermissionsConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **bool** | Defines if the document can be commented or not. | [optional] 
**Chat** | Pointer to **bool** | Defines if the chat functionality is enabled in the document or not. | [optional] 
**Download** | Pointer to **bool** | Defines if the document can be downloaded or only viewed or edited online. | [optional] 
**Edit** | Pointer to **bool** | Defines if the document can be edited or only viewed. | [optional] 
**FillForms** | Pointer to **bool** | Defines if the forms can be filled. | [optional] 
**ModifyFilter** | Pointer to **bool** | Defines if the filter can be applied globally (true) affecting all the other users,  or locally (false), i.e. for the current user only. | [optional] 
**Protect** | Pointer to **bool** | Defines if the Protection tab on the toolbar and the Protect button in the left menu are displayedor hidden. | [optional] 
**Print** | Pointer to **bool** | Defines if the document can be printed or not. | [optional] 
**Review** | Pointer to **bool** | Defines if the document can be reviewed or not. | [optional] 
**Copy** | Pointer to **bool** | Defines if the content can be copied to the clipboard or not. | [optional] 

## Methods

### NewPermissionsConfigDto

`func NewPermissionsConfigDto() *PermissionsConfigDto`

NewPermissionsConfigDto instantiates a new PermissionsConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsConfigDtoWithDefaults

`func NewPermissionsConfigDtoWithDefaults() *PermissionsConfigDto`

NewPermissionsConfigDtoWithDefaults instantiates a new PermissionsConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PermissionsConfigDto) GetComment() bool`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PermissionsConfigDto) GetCommentOk() (*bool, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PermissionsConfigDto) SetComment(v bool)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PermissionsConfigDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetChat

`func (o *PermissionsConfigDto) GetChat() bool`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *PermissionsConfigDto) GetChatOk() (*bool, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *PermissionsConfigDto) SetChat(v bool)`

SetChat sets Chat field to given value.

### HasChat

`func (o *PermissionsConfigDto) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetDownload

`func (o *PermissionsConfigDto) GetDownload() bool`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *PermissionsConfigDto) GetDownloadOk() (*bool, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *PermissionsConfigDto) SetDownload(v bool)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *PermissionsConfigDto) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetEdit

`func (o *PermissionsConfigDto) GetEdit() bool`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *PermissionsConfigDto) GetEditOk() (*bool, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *PermissionsConfigDto) SetEdit(v bool)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *PermissionsConfigDto) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetFillForms

`func (o *PermissionsConfigDto) GetFillForms() bool`

GetFillForms returns the FillForms field if non-nil, zero value otherwise.

### GetFillFormsOk

`func (o *PermissionsConfigDto) GetFillFormsOk() (*bool, bool)`

GetFillFormsOk returns a tuple with the FillForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillForms

`func (o *PermissionsConfigDto) SetFillForms(v bool)`

SetFillForms sets FillForms field to given value.

### HasFillForms

`func (o *PermissionsConfigDto) HasFillForms() bool`

HasFillForms returns a boolean if a field has been set.

### GetModifyFilter

`func (o *PermissionsConfigDto) GetModifyFilter() bool`

GetModifyFilter returns the ModifyFilter field if non-nil, zero value otherwise.

### GetModifyFilterOk

`func (o *PermissionsConfigDto) GetModifyFilterOk() (*bool, bool)`

GetModifyFilterOk returns a tuple with the ModifyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyFilter

`func (o *PermissionsConfigDto) SetModifyFilter(v bool)`

SetModifyFilter sets ModifyFilter field to given value.

### HasModifyFilter

`func (o *PermissionsConfigDto) HasModifyFilter() bool`

HasModifyFilter returns a boolean if a field has been set.

### GetProtect

`func (o *PermissionsConfigDto) GetProtect() bool`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *PermissionsConfigDto) GetProtectOk() (*bool, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *PermissionsConfigDto) SetProtect(v bool)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *PermissionsConfigDto) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetPrint

`func (o *PermissionsConfigDto) GetPrint() bool`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *PermissionsConfigDto) GetPrintOk() (*bool, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *PermissionsConfigDto) SetPrint(v bool)`

SetPrint sets Print field to given value.

### HasPrint

`func (o *PermissionsConfigDto) HasPrint() bool`

HasPrint returns a boolean if a field has been set.

### GetReview

`func (o *PermissionsConfigDto) GetReview() bool`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *PermissionsConfigDto) GetReviewOk() (*bool, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *PermissionsConfigDto) SetReview(v bool)`

SetReview sets Review field to given value.

### HasReview

`func (o *PermissionsConfigDto) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetCopy

`func (o *PermissionsConfigDto) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *PermissionsConfigDto) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *PermissionsConfigDto) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *PermissionsConfigDto) HasCopy() bool`

HasCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


