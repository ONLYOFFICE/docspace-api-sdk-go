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

// checks if the BackupProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupProgress{}

// BackupProgress The state of one backup or restoring job.
type BackupProgress struct {
	// Specifies whether the job has stopped running. This is the field to poll: true means the job will not  change any more, whether it succeeded, failed or was cancelled, and `status` tells which of the three  it is.
	IsCompleted *bool `json:"isCompleted,omitempty"`
	// The share of the job that is already done, from 0 to 100. A job that has only been queued reports 0,  because the work starts when a separate worker service picks it up.
	Progress *int32 `json:"progress,omitempty"`
	// The message of the error that stopped the job. It is an empty string, not null, while the job runs  and after a job that succeeded, so the sign of a failure is a non-empty value - and this is the only  place where the reason is reported.
	Error NullableString `json:"error,omitempty"`
	// A message about a job that stopped without failing: it names the entry inside the archive that lists  the files which could not be read, when a backup finished without some of them, and it says so when  the job was cancelled. It is an empty string otherwise, and it is only ever filled in for a backup  job - a cancelled restoring job leaves it empty.
	Warning NullableString `json:"warning,omitempty"`
	// The link to download the stored archive. It is an empty string until the archive has been uploaded,  and it is only ever filled in for a backup job, never for a restoring one.
	Link NullableString `json:"link,omitempty"`
	// The ID of the portal the job belongs to, or -1 for a job that covers the whole server.
	TenantId *int32 `json:"tenantId,omitempty"`
	// Whether this is a backup or a restoring job, reported as a number rather than as a name.
	BackupProgressEnum *BackupProgressEnum `json:"backupProgressEnum,omitempty"`
	// The state of the job: `Created` while it waits for a worker to pick it up, `Running` while it works,  `Completed` once it has finished on its own, `Canceled` after it was cancelled, and `Failted` when it  stopped on an error, in which case `error` carries the reason. Reported as a number rather than as a  name.
	Status *DistributedTaskStatus `json:"status,omitempty"`
	// The ID of the job. It is the handle to poll this operation with, and for a backup job it also becomes  the `id` of the record in `GET api/2.0/backup/getbackuphistory`.
	TaskId NullableString `json:"taskId,omitempty"`
}

// NewBackupProgress instantiates a new BackupProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupProgress() *BackupProgress {
	this := BackupProgress{}
	return &this
}

// NewBackupProgressWithDefaults instantiates a new BackupProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupProgressWithDefaults() *BackupProgress {
	this := BackupProgress{}
	return &this
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *BackupProgress) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProgress) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *BackupProgress) IsIsCompletedSet() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *BackupProgress) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *BackupProgress) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProgress) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *BackupProgress) IsProgressSet() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *BackupProgress) SetProgress(v int32) {
	o.Progress = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProgress) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProgress) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *BackupProgress) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *BackupProgress) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *BackupProgress) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *BackupProgress) UnsetError() {
	o.Error.Unset()
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProgress) GetWarning() string {
	if o == nil || IsNil(o.Warning.Get()) {
		var ret string
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProgress) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *BackupProgress) IsWarningSet() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableString and assigns it to the Warning field.
func (o *BackupProgress) SetWarning(v string) {
	o.Warning.Set(&v)
}
// SetWarningNil sets the value for Warning to be an explicit nil
func (o *BackupProgress) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *BackupProgress) UnsetWarning() {
	o.Warning.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProgress) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProgress) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *BackupProgress) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *BackupProgress) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *BackupProgress) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *BackupProgress) UnsetLink() {
	o.Link.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *BackupProgress) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProgress) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *BackupProgress) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *BackupProgress) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetBackupProgressEnum returns the BackupProgressEnum field value if set, zero value otherwise.
func (o *BackupProgress) GetBackupProgressEnum() BackupProgressEnum {
	if o == nil || IsNil(o.BackupProgressEnum) {
		var ret BackupProgressEnum
		return ret
	}
	return *o.BackupProgressEnum
}

// GetBackupProgressEnumOk returns a tuple with the BackupProgressEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProgress) GetBackupProgressEnumOk() (*BackupProgressEnum, bool) {
	if o == nil || IsNil(o.BackupProgressEnum) {
		return nil, false
	}
	return o.BackupProgressEnum, true
}

// HasBackupProgressEnum returns a boolean if a field has been set.
func (o *BackupProgress) IsBackupProgressEnumSet() bool {
	if o != nil && !IsNil(o.BackupProgressEnum) {
		return true
	}

	return false
}

// SetBackupProgressEnum gets a reference to the given BackupProgressEnum and assigns it to the BackupProgressEnum field.
func (o *BackupProgress) SetBackupProgressEnum(v BackupProgressEnum) {
	o.BackupProgressEnum = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupProgress) GetStatus() DistributedTaskStatus {
	if o == nil || IsNil(o.Status) {
		var ret DistributedTaskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProgress) GetStatusOk() (*DistributedTaskStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupProgress) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DistributedTaskStatus and assigns it to the Status field.
func (o *BackupProgress) SetStatus(v DistributedTaskStatus) {
	o.Status = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProgress) GetTaskId() string {
	if o == nil || IsNil(o.TaskId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProgress) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *BackupProgress) IsTaskIdSet() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *BackupProgress) SetTaskId(v string) {
	o.TaskId.Set(&v)
}
// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *BackupProgress) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *BackupProgress) UnsetTaskId() {
	o.TaskId.Unset()
}

func (o BackupProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsCompleted) {
		toSerialize["isCompleted"] = o.IsCompleted
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.BackupProgressEnum) {
		toSerialize["backupProgressEnum"] = o.BackupProgressEnum
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	return toSerialize, nil
}

type NullableBackupProgress struct {
	value *BackupProgress
	isSet bool
}

func (v NullableBackupProgress) Get() *BackupProgress {
	return v.value
}

func (v *NullableBackupProgress) Set(val *BackupProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProgress(val *BackupProgress) *NullableBackupProgress {
	return &NullableBackupProgress{value: val, isSet: true}
}

func (v NullableBackupProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

