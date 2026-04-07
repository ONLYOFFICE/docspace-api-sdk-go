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

// checks if the FilesSettingsDtoInternalFormats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSettingsDtoInternalFormats{}

// FilesSettingsDtoInternalFormats The internal file formats.
type FilesSettingsDtoInternalFormats struct {
	Unknown *string `json:"Unknown,omitempty"`
	Archive *string `json:"Archive,omitempty"`
	Video *string `json:"Video,omitempty"`
	Audio *string `json:"Audio,omitempty"`
	Image *string `json:"Image,omitempty"`
	Spreadsheet *string `json:"Spreadsheet,omitempty"`
	Presentation *string `json:"Presentation,omitempty"`
	Document *string `json:"Document,omitempty"`
	Pdf *string `json:"Pdf,omitempty"`
	Diagram *string `json:"Diagram,omitempty"`
}

// NewFilesSettingsDtoInternalFormats instantiates a new FilesSettingsDtoInternalFormats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSettingsDtoInternalFormats() *FilesSettingsDtoInternalFormats {
	this := FilesSettingsDtoInternalFormats{}
	return &this
}

// NewFilesSettingsDtoInternalFormatsWithDefaults instantiates a new FilesSettingsDtoInternalFormats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSettingsDtoInternalFormatsWithDefaults() *FilesSettingsDtoInternalFormats {
	this := FilesSettingsDtoInternalFormats{}
	return &this
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetUnknown() string {
	if o == nil || IsNil(o.Unknown) {
		var ret string
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetUnknownOk() (*string, bool) {
	if o == nil || IsNil(o.Unknown) {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsUnknownSet() bool {
	if o != nil && !IsNil(o.Unknown) {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given string and assigns it to the Unknown field.
func (o *FilesSettingsDtoInternalFormats) SetUnknown(v string) {
	o.Unknown = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetArchive() string {
	if o == nil || IsNil(o.Archive) {
		var ret string
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetArchiveOk() (*string, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsArchiveSet() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given string and assigns it to the Archive field.
func (o *FilesSettingsDtoInternalFormats) SetArchive(v string) {
	o.Archive = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetVideo() string {
	if o == nil || IsNil(o.Video) {
		var ret string
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetVideoOk() (*string, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsVideoSet() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given string and assigns it to the Video field.
func (o *FilesSettingsDtoInternalFormats) SetVideo(v string) {
	o.Video = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetAudio() string {
	if o == nil || IsNil(o.Audio) {
		var ret string
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetAudioOk() (*string, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsAudioSet() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given string and assigns it to the Audio field.
func (o *FilesSettingsDtoInternalFormats) SetAudio(v string) {
	o.Audio = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsImageSet() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *FilesSettingsDtoInternalFormats) SetImage(v string) {
	o.Image = &v
}

// GetSpreadsheet returns the Spreadsheet field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetSpreadsheet() string {
	if o == nil || IsNil(o.Spreadsheet) {
		var ret string
		return ret
	}
	return *o.Spreadsheet
}

// GetSpreadsheetOk returns a tuple with the Spreadsheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetSpreadsheetOk() (*string, bool) {
	if o == nil || IsNil(o.Spreadsheet) {
		return nil, false
	}
	return o.Spreadsheet, true
}

// HasSpreadsheet returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsSpreadsheetSet() bool {
	if o != nil && !IsNil(o.Spreadsheet) {
		return true
	}

	return false
}

// SetSpreadsheet gets a reference to the given string and assigns it to the Spreadsheet field.
func (o *FilesSettingsDtoInternalFormats) SetSpreadsheet(v string) {
	o.Spreadsheet = &v
}

// GetPresentation returns the Presentation field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetPresentation() string {
	if o == nil || IsNil(o.Presentation) {
		var ret string
		return ret
	}
	return *o.Presentation
}

// GetPresentationOk returns a tuple with the Presentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetPresentationOk() (*string, bool) {
	if o == nil || IsNil(o.Presentation) {
		return nil, false
	}
	return o.Presentation, true
}

// HasPresentation returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsPresentationSet() bool {
	if o != nil && !IsNil(o.Presentation) {
		return true
	}

	return false
}

// SetPresentation gets a reference to the given string and assigns it to the Presentation field.
func (o *FilesSettingsDtoInternalFormats) SetPresentation(v string) {
	o.Presentation = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetDocument() string {
	if o == nil || IsNil(o.Document) {
		var ret string
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetDocumentOk() (*string, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsDocumentSet() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given string and assigns it to the Document field.
func (o *FilesSettingsDtoInternalFormats) SetDocument(v string) {
	o.Document = &v
}

// GetPdf returns the Pdf field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetPdf() string {
	if o == nil || IsNil(o.Pdf) {
		var ret string
		return ret
	}
	return *o.Pdf
}

// GetPdfOk returns a tuple with the Pdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetPdfOk() (*string, bool) {
	if o == nil || IsNil(o.Pdf) {
		return nil, false
	}
	return o.Pdf, true
}

// HasPdf returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsPdfSet() bool {
	if o != nil && !IsNil(o.Pdf) {
		return true
	}

	return false
}

// SetPdf gets a reference to the given string and assigns it to the Pdf field.
func (o *FilesSettingsDtoInternalFormats) SetPdf(v string) {
	o.Pdf = &v
}

// GetDiagram returns the Diagram field value if set, zero value otherwise.
func (o *FilesSettingsDtoInternalFormats) GetDiagram() string {
	if o == nil || IsNil(o.Diagram) {
		var ret string
		return ret
	}
	return *o.Diagram
}

// GetDiagramOk returns a tuple with the Diagram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDtoInternalFormats) GetDiagramOk() (*string, bool) {
	if o == nil || IsNil(o.Diagram) {
		return nil, false
	}
	return o.Diagram, true
}

// HasDiagram returns a boolean if a field has been set.
func (o *FilesSettingsDtoInternalFormats) IsDiagramSet() bool {
	if o != nil && !IsNil(o.Diagram) {
		return true
	}

	return false
}

// SetDiagram gets a reference to the given string and assigns it to the Diagram field.
func (o *FilesSettingsDtoInternalFormats) SetDiagram(v string) {
	o.Diagram = &v
}

func (o FilesSettingsDtoInternalFormats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSettingsDtoInternalFormats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unknown) {
		toSerialize["Unknown"] = o.Unknown
	}
	if !IsNil(o.Archive) {
		toSerialize["Archive"] = o.Archive
	}
	if !IsNil(o.Video) {
		toSerialize["Video"] = o.Video
	}
	if !IsNil(o.Audio) {
		toSerialize["Audio"] = o.Audio
	}
	if !IsNil(o.Image) {
		toSerialize["Image"] = o.Image
	}
	if !IsNil(o.Spreadsheet) {
		toSerialize["Spreadsheet"] = o.Spreadsheet
	}
	if !IsNil(o.Presentation) {
		toSerialize["Presentation"] = o.Presentation
	}
	if !IsNil(o.Document) {
		toSerialize["Document"] = o.Document
	}
	if !IsNil(o.Pdf) {
		toSerialize["Pdf"] = o.Pdf
	}
	if !IsNil(o.Diagram) {
		toSerialize["Diagram"] = o.Diagram
	}
	return toSerialize, nil
}

type NullableFilesSettingsDtoInternalFormats struct {
	value *FilesSettingsDtoInternalFormats
	isSet bool
}

func (v NullableFilesSettingsDtoInternalFormats) Get() *FilesSettingsDtoInternalFormats {
	return v.value
}

func (v *NullableFilesSettingsDtoInternalFormats) Set(val *FilesSettingsDtoInternalFormats) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSettingsDtoInternalFormats) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSettingsDtoInternalFormats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSettingsDtoInternalFormats(val *FilesSettingsDtoInternalFormats) *NullableFilesSettingsDtoInternalFormats {
	return &NullableFilesSettingsDtoInternalFormats{value: val, isSet: true}
}

func (v NullableFilesSettingsDtoInternalFormats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSettingsDtoInternalFormats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

