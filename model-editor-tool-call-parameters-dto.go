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
	"gopkg.in/validator.v2"
)

// EditorToolCallParametersDto - The editor tool call parameters.
type EditorToolCallParametersDto struct {
	GenerateDocxToolCallParametersDto *GenerateDocxToolCallParametersDto
	GenerateFormToolCallParametersDto *GenerateFormToolCallParametersDto
	GeneratePresentationToolCallParametersDto *GeneratePresentationToolCallParametersDto
}

// GenerateDocxToolCallParametersDtoAsEditorToolCallParametersDto is a convenience function that returns GenerateDocxToolCallParametersDto wrapped in EditorToolCallParametersDto
func GenerateDocxToolCallParametersDtoAsEditorToolCallParametersDto(v *GenerateDocxToolCallParametersDto) EditorToolCallParametersDto {
	return EditorToolCallParametersDto{
		GenerateDocxToolCallParametersDto: v,
	}
}

// GenerateFormToolCallParametersDtoAsEditorToolCallParametersDto is a convenience function that returns GenerateFormToolCallParametersDto wrapped in EditorToolCallParametersDto
func GenerateFormToolCallParametersDtoAsEditorToolCallParametersDto(v *GenerateFormToolCallParametersDto) EditorToolCallParametersDto {
	return EditorToolCallParametersDto{
		GenerateFormToolCallParametersDto: v,
	}
}

// GeneratePresentationToolCallParametersDtoAsEditorToolCallParametersDto is a convenience function that returns GeneratePresentationToolCallParametersDto wrapped in EditorToolCallParametersDto
func GeneratePresentationToolCallParametersDtoAsEditorToolCallParametersDto(v *GeneratePresentationToolCallParametersDto) EditorToolCallParametersDto {
	return EditorToolCallParametersDto{
		GeneratePresentationToolCallParametersDto: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EditorToolCallParametersDto) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GenerateDocxToolCallParametersDto
	err = newStrictDecoder(data).Decode(&dst.GenerateDocxToolCallParametersDto)
	if err == nil {
		jsonGenerateDocxToolCallParametersDto, _ := json.Marshal(dst.GenerateDocxToolCallParametersDto)
		if string(jsonGenerateDocxToolCallParametersDto) == "{}" { // empty struct
			dst.GenerateDocxToolCallParametersDto = nil
		} else {
			if err = validator.Validate(dst.GenerateDocxToolCallParametersDto); err != nil {
				dst.GenerateDocxToolCallParametersDto = nil
			} else {
				match++
			}
		}
	} else {
		dst.GenerateDocxToolCallParametersDto = nil
	}

	// try to unmarshal data into GenerateFormToolCallParametersDto
	err = newStrictDecoder(data).Decode(&dst.GenerateFormToolCallParametersDto)
	if err == nil {
		jsonGenerateFormToolCallParametersDto, _ := json.Marshal(dst.GenerateFormToolCallParametersDto)
		if string(jsonGenerateFormToolCallParametersDto) == "{}" { // empty struct
			dst.GenerateFormToolCallParametersDto = nil
		} else {
			if err = validator.Validate(dst.GenerateFormToolCallParametersDto); err != nil {
				dst.GenerateFormToolCallParametersDto = nil
			} else {
				match++
			}
		}
	} else {
		dst.GenerateFormToolCallParametersDto = nil
	}

	// try to unmarshal data into GeneratePresentationToolCallParametersDto
	err = newStrictDecoder(data).Decode(&dst.GeneratePresentationToolCallParametersDto)
	if err == nil {
		jsonGeneratePresentationToolCallParametersDto, _ := json.Marshal(dst.GeneratePresentationToolCallParametersDto)
		if string(jsonGeneratePresentationToolCallParametersDto) == "{}" { // empty struct
			dst.GeneratePresentationToolCallParametersDto = nil
		} else {
			if err = validator.Validate(dst.GeneratePresentationToolCallParametersDto); err != nil {
				dst.GeneratePresentationToolCallParametersDto = nil
			} else {
				match++
			}
		}
	} else {
		dst.GeneratePresentationToolCallParametersDto = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GenerateDocxToolCallParametersDto = nil
		dst.GenerateFormToolCallParametersDto = nil
		dst.GeneratePresentationToolCallParametersDto = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EditorToolCallParametersDto)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(EditorToolCallParametersDto)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EditorToolCallParametersDto) MarshalJSON() ([]byte, error) {
	if src.GenerateDocxToolCallParametersDto != nil {
		return json.Marshal(&src.GenerateDocxToolCallParametersDto)
	}

	if src.GenerateFormToolCallParametersDto != nil {
		return json.Marshal(&src.GenerateFormToolCallParametersDto)
	}

	if src.GeneratePresentationToolCallParametersDto != nil {
		return json.Marshal(&src.GeneratePresentationToolCallParametersDto)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EditorToolCallParametersDto) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GenerateDocxToolCallParametersDto != nil {
		return obj.GenerateDocxToolCallParametersDto
	}

	if obj.GenerateFormToolCallParametersDto != nil {
		return obj.GenerateFormToolCallParametersDto
	}

	if obj.GeneratePresentationToolCallParametersDto != nil {
		return obj.GeneratePresentationToolCallParametersDto
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EditorToolCallParametersDto) GetActualInstanceValue() (interface{}) {
	if obj.GenerateDocxToolCallParametersDto != nil {
		return *obj.GenerateDocxToolCallParametersDto
	}

	if obj.GenerateFormToolCallParametersDto != nil {
		return *obj.GenerateFormToolCallParametersDto
	}

	if obj.GeneratePresentationToolCallParametersDto != nil {
		return *obj.GeneratePresentationToolCallParametersDto
	}

	// all schemas are nil
	return nil
}

type NullableEditorToolCallParametersDto struct {
	value *EditorToolCallParametersDto
	isSet bool
}

func (v NullableEditorToolCallParametersDto) Get() *EditorToolCallParametersDto {
	return v.value
}

func (v *NullableEditorToolCallParametersDto) Set(val *EditorToolCallParametersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditorToolCallParametersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditorToolCallParametersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditorToolCallParametersDto(val *EditorToolCallParametersDto) *NullableEditorToolCallParametersDto {
	return &NullableEditorToolCallParametersDto{value: val, isSet: true}
}

func (v NullableEditorToolCallParametersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditorToolCallParametersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


