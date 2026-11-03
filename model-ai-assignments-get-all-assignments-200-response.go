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

// AiAssignmentsGetAllAssignments200Response - struct for AiAssignmentsGetAllAssignments200Response
type AiAssignmentsGetAllAssignments200Response struct {
	AiAssignmentsExpanded *AiAssignmentsExpanded
	MapmapOfStringstring *map[string]string
}

// AiAssignmentsExpandedAsAiAssignmentsGetAllAssignments200Response is a convenience function that returns AiAssignmentsExpanded wrapped in AiAssignmentsGetAllAssignments200Response
func AiAssignmentsExpandedAsAiAssignmentsGetAllAssignments200Response(v *AiAssignmentsExpanded) AiAssignmentsGetAllAssignments200Response {
	return AiAssignmentsGetAllAssignments200Response{
		AiAssignmentsExpanded: v,
	}
}

// map[string]stringAsAiAssignmentsGetAllAssignments200Response is a convenience function that returns map[string]string wrapped in AiAssignmentsGetAllAssignments200Response
func MapmapOfStringstringAsAiAssignmentsGetAllAssignments200Response(v *map[string]string) AiAssignmentsGetAllAssignments200Response {
	return AiAssignmentsGetAllAssignments200Response{
		MapmapOfStringstring: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AiAssignmentsGetAllAssignments200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AiAssignmentsExpanded
	err = newStrictDecoder(data).Decode(&dst.AiAssignmentsExpanded)
	if err == nil {
		jsonAiAssignmentsExpanded, _ := json.Marshal(dst.AiAssignmentsExpanded)
		if string(jsonAiAssignmentsExpanded) == "{}" { // empty struct
			dst.AiAssignmentsExpanded = nil
		} else {
			if err = validator.Validate(dst.AiAssignmentsExpanded); err != nil {
				dst.AiAssignmentsExpanded = nil
			} else {
				match++
			}
		}
	} else {
		dst.AiAssignmentsExpanded = nil
	}

	// try to unmarshal data into MapmapOfStringstring
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringstring)
	if err == nil {
		jsonMapmapOfStringstring, _ := json.Marshal(dst.MapmapOfStringstring)
		if string(jsonMapmapOfStringstring) == "{}" { // empty struct
			dst.MapmapOfStringstring = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringstring); err != nil {
				dst.MapmapOfStringstring = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringstring = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AiAssignmentsExpanded = nil
		dst.MapmapOfStringstring = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AiAssignmentsGetAllAssignments200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(AiAssignmentsGetAllAssignments200Response): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(AiAssignmentsGetAllAssignments200Response)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(AiAssignmentsGetAllAssignments200Response): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(AiAssignmentsGetAllAssignments200Response)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiAssignmentsGetAllAssignments200Response) MarshalJSON() ([]byte, error) {
	if src.AiAssignmentsExpanded != nil {
		return json.Marshal(&src.AiAssignmentsExpanded)
	}

	if src.MapmapOfStringstring != nil {
		return json.Marshal(&src.MapmapOfStringstring)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AiAssignmentsGetAllAssignments200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AiAssignmentsExpanded != nil {
		return obj.AiAssignmentsExpanded
	}

	if obj.MapmapOfStringstring != nil {
		return obj.MapmapOfStringstring
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AiAssignmentsGetAllAssignments200Response) GetActualInstanceValue() (interface{}) {
	if obj.AiAssignmentsExpanded != nil {
		return *obj.AiAssignmentsExpanded
	}

	if obj.MapmapOfStringstring != nil {
		return *obj.MapmapOfStringstring
	}

	// all schemas are nil
	return nil
}

type NullableAiAssignmentsGetAllAssignments200Response struct {
	value *AiAssignmentsGetAllAssignments200Response
	isSet bool
}

func (v NullableAiAssignmentsGetAllAssignments200Response) Get() *AiAssignmentsGetAllAssignments200Response {
	return v.value
}

func (v *NullableAiAssignmentsGetAllAssignments200Response) Set(val *AiAssignmentsGetAllAssignments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAssignmentsGetAllAssignments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAssignmentsGetAllAssignments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAssignmentsGetAllAssignments200Response(val *AiAssignmentsGetAllAssignments200Response) *NullableAiAssignmentsGetAllAssignments200Response {
	return &NullableAiAssignmentsGetAllAssignments200Response{value: val, isSet: true}
}

func (v NullableAiAssignmentsGetAllAssignments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAssignmentsGetAllAssignments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


