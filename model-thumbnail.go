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
	"fmt"
)

// Thumbnail [0 - Waiting, 1 - Created, 2 - Error, 3 - Not required, 4 - Creating]
type Thumbnail int32

// List of Thumbnail
const (
	THUMBNAIL_Waiting Thumbnail = 0
	THUMBNAIL_Created Thumbnail = 1
	THUMBNAIL_Error Thumbnail = 2
	THUMBNAIL_NotRequired Thumbnail = 3
	THUMBNAIL_Creating Thumbnail = 4
)

// All allowed values of Thumbnail enum
var AllowedThumbnailEnumValues = []Thumbnail{
	0,
	1,
	2,
	3,
	4,
}

func (v *Thumbnail) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Thumbnail(value)
	for _, existing := range AllowedThumbnailEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Thumbnail", value)
}

// NewThumbnailFromValue returns a pointer to a valid Thumbnail
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThumbnailFromValue(v int32) (*Thumbnail, error) {
	ev := Thumbnail(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Thumbnail: valid values are %v", v, AllowedThumbnailEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Thumbnail) IsValid() bool {
	for _, existing := range AllowedThumbnailEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Thumbnail value
func (v Thumbnail) Ptr() *Thumbnail {
	return &v
}

type NullableThumbnail struct {
	value *Thumbnail
	isSet bool
}

func (v NullableThumbnail) Get() *Thumbnail {
	return v.value
}

func (v *NullableThumbnail) Set(val *Thumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnail(val *Thumbnail) *NullableThumbnail {
	return &NullableThumbnail{value: val, isSet: true}
}

func (v NullableThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

