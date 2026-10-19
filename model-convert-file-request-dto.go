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

// checks if the ConvertFileRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertFileRequestDto{}

// ConvertFileRequestDto What to convert, where to keep the result, and the parameters the document service is handed. Everything below the  first two fields is the contract of the document service itself, named and shaped exactly as it names it, so that  the portal passes the request on rather than translating it. What the portal can read off the file - its format,  its name and the key of its revision - is not asked for and cannot be overridden.
type ConvertFileRequestDto struct {
	// The id of the portal file to convert, as reported by a folder listing such as `GET api/2.0/files/{folderId}`.
	FileId int32 `json:"fileId"`
	// The id of the folder to save the converted file in, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`.
	FolderId NullableInt32 `json:"folderId,omitempty"`
	// The encoding of a csv or txt source, as a code page number: 1251 for Cyrillic, 65001 for UTF-8, and so on.  Leaving it out lets the document service guess, which garbles non-latin text.
	CodePage NullableInt32 `json:"codePage,omitempty"`
	// The character separating the values of a csv source: 0 - none, 1 - tab, 2 - semicolon, 3 - colon, 4 - comma,  5 - space.
	Delimiter NullableInt32 `json:"delimiter,omitempty"`
	// The layout of forms printed as pdf documents or images.
	DocumentLayout *ConvertDocumentLayoutRequest `json:"documentLayout,omitempty"`
	// How the text of a pdf, xps or oxps source is read back.
	DocumentRenderer *ConvertDocumentRendererRequest `json:"documentRenderer,omitempty"`
	// The extension of the format to convert to, without the dot. The formats a source can become are listed under  `extsConvertible` in `GET api/2.0/files/settings`, keyed by the source extension, both with a leading dot there.
	Outputtype NullableString `json:"outputtype"`
	// The password of the source document, for a document that is protected with one.
	Password NullableString `json:"password,omitempty"`
	// The culture code whose currency, date and time formats a spreadsheet converted to PDF is shown in.
	Region NullableString `json:"region,omitempty"`
	// How a thumbnail is fitted, when the target format is an image.
	Thumbnail *ConvertThumbnailRequest `json:"thumbnail,omitempty"`
	// How a spreadsheet is laid out on the page, when it is converted to PDF.
	SpreadsheetLayout *ConvertSpreadsheetLayoutRequest `json:"spreadsheetLayout,omitempty"`
	// The PDF settings, used when the target format is PDF.
	Pdf *ConvertPdfRequest `json:"pdf,omitempty"`
}

type _ConvertFileRequestDto ConvertFileRequestDto

// NewConvertFileRequestDto instantiates a new ConvertFileRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertFileRequestDto(fileId int32, outputtype NullableString) *ConvertFileRequestDto {
	this := ConvertFileRequestDto{}
	this.FileId = fileId
	this.Outputtype = outputtype
	return &this
}

// NewConvertFileRequestDtoWithDefaults instantiates a new ConvertFileRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertFileRequestDtoWithDefaults() *ConvertFileRequestDto {
	this := ConvertFileRequestDto{}
	return &this
}

// GetFileId returns the FileId field value
func (o *ConvertFileRequestDto) GetFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *ConvertFileRequestDto) SetFileId(v int32) {
	o.FileId = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertFileRequestDto) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableInt32 and assigns it to the FolderId field.
func (o *ConvertFileRequestDto) SetFolderId(v int32) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ConvertFileRequestDto) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ConvertFileRequestDto) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetCodePage returns the CodePage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertFileRequestDto) GetCodePage() int32 {
	if o == nil || IsNil(o.CodePage.Get()) {
		var ret int32
		return ret
	}
	return *o.CodePage.Get()
}

// GetCodePageOk returns a tuple with the CodePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetCodePageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodePage.Get(), o.CodePage.IsSet()
}

// HasCodePage returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsCodePageSet() bool {
	if o != nil && o.CodePage.IsSet() {
		return true
	}

	return false
}

// SetCodePage gets a reference to the given NullableInt32 and assigns it to the CodePage field.
func (o *ConvertFileRequestDto) SetCodePage(v int32) {
	o.CodePage.Set(&v)
}
// SetCodePageNil sets the value for CodePage to be an explicit nil
func (o *ConvertFileRequestDto) SetCodePageNil() {
	o.CodePage.Set(nil)
}

// UnsetCodePage ensures that no value is present for CodePage, not even an explicit nil
func (o *ConvertFileRequestDto) UnsetCodePage() {
	o.CodePage.Unset()
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertFileRequestDto) GetDelimiter() int32 {
	if o == nil || IsNil(o.Delimiter.Get()) {
		var ret int32
		return ret
	}
	return *o.Delimiter.Get()
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetDelimiterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delimiter.Get(), o.Delimiter.IsSet()
}

// HasDelimiter returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsDelimiterSet() bool {
	if o != nil && o.Delimiter.IsSet() {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given NullableInt32 and assigns it to the Delimiter field.
func (o *ConvertFileRequestDto) SetDelimiter(v int32) {
	o.Delimiter.Set(&v)
}
// SetDelimiterNil sets the value for Delimiter to be an explicit nil
func (o *ConvertFileRequestDto) SetDelimiterNil() {
	o.Delimiter.Set(nil)
}

// UnsetDelimiter ensures that no value is present for Delimiter, not even an explicit nil
func (o *ConvertFileRequestDto) UnsetDelimiter() {
	o.Delimiter.Unset()
}

// GetDocumentLayout returns the DocumentLayout field value if set, zero value otherwise.
func (o *ConvertFileRequestDto) GetDocumentLayout() ConvertDocumentLayoutRequest {
	if o == nil || IsNil(o.DocumentLayout) {
		var ret ConvertDocumentLayoutRequest
		return ret
	}
	return *o.DocumentLayout
}

// GetDocumentLayoutOk returns a tuple with the DocumentLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetDocumentLayoutOk() (*ConvertDocumentLayoutRequest, bool) {
	if o == nil || IsNil(o.DocumentLayout) {
		return nil, false
	}
	return o.DocumentLayout, true
}

// HasDocumentLayout returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsDocumentLayoutSet() bool {
	if o != nil && !IsNil(o.DocumentLayout) {
		return true
	}

	return false
}

// SetDocumentLayout gets a reference to the given ConvertDocumentLayoutRequest and assigns it to the DocumentLayout field.
func (o *ConvertFileRequestDto) SetDocumentLayout(v ConvertDocumentLayoutRequest) {
	o.DocumentLayout = &v
}

// GetDocumentRenderer returns the DocumentRenderer field value if set, zero value otherwise.
func (o *ConvertFileRequestDto) GetDocumentRenderer() ConvertDocumentRendererRequest {
	if o == nil || IsNil(o.DocumentRenderer) {
		var ret ConvertDocumentRendererRequest
		return ret
	}
	return *o.DocumentRenderer
}

// GetDocumentRendererOk returns a tuple with the DocumentRenderer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetDocumentRendererOk() (*ConvertDocumentRendererRequest, bool) {
	if o == nil || IsNil(o.DocumentRenderer) {
		return nil, false
	}
	return o.DocumentRenderer, true
}

// HasDocumentRenderer returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsDocumentRendererSet() bool {
	if o != nil && !IsNil(o.DocumentRenderer) {
		return true
	}

	return false
}

// SetDocumentRenderer gets a reference to the given ConvertDocumentRendererRequest and assigns it to the DocumentRenderer field.
func (o *ConvertFileRequestDto) SetDocumentRenderer(v ConvertDocumentRendererRequest) {
	o.DocumentRenderer = &v
}

// GetOutputtype returns the Outputtype field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConvertFileRequestDto) GetOutputtype() string {
	if o == nil || o.Outputtype.Get() == nil {
		var ret string
		return ret
	}

	return *o.Outputtype.Get()
}

// GetOutputtypeOk returns a tuple with the Outputtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetOutputtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputtype.Get(), o.Outputtype.IsSet()
}

// SetOutputtype sets field value
func (o *ConvertFileRequestDto) SetOutputtype(v string) {
	o.Outputtype.Set(&v)
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertFileRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ConvertFileRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ConvertFileRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ConvertFileRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertFileRequestDto) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertFileRequestDto) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsRegionSet() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *ConvertFileRequestDto) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *ConvertFileRequestDto) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *ConvertFileRequestDto) UnsetRegion() {
	o.Region.Unset()
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *ConvertFileRequestDto) GetThumbnail() ConvertThumbnailRequest {
	if o == nil || IsNil(o.Thumbnail) {
		var ret ConvertThumbnailRequest
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetThumbnailOk() (*ConvertThumbnailRequest, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsThumbnailSet() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given ConvertThumbnailRequest and assigns it to the Thumbnail field.
func (o *ConvertFileRequestDto) SetThumbnail(v ConvertThumbnailRequest) {
	o.Thumbnail = &v
}

// GetSpreadsheetLayout returns the SpreadsheetLayout field value if set, zero value otherwise.
func (o *ConvertFileRequestDto) GetSpreadsheetLayout() ConvertSpreadsheetLayoutRequest {
	if o == nil || IsNil(o.SpreadsheetLayout) {
		var ret ConvertSpreadsheetLayoutRequest
		return ret
	}
	return *o.SpreadsheetLayout
}

// GetSpreadsheetLayoutOk returns a tuple with the SpreadsheetLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetSpreadsheetLayoutOk() (*ConvertSpreadsheetLayoutRequest, bool) {
	if o == nil || IsNil(o.SpreadsheetLayout) {
		return nil, false
	}
	return o.SpreadsheetLayout, true
}

// HasSpreadsheetLayout returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsSpreadsheetLayoutSet() bool {
	if o != nil && !IsNil(o.SpreadsheetLayout) {
		return true
	}

	return false
}

// SetSpreadsheetLayout gets a reference to the given ConvertSpreadsheetLayoutRequest and assigns it to the SpreadsheetLayout field.
func (o *ConvertFileRequestDto) SetSpreadsheetLayout(v ConvertSpreadsheetLayoutRequest) {
	o.SpreadsheetLayout = &v
}

// GetPdf returns the Pdf field value if set, zero value otherwise.
func (o *ConvertFileRequestDto) GetPdf() ConvertPdfRequest {
	if o == nil || IsNil(o.Pdf) {
		var ret ConvertPdfRequest
		return ret
	}
	return *o.Pdf
}

// GetPdfOk returns a tuple with the Pdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertFileRequestDto) GetPdfOk() (*ConvertPdfRequest, bool) {
	if o == nil || IsNil(o.Pdf) {
		return nil, false
	}
	return o.Pdf, true
}

// HasPdf returns a boolean if a field has been set.
func (o *ConvertFileRequestDto) IsPdfSet() bool {
	if o != nil && !IsNil(o.Pdf) {
		return true
	}

	return false
}

// SetPdf gets a reference to the given ConvertPdfRequest and assigns it to the Pdf field.
func (o *ConvertFileRequestDto) SetPdf(v ConvertPdfRequest) {
	o.Pdf = &v
}

func (o ConvertFileRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertFileRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileId"] = o.FileId
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.CodePage.IsSet() {
		toSerialize["codePage"] = o.CodePage.Get()
	}
	if o.Delimiter.IsSet() {
		toSerialize["delimiter"] = o.Delimiter.Get()
	}
	if !IsNil(o.DocumentLayout) {
		toSerialize["documentLayout"] = o.DocumentLayout
	}
	if !IsNil(o.DocumentRenderer) {
		toSerialize["documentRenderer"] = o.DocumentRenderer
	}
	toSerialize["outputtype"] = o.Outputtype.Get()
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.SpreadsheetLayout) {
		toSerialize["spreadsheetLayout"] = o.SpreadsheetLayout
	}
	if !IsNil(o.Pdf) {
		toSerialize["pdf"] = o.Pdf
	}
	return toSerialize, nil
}

func (o *ConvertFileRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileId",
		"outputtype",
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

	varConvertFileRequestDto := _ConvertFileRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConvertFileRequestDto)

	if err != nil {
		return err
	}

	*o = ConvertFileRequestDto(varConvertFileRequestDto)

	return err
}

type NullableConvertFileRequestDto struct {
	value *ConvertFileRequestDto
	isSet bool
}

func (v NullableConvertFileRequestDto) Get() *ConvertFileRequestDto {
	return v.value
}

func (v *NullableConvertFileRequestDto) Set(val *ConvertFileRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertFileRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertFileRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertFileRequestDto(val *ConvertFileRequestDto) *NullableConvertFileRequestDto {
	return &NullableConvertFileRequestDto{value: val, isSet: true}
}

func (v NullableConvertFileRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertFileRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

