// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
)

// checks if the AiFileEntryDtoAllOfSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileEntryDtoAllOfSecurity{}

// AiFileEntryDtoAllOfSecurity What the calling account may do with this entry, one flag per action, and the cheapest way to decide which  operations to offer without trying them. The flags already take the room's settings and the account's role  into account.
type AiFileEntryDtoAllOfSecurity struct {
	Read *bool `json:"Read,omitempty"`
	Comment *bool `json:"Comment,omitempty"`
	FillForms *bool `json:"FillForms,omitempty"`
	Review *bool `json:"Review,omitempty"`
	Create *bool `json:"Create,omitempty"`
	CreateFrom *bool `json:"CreateFrom,omitempty"`
	Edit *bool `json:"Edit,omitempty"`
	Delete *bool `json:"Delete,omitempty"`
	CustomFilter *bool `json:"CustomFilter,omitempty"`
	EditRoom *bool `json:"EditRoom,omitempty"`
	Rename *bool `json:"Rename,omitempty"`
	ReadHistory *bool `json:"ReadHistory,omitempty"`
	Lock *bool `json:"Lock,omitempty"`
	EditHistory *bool `json:"EditHistory,omitempty"`
	CopyTo *bool `json:"CopyTo,omitempty"`
	Copy *bool `json:"Copy,omitempty"`
	MoveTo *bool `json:"MoveTo,omitempty"`
	Move *bool `json:"Move,omitempty"`
	Pin *bool `json:"Pin,omitempty"`
	Mute *bool `json:"Mute,omitempty"`
	EditAccess *bool `json:"EditAccess,omitempty"`
	Duplicate *bool `json:"Duplicate,omitempty"`
	SubmitToFormGallery *bool `json:"SubmitToFormGallery,omitempty"`
	Download *bool `json:"Download,omitempty"`
	Convert *bool `json:"Convert,omitempty"`
	CopySharedLink *bool `json:"CopySharedLink,omitempty"`
	ReadLinks *bool `json:"ReadLinks,omitempty"`
	Reconnect *bool `json:"Reconnect,omitempty"`
	CreateRoomFrom *bool `json:"CreateRoomFrom,omitempty"`
	CopyLink *bool `json:"CopyLink,omitempty"`
	Embed *bool `json:"Embed,omitempty"`
	ChangeOwner *bool `json:"ChangeOwner,omitempty"`
	IndexExport *bool `json:"IndexExport,omitempty"`
	StartFilling *bool `json:"StartFilling,omitempty"`
	FillingStatus *bool `json:"FillingStatus,omitempty"`
	ResetFilling *bool `json:"ResetFilling,omitempty"`
	StopFilling *bool `json:"StopFilling,omitempty"`
	OpenForm *bool `json:"OpenForm,omitempty"`
	EditInternal *bool `json:"EditInternal,omitempty"`
	EditExpiration *bool `json:"EditExpiration,omitempty"`
	Vectorization *bool `json:"Vectorization,omitempty"`
	AskAi *bool `json:"AskAi,omitempty"`
	UseChat *bool `json:"UseChat,omitempty"`
	UpdateXlsx *bool `json:"UpdateXlsx,omitempty"`
	AnalyzeResponses *bool `json:"AnalyzeResponses,omitempty"`
	CanUseAi *bool `json:"CanUseAi,omitempty"`
	HistoryExport *bool `json:"HistoryExport,omitempty"`
}

// NewAiFileEntryDtoAllOfSecurity instantiates a new AiFileEntryDtoAllOfSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileEntryDtoAllOfSecurity() *AiFileEntryDtoAllOfSecurity {
	this := AiFileEntryDtoAllOfSecurity{}
	return &this
}

// NewAiFileEntryDtoAllOfSecurityWithDefaults instantiates a new AiFileEntryDtoAllOfSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileEntryDtoAllOfSecurityWithDefaults() *AiFileEntryDtoAllOfSecurity {
	this := AiFileEntryDtoAllOfSecurity{}
	return &this
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsReadSet() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *AiFileEntryDtoAllOfSecurity) SetRead(v bool) {
	o.Read = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetComment() bool {
	if o == nil || IsNil(o.Comment) {
		var ret bool
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCommentSet() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given bool and assigns it to the Comment field.
func (o *AiFileEntryDtoAllOfSecurity) SetComment(v bool) {
	o.Comment = &v
}

// GetFillForms returns the FillForms field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetFillForms() bool {
	if o == nil || IsNil(o.FillForms) {
		var ret bool
		return ret
	}
	return *o.FillForms
}

// GetFillFormsOk returns a tuple with the FillForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetFillFormsOk() (*bool, bool) {
	if o == nil || IsNil(o.FillForms) {
		return nil, false
	}
	return o.FillForms, true
}

// HasFillForms returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsFillFormsSet() bool {
	if o != nil && !IsNil(o.FillForms) {
		return true
	}

	return false
}

// SetFillForms gets a reference to the given bool and assigns it to the FillForms field.
func (o *AiFileEntryDtoAllOfSecurity) SetFillForms(v bool) {
	o.FillForms = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetReview() bool {
	if o == nil || IsNil(o.Review) {
		var ret bool
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsReviewSet() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given bool and assigns it to the Review field.
func (o *AiFileEntryDtoAllOfSecurity) SetReview(v bool) {
	o.Review = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCreate() bool {
	if o == nil || IsNil(o.Create) {
		var ret bool
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCreateSet() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given bool and assigns it to the Create field.
func (o *AiFileEntryDtoAllOfSecurity) SetCreate(v bool) {
	o.Create = &v
}

// GetCreateFrom returns the CreateFrom field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCreateFrom() bool {
	if o == nil || IsNil(o.CreateFrom) {
		var ret bool
		return ret
	}
	return *o.CreateFrom
}

// GetCreateFromOk returns a tuple with the CreateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCreateFromOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateFrom) {
		return nil, false
	}
	return o.CreateFrom, true
}

// HasCreateFrom returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCreateFromSet() bool {
	if o != nil && !IsNil(o.CreateFrom) {
		return true
	}

	return false
}

// SetCreateFrom gets a reference to the given bool and assigns it to the CreateFrom field.
func (o *AiFileEntryDtoAllOfSecurity) SetCreateFrom(v bool) {
	o.CreateFrom = &v
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEdit() bool {
	if o == nil || IsNil(o.Edit) {
		var ret bool
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditOk() (*bool, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditSet() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given bool and assigns it to the Edit field.
func (o *AiFileEntryDtoAllOfSecurity) SetEdit(v bool) {
	o.Edit = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetDelete() bool {
	if o == nil || IsNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsDeleteSet() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *AiFileEntryDtoAllOfSecurity) SetDelete(v bool) {
	o.Delete = &v
}

// GetCustomFilter returns the CustomFilter field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCustomFilter() bool {
	if o == nil || IsNil(o.CustomFilter) {
		var ret bool
		return ret
	}
	return *o.CustomFilter
}

// GetCustomFilterOk returns a tuple with the CustomFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCustomFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.CustomFilter) {
		return nil, false
	}
	return o.CustomFilter, true
}

// HasCustomFilter returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCustomFilterSet() bool {
	if o != nil && !IsNil(o.CustomFilter) {
		return true
	}

	return false
}

// SetCustomFilter gets a reference to the given bool and assigns it to the CustomFilter field.
func (o *AiFileEntryDtoAllOfSecurity) SetCustomFilter(v bool) {
	o.CustomFilter = &v
}

// GetEditRoom returns the EditRoom field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEditRoom() bool {
	if o == nil || IsNil(o.EditRoom) {
		var ret bool
		return ret
	}
	return *o.EditRoom
}

// GetEditRoomOk returns a tuple with the EditRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditRoomOk() (*bool, bool) {
	if o == nil || IsNil(o.EditRoom) {
		return nil, false
	}
	return o.EditRoom, true
}

// HasEditRoom returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditRoomSet() bool {
	if o != nil && !IsNil(o.EditRoom) {
		return true
	}

	return false
}

// SetEditRoom gets a reference to the given bool and assigns it to the EditRoom field.
func (o *AiFileEntryDtoAllOfSecurity) SetEditRoom(v bool) {
	o.EditRoom = &v
}

// GetRename returns the Rename field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetRename() bool {
	if o == nil || IsNil(o.Rename) {
		var ret bool
		return ret
	}
	return *o.Rename
}

// GetRenameOk returns a tuple with the Rename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetRenameOk() (*bool, bool) {
	if o == nil || IsNil(o.Rename) {
		return nil, false
	}
	return o.Rename, true
}

// HasRename returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsRenameSet() bool {
	if o != nil && !IsNil(o.Rename) {
		return true
	}

	return false
}

// SetRename gets a reference to the given bool and assigns it to the Rename field.
func (o *AiFileEntryDtoAllOfSecurity) SetRename(v bool) {
	o.Rename = &v
}

// GetReadHistory returns the ReadHistory field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetReadHistory() bool {
	if o == nil || IsNil(o.ReadHistory) {
		var ret bool
		return ret
	}
	return *o.ReadHistory
}

// GetReadHistoryOk returns a tuple with the ReadHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetReadHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadHistory) {
		return nil, false
	}
	return o.ReadHistory, true
}

// HasReadHistory returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsReadHistorySet() bool {
	if o != nil && !IsNil(o.ReadHistory) {
		return true
	}

	return false
}

// SetReadHistory gets a reference to the given bool and assigns it to the ReadHistory field.
func (o *AiFileEntryDtoAllOfSecurity) SetReadHistory(v bool) {
	o.ReadHistory = &v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetLock() bool {
	if o == nil || IsNil(o.Lock) {
		var ret bool
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetLockOk() (*bool, bool) {
	if o == nil || IsNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsLockSet() bool {
	if o != nil && !IsNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given bool and assigns it to the Lock field.
func (o *AiFileEntryDtoAllOfSecurity) SetLock(v bool) {
	o.Lock = &v
}

// GetEditHistory returns the EditHistory field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEditHistory() bool {
	if o == nil || IsNil(o.EditHistory) {
		var ret bool
		return ret
	}
	return *o.EditHistory
}

// GetEditHistoryOk returns a tuple with the EditHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.EditHistory) {
		return nil, false
	}
	return o.EditHistory, true
}

// HasEditHistory returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditHistorySet() bool {
	if o != nil && !IsNil(o.EditHistory) {
		return true
	}

	return false
}

// SetEditHistory gets a reference to the given bool and assigns it to the EditHistory field.
func (o *AiFileEntryDtoAllOfSecurity) SetEditHistory(v bool) {
	o.EditHistory = &v
}

// GetCopyTo returns the CopyTo field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCopyTo() bool {
	if o == nil || IsNil(o.CopyTo) {
		var ret bool
		return ret
	}
	return *o.CopyTo
}

// GetCopyToOk returns a tuple with the CopyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCopyToOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyTo) {
		return nil, false
	}
	return o.CopyTo, true
}

// HasCopyTo returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCopyToSet() bool {
	if o != nil && !IsNil(o.CopyTo) {
		return true
	}

	return false
}

// SetCopyTo gets a reference to the given bool and assigns it to the CopyTo field.
func (o *AiFileEntryDtoAllOfSecurity) SetCopyTo(v bool) {
	o.CopyTo = &v
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCopy() bool {
	if o == nil || IsNil(o.Copy) {
		var ret bool
		return ret
	}
	return *o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.Copy) {
		return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCopySet() bool {
	if o != nil && !IsNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given bool and assigns it to the Copy field.
func (o *AiFileEntryDtoAllOfSecurity) SetCopy(v bool) {
	o.Copy = &v
}

// GetMoveTo returns the MoveTo field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetMoveTo() bool {
	if o == nil || IsNil(o.MoveTo) {
		var ret bool
		return ret
	}
	return *o.MoveTo
}

// GetMoveToOk returns a tuple with the MoveTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetMoveToOk() (*bool, bool) {
	if o == nil || IsNil(o.MoveTo) {
		return nil, false
	}
	return o.MoveTo, true
}

// HasMoveTo returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsMoveToSet() bool {
	if o != nil && !IsNil(o.MoveTo) {
		return true
	}

	return false
}

// SetMoveTo gets a reference to the given bool and assigns it to the MoveTo field.
func (o *AiFileEntryDtoAllOfSecurity) SetMoveTo(v bool) {
	o.MoveTo = &v
}

// GetMove returns the Move field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetMove() bool {
	if o == nil || IsNil(o.Move) {
		var ret bool
		return ret
	}
	return *o.Move
}

// GetMoveOk returns a tuple with the Move field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetMoveOk() (*bool, bool) {
	if o == nil || IsNil(o.Move) {
		return nil, false
	}
	return o.Move, true
}

// HasMove returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsMoveSet() bool {
	if o != nil && !IsNil(o.Move) {
		return true
	}

	return false
}

// SetMove gets a reference to the given bool and assigns it to the Move field.
func (o *AiFileEntryDtoAllOfSecurity) SetMove(v bool) {
	o.Move = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetPin() bool {
	if o == nil || IsNil(o.Pin) {
		var ret bool
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetPinOk() (*bool, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsPinSet() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given bool and assigns it to the Pin field.
func (o *AiFileEntryDtoAllOfSecurity) SetPin(v bool) {
	o.Pin = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *AiFileEntryDtoAllOfSecurity) SetMute(v bool) {
	o.Mute = &v
}

// GetEditAccess returns the EditAccess field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEditAccess() bool {
	if o == nil || IsNil(o.EditAccess) {
		var ret bool
		return ret
	}
	return *o.EditAccess
}

// GetEditAccessOk returns a tuple with the EditAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.EditAccess) {
		return nil, false
	}
	return o.EditAccess, true
}

// HasEditAccess returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditAccessSet() bool {
	if o != nil && !IsNil(o.EditAccess) {
		return true
	}

	return false
}

// SetEditAccess gets a reference to the given bool and assigns it to the EditAccess field.
func (o *AiFileEntryDtoAllOfSecurity) SetEditAccess(v bool) {
	o.EditAccess = &v
}

// GetDuplicate returns the Duplicate field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetDuplicate() bool {
	if o == nil || IsNil(o.Duplicate) {
		var ret bool
		return ret
	}
	return *o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetDuplicateOk() (*bool, bool) {
	if o == nil || IsNil(o.Duplicate) {
		return nil, false
	}
	return o.Duplicate, true
}

// HasDuplicate returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsDuplicateSet() bool {
	if o != nil && !IsNil(o.Duplicate) {
		return true
	}

	return false
}

// SetDuplicate gets a reference to the given bool and assigns it to the Duplicate field.
func (o *AiFileEntryDtoAllOfSecurity) SetDuplicate(v bool) {
	o.Duplicate = &v
}

// GetSubmitToFormGallery returns the SubmitToFormGallery field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetSubmitToFormGallery() bool {
	if o == nil || IsNil(o.SubmitToFormGallery) {
		var ret bool
		return ret
	}
	return *o.SubmitToFormGallery
}

// GetSubmitToFormGalleryOk returns a tuple with the SubmitToFormGallery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetSubmitToFormGalleryOk() (*bool, bool) {
	if o == nil || IsNil(o.SubmitToFormGallery) {
		return nil, false
	}
	return o.SubmitToFormGallery, true
}

// HasSubmitToFormGallery returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsSubmitToFormGallerySet() bool {
	if o != nil && !IsNil(o.SubmitToFormGallery) {
		return true
	}

	return false
}

// SetSubmitToFormGallery gets a reference to the given bool and assigns it to the SubmitToFormGallery field.
func (o *AiFileEntryDtoAllOfSecurity) SetSubmitToFormGallery(v bool) {
	o.SubmitToFormGallery = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsDownloadSet() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *AiFileEntryDtoAllOfSecurity) SetDownload(v bool) {
	o.Download = &v
}

// GetConvert returns the Convert field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetConvert() bool {
	if o == nil || IsNil(o.Convert) {
		var ret bool
		return ret
	}
	return *o.Convert
}

// GetConvertOk returns a tuple with the Convert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.Convert) {
		return nil, false
	}
	return o.Convert, true
}

// HasConvert returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsConvertSet() bool {
	if o != nil && !IsNil(o.Convert) {
		return true
	}

	return false
}

// SetConvert gets a reference to the given bool and assigns it to the Convert field.
func (o *AiFileEntryDtoAllOfSecurity) SetConvert(v bool) {
	o.Convert = &v
}

// GetCopySharedLink returns the CopySharedLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCopySharedLink() bool {
	if o == nil || IsNil(o.CopySharedLink) {
		var ret bool
		return ret
	}
	return *o.CopySharedLink
}

// GetCopySharedLinkOk returns a tuple with the CopySharedLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCopySharedLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.CopySharedLink) {
		return nil, false
	}
	return o.CopySharedLink, true
}

// HasCopySharedLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCopySharedLinkSet() bool {
	if o != nil && !IsNil(o.CopySharedLink) {
		return true
	}

	return false
}

// SetCopySharedLink gets a reference to the given bool and assigns it to the CopySharedLink field.
func (o *AiFileEntryDtoAllOfSecurity) SetCopySharedLink(v bool) {
	o.CopySharedLink = &v
}

// GetReadLinks returns the ReadLinks field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetReadLinks() bool {
	if o == nil || IsNil(o.ReadLinks) {
		var ret bool
		return ret
	}
	return *o.ReadLinks
}

// GetReadLinksOk returns a tuple with the ReadLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetReadLinksOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadLinks) {
		return nil, false
	}
	return o.ReadLinks, true
}

// HasReadLinks returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsReadLinksSet() bool {
	if o != nil && !IsNil(o.ReadLinks) {
		return true
	}

	return false
}

// SetReadLinks gets a reference to the given bool and assigns it to the ReadLinks field.
func (o *AiFileEntryDtoAllOfSecurity) SetReadLinks(v bool) {
	o.ReadLinks = &v
}

// GetReconnect returns the Reconnect field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetReconnect() bool {
	if o == nil || IsNil(o.Reconnect) {
		var ret bool
		return ret
	}
	return *o.Reconnect
}

// GetReconnectOk returns a tuple with the Reconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetReconnectOk() (*bool, bool) {
	if o == nil || IsNil(o.Reconnect) {
		return nil, false
	}
	return o.Reconnect, true
}

// HasReconnect returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsReconnectSet() bool {
	if o != nil && !IsNil(o.Reconnect) {
		return true
	}

	return false
}

// SetReconnect gets a reference to the given bool and assigns it to the Reconnect field.
func (o *AiFileEntryDtoAllOfSecurity) SetReconnect(v bool) {
	o.Reconnect = &v
}

// GetCreateRoomFrom returns the CreateRoomFrom field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCreateRoomFrom() bool {
	if o == nil || IsNil(o.CreateRoomFrom) {
		var ret bool
		return ret
	}
	return *o.CreateRoomFrom
}

// GetCreateRoomFromOk returns a tuple with the CreateRoomFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCreateRoomFromOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateRoomFrom) {
		return nil, false
	}
	return o.CreateRoomFrom, true
}

// HasCreateRoomFrom returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCreateRoomFromSet() bool {
	if o != nil && !IsNil(o.CreateRoomFrom) {
		return true
	}

	return false
}

// SetCreateRoomFrom gets a reference to the given bool and assigns it to the CreateRoomFrom field.
func (o *AiFileEntryDtoAllOfSecurity) SetCreateRoomFrom(v bool) {
	o.CreateRoomFrom = &v
}

// GetCopyLink returns the CopyLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCopyLink() bool {
	if o == nil || IsNil(o.CopyLink) {
		var ret bool
		return ret
	}
	return *o.CopyLink
}

// GetCopyLinkOk returns a tuple with the CopyLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCopyLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyLink) {
		return nil, false
	}
	return o.CopyLink, true
}

// HasCopyLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCopyLinkSet() bool {
	if o != nil && !IsNil(o.CopyLink) {
		return true
	}

	return false
}

// SetCopyLink gets a reference to the given bool and assigns it to the CopyLink field.
func (o *AiFileEntryDtoAllOfSecurity) SetCopyLink(v bool) {
	o.CopyLink = &v
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEmbed() bool {
	if o == nil || IsNil(o.Embed) {
		var ret bool
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEmbedOk() (*bool, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEmbedSet() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given bool and assigns it to the Embed field.
func (o *AiFileEntryDtoAllOfSecurity) SetEmbed(v bool) {
	o.Embed = &v
}

// GetChangeOwner returns the ChangeOwner field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetChangeOwner() bool {
	if o == nil || IsNil(o.ChangeOwner) {
		var ret bool
		return ret
	}
	return *o.ChangeOwner
}

// GetChangeOwnerOk returns a tuple with the ChangeOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetChangeOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.ChangeOwner) {
		return nil, false
	}
	return o.ChangeOwner, true
}

// HasChangeOwner returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsChangeOwnerSet() bool {
	if o != nil && !IsNil(o.ChangeOwner) {
		return true
	}

	return false
}

// SetChangeOwner gets a reference to the given bool and assigns it to the ChangeOwner field.
func (o *AiFileEntryDtoAllOfSecurity) SetChangeOwner(v bool) {
	o.ChangeOwner = &v
}

// GetIndexExport returns the IndexExport field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetIndexExport() bool {
	if o == nil || IsNil(o.IndexExport) {
		var ret bool
		return ret
	}
	return *o.IndexExport
}

// GetIndexExportOk returns a tuple with the IndexExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetIndexExportOk() (*bool, bool) {
	if o == nil || IsNil(o.IndexExport) {
		return nil, false
	}
	return o.IndexExport, true
}

// HasIndexExport returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsIndexExportSet() bool {
	if o != nil && !IsNil(o.IndexExport) {
		return true
	}

	return false
}

// SetIndexExport gets a reference to the given bool and assigns it to the IndexExport field.
func (o *AiFileEntryDtoAllOfSecurity) SetIndexExport(v bool) {
	o.IndexExport = &v
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling) {
		var ret bool
		return ret
	}
	return *o.StartFilling
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetStartFillingOk() (*bool, bool) {
	if o == nil || IsNil(o.StartFilling) {
		return nil, false
	}
	return o.StartFilling, true
}

// HasStartFilling returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsStartFillingSet() bool {
	if o != nil && !IsNil(o.StartFilling) {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given bool and assigns it to the StartFilling field.
func (o *AiFileEntryDtoAllOfSecurity) SetStartFilling(v bool) {
	o.StartFilling = &v
}

// GetFillingStatus returns the FillingStatus field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetFillingStatus() bool {
	if o == nil || IsNil(o.FillingStatus) {
		var ret bool
		return ret
	}
	return *o.FillingStatus
}

// GetFillingStatusOk returns a tuple with the FillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetFillingStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FillingStatus) {
		return nil, false
	}
	return o.FillingStatus, true
}

// HasFillingStatus returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsFillingStatusSet() bool {
	if o != nil && !IsNil(o.FillingStatus) {
		return true
	}

	return false
}

// SetFillingStatus gets a reference to the given bool and assigns it to the FillingStatus field.
func (o *AiFileEntryDtoAllOfSecurity) SetFillingStatus(v bool) {
	o.FillingStatus = &v
}

// GetResetFilling returns the ResetFilling field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetResetFilling() bool {
	if o == nil || IsNil(o.ResetFilling) {
		var ret bool
		return ret
	}
	return *o.ResetFilling
}

// GetResetFillingOk returns a tuple with the ResetFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetResetFillingOk() (*bool, bool) {
	if o == nil || IsNil(o.ResetFilling) {
		return nil, false
	}
	return o.ResetFilling, true
}

// HasResetFilling returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsResetFillingSet() bool {
	if o != nil && !IsNil(o.ResetFilling) {
		return true
	}

	return false
}

// SetResetFilling gets a reference to the given bool and assigns it to the ResetFilling field.
func (o *AiFileEntryDtoAllOfSecurity) SetResetFilling(v bool) {
	o.ResetFilling = &v
}

// GetStopFilling returns the StopFilling field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetStopFilling() bool {
	if o == nil || IsNil(o.StopFilling) {
		var ret bool
		return ret
	}
	return *o.StopFilling
}

// GetStopFillingOk returns a tuple with the StopFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetStopFillingOk() (*bool, bool) {
	if o == nil || IsNil(o.StopFilling) {
		return nil, false
	}
	return o.StopFilling, true
}

// HasStopFilling returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsStopFillingSet() bool {
	if o != nil && !IsNil(o.StopFilling) {
		return true
	}

	return false
}

// SetStopFilling gets a reference to the given bool and assigns it to the StopFilling field.
func (o *AiFileEntryDtoAllOfSecurity) SetStopFilling(v bool) {
	o.StopFilling = &v
}

// GetOpenForm returns the OpenForm field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetOpenForm() bool {
	if o == nil || IsNil(o.OpenForm) {
		var ret bool
		return ret
	}
	return *o.OpenForm
}

// GetOpenFormOk returns a tuple with the OpenForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetOpenFormOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenForm) {
		return nil, false
	}
	return o.OpenForm, true
}

// HasOpenForm returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsOpenFormSet() bool {
	if o != nil && !IsNil(o.OpenForm) {
		return true
	}

	return false
}

// SetOpenForm gets a reference to the given bool and assigns it to the OpenForm field.
func (o *AiFileEntryDtoAllOfSecurity) SetOpenForm(v bool) {
	o.OpenForm = &v
}

// GetEditInternal returns the EditInternal field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEditInternal() bool {
	if o == nil || IsNil(o.EditInternal) {
		var ret bool
		return ret
	}
	return *o.EditInternal
}

// GetEditInternalOk returns a tuple with the EditInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.EditInternal) {
		return nil, false
	}
	return o.EditInternal, true
}

// HasEditInternal returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditInternalSet() bool {
	if o != nil && !IsNil(o.EditInternal) {
		return true
	}

	return false
}

// SetEditInternal gets a reference to the given bool and assigns it to the EditInternal field.
func (o *AiFileEntryDtoAllOfSecurity) SetEditInternal(v bool) {
	o.EditInternal = &v
}

// GetEditExpiration returns the EditExpiration field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetEditExpiration() bool {
	if o == nil || IsNil(o.EditExpiration) {
		var ret bool
		return ret
	}
	return *o.EditExpiration
}

// GetEditExpirationOk returns a tuple with the EditExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetEditExpirationOk() (*bool, bool) {
	if o == nil || IsNil(o.EditExpiration) {
		return nil, false
	}
	return o.EditExpiration, true
}

// HasEditExpiration returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsEditExpirationSet() bool {
	if o != nil && !IsNil(o.EditExpiration) {
		return true
	}

	return false
}

// SetEditExpiration gets a reference to the given bool and assigns it to the EditExpiration field.
func (o *AiFileEntryDtoAllOfSecurity) SetEditExpiration(v bool) {
	o.EditExpiration = &v
}

// GetVectorization returns the Vectorization field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetVectorization() bool {
	if o == nil || IsNil(o.Vectorization) {
		var ret bool
		return ret
	}
	return *o.Vectorization
}

// GetVectorizationOk returns a tuple with the Vectorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetVectorizationOk() (*bool, bool) {
	if o == nil || IsNil(o.Vectorization) {
		return nil, false
	}
	return o.Vectorization, true
}

// HasVectorization returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsVectorizationSet() bool {
	if o != nil && !IsNil(o.Vectorization) {
		return true
	}

	return false
}

// SetVectorization gets a reference to the given bool and assigns it to the Vectorization field.
func (o *AiFileEntryDtoAllOfSecurity) SetVectorization(v bool) {
	o.Vectorization = &v
}

// GetAskAi returns the AskAi field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetAskAi() bool {
	if o == nil || IsNil(o.AskAi) {
		var ret bool
		return ret
	}
	return *o.AskAi
}

// GetAskAiOk returns a tuple with the AskAi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetAskAiOk() (*bool, bool) {
	if o == nil || IsNil(o.AskAi) {
		return nil, false
	}
	return o.AskAi, true
}

// HasAskAi returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsAskAiSet() bool {
	if o != nil && !IsNil(o.AskAi) {
		return true
	}

	return false
}

// SetAskAi gets a reference to the given bool and assigns it to the AskAi field.
func (o *AiFileEntryDtoAllOfSecurity) SetAskAi(v bool) {
	o.AskAi = &v
}

// GetUseChat returns the UseChat field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetUseChat() bool {
	if o == nil || IsNil(o.UseChat) {
		var ret bool
		return ret
	}
	return *o.UseChat
}

// GetUseChatOk returns a tuple with the UseChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetUseChatOk() (*bool, bool) {
	if o == nil || IsNil(o.UseChat) {
		return nil, false
	}
	return o.UseChat, true
}

// HasUseChat returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsUseChatSet() bool {
	if o != nil && !IsNil(o.UseChat) {
		return true
	}

	return false
}

// SetUseChat gets a reference to the given bool and assigns it to the UseChat field.
func (o *AiFileEntryDtoAllOfSecurity) SetUseChat(v bool) {
	o.UseChat = &v
}

// GetUpdateXlsx returns the UpdateXlsx field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetUpdateXlsx() bool {
	if o == nil || IsNil(o.UpdateXlsx) {
		var ret bool
		return ret
	}
	return *o.UpdateXlsx
}

// GetUpdateXlsxOk returns a tuple with the UpdateXlsx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetUpdateXlsxOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateXlsx) {
		return nil, false
	}
	return o.UpdateXlsx, true
}

// HasUpdateXlsx returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsUpdateXlsxSet() bool {
	if o != nil && !IsNil(o.UpdateXlsx) {
		return true
	}

	return false
}

// SetUpdateXlsx gets a reference to the given bool and assigns it to the UpdateXlsx field.
func (o *AiFileEntryDtoAllOfSecurity) SetUpdateXlsx(v bool) {
	o.UpdateXlsx = &v
}

// GetAnalyzeResponses returns the AnalyzeResponses field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetAnalyzeResponses() bool {
	if o == nil || IsNil(o.AnalyzeResponses) {
		var ret bool
		return ret
	}
	return *o.AnalyzeResponses
}

// GetAnalyzeResponsesOk returns a tuple with the AnalyzeResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetAnalyzeResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.AnalyzeResponses) {
		return nil, false
	}
	return o.AnalyzeResponses, true
}

// HasAnalyzeResponses returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsAnalyzeResponsesSet() bool {
	if o != nil && !IsNil(o.AnalyzeResponses) {
		return true
	}

	return false
}

// SetAnalyzeResponses gets a reference to the given bool and assigns it to the AnalyzeResponses field.
func (o *AiFileEntryDtoAllOfSecurity) SetAnalyzeResponses(v bool) {
	o.AnalyzeResponses = &v
}

// GetCanUseAi returns the CanUseAi field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetCanUseAi() bool {
	if o == nil || IsNil(o.CanUseAi) {
		var ret bool
		return ret
	}
	return *o.CanUseAi
}

// GetCanUseAiOk returns a tuple with the CanUseAi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetCanUseAiOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseAi) {
		return nil, false
	}
	return o.CanUseAi, true
}

// HasCanUseAi returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsCanUseAiSet() bool {
	if o != nil && !IsNil(o.CanUseAi) {
		return true
	}

	return false
}

// SetCanUseAi gets a reference to the given bool and assigns it to the CanUseAi field.
func (o *AiFileEntryDtoAllOfSecurity) SetCanUseAi(v bool) {
	o.CanUseAi = &v
}

// GetHistoryExport returns the HistoryExport field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfSecurity) GetHistoryExport() bool {
	if o == nil || IsNil(o.HistoryExport) {
		var ret bool
		return ret
	}
	return *o.HistoryExport
}

// GetHistoryExportOk returns a tuple with the HistoryExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfSecurity) GetHistoryExportOk() (*bool, bool) {
	if o == nil || IsNil(o.HistoryExport) {
		return nil, false
	}
	return o.HistoryExport, true
}

// HasHistoryExport returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfSecurity) IsHistoryExportSet() bool {
	if o != nil && !IsNil(o.HistoryExport) {
		return true
	}

	return false
}

// SetHistoryExport gets a reference to the given bool and assigns it to the HistoryExport field.
func (o *AiFileEntryDtoAllOfSecurity) SetHistoryExport(v bool) {
	o.HistoryExport = &v
}

func (o AiFileEntryDtoAllOfSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileEntryDtoAllOfSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Read) {
		toSerialize["Read"] = o.Read
	}
	if !IsNil(o.Comment) {
		toSerialize["Comment"] = o.Comment
	}
	if !IsNil(o.FillForms) {
		toSerialize["FillForms"] = o.FillForms
	}
	if !IsNil(o.Review) {
		toSerialize["Review"] = o.Review
	}
	if !IsNil(o.Create) {
		toSerialize["Create"] = o.Create
	}
	if !IsNil(o.CreateFrom) {
		toSerialize["CreateFrom"] = o.CreateFrom
	}
	if !IsNil(o.Edit) {
		toSerialize["Edit"] = o.Edit
	}
	if !IsNil(o.Delete) {
		toSerialize["Delete"] = o.Delete
	}
	if !IsNil(o.CustomFilter) {
		toSerialize["CustomFilter"] = o.CustomFilter
	}
	if !IsNil(o.EditRoom) {
		toSerialize["EditRoom"] = o.EditRoom
	}
	if !IsNil(o.Rename) {
		toSerialize["Rename"] = o.Rename
	}
	if !IsNil(o.ReadHistory) {
		toSerialize["ReadHistory"] = o.ReadHistory
	}
	if !IsNil(o.Lock) {
		toSerialize["Lock"] = o.Lock
	}
	if !IsNil(o.EditHistory) {
		toSerialize["EditHistory"] = o.EditHistory
	}
	if !IsNil(o.CopyTo) {
		toSerialize["CopyTo"] = o.CopyTo
	}
	if !IsNil(o.Copy) {
		toSerialize["Copy"] = o.Copy
	}
	if !IsNil(o.MoveTo) {
		toSerialize["MoveTo"] = o.MoveTo
	}
	if !IsNil(o.Move) {
		toSerialize["Move"] = o.Move
	}
	if !IsNil(o.Pin) {
		toSerialize["Pin"] = o.Pin
	}
	if !IsNil(o.Mute) {
		toSerialize["Mute"] = o.Mute
	}
	if !IsNil(o.EditAccess) {
		toSerialize["EditAccess"] = o.EditAccess
	}
	if !IsNil(o.Duplicate) {
		toSerialize["Duplicate"] = o.Duplicate
	}
	if !IsNil(o.SubmitToFormGallery) {
		toSerialize["SubmitToFormGallery"] = o.SubmitToFormGallery
	}
	if !IsNil(o.Download) {
		toSerialize["Download"] = o.Download
	}
	if !IsNil(o.Convert) {
		toSerialize["Convert"] = o.Convert
	}
	if !IsNil(o.CopySharedLink) {
		toSerialize["CopySharedLink"] = o.CopySharedLink
	}
	if !IsNil(o.ReadLinks) {
		toSerialize["ReadLinks"] = o.ReadLinks
	}
	if !IsNil(o.Reconnect) {
		toSerialize["Reconnect"] = o.Reconnect
	}
	if !IsNil(o.CreateRoomFrom) {
		toSerialize["CreateRoomFrom"] = o.CreateRoomFrom
	}
	if !IsNil(o.CopyLink) {
		toSerialize["CopyLink"] = o.CopyLink
	}
	if !IsNil(o.Embed) {
		toSerialize["Embed"] = o.Embed
	}
	if !IsNil(o.ChangeOwner) {
		toSerialize["ChangeOwner"] = o.ChangeOwner
	}
	if !IsNil(o.IndexExport) {
		toSerialize["IndexExport"] = o.IndexExport
	}
	if !IsNil(o.StartFilling) {
		toSerialize["StartFilling"] = o.StartFilling
	}
	if !IsNil(o.FillingStatus) {
		toSerialize["FillingStatus"] = o.FillingStatus
	}
	if !IsNil(o.ResetFilling) {
		toSerialize["ResetFilling"] = o.ResetFilling
	}
	if !IsNil(o.StopFilling) {
		toSerialize["StopFilling"] = o.StopFilling
	}
	if !IsNil(o.OpenForm) {
		toSerialize["OpenForm"] = o.OpenForm
	}
	if !IsNil(o.EditInternal) {
		toSerialize["EditInternal"] = o.EditInternal
	}
	if !IsNil(o.EditExpiration) {
		toSerialize["EditExpiration"] = o.EditExpiration
	}
	if !IsNil(o.Vectorization) {
		toSerialize["Vectorization"] = o.Vectorization
	}
	if !IsNil(o.AskAi) {
		toSerialize["AskAi"] = o.AskAi
	}
	if !IsNil(o.UseChat) {
		toSerialize["UseChat"] = o.UseChat
	}
	if !IsNil(o.UpdateXlsx) {
		toSerialize["UpdateXlsx"] = o.UpdateXlsx
	}
	if !IsNil(o.AnalyzeResponses) {
		toSerialize["AnalyzeResponses"] = o.AnalyzeResponses
	}
	if !IsNil(o.CanUseAi) {
		toSerialize["CanUseAi"] = o.CanUseAi
	}
	if !IsNil(o.HistoryExport) {
		toSerialize["HistoryExport"] = o.HistoryExport
	}
	return toSerialize, nil
}

type NullableAiFileEntryDtoAllOfSecurity struct {
	value *AiFileEntryDtoAllOfSecurity
	isSet bool
}

func (v NullableAiFileEntryDtoAllOfSecurity) Get() *AiFileEntryDtoAllOfSecurity {
	return v.value
}

func (v *NullableAiFileEntryDtoAllOfSecurity) Set(val *AiFileEntryDtoAllOfSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryDtoAllOfSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryDtoAllOfSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryDtoAllOfSecurity(val *AiFileEntryDtoAllOfSecurity) *NullableAiFileEntryDtoAllOfSecurity {
	return &NullableAiFileEntryDtoAllOfSecurity{value: val, isSet: true}
}

func (v NullableAiFileEntryDtoAllOfSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryDtoAllOfSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

