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
	"bytes"
	"fmt"
)

// checks if the AiThreadMessageLikeContentAnyOfInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLikeContentAnyOfInner{}

// AiThreadMessageLikeContentAnyOfInner struct for AiThreadMessageLikeContentAnyOfInner
type AiThreadMessageLikeContentAnyOfInner struct {
	// Part kind: `text`, `reasoning`, `source`, `image`, `file`, `audio`, `data` or `tool-call`.
	Type string `json:"type"`
	// The text of a `text` or `reasoning` part.
	Text *string `json:"text,omitempty"`
	// Id of the part this one belongs to, on `text`, `reasoning`, `source` and `tool-call` parts. On a `reasoning` part the chat library keeps the provider's thinking signature here.
	ParentId *string `json:"parentId,omitempty"`
	// A `source` part's source kind; always `url`.
	SourceType *string `json:"sourceType,omitempty"`
	// A `source` part's id.
	Id *string `json:"id,omitempty"`
	// A `source` part's URL.
	Url *string `json:"url,omitempty"`
	// A `source` part's title.
	Title *string `json:"title,omitempty"`
	// An `image` part's picture, as a URL (usually a `data:` URL). For an image the user attached, a JSON string `{ ref, title, kind: image }` naming the stored attachment instead.
	Image *string `json:"image,omitempty"`
	// A `file` part's media type. For a file the user attached, a JSON string `{ ref, title, kind: file, path, type }` naming the stored attachment instead - the chat library carries the reference here.
	MimeType *string `json:"mimeType,omitempty"`
	// A `file` part's content; empty when the part references an attachment. (A `data` part would carry its payload here too, as any JSON value; neither the chat library nor this service writes one.)
	Data *string `json:"data,omitempty"`
	// The file name of a `file` or `image` part.
	Filename *string `json:"filename,omitempty"`
	// An `image` part's display name, or a `data` part's name.
	Name *string `json:"name,omitempty"`
	Audio *AiThreadMessageLikeContentAnyOfInnerAudio `json:"audio,omitempty"`
	// A `tool-call` part's call id.
	ToolCallId *string `json:"toolCallId,omitempty"`
	// A `tool-call` part's tool name.
	ToolName *string `json:"toolName,omitempty"`
	// A `tool-call` part's arguments.
	Args map[string]interface{} `json:"args,omitempty"`
	// A `tool-call` part's arguments as the model streamed them, JSON text.
	ArgsText *string `json:"argsText,omitempty"`
	Result interface{} `json:"result,omitempty"`
	// Whether a `tool-call` part's result is an error.
	IsError *bool `json:"isError,omitempty"`
	Artifact interface{} `json:"artifact,omitempty"`
	Interrupt *AiThreadMessageLikeContentAnyOfInnerInterrupt `json:"interrupt,omitempty"`
	// Messages of a sub-agent run inside a `tool-call` part.
	Messages []map[string]interface{} `json:"messages,omitempty"`
	// Provider-specific data of a `tool-call` part - for a Gemini call, `{ thoughtSignature }`, which has to be sent back with the call.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _AiThreadMessageLikeContentAnyOfInner AiThreadMessageLikeContentAnyOfInner

// NewAiThreadMessageLikeContentAnyOfInner instantiates a new AiThreadMessageLikeContentAnyOfInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLikeContentAnyOfInner(type_ string) *AiThreadMessageLikeContentAnyOfInner {
	this := AiThreadMessageLikeContentAnyOfInner{}
	this.Type = type_
	return &this
}

// NewAiThreadMessageLikeContentAnyOfInnerWithDefaults instantiates a new AiThreadMessageLikeContentAnyOfInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeContentAnyOfInnerWithDefaults() *AiThreadMessageLikeContentAnyOfInner {
	this := AiThreadMessageLikeContentAnyOfInner{}
	return &this
}

// GetType returns the Type field value
func (o *AiThreadMessageLikeContentAnyOfInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiThreadMessageLikeContentAnyOfInner) SetType(v string) {
	o.Type = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsTextSet() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetText(v string) {
	o.Text = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsParentIdSet() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetParentId(v string) {
	o.ParentId = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsSourceTypeSet() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetSourceType(v string) {
	o.SourceType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsUrlSet() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetUrl(v string) {
	o.Url = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetTitle(v string) {
	o.Title = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsImageSet() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetImage(v string) {
	o.Image = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsMimeTypeSet() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetMimeType(v string) {
	o.MimeType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetData(v string) {
	o.Data = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsFilenameSet() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetFilename(v string) {
	o.Filename = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetName(v string) {
	o.Name = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetAudio() AiThreadMessageLikeContentAnyOfInnerAudio {
	if o == nil || IsNil(o.Audio) {
		var ret AiThreadMessageLikeContentAnyOfInnerAudio
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetAudioOk() (*AiThreadMessageLikeContentAnyOfInnerAudio, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsAudioSet() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given AiThreadMessageLikeContentAnyOfInnerAudio and assigns it to the Audio field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetAudio(v AiThreadMessageLikeContentAnyOfInnerAudio) {
	o.Audio = &v
}

// GetToolCallId returns the ToolCallId field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetToolCallId() string {
	if o == nil || IsNil(o.ToolCallId) {
		var ret string
		return ret
	}
	return *o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetToolCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ToolCallId) {
		return nil, false
	}
	return o.ToolCallId, true
}

// HasToolCallId returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsToolCallIdSet() bool {
	if o != nil && !IsNil(o.ToolCallId) {
		return true
	}

	return false
}

// SetToolCallId gets a reference to the given string and assigns it to the ToolCallId field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetToolCallId(v string) {
	o.ToolCallId = &v
}

// GetToolName returns the ToolName field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetToolName() string {
	if o == nil || IsNil(o.ToolName) {
		var ret string
		return ret
	}
	return *o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetToolNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToolName) {
		return nil, false
	}
	return o.ToolName, true
}

// HasToolName returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsToolNameSet() bool {
	if o != nil && !IsNil(o.ToolName) {
		return true
	}

	return false
}

// SetToolName gets a reference to the given string and assigns it to the ToolName field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetToolName(v string) {
	o.ToolName = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetArgs() map[string]interface{} {
	if o == nil || IsNil(o.Args) {
		var ret map[string]interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return map[string]interface{}{}, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsArgsSet() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetArgs(v map[string]interface{}) {
	o.Args = v
}

// GetArgsText returns the ArgsText field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsText() string {
	if o == nil || IsNil(o.ArgsText) {
		var ret string
		return ret
	}
	return *o.ArgsText
}

// GetArgsTextOk returns a tuple with the ArgsText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsTextOk() (*string, bool) {
	if o == nil || IsNil(o.ArgsText) {
		return nil, false
	}
	return o.ArgsText, true
}

// HasArgsText returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsArgsTextSet() bool {
	if o != nil && !IsNil(o.ArgsText) {
		return true
	}

	return false
}

// SetArgsText gets a reference to the given string and assigns it to the ArgsText field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetArgsText(v string) {
	o.ArgsText = &v
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiThreadMessageLikeContentAnyOfInner) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiThreadMessageLikeContentAnyOfInner) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsResultSet() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given interface{} and assigns it to the Result field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetResult(v interface{}) {
	o.Result = v
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsIsErrorSet() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetIsError(v bool) {
	o.IsError = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiThreadMessageLikeContentAnyOfInner) GetArtifact() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiThreadMessageLikeContentAnyOfInner) GetArtifactOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return &o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsArtifactSet() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given interface{} and assigns it to the Artifact field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetArtifact(v interface{}) {
	o.Artifact = v
}

// GetInterrupt returns the Interrupt field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetInterrupt() AiThreadMessageLikeContentAnyOfInnerInterrupt {
	if o == nil || IsNil(o.Interrupt) {
		var ret AiThreadMessageLikeContentAnyOfInnerInterrupt
		return ret
	}
	return *o.Interrupt
}

// GetInterruptOk returns a tuple with the Interrupt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetInterruptOk() (*AiThreadMessageLikeContentAnyOfInnerInterrupt, bool) {
	if o == nil || IsNil(o.Interrupt) {
		return nil, false
	}
	return o.Interrupt, true
}

// HasInterrupt returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsInterruptSet() bool {
	if o != nil && !IsNil(o.Interrupt) {
		return true
	}

	return false
}

// SetInterrupt gets a reference to the given AiThreadMessageLikeContentAnyOfInnerInterrupt and assigns it to the Interrupt field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetInterrupt(v AiThreadMessageLikeContentAnyOfInnerInterrupt) {
	o.Interrupt = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMessages() []map[string]interface{} {
	if o == nil || IsNil(o.Messages) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMessagesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsMessagesSet() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []map[string]interface{} and assigns it to the Messages field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetMessages(v []map[string]interface{}) {
	o.Messages = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsMetadataSet() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o AiThreadMessageLikeContentAnyOfInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLikeContentAnyOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Audio) {
		toSerialize["audio"] = o.Audio
	}
	if !IsNil(o.ToolCallId) {
		toSerialize["toolCallId"] = o.ToolCallId
	}
	if !IsNil(o.ToolName) {
		toSerialize["toolName"] = o.ToolName
	}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.ArgsText) {
		toSerialize["argsText"] = o.ArgsText
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if o.Artifact != nil {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.Interrupt) {
		toSerialize["interrupt"] = o.Interrupt
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *AiThreadMessageLikeContentAnyOfInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAiThreadMessageLikeContentAnyOfInner := _AiThreadMessageLikeContentAnyOfInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLikeContentAnyOfInner)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLikeContentAnyOfInner(varAiThreadMessageLikeContentAnyOfInner)

	return err
}

type NullableAiThreadMessageLikeContentAnyOfInner struct {
	value *AiThreadMessageLikeContentAnyOfInner
	isSet bool
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) Get() *AiThreadMessageLikeContentAnyOfInner {
	return v.value
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) Set(val *AiThreadMessageLikeContentAnyOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeContentAnyOfInner(val *AiThreadMessageLikeContentAnyOfInner) *NullableAiThreadMessageLikeContentAnyOfInner {
	return &NullableAiThreadMessageLikeContentAnyOfInner{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

